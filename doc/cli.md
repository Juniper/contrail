# API Server command line client

API Server CLI is contained in both `contrail` and `contrailcli` executables.
It consists of following commands:
- schema
- show
- list
- create
- set
- update
- sync
- rm
- delete

Find information about them invoking `help`:

	contrailcli <command> -h

## Configuration

CLI reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/agent/agent.go) as the `Config` structure.
Note that Agent-specific fields, such as `watcher` and `tasks` are not required - to be changed in the future.

Example configuration can be found [here](../tools/contrailcli.yml).  

## Running

Usage examples for available commands are presented below.

### Schema command

Schema command shows schema for specified resource.

	contrailcli schema virtual_network -c integration/contrailcli.yml

Schema command output (`integration/cli/testdata/virtual_network_schema.yml`):
``` yaml
#  
- kind: virtual_network
  data: 
	display_name:  #  (string) 
	mac_aging_time: 300 #  (integer) 
	mac_move_control:  #  (object) 
	pbb_evpn_enable: False #  (boolean) 
	provider_properties:  #  (object) 
	is_shared:  #  (boolean) 
	parent_uuid:  #  (string) 
	port_security_enabled: True #  (boolean) 
	virtual_network_network_id:  #  (integer) 
	annotations:  #  (object) 
	external_ipam:  #  (boolean) 
	fq_name:  #  (array) 
	layer2_control_word: False #  (boolean) 
	mac_limit_control:  #  (object) 
	router_external:  #  (boolean) 
	uuid:  #  (string) 
	mac_learning_enabled: False #  (boolean) 
	address_allocation_mode:  #  (string) 
	pbb_etree_enable: False #  (boolean) 
	virtual_network_properties:  #  (object) 
	id_perms:  #  (object) 
	perms2:  #  (object) 
	ecmp_hashing_include_fields:  #  (object) 
	export_route_target_list:  #  (object) 
	flood_unknown_unicast: False #  (boolean) 
	import_route_target_list:  #  (object) 
	multi_policy_service_chains_enabled:  #  (boolean) 
	parent_type:  #  (string) 
	route_target_list:  #  (object)
```

### Show command

Show command shows data of specified resource.

	contrailcli show virtual_network first-uuid -c integration/contrailcli.yml

Show command output (`integration/cli/testdata/virtual_networks_showed.yml`):
``` yaml
kind: virtual_network
data:
  virtual_network:
    - annotations: {}
      ecmp_hashing_include_fields:
        destination_ip: false
        destination_port: false
        hashing_configured: false
        ip_protocol: false
        source_ip: false
        source_port: false
      export_route_target_list: {}
      external_ipam: false
      flood_unknown_unicast: false
      fq_name:
      - vn-blue
      id_perms:
        enable: false
        permissions: {}
        user_visible: false
      import_route_target_list: {}
      is_shared: false
      layer2_control_word: false
      mac_learning_enabled: false
      mac_limit_control: {}
      mac_move_control: {}
      multi_policy_service_chains_enabled: false
      pbb_etree_enable: false
      pbb_evpn_enable: false
      perms2:
        owner: admin
      port_security_enabled: false
      provider_properties: {}
      route_target_list: {}
      router_external: false
      uuid: first-uuid
      virtual_network_properties:
        allow_transit: false
        mirror_destination: false
```

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli show "" "" -c integration/contrailcli.yml

Show command output (shortened):
```
Show command possible usages:

contrail show base $UUID
contrail show has_node $UUID
contrail show has_status $UUID
contrail show access_control_list $UUID
contrail show address_group $UUID
contrail show alarm $UUID
contrail show alias_ip_pool $UUID
contrail show alias_ip $UUID
contrail show analytics_node $UUID
contrail show api_access_list $UUID
contrail show application_policy_set $UUID
contrail show bgp_as_a_service $UUID
contrail show bgp_router $UUID
contrail show bgpvpn $UUID
contrail show bridge_domain $UUID
...
```

### List command

List command lists data of specified resources.

There are also multiple parameters available, such as filters.
Display them with `contrailcli list -h` command.

	contrailcli list virtual_network -c integration/contrailcli.yml

