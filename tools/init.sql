



create table access_control_list (
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `access_control_list_hash` text,
    `acl_rule` text,
    `dynamic` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `uuid` varchar(255),
     primary key(uuid));


create table address_group (
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `address_group_prefix` text,
    `uuid` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
     primary key(uuid));


create table alarm (
    `uve_key` text,
    `display_name` varchar(255),
    `fq_name` text,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `alarm_rules` text,
    `alarm_severity` int,
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `uuid` varchar(255),
     primary key(uuid));


create table alias_ip_pool (
    `uuid` varchar(255),
    `fq_name` text,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table alias_ip (
    `uuid` varchar(255),
    `fq_name` text,
    `alias_ip_address` varchar(255),
    `alias_ip_address_family` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
     primary key(uuid));


create table analytics_node (
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `analytics_node_ip_address` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
     primary key(uuid));


create table api_access_list (
    `rbac_rule` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table application_policy_set (
    `uuid` varchar(255),
    `fq_name` text,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `all_applications` bool,
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table bgp_as_a_service (
    `autonomous_system` int,
    `uuid` varchar(255),
    `display_name` varchar(255),
    `bgpaas_suppress_route_advertisement` bool,
    `bgpaas_session_attributes` varchar(255),
    `bgpaas_ipv4_mapped_ipv6_nexthop` bool,
    `bgpaas_ip_address` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `bgpaas_shared` bool,
     primary key(uuid));


create table bgp_router (
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table bgpvpn (
    `bgpvpn_type` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `key_value_pair` text,
    `uuid` varchar(255),
    `route_target` text,
    `export_route_target_list_route_target` text,
    `display_name` varchar(255),
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `route_target_list_route_target` text,
    `fq_name` text,
     primary key(uuid));


create table bridge_domain (
    `isid` int,
    `mac_learning_enabled` bool,
    `uuid` varchar(255),
    `key_value_pair` text,
    `mac_aging_time` int,
    `mac_move_time_window` int,
    `mac_move_limit` int,
    `mac_move_limit_action` varchar(255),
    `mac_limit` int,
    `mac_limit_action` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `display_name` varchar(255),
     primary key(uuid));


create table config_node (
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `config_node_ip_address` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table config_root (
    `uuid` varchar(255),
    `fq_name` text,
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
     primary key(uuid));


create table customer_attachment (
    `uuid` varchar(255),
    `fq_name` text,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
     primary key(uuid));


create table database_node (
    `uuid` varchar(255),
    `database_node_ip_address` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
     primary key(uuid));


create table discovery_service_assignment (
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
     primary key(uuid));


create table domain (
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `project_limit` int,
    `virtual_network_limit` int,
    `security_group_limit` int,
    `uuid` varchar(255),
     primary key(uuid));


create table dsa_rule (
    `fq_name` text,
    `subscriber` text,
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `ep_version` varchar(255),
    `ep_id` varchar(255),
    `ep_type` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `uuid` varchar(255),
     primary key(uuid));


create table e2_service_provider (
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `e2_service_provider_promiscuous` bool,
    `uuid` varchar(255),
    `fq_name` text,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
     primary key(uuid));


create table firewall_policy (
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `display_name` varchar(255),
     primary key(uuid));


create table firewall_rule (
    `display_name` varchar(255),
    `uuid` varchar(255),
    `tag_ids` text,
    `virtual_network` varchar(255),
    `any` bool,
    `address_group` varchar(255),
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `tags` text,
    `protocol_id` int,
    `protocol` varchar(255),
    `start_port` int,
    `end_port` int,
    `src_ports_start_port` int,
    `src_ports_end_port` int,
    `direction` varchar(255),
    `match_tags` text,
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `endpoint_1_any` bool,
    `endpoint_1_address_group` varchar(255),
    `endpoint_1_subnet_ip_prefix` varchar(255),
    `endpoint_1_subnet_ip_prefix_len` int,
    `endpoint_1_tags` text,
    `endpoint_1_tag_ids` text,
    `endpoint_1_virtual_network` varchar(255),
    `assign_routing_instance` varchar(255),
    `analyzer_mac_address` varchar(255),
    `nic_assisted_mirroring_vlan` int,
    `nic_assisted_mirroring` bool,
    `analyzer_name` varchar(255),
    `analyzer_ip_address` varchar(255),
    `encapsulation` varchar(255),
    `nh_mode` varchar(255),
    `routing_instance` varchar(255),
    `juniper_header` bool,
    `vni` int,
    `vtep_dst_ip_address` varchar(255),
    `vtep_dst_mac_address` varchar(255),
    `udp_port` int,
    `simple_action` varchar(255),
    `apply_service` text,
    `gateway_name` varchar(255),
    `log` bool,
    `alert` bool,
    `qos_action` varchar(255),
    `tag_type` text,
    `key_value_pair` text,
     primary key(uuid));


create table floating_ip_pool (
    `uuid` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `subnet_uuid` text,
     primary key(uuid));


create table floating_ip (
    `uuid` varchar(255),
    `fq_name` text,
    `floating_ip_address_family` varchar(255),
    `floating_ip_port_mappings` text,
    `floating_ip_port_mappings_enable` bool,
    `floating_ip_traffic_direction` varchar(255),
    `display_name` varchar(255),
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `floating_ip_is_virtual_ip` bool,
    `floating_ip_address` varchar(255),
    `floating_ip_fixed_ip_address` varchar(255),
    `key_value_pair` text,
     primary key(uuid));


create table forwarding_class (
    `uuid` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `forwarding_class_dscp` int,
    `forwarding_class_mpls_exp` int,
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `forwarding_class_vlan_priority` int,
    `forwarding_class_id` int,
     primary key(uuid));


create table global_qos_config (
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `control` int,
    `analytics` int,
    `dns` int,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
     primary key(uuid));


create table global_system_config (
    `mac_move_time_window` int,
    `mac_move_limit` int,
    `mac_move_limit_action` varchar(255),
    `plugin_property` text,
    `user_defined_log_statistics` text,
    `autonomous_system` int,
    `port_start` int,
    `port_end` int,
    `subnet` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `config_version` varchar(255),
    `ibgp_auto_mesh` bool,
    `enable` bool,
    `end_of_rib_timeout` int,
    `bgp_helper_enable` bool,
    `xmpp_helper_enable` bool,
    `restart_time` int,
    `long_lived_restart_time` int,
    `fq_name` text,
    `display_name` varchar(255),
    `key_value_pair` text,
    `alarm_enable` bool,
    `mac_aging_time` int,
    `bgp_always_compare_med` bool,
    `mac_limit` int,
    `mac_limit_action` varchar(255),
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `id_perms_enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
     primary key(uuid));


create table global_vrouter_config (
    `encapsulation` text,
    `uuid` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `flow_aging_timeout` text,
    `forwarding_mode` varchar(255),
    `flow_export_rate` int,
    `vxlan_network_identifier_mode` varchar(255),
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `fq_name` text,
    `source_port` bool,
    `destination_port` bool,
    `destination_ip` bool,
    `ip_protocol` bool,
    `source_ip` bool,
    `hashing_configured` bool,
    `linklocal_service_entry` text,
    `enable_security_logging` bool,
     primary key(uuid));


create table instance_ip (
    `instance_ip_mode` varchar(255),
    `instance_ip_local_ip` bool,
    `instance_ip_address` varchar(255),
    `subnet_uuid` varchar(255),
    `service_instance_ip` bool,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `fq_name` text,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `key_value_pair` text,
    `display_name` varchar(255),
    `service_health_check_ip` bool,
    `ip_prefix` varchar(255),
    `ip_prefix_len` int,
    `instance_ip_family` varchar(255),
    `instance_ip_secondary` bool,
    `uuid` varchar(255),
     primary key(uuid));


create table interface_route_table (
    `uuid` varchar(255),
    `fq_name` text,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `route` text,
     primary key(uuid));


create table loadbalancer_healthmonitor (
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `display_name` varchar(255),
    `key_value_pair` text,
    `expected_codes` varchar(255),
    `max_retries` int,
    `http_method` varchar(255),
    `admin_state` bool,
    `timeout` int,
    `url_path` varchar(255),
    `monitor_type` varchar(255),
    `delay` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table loadbalancer_listener (
    `uuid` varchar(255),
    `protocol` varchar(255),
    `connection_limit` int,
    `admin_state` bool,
    `sni_containers` text,
    `protocol_port` int,
    `default_tls_container` varchar(255),
    `fq_name` text,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table loadbalancer_member (
    `status` varchar(255),
    `status_description` varchar(255),
    `weight` int,
    `admin_state` bool,
    `address` varchar(255),
    `protocol_port` int,
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table loadbalancer_pool (
    `session_persistence` varchar(255),
    `admin_state` bool,
    `persistence_cookie_name` varchar(255),
    `status_description` varchar(255),
    `loadbalancer_method` varchar(255),
    `status` varchar(255),
    `protocol` varchar(255),
    `subnet_id` varchar(255),
    `loadbalancer_pool_provider` varchar(255),
    `display_name` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `key_value_pair` text,
    `annotations_key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
     primary key(uuid));


create table loadbalancer (
    `status` varchar(255),
    `provisioning_status` varchar(255),
    `admin_state` bool,
    `vip_address` varchar(255),
    `vip_subnet_id` varchar(255),
    `operating_status` varchar(255),
    `loadbalancer_provider` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
     primary key(uuid));


create table logical_interface (
    `logical_interface_vlan_tag` int,
    `logical_interface_type` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table logical_router (
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `vxlan_network_identifier` varchar(255),
    `route_target` text,
    `display_name` varchar(255),
     primary key(uuid));


create table namespace (
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner_access` int,
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `ip_prefix` varchar(255),
    `ip_prefix_len` int,
    `uuid` varchar(255),
    `fq_name` text,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
     primary key(uuid));


create table network_device_config (
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table network_ipam (
    `ipam_subnet_method` varchar(255),
    `uuid` varchar(255),
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `ipam_dns_method` varchar(255),
    `ip_address` varchar(255),
    `virtual_dns_server_name` varchar(255),
    `dhcp_option` text,
    `route` text,
    `ip_prefix` varchar(255),
    `ip_prefix_len` int,
    `ipam_method` varchar(255),
    `ipam_subnets` text,
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
     primary key(uuid));


create table network_policy (
    `uuid` varchar(255),
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `policy_rule` text,
     primary key(uuid));


create table peering_policy (
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `uuid` varchar(255),
    `peering_service` varchar(255),
    `fq_name` text,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
     primary key(uuid));


create table physical_interface (
    `ethernet_segment_identifier` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table physical_router (
    `physical_router_role` varchar(255),
    `username` varchar(255),
    `password` varchar(255),
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `display_name` varchar(255),
    `fq_name` text,
    `key_value_pair` text,
    `physical_router_management_ip` varchar(255),
    `physical_router_product_name` varchar(255),
    `physical_router_lldp` bool,
    `physical_router_snmp` bool,
    `service_port` text,
    `physical_router_loopback_ip` varchar(255),
    `physical_router_dataplane_ip` varchar(255),
    `uuid` varchar(255),
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `v3_security_level` varchar(255),
    `v3_engine_id` varchar(255),
    `v3_security_name` varchar(255),
    `v3_context` varchar(255),
    `v3_privacy_password` varchar(255),
    `v3_context_engine_id` varchar(255),
    `v3_engine_boots` int,
    `timeout` int,
    `v3_engine_time` int,
    `v2_community` varchar(255),
    `version` int,
    `v3_privacy_protocol` varchar(255),
    `v3_authentication_protocol` varchar(255),
    `v3_authentication_password` varchar(255),
    `retries` int,
    `v3_security_engine_id` varchar(255),
    `local_port` int,
    `physical_router_vendor_name` varchar(255),
    `physical_router_vnc_managed` bool,
    `physical_router_image_uri` varchar(255),
    `resource` text,
    `server_port` int,
    `server_ip` varchar(255),
     primary key(uuid));


create table policy_management (
    `uuid` varchar(255),
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table port_tuple (
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `uuid` varchar(255),
     primary key(uuid));


create table project (
    `alarm_enable` bool,
    `uuid` varchar(255),
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `vxlan_routing` bool,
    `loadbalancer_healthmonitor` int,
    `virtual_ip` int,
    `security_group` int,
    `virtual_router` int,
    `global_vrouter_config` int,
    `instance_ip` int,
    `bgp_router` int,
    `virtual_machine_interface` int,
    `defaults` int,
    `subnet` int,
    `network_policy` int,
    `service_instance` int,
    `access_control_list` int,
    `security_group_rule` int,
    `floating_ip` int,
    `route_table` int,
    `network_ipam` int,
    `service_template` int,
    `loadbalancer_pool` int,
    `virtual_network` int,
    `virtual_DNS` int,
    `security_logging_object` int,
    `virtual_DNS_record` int,
    `logical_router` int,
    `floating_ip_pool` int,
    `loadbalancer_member` int,
    `fq_name` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table provider_attachment (
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `uuid` varchar(255),
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
     primary key(uuid));


create table qos_config (
    `fq_name` text,
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `mpls_exp_entries` text,
    `vlan_priority_entries` text,
    `default_forwarding_class_id` int,
    `dscp_entries` text,
    `uuid` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `qos_config_type` varchar(255),
     primary key(uuid));


create table qos_queue (
    `max_bandwidth` int,
    `min_bandwidth` int,
    `uuid` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `qos_queue_identifier` int,
     primary key(uuid));


create table route_aggregate (
    `uuid` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
     primary key(uuid));


create table route_table (
    `route` text,
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
     primary key(uuid));


create table route_target (
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
     primary key(uuid));


create table routing_instance (
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table routing_policy (
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `uuid` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
     primary key(uuid));


create table security_group (
    `policy_rule` text,
    `configured_security_group_id` int,
    `security_group_id` int,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `display_name` varchar(255),
    `last_modified` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `key_value_pair` text,
     primary key(uuid));


create table security_logging_object (
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `rule` text,
    `security_logging_object_rate` int,
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table service_appliance (
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `key_value_pair` text,
    `fq_name` text,
    `uuid` varchar(255),
    `display_name` varchar(255),
    `annotations_key_value_pair` text,
    `username` varchar(255),
    `password` varchar(255),
    `service_appliance_ip_address` varchar(255),
     primary key(uuid));


create table service_appliance_set (
    `key_value_pair` text,
    `service_appliance_ha_mode` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `fq_name` text,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `enable` bool,
    `display_name` varchar(255),
    `service_appliance_driver` varchar(255),
    `annotations_key_value_pair` text,
    `uuid` varchar(255),
     primary key(uuid));


create table service_connection_module (
    `service_type` varchar(255),
    `e2_service` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
     primary key(uuid));


create table service_endpoint (
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
     primary key(uuid));


create table service_group (
    `service_group_firewall_service_list` text,
    `fq_name` text,
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `uuid` varchar(255),
     primary key(uuid));


create table service_health_check (
    `delayUsecs` int,
    `enabled` bool,
    `max_retries` int,
    `timeout` int,
    `url_path` varchar(255),
    `timeoutUsecs` int,
    `delay` int,
    `health_check_type` varchar(255),
    `http_method` varchar(255),
    `expected_codes` varchar(255),
    `monitor_type` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `last_modified` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
     primary key(uuid));


create table service_instance (
    `auto_policy` bool,
    `left_ip_address` varchar(255),
    `right_virtual_network` varchar(255),
    `max_instances` int,
    `auto_scale` bool,
    `interface_list` text,
    `left_virtual_network` varchar(255),
    `availability_zone` varchar(255),
    `right_ip_address` varchar(255),
    `management_virtual_network` varchar(255),
    `ha_mode` varchar(255),
    `virtual_router_id` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `service_instance_bindings` text,
     primary key(uuid));


create table service_object (
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table service_template (
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `version` int,
    `instance_data` varchar(255),
    `interface_type` text,
    `flavor` varchar(255),
    `image_name` varchar(255),
    `service_mode` varchar(255),
    `service_scaling` bool,
    `availability_zone_enable` bool,
    `ordered_interfaces` bool,
    `service_virtualization_type` varchar(255),
    `service_type` varchar(255),
    `vrouter_instance_type` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table subnet (
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table tag (
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `tag_type_name` varchar(255),
    `tag_value` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `tag_id` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
     primary key(uuid));


create table tag_type (
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `tag_type_id` varchar(255),
    `key_value_pair` text,
     primary key(uuid));



create table virtual_DNS_record (
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `record_name` varchar(255),
    `record_class` varchar(255),
    `record_data` varchar(255),
    `record_type` varchar(255),
    `record_ttl_seconds` int,
    `record_mx_preference` int,
    `uuid` varchar(255),
    `fq_name` text,
     primary key(uuid));


create table virtual_DNS (
    `external_visible` bool,
    `next_virtual_DNS` varchar(255),
    `dynamic_records_from_client` bool,
    `reverse_resolution` bool,
    `default_ttl_seconds` int,
    `record_order` varchar(255),
    `floating_ip_record` varchar(255),
    `domain_name` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
     primary key(uuid));


create table virtual_ip (
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `protocol_port` int,
    `status` varchar(255),
    `protocol` varchar(255),
    `address` varchar(255),
    `subnet_id` varchar(255),
    `persistence_cookie_name` varchar(255),
    `persistence_type` varchar(255),
    `connection_limit` int,
    `admin_state` bool,
    `status_description` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
     primary key(uuid));


create table virtual_machine_interface (
    `key_value_pair` text,
    `dhcp_option` text,
    `virtual_machine_interface_fat_flow_protocols` text,
    `port_security_enabled` bool,
    `virtual_machine_interface_device_owner` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `display_name` varchar(255),
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `mac_address` text,
    `virtual_machine_interface_bindings` text,
    `vlan_tag_based_bridge_domain` bool,
    `virtual_machine_interface_disable_policy` bool,
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `vrf_assign_rule` text,
    `service_interface_type` varchar(255),
    `sub_interface_vlan_tag` int,
    `local_preference` int,
    `traffic_direction` varchar(255),
    `analyzer_mac_address` varchar(255),
    `nic_assisted_mirroring` bool,
    `juniper_header` bool,
    `nh_mode` varchar(255),
    `nic_assisted_mirroring_vlan` int,
    `udp_port` int,
    `analyzer_ip_address` varchar(255),
    `analyzer_name` varchar(255),
    `encapsulation` varchar(255),
    `routing_instance` varchar(255),
    `vni` int,
    `vtep_dst_ip_address` varchar(255),
    `vtep_dst_mac_address` varchar(255),
    `ip_protocol` bool,
    `source_ip` bool,
    `hashing_configured` bool,
    `source_port` bool,
    `destination_port` bool,
    `destination_ip` bool,
    `route` text,
    `allowed_address_pair` text,
     primary key(uuid));


create table virtual_machine (
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `uuid` varchar(255),
     primary key(uuid));


create table virtual_network (
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `display_name` varchar(255),
    `rpf` varchar(255),
    `forwarding_mode` varchar(255),
    `allow_transit` bool,
    `network_id` int,
    `mirror_destination` bool,
    `vxlan_network_identifier` int,
    `virtual_network_network_id` int,
    `router_external` bool,
    `segmentation_id` int,
    `physical_network` varchar(255),
    `pbb_etree_enable` bool,
    `fq_name` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `pbb_evpn_enable` bool,
    `route_target` text,
    `layer2_control_word` bool,
    `key_value_pair` text,
    `mac_move_limit_action` varchar(255),
    `mac_move_time_window` int,
    `mac_move_limit` int,
    `is_shared` bool,
    `route_target_list_route_target` text,
    `multi_policy_service_chains_enabled` bool,
    `mac_aging_time` int,
    `export_route_target_list_route_target` text,
    `flood_unknown_unicast` bool,
    `port_security_enabled` bool,
    `uuid` varchar(255),
    `destination_ip` bool,
    `ip_protocol` bool,
    `source_ip` bool,
    `hashing_configured` bool,
    `source_port` bool,
    `destination_port` bool,
    `address_allocation_mode` varchar(255),
    `mac_learning_enabled` bool,
    `external_ipam` bool,
    `mac_limit` int,
    `mac_limit_action` varchar(255),
     primary key(uuid));


create table virtual_router (
    `virtual_router_dpdk_enabled` bool,
    `virtual_router_ip_address` varchar(255),
    `key_value_pair` text,
    `uuid` varchar(255),
    `fq_name` text,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `virtual_router_type` varchar(255),
    `display_name` varchar(255),
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
     primary key(uuid));


create table appformix_node_role (
    `provisioning_start_time` varchar(255),
    `provisioning_state` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `uuid` varchar(255),
    `provisioning_log` text,
    `provisioning_progress` int,
    `display_name` varchar(255),
    `key_value_pair` text,
    `provisioning_progress_stage` varchar(255),
     primary key(uuid));


create table contrail_analytics_database_node_role (
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `provisioning_progress` int,
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `group_access` int,
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `provisioning_log` text,
    `provisioning_progress_stage` varchar(255),
     primary key(uuid));


create table contrail_analytics_node (
    `provisioning_start_time` varchar(255),
    `provisioning_state` varchar(255),
    `fq_name` text,
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `uuid` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `provisioning_log` text,
    `provisioning_progress` int,
     primary key(uuid));


create table contrail_cluster (
    `config_audit_ttl` varchar(255),
    `contrail_webui` varchar(255),
    `default_vrouter_bond_interface` varchar(255),
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `data_ttl` varchar(255),
    `flow_ttl` varchar(255),
    `statistics_ttl` varchar(255),
    `uuid` varchar(255),
    `default_vrouter_bond_interface_members` varchar(255),
    `display_name` varchar(255),
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `default_gateway` varchar(255),
    `key_value_pair` text,
    `fq_name` text,
     primary key(uuid));


create table contrail_controller_node_role (
    `key_value_pair` text,
    `uuid` varchar(255),
    `provisioning_log` text,
    `provisioning_progress_stage` varchar(255),
    `provisioning_start_time` varchar(255),
    `fq_name` text,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `display_name` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `provisioning_progress` int,
    `provisioning_state` varchar(255),
     primary key(uuid));


create table controller_node_role (
    `provisioning_log` text,
    `provisioning_progress` int,
    `storage_management_bond_interface_members` varchar(255),
    `display_name` varchar(255),
    `performance_drives` varchar(255),
    `last_modified` varchar(255),
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `uuid` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `capacity_drives` varchar(255),
    `internalapi_bond_interface_members` varchar(255),
    `key_value_pair` text,
    `provisioning_start_time` varchar(255),
    `provisioning_state` varchar(255),
    `fq_name` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
     primary key(uuid));


create table dashboard (
    `share` text,
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `container_config` varchar(255),
    `uuid` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `permissions_owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
     primary key(uuid));


create table kubernetes_cluster (
    `uuid` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `contrail_cluster_id` varchar(255),
    `kuberunetes_dashboard` varchar(255),
    `key_value_pair` text,
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
     primary key(uuid));


create table kubernetes_node (
    `provisioning_progress_stage` varchar(255),
    `provisioning_state` varchar(255),
    `fq_name` text,
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `description` varchar(255),
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `provisioning_log` text,
    `provisioning_progress` int,
    `provisioning_start_time` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `uuid` varchar(255),
     primary key(uuid));


create table location (
    `private_ospd_user_name` varchar(255),
    `gcp_account_info` varchar(255),
    `gcp_subnet` varchar(255),
    `uuid` varchar(255),
    `private_ntp_hosts` varchar(255),
    `aws_subnet` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `display_name` varchar(255),
    `provisioning_start_time` varchar(255),
    `private_dns_servers` varchar(255),
    `private_ospd_vm_ram_mb` varchar(255),
    `private_redhat_subscription_user` varchar(255),
    `aws_secret_key` varchar(255),
    `private_ospd_package_url` varchar(255),
    `private_ospd_user_password` varchar(255),
    `private_redhat_subscription_key` varchar(255),
    `gcp_region` varchar(255),
    `type` varchar(255),
    `key_value_pair` text,
    `provisioning_state` varchar(255),
    `provisioning_progress` int,
    `private_redhat_subscription_pasword` varchar(255),
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `provisioning_progress_stage` varchar(255),
    `provisioning_log` text,
    `private_ospd_vm_disk_gb` varchar(255),
    `private_ospd_vm_name` varchar(255),
    `private_ospd_vm_vcpus` varchar(255),
    `private_redhat_pool_id` varchar(255),
    `fq_name` text,
    `gcp_asn` int,
    `aws_access_key` varchar(255),
    `aws_region` varchar(255),
     primary key(uuid));


create table node (
    `mac_address` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `private_machine_properties` varchar(255),
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `uuid` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` text,
    `hostname` varchar(255),
    `password` varchar(255),
    `ssh_key` varchar(255),
    `gcp_machine_type` varchar(255),
    `private_power_management_password` varchar(255),
    `fq_name` text,
    `username` varchar(255),
    `aws_ami` varchar(255),
    `gcp_image` varchar(255),
    `private_machine_state` varchar(255),
    `private_power_management_ip` varchar(255),
    `private_power_management_username` varchar(255),
    `ip_address` varchar(255),
    `type` varchar(255),
    `aws_instance_type` varchar(255),
     primary key(uuid));


create table openstack_cluster (
    `fq_name` text,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `contrail_cluster_id` varchar(255),
    `default_capacity_drives` varchar(255),
    `default_storage_backend_bond_interface_members` varchar(255),
    `external_allocation_pool_start` varchar(255),
    `public_gateway` varchar(255),
    `public_ip` varchar(255),
    `uuid` varchar(255),
    `display_name` varchar(255),
    `admin_password` varchar(255),
    `default_osd_drives` varchar(255),
    `default_performance_drives` varchar(255),
    `provisioning_log` text,
    `provisioning_progress` int,
    `provisioning_progress_stage` varchar(255),
    `key_value_pair` text,
    `global_access` int,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `provisioning_start_time` varchar(255),
    `provisioning_state` varchar(255),
    `default_journal_drives` varchar(255),
    `default_storage_access_bond_interface_members` varchar(255),
    `openstack_webui` varchar(255),
    `external_allocation_pool_end` varchar(255),
    `external_net_cidr` varchar(255),
     primary key(uuid));


create table openstack_compute_node_role (
    `fq_name` text,
    `uuid` varchar(255),
    `key_value_pair` text,
    `provisioning_start_time` varchar(255),
    `provisioning_state` varchar(255),
    `default_gateway` varchar(255),
    `vrouter_bond_interface_members` varchar(255),
    `vrouter_type` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `global_access` int,
    `share` text,
    `provisioning_progress` int,
    `provisioning_progress_stage` varchar(255),
    `vrouter_bond_interface` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `permissions_owner` varchar(255),
    `permissions_owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `display_name` varchar(255),
    `provisioning_log` text,
     primary key(uuid));


create table openstack_storage_node_role (
    `storage_backend_bond_interface_members` varchar(255),
    `uuid` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_progress` int,
    `fq_name` text,
    `display_name` varchar(255),
    `key_value_pair` text,
    `provisioning_log` text,
    `journal_drives` varchar(255),
    `osd_drives` varchar(255),
    `storage_access_bond_interface_members` varchar(255),
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `owner` varchar(255),
    `owner_access` int,
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `enable` bool,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
     primary key(uuid));


create table vpn_group (
    `fq_name` text,
    `provisioning_start_time` varchar(255),
    `provisioning_log` text,
    `provisioning_progress` int,
    `enable` bool,
    `description` varchar(255),
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `display_name` varchar(255),
    `uuid` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_state` varchar(255),
    `type` varchar(255),
    `key_value_pair` text,
    `share` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
     primary key(uuid));


create table widget (
    `created` varchar(255),
    `creator` varchar(255),
    `user_visible` bool,
    `last_modified` varchar(255),
    `other_access` int,
    `group` varchar(255),
    `group_access` int,
    `owner` varchar(255),
    `owner_access` int,
    `enable` bool,
    `description` varchar(255),
    `display_name` varchar(255),
    `content_config` varchar(255),
    `uuid` varchar(255),
    `key_value_pair` text,
    `perms2_owner` varchar(255),
    `perms2_owner_access` int,
    `global_access` int,
    `share` text,
    `fq_name` text,
    `container_config` varchar(255),
    `layout_config` varchar(255),
     primary key(uuid));








create table parent_access_control_list_security_group (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references access_control_list(uuid), 
    foreign key (`to`) references security_group(uuid) );

create table parent_access_control_list_virtual_network (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references access_control_list(uuid), 
    foreign key (`to`) references virtual_network(uuid) );





create table parent_address_group_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references address_group(uuid), 
    foreign key (`to`) references project(uuid) );

create table parent_address_group_policy_management (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references address_group(uuid), 
    foreign key (`to`) references policy_management(uuid) );





create table parent_alarm_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references alarm(uuid), 
    foreign key (`to`) references global_system_config(uuid) );

create table parent_alarm_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references alarm(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_alias_ip_pool_virtual_network (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references alias_ip_pool(uuid), 
    foreign key (`to`) references virtual_network(uuid) );




create table ref_alias_ip_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references alias_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );

create table ref_alias_ip_project (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references alias_ip(uuid) on delete cascade, 
    foreign key (`to`) references project(uuid) );


create table parent_alias_ip_alias_ip_pool (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references alias_ip(uuid), 
    foreign key (`to`) references alias_ip_pool(uuid) );





create table parent_analytics_node_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references analytics_node(uuid), 
    foreign key (`to`) references global_system_config(uuid) );





create table parent_api_access_list_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references api_access_list(uuid), 
    foreign key (`to`) references project(uuid) );

create table parent_api_access_list_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references api_access_list(uuid), 
    foreign key (`to`) references global_system_config(uuid) );

create table parent_api_access_list_domain (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references api_access_list(uuid), 
    foreign key (`to`) references domain(uuid) );




create table ref_application_policy_set_firewall_policy (
    `from` varchar(255), 
    `to` varchar(255),
    `sequence` varchar(255),
     foreign key (`from`) references application_policy_set(uuid) on delete cascade, 
    foreign key (`to`) references firewall_policy(uuid) );

create table ref_application_policy_set_global_vrouter_config (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references application_policy_set(uuid) on delete cascade, 
    foreign key (`to`) references global_vrouter_config(uuid) );


create table parent_application_policy_set_policy_management (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references application_policy_set(uuid), 
    foreign key (`to`) references policy_management(uuid) );

create table parent_application_policy_set_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references application_policy_set(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_bgp_as_a_service_service_health_check (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references bgp_as_a_service(uuid) on delete cascade, 
    foreign key (`to`) references service_health_check(uuid) );

create table ref_bgp_as_a_service_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references bgp_as_a_service(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );


create table parent_bgp_as_a_service_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references bgp_as_a_service(uuid), 
    foreign key (`to`) references project(uuid) );









create table parent_bgpvpn_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references bgpvpn(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_bridge_domain_virtual_network (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references bridge_domain(uuid), 
    foreign key (`to`) references virtual_network(uuid) );





create table parent_config_node_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references config_node(uuid), 
    foreign key (`to`) references global_system_config(uuid) );




create table ref_config_root_tag (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references config_root(uuid) on delete cascade, 
    foreign key (`to`) references tag(uuid) );





create table ref_customer_attachment_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references customer_attachment(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );

create table ref_customer_attachment_floating_ip (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references customer_attachment(uuid) on delete cascade, 
    foreign key (`to`) references floating_ip(uuid) );






create table parent_database_node_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references database_node(uuid), 
    foreign key (`to`) references global_system_config(uuid) );









