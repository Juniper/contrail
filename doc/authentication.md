# Authentication

API Server support OpenStack Keystone v3 authentiation.
Keystone supports various backend such as LDAP etc.

Configuraion Example

```
# Keystone configuraion
keystone:
    authurl: http://localhost:5000/v3
```

```
# Keystone configuraion
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