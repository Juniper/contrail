# Authentication

API Server support OpenStack Keystone v3 authentication.
Keystone supports various backend such as LDAP etc.

Configuration Example

```yaml
# Keystone configuration
keystone:
    authurl: http://localhost:5000/v3
    service_user:
        id: goapi
        password: goapi
        project_name: service
        domain_id: default
```

```yaml
# Keystone configuration
keystone:
    local: true
    assignment:
        type: static
        file: ./test_data/keystone.yml
    store:
        type: memory
        expire: 3600
    authurl: http://localhost:9091/v3
```
