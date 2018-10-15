# Authentication

API Server support OpenStack Keystone v3 authentication.
Keystone supports various backend such as LDAP etc.

Configuration Example

```yaml
# Keystone configuration
keystone:
    authurl: http://localhost:5000/v3
```

```yaml
# Keystone configuration
keystone:
    local: true # Enable local keystone v3. This is only for testing now.
    assignment:
        type: static
        file: ./test_data/keystone.yml
    store:
        type: memory
        expire: 3600
    authurl: http://localhost:9091/v3
```
