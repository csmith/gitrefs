package gitrefs

import "testing"

const (
	commit1 = "a1b2c3d4e5f6"
	commit2 = "f6e5d4c3b2a1"
	commit3 = "1a2b3c4d5e6f"
)

func Test_latestTag(t *testing.T) {
	tests := []struct {
		name     string
		refs     map[string]string
		prefix   string
		wantTag  string
		wantHash string
		wantErr  bool
	}{
		{"No refs", map[string]string{}, "", "", "", true},
		{"No tags", map[string]string{
			"refs/heads/master": commit1,
			"refs/heads/v1.0.0": commit2,
		}, "", "", "", true},
		{"Single tag", map[string]string{
			"refs/heads/master": commit1,
			"refs/heads/v1.0.0": commit2,
			"refs/tags/v1.0.0":  commit3,
		}, "", "v1.0.0", commit3, false},
		{"Multiple tags", map[string]string{
			"refs/tags/v0.9.9": commit1,
			"refs/tags/v1.0.0": commit2,
			"refs/tags/v1.0.1": commit3,
		}, "", "v1.0.1", commit3, false},
		{"Multiple tags with prefix", map[string]string{
			"refs/tags/release-0.9.9": commit1,
			"refs/tags/release-1.0.0": commit2,
			"refs/tags/release-1.0.1": commit3,
		}, "release-", "release-1.0.1", commit3, false},
		{"Multiple tags some with prefix", map[string]string{
			"refs/tags/v0.9.9":        commit1,
			"refs/tags/v1.0.0":        commit2,
			"refs/tags/release-1.0.1": commit3,
		}, "release-", "release-1.0.1", commit3, false},
		{"Pre-release tags", map[string]string{
			"refs/tags/v0.9.9":        commit1,
			"refs/tags/v1.0.0-rc1":    commit2,
			"refs/tags/v1.0.0-alpha7": commit3,
		}, "", "v0.9.9", commit1, false},
		{"Equal tags", map[string]string{
			"refs/tags/v1.0.0": commit1,
			"refs/tags/1.0.0":  commit2,
		}, "", "1.0.0", commit2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, hash, err := latestTag(tt.refs, tt.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("latestTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tag != tt.wantTag {
				t.Errorf("latestTag() tag = %v, wantTag %v", tag, tt.wantTag)
			}
			if hash != tt.wantHash {
				t.Errorf("latestTag() hash = %v, wantTag %v", hash, tt.wantHash)
			}
		})
	}
}