create table parent_domain_config_root (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references domain(uuid), 
    foreign key (`to`) references config_root(uuid) );





create table parent_dsa_rule_discovery_service_assignment (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references dsa_rule(uuid), 
    foreign key (`to`) references discovery_service_assignment(uuid) );




create table ref_e2_service_provider_physical_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references e2_service_provider(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid) );

create table ref_e2_service_provider_peering_policy (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references e2_service_provider(uuid) on delete cascade, 
    foreign key (`to`) references peering_policy(uuid) );





create table ref_firewall_policy_firewall_rule (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references firewall_policy(uuid) on delete cascade, 
    foreign key (`to`) references firewall_rule(uuid) );

create table ref_firewall_policy_security_logging_object (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references firewall_policy(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid) );


create table parent_firewall_policy_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references firewall_policy(uuid), 
    foreign key (`to`) references project(uuid) );

create table parent_firewall_policy_policy_management (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references firewall_policy(uuid), 
    foreign key (`to`) references policy_management(uuid) );




create table ref_firewall_rule_service_group (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references service_group(uuid) );

create table ref_firewall_rule_address_group (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references address_group(uuid) );

create table ref_firewall_rule_security_logging_object (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid) );

create table ref_firewall_rule_virtual_network (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid) );


create table parent_firewall_rule_policy_management (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references firewall_rule(uuid), 
    foreign key (`to`) references policy_management(uuid) );