List command output (`integration/cli/testdata/virtual_networks_listed.yml`):
``` yaml
- kind: virtual_network
  data:
    virtual-networks:
    - annotations: {}
      ecmp_hashing_include_fields:
        destination_ip: false
        destination_port: false
        hashing_configured: false
        ip_protocol: false
        source_ip: false
        source_port: false
      export_route_target_list: {}
      external_ipam: false
      flood_unknown_unicast: false
      fq_name:
      - vn-blue
      id_perms:
        enable: false
        permissions: {}
        user_visible: false
      import_route_target_list: {}
      is_shared: false
      layer2_control_word: false
      mac_learning_enabled: false
      mac_limit_control: {}
      mac_move_control: {}
      multi_policy_service_chains_enabled: false
      pbb_etree_enable: false
      pbb_evpn_enable: false
      perms2:
        owner: admin
      port_security_enabled: false
      provider_properties: {}
      route_target_list: {}
      router_external: false
      uuid: first-uuid
      virtual_network_properties:
        allow_transit: false
        mirror_destination: false
    - annotations: {}
      ecmp_hashing_include_fields:
        destination_ip: false
        destination_port: false
        hashing_configured: false
        ip_protocol: false
        source_ip: false
        source_port: false
      export_route_target_list: {}
      external_ipam: true
      flood_unknown_unicast: true
      fq_name:
      - vn-red
      id_perms:
        enable: false
        permissions: {}
        user_visible: false
      import_route_target_list: {}
      is_shared: true
      layer2_control_word: true
      mac_learning_enabled: true
      mac_limit_control: {}
      mac_move_control: {}
      multi_policy_service_chains_enabled: true
      pbb_etree_enable: true
      pbb_evpn_enable: true
      perms2:
        owner: admin
      port_security_enabled: true
      provider_properties: {}
      route_target_list: {}
      router_external: true
      uuid: second-uuid
      virtual_network_properties:
        allow_transit: false
        mirror_destination: false
```

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli list "" -c integration/contrailcli.yml

List command output (shortened):
```
List command possible usages:

contrail list base
contrail list has_node
contrail list has_status
contrail list access_control_list
contrail list address_group
contrail list alarm
contrail list alias_ip_pool
contrail list alias_ip
contrail list analytics_node
contrail list api_access_list
contrail list application_policy_set
contrail list bgp_as_a_service
contrail list bgp_router
contrail list bgpvpn
contrail list bridge_domain
...
```

### Create command

Create command creates resources defined in given YAML file.

	contrailcli create integration/cli/testdata/virtual_networks.yml -c integration/contrailcli.yml
	
Input file content (`integration/cli/testdata/virtual_networks.yml`):
``` yaml
- kind: virtual_network
  data:
  - annotations: {}
    ecmp_hashing_include_fields:
      destination_ip: false
      destination_port: false
      hashing_configured: false
      ip_protocol: false
      source_ip: false
      source_port: false
    export_route_target_list: {}
    external_ipam: false
    flood_unknown_unicast: false
    fq_name:
    - vn-blue
    id_perms:
      enable: false
      permissions: {}
      user_visible: false
    import_route_target_list: {}
    is_shared: false
    layer2_control_word: false
    mac_learning_enabled: false
    mac_limit_control: {}
    mac_move_control: {}
    multi_policy_service_chains_enabled: false
    pbb_etree_enable: false
    pbb_evpn_enable: false
    perms2:
      owner: admin
    port_security_enabled: false
    provider_properties: {}
    route_target_list: {}
    router_external: false
    uuid: first-uuid
    virtual_network_properties:
      allow_transit: false
      mirror_destination: false
  - annotations: {}
    ecmp_hashing_include_fields:
      destination_ip: false
      destination_port: false
      hashing_configured: false
      ip_protocol: false
      source_ip: false
      source_port: false
    export_route_target_list: {}
    external_ipam: true
    flood_unknown_unicast: true
    fq_name:
    - vn-red
    id_perms:
      enable: false
      permissions: {}
      user_visible: false
    import_route_target_list: {}
    is_shared: true
    layer2_control_word: true
    mac_learning_enabled: true
    mac_limit_control: {}
    mac_move_control: {}
    multi_policy_service_chains_enabled: true
    pbb_etree_enable: true
    pbb_evpn_enable: true
    perms2:
      owner: admin
    port_security_enabled: true
    provider_properties: {}
    route_target_list: {}
    router_external: true
    uuid: second-uuid
    virtual_network_properties:
      allow_transit: false
      mirror_destination: false
```

Create command output after successful operation should be the same as input content.

### Set command

Set updates properties of specified resource.

	contrailcli set virtual_network first-uuid "external_ipam: true" -c integration/contrailcli.yml

Set command output (`integration/cli/testdata/virtual_networks_set_output.yml`):
``` yaml
uri: /virtual-network/first-uuid
uuid: first-uuid
```

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli set "" "" "" -c integration/contrailcli.yml

