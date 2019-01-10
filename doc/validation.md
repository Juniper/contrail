# Validation
## Schema Based Validation

Schema based validation is done in `ContrailService` which is the first one in the inner service chain.
Validation function for each resource is generated from [template file](tools/templates/contrail/type_validation.tmpl).
Each generated function may be overwriten or extended in [this file](pkg/models/validation.go).

### Enums

Each property may have `enum` parameter defined which allows only specified values.

### Format

Each property may have `format` parameter defined which assures value in proper format is being stored.
Definitions of format are located in [source code](pkg/models/basemodels/validation.go).
New formats must be added to the same file.

```yaml
definitions:

```

##### Supported formats:
- date-time
- ipv4
- mac
- hostname
- base64
- positive_int_as_string
- service_interface_type_format

## Schema Example

```yaml
definitions:
  ActionType:
    type: string
    enum:
    - reject
    - accept
    - next
  IpamSubnetType:
    created:
      description: timestamp when subnet object gets created
      format: date-time
      presence: optional
      type: string
```

## Complex type specific validation

More complex type specific validation is done in `ContrailTypeLogicService`.
This type of validation must be added to appropriate files in [this folder](pkg/types).