create table parent_firewall_rule_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references firewall_rule(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_floating_ip_pool_virtual_network (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references floating_ip_pool(uuid), 
    foreign key (`to`) references virtual_network(uuid) );




create table ref_floating_ip_project (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references floating_ip(uuid) on delete cascade, 
    foreign key (`to`) references project(uuid) );

create table ref_floating_ip_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references floating_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );


create table parent_floating_ip_instance_ip (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references floating_ip(uuid), 
    foreign key (`to`) references instance_ip(uuid) );

create table parent_floating_ip_floating_ip_pool (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references floating_ip(uuid), 
    foreign key (`to`) references floating_ip_pool(uuid) );




create table ref_forwarding_class_qos_queue (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references forwarding_class(uuid) on delete cascade, 
    foreign key (`to`) references qos_queue(uuid) );


create table parent_forwarding_class_global_qos_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references forwarding_class(uuid), 
    foreign key (`to`) references global_qos_config(uuid) );





create table parent_global_qos_config_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references global_qos_config(uuid), 
    foreign key (`to`) references global_system_config(uuid) );




create table ref_global_system_config_bgp_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references global_system_config(uuid) on delete cascade, 
    foreign key (`to`) references bgp_router(uuid) );


create table parent_global_system_config_config_root (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references global_system_config(uuid), 
    foreign key (`to`) references config_root(uuid) );





