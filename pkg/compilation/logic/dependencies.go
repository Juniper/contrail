package logic

// ReactionsYAML reaction map stored as yaml
const ReactionsYAML = `
intent-compiler:
  routing-instance:
    self:
    - virtual-network

  virtual-machine-interface:
    self:
    - virtual-machine
    - port-tuple
    - virtual-network
    - bgp-as-a-service
    virtual-network:
    - virtual-machine
    - port-tuple
    - bgp-as-a-service
    logical-router:
    - virtual-network
    instance-ip:
    - virtual-machine
    - port-tuple
    - bgp-as-a-service
    - virtual-network
    floating-ip:
    - virtual-machine
    - port-tuple
    alias-ip:
    - virtual-machine
    - port-tuple
    virtual-machine:
    - virtual-network
    port-tuple:
    - virtual-network

  virtual-network:
    self:
    - network-policy
    - route-table
    - virtual-network
    routing-instance:
    - network-policy

  virtual-machine:
    self:
    - service-instance
    virtual-machine-interface:
    - service-instance
    service-instance:
    - virtual-machine-interface

  port-tuple:
    self:
    - service-instance
    virtual-machine-interface:
    - service-instance
    service-instance:
    - virtual-machine-interface

  service-instance:
    self:
    - network-policy
    - virtual-machine
    - port-tuple
    route-table:
    - network-policy
    routing-policy:
    - network-policy
    route-aggregate:
    - network-policy
    virtual-machine:
    - network-policy
    port-tuple:
    - network-policy
    network-policy:
    - virtual-machine
    - port-tuple

  network-policy:
    self:
    - security-logging-object
    - virtual-network
    - network-policy
    - service-instance
    service-instance:
    - virtual-network
    network-policy:
    - virtual-network
    virtual-network:
    - virtual-network
    - network-policy
    - service-instance

  security-group:
      self:
      - security-group
      - security-logging-object
      security-group:

  route-table:
    self:
    - virtual-network
    - service-instance
    - logical-router
    virtual-network:
    - service-instance
    logical-router:
    - service-instance

  logical-router:
    self:
    - route-table

  floating-ip:
    self:
    - virtual-machine-interface

  alias-ip:
    self:
    - virtual-machine-interface

  instance-ip:
    self:
    - virtual-machine-interface

  bgp-as-a-service:
    self:
    - bgp-router
    virtual-machine-interface:
    - bgp-router

  routing-policy:
    self:
    - service-instance
    - virtual-network

  route-aggregate:
    self:
    - service-instance

  bgpvpn:
    self:
    - virtual-network
    - logical-router
`
