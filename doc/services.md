# Services

## Service chain

Services are arranged in a chain, where service on the left can call service on the right.
Services does not act as middleware interceptors. They have to be called explicitly and returned
value does not propagate to the next service, but it goes back to the caller on the left.

### The chain

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
- **Purpose:** Operations on request data


### Ref update to update service
### Sanitizer service
### Contrail specific type logic service
### Quota checker service
### Database service

## Other services

### Cache Service