create table parent_global_vrouter_config_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references global_vrouter_config(uuid), 
    foreign key (`to`) references global_system_config(uuid) );




create table ref_instance_ip_network_ipam (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references network_ipam(uuid) );

create table ref_instance_ip_virtual_network (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid) );

create table ref_instance_ip_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );

create table ref_instance_ip_physical_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid) );

create table ref_instance_ip_virtual_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_router(uuid) );





create table ref_interface_route_table_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
    `interface_type` varchar(255),
     foreign key (`from`) references interface_route_table(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );


create table parent_interface_route_table_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references interface_route_table(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_loadbalancer_healthmonitor_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_healthmonitor(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_loadbalancer_listener_loadbalancer (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_listener(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer(uuid) );


create table parent_loadbalancer_listener_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_listener(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_loadbalancer_member_loadbalancer_pool (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_member(uuid), 
    foreign key (`to`) references loadbalancer_pool(uuid) );




create table ref_loadbalancer_pool_loadbalancer_listener (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer_listener(uuid) );

create table ref_loadbalancer_pool_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );

create table ref_loadbalancer_pool_loadbalancer_healthmonitor (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer_healthmonitor(uuid) );

create table ref_loadbalancer_pool_service_appliance_set (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references service_appliance_set(uuid) );

create table ref_loadbalancer_pool_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );


create table parent_loadbalancer_pool_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer_pool(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_loadbalancer_service_appliance_set (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer(uuid) on delete cascade, 
    foreign key (`to`) references service_appliance_set(uuid) );

create table ref_loadbalancer_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );

create table ref_loadbalancer_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );


create table parent_loadbalancer_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references loadbalancer(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_logical_interface_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );


create table parent_logical_interface_physical_router (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references logical_interface(uuid), 
    foreign key (`to`) references physical_router(uuid) );

create table parent_logical_interface_physical_interface (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references logical_interface(uuid), 
    foreign key (`to`) references physical_interface(uuid) );




create table ref_logical_router_route_table (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references route_table(uuid) );

create table ref_logical_router_virtual_network (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid) );

