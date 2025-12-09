# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.26.0] - 2025-12-09
### Added
- Added `requirement` field to dependency response
- Added `StatusResponse.DB` nested message with schema version and creation timestamp fields
- Added `StatusResponse.Server` nested message with server version information
- Added optional `db` field to `StatusResponse` for database version information
- Added `server` field to `StatusResponse` for general server information

## [0.25.0] - 2025-10-13
### Added
- Added `error_message` and `error_code` fields to all cryptography component response messages for enhanced error handling
- Added `ErrorCode` enum with standardized error codes: `INVALID_PURL`, `COMPONENT_NOT_FOUND`, `NO_INFO`, `INVALID_SEMVER`
- Added error handling documentation section to cryptography API README.md
- Added error response examples to cryptography API documentation demonstrating error field usage

### Changed
- Enhanced cryptography API documentation with comprehensive error handling examples and field descriptions
- Updated JSON schema examples in `ComponentsAlgorithmsResponse` and `ComponentAlgorithmsResponse` to demonstrate error scenarios
- Updated cryptography README.md with error codes reference table and usage notes

## [0.24.0] - 2025-09-24
### Added
- Added gRPC `GetComponentsIssues` and REST endpoint POST `/v2/semgrep/issues/components` for Semgrep security analysis
- Added gRPC `GetComponentIssues` and REST endpoint GET `/v2/semgrep/issues/component` for single component Semgrep analysis
- Added new response message types `ComponentsIssueResponse` and `ComponentIssueResponse` for enhanced component handling
- Added `ComponentIssueInfo` message with component-specific security analysis information
- Added `requirement` field to Semgrep component information for version requirement specifications
- Added JSON schema examples to Semgrep response messages for improved API documentation

### Changed
- Enhanced Semgrep protobuf definitions with comprehensive service and message documentation
- Refactored Semgrep protobuf structure by extracting `Issue` and `File` messages as top-level messages
- Updated OpenAPI schema with realistic JSON response examples for Semgrep endpoints

### Deprecated
- Deprecated gRPC `GetIssues` method (use `GetComponentsIssues` instead)
- Deprecated `SemgrepResponse` message type (use `ComponentsIssueResponse` or `ComponentIssueResponse` instead)

## [0.23.0] - 2025-09-22
### Added
- Added gRPC `GetCountryContributorsByComponents` and REST endpoint POST `/v2/geoprovenance/countries/components`
- Added gRPC `GetCountryContributorsByComponent` and REST endpoint GET `/v2/geoprovenance/countries/component`
- Added gRPC `GetOriginByComponents` and REST endpoint POST `/v2/geoprovenance/origin/components`
- Added gRPC `GetOriginByComponent` and REST endpoint GET `/v2/geoprovenance/origin/component`
- Added comprehensive documentation to geo-provenance protobuf service
- Added geo-provenance API documentation (README.md)
- Added JSON schema examples to geo-provenance response messages
- Added new response message types `ComponentsContributorResponse` and `ComponentsOriginResponse` for enhanced component handling
### Changed
- Enhanced geo-provenance protobuf definitions with comprehensive service and message documentation
- Updated OpenAPI schema with realistic JSON response examples for geo-provenance endpoints
- Enhanced field documentation across all geo-provenance message types
### Deprecated
- Deprecated gRPC `GetComponentContributors` method (use `GetCountryContributorsByComponents` instead)
- Deprecated gRPC `GetComponentOrigin` method (use `GetOriginByComponents` instead)
- Deprecated `ContributorResponse` and `OriginResponse` message types (use new component-based response types instead)

## [0.22.0] - 2025-09-22

## [0.21.0] - 2025-09-18
### Added
- Added README.md documentation for Components Service API v2
- Added JSON schema examples to components protobuf definitions for improved API documentation
### Changed
- Updated components protobuf message structure with proper field naming consistency
- Enhanced components response messages with JSON schema examples for OpenAPI generation

## [0.20.1] - 2025-09-15
### Fixed
- Fixed `TransitiveDependencyRequest` protobuf definition

## [0.20.0] - 2025-09-15
### Added
- Added dependency service documentation
### Changed
- **BREAKING CHANGE**: Changed `/v2/dependencies/transitive` to `/v2/dependencies/transitive/components`
- Deprecated `/v2/dependencies/dependencies` endpoint (use `/v2/licenses/components` instead)

## [0.19.0] - 2025-09-08
### Added
- Added gRPC `GetComponentAlgorithms` and REST endpoint GET `/v2/cryptography/algorithms/component`
- Added gRPC `GetComponentsAlgorithms` and REST endpoint POST `/v2/cryptography/algorithms/components`
- Added gRPC `GetComponentAlgorithmsInRange` and REST endpoint GET `/v2/cryptography/algorithms/range/component`
- Added gRPC `GetComponentsAlgorithmsInRange` and REST endpoint POST `/v2/cryptography/algorithms/range/components`
- Added gRPC `ComponentVersionsInRange` and REST endpoint GET `/v2/cryptography/algorithms/versions/range/component`
- Added gRPC `ComponentsVersionsInRange` and REST endpoint POST `/v2/cryptography/algorithms/versions/range/components`
- Added gRPC `ComponentHintsInRange` and REST endpoint GET `/v2/cryptography/hints/component`
- Added gRPC `ComponentsHintsInRange` and REST endpoint POST `/v2/cryptography/hints/components`
- Added comprehensive documentation to cryptography protobuf service
- Added cryptography API documentation (README.md)
- Added JSON schema examples to all cryptography response messages
- Added `requirement` field to `ComponentAlgorithms` and `ComponentHints` messages for consistency with vulnerability service
- Enhanced cryptography service endpoints with detailed method documentation
### Changed  
- Improved cryptography protobuf definitions with comprehensive service and message documentation
- Updated OpenAPI schema with realistic JSON response examples
- Enhanced field documentation across all cryptography message types
### Fixed
- Fixed GitHub URL typo in cryptography service contact information

