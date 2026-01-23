# ADR 001: Adopt Branch-Based API Versioning

* Status: Accepted
* Owner: @bschimke95

## Purpose

This ADR establishes the versioning strategy for the k8sd API repository. Specifically, we need to decide how to organize and manage multiple major API versions in the codebase, whether to use separate directories for each version (like Kubernetes) or separate Git branches (like Go modules and microcluster).

**The decision**: We will adopt **branch-based versioning** where the `main` branch contains the latest major version, and older major versions are maintained on separate branches (e.g., `v1`). API files will live directly in `api/*.go`, and the `go.mod` will include the version in the module path for v2+ (e.g., `github.com/canonical/k8s-snap-api/v2`).

## Context

The k8sd API repository previously used a **directory-based versioning** approach where each major version lived in separate directories (`api/v1/`, `api/v2/`) within the same repository. This approach aligns with the Kubernetes API versioning model but does not quite fit the k8sd use case.

### Problems with Directory-Based Versioning

#### Kubernetes-Style API Groups Don't Apply

Kubernetes has the concept of API groups, which allows different parts of their API to evolve independently. This makes sense for their directory-based structure (`apps/v1`, `networking.k8s.io/v1`, etc.).

However, k8sd does not have API groups. This means:

* Either the full API needs to be copied to a new major version directory
* Or awkward dependencies between `v{N}` and `v{N+1}` arise if types are reused

#### Tagging Challenges

When a type shared between versions needs updating, we face a dilemma:

* Update the `v{N}` struct definition and create a `v{N+1}` tag (but this reflects the change in v2 as well, which is confusing)
* Copy the entire structure to `v{N+1}` (treating it as a breaking change even if it isn't)

#### Not Idiomatic for Go Modules

The Go module system expects v2+ modules to include the version in the module path:

* Directory-based creates awkward imports: `github.com/canonical/k8s-snap-api/api/v2`
* The relationship between directories and module versions is unclear
* Import paths become unnecessarily verbose

#### Complexity for Small Teams

Given k8sd's size and contributor count, there's generally no need for complex feature lifecycle management through alpha → beta → stable as done in upstream Kubernetes. A simpler versioning scheme is more appropriate.

The [microcluster](https://github.com/canonical/microcluster) project successfully uses branch-based versioning, which is the idiomatic Go versioning scheme. So, this approach has proven effective for similar projects and team-sizes within Canonical.

## Decision

We will adopt **branch-based versioning** following Go module best practices and the microcluster model.

### New Versioning Strategy

**Branch-Based Versioning**:

1. **`main` branch** contains the latest major version (currently v2)
2. **Version branches** (e.g. `v1`) are created to preserve older major versions
3. **API files** live directly in `api/*.go` (not in versioned subdirectories)
4. **The `go.mod`** includes major version in module path for v2+: `github.com/canonical/k8s-snap-api/v2`
5. **Breaking changes** trigger creation of a new version branch and module path update on `main`

### How It Works

**When creating a new major version (e.g., v2 → v3):**

1. Create a `v2` branch from current `main`
2. Keep developing v3 on `main`
3. Update `go.mod` on `main` to `github.com/canonical/k8s-snap-api/v3`
4. Tag the first v3 release as `v3.0.0`

**Backporting to older versions:**

* Bug fixes: Develop on `main` first, then cherry-pick to older version branches
* Security fixes: Must be backported to all supported versions
* Features: Generally only on `main`, backport only if critical for older versions
* Breaking changes: Never backport; create a new major version

**Import paths:**

* v1: `github.com/canonical/k8s-snap-api` (on `v1` branch)
* v2: `github.com/canonical/k8s-snap-api/v2` (on `main` branch)
* v3: `github.com/canonical/k8s-snap-api/v3` (on `main` branch, when created)

## Alternatives Considered

### Alternative 1: Keep Directory-Based Versioning (api/v1/, api/v2/)

**Description**: Continue with the previous approach of separate directories for each version.

**Rejected Because**:

* Not idiomatic for Go modules (v2+ should have version in module path)
* Creates awkward import paths: `github.com/canonical/k8s-snap-api/api/v2`
* Harder to share code between versions (leads to duplication or awkward dependencies)
* The relationship between directories and module versions is unclear
* Tagging becomes confusing when shared types are updated

### Alternative 2: Monorepo with Separate Modules

**Description**: Use separate `go.mod` files for each version in subdirectories, creating multiple independent modules within one repository.

**Rejected Because**:

* More complex to manage multiple modules in one repository
* Tag management becomes more complex (need to prefix tags with subdirectory paths)
* Not necessary for k8sd's use case where versions are sequential, not parallel development
* Adds unnecessary complexity for a small team

### Alternative 3: Separate Repository Per Version

**Description**: Create entirely separate repositories for each major version (k8s-snap-api, k8s-snap-api-v2, etc.).

**Rejected Because**:

* Creates organizational overhead in managing multiple repositories
* Makes cross-version code review and comparison difficult
* Fragments the project's history and documentation
* Overkill for the scale and team size of k8sd

## Consequences

### Advantages

1. **Idiomatic Go**: Follows Go module versioning best practices with version in module path for v2+
2. **Simpler structure**: No nested version directories, just `api/*.go` files on each branch
3. **Clear evolution**: `main` is always the latest version, older versions preserved on branches

### Trade-offs

1. **Branch management required**: Need to create and maintain version branches when introducing breaking changes

### Migration from Directory-Based to Branch-Based

The migration involves:

1. Create `v1` branch preserving the old `api/v1/*` structure
2. Move `api/v2/*` files to `api/*` on `main`
3. Update `go.mod` on `main` to `github.com/canonical/k8s-snap-api/v2`
4. Tag `v2.0.0` on `main`
5. Update client libraries to import `/v2` path

## References

* [k8sd API Versioning Principles](../api-versioning.md)
* [Microcluster API Versioning](https://github.com/canonical/microcluster)
* [Go Modules Reference - Version Numbers](https://go.dev/ref/mod#versions)
* [Semantic Versioning 2.0.0](https://semver.org/)
