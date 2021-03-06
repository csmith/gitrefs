package gitrefs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-version"
)

// LatestTag attempts to find the highest semver tag from the repository at the given url.
// Returns the tag and the commit hash of that tag, or an error if no semver tags are found.
// Tags that can't be parsed as semver are silently ignored.
func LatestTag(url string, options ...Option) (string, string, error) {
	refs, err := Fetch(url, append(options, TagsOnly())...)
	if err != nil {
		return "", "", err
	}

	best := version.Must(version.NewVersion("0.0.0"))
	bestTag := ""
	bestHash := ""
	for r := range refs {
		tag := strings.TrimPrefix(r, tagPrefix)
		v, err := version.NewVersion(tag)
		if err == nil && v.GreaterThanOrEqual(best) && v.Prerelease() == "" {
			best = v
			bestTag = tag
			bestHash = refs[r]
		}
	}

	if bestTag == "" {
		return "", "", fmt.Errorf("no tags found")
	}

	return bestTag, bestHash, nil
}
