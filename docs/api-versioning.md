# k8sd API Versioning Principles

## Overview

The k8sd API is the core interface for managing Canonical Kubernetes clusters. This document outlines the versioning strategy and guidelines for managing major, minor, and patch versions of the API.

## Versioning Strategy

The k8sd API uses a **branch-based versioning approach** following Go module best practices:

1. **The `main` branch** contains the latest major version (currently v2)
2. **Release branches** (e.g. `v1`) for older versions are created when a new major version is introduced
3. **Git tags** track minor and patch versions in the format `vX.Y.Z`
4. **The `go.mod` file** includes the major version in the module path for v2+

### Version Format

The complete version follows semantic versioning: `vX.Y.Z`

- `X` = Major version (reflected in branch name and module path)
- `Y` = Minor version (tracked via Git tags)
- `Z` = Patch version (tracked via Git tags)

**Example tags:**

- `v1.0.0` - Initial release of v1 API (on `v1` branch)
- `v1.1.0` - v1 API with backward-compatible additions (on `v1` branch)
- `v1.1.1` - v1 API with bug fixes (on `v1` branch)
- `v2.0.0` - Initial release of v2 API (on `main` branch)
- `v2.1.0` - v2 API with backward-compatible additions (on `main` branch)

### How It Works

1. **Active Development on `main`**: The `main` branch always contains the latest major version (e.g., v2)
2. **API Files in `api/` Directory**: All API types and definitions live directly in `api/*.go` files
3. **Module Path with Version**: The `go.mod` declares the module as `github.com/canonical/k8s-snap-api/v2`
4. **Branch for Each Major Version**: When v3 is needed, the current `main` (v2) becomes the `v2` branch
5. **Backports for Older Versions**: Bug fixes and features can be backported from `main` to older version branches

Note that this structure was adopted for `v2`. The `v1` branch has a slightly different structure.
See [ADR 001: Adopt Branch-Based API Versioning] for details.

## When to Create a New Major Version (New Branch)

A new major version **MUST** be created when making **breaking changes** to the API. Breaking changes include:

### 1. Removing Fields or Endpoints

- Removing fields from request or response structures
- Removing RPC endpoints

### 2. Changing Field Types or Semantics

- Changing the data type of an existing field (e.g., `string` → `int`)
- Changing the meaning or behavior of a field
- Changing validation rules that would reject previously valid values

### 3. Renaming Fields or Endpoints

- Renaming API fields (JSON/YAML keys)
- Renaming RPC endpoints
- Changing field serialization names

### 4. Changing Default Behaviors

- Modifying default values in ways that change cluster behavior

### 5. Removing or Changing Enum Values

- Removing allowed values from fields with restricted sets
- Example: Removing `k8s-dqlite` from the allowed `DatastoreType` values

### 6. Making Required Fields Optional or Vice Versa

- Changes that affect whether fields must be provided

**When a breaking change is introduced:**

1. **Create a release branch** for the current major version:

   ```bash
   # If currently on main with v2
   git checkout main
   git checkout -b v2
   git push origin v2
   ```

2. **Update `go.mod` on `main`** to the next major version:

   ```go
   module github.com/canonical/k8s-snap-api/v3
   ```

   Note that `v1` does not contain a suffix.

3. **Make breaking changes** in `api/*.go` files on `main`

4. **Tag the first release** of the new version:

   ```bash
   git tag -a v3.0.0 -m "Release v3.0.0: Breaking changes description"
   git push origin v3.0.0
   ```

5. **Update import paths** in consuming code:

   ```go
   // Old
   import "github.com/canonical/k8s-snap-api/v2/api"

   // New
   import "github.com/canonical/k8s-snap-api/v3/api"
   ```

## When to Create a New Minor Version

Minor versions are tracked via Git tags **without creating a new branch**. A minor version bump (e.g., `v2.0.0` → `v2.1.0`) is appropriate for backward-compatible changes that add functionality:

### 1. Adding New Optional Fields

- Adding new fields to structs that have default values
- New fields must be optional and must not change existing behavior when omitted
- Files are modified on the appropriate branch (e.g., `main` for v2, or `v1` branch for v1)

### 2. Adding New Endpoints

- Introducing new RPC methods while keeping existing ones unchanged
- New `rpc_*.go` files are added to the `api/` directory

### 3. Adding New Enum Values

- Adding new allowed values to existing fields with restricted sets
- Existing values continue to work as before

### 4. Adding New Types

- Introducing new types that don't modify existing types
- New `type_*.go` files are added to the `api/` directory

**When adding backward-compatible features:**

1. Modify files in `api/` on the appropriate branch
2. Commit the changes
3. Tag the release: `v{N}.{Y+1}.0`
4. Push the tag

## When to Create a New Patch Version

Patch versions are tracked via Git tags **without creating a new branch**. A patch version bump (e.g., `v2.1.0` → `v2.1.1`) is appropriate for backward-compatible bug fixes:

### 1. Documentation Fixes

- Correcting comments, documentation, or examples
- Clarifying field descriptions

### 3. Code Quality

- Repository changes that don't change the API contract

**When fixing bugs:**

1. Fix the bug in `api/` files on the appropriate branch
2. Commit the fix
3. Tag the release: `v{N}.{Y}.{Z+1}`
4. Push the tag


<!-- Links -->

[ADR 001: Adopt Branch-Based API Versioning]: (./adr/001-branch-based-versioning.md)
