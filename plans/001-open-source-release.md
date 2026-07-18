# #001: Prepare repository for open-source release
**Type**: chore
**Module**: (root)
**Status**: in-progress
**GitHub**: https://github.com/mainvec/ugo/issues/1
**Branch**: chore/1-open-source-release

## Progress

- [x] GitHub issue created (#1)
- [x] Plan file written (`plans/001-open-source-release.md`)
- [x] Registry updated (`plans/registry.json`)
- [x] Branch created (`chore/1-open-source-release`)
- [x] T1: Audit public-release readiness
- [x] T2: Document the project and contribution process
- [x] T3: Add continuous integration and repository metadata
- [x] Tests passing (`go test ./...`)
- [ ] PR opened (`Closes #1`)

---

## Problem / Goal

The repository is public and MIT-licensed, but it lacks sufficient project documentation,
contributor and security guidance, and automated checks for a dependable open-source release.

## Approach

Document the supported Go version, installation, package scope, stability expectations, and
contribution workflow. Add focused GitHub community files and a CI workflow that enforces
formatting, vetting, and tests while preserving the standard-library-only dependency policy.

---

## Tasks

### T1: Audit public-release readiness
**Status**: done
**Notes**: Confirmed the repository is public, GitHub detects the MIT license, all tests and
`go vet` pass, no third-party modules are required, and no common credential markers were found.
The existing deletion of obsolete `go.sum` entries is user work and will be preserved.

### T2: Document the project and contribution process
**Status**: done
**Notes**: Replaced the placeholder README; added contribution, conduct, and security guidance;
and preserved the upstream MIT license for the adapted CLI package in a third-party notice.

### T3: Add continuous integration and repository metadata
**Status**: done
**Notes**: Added pinned GitHub Actions checks, Dependabot configuration, issue forms, and a pull
request template. Updated the live repository description, homepage, and topics, and enabled
private vulnerability reporting. All 26 tests, formatting checks, and `go vet` pass.