# REST API

## Introduction

API Server provides REST API endpoints.
For each resource type, the following APIs are available:

- Create a resource
- Update a resource
- Delete a resource given its UUID
- Read a resource given its UUID
- List resources of given type

In addition, the following APIs are also available:

- Listing all resource types
- Convert FQ name to UUID
- Convert UUID to FQ name
- Add/Delete/Update a reference between two objects
- Relax a reference between two objects

## OpenAPI

We have OpenAPI version 2.0 definitions in public directory.
You can also take a look detail REST API spec using OpenAPI document generator.

Example.

```bash
npm install -g spectacle-docs
spectacle -d public/openapi.json
```

## Creating a resource

To create a resource, a POST has to be issued on the collection URL. So for a resource of type example-resource,

```text
METHOD: POST
URL: http://<ip>:<port>/example_resources/
BODY: JSON representation of example-resource type
RESPONSE: UUID and href of created resource
```

Example request

```bash
curl -X POST -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" -d '{"virtual-network": {"parent_type": "project", "fq_name": ["default-domain", "admin", "vn-blue"], "network_ipam_refs": [{"attr": {"ipam_subnets": [{"subnet": {"ip_prefix": "10.1.1.0", "ip_prefix_len": 24}}]}, "to": ["default-domain", "default-project", "default-network-ipam"]}]}}' http://10.84.14.2:8082/virtual-networks
```

Response

```json
{
    "virtual-network": {
        "fq_name": [
            "default-domain",
            "admin",
            "vn-blue"
        ],
        "href": "http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee",
        "name": "vn-blue",
        "parent_href": "http://10.84.14.2:8082/project/df7649a6-3e2c-4982-b0c3-4b5038eef587",
        "parent_uuid": "df7649a6-3e2c-4982-b0c3-4b5038eef587",
        "uuid": "8c84ff8a-30ac-4136-99d9-f0d9662f3eee"
    }
}
```

## Reading a resource

To read a resource, a GET has to be issued on the resource URL.

```text
METHOD: GET
URL: http://<ip>:<port>/example_resource/<example-resource-uuid>
BODY: None
RESPONSE: JSON representation of the resource
```

Example Request

```bash
curl -X GET -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee
```

Response

```json
{
    "virtual-network": {
        "access_control_lists": [
            {
                "href": "http://10.84.14.2:8082/access-control-list/24b9c337-7be8-4883-a9a0-60197edf64e4",
                "to": [
                    "default-domain",
                    "admin",
                    "vn-blue",
                    "vn-blue"
                ],
                "uuid": "24b9c337-7be8-4883-a9a0-60197edf64e4"
            }
        ],
        "fq_name": [
            "default-domain",
            "admin",
            "vn-blue"
        ],
        "href": "http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee",
        "id_perms": {
            "created": "2013-09-13T00:26:05.290644",
            "description": null,
            "enable": true,
            "last_modified": "2013-09-13T00:47:41.219833",
            "permissions": {
                "group": "cloud-admin-group",
                "group_access": 7,
                "other_access": 7,
                "owner": "cloud-admin",
                "owner_access": 7
            },
            "uuid": {
                "uuid_lslong": 11086156774262128366,
                "uuid_mslong": 10125498831222882614
            }
        },
        "name": "vn-blue",
        "network_ipam_refs": [
            {
                "attr": {
                    "ipam_subnets": [
                        {
                            "default_gateway": "10.1.1.254",
                            "subnet": {
                                "ip_prefix": "10.1.1.0",
                                "ip_prefix_len": 24
                            }
                        }
                    ]
                },
                "href": "http://10.84.14.2:8082/network-ipam/a01b486e-2c3e-47df-811c-440e59417ed8",
                "to": [
                    "default-domain",
                    "default-project",
                    "default-network-ipam"
                ],
                "uuid": "a01b486e-2c3e-47df-811c-440e59417ed8"
            }
        ],
        "network_policy_refs": [
            {
                "attr": {
                    "sequence": {
                        "major": 0,
                        "minor": 0
                    }
                },
                "href": "http://10.84.14.2:8082/network-policy/f215a3ec-5cbd-4310-91f4-7bbca52b27bd",
                "to": [
                    "default-domain",
                    "admin",
                    "policy-red-blue"
                ],
                "uuid": "f215a3ec-5cbd-4310-91f4-7bbca52b27bd"
            }
        ],
        "parent_href": "http://10.84.14.2:8082/project/df7649a6-3e2c-4982-b0c3-4b5038eef587",
        "parent_type": "project",
        "parent_uuid": "df7649a6-3e2c-4982-b0c3-4b5038eef587",
        "routing_instances": [
            {
                "href": "http://10.84.14.2:8082/routing-instance/732567fd-8607-4045-b6c0-ff4109d3e0fb",
                "to": [
                    "default-domain",
                    "admin",
                    "vn-blue",
                    "vn-blue"
                ],
                "uuid": "732567fd-8607-4045-b6c0-ff4109d3e0fb"
            }
        ],
        "uuid": "8c84ff8a-30ac-4136-99d9-f0d9662f3eee",
        "virtual_network_properties": {
            "extend_to_external_routers": null,
            "network_id": 4,
            "vxlan_network_identifier": null
        }
    }
}
```

