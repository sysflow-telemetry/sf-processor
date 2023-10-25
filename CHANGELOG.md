# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/).

> **Types of changes:**
>
> - **Added**: for new features.
> - **Changed**: for changes in existing functionality.
> - **Deprecated**: for soon-to-be removed features.
> - **Removed**: for now removed features.
> - **Fixed**: for any bug fixes.
> - **Security**: in case of vulnerabilities.

## [Unreleased]

## [0.6.0] - 2023-09-30

### Added

- Experimental Sigma rules support
- Multi-language rules engine support

### Changed

- Bumped UBI version to 9.2-755
- Bump sf-apis to 0.6.0

### Security

- CVE-2022-41723: golang.org/x/net Uncontrolled Resource Consumption (updated to 0.7.0)
- CVE-2022-27664: golang.org/x/net/http2 Denial of Service vulnerability (updated to 0.0.0-20220906165146-f3363e06e74c)
- CVE-2022-32149: Denial of service in golang.org/x/text/language (updated to 0.3.8)
- CVE-2022-41721: golang.org/x/net/http2/h2c vulnerable to request smuggling attack (updated to 0.1.1-0.20221104162952-702349b0e862)
- CVE-2022-28948: gopkg.in/yaml.v3 Denial of Service (updated to 3.0.0-20220521103104-8f96da9f5d5e)

## [0.5.1] - 2023-06-07

### Added

- Add multi-driver support

### Changed

- Bumped UBI version to 8.8-854
- Bump sf-apis to 0.5.1

### Fixed

- Fix off-by-1 JSON ports encoding
- Add correct formatting to mapPortList in JSON output

## [0.5.0] - 2022-10-17

### Added

- Add support for k8s pod and event objects
- Add jsonpath expression support for policy engine

### Changed

- Bumped UBI version to 8.6-943.1665521450

### Fixed

- Fix bug in exists predicate
- Fix `open_read` and `open_write` macros in ttps.yaml

## [0.4.4] - 2022-08-01

### Added

- Add rate limiting filter with time decaying

### Changed

- Bump UBI to 8.6-855
- Update reference to sf-apis

### Fixed

- Fix exists predicate
- Fix handling of integers and booleans in MatStr function

## [0.4.3] - 2022-06-21

### Changed

- Update systemd service to include plugindir argument

## [0.4.2] - 2022-06-10

### Changed

- Add missing host field to ECS encoder

## [0.4.1] - 2022-05-26

### Changed

- Bumped UBI version to 8.6-754
- Removed binary package's dkms requirement

## [0.4.0] - 2022-02-18

### Added

- Support for pluggable actions for policy engine
- Support for asynchonous policy engine with thread pooling
- Packaging in deb, rpm, and targz formats
- Added 14 new MITRE TTP tagging rules
- Added support for quiet logging mode
- Added plugin builder image to support plugin development and releases

### Changed

- Added contextual sysflow structure, removed global cache and cache synchronization primitives; refactored handler interface
- Changed cache keys to OID types
- BREAKING Changed policy engine modes and action verbs (update policy yaml rule declarations to remove `action` attribute if used with `alert` or `tag` verbs)
  - `alert` and `enrich` are now policy engine modes, and `action` in policy rule declaration is now used for calling action handling plugins
- Updated the short union strings from gogen-avro
- Updated CI to automate packaging or release assets with release notes
- Bump go version to go1.17.7
- BREAKING Added support for architecture-dependent build (darwin, linux), due to [changes in go 1.17 net](https://github.com/golang/go/commit/e97d8eb027c0067f757860b6f766644de15941f2) package
- Updated findings short description formatting and name convention

### Fixed

- Fixed cache coherence and race condition when updating the cache in the processor plugin; splits the processor plugin into two plugins, reader (which builds the cache) and processor (only reads from cache)
- Fixed stream socket reader issue introduced with the upgrade to go 1.17

### Security

- Updated IBM Findings SDK to fix [CVE-2020-26160](https://github.com/advisories/GHSA-w73w-5m7g-f7qc)

## [0.3.1] - 2021-09-29

### Changed

- Bumped UBI version to 8.4-211.

## [0.3.0] - 2021-09-20

### Added

- Support for pluggable export protocols
- Elastic Common Schema (ECS) export format and Elasticsearch integration
- Export to IBM Findings API
- MITRE ATT&CK ttp tagging policy
- Support for pipeline forking (tee feature)
- Custom S3 prefix to Findings exporter

### Changed

- Moved away from Dockerhub CI.
- Optimized JSON export
- Updated dependencies to latest `sf-apis`
- Updated sample policies
- Refactoring of processor and handling APIs

### Fixed

- Fixes bugs in policy engine related to lists containing quoted strings
- Fixes several issues in policy engine field mapping

### Removed

- Support for flat JSON schema

## [0.2.2] - 2020-12-07

### Changed

- Updated dependencies to latest `sf-apis`.

## [0.2.1] - 2020-12-02

### Fixed

- Fixes `sf.file.oid` and `sf.file.newoid` attribute mapping.

## [0.2.0] - 2020-12-01

### Added

- Adds lists and macro preprocessing to deal with usage before declarations in input policy language.
- Adds empty handling for process flow objects.
- Adds `endswith` binary operator to policy expression language.
- Added initial documentation.

### Changed

- Updates the grammar and intepreter to support falco policies.
- Several refactorings and performance optimizations in policy engine.
- Tuned filter policy for k8s clusters.

### Fixed

- Fixes module names and package paths.

## [0.1.0] - 2020-10-30

### Added

- First release of SysFlow Processor.

[Unreleased]: https://github.com/sysflow-telemetry/sf-processor/compare/0.6.0...HEAD
[0.6.0]: https://github.com/sysflow-telemetry/sf-processor/compare/0.5.1...0.6.0
[0.5.1]: https://github.com/sysflow-telemetry/sf-processor/compare/0.5.0...0.5.1
[0.5.0]: https://github.com/sysflow-telemetry/sf-processor/compare/0.4.4...0.5.0
[0.4.4]: https://github.com/sysflow-telemetry/sf-processor/compare/0.4.3...0.4.4
[0.4.3]: https://github.com/sysflow-telemetry/sf-processor/compare/0.4.2...0.4.3
[0.4.2]: https://github.com/sysflow-telemetry/sf-processor/compare/0.4.1...0.4.2
[0.4.1]: https://github.com/sysflow-telemetry/sf-processor/compare/0.4.0...0.4.1
[0.4.0]: https://github.com/sysflow-telemetry/sf-processor/compare/0.3.1...0.4.0
[0.3.1]: https://github.com/sysflow-telemetry/sf-processor/compare/0.2.2...0.3.1
[0.3.0]: https://github.com/sysflow-telemetry/sf-processor/compare/0.2.2...0.3.0
[0.2.2]: https://github.com/sysflow-telemetry/sf-processor/compare/0.2.1...0.2.2
[0.2.1]: https://github.com/sysflow-telemetry/sf-processor/compare/0.2.0...0.2.1
[0.2.0]: https://github.com/sysflow-telemetry/sf-processor/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/sysflow-telemetry/sf-processor/releases/tag/0.1.0
