SET GLOBAL group_concat_max_len=100000;

create table metadata (
    `uuid` varchar(255),
    `type` varchar(255),
    `fq_name` varchar(255) unique,
    primary key (`uuid`),
    index fq_name_index (`fq_name`)
 ) CHARACTER SET utf8mb4;




create table access_control_list (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `access_control_list_hash` int,
    `dynamic` bool,
    `acl_rule` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table address_group (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `subnet` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table alarm (
    `uve_key` json,
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `alarm_severity` int,
    `or_list` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table alias_ip_pool (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table alias_ip (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `alias_ip_address_family` varchar(255),
    `alias_ip_address` varchar(255),
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table analytics_node (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `analytics_node_ip_address` varchar(255),
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table api_access_list (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `rbac_rule` json,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table application_policy_set (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `all_applications` bool,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table bgp_as_a_service (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `bgpaas_suppress_route_advertisement` bool,
    `bgpaas_shared` bool,
    `bgpaas_session_attributes` varchar(255),
    `bgpaas_ipv4_mapped_ipv6_nexthop` bool,
    `bgpaas_ip_address` varchar(255),
    `autonomous_system` int,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table bgp_router (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table bgpvpn (
    `uuid` varchar(255),
    `route_target` json,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `import_route_target_list_route_target` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `export_route_target_list_route_target` json,
    `display_name` varchar(255),
    `bgpvpn_type` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table bridge_domain (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `mac_move_time_window` int,
    `mac_move_limit_action` varchar(255),
    `mac_move_limit` int,
    `mac_limit_action` varchar(255),
    `mac_limit` int,
    `mac_learning_enabled` bool,
    `mac_aging_time` int,
    `isid` int,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table config_node (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `config_node_ip_address` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table config_root (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table customer_attachment (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table database_node (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `database_node_ip_address` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table discovery_service_assignment (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table domain (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `virtual_network_limit` int,
    `security_group_limit` int,
    `project_limit` int,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table dsa_rule (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `subscriber` json,
    `ep_version` varchar(255),
    `ep_type` varchar(255),
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `ep_id` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table e2_service_provider (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `e2_service_provider_promiscuous` bool,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table firewall_policy (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table firewall_rule (
    `uuid` varchar(255),
    `start_port` int,
    `end_port` int,
    `protocol_id` int,
    `protocol` varchar(255),
    `dst_ports_start_port` int,
    `dst_ports_end_port` int,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `tag_list` json,
    `tag_type` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `virtual_network` varchar(255),
    `tags` json,
    `tag_ids` json,
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `any` bool,
    `address_group` varchar(255),
    `endpoint_1_virtual_network` varchar(255),
    `endpoint_1_tags` json,
    `endpoint_1_tag_ids` json,
    `subnet_ip_prefix_len` int,
    `subnet_ip_prefix` varchar(255),
    `endpoint_1_any` bool,
    `endpoint_1_address_group` varchar(255),
    `display_name` varchar(255),
    `direction` varchar(255),
    `key_value_pair` json,
    `simple_action` varchar(255),
    `qos_action` varchar(255),
    `udp_port` int,
    `vtep_dst_mac_address` varchar(255),
    `vtep_dst_ip_address` varchar(255),
    `vni` int,
    `routing_instance` varchar(255),
    `nic_assisted_mirroring_vlan` int,
    `nic_assisted_mirroring` bool,
    `nh_mode` varchar(255),
    `juniper_header` bool,
    `encapsulation` varchar(255),
    `analyzer_name` varchar(255),
    `analyzer_mac_address` varchar(255),
    `analyzer_ip_address` varchar(255),
    `log` bool,
    `gateway_name` varchar(255),
    `assign_routing_instance` varchar(255),
    `apply_service` json,
    `alert` bool,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table floating_ip_pool (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `subnet_uuid` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table floating_ip (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `floating_ip_traffic_direction` varchar(255),
    `port_mappings` json,
    `floating_ip_port_mappings_enable` bool,
    `floating_ip_is_virtual_ip` bool,
    `floating_ip_fixed_ip_address` varchar(255),
    `floating_ip_address_family` varchar(255),
    `floating_ip_address` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table forwarding_class (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `forwarding_class_vlan_priority` int,
    `forwarding_class_mpls_exp` int,
    `forwarding_class_id` int,
    `forwarding_class_dscp` int,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table global_qos_config (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `dns` int,
    `control` int,
    `analytics` int,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table global_system_config (
    `uuid` varchar(255),
    `statlist` json,
    `plugin_property` json,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `mac_move_time_window` int,
    `mac_move_limit_action` varchar(255),
    `mac_move_limit` int,
    `mac_limit_action` varchar(255),
    `mac_limit` int,
    `mac_aging_time` int,
    `subnet` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `ibgp_auto_mesh` bool,
    `xmpp_helper_enable` bool,
    `restart_time` int,
    `long_lived_restart_time` int,
    `end_of_rib_timeout` int,
    `graceful_restart_parameters_enable` bool,
    `bgp_helper_enable` bool,
    `fq_name` json,
    `display_name` varchar(255),
    `config_version` varchar(255),
    `port_start` int,
    `port_end` int,
    `bgp_always_compare_med` bool,
    `autonomous_system` int,
    `key_value_pair` json,
    `alarm_enable` bool,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table global_vrouter_config (
    `vxlan_network_identifier_mode` varchar(255),
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `linklocal_service_entry` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `forwarding_mode` varchar(255),
    `flow_export_rate` int,
    `flow_aging_timeout` json,
    `encapsulation` json,
    `enable_security_logging` bool,
    `source_port` bool,
    `source_ip` bool,
    `ip_protocol` bool,
    `hashing_configured` bool,
    `destination_port` bool,
    `destination_ip` bool,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table instance_ip (
    `uuid` varchar(255),
    `subnet_uuid` varchar(255),
    `service_instance_ip` bool,
    `service_health_check_ip` bool,
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `instance_ip_secondary` bool,
    `instance_ip_mode` varchar(255),
    `instance_ip_local_ip` bool,
    `instance_ip_family` varchar(255),
    `instance_ip_address` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table interface_route_table (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `route` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table loadbalancer_healthmonitor (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `url_path` varchar(255),
    `timeout` int,
    `monitor_type` varchar(255),
    `max_retries` int,
    `http_method` varchar(255),
    `expected_codes` varchar(255),
    `delay` int,
    `admin_state` bool,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table loadbalancer_listener (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `sni_containers` json,
    `protocol_port` int,
    `protocol` varchar(255),
    `default_tls_container` varchar(255),
    `connection_limit` int,
    `admin_state` bool,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table loadbalancer_member (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `weight` int,
    `status_description` varchar(255),
    `status` varchar(255),
    `protocol_port` int,
    `admin_state` bool,
    `address` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table loadbalancer_pool (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `loadbalancer_pool_provider` varchar(255),
    `subnet_id` varchar(255),
    `status_description` varchar(255),
    `status` varchar(255),
    `session_persistence` varchar(255),
    `protocol` varchar(255),
    `persistence_cookie_name` varchar(255),
    `loadbalancer_method` varchar(255),
    `admin_state` bool,
    `key_value_pair` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `annotations_key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table loadbalancer (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `loadbalancer_provider` varchar(255),
    `vip_subnet_id` varchar(255),
    `vip_address` varchar(255),
    `status` varchar(255),
    `provisioning_status` varchar(255),
    `operating_status` varchar(255),
    `admin_state` bool,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table logical_interface (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `logical_interface_vlan_tag` int,
    `logical_interface_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table logical_router (
    `vxlan_network_identifier` varchar(255),
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `route_target` json,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table namespace (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table network_device_config (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table network_ipam (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `ipam_method` varchar(255),
    `virtual_dns_server_name` varchar(255),
    `ip_address` varchar(255),
    `ipam_dns_method` varchar(255),
    `route` json,
    `dhcp_option` json,
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `subnets` json,
    `ipam_subnet_method` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table network_policy (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `policy_rule` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table peering_policy (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `peering_service` varchar(255),
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table physical_interface (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `ethernet_segment_identifier` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table physical_router (
    `uuid` varchar(255),
    `server_port` int,
    `server_ip` varchar(255),
    `resource` json,
    `physical_router_vnc_managed` bool,
    `physical_router_vendor_name` varchar(255),
    `username` varchar(255),
    `password` varchar(255),
    `version` int,
    `v3_security_name` varchar(255),
    `v3_security_level` varchar(255),
    `v3_security_engine_id` varchar(255),
    `v3_privacy_protocol` varchar(255),
    `v3_privacy_password` varchar(255),
    `v3_engine_time` int,
    `v3_engine_id` varchar(255),
    `v3_engine_boots` int,
    `v3_context_engine_id` varchar(255),
    `v3_context` varchar(255),
    `v3_authentication_protocol` varchar(255),
    `v3_authentication_password` varchar(255),
    `v2_community` varchar(255),
    `timeout` int,
    `retries` int,
    `local_port` int,
    `physical_router_snmp` bool,
    `physical_router_role` varchar(255),
    `physical_router_product_name` varchar(255),
    `physical_router_management_ip` varchar(255),
    `physical_router_loopback_ip` varchar(255),
    `physical_router_lldp` bool,
    `service_port` json,
    `physical_router_image_uri` varchar(255),
    `physical_router_dataplane_ip` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table policy_management (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table port_tuple (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table project (
    `vxlan_routing` bool,
    `uuid` varchar(255),
    `virtual_router` int,
    `virtual_network` int,
    `virtual_machine_interface` int,
    `virtual_ip` int,
    `virtual_DNS_record` int,
    `virtual_DNS` int,
    `subnet` int,
    `service_template` int,
    `service_instance` int,
    `security_logging_object` int,
    `security_group_rule` int,
    `security_group` int,
    `route_table` int,
    `network_policy` int,
    `network_ipam` int,
    `logical_router` int,
    `loadbalancer_pool` int,
    `loadbalancer_member` int,
    `loadbalancer_healthmonitor` int,
    `instance_ip` int,
    `global_vrouter_config` int,
    `floating_ip_pool` int,
    `floating_ip` int,
    `defaults` int,
    `bgp_router` int,
    `access_control_list` int,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
    `alarm_enable` bool,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table provider_attachment (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table qos_config (
    `qos_id_forwarding_class_pair` json,
    `uuid` varchar(255),
    `qos_config_type` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `mpls_exp_entries_qos_id_forwarding_class_pair` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `dscp_entries_qos_id_forwarding_class_pair` json,
    `display_name` varchar(255),
    `default_forwarding_class_id` int,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table qos_queue (
    `uuid` varchar(255),
    `qos_queue_identifier` int,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `min_bandwidth` int,
    `max_bandwidth` int,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table route_aggregate (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table route_table (
    `uuid` varchar(255),
    `route` json,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table route_target (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table routing_instance (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table routing_policy (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table security_group (
    `uuid` varchar(255),
    `security_group_id` int,
    `policy_rule` json,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `configured_security_group_id` int,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table security_logging_object (
    `uuid` varchar(255),
    `rule` json,
    `security_logging_object_rate` int,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_appliance (
    `uuid` varchar(255),
    `username` varchar(255),
    `password` varchar(255),
    `key_value_pair` json,
    `service_appliance_ip_address` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `annotations_key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_appliance_set (
    `uuid` varchar(255),
    `key_value_pair` json,
    `service_appliance_ha_mode` varchar(255),
    `service_appliance_driver` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `annotations_key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_connection_module (
    `uuid` varchar(255),
    `service_type` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `e2_service` varchar(255),
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_endpoint (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_group (
    `uuid` varchar(255),
    `firewall_service` json,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_health_check (
    `uuid` varchar(255),
    `url_path` varchar(255),
    `timeoutUsecs` int,
    `timeout` int,
    `monitor_type` varchar(255),
    `max_retries` int,
    `http_method` varchar(255),
    `health_check_type` varchar(255),
    `expected_codes` varchar(255),
    `enabled` bool,
    `delayUsecs` int,
    `delay` int,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_instance (
    `uuid` varchar(255),
    `virtual_router_id` varchar(255),
    `max_instances` int,
    `auto_scale` bool,
    `right_virtual_network` varchar(255),
    `right_ip_address` varchar(255),
    `management_virtual_network` varchar(255),
    `left_virtual_network` varchar(255),
    `left_ip_address` varchar(255),
    `interface_list` json,
    `ha_mode` varchar(255),
    `availability_zone` varchar(255),
    `auto_policy` bool,
    `key_value_pair` json,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `annotations_key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_object (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table service_template (
    `uuid` varchar(255),
    `vrouter_instance_type` varchar(255),
    `version` int,
    `service_virtualization_type` varchar(255),
    `service_type` varchar(255),
    `service_scaling` bool,
    `service_mode` varchar(255),
    `ordered_interfaces` bool,
    `interface_type` json,
    `instance_data` varchar(255),
    `image_name` varchar(255),
    `flavor` varchar(255),
    `availability_zone_enable` bool,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table subnet (
    `uuid` varchar(255),
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table tag (
    `uuid` varchar(255),
    `tag_value` varchar(255),
    `tag_type_name` varchar(255),
    `tag_id` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table tag_type (
    `uuid` varchar(255),
    `tag_type_id` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table user (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `password` varchar(255),
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_DNS_record (
    `record_type` varchar(255),
    `record_ttl_seconds` int,
    `record_name` varchar(255),
    `record_mx_preference` int,
    `record_data` varchar(255),
    `record_class` varchar(255),
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_DNS (
    `reverse_resolution` bool,
    `record_order` varchar(255),
    `next_virtual_DNS` varchar(255),
    `floating_ip_record` varchar(255),
    `external_visible` bool,
    `dynamic_records_from_client` bool,
    `domain_name` varchar(255),
    `default_ttl_seconds` int,
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_ip (
    `subnet_id` varchar(255),
    `status_description` varchar(255),
    `status` varchar(255),
    `protocol_port` int,
    `protocol` varchar(255),
    `persistence_type` varchar(255),
    `persistence_cookie_name` varchar(255),
    `connection_limit` int,
    `admin_state` bool,
    `address` varchar(255),
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_machine_interface (
    `vrf_assign_rule` json,
    `vlan_tag_based_bridge_domain` bool,
    `sub_interface_vlan_tag` int,
    `service_interface_type` varchar(255),
    `local_preference` int,
    `traffic_direction` varchar(255),
    `udp_port` int,
    `vtep_dst_mac_address` varchar(255),
    `vtep_dst_ip_address` varchar(255),
    `vni` int,
    `routing_instance` varchar(255),
    `nic_assisted_mirroring_vlan` int,
    `nic_assisted_mirroring` bool,
    `nh_mode` varchar(255),
    `juniper_header` bool,
    `encapsulation` varchar(255),
    `analyzer_name` varchar(255),
    `analyzer_mac_address` varchar(255),
    `analyzer_ip_address` varchar(255),
    `mac_address` json,
    `route` json,
    `fat_flow_protocol` json,
    `virtual_machine_interface_disable_policy` bool,
    `dhcp_option` json,
    `virtual_machine_interface_device_owner` varchar(255),
    `key_value_pair` json,
    `allowed_address_pair` json,
    `uuid` varchar(255),
    `port_security_enabled` bool,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `source_port` bool,
    `source_ip` bool,
    `ip_protocol` bool,
    `hashing_configured` bool,
    `destination_port` bool,
    `destination_ip` bool,
    `display_name` varchar(255),
    `annotations_key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_machine (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_network (
    `vxlan_network_identifier` int,
    `rpf` varchar(255),
    `network_id` int,
    `mirror_destination` bool,
    `forwarding_mode` varchar(255),
    `allow_transit` bool,
    `virtual_network_network_id` int,
    `uuid` varchar(255),
    `router_external` bool,
    `route_target` json,
    `segmentation_id` int,
    `physical_network` varchar(255),
    `port_security_enabled` bool,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `pbb_evpn_enable` bool,
    `pbb_etree_enable` bool,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `multi_policy_service_chains_enabled` bool,
    `mac_move_time_window` int,
    `mac_move_limit_action` varchar(255),
    `mac_move_limit` int,
    `mac_limit_action` varchar(255),
    `mac_limit` int,
    `mac_learning_enabled` bool,
    `mac_aging_time` int,
    `layer2_control_word` bool,
    `is_shared` bool,
    `import_route_target_list_route_target` json,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `flood_unknown_unicast` bool,
    `external_ipam` bool,
    `export_route_target_list_route_target` json,
    `source_port` bool,
    `source_ip` bool,
    `ip_protocol` bool,
    `hashing_configured` bool,
    `destination_port` bool,
    `destination_ip` bool,
    `display_name` varchar(255),
    `key_value_pair` json,
    `address_allocation_mode` varchar(255),
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table virtual_router (
    `virtual_router_type` varchar(255),
    `virtual_router_ip_address` varchar(255),
    `virtual_router_dpdk_enabled` bool,
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table appformix_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table baremetal_node (
    `uuid` varchar(255),
    `updated_at` varchar(255),
    `target_provision_state` varchar(255),
    `target_power_state` varchar(255),
    `provision_state` varchar(255),
    `power_state` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `name` varchar(255),
    `maintenance_reason` varchar(255),
    `maintenance` bool,
    `last_error` varchar(255),
    `instance_uuid` varchar(255),
    `vcpus` varchar(255),
    `swap_mb` varchar(255),
    `root_gb` varchar(255),
    `nova_host_id` varchar(255),
    `memory_mb` varchar(255),
    `local_gb` varchar(255),
    `image_source` varchar(255),
    `display_name` varchar(255),
    `capabilities` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `ipmi_username` varchar(255),
    `ipmi_password` varchar(255),
    `ipmi_address` varchar(255),
    `deploy_ramdisk` varchar(255),
    `deploy_kernel` varchar(255),
    `_display_name` varchar(255),
    `created_at` varchar(255),
    `console_enabled` bool,
    `bm_properties_memory_mb` int,
    `disk_gb` int,
    `cpu_count` int,
    `cpu_arch` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table baremetal_port (
    `uuid` varchar(255),
    `updated_at` varchar(255),
    `pxe_enabled` bool,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `node` varchar(255),
    `mac_address` varchar(255),
    `switch_info` varchar(255),
    `switch_id` varchar(255),
    `port_id` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `created_at` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_analytics_database_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_analytics_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_cluster (
    `uuid` varchar(255),
    `statistics_ttl` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `provisioner_type` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `orchestrator` varchar(255),
    `openstack` varchar(255),
    `kubernetes_master` varchar(255),
    `kubernetes` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `flow_ttl` varchar(255),
    `display_name` varchar(255),
    `default_vrouter_bond_interface_members` varchar(255),
    `default_vrouter_bond_interface` varchar(255),
    `default_gateway` varchar(255),
    `data_ttl` varchar(255),
    `contrail_webui` varchar(255),
    `contrail_vrouter` varchar(255),
    `contrail_control` varchar(255),
    `contrail_configdb` varchar(255),
    `contrail_config` varchar(255),
    `contrail_analyticsdb` varchar(255),
    `contrail_analytics` varchar(255),
    `config_audit_ttl` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_config_database_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_config_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_control_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_storage_node (
    `uuid` varchar(255),
    `storage_backend_bond_interface_members` varchar(255),
    `storage_access_bond_interface_members` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `osd_drives` varchar(255),
    `journal_drives` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_vrouter_node (
    `vrouter_type` varchar(255),
    `vrouter_bond_interface_members` varchar(255),
    `vrouter_bond_interface` varchar(255),
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `default_gateway` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table contrail_controller_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table dashboard (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `container_config` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table flavor (
    `vcpus` int,
    `uuid` varchar(255),
    `swap` int,
    `rxtx_factor` int,
    `ram` int,
    `property` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `name` varchar(255),
    `type` varchar(255),
    `rel` varchar(255),
    `href` varchar(255),
    `is_public` bool,
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `id` varchar(255),
    `fq_name` json,
    `ephemeral` int,
    `display_name` varchar(255),
    `disk` int,
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table os_image (
    `visibility` varchar(255),
    `uuid` varchar(255),
    `updated_at` varchar(255),
    `tags` varchar(255),
    `status` varchar(255),
    `size` int,
    `protected` bool,
    `property` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `_owner` varchar(255),
    `name` varchar(255),
    `min_ram` int,
    `min_disk` int,
    `location` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `id` varchar(255),
    `fq_name` json,
    `file` varchar(255),
    `display_name` varchar(255),
    `disk_format` varchar(255),
    `created_at` varchar(255),
    `container_format` varchar(255),
    `checksum` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table keypair (
    `uuid` varchar(255),
    `public_key` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `name` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table kubernetes_master_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table kubernetes_node (
    `uuid` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table location (
    `uuid` varchar(255),
    `type` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `private_redhat_subscription_user` varchar(255),
    `private_redhat_subscription_pasword` varchar(255),
    `private_redhat_subscription_key` varchar(255),
    `private_redhat_pool_id` varchar(255),
    `private_ospd_vm_vcpus` varchar(255),
    `private_ospd_vm_ram_mb` varchar(255),
    `private_ospd_vm_name` varchar(255),
    `private_ospd_vm_disk_gb` varchar(255),
    `private_ospd_user_password` varchar(255),
    `private_ospd_user_name` varchar(255),
    `private_ospd_package_url` varchar(255),
    `private_ntp_hosts` varchar(255),
    `private_dns_servers` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `gcp_subnet` varchar(255),
    `gcp_region` varchar(255),
    `gcp_asn` int,
    `gcp_account_info` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `aws_subnet` varchar(255),
    `aws_secret_key` varchar(255),
    `aws_region` varchar(255),
    `aws_access_key` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table node (
    `uuid` varchar(255),
    `username` varchar(255),
    `type` varchar(255),
    `ssh_key` tinytext,
    `private_machine_state` varchar(255),
    `private_machine_properties` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `password` varchar(255),
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `mac_address` varchar(255),
    `ipmi_username` varchar(255),
    `ipmi_password` varchar(255),
    `ipmi_address` varchar(255),
    `ip_address` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `hostname` varchar(255),
    `gcp_machine_type` varchar(255),
    `gcp_image` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `aws_instance_type` varchar(255),
    `aws_ami` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table server (
    `uuid` varchar(255),
    `user_id` int,
    `updated` varchar(255),
    `tenant_id` varchar(255),
    `status` varchar(255),
    `progress` int,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `name` varchar(255),
    `locked` bool,
    `type` varchar(255),
    `rel` varchar(255),
    `href` varchar(255),
    `id` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `_id` varchar(255),
    `host_status` varchar(255),
    `hostId` varchar(255),
    `fq_name` json,
    `links_type` varchar(255),
    `links_rel` varchar(255),
    `links_href` varchar(255),
    `flavor_id` varchar(255),
    `display_name` varchar(255),
    `_created` varchar(255),
    `config_drive` bool,
    `key_value_pair` json,
    `addr` varchar(255),
    `accessIPv6` varchar(255),
    `accessIPv4` varchar(255),
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table vpn_group (
    `uuid` varchar(255),
    `type` varchar(255),
    `provisioning_state` varchar(255),
    `provisioning_start_time` varchar(255),
    `provisioning_progress_stage` varchar(255),
    `provisioning_progress` int,
    `provisioning_log` text,
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;


create table widget (
    `uuid` varchar(255),
    `share` json,
    `owner_access` int,
    `owner` varchar(255),
    `global_access` int,
    `parent_uuid` varchar(255),
    `parent_type` varchar(255),
    `layout_config` varchar(255),
    `user_visible` bool,
    `permissions_owner_access` int,
    `permissions_owner` varchar(255),
    `other_access` int,
    `group_access` int,
    `group` varchar(255),
    `last_modified` varchar(255),
    `enable` bool,
    `description` varchar(255),
    `creator` varchar(255),
    `created` varchar(255),
    `fq_name` json,
    `display_name` varchar(255),
    `content_config` varchar(255),
    `container_config` varchar(255),
    `key_value_pair` json,
     primary key(`uuid`),
    index parent_uuid_index (`parent_uuid`)
    ) CHARACTER SET utf8mb4;







create table tenant_share_access_control_list (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references access_control_list(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_access_control_list (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references access_control_list(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_address_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references address_group(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_address_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references address_group(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_alarm (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references alarm(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_alarm (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references alarm(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_alias_ip_pool (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references alias_ip_pool(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_alias_ip_pool (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references alias_ip_pool(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_alias_ip_project (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references alias_ip(uuid) on delete cascade, 
    foreign key (`to`) references project(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_alias_ip_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references alias_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_alias_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references alias_ip(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_alias_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references alias_ip(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_analytics_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references analytics_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_analytics_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references analytics_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_api_access_list (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references api_access_list(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_api_access_list (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references api_access_list(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_application_policy_set_firewall_policy (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `sequence` varchar(255),
     foreign key (`from`) references application_policy_set(uuid) on delete cascade, 
    foreign key (`to`) references firewall_policy(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_application_policy_set_global_vrouter_config (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references application_policy_set(uuid) on delete cascade, 
    foreign key (`to`) references global_vrouter_config(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_application_policy_set (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references application_policy_set(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_application_policy_set (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references application_policy_set(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_bgp_as_a_service_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references bgp_as_a_service(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_bgp_as_a_service_service_health_check (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references bgp_as_a_service(uuid) on delete cascade, 
    foreign key (`to`) references service_health_check(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_bgp_as_a_service (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bgp_as_a_service(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_bgp_as_a_service (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bgp_as_a_service(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_bgp_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bgp_router(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_bgp_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bgp_router(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_bgpvpn (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bgpvpn(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_bgpvpn (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bgpvpn(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_bridge_domain (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bridge_domain(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_bridge_domain (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references bridge_domain(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_config_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references config_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_config_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references config_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_config_root_tag (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references config_root(uuid) on delete cascade, 
    foreign key (`to`) references tag(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_config_root (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references config_root(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_config_root (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references config_root(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_customer_attachment_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references customer_attachment(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_customer_attachment_floating_ip (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references customer_attachment(uuid) on delete cascade, 
    foreign key (`to`) references floating_ip(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_customer_attachment (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references customer_attachment(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_customer_attachment (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references customer_attachment(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_database_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references database_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_database_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references database_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_discovery_service_assignment (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references discovery_service_assignment(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_discovery_service_assignment (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references discovery_service_assignment(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_domain (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references domain(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_domain (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references domain(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_dsa_rule (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references dsa_rule(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_dsa_rule (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references dsa_rule(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_e2_service_provider_physical_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references e2_service_provider(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_e2_service_provider_peering_policy (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references e2_service_provider(uuid) on delete cascade, 
    foreign key (`to`) references peering_policy(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_e2_service_provider (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references e2_service_provider(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_e2_service_provider (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references e2_service_provider(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_firewall_policy_firewall_rule (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `sequence` varchar(255),
     foreign key (`from`) references firewall_policy(uuid) on delete cascade, 
    foreign key (`to`) references firewall_rule(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_firewall_policy_security_logging_object (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references firewall_policy(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_firewall_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references firewall_policy(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_firewall_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references firewall_policy(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_firewall_rule_security_logging_object (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_firewall_rule_virtual_network (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_firewall_rule_service_group (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references service_group(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_firewall_rule_address_group (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references firewall_rule(uuid) on delete cascade, 
    foreign key (`to`) references address_group(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_firewall_rule (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references firewall_rule(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_firewall_rule (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references firewall_rule(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_floating_ip_pool (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references floating_ip_pool(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_floating_ip_pool (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references floating_ip_pool(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_floating_ip_project (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references floating_ip(uuid) on delete cascade, 
    foreign key (`to`) references project(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_floating_ip_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references floating_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_floating_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references floating_ip(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_floating_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references floating_ip(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_forwarding_class_qos_queue (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references forwarding_class(uuid) on delete cascade, 
    foreign key (`to`) references qos_queue(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_forwarding_class (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references forwarding_class(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_forwarding_class (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references forwarding_class(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_global_qos_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references global_qos_config(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_global_qos_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references global_qos_config(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_global_system_config_bgp_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references global_system_config(uuid) on delete cascade, 
    foreign key (`to`) references bgp_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_global_system_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references global_system_config(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_global_system_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references global_system_config(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_global_vrouter_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references global_vrouter_config(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_global_vrouter_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references global_vrouter_config(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_instance_ip_network_ipam (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references network_ipam(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_instance_ip_virtual_network (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_instance_ip_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_instance_ip_physical_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_instance_ip_virtual_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references instance_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_instance_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references instance_ip(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_instance_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references instance_ip(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_interface_route_table_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `interface_type` varchar(255),
     foreign key (`from`) references interface_route_table(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_interface_route_table (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references interface_route_table(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_interface_route_table (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references interface_route_table(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_loadbalancer_healthmonitor (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_healthmonitor(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_loadbalancer_healthmonitor (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_healthmonitor(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_loadbalancer_listener_loadbalancer (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer_listener(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_loadbalancer_listener (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_listener(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_loadbalancer_listener (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_listener(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_loadbalancer_member (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_member(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_loadbalancer_member (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_member(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_loadbalancer_pool_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_loadbalancer_pool_loadbalancer_healthmonitor (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer_healthmonitor(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_loadbalancer_pool_service_appliance_set (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references service_appliance_set(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_loadbalancer_pool_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_loadbalancer_pool_loadbalancer_listener (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer_pool(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer_listener(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_loadbalancer_pool (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_pool(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_loadbalancer_pool (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer_pool(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_loadbalancer_service_appliance_set (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer(uuid) on delete cascade, 
    foreign key (`to`) references service_appliance_set(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_loadbalancer_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_loadbalancer_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references loadbalancer(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_loadbalancer (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_loadbalancer (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references loadbalancer(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_logical_interface_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_logical_interface (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references logical_interface(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_logical_interface (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references logical_interface(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_logical_router_route_target (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references route_target(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_logical_router_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_logical_router_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_logical_router_route_table (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references route_table(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_logical_router_virtual_network (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_logical_router_physical_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_logical_router_bgpvpn (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references logical_router(uuid) on delete cascade, 
    foreign key (`to`) references bgpvpn(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_logical_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references logical_router(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_logical_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references logical_router(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_namespace (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references namespace(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_namespace (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references namespace(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_network_device_config_physical_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references network_device_config(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_network_device_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references network_device_config(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_network_device_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references network_device_config(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_network_ipam_virtual_DNS (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references network_ipam(uuid) on delete cascade, 
    foreign key (`to`) references virtual_DNS(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_network_ipam (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references network_ipam(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_network_ipam (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references network_ipam(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_network_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references network_policy(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_network_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references network_policy(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_peering_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references peering_policy(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_peering_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references peering_policy(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_physical_interface_physical_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references physical_interface(uuid) on delete cascade, 
    foreign key (`to`) references physical_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_physical_interface (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references physical_interface(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_physical_interface (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references physical_interface(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_physical_router_virtual_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references physical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_physical_router_virtual_network (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references physical_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_physical_router_bgp_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references physical_router(uuid) on delete cascade, 
    foreign key (`to`) references bgp_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_physical_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references physical_router(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_physical_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references physical_router(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_policy_management (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references policy_management(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_policy_management (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references policy_management(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_port_tuple (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references port_tuple(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_port_tuple (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references port_tuple(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_project_application_policy_set (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references application_policy_set(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_project_floating_ip_pool (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references floating_ip_pool(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_project_alias_ip_pool (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references alias_ip_pool(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_project_namespace (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `ip_prefix_len` int,
    `ip_prefix` varchar(255),
     foreign key (`from`) references project(uuid) on delete cascade, 
    foreign key (`to`) references namespace(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_project (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references project(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_project (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references project(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_provider_attachment_virtual_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references provider_attachment(uuid) on delete cascade, 
    foreign key (`to`) references virtual_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_provider_attachment (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references provider_attachment(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_provider_attachment (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references provider_attachment(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_qos_config_global_system_config (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references qos_config(uuid) on delete cascade, 
    foreign key (`to`) references global_system_config(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_qos_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references qos_config(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_qos_config (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references qos_config(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_qos_queue (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references qos_queue(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_qos_queue (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references qos_queue(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_route_aggregate_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `interface_type` varchar(255),
     foreign key (`from`) references route_aggregate(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_route_aggregate (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references route_aggregate(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_route_aggregate (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references route_aggregate(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_route_table (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references route_table(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_route_table (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references route_table(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_route_target (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references route_target(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_route_target (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references route_target(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_routing_instance (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references routing_instance(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_routing_instance (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references routing_instance(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_routing_policy_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `right_sequence` varchar(255),
    `left_sequence` varchar(255),
     foreign key (`from`) references routing_policy(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_routing_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references routing_policy(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_routing_policy (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references routing_policy(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_security_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references security_group(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_security_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references security_group(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_security_logging_object_security_group (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `rule` json,
     foreign key (`from`) references security_logging_object(uuid) on delete cascade, 
    foreign key (`to`) references security_group(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_security_logging_object_network_policy (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `rule` json,
     foreign key (`from`) references security_logging_object(uuid) on delete cascade, 
    foreign key (`to`) references network_policy(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_security_logging_object (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references security_logging_object(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_security_logging_object (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references security_logging_object(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_service_appliance_physical_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `interface_type` varchar(255),
     foreign key (`from`) references service_appliance(uuid) on delete cascade, 
    foreign key (`to`) references physical_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_service_appliance (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_appliance(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_appliance (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_appliance(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_service_appliance_set (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_appliance_set(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_appliance_set (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_appliance_set(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_service_connection_module_service_object (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references service_connection_module(uuid) on delete cascade, 
    foreign key (`to`) references service_object(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_service_connection_module (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_connection_module(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_connection_module (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_connection_module(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_service_endpoint_service_object (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references service_endpoint(uuid) on delete cascade, 
    foreign key (`to`) references service_object(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_service_endpoint_service_connection_module (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references service_endpoint(uuid) on delete cascade, 
    foreign key (`to`) references service_connection_module(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_service_endpoint_physical_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references service_endpoint(uuid) on delete cascade, 
    foreign key (`to`) references physical_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_service_endpoint (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_endpoint(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_endpoint (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_endpoint(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_service_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_group(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_group(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_service_health_check_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `interface_type` varchar(255),
     foreign key (`from`) references service_health_check(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_service_health_check (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_health_check(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_health_check (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_health_check(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_service_instance_service_template (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references service_instance(uuid) on delete cascade, 
    foreign key (`to`) references service_template(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_service_instance_instance_ip (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `interface_type` varchar(255),
     foreign key (`from`) references service_instance(uuid) on delete cascade, 
    foreign key (`to`) references instance_ip(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_service_instance (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_instance(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_instance (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_instance(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_service_object (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_object(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_object (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_object(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_service_template_service_appliance_set (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references service_template(uuid) on delete cascade, 
    foreign key (`to`) references service_appliance_set(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_service_template (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_template(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_service_template (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references service_template(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_subnet_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references subnet(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_subnet (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references subnet(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_subnet (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references subnet(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_tag_tag_type (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references tag(uuid) on delete cascade, 
    foreign key (`to`) references tag_type(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_tag (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references tag(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_tag (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references tag(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_tag_type (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references tag_type(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_tag_type (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references tag_type(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_user (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references user(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_user (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references user(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_virtual_DNS_record (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_DNS_record(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_DNS_record (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_DNS_record(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_virtual_DNS (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_DNS(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_DNS (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_DNS(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_virtual_ip_loadbalancer_pool (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_ip(uuid) on delete cascade, 
    foreign key (`to`) references loadbalancer_pool(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_ip_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_ip(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_virtual_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_ip(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_ip (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_ip(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_virtual_machine_interface_bgp_router (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references bgp_router(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_routing_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `ipv6_service_chain_address` varchar(255),
    `direction` varchar(255),
    `mpls_label` int,
    `vlan_tag` int,
    `src_mac` varchar(255),
    `service_chain_address` varchar(255),
    `dst_mac` varchar(255),
    `protocol` varchar(255),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references routing_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_port_tuple (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references port_tuple(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_security_group (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references security_group(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_virtual_machine_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_security_logging_object (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_virtual_network (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_virtual_machine (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_bridge_domain (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `vlan_tag` int,
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references bridge_domain(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_service_endpoint (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references service_endpoint(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_interface_route_table (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references interface_route_table(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_qos_config (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references qos_config(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_physical_interface (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references physical_interface(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_machine_interface_service_health_check (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine_interface(uuid) on delete cascade, 
    foreign key (`to`) references service_health_check(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_virtual_machine_interface (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_machine_interface(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_machine_interface (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_machine_interface(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_virtual_machine_service_instance (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_machine(uuid) on delete cascade, 
    foreign key (`to`) references service_instance(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_virtual_machine (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_machine(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_machine (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_machine(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_virtual_network_qos_config (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references qos_config(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_network_route_table (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references route_table(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_network_virtual_network (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references virtual_network(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_network_bgpvpn (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references bgpvpn(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_network_network_ipam (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `ipam_subnets` json,
    `route` json,
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references network_ipam(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_network_security_logging_object (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references security_logging_object(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_network_network_policy (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `start_time` varchar(255),
    `off_interval` varchar(255),
    `on_interval` varchar(255),
    `end_time` varchar(255),
    `minor` int,
    `major` int,
     foreign key (`from`) references virtual_network(uuid) on delete cascade, 
    foreign key (`to`) references network_policy(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_virtual_network (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_network(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_network (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_network(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_virtual_router_network_ipam (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
    `subnet` json,
    `allocation_pools` json,
     foreign key (`from`) references virtual_router(uuid) on delete cascade, 
    foreign key (`to`) references network_ipam(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;

create table ref_virtual_router_virtual_machine (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references virtual_router(uuid) on delete cascade, 
    foreign key (`to`) references virtual_machine(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_virtual_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_router(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_virtual_router (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references virtual_router(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_appformix_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references appformix_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_appformix_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references appformix_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_appformix_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references appformix_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_baremetal_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references baremetal_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_baremetal_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references baremetal_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_baremetal_port (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references baremetal_port(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_baremetal_port (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references baremetal_port(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_analytics_database_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_analytics_database_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_analytics_database_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_analytics_database_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_analytics_database_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_analytics_database_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_analytics_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_analytics_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_analytics_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_analytics_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_analytics_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_analytics_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_contrail_cluster (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_cluster(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_cluster (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_cluster(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_config_database_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_config_database_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_config_database_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_config_database_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_config_database_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_config_database_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_config_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_config_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_config_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_config_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_config_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_config_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_control_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_control_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_control_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_control_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_control_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_control_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_storage_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_storage_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_storage_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_storage_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_storage_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_storage_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_vrouter_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_vrouter_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_vrouter_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_vrouter_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_vrouter_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_vrouter_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_contrail_controller_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references contrail_controller_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_contrail_controller_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_controller_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_contrail_controller_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references contrail_controller_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_dashboard (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references dashboard(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_dashboard (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references dashboard(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_flavor (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references flavor(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_flavor (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references flavor(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_os_image (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references os_image(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_os_image (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references os_image(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_keypair (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references keypair(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_keypair (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references keypair(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_kubernetes_master_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references kubernetes_master_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_kubernetes_master_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references kubernetes_master_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_kubernetes_master_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references kubernetes_master_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_kubernetes_node_node (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references kubernetes_node(uuid) on delete cascade, 
    foreign key (`to`) references node(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_kubernetes_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references kubernetes_node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_kubernetes_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references kubernetes_node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_location (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references location(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_location (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references location(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references node(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_node (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references node(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_server (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references server(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_server (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references server(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;




create table ref_vpn_group_location (
    `from` varchar(255),
    `to` varchar(255),
    primary key (`from`,`to`),
     foreign key (`from`) references vpn_group(uuid) on delete cascade, 
    foreign key (`to`) references location(uuid),
    index from_index (`from`)) CHARACTER SET utf8mb4;


create table tenant_share_vpn_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references vpn_group(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_vpn_group (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references vpn_group(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;





create table tenant_share_widget (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references widget(uuid) on delete cascade,
    foreign key (`to`) references project(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;

create table domain_share_widget (
    `uuid` varchar(255),
    `to` varchar(255),
    primary key (`uuid`,`to`),
    `access` integer,
    foreign key (`uuid`) references widget(uuid) on delete cascade,
    foreign key (`to`) references domain(uuid) on delete cascade,
    index uuid_index (`uuid`),
    index to_index (`to`)
    ) CHARACTER SET utf8mb4;


