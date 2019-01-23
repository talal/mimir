# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v1.2.0 - 2019-01-23
### Added
- Highlight inaccessible directory path in red.

## v1.1.2 - 2019-01-10
### Fixed
- Home path not being converted to `~` in some cases.

## v1.1.1 - 2018-12-07
### Fixed
- "slice bounds out of range" error that occurred when the `pwd` was more than
  6 directories deep and began with a `/`.

## v1.1.0 - 2018-11-09
### Added
- Kubernetes context and OpenStack cloud info can be turned off by specifying
  the respective flags to 'false': `MIMIR_KUBE` and `MIMIR_OS_CLOUD`.

### Changed
- If `CURRENT_OS_CLOUD` env variable is not available then the OpenStack cloud
  info is shown using the standard OpenStack environment variables. The scope
  of the info depends on the environment variables that are available.

## v1.0.0 - 2018-11-02

Initial release.
