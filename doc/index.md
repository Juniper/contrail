# Preface

This is a design document for refactoring Contrail using Go.
The goal of this change is to provide

- Simplify architecture and operation experience
- Higher performance
- Data collections
- Maintainability

# Review process

see [Review Doc](../REVIEW.md)

# Architecture

This diagram shows the overall architecture.

API Server provides REST API and gRPC API.
Internal logic depending on RDBMS and utilize the power
of the relational model, so that developers can focus on implementing logic.
Sync process replicates RDBMS data to [etcd](https://github.com/etcd-io/etcd) using Replication mechanism.

Intent compiler get updates from etcd, evaluate configuration changes and dependency,
and generate config object from the API resource model.

Existing contrail processes such as Control Node, Device manager,
Kube Managers will be capable of watching for any updates in etcd.

![Architecture](./images/architecture.svg "Architecture")

# Process management

To simplify deployment,
we take a "Single binary" approach for Go-related project. We manage multiple goroutine
as micro-services, and run them based on a switch in configuration.
Each process communicate with each other using Service Interface defined above so that we can
switch internal functional call or gRPC call depending on where the other processes running on.

![Process model](./images/process.svg "Processs")

see [Related source code](../pkg/cmd/contrail/run.go)

# Schema

We manage API Schema using the YAML format.
[schemas directory](../schemas) contains schemas.

In YAML format, the schema as following properties.

- id: unique schema ID
- extends: You can specify a list of abstract schema
- parents: parents resources
- references: many to many relations
- prefix: REST API prefix
- schema: JSON Schema

This is a sample schema.

``` YAML
extends:
- base
id: virtual_network
parents:
  project:
    description: Virtual network is collection of end points (interface or ip(s) or
      MAC(s)) that can talk to each other by default. It is collection of subnets
      connected by implicit router which default gateway in each subnet.
    operations: CRUD
    presence: optional
plural: virtual_networks
prefix: /
references:
  network_ipam:
    $ref: types.json#definitions/VnSubnetsType
    description: Reference to network-ipam this network is using. It has list of subnets
      that are to be used as property of the reference.
    operations: CRUD
    presence: required
schema:
  properties:
    external_ipam:
      description: IP address assignment to VM is done statically, outside of (external
        to) Contrail Ipam. vCenter only feature.
      operations: CRUD
      presence: optional
      type: boolean
    fabric_snat:
      default: false
      description: Provide connectivity to underlay network by port mapping
      operations: CRUD
      presence: optional
      type: boolean
  required: []
  type: object
```

# Code generation

# Models

# API Server

# Intent Compiler

# Sync Process
