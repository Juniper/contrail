# Services in the Service Chain

Common source code: [common.go](../pkg/services/common.go)

The chain overview:

    ⇄ REST API / gRPC
    ↓
    ↓ ContrailService
    ↓ RefUpdateToUpdateService
    ↓ SanitizerService
    ↓ RBACService
    ↓ QuotaCheckerService
    ↓ etcd.NotifierService (optional)
    ↓ NeutronService (optional)
    ↓ ContrailTypeLogicService
    ↓ db.DBService
    ↓
    ↺ Database

## Contrail service

- **Caller:** REST/gRPC framework
- **Purpose:** Operations on request's payload: common operations, schema-based validation, deserialize payload to resource structure.
- **Source code:** [service.go.tmpl](../pkg/services/service.go.tmpl) [service_common.go.tmpl](../pkg/services/service_common.go.tmpl)

Contrail Service is registered as API request handler.

## Ref update to update service

- **Caller:** ContrailService
- **Purpose:** Translate reference update to in-transaction resource update.
- **Source code:** [service_interface.go.tmpl](../pkg/services/service_interface.go.tmpl)

RefUpdate is a special endpoint which is used only for add and delete references.
It is risky to make such changes outside of transaction. RefUpdateToUpdate
translates add/delete reference to in-transaction resource update.

## Sanitizer service

- **Caller:** RefUpdateToUpdateService
- **Purpose:** Fills up missing properties based on resources logic and metadata.
- **Source code:** [sanitizer_service.go.tmpl](../pkg/services/sanitizer_service.go.tmpl) [sanitizer.go](../pkg/services/sanitizer.go)

Sanitizer complement properties like: refs or display name by creating or updating resources.

## RBAC service

- **Caller:** Sanitizer service
- **Purpose:** Check whether resource access is allowed based on RBAC configuration.
- **Source code:** [rbac_service.go.tmpl](../pkg/services/rbac_service.go.tmpl) [rbac.go](../pkg/services/rbac.go)

RBAC does role based access control on resource operations. If a user has not any role which will allow a particular
operation, RBAC service won't allow the user to do that resource operation.

## Contrail type logic service

- **Caller:** RBAC service
- **Purpose:** Implements business logic specific to each type (model).
- **Source code:** [service.go](../pkg/types/service.go)

Here lives business logic specific for each type.

## Quota checker service

- **Caller:** ContrailTypeLogicService
- **Purpose:** Checks if the resource's quantity has been exceeded.
- **Source code:** [base_quota_getter.go.tmpl](../pkg/services/base_quota_getter.go.tmpl), [base_quota_counter.go.tmpl](../pkg/services/base_quota_counter.go.tmpl), [quota_checker_service.go.tmpl](../pkg/services/quota_checker_service.go.tmpl)

Quota is a maximum limit for creation new resources.

Quota checker service is composed of two parts: quota limit getter and counter.
Limit getter implements quota limit retrieval and counter implements counting logic.

## etcd notifier service

- **Caller:** QuotaCheckerService
- **Purpose:** Notification bus for etcd.
- **Source code:** [etcdserviceif.go.tmpl](../pkg/db/etcd/etcdserviceif.go.tmpl)

Notifier service uses etcd server for pushing change notification.
Other Contrail microservices can observe etcd events and react on changing resources.

Notifier is optional and can be disabled in config yaml file by setting `server.notify_etcd: false`.
etcd notifier is a temporary substitute for Sync service.

## Database service

- **Caller:** QuotaCheckerService or ContrailTypeLogicService
- **Purpose:** High level abstraction for database driver.
- **Source code:** [db.go](../pkg/db/db.go)

Database service provides query builder by exposing high level abstraction methods. Works with PostgreSQL driver.

Database service can be accessed through Read and Write services located
in ContrailService and ContrailTypeLogicService.

## Other services

Services which play significant role in the project but are not part of the Service Chain.

### Cache Service

- **Caller:** Intent compiler
- **Source code:** [intent cache.go](../pkg/compilation/intent/cache.go)

Cache is heavily used in Intent Compiler (outside of Service Chain).

Current implementation of cache is responsible for:

- Storing and removing objects from cache
- Updating object and its references in cache
