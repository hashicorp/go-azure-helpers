# Versioning Guidelines

This document outlines the semantic versioning practices used in this repository.

## Semantic Versioning Overview

This repository follows [Semantic Versioning 2.0.0](https://semver.org/) (SemVer) principles. Version numbers follow the format X.Y.Z (Major.Minor.Patch):

- **Major version (X)**: Incremented for incompatible API changes
- **Minor version (Y)**: Incremented for backward-compatible functionality additions
- **Patch version (Z)**: Incremented for backward-compatible bug fixes

## Current Versioning Status

The project is currently in the pre-1.0 phase (0.x.y), which according to SemVer means the API is not considered stable and breaking changes could technically occur in minor versions. However, we tend to follow conservative versioning even in this development phase.

## Version Change Types

### Major Version Changes (X.y.z)

While the repository hasn't yet had a major version bump (still in 0.x.x series), the following would typically constitute a major version change:

- Breaking API changes such as:
  - Renaming or removing public functions, types, or methods
  - Changing function signatures in incompatible ways
  - Modifying the behavior of existing functions in ways that would break dependent code
  - Significant restructuring of the package organization
  - Changes to the minimum required Go version that would force users to upgrade

**Future Major Release Scenarios**:
- Complete refactoring of resource ID handling paradigm
- Redesigning authentication interfaces
- Removing deprecated functionality that was marked for removal in previous versions
- Changing the minimum Go version requirement to a much newer release

### Minor Version Changes (x.Y.z)

Minor version increases are used for new functionality added in a backward-compatible manner.

**Real Examples**:
- v0.70.0: Added new functionality to recaser for better resource ID type handling
- v0.71.0: Added logic app ID functionality for differentiation in app service
- v0.69.0: Added new common IDs
- v0.67.0: Added new schema functions

**Common Patterns**:
- Adding new resource types
- Adding new helper functions
- Implementing new authentication methods
- Adding new schema validation functions
- Extending existing functionality in backward-compatible ways

### Patch Version Changes (x.y.Z)

Patch versions are incremented for backward-compatible bug fixes and minor improvements.

**Real Examples**:
- v0.71.1: Dependency updates and maintenance (`go get -u -t ./... && go mod tidy`)
- v0.70.1: Adding schema function for resourceid string elements (enhancement to existing functionality)

**Common Patterns**:
- Bug fixes
- Documentation updates
- Performance improvements that don't affect the API
- Minor dependency updates
- Non-breaking enhancements to existing features

## Versioning Process

When contributing to this repository, consider the following:

1. If your change breaks existing functionality, it likely requires a major version bump (or careful consideration if we're pre-1.0)
2. If your change adds new functionality without breaking existing code, it requires a minor version bump
3. If your change fixes bugs or makes minor improvements without changing the API, it requires a patch version bump

## Pre-1.0 Considerations

While in the 0.x.y phase:

- The API is still considered under development
- According to strict SemVer, breaking changes could occur in minor version increments
- However, our practice shows we try to maintain stability even in this phase
- Breaking changes are generally avoided when possible, even in this early development stage

## Transitioning to 1.0.0

When the project reaches v1.0.0, it will signal:

- API stability
- A commitment to maintain backward compatibility until the next major version
- A more strict application of SemVer principles

After reaching 1.0.0, any breaking changes will require incrementing the major version number.