create table ref_logical_router_physical_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid) );

create table ref_logical_router_bgpvpn (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references bgpvpn(uuid) );

create table ref_logical_router_route_target (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references route_target(uuid) );

create table ref_logical_router_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );

create table ref_logical_router_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );


create table parent_logical_router_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references logical_router(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_namespace_domain (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references namespace(uuid), 
    foreign key (`to`) references domain(uuid) );




create table ref_network_device_config_physical_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references network_device_config(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid) );





create table ref_network_ipam_virtual_DNS (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references network_ipam(uuid) on delete cascade, 
    foreign key (`to`) references virtual_DNS(uuid) );


create table parent_network_ipam_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references network_ipam(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_network_policy_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references network_policy(uuid), 
    foreign key (`to`) references project(uuid) );








create table ref_physical_interface_physical_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references physical_interface(uuid) on delete cascade, 
    foreign key (`to`) references physical_interface(uuid) );


create table parent_physical_interface_physical_router (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references physical_interface(uuid), 
    foreign key (`to`) references physical_router(uuid) );




create table ref_physical_router_virtual_network (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references physical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid) );

create table ref_physical_router_bgp_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references physical_router(uuid) on delete cascade, 
    foreign key (`to`) references bgp_router(uuid) );

create table ref_physical_router_virtual_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references physical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_router(uuid) );


