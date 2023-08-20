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
	return LatestTagIgnoringPrefix(url, "", options...)
}

// LatestTagIgnoringPrefix behaves the same as [LatestTag], but ignores the
// given prefix when trying to parse tags as semver. This can be useful if
// releases are tagged with "release-" or similar.
func LatestTagIgnoringPrefix(url string, prefix string, options ...Option) (string, string, error) {
	refs, err := Fetch(url, append(options, TagsOnly())...)
	if err != nil {
		return "", "", err
	}

	return latestTag(refs, prefix)
}

func latestTag(refs map[string]string, prefix string) (string, string, error) {
	best := version.Must(version.NewVersion("0.0.0"))
	bestTag := ""
	bestHash := ""
	for r := range refs {
		tag := strings.TrimPrefix(r, tagPrefix)
		v, err := version.NewVersion(strings.TrimPrefix(tag, prefix))
		if err == nil && v.Prerelease() == "" {
			if v.GreaterThan(best) || (v.Equal(best) && strings.Compare(v.Original(), best.Original()) < 0) {
				best = v
				bestTag = tag
				bestHash = refs[r]
			}
		}
	}

	if bestTag == "" {
		return "", "", fmt.Errorf("no tags found")
	}

	return bestTag, bestHash, nil
}
