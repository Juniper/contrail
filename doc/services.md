# Services

Sourcecode:

- Structs [common.go](../pkg/services/common.go)
- REST endpoints [service.tmpl](../tools/templates/service.tmpl)
- Interfaces [service_interface.tmpl](../tools/templates/service_interface.tmpl)

## The chain

Each service on top, can call next service and get its returned value.
Services does not act as middleware interceptors. Each service have to be called explicitly
and returned value does not propagate to the next service, but it goes back to the caller.
After whole round trip response will be send back to the client.

```
⇄ RESTAPI / gRPC
↓
↓ ContrailService
↓ RefUpdateToUpdateService
↓ SanitizerService
↓ ContrailTypeLogicService
↓ QuotaCheckerService
↓ NotifierService (optional)
↓ DBService
↓
↺ Database
```

### Contrail service

- **Caller:** REST/gRPC framework
- **Purpose:** Operations on request payload, schema validation, bind payload to proper structure.

Contrail Service is registered as API request handler.

Has access to:
- Type validator
- In transaction doer
- Metadata getter
- Int pool allocator

### Ref update to update service

- **Caller:** ContrailService
- **Purpose:** Translate reference update to in-transaction resource update.

RefUpdate is a special endpoint which is used only for add and delete tag references.
It is risky to make such changes outside of transaction. RefUpdateToUpdate
translates add/delete ref to in-transaction resource update.

Has access to:
- In transaction doer
- Read service

### Sanitizer service

- **Caller:** RefUpdateToUpdateService
- **Purpose:** Fills up missing properties based on resources logic and metadata.
- **Sourcecode:** [sanitizer_service.tmpl](../tools/templates/sanitizer_service.tmpl)

Sanitizer fills gaps in properties like: refs or display name by creating or updating resources.

Has access to:
- Metadata getter

### Contrail specific type logic service

- **Caller:** SanitizerService
- **Purpose:** Implements business logic specific to each type (model).
- **Sourcecode:** [type logic](../pkg/types)

Contrail type logic is the most important part of the service chain.
Here lives business logic specific for each specific type.
Type logic contains preparing model from request object, specific validation,
interactions with other objects and finaly actions on database.

Has access to:
- In transaction doer
- Metadata getter
- Read service
- Write service
- Address manager
- Int pool allocator

### Quota checker service

- **Caller:** ContrailTypeLogicService
- **Purpose:** Checks if the resource limit has been exceeded.
- **Sourcecode:** [quota getter](../tools/templates/base_quota_getter.tmpl), [quota counter](../tools/templates/base_quota_counter.tmpl), [quota checker](../tools/templates/quota_checker_service.tmpl)

Quota is a resource maximum limit which can be created for a tenant or a project.

Quota checker service is compose of two parts: quota limit getter and counter.
Limit getter implements quota limit retrieval and counter implements counting logic.

### Notifier service

- **Caller:** QuotaCheckerService
- **Purpose:** Notification bus for other microservices.
- **Sourcecode:** [notifier](../tools/templates/etcdserviceif.tmpl)

Notifier service uses etcd server for pushing change notification.
Other Contrail microservices can observe etcd events and react on changing resources.

Notifier is optional and can be disabled in config yaml file by setting `notify_etcd: false`.

### Database service

- **Caller:** QuotaCheckerService or ContrailTypeLogicService
- **Purpose:** High level abstraction for database driver.
- **Sourcecode:** [base db](../pkg/db/basedb)

Database service provides query builder by exposing high level abstraction methods.
Works with MySQL and PostgreSQL drivers.

Database service can be accessed through Read and Write services located
in ContrailService and ContrailTypeLogicService.

## Other services

Services which play significant role in project but are not part of the service chain.

### Cache Service

- **Caller:** Intent compiler
- **Sourcecode:** [Intent Cache](../pkg/compilation/intent/cache.go)

Cache is heavily used in Intent Compiler (outside of service chain).

Current implementation of cache is responsible for:
- Store, update and remove objects
- Update references when child objects have changed
