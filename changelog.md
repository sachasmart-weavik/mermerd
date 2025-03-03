# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html) (after version 0.0.5).

## [0.4.1] - 2022-09-28
### Fixed
- Fix wrong column format for `is_primary` ([Issue #24](https://github.com/KarnerTh/mermerd/issues/24))

## [0.4.0] - 2022-09-11
### Added
- Support variable expansion ([Issue #16](https://github.com/KarnerTh/mermerd/issues/16))
- Added support for attribute keys ([Issue #20](https://github.com/KarnerTh/mermerd/issues/20))

### Changed
- mermerd will now by default add the attribute key if applicable (PK or FK). If this is undesired, it can be
  disabled by the `--omitAttributeKeys` flag or the `omitAttributeKeys` config (example is in the readme).

## [0.3.0] - 2022-09-02
### Added
- Added --selectedTables switch ([PR #12](https://github.com/KarnerTh/mermerd/pull/12))
- Added support for MSSQL ([Issue #13](https://github.com/KarnerTh/mermerd/issues/13))

### Changed
- go 1.19 is now used
- updated dependencies

### Fixed
- Fixed some typos and documentation

## [0.2.1] - 2022-06-03
### Fixed
- Embed the template file into the binary ([Issue #10](https://github.com/KarnerTh/mermerd/issues/10))

## [0.2.0] - 2022-06-01
### Added
- A `--debug` flag/config to show debug information
- A `--omitConstraintLabels` flag/config to toggle the new constraint labels

### Changed
- The column name is now displayed as the constraint label (can be switched off)

### Fixed
- Sub query for constraints returned multiple items  ([Issue #8](https://github.com/KarnerTh/mermerd/issues/8))

## [0.1.0] - 2022-04-15
### Added
- Mermerd is available via the go tools

### Changed
- go 1.18 is now used

### Fixed
- MySQL query fix for constraints ([Issue #7](https://github.com/KarnerTh/mermerd/issues/7))

## [0.0.5] - 2022-03-17
### Added
- New config: allow surrounding output with mermerd backticks ([PR #4](https://github.com/KarnerTh/mermerd/pull/4))

## [0.0.4] - 2022-03-14
### Added
- Licence

### Fixed
- Do not require a global configuration file

## [0.0.3] - 2022-03-12
### Added
- Possibility to opt in for all tables
- Start mermerd with a predefined run config
- Add version command
- Show version number in intro header

### Changed
- Improved help command output
- Exit with error code 1 on failure
- Fully POSIX-compliant flags (including short & long versions)
- the parameter for the connection string suggestions (previously `connectionStrings`) was renamed to
  `connectionStringSuggestions`
- the flag `-ac` was replaced with `--showAllConstraints`

### Removed
- `.mermerd` configuration file is not automatically created on first use anymore

## [0.0.2] - 2022-01-30
### Added
- Configurable suggestions for connection string input

### Changed
- improved one to many constraint detection for mysql
- improved one to many constraint detection for postgres

## [0.0.1] - 2022-01-17
### Added
- Initial release of mermerd

[0.4.1]: https://github.com/KarnerTh/mermerd/releases/tag/v0.4.1

[0.4.0]: https://github.com/KarnerTh/mermerd/releases/tag/v0.4.0

[0.3.0]: https://github.com/KarnerTh/mermerd/releases/tag/v0.3.0

[0.2.1]: https://github.com/KarnerTh/mermerd/releases/tag/v0.2.1

[0.2.0]: https://github.com/KarnerTh/mermerd/releases/tag/v0.2.0

[0.1.0]: https://github.com/KarnerTh/mermerd/releases/tag/v0.1.0

[0.0.5]: https://github.com/KarnerTh/mermerd/releases/tag/v0.0.5

[0.0.4]: https://github.com/KarnerTh/mermerd/releases/tag/v0.0.4

[0.0.3]: https://github.com/KarnerTh/mermerd/releases/tag/v0.0.3

[0.0.2]: https://github.com/KarnerTh/mermerd/releases/tag/v0.0.2

[0.0.1]: https://github.com/KarnerTh/mermerd/releases/tag/v0.0.1