## [0.18.0] - 2025-09-04
### Changed
- Removed `/api` prefix from REST endpoints

## [0.17.0] - 2025-08-29
### Added
- Added `json_name` to protobuf to keep compatibility between REST and gRPC protocol


## [0.16.0] - 2025-08-29
### Added
- Enhanced swagger documentation with examples for dependencies protobuf
- Added request/response examples to transitive dependencies API
### Changed
- Updated REST endpoint paths: removed `/api/` prefix from vulnerability and dependency endpoints (now `/v2/...` instead of `/api/v2/...`)
- Changed Component Echo endpoints from GET to POST

## [0.15.0] - 2025-08-26
### Added
- Added gRPC `GetComponentCpes` and REST endpoint GET `/api/v2/vulnerabilities/cpes/component`
- Added gRPC `GetComponentsCpes` and REST endpoint POST `/api/v2/vulnerabilities/cpes/components`
- Added gRPC `GetComponentVulnerabilities` and REST endpoint GET `/api/v2/vulnerabilities/component`
- Added gRPC `GetComponentsVulnerabilities` and REST endpoint POST `/api/v2/vulnerabilities/components`
- Added vulnerability API documentation

## [0.14.0] - 2025-08-18
### Added
- Added version to component response

## [0.13.0] - 2025-08-13
### Added
- Added License API v2 service

### Changed
- Deprecated PurlRequest, replaced by ComponentsRequest

## [0.12.0] - 2025-07-11
### Changed
- gRPC-Gateway v2 migration completed
- Upgraded Python base image

### Fixed
- Replace protoc-gen-swagger with protoc-gen-openapiv2 for Python

## [0.11.0] - 2025-07-02
### Changed
- Downgraded protoc versions to support python 3.8

## [0.10.0] - 2025-07-01
### Added
- Added protobuf dependency management
- Added protoc-gen-openapiv2 support

## [0.9.0] - 2025-07-01
### Added
- Generated go API definitions

## [0.8.0] - 2025-06-30
### Added
- Added Docker build support
- Enhanced folder hashing endpoints

## [0.7.2] - 2025-04-25
### Added
- Updated API paths

## [0.7.1] - 2025-04-25
### Added
- Updated geoprovenance API paths

## [0.7.0] - 2025-04-24
### Added
- Added Geo Provenance endpoints

## [0.6.0] - 2025-04-01
### Added
- Added folder hashing endpoint

## [0.5.0] - 2024-12-04
### Added
- Extended crypto endpoints
  - library/sdk hints

## [0.4.0] - 2024-10-15
### Added
- Added crypto range endpoints

## [0.3.0] - 2024-08-23
### Added
- Initial crypto endpoint

## [0.2.0] - 2023-10-27
### Added
- Added the following endpoints:
  - Semgrep

## [0.1.0] - 2023-03-08
### Added
- Added the following endpoints:
  - Cryptography
  - Components
  - Dependencies
  - Vulnerabilities
- Added REST endpoint support for each service also

[Unreleased]: https://github.com/scanoss/papi/compare/v0.12.0...HEAD
[0.26.0]: https://github.com/scanoss/papi/compare/v0.25.0...v0.26.0
[0.25.0]: https://github.com/scanoss/papi/compare/v0.24.0...v0.25.0
[0.24.0]: https://github.com/scanoss/papi/compare/v0.23.0...v0.24.0
[0.23.0]: https://github.com/scanoss/papi/compare/v0.22.0...v0.23.0
[0.22.0]: https://github.com/scanoss/papi/compare/v0.21.0...v0.22.0
[0.21.0]: https://github.com/scanoss/papi/compare/v0.20.1...v0.21.0
[0.20.1]: https://github.com/scanoss/papi/compare/v0.20.0...v0.20.1
[0.20.0]: https://github.com/scanoss/papi/compare/v0.19.0...v0.20.0
[0.19.0]: https://github.com/scanoss/papi/compare/v0.18.0...v0.19.0
[0.18.0]: https://github.com/scanoss/papi/compare/v0.17.0...v0.18.0
[0.17.0]: https://github.com/scanoss/papi/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/scanoss/papi/compare/v0.15.0...v0.16.0
[0.15.0]: https://github.com/scanoss/papi/compare/v0.14.0...v0.15.0
[0.14.0]: https://github.com/scanoss/papi/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/scanoss/papi/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/scanoss/papi/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/scanoss/papi/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/scanoss/papi/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/scanoss/papi/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/scanoss/papi/compare/v0.7.2...v0.8.0
[0.7.2]: https://github.com/scanoss/papi/compare/v0.7.1...v0.7.2
[0.7.1]: https://github.com/scanoss/papi/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/scanoss/papi/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/scanoss/papi/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/scanoss/papi/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/scanoss/papi/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/scanoss/papi/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/scanoss/papi/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/scanoss/papi/compare/v0.0.1...v0.1.0
