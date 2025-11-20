# Changelog

## 1.6.0 - 2025-11-20

- Added `WithAuth` option to pass basic auth credentials.

## 1.5.0 - 2025-11-06

- Filter out dereferenced tags (thanks @Greboid)

## 1.4.1 - 2025-08-25

- Reduced minimum go version to 1.16

## 1.4.0 - 2025-07-08

- Minor dependency update.

## 1.3.0 - 2024-12-17

- Added `WithContext` option.
- Minor dependency update.

## 1.2.0 - 2023-08-20

- Added `GetLatestTagIgnoringPrefix`, which will ignore a user-defined prefix 
  when parsing the semver from a tag. e.g. given a prefix of "release-", a ref
  of `refs/tags/release-v1.0.0` will parse as semver `1.0.0`.

## 1.1.1 - 2023-07-30

- Removed debug print statement from GetLatestTag.

## 1.1.0 - 2023-07-30

- GetLatestTag will now return a stable answer if there are multiple, equivalent
  tags in the repository (e.g. `refs/tags/v1.0.0` and `refs/tags/1.0.0`). Previously
  the returned tag would be randomly selected.
