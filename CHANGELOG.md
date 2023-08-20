# v1.2.0 (Sun Aug 20 2023)

- Added `GetLatestTagIgnoringPrefix`, which will ignore a user-defined prefix 
  when parsing the semver from a tag. e.g. given a prefix of "release-", a ref
  of `refs/tags/release-v1.0.0` will parse as semver `1.0.0`.

# v1.1.1 (Fri Jul 30 2023)

- Removed debug print statement from GetLatestTag.

# v1.1.0 (Fri Jul 30 2023)

- GetLatestTag will now return a stable answer if there are multiple, equivalent
  tags in the repository (e.g. `refs/tags/v1.0.0` and `refs/tags/1.0.0`). Previously
  the returned tag would be randomly selected.