## Updating a resource

To update a resource, a PUT has to be issued on the resource URL.

```text
METHOD: PUT
URL: http://<ip>:<port>/example_resource/<example-resource-uuid>
BODY: JSON representation of resource attributes that are changing
RESPONSE: UUID and href of updated resource
```

References to other resources are specified as a list of dictionaries with “to” and “attr” keys where “to” is the fully-qualified name of the resource being referred to and “attr” is the data associated with the relation (if any).

Example request

```bash
curl -X PUT -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" -d '{"virtual-network": {"fq_name": ["default-domain", "admin", "vn-blue"],"network_policy_refs": [{"to": ["default-domain", "admin", "policy-red-blue"], "attr":{"sequence":{"major":0, "minor": 0}}}]}}' http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee
```

Response

```json
{
    "virtual-network": {
        "href": "http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee",
        "uuid": "8c84ff8a-30ac-4136-99d9-f0d9662f3eee"
    }
}
```

## Deleting a resource

To delete a resource, a DELETE has to be issued on the resource URL

```text
METHOD: DELETE
URL: http://<ip>:<port>/example_resource/<example-resource-uuid>
BODY: None
RESPONSE: None
```

Example Request

```bash
curl -X DELETE -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" http://10.84.14.2:8082/virtual-network/47a91732-629b-4cbe-9aa5-45ba4d7b0e99
```

Response None

## Listing Resources

To list a set of resources, a GET has to be issued on the collection URL with an optional query parameter mentioning the parent resource that contains this collection. If parent resource is not mentioned, a resource named `default-<parent-type>` is assumed.

For detail on available query parameters, such as _filters_ and _fields_ see:

- Output of `contrailcli list -h` command   
- Query parameter keys in block of constants in [api.go file](../pkg/services/baseservices/api.go)
- Definitions of `contrailcli list` command flags using query parameters underneath [contrailcli/list.go file](../pkg/cmd/contrailcli/list.go)

```text
METHOD: GET
URL: http://<ip>:<port>/example_resources
BODY: None
RESPONSE: JSON list of UUID and href of collection if detail not specified, else JSON list of collection dicts
```

Example Request

```bash
curl -X GET -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" http://10.84.14.2:8082/virtual-networks
```

Example Response