create table parent_physical_router_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references physical_router(uuid), 
    foreign key (`to`) references global_system_config(uuid) );

create table parent_physical_router_location (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references physical_router(uuid), 
    foreign key (`to`) references location(uuid) );









create table parent_port_tuple_service_instance (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references port_tuple(uuid), 
    foreign key (`to`) references service_instance(uuid) );




create table ref_project_alias_ip_pool (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references alias_ip_pool(uuid) );

create table ref_project_namespace (
    `from` varchar(255), 
    `to` varchar(255),
    `ip_prefix` varchar(255),
    `ip_prefix_len` int,
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references namespace(uuid) );

create table ref_project_application_policy_set (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references application_policy_set(uuid) );

create table ref_project_floating_ip_pool (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references floating_ip_pool(uuid) );


create table parent_project_domain (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references project(uuid), 
    foreign key (`to`) references domain(uuid) );




create table ref_provider_attachment_virtual_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references provider_attachment(uuid) on delete cascade, 
    foreign key (`to`) references virtual_router(uuid) );





create table ref_qos_config_global_system_config (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references qos_config(uuid) on delete cascade, 
    foreign key (`to`) references global_system_config(uuid) );


create table parent_qos_config_global_qos_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references qos_config(uuid), 
    foreign key (`to`) references global_qos_config(uuid) );

