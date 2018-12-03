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
↓ DBService
↓
↺ Database
```

### Contrail service

- **Caller:** REST/gRPC framework
- **Purpose:** Operations on request payload, schema validation, bind payload to proper structure.

Contrail Service is registered as API request handler.

Has access to:
- Metadata getter
- Type validator
- In transaction doer
- Int pool allocator

### Ref update to update service

- **Caller:** ContrailService
- **Purpose:** Translate reference update to in-transaction resource update.

RefUpdate is a special endpoint which is used only for add and delete tag references.
It is risky to make such changes outside of transaction. RefUpdateToUpdate
translates add/delete ref to in-transaction resource update.

Has access to:
- Read service
- In transaction doer

### Sanitizer service

- **Caller:** RefUpdateToUpdateService
- **Purpose:** Fills up missing properties based on resources logic and metadata.



### Contrail specific type logic service
### Quota checker service
### Database service

## Other services

### Cache Service