```json
{
    "virtual-networks": [
        {
            "fq_name": [
                "default-domain",
                "admin",
                "vn-red"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/47a91732-629b-4cbe-9aa5-45ba4d7b0e99",
            "uuid": "47a91732-629b-4cbe-9aa5-45ba4d7b0e99"
        },
        {
            "fq_name": [
                "default-domain",
                "admin",
                "vn-blue"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee",
            "uuid": "8c84ff8a-30ac-4136-99d9-f0d9662f3eee"
        },
        {
            "fq_name": [
                "default-domain",
                "default-project",
                "ip-fabric"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/aad9e80a-8638-449f-a484-5d1bfd58065c",
            "uuid": "aad9e80a-8638-449f-a484-5d1bfd58065c"
        },
        {
            "fq_name": [
                "default-domain",
                "default-project",
                "default-virtual-network"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/d44a51b0-f2d8-4644-aee0-fe856f970683",
            "uuid": "d44a51b0-f2d8-4644-aee0-fe856f970683"
        },
        {
            "fq_name": [
                "default-domain",
                "default-project",
                "__link_local__"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/f423b6c8-deb6-4325-9035-15a8c8bb0a0d",
            "uuid": "f423b6c8-deb6-4325-9035-15a8c8bb0a0d"
        }
    ]
}
```

Example with pagination limit
``` shell
curl -X GET -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" http://10.84.14.2:8082/virtual-networks?page_limit=1
```

``` javascript
{
    "virtual-networks": [
        {
            "fq_name": [
                "default-domain",
                "admin",
                "vn-red"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/47a91732-629b-4cbe-9aa5-45ba4d7b0e99",
            "uuid": "47a91732-629b-4cbe-9aa5-45ba4d7b0e99"
        }
    ]
 }
```

Example with using of pagination limit and marker
``` shell
curl -X GET -H "X-Auth-Token: $OS_TOKEN" -H "Content-Type: application/json; charset=UTF-8" http://10.84.14.2:8082/virtual-networks?page_limit=1&page_marker=47a91732-629b-4cbe-9aa5-45ba4d7b0e99
```

``` javascript
{
    "virtual-networks": [
        {
            "fq_name": [
                "default-domain",
                "admin",
                "vn-blue"
            ],
            "href": "http://10.84.14.2:8082/virtual-network/8c84ff8a-30ac-4136-99d9-f0d9662f3eee",
            "uuid": "8c84ff8a-30ac-4136-99d9-f0d9662f3eee"
        }
    ]
 }
```

## Relaxing a reference between two resources

Relaxing a reference makes it possible to delete the referred resource even when the reference exists.
When a resource is deleted, all relaxed references to it are deleted as well.

```text
METHOD: POST
URL: http://<ip>:<port>/ref-relax-for-delete
BODY: {
  uuid: "<referring-resource-uuid>",
  ref-uuid: "<referred-resource-uuid>,
}
RESPONSE: {
  uuid: "<referring-resource-uuid>"
}
```


## Sync API

Sync API support creating or deleting multiple resoruces in one time.
This API tries to update resources if same resource with UUID has already exists.

POST
/sync

Body

```json
{
  "resources": [
    {
      "kind": "project",
      "data": {
        "admin_project": {
          "fq_name": [
            "default",
            "admin"
          ],
          "uuid": "admin_project_uuid"
        }
      }
    },
    {
      "kind": "network_policy",
      "data": {
        "fq_name": [
          "default",
          "admin",
          "policy1"
        ],
        "uuid": "network_policy_uuid",
        "parent_type": "project",
        "parent_uuid": "admin_project_uuid",
        "network_policy_entries": {
          "policy_rule": [
            {
              "direction": "<",
              "protocol": "tcp",
              "rule_sequence": {
                "major": 4,
                "minor": 1
              }
            }
          ]
        }
      }
    }
  ]
}
```

### Error handling

Sync API returns following errors:

1. 400 - events has circular dependencies (cannot be sorted)
1. 404, 409, 400 - if there was an error while processing an event (return code is the same as it would be if using resource's endpoint)
1. 500 - there was en error in calling Process method on event (in case of bug in event_encoding)