Set command output (shortened):
```
Set command possible usages:

contrail set base $UUID $YAML
contrail set has_node $UUID $YAML
contrail set has_status $UUID $YAML
contrail set access_control_list $UUID $YAML
contrail set address_group $UUID $YAML
contrail set alarm $UUID $YAML
contrail set alias_ip_pool $UUID $YAML
contrail set alias_ip $UUID $YAML
contrail set analytics_node $UUID $YAML
contrail set api_access_list $UUID $YAML
contrail set application_policy_set $UUID $YAML
contrail set bgp_as_a_service $UUID $YAML
contrail set bgp_router $UUID $YAML
contrail set bgpvpn $UUID $YAML
contrail set bridge_domain $UUID $YAML
...
```

### Update command

Update updates resources with data defined in given YAML file.

	contrailcli update integration/cli/testdata/virtual_networks_update.yml -c integration/contrailcli.yml

Input file content (`integration/cli/testdata/virtual_networks_update.yml`):
``` yaml
- kind: virtual_network
  data:
  - external_ipam: true
    flood_unknown_unicast: true
    uuid: first-uuid
```

Update command output (`integration/cli/testdata/virtual_networks_update_output.yml`):
``` yaml
- kind: virtual_network
  data:
  - uri: /virtual-network/first-uuid
    uuid: first-uuid
```

### Sync command

Sync synchronises resources with data defined in given YAML file.
It creates new resource for every not already existing resource.

	contrailcli sync integration/cli/testdata/virtual_networks_update.yml -c integration/contrailcli.yml

Input file content (`integration/cli/testdata/virtual_networks_update.yml`):
``` yaml
- kind: virtual_network
  data:
  - external_ipam: true
    flood_unknown_unicast: true
    uuid: first-uuid
```

Update command output (`integration/cli/testdata/virtual_networks_update_output.yml`):
``` yaml
- kind: virtual_network
  data:
  - uri: /virtual-network/first-uuid
    uuid: first-uuid
```

### Rm command

Rm removes a resource with specified UUID.

	contrailcli rm virtual_network second-uuid -c integration/contrailcli.yml
	
Rm command output is empty on successful operation.

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli rm -c integration/contrailcli.yml

Rm command output (shortened):
```
Remove command possible usages:

contrail rm base $UUID
contrail rm has_node $UUID
contrail rm has_status $UUID
contrail rm access_control_list $UUID
contrail rm address_group $UUID
contrail rm alarm $UUID
contrail rm alias_ip_pool $UUID
contrail rm alias_ip $UUID
contrail rm analytics_node $UUID
contrail rm api_access_list $UUID
contrail rm application_policy_set $UUID
contrail rm bgp_as_a_service $UUID
contrail rm bgp_router $UUID
contrail rm bgpvpn $UUID
contrail rm bridge_domain $UUID
...
```

### Delete command

Delete removes resources specified in given YAML file.

	contrailcli delete integration/cli/testdata/virtual_networks.yml
	
Input file content (`integration/cli/testdata/virtual_networks.yml`):
``` yaml
- kind: virtual_network
  data:
  - annotations: {}
    ecmp_hashing_include_fields:
      destination_ip: false
      destination_port: false
      hashing_configured: false
      ip_protocol: false
      source_ip: false
      source_port: false
    export_route_target_list: {}
    external_ipam: false
    flood_unknown_unicast: false
    fq_name:
    - vn-blue
    id_perms:
      enable: false
      permissions: {}
      user_visible: false
    import_route_target_list: {}
    is_shared: false
    layer2_control_word: false
    mac_learning_enabled: false
    mac_limit_control: {}
    mac_move_control: {}
    multi_policy_service_chains_enabled: false
    pbb_etree_enable: false
    pbb_evpn_enable: false
    perms2:
      owner: admin
    port_security_enabled: false
    provider_properties: {}
    route_target_list: {}
    router_external: false
    uuid: first-uuid
    virtual_network_properties:
      allow_transit: false
      mirror_destination: false
  - annotations: {}
    ecmp_hashing_include_fields:
      destination_ip: false
      destination_port: false
      hashing_configured: false
      ip_protocol: false
      source_ip: false
      source_port: false
    export_route_target_list: {}
    external_ipam: true
    flood_unknown_unicast: true
    fq_name:
    - vn-red
    id_perms:
      enable: false
      permissions: {}
      user_visible: false
    import_route_target_list: {}
    is_shared: true
    layer2_control_word: true
    mac_learning_enabled: true
    mac_limit_control: {}
    mac_move_control: {}
    multi_policy_service_chains_enabled: true
    pbb_etree_enable: true
    pbb_evpn_enable: true
    perms2:
      owner: admin
    port_security_enabled: true
    provider_properties: {}
    route_target_list: {}
    router_external: true
    uuid: second-uuid
    virtual_network_properties:
      allow_transit: false
      mirror_destination: false
```

Delete command returns no output.
