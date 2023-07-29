# v1.1.0 (Fri Jul 30 2021)

- GetLatestTag will now return a stable answer if there are multiple, equivalent
  tags in the repository (e.g. `refs/tags/v1.0.0` and `refs/tags/1.0.0`). Previously
  the returned tag would be randomly selected.