create table parent_qos_config_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references qos_config(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_qos_queue_global_qos_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references qos_queue(uuid), 
    foreign key (`to`) references global_qos_config(uuid) );




create table ref_route_aggregate_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references route_aggregate(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );


create table parent_route_aggregate_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references route_aggregate(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_route_table_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references route_table(uuid), 
    foreign key (`to`) references project(uuid) );









create table parent_routing_instance_virtual_network (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references routing_instance(uuid), 
    foreign key (`to`) references virtual_network(uuid) );




create table ref_routing_policy_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
    `right_sequence` varchar(255),
    `left_sequence` varchar(255),
     foreign key (`from`) references routing_policy(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );


create table parent_routing_policy_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references routing_policy(uuid), 
    foreign key (`to`) references project(uuid) );





create table parent_security_group_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references security_group(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_security_logging_object_security_group (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references security_logging_object(uuid) on delete cascade, 
    foreign key (`to`) references security_group(uuid) );

create table ref_security_logging_object_network_policy (
    `from` varchar(255), 
    `to` varchar(255),
    `rule` text,
     foreign key (`from`) references security_logging_object(uuid) on delete cascade, 
    foreign key (`to`) references network_policy(uuid) );


create table parent_security_logging_object_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references security_logging_object(uuid), 
    foreign key (`to`) references project(uuid) );

create table parent_security_logging_object_global_vrouter_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references security_logging_object(uuid), 
    foreign key (`to`) references global_vrouter_config(uuid) );




create table ref_service_appliance_physical_interface (
    `from` varchar(255), 
    `to` varchar(255),
    `interface_type` varchar(255),
     foreign key (`from`) references service_appliance(uuid) on delete cascade, 
    foreign key (`to`) references physical_interface(uuid) );


create table parent_service_appliance_service_appliance_set (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_appliance(uuid), 
    foreign key (`to`) references service_appliance_set(uuid) );





create table parent_service_appliance_set_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_appliance_set(uuid), 
    foreign key (`to`) references global_system_config(uuid) );




