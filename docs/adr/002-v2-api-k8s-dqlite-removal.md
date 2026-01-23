# ADR 002: Create v2 API Without k8s-dqlite Datastore

* Status: Accepted
* Owner: @bschimke95

## Purpose

This ADR establishes whether we should create a new major API version to remove k8s-dqlite support from the k8sd API. Specifically, we need to decide how to handle the deprecation and removal of k8s-dqlite fields.

**The decision**: We will create a new v2 API that removes all k8s-dqlite related fields and configuration options.

## Context

Canonical Kubernetes has historically supported two datastore options for storing cluster state:

1. **k8s-dqlite**: A custom distributed database based on dqlite (a distributed SQLite implementation)
2. **etcd**: The standard Kubernetes datastore used by most Kubernetes distributions

### Deprecation Decision

After careful consideration, the decision was made to deprecate k8s-dqlite in Canonical Kubernetes 1.35 with full removal planned for version 1.36.

### API Impact

The current v1 API includes several fields related to k8s-dqlite configuration:

* `BootstrapConfig.K8sDqlitePort`: Port configuration for k8s-dqlite
* `BootstrapConfig.DatastoreType`: Accepts "k8s-dqlite", "etcd", or "external"
* `BootstrapConfig.ExtraNodeK8sDqliteArgs`: Additional arguments for k8s-dqlite
* `ControlPlaneJoinConfig.ExtraNodeK8sDqliteArgs`: Additional arguments for k8s-dqlite

Removing these fields from the existing v1 API would be a breaking change that violates API compatibility guarantees as outlined in the [k8sd API versioning principles](../api-versioning.md).

## Decision

We will **create a new v2 API** that removes all k8s-dqlite related fields and configuration options, while maintaining the v1 API for a transition period in a separate branch.

### v1 API Changes

* Mark `K8sDqlitePort` from `BootstrapConfig` as deprecated.
* Mark `ExtraNodeK8sDqliteArgs` from `BootstrapConfig` and `ControlPlaneJoinConfig` as deprecated.
* Add deprecation notice for the `k8s-dqlite` option to the `DatabaseType` config

### v2 API Changes

1. Add deprecation notice to k8s-dqlite fields
2. Cut `v1` branch
3. Remove `K8sDqlitePort` from `BootstrapConfig`
4. Remove `ExtraNodeK8sDqliteArgs` from `BootstrapConfig` and `ControlPlaneJoinConfig`
5. Change allowed values from "k8s-dqlite | etcd | external" to "etcd | external" for the `DatabaseType` config
6. Create `v2.0.0` tag

## Alternatives Considered

### Alternative 1: Keep k8s-dqlite Fields but Mark as Deprecated

**Description**: Keep all k8s-dqlite fields in v1 API with deprecation warnings and ignore them in the implementation.

**Rejected Because**:

* Creates confusion for new users who might not notice deprecation warnings
* Maintains technical debt in the API surface area
* Could lead to bugs if users attempt to use deprecated fields expecting them to work
* Doesn't provide a clear signal about the direction of the product
* Violates the principle of removing unused code

### Alternative 2: Remove Fields from v1 API (Breaking Change)

**Description**: Remove k8s-dqlite fields from v1 API in a breaking manner without creating a new version.

**Rejected Because**:

* Violates API stability guarantees and semantic versioning principles

## Consequences

### Advantages

1. **Cleaner API**: Removes deprecated functionality, making the API easier to understand for new users

### Trade-offs

1. **Migration effort required**: Clients must update imports and remove k8s-dqlite configuration from their code

## References

* [k8sd API Versioning Principles](../api-versioning.md)
* [ADR 001: Adopt Branch-Based API Versioning](./001-branch-based-versioning.md)