create table ref_service_connection_module_service_object (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_connection_module(uuid) on delete cascade, 
    foreign key (`to`) references service_object(uuid) );





create table ref_service_endpoint_service_connection_module (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_endpoint(uuid) on delete cascade, 
    foreign key (`to`) references service_connection_module(uuid) );

create table ref_service_endpoint_physical_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_endpoint(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid) );

create table ref_service_endpoint_service_object (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_endpoint(uuid) on delete cascade, 
    foreign key (`to`) references service_object(uuid) );






create table parent_service_group_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_group(uuid), 
    foreign key (`to`) references project(uuid) );

create table parent_service_group_policy_management (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_group(uuid), 
    foreign key (`to`) references policy_management(uuid) );




create table ref_service_health_check_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_health_check(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );


create table parent_service_health_check_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_health_check(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_service_instance_instance_ip (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_instance(uuid) on delete cascade, 
    foreign key (`to`) references instance_ip(uuid) );

create table ref_service_instance_service_template (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_instance(uuid) on delete cascade, 
    foreign key (`to`) references service_template(uuid) );


create table parent_service_instance_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_instance(uuid), 
    foreign key (`to`) references project(uuid) );








create table ref_service_template_service_appliance_set (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references service_template(uuid) on delete cascade, 
    foreign key (`to`) references service_appliance_set(uuid) );


create table parent_service_template_domain (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references service_template(uuid), 
    foreign key (`to`) references domain(uuid) );




create table ref_subnet_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references subnet(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );





create table ref_tag_tag_type (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references tag(uuid) on delete cascade, 
    foreign key (`to`) references tag_type(uuid) );


create table parent_tag_config_root (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references tag(uuid), 
    foreign key (`to`) references config_root(uuid) );

create table parent_tag_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references tag(uuid), 
    foreign key (`to`) references project(uuid) );










create table parent_virtual_DNS_record_virtual_DNS (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_DNS_record(uuid), 
    foreign key (`to`) references virtual_DNS(uuid) );





create table parent_virtual_DNS_domain (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_DNS(uuid), 
    foreign key (`to`) references domain(uuid) );




create table ref_virtual_ip_loadbalancer_pool (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_ip(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer_pool(uuid) );

create table ref_virtual_ip_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );


create table parent_virtual_ip_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_ip(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_virtual_machine_interface_physical_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references physical_interface(uuid) );

create table ref_virtual_machine_interface_interface_route_table (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references interface_route_table(uuid) );

create table ref_virtual_machine_interface_routing_instance (
    `from` varchar(255), 
    `to` varchar(255),
    `protocol` varchar(255),
    `ipv6_service_chain_address` varchar(255),
    `direction` varchar(255),
    `mpls_label` int,
    `vlan_tag` int,
    `src_mac` varchar(255),
    `service_chain_address` varchar(255),
    `dst_mac` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references routing_instance(uuid) );

create table ref_virtual_machine_interface_port_tuple (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references port_tuple(uuid) );

create table ref_virtual_machine_interface_security_logging_object (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid) );

create table ref_virtual_machine_interface_security_group (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references security_group(uuid) );

create table ref_virtual_machine_interface_bridge_domain (
    `from` varchar(255), 
    `to` varchar(255),
    `vlan_tag` int,
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references bridge_domain(uuid) );

create table ref_virtual_machine_interface_virtual_network (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid) );

create table ref_virtual_machine_interface_virtual_machine_interface (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid) );

create table ref_virtual_machine_interface_qos_config (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references qos_config(uuid) );

create table ref_virtual_machine_interface_service_health_check (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references service_health_check(uuid) );

create table ref_virtual_machine_interface_virtual_machine (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine(uuid) );

create table ref_virtual_machine_interface_bgp_router (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references bgp_router(uuid) );

create table ref_virtual_machine_interface_service_endpoint (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references service_endpoint(uuid) );


create table parent_virtual_machine_interface_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid), 
    foreign key (`to`) references project(uuid) );

create table parent_virtual_machine_interface_virtual_machine (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid), 
    foreign key (`to`) references virtual_machine(uuid) );

create table parent_virtual_machine_interface_virtual_router (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid), 
    foreign key (`to`) references virtual_router(uuid) );




create table ref_virtual_machine_service_instance (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_machine(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid) );





create table ref_virtual_network_network_ipam (
    `from` varchar(255), 
    `to` varchar(255),
    `ipam_subnets` text,
    `route` text,
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references network_ipam(uuid) );

create table ref_virtual_network_security_logging_object (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid) );

create table ref_virtual_network_network_policy (
    `from` varchar(255), 
    `to` varchar(255),
    `on_interval` varchar(255),
    `end_time` varchar(255),
    `start_time` varchar(255),
    `off_interval` varchar(255),
    `major` int,
    `minor` int,
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references network_policy(uuid) );

create table ref_virtual_network_qos_config (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references qos_config(uuid) );

create table ref_virtual_network_route_table (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references route_table(uuid) );

create table ref_virtual_network_virtual_network (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid) );

create table ref_virtual_network_bgpvpn (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references bgpvpn(uuid) );


create table parent_virtual_network_project (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_network(uuid), 
    foreign key (`to`) references project(uuid) );




create table ref_virtual_router_virtual_machine (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references virtual_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine(uuid) );

create table ref_virtual_router_network_ipam (
    `from` varchar(255), 
    `to` varchar(255),
    `subnet` text,
    `allocation_pools` text,
     foreign key (`from`) references virtual_router(uuid) on delete cascade, 
    foreign key (`to`) references network_ipam(uuid) );


create table parent_virtual_router_global_system_config (
    `from` varchar(255) unique, 
    `to` varchar(255),
     foreign key (`from`) references virtual_router(uuid), 
    foreign key (`to`) references global_system_config(uuid) );




























































create table ref_vpn_group_location (
    `from` varchar(255), 
    `to` varchar(255),
     foreign key (`from`) references vpn_group(uuid) on delete cascade, 
    foreign key (`to`) references location(uuid) );







