create table metadata (
    "uuid" varchar(255),
    "type" varchar(255),
    "fq_name" varchar(255) unique,
    primary key ("uuid"));

create index fq_name_index on metadata ("fq_name");




create table "access_control_list" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "access_control_list_hash" int,
    "dynamic" bool,
    "acl_rule" json,
     primary key("uuid"));

create index access_control_list_parent_uuid_index on "access_control_list" ("parent_uuid");


create table "address_group" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "subnet" json,
     primary key("uuid"));

create index address_group_parent_uuid_index on "address_group" ("parent_uuid");


create table "alarm" (
    "uve_key" json,
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "alarm_severity" int,
    "or_list" json,
     primary key("uuid"));

create index alarm_parent_uuid_index on "alarm" ("parent_uuid");


create table "alias_ip_pool" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index alias_ip_pool_parent_uuid_index on "alias_ip_pool" ("parent_uuid");


create table "alias_ip" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "alias_ip_address_family" varchar(255),
    "alias_ip_address" varchar(255),
     primary key("uuid"));

create index alias_ip_parent_uuid_index on "alias_ip" ("parent_uuid");


create table "analytics_node" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "analytics_node_ip_address" varchar(255),
     primary key("uuid"));

create index analytics_node_parent_uuid_index on "analytics_node" ("parent_uuid");


create table "api_access_list" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "rbac_rule" json,
    "key_value_pair" json,
     primary key("uuid"));

create index api_access_list_parent_uuid_index on "api_access_list" ("parent_uuid");


create table "application_policy_set" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "all_applications" bool,
     primary key("uuid"));

create index application_policy_set_parent_uuid_index on "application_policy_set" ("parent_uuid");


create table "bgp_as_a_service" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "bgpaas_suppress_route_advertisement" bool,
    "bgpaas_shared" bool,
    "bgpaas_session_attributes" varchar(255),
    "bgpaas_ipv4_mapped_ipv6_nexthop" bool,
    "bgpaas_ip_address" varchar(255),
    "autonomous_system" int,
    "key_value_pair" json,
     primary key("uuid"));

create index bgp_as_a_service_parent_uuid_index on "bgp_as_a_service" ("parent_uuid");


create table "bgp_router" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index bgp_router_parent_uuid_index on "bgp_router" ("parent_uuid");


create table "bgpvpn" (
    "uuid" varchar(255),
    "route_target" json,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "import_route_target_list_route_target" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "export_route_target_list_route_target" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "bgpvpn_type" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index bgpvpn_parent_uuid_index on "bgpvpn" ("parent_uuid");


create table "bridge_domain" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "mac_move_time_window" int,
    "mac_move_limit_action" varchar(255),
    "mac_move_limit" int,
    "mac_limit_action" varchar(255),
    "mac_limit" int,
    "mac_learning_enabled" bool,
    "mac_aging_time" int,
    "isid" int,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index bridge_domain_parent_uuid_index on "bridge_domain" ("parent_uuid");


create table "config_node" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "config_node_ip_address" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index config_node_parent_uuid_index on "config_node" ("parent_uuid");


create table "config_root" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index config_root_parent_uuid_index on "config_root" ("parent_uuid");


create table "customer_attachment" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index customer_attachment_parent_uuid_index on "customer_attachment" ("parent_uuid");


create table "database_node" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "database_node_ip_address" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index database_node_parent_uuid_index on "database_node" ("parent_uuid");


create table "discovery_service_assignment" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index discovery_service_assignment_parent_uuid_index on "discovery_service_assignment" ("parent_uuid");


create table "domain" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "virtual_network_limit" int,
    "security_group_limit" int,
    "project_limit" int,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index domain_parent_uuid_index on "domain" ("parent_uuid");


create table "dsa_rule" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "subscriber" json,
    "ep_version" varchar(255),
    "ep_type" varchar(255),
    "ip_prefix_len" int,
    "ip_prefix" varchar(255),
    "ep_id" varchar(255),
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index dsa_rule_parent_uuid_index on "dsa_rule" ("parent_uuid");


create table "e2_service_provider" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "e2_service_provider_promiscuous" bool,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index e2_service_provider_parent_uuid_index on "e2_service_provider" ("parent_uuid");


create table "firewall_policy" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index firewall_policy_parent_uuid_index on "firewall_policy" ("parent_uuid");


create table "firewall_rule" (
    "uuid" varchar(255),
    "start_port" int,
    "end_port" int,
    "protocol_id" int,
    "protocol" varchar(255),
    "dst_ports_start_port" int,
    "dst_ports_end_port" int,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "tag_list" json,
    "tag_type" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "virtual_network" varchar(255),
    "tags" json,
    "tag_ids" json,
    "ip_prefix_len" int,
    "ip_prefix" varchar(255),
    "any" bool,
    "address_group" varchar(255),
    "endpoint_1_virtual_network" varchar(255),
    "endpoint_1_tags" json,
    "endpoint_1_tag_ids" json,
    "subnet_ip_prefix_len" int,
    "subnet_ip_prefix" varchar(255),
    "endpoint_1_any" bool,
    "endpoint_1_address_group" varchar(255),
    "display_name" varchar(255),
    "direction" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "simple_action" varchar(255),
    "qos_action" varchar(255),
    "udp_port" int,
    "vtep_dst_mac_address" varchar(255),
    "vtep_dst_ip_address" varchar(255),
    "vni" int,
    "routing_instance" varchar(255),
    "nic_assisted_mirroring_vlan" int,
    "nic_assisted_mirroring" bool,
    "nh_mode" varchar(255),
    "juniper_header" bool,
    "encapsulation" varchar(255),
    "analyzer_name" varchar(255),
    "analyzer_mac_address" varchar(255),
    "analyzer_ip_address" varchar(255),
    "log" bool,
    "gateway_name" varchar(255),
    "assign_routing_instance" varchar(255),
    "apply_service" json,
    "alert" bool,
     primary key("uuid"));

create index firewall_rule_parent_uuid_index on "firewall_rule" ("parent_uuid");


create table "floating_ip_pool" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "subnet_uuid" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index floating_ip_pool_parent_uuid_index on "floating_ip_pool" ("parent_uuid");


create table "floating_ip" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "floating_ip_traffic_direction" varchar(255),
    "port_mappings" json,
    "floating_ip_port_mappings_enable" bool,
    "floating_ip_is_virtual_ip" bool,
    "floating_ip_fixed_ip_address" varchar(255),
    "floating_ip_address_family" varchar(255),
    "floating_ip_address" varchar(255),
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index floating_ip_parent_uuid_index on "floating_ip" ("parent_uuid");


create table "forwarding_class" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "forwarding_class_vlan_priority" int,
    "forwarding_class_mpls_exp" int,
    "forwarding_class_id" int,
    "forwarding_class_dscp" int,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index forwarding_class_parent_uuid_index on "forwarding_class" ("parent_uuid");


create table "global_qos_config" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "dns" int,
    "control" int,
    "analytics" int,
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index global_qos_config_parent_uuid_index on "global_qos_config" ("parent_uuid");


create table "global_system_config" (
    "uuid" varchar(255),
    "statlist" json,
    "plugin_property" json,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "mac_move_time_window" int,
    "mac_move_limit_action" varchar(255),
    "mac_move_limit" int,
    "mac_limit_action" varchar(255),
    "mac_limit" int,
    "mac_aging_time" int,
    "subnet" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "ibgp_auto_mesh" bool,
    "xmpp_helper_enable" bool,
    "restart_time" int,
    "long_lived_restart_time" int,
    "end_of_rib_timeout" int,
    "graceful_restart_parameters_enable" bool,
    "bgp_helper_enable" bool,
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "config_version" varchar(255),
    "port_start" int,
    "port_end" int,
    "bgp_always_compare_med" bool,
    "autonomous_system" int,
    "key_value_pair" json,
    "alarm_enable" bool,
     primary key("uuid"));

create index global_system_config_parent_uuid_index on "global_system_config" ("parent_uuid");


create table "global_vrouter_config" (
    "vxlan_network_identifier_mode" varchar(255),
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "linklocal_service_entry" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "forwarding_mode" varchar(255),
    "flow_export_rate" int,
    "flow_aging_timeout" json,
    "encapsulation" json,
    "enable_security_logging" bool,
    "source_port" bool,
    "source_ip" bool,
    "ip_protocol" bool,
    "hashing_configured" bool,
    "destination_port" bool,
    "destination_ip" bool,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index global_vrouter_config_parent_uuid_index on "global_vrouter_config" ("parent_uuid");


create table "instance_ip" (
    "uuid" varchar(255),
    "subnet_uuid" varchar(255),
    "service_instance_ip" bool,
    "service_health_check_ip" bool,
    "ip_prefix_len" int,
    "ip_prefix" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "instance_ip_secondary" bool,
    "instance_ip_mode" varchar(255),
    "instance_ip_local_ip" bool,
    "instance_ip_family" varchar(255),
    "instance_ip_address" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index instance_ip_parent_uuid_index on "instance_ip" ("parent_uuid");


create table "interface_route_table" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "route" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index interface_route_table_parent_uuid_index on "interface_route_table" ("parent_uuid");


create table "loadbalancer_healthmonitor" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "url_path" varchar(255),
    "timeout" int,
    "monitor_type" varchar(255),
    "max_retries" int,
    "http_method" varchar(255),
    "expected_codes" varchar(255),
    "delay" int,
    "admin_state" bool,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index loadbalancer_healthmonitor_parent_uuid_index on "loadbalancer_healthmonitor" ("parent_uuid");


create table "loadbalancer_listener" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "sni_containers" json,
    "protocol_port" int,
    "protocol" varchar(255),
    "default_tls_container" varchar(255),
    "connection_limit" int,
    "admin_state" bool,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index loadbalancer_listener_parent_uuid_index on "loadbalancer_listener" ("parent_uuid");


create table "loadbalancer_member" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "weight" int,
    "status_description" varchar(255),
    "status" varchar(255),
    "protocol_port" int,
    "admin_state" bool,
    "address" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index loadbalancer_member_parent_uuid_index on "loadbalancer_member" ("parent_uuid");


create table "loadbalancer_pool" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "loadbalancer_pool_provider" varchar(255),
    "subnet_id" varchar(255),
    "status_description" varchar(255),
    "status" varchar(255),
    "session_persistence" varchar(255),
    "protocol" varchar(255),
    "persistence_cookie_name" varchar(255),
    "loadbalancer_method" varchar(255),
    "admin_state" bool,
    "key_value_pair" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "annotations_key_value_pair" json,
     primary key("uuid"));

create index loadbalancer_pool_parent_uuid_index on "loadbalancer_pool" ("parent_uuid");


create table "loadbalancer" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "loadbalancer_provider" varchar(255),
    "vip_subnet_id" varchar(255),
    "vip_address" varchar(255),
    "status" varchar(255),
    "provisioning_status" varchar(255),
    "operating_status" varchar(255),
    "admin_state" bool,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index loadbalancer_parent_uuid_index on "loadbalancer" ("parent_uuid");


create table "logical_interface" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "logical_interface_vlan_tag" int,
    "logical_interface_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index logical_interface_parent_uuid_index on "logical_interface" ("parent_uuid");


create table "logical_router" (
    "vxlan_network_identifier" varchar(255),
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "route_target" json,
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index logical_router_parent_uuid_index on "logical_router" ("parent_uuid");


create table "namespace" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "ip_prefix_len" int,
    "ip_prefix" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index namespace_parent_uuid_index on "namespace" ("parent_uuid");


create table "network_device_config" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index network_device_config_parent_uuid_index on "network_device_config" ("parent_uuid");


create table "network_ipam" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "ipam_method" varchar(255),
    "virtual_dns_server_name" varchar(255),
    "ip_address" varchar(255),
    "ipam_dns_method" varchar(255),
    "route" json,
    "dhcp_option" json,
    "ip_prefix_len" int,
    "ip_prefix" varchar(255),
    "subnets" json,
    "ipam_subnet_method" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index network_ipam_parent_uuid_index on "network_ipam" ("parent_uuid");


create table "network_policy" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "policy_rule" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index network_policy_parent_uuid_index on "network_policy" ("parent_uuid");


create table "peering_policy" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "peering_service" varchar(255),
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index peering_policy_parent_uuid_index on "peering_policy" ("parent_uuid");


create table "physical_interface" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "ethernet_segment_identifier" varchar(255),
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index physical_interface_parent_uuid_index on "physical_interface" ("parent_uuid");


create table "physical_router" (
    "uuid" varchar(255),
    "server_port" int,
    "server_ip" varchar(255),
    "resource" json,
    "physical_router_vnc_managed" bool,
    "physical_router_vendor_name" varchar(255),
    "username" varchar(255),
    "password" varchar(255),
    "version" int,
    "v3_security_name" varchar(255),
    "v3_security_level" varchar(255),
    "v3_security_engine_id" varchar(255),
    "v3_privacy_protocol" varchar(255),
    "v3_privacy_password" varchar(255),
    "v3_engine_time" int,
    "v3_engine_id" varchar(255),
    "v3_engine_boots" int,
    "v3_context_engine_id" varchar(255),
    "v3_context" varchar(255),
    "v3_authentication_protocol" varchar(255),
    "v3_authentication_password" varchar(255),
    "v2_community" varchar(255),
    "timeout" int,
    "retries" int,
    "local_port" int,
    "physical_router_snmp" bool,
    "physical_router_role" varchar(255),
    "physical_router_product_name" varchar(255),
    "physical_router_management_ip" varchar(255),
    "physical_router_loopback_ip" varchar(255),
    "physical_router_lldp" bool,
    "service_port" json,
    "physical_router_image_uri" varchar(255),
    "physical_router_dataplane_ip" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index physical_router_parent_uuid_index on "physical_router" ("parent_uuid");


create table "policy_management" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index policy_management_parent_uuid_index on "policy_management" ("parent_uuid");


create table "port_tuple" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index port_tuple_parent_uuid_index on "port_tuple" ("parent_uuid");


create table "project" (
    "vxlan_routing" bool,
    "uuid" varchar(255),
    "virtual_router" int,
    "virtual_network" int,
    "virtual_machine_interface" int,
    "virtual_ip" int,
    "virtual_DNS_record" int,
    "virtual_DNS" int,
    "subnet" int,
    "service_template" int,
    "service_instance" int,
    "security_logging_object" int,
    "security_group_rule" int,
    "security_group" int,
    "route_table" int,
    "network_policy" int,
    "network_ipam" int,
    "logical_router" int,
    "loadbalancer_pool" int,
    "loadbalancer_member" int,
    "loadbalancer_healthmonitor" int,
    "instance_ip" int,
    "global_vrouter_config" int,
    "floating_ip_pool" int,
    "floating_ip" int,
    "defaults" int,
    "bgp_router" int,
    "access_control_list" int,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "alarm_enable" bool,
     primary key("uuid"));

create index project_parent_uuid_index on "project" ("parent_uuid");


create table "provider_attachment" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index provider_attachment_parent_uuid_index on "provider_attachment" ("parent_uuid");


create table "qos_config" (
    "qos_id_forwarding_class_pair" json,
    "uuid" varchar(255),
    "qos_config_type" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "mpls_exp_entries_qos_id_forwarding_class_pair" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "dscp_entries_qos_id_forwarding_class_pair" json,
    "display_name" varchar(255),
    "default_forwarding_class_id" int,
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index qos_config_parent_uuid_index on "qos_config" ("parent_uuid");


create table "qos_queue" (
    "uuid" varchar(255),
    "qos_queue_identifier" int,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "min_bandwidth" int,
    "max_bandwidth" int,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index qos_queue_parent_uuid_index on "qos_queue" ("parent_uuid");


create table "route_aggregate" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index route_aggregate_parent_uuid_index on "route_aggregate" ("parent_uuid");


create table "route_table" (
    "uuid" varchar(255),
    "route" json,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index route_table_parent_uuid_index on "route_table" ("parent_uuid");


create table "route_target" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index route_target_parent_uuid_index on "route_target" ("parent_uuid");


create table "routing_instance" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index routing_instance_parent_uuid_index on "routing_instance" ("parent_uuid");


create table "routing_policy" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index routing_policy_parent_uuid_index on "routing_policy" ("parent_uuid");


create table "security_group" (
    "uuid" varchar(255),
    "security_group_id" int,
    "policy_rule" json,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configured_security_group_id" int,
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index security_group_parent_uuid_index on "security_group" ("parent_uuid");


create table "security_logging_object" (
    "uuid" varchar(255),
    "rule" json,
    "security_logging_object_rate" int,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index security_logging_object_parent_uuid_index on "security_logging_object" ("parent_uuid");


create table "service_appliance" (
    "uuid" varchar(255),
    "username" varchar(255),
    "password" varchar(255),
    "key_value_pair" json,
    "service_appliance_ip_address" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "annotations_key_value_pair" json,
     primary key("uuid"));

create index service_appliance_parent_uuid_index on "service_appliance" ("parent_uuid");


create table "service_appliance_set" (
    "uuid" varchar(255),
    "key_value_pair" json,
    "service_appliance_ha_mode" varchar(255),
    "service_appliance_driver" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "annotations_key_value_pair" json,
     primary key("uuid"));

create index service_appliance_set_parent_uuid_index on "service_appliance_set" ("parent_uuid");


create table "service_connection_module" (
    "uuid" varchar(255),
    "service_type" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "e2_service" varchar(255),
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index service_connection_module_parent_uuid_index on "service_connection_module" ("parent_uuid");


create table "service_endpoint" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index service_endpoint_parent_uuid_index on "service_endpoint" ("parent_uuid");


create table "service_group" (
    "uuid" varchar(255),
    "firewall_service" json,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index service_group_parent_uuid_index on "service_group" ("parent_uuid");


create table "service_health_check" (
    "uuid" varchar(255),
    "url_path" varchar(255),
    "timeoutUsecs" int,
    "timeout" int,
    "monitor_type" varchar(255),
    "max_retries" int,
    "http_method" varchar(255),
    "health_check_type" varchar(255),
    "expected_codes" varchar(255),
    "enabled" bool,
    "delayUsecs" int,
    "delay" int,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index service_health_check_parent_uuid_index on "service_health_check" ("parent_uuid");


create table "service_instance" (
    "uuid" varchar(255),
    "virtual_router_id" varchar(255),
    "max_instances" int,
    "auto_scale" bool,
    "right_virtual_network" varchar(255),
    "right_ip_address" varchar(255),
    "management_virtual_network" varchar(255),
    "left_virtual_network" varchar(255),
    "left_ip_address" varchar(255),
    "interface_list" json,
    "ha_mode" varchar(255),
    "availability_zone" varchar(255),
    "auto_policy" bool,
    "key_value_pair" json,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "annotations_key_value_pair" json,
     primary key("uuid"));

create index service_instance_parent_uuid_index on "service_instance" ("parent_uuid");


create table "service_object" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index service_object_parent_uuid_index on "service_object" ("parent_uuid");


create table "service_template" (
    "uuid" varchar(255),
    "vrouter_instance_type" varchar(255),
    "version" int,
    "service_virtualization_type" varchar(255),
    "service_type" varchar(255),
    "service_scaling" bool,
    "service_mode" varchar(255),
    "ordered_interfaces" bool,
    "interface_type" json,
    "instance_data" varchar(255),
    "image_name" varchar(255),
    "flavor" varchar(255),
    "availability_zone_enable" bool,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index service_template_parent_uuid_index on "service_template" ("parent_uuid");


create table "subnet" (
    "uuid" varchar(255),
    "ip_prefix_len" int,
    "ip_prefix" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index subnet_parent_uuid_index on "subnet" ("parent_uuid");


create table "tag" (
    "uuid" varchar(255),
    "tag_value" varchar(255),
    "tag_type_name" varchar(255),
    "tag_id" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index tag_parent_uuid_index on "tag" ("parent_uuid");


create table "tag_type" (
    "uuid" varchar(255),
    "tag_type_id" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index tag_type_parent_uuid_index on "tag_type" ("parent_uuid");


create table "user" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "password" varchar(255),
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index user_parent_uuid_index on "user" ("parent_uuid");


create table "virtual_DNS_record" (
    "record_type" varchar(255),
    "record_ttl_seconds" int,
    "record_name" varchar(255),
    "record_mx_preference" int,
    "record_data" varchar(255),
    "record_class" varchar(255),
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index virtual_DNS_record_parent_uuid_index on "virtual_DNS_record" ("parent_uuid");


create table "virtual_DNS" (
    "reverse_resolution" bool,
    "record_order" varchar(255),
    "next_virtual_DNS" varchar(255),
    "floating_ip_record" varchar(255),
    "external_visible" bool,
    "dynamic_records_from_client" bool,
    "domain_name" varchar(255),
    "default_ttl_seconds" int,
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index virtual_DNS_parent_uuid_index on "virtual_DNS" ("parent_uuid");


create table "virtual_ip" (
    "subnet_id" varchar(255),
    "status_description" varchar(255),
    "status" varchar(255),
    "protocol_port" int,
    "protocol" varchar(255),
    "persistence_type" varchar(255),
    "persistence_cookie_name" varchar(255),
    "connection_limit" int,
    "admin_state" bool,
    "address" varchar(255),
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index virtual_ip_parent_uuid_index on "virtual_ip" ("parent_uuid");


create table "virtual_machine_interface" (
    "vrf_assign_rule" json,
    "vlan_tag_based_bridge_domain" bool,
    "sub_interface_vlan_tag" int,
    "service_interface_type" varchar(255),
    "local_preference" int,
    "traffic_direction" varchar(255),
    "udp_port" int,
    "vtep_dst_mac_address" varchar(255),
    "vtep_dst_ip_address" varchar(255),
    "vni" int,
    "routing_instance" varchar(255),
    "nic_assisted_mirroring_vlan" int,
    "nic_assisted_mirroring" bool,
    "nh_mode" varchar(255),
    "juniper_header" bool,
    "encapsulation" varchar(255),
    "analyzer_name" varchar(255),
    "analyzer_mac_address" varchar(255),
    "analyzer_ip_address" varchar(255),
    "mac_address" json,
    "route" json,
    "fat_flow_protocol" json,
    "virtual_machine_interface_disable_policy" bool,
    "dhcp_option" json,
    "virtual_machine_interface_device_owner" varchar(255),
    "key_value_pair" json,
    "allowed_address_pair" json,
    "uuid" varchar(255),
    "port_security_enabled" bool,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "source_port" bool,
    "source_ip" bool,
    "ip_protocol" bool,
    "hashing_configured" bool,
    "destination_port" bool,
    "destination_ip" bool,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "annotations_key_value_pair" json,
     primary key("uuid"));

create index virtual_machine_interface_parent_uuid_index on "virtual_machine_interface" ("parent_uuid");


create table "virtual_machine" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index virtual_machine_parent_uuid_index on "virtual_machine" ("parent_uuid");


create table "virtual_network" (
    "vxlan_network_identifier" int,
    "rpf" varchar(255),
    "network_id" int,
    "mirror_destination" bool,
    "forwarding_mode" varchar(255),
    "allow_transit" bool,
    "virtual_network_network_id" int,
    "uuid" varchar(255),
    "router_external" bool,
    "route_target" json,
    "segmentation_id" int,
    "physical_network" varchar(255),
    "port_security_enabled" bool,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "pbb_evpn_enable" bool,
    "pbb_etree_enable" bool,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "multi_policy_service_chains_enabled" bool,
    "mac_move_time_window" int,
    "mac_move_limit_action" varchar(255),
    "mac_move_limit" int,
    "mac_limit_action" varchar(255),
    "mac_limit" int,
    "mac_learning_enabled" bool,
    "mac_aging_time" int,
    "layer2_control_word" bool,
    "is_shared" bool,
    "import_route_target_list_route_target" json,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "flood_unknown_unicast" bool,
    "external_ipam" bool,
    "export_route_target_list_route_target" json,
    "source_port" bool,
    "source_ip" bool,
    "ip_protocol" bool,
    "hashing_configured" bool,
    "destination_port" bool,
    "destination_ip" bool,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
    "address_allocation_mode" varchar(255),
     primary key("uuid"));

create index virtual_network_parent_uuid_index on "virtual_network" ("parent_uuid");


create table "virtual_router" (
    "virtual_router_type" varchar(255),
    "virtual_router_ip_address" varchar(255),
    "virtual_router_dpdk_enabled" bool,
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index virtual_router_parent_uuid_index on "virtual_router" ("parent_uuid");


create table "appformix_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index appformix_node_parent_uuid_index on "appformix_node" ("parent_uuid");


create table "baremetal_node" (
    "uuid" varchar(255),
    "updated_at" varchar(255),
    "target_provision_state" varchar(255),
    "target_power_state" varchar(255),
    "provision_state" varchar(255),
    "power_state" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "name" varchar(255),
    "maintenance_reason" varchar(255),
    "maintenance" bool,
    "last_error" varchar(255),
    "instance_uuid" varchar(255),
    "vcpus" varchar(255),
    "swap_mb" varchar(255),
    "root_gb" varchar(255),
    "nova_host_id" varchar(255),
    "memory_mb" varchar(255),
    "local_gb" varchar(255),
    "image_source" varchar(255),
    "display_name" varchar(255),
    "capabilities" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "ipmi_username" varchar(255),
    "ipmi_password" varchar(255),
    "ipmi_address" varchar(255),
    "deploy_ramdisk" varchar(255),
    "deploy_kernel" varchar(255),
    "_display_name" varchar(255),
    "created_at" varchar(255),
    "console_enabled" bool,
    "configuration_version" bigint,
    "bm_properties_memory_mb" int,
    "disk_gb" int,
    "cpu_count" int,
    "cpu_arch" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index baremetal_node_parent_uuid_index on "baremetal_node" ("parent_uuid");


create table "baremetal_port" (
    "uuid" varchar(255),
    "updated_at" varchar(255),
    "pxe_enabled" bool,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "node" varchar(255),
    "mac_address" varchar(255),
    "switch_info" varchar(255),
    "switch_id" varchar(255),
    "port_id" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "created_at" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index baremetal_port_parent_uuid_index on "baremetal_port" ("parent_uuid");


create table "contrail_analytics_database_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_analytics_database_node_parent_uuid_index on "contrail_analytics_database_node" ("parent_uuid");


create table "contrail_analytics_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_analytics_node_parent_uuid_index on "contrail_analytics_node" ("parent_uuid");


create table "contrail_cluster" (
    "uuid" varchar(255),
    "statistics_ttl" varchar(255),
    "rabbitmq_port" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "provisioner_type" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "orchestrator" varchar(255),
    "openstack_internal_vip_interface" varchar(255),
    "openstack_internal_vip" varchar(255),
    "openstack_external_vip_interface" varchar(255),
    "openstack_external_vip" varchar(255),
    "openstack_enable_haproxy" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "flow_ttl" varchar(255),
    "display_name" varchar(255),
    "default_vrouter_bond_interface_members" varchar(255),
    "default_vrouter_bond_interface" varchar(255),
    "default_gateway" varchar(255),
    "contrail_version" varchar(255),
    "container_registry" varchar(255),
    "configuration_version" bigint,
    "config_audit_ttl" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_cluster_parent_uuid_index on "contrail_cluster" ("parent_uuid");


create table "contrail_config_database_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_config_database_node_parent_uuid_index on "contrail_config_database_node" ("parent_uuid");


create table "contrail_config_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_config_node_parent_uuid_index on "contrail_config_node" ("parent_uuid");


create table "contrail_control_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_control_node_parent_uuid_index on "contrail_control_node" ("parent_uuid");


create table "contrail_storage_node" (
    "uuid" varchar(255),
    "storage_backend_bond_interface_members" varchar(255),
    "storage_access_bond_interface_members" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "osd_drives" varchar(255),
    "journal_drives" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_storage_node_parent_uuid_index on "contrail_storage_node" ("parent_uuid");


create table "contrail_vrouter_node" (
    "vrouter_type" varchar(255),
    "vrouter_bond_interface_members" varchar(255),
    "vrouter_bond_interface" varchar(255),
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "default_gateway" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_vrouter_node_parent_uuid_index on "contrail_vrouter_node" ("parent_uuid");


create table "contrail_webui_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index contrail_webui_node_parent_uuid_index on "contrail_webui_node" ("parent_uuid");


create table "dashboard" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "container_config" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index dashboard_parent_uuid_index on "dashboard" ("parent_uuid");


create table "flavor" (
    "vcpus" int,
    "uuid" varchar(255),
    "swap" int,
    "rxtx_factor" int,
    "ram" int,
    "property" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "name" varchar(255),
    "type" varchar(255),
    "rel" varchar(255),
    "href" varchar(255),
    "is_public" bool,
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "id" varchar(255),
    "fq_name" json,
    "ephemeral" int,
    "display_name" varchar(255),
    "disk" int,
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index flavor_parent_uuid_index on "flavor" ("parent_uuid");


create table "os_image" (
    "visibility" varchar(255),
    "uuid" varchar(255),
    "updated_at" varchar(255),
    "tags" varchar(255),
    "status" varchar(255),
    "size" int,
    "protected" bool,
    "property" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "_owner" varchar(255),
    "name" varchar(255),
    "min_ram" int,
    "min_disk" int,
    "location" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "id" varchar(255),
    "fq_name" json,
    "file" varchar(255),
    "display_name" varchar(255),
    "disk_format" varchar(255),
    "created_at" varchar(255),
    "container_format" varchar(255),
    "configuration_version" bigint,
    "checksum" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index os_image_parent_uuid_index on "os_image" ("parent_uuid");


create table "keypair" (
    "uuid" varchar(255),
    "public_key" tinytext,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "name" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index keypair_parent_uuid_index on "keypair" ("parent_uuid");


create table "kubernetes_master_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index kubernetes_master_node_parent_uuid_index on "kubernetes_master_node" ("parent_uuid");


create table "kubernetes_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index kubernetes_node_parent_uuid_index on "kubernetes_node" ("parent_uuid");


create table "location" (
    "uuid" varchar(255),
    "type" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "private_redhat_subscription_user" varchar(255),
    "private_redhat_subscription_pasword" varchar(255),
    "private_redhat_subscription_key" varchar(255),
    "private_redhat_pool_id" varchar(255),
    "private_ospd_vm_vcpus" varchar(255),
    "private_ospd_vm_ram_mb" varchar(255),
    "private_ospd_vm_name" varchar(255),
    "private_ospd_vm_disk_gb" varchar(255),
    "private_ospd_user_password" varchar(255),
    "private_ospd_user_name" varchar(255),
    "private_ospd_package_url" varchar(255),
    "private_ntp_hosts" varchar(255),
    "private_dns_servers" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "gcp_subnet" varchar(255),
    "gcp_region" varchar(255),
    "gcp_asn" int,
    "gcp_account_info" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "aws_subnet" varchar(255),
    "aws_secret_key" varchar(255),
    "aws_region" varchar(255),
    "aws_access_key" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index location_parent_uuid_index on "location" ("parent_uuid");


create table "node" (
    "uuid" varchar(255),
    "username" varchar(255),
    "type" varchar(255),
    "ssh_key" text,
    "private_machine_state" varchar(255),
    "private_machine_properties" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "password" varchar(255),
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "mac_address" varchar(255),
    "ip_address" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "hostname" varchar(255),
    "gcp_machine_type" varchar(255),
    "gcp_image" varchar(255),
    "fq_name" json,
    "ipmi_username" varchar(255),
    "ipmi_password" varchar(255),
    "ipmi_address" varchar(255),
    "deploy_ramdisk" varchar(255),
    "deploy_kernel" varchar(255),
    "display_name" varchar(255),
    "configuration_version" bigint,
    "memory_mb" int,
    "disk_gb" int,
    "cpu_count" int,
    "cpu_arch" varchar(255),
    "aws_instance_type" varchar(255),
    "aws_ami" varchar(255),
    "key_value_pair" json,
     primary key("uuid"));

create index node_parent_uuid_index on "node" ("parent_uuid");


create table "openstack_compute_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index openstack_compute_node_parent_uuid_index on "openstack_compute_node" ("parent_uuid");


create table "openstack_control_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index openstack_control_node_parent_uuid_index on "openstack_control_node" ("parent_uuid");


create table "openstack_monitoring_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index openstack_monitoring_node_parent_uuid_index on "openstack_monitoring_node" ("parent_uuid");


create table "openstack_network_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index openstack_network_node_parent_uuid_index on "openstack_network_node" ("parent_uuid");


create table "openstack_storage_node" (
    "uuid" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index openstack_storage_node_parent_uuid_index on "openstack_storage_node" ("parent_uuid");


create table "port" (
    "uuid" varchar(255),
    "pxe_enabled" bool,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "node_uuid" varchar(255),
    "mac_address" varchar(255),
    "switch_info" varchar(255),
    "switch_id" varchar(255),
    "port_id" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index port_parent_uuid_index on "port" ("parent_uuid");


create table "server" (
    "uuid" varchar(255),
    "user_id" int,
    "updated" varchar(255),
    "tenant_id" varchar(255),
    "status" varchar(255),
    "progress" int,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "name" varchar(255),
    "locked" bool,
    "type" varchar(255),
    "rel" varchar(255),
    "href" varchar(255),
    "id" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "_id" varchar(255),
    "host_status" varchar(255),
    "hostId" varchar(255),
    "fq_name" json,
    "links_type" varchar(255),
    "links_rel" varchar(255),
    "links_href" varchar(255),
    "flavor_id" varchar(255),
    "display_name" varchar(255),
    "_created" varchar(255),
    "configuration_version" bigint,
    "config_drive" bool,
    "key_value_pair" json,
    "addr" varchar(255),
    "accessIPv6" varchar(255),
    "accessIPv4" varchar(255),
     primary key("uuid"));

create index server_parent_uuid_index on "server" ("parent_uuid");


create table "vpn_group" (
    "uuid" varchar(255),
    "type" varchar(255),
    "provisioning_state" varchar(255),
    "provisioning_start_time" varchar(255),
    "provisioning_progress_stage" varchar(255),
    "provisioning_progress" int,
    "provisioning_log" text,
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index vpn_group_parent_uuid_index on "vpn_group" ("parent_uuid");


create table "widget" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" int,
    "owner" varchar(255),
    "global_access" int,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "layout_config" varchar(255),
    "user_visible" bool,
    "permissions_owner_access" int,
    "permissions_owner" varchar(255),
    "other_access" int,
    "group_access" int,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "content_config" varchar(255),
    "container_config" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index widget_parent_uuid_index on "widget" ("parent_uuid");







create table tenant_share_access_control_list (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "access_control_list"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_access_control_list_id on tenant_share_access_control_list("uuid");
create index index_t_access_control_list_to on tenant_share_access_control_list("to");

create table domain_share_access_control_list (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "access_control_list"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_access_control_list_id on domain_share_access_control_list("uuid");
create index index_d_access_control_list_to on domain_share_access_control_list("to");





create table tenant_share_address_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "address_group"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_address_group_id on tenant_share_address_group("uuid");
create index index_t_address_group_to on tenant_share_address_group("to");

create table domain_share_address_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "address_group"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_address_group_id on domain_share_address_group("uuid");
create index index_d_address_group_to on domain_share_address_group("to");





create table tenant_share_alarm (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "alarm"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_alarm_id on tenant_share_alarm("uuid");
create index index_t_alarm_to on tenant_share_alarm("to");

create table domain_share_alarm (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "alarm"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_alarm_id on domain_share_alarm("uuid");
create index index_d_alarm_to on domain_share_alarm("to");





create table tenant_share_alias_ip_pool (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "alias_ip_pool"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_alias_ip_pool_id on tenant_share_alias_ip_pool("uuid");
create index index_t_alias_ip_pool_to on tenant_share_alias_ip_pool("to");

create table domain_share_alias_ip_pool (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "alias_ip_pool"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_alias_ip_pool_id on domain_share_alias_ip_pool("uuid");
create index index_d_alias_ip_pool_to on domain_share_alias_ip_pool("to");




create table ref_alias_ip_project (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "alias_ip"(uuid) on delete cascade, 
    foreign key ("to") references "project"(uuid));

create index index_alias_ip_project on ref_alias_ip_project ("from");

create table ref_alias_ip_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "alias_ip"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_alias_ip_virtual_machine_interface on ref_alias_ip_virtual_machine_interface ("from");


create table tenant_share_alias_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "alias_ip"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_alias_ip_id on tenant_share_alias_ip("uuid");
create index index_t_alias_ip_to on tenant_share_alias_ip("to");

create table domain_share_alias_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "alias_ip"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_alias_ip_id on domain_share_alias_ip("uuid");
create index index_d_alias_ip_to on domain_share_alias_ip("to");





create table tenant_share_analytics_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "analytics_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_analytics_node_id on tenant_share_analytics_node("uuid");
create index index_t_analytics_node_to on tenant_share_analytics_node("to");

create table domain_share_analytics_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "analytics_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_analytics_node_id on domain_share_analytics_node("uuid");
create index index_d_analytics_node_to on domain_share_analytics_node("to");





create table tenant_share_api_access_list (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "api_access_list"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_api_access_list_id on tenant_share_api_access_list("uuid");
create index index_t_api_access_list_to on tenant_share_api_access_list("to");

create table domain_share_api_access_list (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "api_access_list"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_api_access_list_id on domain_share_api_access_list("uuid");
create index index_d_api_access_list_to on domain_share_api_access_list("to");




create table ref_application_policy_set_firewall_policy (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "sequence" varchar(255),
     foreign key ("from") references "application_policy_set"(uuid) on delete cascade, 
    foreign key ("to") references "firewall_policy"(uuid));

create index index_application_policy_set_firewall_policy on ref_application_policy_set_firewall_policy ("from");

create table ref_application_policy_set_global_vrouter_config (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "application_policy_set"(uuid) on delete cascade, 
    foreign key ("to") references "global_vrouter_config"(uuid));

create index index_application_policy_set_global_vrouter_config on ref_application_policy_set_global_vrouter_config ("from");


create table tenant_share_application_policy_set (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "application_policy_set"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_application_policy_set_id on tenant_share_application_policy_set("uuid");
create index index_t_application_policy_set_to on tenant_share_application_policy_set("to");

create table domain_share_application_policy_set (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "application_policy_set"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_application_policy_set_id on domain_share_application_policy_set("uuid");
create index index_d_application_policy_set_to on domain_share_application_policy_set("to");




create table ref_bgp_as_a_service_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "bgp_as_a_service"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_bgp_as_a_service_virtual_machine_interface on ref_bgp_as_a_service_virtual_machine_interface ("from");

create table ref_bgp_as_a_service_service_health_check (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "bgp_as_a_service"(uuid) on delete cascade, 
    foreign key ("to") references "service_health_check"(uuid));

create index index_bgp_as_a_service_service_health_check on ref_bgp_as_a_service_service_health_check ("from");


create table tenant_share_bgp_as_a_service (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bgp_as_a_service"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_bgp_as_a_service_id on tenant_share_bgp_as_a_service("uuid");
create index index_t_bgp_as_a_service_to on tenant_share_bgp_as_a_service("to");

create table domain_share_bgp_as_a_service (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bgp_as_a_service"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_bgp_as_a_service_id on domain_share_bgp_as_a_service("uuid");
create index index_d_bgp_as_a_service_to on domain_share_bgp_as_a_service("to");





create table tenant_share_bgp_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bgp_router"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_bgp_router_id on tenant_share_bgp_router("uuid");
create index index_t_bgp_router_to on tenant_share_bgp_router("to");

create table domain_share_bgp_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bgp_router"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_bgp_router_id on domain_share_bgp_router("uuid");
create index index_d_bgp_router_to on domain_share_bgp_router("to");





create table tenant_share_bgpvpn (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bgpvpn"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_bgpvpn_id on tenant_share_bgpvpn("uuid");
create index index_t_bgpvpn_to on tenant_share_bgpvpn("to");

create table domain_share_bgpvpn (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bgpvpn"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_bgpvpn_id on domain_share_bgpvpn("uuid");
create index index_d_bgpvpn_to on domain_share_bgpvpn("to");





create table tenant_share_bridge_domain (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bridge_domain"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_bridge_domain_id on tenant_share_bridge_domain("uuid");
create index index_t_bridge_domain_to on tenant_share_bridge_domain("to");

create table domain_share_bridge_domain (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "bridge_domain"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_bridge_domain_id on domain_share_bridge_domain("uuid");
create index index_d_bridge_domain_to on domain_share_bridge_domain("to");





create table tenant_share_config_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "config_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_config_node_id on tenant_share_config_node("uuid");
create index index_t_config_node_to on tenant_share_config_node("to");

create table domain_share_config_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "config_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_config_node_id on domain_share_config_node("uuid");
create index index_d_config_node_to on domain_share_config_node("to");




create table ref_config_root_tag (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "config_root"(uuid) on delete cascade, 
    foreign key ("to") references "tag"(uuid));

create index index_config_root_tag on ref_config_root_tag ("from");


create table tenant_share_config_root (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "config_root"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_config_root_id on tenant_share_config_root("uuid");
create index index_t_config_root_to on tenant_share_config_root("to");

create table domain_share_config_root (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "config_root"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_config_root_id on domain_share_config_root("uuid");
create index index_d_config_root_to on domain_share_config_root("to");




create table ref_customer_attachment_floating_ip (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "customer_attachment"(uuid) on delete cascade, 
    foreign key ("to") references "floating_ip"(uuid));

create index index_customer_attachment_floating_ip on ref_customer_attachment_floating_ip ("from");

create table ref_customer_attachment_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "customer_attachment"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_customer_attachment_virtual_machine_interface on ref_customer_attachment_virtual_machine_interface ("from");


create table tenant_share_customer_attachment (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "customer_attachment"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_customer_attachment_id on tenant_share_customer_attachment("uuid");
create index index_t_customer_attachment_to on tenant_share_customer_attachment("to");

create table domain_share_customer_attachment (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "customer_attachment"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_customer_attachment_id on domain_share_customer_attachment("uuid");
create index index_d_customer_attachment_to on domain_share_customer_attachment("to");





create table tenant_share_database_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "database_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_database_node_id on tenant_share_database_node("uuid");
create index index_t_database_node_to on tenant_share_database_node("to");

create table domain_share_database_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "database_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_database_node_id on domain_share_database_node("uuid");
create index index_d_database_node_to on domain_share_database_node("to");





create table tenant_share_discovery_service_assignment (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "discovery_service_assignment"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_discovery_service_assignment_id on tenant_share_discovery_service_assignment("uuid");
create index index_t_discovery_service_assignment_to on tenant_share_discovery_service_assignment("to");

create table domain_share_discovery_service_assignment (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "discovery_service_assignment"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_discovery_service_assignment_id on domain_share_discovery_service_assignment("uuid");
create index index_d_discovery_service_assignment_to on domain_share_discovery_service_assignment("to");





create table tenant_share_domain (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "domain"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_domain_id on tenant_share_domain("uuid");
create index index_t_domain_to on tenant_share_domain("to");

create table domain_share_domain (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "domain"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_domain_id on domain_share_domain("uuid");
create index index_d_domain_to on domain_share_domain("to");





create table tenant_share_dsa_rule (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "dsa_rule"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_dsa_rule_id on tenant_share_dsa_rule("uuid");
create index index_t_dsa_rule_to on tenant_share_dsa_rule("to");

create table domain_share_dsa_rule (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "dsa_rule"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_dsa_rule_id on domain_share_dsa_rule("uuid");
create index index_d_dsa_rule_to on domain_share_dsa_rule("to");




create table ref_e2_service_provider_physical_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "e2_service_provider"(uuid) on delete cascade, 
    foreign key ("to") references "physical_router"(uuid));

create index index_e2_service_provider_physical_router on ref_e2_service_provider_physical_router ("from");

create table ref_e2_service_provider_peering_policy (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "e2_service_provider"(uuid) on delete cascade, 
    foreign key ("to") references "peering_policy"(uuid));

create index index_e2_service_provider_peering_policy on ref_e2_service_provider_peering_policy ("from");


create table tenant_share_e2_service_provider (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "e2_service_provider"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_e2_service_provider_id on tenant_share_e2_service_provider("uuid");
create index index_t_e2_service_provider_to on tenant_share_e2_service_provider("to");

create table domain_share_e2_service_provider (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "e2_service_provider"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_e2_service_provider_id on domain_share_e2_service_provider("uuid");
create index index_d_e2_service_provider_to on domain_share_e2_service_provider("to");




create table ref_firewall_policy_firewall_rule (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "sequence" varchar(255),
     foreign key ("from") references "firewall_policy"(uuid) on delete cascade, 
    foreign key ("to") references "firewall_rule"(uuid));

create index index_firewall_policy_firewall_rule on ref_firewall_policy_firewall_rule ("from");

create table ref_firewall_policy_security_logging_object (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "firewall_policy"(uuid) on delete cascade, 
    foreign key ("to") references "security_logging_object"(uuid));

create index index_firewall_policy_security_logging_object on ref_firewall_policy_security_logging_object ("from");


create table tenant_share_firewall_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "firewall_policy"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_firewall_policy_id on tenant_share_firewall_policy("uuid");
create index index_t_firewall_policy_to on tenant_share_firewall_policy("to");

create table domain_share_firewall_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "firewall_policy"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_firewall_policy_id on domain_share_firewall_policy("uuid");
create index index_d_firewall_policy_to on domain_share_firewall_policy("to");




create table ref_firewall_rule_address_group (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "firewall_rule"(uuid) on delete cascade, 
    foreign key ("to") references "address_group"(uuid));

create index index_firewall_rule_address_group on ref_firewall_rule_address_group ("from");

create table ref_firewall_rule_security_logging_object (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "firewall_rule"(uuid) on delete cascade, 
    foreign key ("to") references "security_logging_object"(uuid));

create index index_firewall_rule_security_logging_object on ref_firewall_rule_security_logging_object ("from");

create table ref_firewall_rule_virtual_network (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "firewall_rule"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_network"(uuid));

create index index_firewall_rule_virtual_network on ref_firewall_rule_virtual_network ("from");

create table ref_firewall_rule_service_group (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "firewall_rule"(uuid) on delete cascade, 
    foreign key ("to") references "service_group"(uuid));

create index index_firewall_rule_service_group on ref_firewall_rule_service_group ("from");


create table tenant_share_firewall_rule (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "firewall_rule"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_firewall_rule_id on tenant_share_firewall_rule("uuid");
create index index_t_firewall_rule_to on tenant_share_firewall_rule("to");

create table domain_share_firewall_rule (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "firewall_rule"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_firewall_rule_id on domain_share_firewall_rule("uuid");
create index index_d_firewall_rule_to on domain_share_firewall_rule("to");





create table tenant_share_floating_ip_pool (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "floating_ip_pool"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_floating_ip_pool_id on tenant_share_floating_ip_pool("uuid");
create index index_t_floating_ip_pool_to on tenant_share_floating_ip_pool("to");

create table domain_share_floating_ip_pool (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "floating_ip_pool"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_floating_ip_pool_id on domain_share_floating_ip_pool("uuid");
create index index_d_floating_ip_pool_to on domain_share_floating_ip_pool("to");




create table ref_floating_ip_project (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "floating_ip"(uuid) on delete cascade, 
    foreign key ("to") references "project"(uuid));

create index index_floating_ip_project on ref_floating_ip_project ("from");

create table ref_floating_ip_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "floating_ip"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_floating_ip_virtual_machine_interface on ref_floating_ip_virtual_machine_interface ("from");


create table tenant_share_floating_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "floating_ip"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_floating_ip_id on tenant_share_floating_ip("uuid");
create index index_t_floating_ip_to on tenant_share_floating_ip("to");

create table domain_share_floating_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "floating_ip"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_floating_ip_id on domain_share_floating_ip("uuid");
create index index_d_floating_ip_to on domain_share_floating_ip("to");




create table ref_forwarding_class_qos_queue (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "forwarding_class"(uuid) on delete cascade, 
    foreign key ("to") references "qos_queue"(uuid));

create index index_forwarding_class_qos_queue on ref_forwarding_class_qos_queue ("from");


create table tenant_share_forwarding_class (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "forwarding_class"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_forwarding_class_id on tenant_share_forwarding_class("uuid");
create index index_t_forwarding_class_to on tenant_share_forwarding_class("to");

create table domain_share_forwarding_class (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "forwarding_class"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_forwarding_class_id on domain_share_forwarding_class("uuid");
create index index_d_forwarding_class_to on domain_share_forwarding_class("to");





create table tenant_share_global_qos_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "global_qos_config"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_global_qos_config_id on tenant_share_global_qos_config("uuid");
create index index_t_global_qos_config_to on tenant_share_global_qos_config("to");

create table domain_share_global_qos_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "global_qos_config"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_global_qos_config_id on domain_share_global_qos_config("uuid");
create index index_d_global_qos_config_to on domain_share_global_qos_config("to");




create table ref_global_system_config_bgp_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "global_system_config"(uuid) on delete cascade, 
    foreign key ("to") references "bgp_router"(uuid));

create index index_global_system_config_bgp_router on ref_global_system_config_bgp_router ("from");


create table tenant_share_global_system_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "global_system_config"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_global_system_config_id on tenant_share_global_system_config("uuid");
create index index_t_global_system_config_to on tenant_share_global_system_config("to");

create table domain_share_global_system_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "global_system_config"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_global_system_config_id on domain_share_global_system_config("uuid");
create index index_d_global_system_config_to on domain_share_global_system_config("to");





create table tenant_share_global_vrouter_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "global_vrouter_config"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_global_vrouter_config_id on tenant_share_global_vrouter_config("uuid");
create index index_t_global_vrouter_config_to on tenant_share_global_vrouter_config("to");

create table domain_share_global_vrouter_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "global_vrouter_config"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_global_vrouter_config_id on domain_share_global_vrouter_config("uuid");
create index index_d_global_vrouter_config_to on domain_share_global_vrouter_config("to");




create table ref_instance_ip_network_ipam (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "instance_ip"(uuid) on delete cascade, 
    foreign key ("to") references "network_ipam"(uuid));

create index index_instance_ip_network_ipam on ref_instance_ip_network_ipam ("from");

create table ref_instance_ip_virtual_network (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "instance_ip"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_network"(uuid));

create index index_instance_ip_virtual_network on ref_instance_ip_virtual_network ("from");

create table ref_instance_ip_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "instance_ip"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_instance_ip_virtual_machine_interface on ref_instance_ip_virtual_machine_interface ("from");

create table ref_instance_ip_physical_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "instance_ip"(uuid) on delete cascade, 
    foreign key ("to") references "physical_router"(uuid));

create index index_instance_ip_physical_router on ref_instance_ip_physical_router ("from");

create table ref_instance_ip_virtual_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "instance_ip"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_router"(uuid));

create index index_instance_ip_virtual_router on ref_instance_ip_virtual_router ("from");


create table tenant_share_instance_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "instance_ip"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_instance_ip_id on tenant_share_instance_ip("uuid");
create index index_t_instance_ip_to on tenant_share_instance_ip("to");

create table domain_share_instance_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "instance_ip"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_instance_ip_id on domain_share_instance_ip("uuid");
create index index_d_instance_ip_to on domain_share_instance_ip("to");




create table ref_interface_route_table_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "interface_type" varchar(255),
     foreign key ("from") references "interface_route_table"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_interface_route_table_service_instance on ref_interface_route_table_service_instance ("from");


create table tenant_share_interface_route_table (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "interface_route_table"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_interface_route_table_id on tenant_share_interface_route_table("uuid");
create index index_t_interface_route_table_to on tenant_share_interface_route_table("to");

create table domain_share_interface_route_table (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "interface_route_table"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_interface_route_table_id on domain_share_interface_route_table("uuid");
create index index_d_interface_route_table_to on domain_share_interface_route_table("to");





create table tenant_share_loadbalancer_healthmonitor (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_healthmonitor"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_loadbalancer_healthmonitor_id on tenant_share_loadbalancer_healthmonitor("uuid");
create index index_t_loadbalancer_healthmonitor_to on tenant_share_loadbalancer_healthmonitor("to");

create table domain_share_loadbalancer_healthmonitor (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_healthmonitor"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_loadbalancer_healthmonitor_id on domain_share_loadbalancer_healthmonitor("uuid");
create index index_d_loadbalancer_healthmonitor_to on domain_share_loadbalancer_healthmonitor("to");




create table ref_loadbalancer_listener_loadbalancer (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer_listener"(uuid) on delete cascade, 
    foreign key ("to") references "loadbalancer"(uuid));

create index index_loadbalancer_listener_loadbalancer on ref_loadbalancer_listener_loadbalancer ("from");


create table tenant_share_loadbalancer_listener (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_listener"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_loadbalancer_listener_id on tenant_share_loadbalancer_listener("uuid");
create index index_t_loadbalancer_listener_to on tenant_share_loadbalancer_listener("to");

create table domain_share_loadbalancer_listener (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_listener"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_loadbalancer_listener_id on domain_share_loadbalancer_listener("uuid");
create index index_d_loadbalancer_listener_to on domain_share_loadbalancer_listener("to");





create table tenant_share_loadbalancer_member (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_member"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_loadbalancer_member_id on tenant_share_loadbalancer_member("uuid");
create index index_t_loadbalancer_member_to on tenant_share_loadbalancer_member("to");

create table domain_share_loadbalancer_member (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_member"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_loadbalancer_member_id on domain_share_loadbalancer_member("uuid");
create index index_d_loadbalancer_member_to on domain_share_loadbalancer_member("to");




create table ref_loadbalancer_pool_loadbalancer_healthmonitor (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer_pool"(uuid) on delete cascade, 
    foreign key ("to") references "loadbalancer_healthmonitor"(uuid));

create index index_loadbalancer_pool_loadbalancer_healthmonitor on ref_loadbalancer_pool_loadbalancer_healthmonitor ("from");

create table ref_loadbalancer_pool_service_appliance_set (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer_pool"(uuid) on delete cascade, 
    foreign key ("to") references "service_appliance_set"(uuid));

create index index_loadbalancer_pool_service_appliance_set on ref_loadbalancer_pool_service_appliance_set ("from");

create table ref_loadbalancer_pool_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer_pool"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_loadbalancer_pool_virtual_machine_interface on ref_loadbalancer_pool_virtual_machine_interface ("from");

create table ref_loadbalancer_pool_loadbalancer_listener (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer_pool"(uuid) on delete cascade, 
    foreign key ("to") references "loadbalancer_listener"(uuid));

create index index_loadbalancer_pool_loadbalancer_listener on ref_loadbalancer_pool_loadbalancer_listener ("from");

create table ref_loadbalancer_pool_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer_pool"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_loadbalancer_pool_service_instance on ref_loadbalancer_pool_service_instance ("from");


create table tenant_share_loadbalancer_pool (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_pool"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_loadbalancer_pool_id on tenant_share_loadbalancer_pool("uuid");
create index index_t_loadbalancer_pool_to on tenant_share_loadbalancer_pool("to");

create table domain_share_loadbalancer_pool (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer_pool"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_loadbalancer_pool_id on domain_share_loadbalancer_pool("uuid");
create index index_d_loadbalancer_pool_to on domain_share_loadbalancer_pool("to");




create table ref_loadbalancer_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_loadbalancer_virtual_machine_interface on ref_loadbalancer_virtual_machine_interface ("from");

create table ref_loadbalancer_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_loadbalancer_service_instance on ref_loadbalancer_service_instance ("from");

create table ref_loadbalancer_service_appliance_set (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "loadbalancer"(uuid) on delete cascade, 
    foreign key ("to") references "service_appliance_set"(uuid));

create index index_loadbalancer_service_appliance_set on ref_loadbalancer_service_appliance_set ("from");


create table tenant_share_loadbalancer (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_loadbalancer_id on tenant_share_loadbalancer("uuid");
create index index_t_loadbalancer_to on tenant_share_loadbalancer("to");

create table domain_share_loadbalancer (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "loadbalancer"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_loadbalancer_id on domain_share_loadbalancer("uuid");
create index index_d_loadbalancer_to on domain_share_loadbalancer("to");




create table ref_logical_interface_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_interface"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_logical_interface_virtual_machine_interface on ref_logical_interface_virtual_machine_interface ("from");


create table tenant_share_logical_interface (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "logical_interface"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_logical_interface_id on tenant_share_logical_interface("uuid");
create index index_t_logical_interface_to on tenant_share_logical_interface("to");

create table domain_share_logical_interface (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "logical_interface"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_logical_interface_id on domain_share_logical_interface("uuid");
create index index_d_logical_interface_to on domain_share_logical_interface("to");




create table ref_logical_router_route_target (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "route_target"(uuid));

create index index_logical_router_route_target on ref_logical_router_route_target ("from");

create table ref_logical_router_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_logical_router_virtual_machine_interface on ref_logical_router_virtual_machine_interface ("from");

create table ref_logical_router_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_logical_router_service_instance on ref_logical_router_service_instance ("from");

create table ref_logical_router_route_table (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "route_table"(uuid));

create index index_logical_router_route_table on ref_logical_router_route_table ("from");

create table ref_logical_router_virtual_network (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_network"(uuid));

create index index_logical_router_virtual_network on ref_logical_router_virtual_network ("from");

create table ref_logical_router_physical_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "physical_router"(uuid));

create index index_logical_router_physical_router on ref_logical_router_physical_router ("from");

create table ref_logical_router_bgpvpn (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "logical_router"(uuid) on delete cascade, 
    foreign key ("to") references "bgpvpn"(uuid));

create index index_logical_router_bgpvpn on ref_logical_router_bgpvpn ("from");


create table tenant_share_logical_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "logical_router"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_logical_router_id on tenant_share_logical_router("uuid");
create index index_t_logical_router_to on tenant_share_logical_router("to");

create table domain_share_logical_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "logical_router"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_logical_router_id on domain_share_logical_router("uuid");
create index index_d_logical_router_to on domain_share_logical_router("to");





create table tenant_share_namespace (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "namespace"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_namespace_id on tenant_share_namespace("uuid");
create index index_t_namespace_to on tenant_share_namespace("to");

create table domain_share_namespace (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "namespace"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_namespace_id on domain_share_namespace("uuid");
create index index_d_namespace_to on domain_share_namespace("to");




create table ref_network_device_config_physical_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "network_device_config"(uuid) on delete cascade, 
    foreign key ("to") references "physical_router"(uuid));

create index index_network_device_config_physical_router on ref_network_device_config_physical_router ("from");


create table tenant_share_network_device_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "network_device_config"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_network_device_config_id on tenant_share_network_device_config("uuid");
create index index_t_network_device_config_to on tenant_share_network_device_config("to");

create table domain_share_network_device_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "network_device_config"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_network_device_config_id on domain_share_network_device_config("uuid");
create index index_d_network_device_config_to on domain_share_network_device_config("to");




create table ref_network_ipam_virtual_DNS (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "network_ipam"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_DNS"(uuid));

create index index_network_ipam_virtual_DNS on ref_network_ipam_virtual_DNS ("from");


create table tenant_share_network_ipam (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "network_ipam"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_network_ipam_id on tenant_share_network_ipam("uuid");
create index index_t_network_ipam_to on tenant_share_network_ipam("to");

create table domain_share_network_ipam (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "network_ipam"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_network_ipam_id on domain_share_network_ipam("uuid");
create index index_d_network_ipam_to on domain_share_network_ipam("to");





create table tenant_share_network_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "network_policy"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_network_policy_id on tenant_share_network_policy("uuid");
create index index_t_network_policy_to on tenant_share_network_policy("to");

create table domain_share_network_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "network_policy"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_network_policy_id on domain_share_network_policy("uuid");
create index index_d_network_policy_to on domain_share_network_policy("to");





create table tenant_share_peering_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "peering_policy"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_peering_policy_id on tenant_share_peering_policy("uuid");
create index index_t_peering_policy_to on tenant_share_peering_policy("to");

create table domain_share_peering_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "peering_policy"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_peering_policy_id on domain_share_peering_policy("uuid");
create index index_d_peering_policy_to on domain_share_peering_policy("to");




create table ref_physical_interface_physical_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "physical_interface"(uuid) on delete cascade, 
    foreign key ("to") references "physical_interface"(uuid));

create index index_physical_interface_physical_interface on ref_physical_interface_physical_interface ("from");


create table tenant_share_physical_interface (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "physical_interface"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_physical_interface_id on tenant_share_physical_interface("uuid");
create index index_t_physical_interface_to on tenant_share_physical_interface("to");

create table domain_share_physical_interface (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "physical_interface"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_physical_interface_id on domain_share_physical_interface("uuid");
create index index_d_physical_interface_to on domain_share_physical_interface("to");




create table ref_physical_router_virtual_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "physical_router"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_router"(uuid));

create index index_physical_router_virtual_router on ref_physical_router_virtual_router ("from");

create table ref_physical_router_virtual_network (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "physical_router"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_network"(uuid));

create index index_physical_router_virtual_network on ref_physical_router_virtual_network ("from");

create table ref_physical_router_bgp_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "physical_router"(uuid) on delete cascade, 
    foreign key ("to") references "bgp_router"(uuid));

create index index_physical_router_bgp_router on ref_physical_router_bgp_router ("from");


create table tenant_share_physical_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "physical_router"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_physical_router_id on tenant_share_physical_router("uuid");
create index index_t_physical_router_to on tenant_share_physical_router("to");

create table domain_share_physical_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "physical_router"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_physical_router_id on domain_share_physical_router("uuid");
create index index_d_physical_router_to on domain_share_physical_router("to");





create table tenant_share_policy_management (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "policy_management"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_policy_management_id on tenant_share_policy_management("uuid");
create index index_t_policy_management_to on tenant_share_policy_management("to");

create table domain_share_policy_management (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "policy_management"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_policy_management_id on domain_share_policy_management("uuid");
create index index_d_policy_management_to on domain_share_policy_management("to");





create table tenant_share_port_tuple (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "port_tuple"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_port_tuple_id on tenant_share_port_tuple("uuid");
create index index_t_port_tuple_to on tenant_share_port_tuple("to");

create table domain_share_port_tuple (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "port_tuple"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_port_tuple_id on domain_share_port_tuple("uuid");
create index index_d_port_tuple_to on domain_share_port_tuple("to");




create table ref_project_floating_ip_pool (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "project"(uuid) on delete cascade, 
    foreign key ("to") references "floating_ip_pool"(uuid));

create index index_project_floating_ip_pool on ref_project_floating_ip_pool ("from");

create table ref_project_alias_ip_pool (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "project"(uuid) on delete cascade, 
    foreign key ("to") references "alias_ip_pool"(uuid));

create index index_project_alias_ip_pool on ref_project_alias_ip_pool ("from");

create table ref_project_namespace (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "ip_prefix" varchar(255),
    "ip_prefix_len" int,
     foreign key ("from") references "project"(uuid) on delete cascade, 
    foreign key ("to") references "namespace"(uuid));

create index index_project_namespace on ref_project_namespace ("from");

create table ref_project_application_policy_set (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "project"(uuid) on delete cascade, 
    foreign key ("to") references "application_policy_set"(uuid));

create index index_project_application_policy_set on ref_project_application_policy_set ("from");


create table tenant_share_project (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "project"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_project_id on tenant_share_project("uuid");
create index index_t_project_to on tenant_share_project("to");

create table domain_share_project (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "project"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_project_id on domain_share_project("uuid");
create index index_d_project_to on domain_share_project("to");




create table ref_provider_attachment_virtual_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "provider_attachment"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_router"(uuid));

create index index_provider_attachment_virtual_router on ref_provider_attachment_virtual_router ("from");


create table tenant_share_provider_attachment (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "provider_attachment"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_provider_attachment_id on tenant_share_provider_attachment("uuid");
create index index_t_provider_attachment_to on tenant_share_provider_attachment("to");

create table domain_share_provider_attachment (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "provider_attachment"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_provider_attachment_id on domain_share_provider_attachment("uuid");
create index index_d_provider_attachment_to on domain_share_provider_attachment("to");




create table ref_qos_config_global_system_config (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "qos_config"(uuid) on delete cascade, 
    foreign key ("to") references "global_system_config"(uuid));

create index index_qos_config_global_system_config on ref_qos_config_global_system_config ("from");


create table tenant_share_qos_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "qos_config"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_qos_config_id on tenant_share_qos_config("uuid");
create index index_t_qos_config_to on tenant_share_qos_config("to");

create table domain_share_qos_config (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "qos_config"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_qos_config_id on domain_share_qos_config("uuid");
create index index_d_qos_config_to on domain_share_qos_config("to");





create table tenant_share_qos_queue (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "qos_queue"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_qos_queue_id on tenant_share_qos_queue("uuid");
create index index_t_qos_queue_to on tenant_share_qos_queue("to");

create table domain_share_qos_queue (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "qos_queue"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_qos_queue_id on domain_share_qos_queue("uuid");
create index index_d_qos_queue_to on domain_share_qos_queue("to");




create table ref_route_aggregate_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "interface_type" varchar(255),
     foreign key ("from") references "route_aggregate"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_route_aggregate_service_instance on ref_route_aggregate_service_instance ("from");


create table tenant_share_route_aggregate (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "route_aggregate"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_route_aggregate_id on tenant_share_route_aggregate("uuid");
create index index_t_route_aggregate_to on tenant_share_route_aggregate("to");

create table domain_share_route_aggregate (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "route_aggregate"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_route_aggregate_id on domain_share_route_aggregate("uuid");
create index index_d_route_aggregate_to on domain_share_route_aggregate("to");





create table tenant_share_route_table (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "route_table"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_route_table_id on tenant_share_route_table("uuid");
create index index_t_route_table_to on tenant_share_route_table("to");

create table domain_share_route_table (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "route_table"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_route_table_id on domain_share_route_table("uuid");
create index index_d_route_table_to on domain_share_route_table("to");





create table tenant_share_route_target (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "route_target"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_route_target_id on tenant_share_route_target("uuid");
create index index_t_route_target_to on tenant_share_route_target("to");

create table domain_share_route_target (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "route_target"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_route_target_id on domain_share_route_target("uuid");
create index index_d_route_target_to on domain_share_route_target("to");





create table tenant_share_routing_instance (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "routing_instance"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_routing_instance_id on tenant_share_routing_instance("uuid");
create index index_t_routing_instance_to on tenant_share_routing_instance("to");

create table domain_share_routing_instance (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "routing_instance"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_routing_instance_id on domain_share_routing_instance("uuid");
create index index_d_routing_instance_to on domain_share_routing_instance("to");




create table ref_routing_policy_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "right_sequence" varchar(255),
    "left_sequence" varchar(255),
     foreign key ("from") references "routing_policy"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_routing_policy_service_instance on ref_routing_policy_service_instance ("from");


create table tenant_share_routing_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "routing_policy"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_routing_policy_id on tenant_share_routing_policy("uuid");
create index index_t_routing_policy_to on tenant_share_routing_policy("to");

create table domain_share_routing_policy (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "routing_policy"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_routing_policy_id on domain_share_routing_policy("uuid");
create index index_d_routing_policy_to on domain_share_routing_policy("to");





create table tenant_share_security_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "security_group"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_security_group_id on tenant_share_security_group("uuid");
create index index_t_security_group_to on tenant_share_security_group("to");

create table domain_share_security_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "security_group"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_security_group_id on domain_share_security_group("uuid");
create index index_d_security_group_to on domain_share_security_group("to");




create table ref_security_logging_object_security_group (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "rule" json,
     foreign key ("from") references "security_logging_object"(uuid) on delete cascade, 
    foreign key ("to") references "security_group"(uuid));

create index index_security_logging_object_security_group on ref_security_logging_object_security_group ("from");

create table ref_security_logging_object_network_policy (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "rule" json,
     foreign key ("from") references "security_logging_object"(uuid) on delete cascade, 
    foreign key ("to") references "network_policy"(uuid));

create index index_security_logging_object_network_policy on ref_security_logging_object_network_policy ("from");


create table tenant_share_security_logging_object (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "security_logging_object"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_security_logging_object_id on tenant_share_security_logging_object("uuid");
create index index_t_security_logging_object_to on tenant_share_security_logging_object("to");

create table domain_share_security_logging_object (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "security_logging_object"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_security_logging_object_id on domain_share_security_logging_object("uuid");
create index index_d_security_logging_object_to on domain_share_security_logging_object("to");




create table ref_service_appliance_physical_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "interface_type" varchar(255),
     foreign key ("from") references "service_appliance"(uuid) on delete cascade, 
    foreign key ("to") references "physical_interface"(uuid));

create index index_service_appliance_physical_interface on ref_service_appliance_physical_interface ("from");


create table tenant_share_service_appliance (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_appliance"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_appliance_id on tenant_share_service_appliance("uuid");
create index index_t_service_appliance_to on tenant_share_service_appliance("to");

create table domain_share_service_appliance (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_appliance"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_appliance_id on domain_share_service_appliance("uuid");
create index index_d_service_appliance_to on domain_share_service_appliance("to");





create table tenant_share_service_appliance_set (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_appliance_set"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_appliance_set_id on tenant_share_service_appliance_set("uuid");
create index index_t_service_appliance_set_to on tenant_share_service_appliance_set("to");

create table domain_share_service_appliance_set (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_appliance_set"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_appliance_set_id on domain_share_service_appliance_set("uuid");
create index index_d_service_appliance_set_to on domain_share_service_appliance_set("to");




create table ref_service_connection_module_service_object (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "service_connection_module"(uuid) on delete cascade, 
    foreign key ("to") references "service_object"(uuid));

create index index_service_connection_module_service_object on ref_service_connection_module_service_object ("from");


create table tenant_share_service_connection_module (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_connection_module"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_connection_module_id on tenant_share_service_connection_module("uuid");
create index index_t_service_connection_module_to on tenant_share_service_connection_module("to");

create table domain_share_service_connection_module (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_connection_module"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_connection_module_id on domain_share_service_connection_module("uuid");
create index index_d_service_connection_module_to on domain_share_service_connection_module("to");




create table ref_service_endpoint_service_connection_module (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "service_endpoint"(uuid) on delete cascade, 
    foreign key ("to") references "service_connection_module"(uuid));

create index index_service_endpoint_service_connection_module on ref_service_endpoint_service_connection_module ("from");

create table ref_service_endpoint_physical_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "service_endpoint"(uuid) on delete cascade, 
    foreign key ("to") references "physical_router"(uuid));

create index index_service_endpoint_physical_router on ref_service_endpoint_physical_router ("from");

create table ref_service_endpoint_service_object (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "service_endpoint"(uuid) on delete cascade, 
    foreign key ("to") references "service_object"(uuid));

create index index_service_endpoint_service_object on ref_service_endpoint_service_object ("from");


create table tenant_share_service_endpoint (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_endpoint"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_endpoint_id on tenant_share_service_endpoint("uuid");
create index index_t_service_endpoint_to on tenant_share_service_endpoint("to");

create table domain_share_service_endpoint (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_endpoint"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_endpoint_id on domain_share_service_endpoint("uuid");
create index index_d_service_endpoint_to on domain_share_service_endpoint("to");





create table tenant_share_service_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_group"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_group_id on tenant_share_service_group("uuid");
create index index_t_service_group_to on tenant_share_service_group("to");

create table domain_share_service_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_group"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_group_id on domain_share_service_group("uuid");
create index index_d_service_group_to on domain_share_service_group("to");




create table ref_service_health_check_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "interface_type" varchar(255),
     foreign key ("from") references "service_health_check"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_service_health_check_service_instance on ref_service_health_check_service_instance ("from");


create table tenant_share_service_health_check (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_health_check"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_health_check_id on tenant_share_service_health_check("uuid");
create index index_t_service_health_check_to on tenant_share_service_health_check("to");

create table domain_share_service_health_check (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_health_check"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_health_check_id on domain_share_service_health_check("uuid");
create index index_d_service_health_check_to on domain_share_service_health_check("to");




create table ref_service_instance_service_template (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "service_instance"(uuid) on delete cascade, 
    foreign key ("to") references "service_template"(uuid));

create index index_service_instance_service_template on ref_service_instance_service_template ("from");

create table ref_service_instance_instance_ip (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "interface_type" varchar(255),
     foreign key ("from") references "service_instance"(uuid) on delete cascade, 
    foreign key ("to") references "instance_ip"(uuid));

create index index_service_instance_instance_ip on ref_service_instance_instance_ip ("from");


create table tenant_share_service_instance (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_instance"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_instance_id on tenant_share_service_instance("uuid");
create index index_t_service_instance_to on tenant_share_service_instance("to");

create table domain_share_service_instance (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_instance"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_instance_id on domain_share_service_instance("uuid");
create index index_d_service_instance_to on domain_share_service_instance("to");





create table tenant_share_service_object (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_object"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_object_id on tenant_share_service_object("uuid");
create index index_t_service_object_to on tenant_share_service_object("to");

create table domain_share_service_object (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_object"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_object_id on domain_share_service_object("uuid");
create index index_d_service_object_to on domain_share_service_object("to");




create table ref_service_template_service_appliance_set (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "service_template"(uuid) on delete cascade, 
    foreign key ("to") references "service_appliance_set"(uuid));

create index index_service_template_service_appliance_set on ref_service_template_service_appliance_set ("from");


create table tenant_share_service_template (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_template"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_service_template_id on tenant_share_service_template("uuid");
create index index_t_service_template_to on tenant_share_service_template("to");

create table domain_share_service_template (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "service_template"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_service_template_id on domain_share_service_template("uuid");
create index index_d_service_template_to on domain_share_service_template("to");




create table ref_subnet_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "subnet"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_subnet_virtual_machine_interface on ref_subnet_virtual_machine_interface ("from");


create table tenant_share_subnet (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "subnet"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_subnet_id on tenant_share_subnet("uuid");
create index index_t_subnet_to on tenant_share_subnet("to");

create table domain_share_subnet (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "subnet"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_subnet_id on domain_share_subnet("uuid");
create index index_d_subnet_to on domain_share_subnet("to");




create table ref_tag_tag_type (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "tag"(uuid) on delete cascade, 
    foreign key ("to") references "tag_type"(uuid));

create index index_tag_tag_type on ref_tag_tag_type ("from");


create table tenant_share_tag (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "tag"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_tag_id on tenant_share_tag("uuid");
create index index_t_tag_to on tenant_share_tag("to");

create table domain_share_tag (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "tag"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_tag_id on domain_share_tag("uuid");
create index index_d_tag_to on domain_share_tag("to");





create table tenant_share_tag_type (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "tag_type"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_tag_type_id on tenant_share_tag_type("uuid");
create index index_t_tag_type_to on tenant_share_tag_type("to");

create table domain_share_tag_type (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "tag_type"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_tag_type_id on domain_share_tag_type("uuid");
create index index_d_tag_type_to on domain_share_tag_type("to");





create table tenant_share_user (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "user"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_user_id on tenant_share_user("uuid");
create index index_t_user_to on tenant_share_user("to");

create table domain_share_user (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "user"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_user_id on domain_share_user("uuid");
create index index_d_user_to on domain_share_user("to");





create table tenant_share_virtual_DNS_record (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_DNS_record"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_DNS_record_id on tenant_share_virtual_DNS_record("uuid");
create index index_t_virtual_DNS_record_to on tenant_share_virtual_DNS_record("to");

create table domain_share_virtual_DNS_record (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_DNS_record"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_DNS_record_id on domain_share_virtual_DNS_record("uuid");
create index index_d_virtual_DNS_record_to on domain_share_virtual_DNS_record("to");





create table tenant_share_virtual_DNS (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_DNS"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_DNS_id on tenant_share_virtual_DNS("uuid");
create index index_t_virtual_DNS_to on tenant_share_virtual_DNS("to");

create table domain_share_virtual_DNS (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_DNS"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_DNS_id on domain_share_virtual_DNS("uuid");
create index index_d_virtual_DNS_to on domain_share_virtual_DNS("to");




create table ref_virtual_ip_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_ip"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_virtual_ip_virtual_machine_interface on ref_virtual_ip_virtual_machine_interface ("from");

create table ref_virtual_ip_loadbalancer_pool (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_ip"(uuid) on delete cascade, 
    foreign key ("to") references "loadbalancer_pool"(uuid));

create index index_virtual_ip_loadbalancer_pool on ref_virtual_ip_loadbalancer_pool ("from");


create table tenant_share_virtual_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_ip"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_ip_id on tenant_share_virtual_ip("uuid");
create index index_t_virtual_ip_to on tenant_share_virtual_ip("to");

create table domain_share_virtual_ip (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_ip"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_ip_id on domain_share_virtual_ip("uuid");
create index index_d_virtual_ip_to on domain_share_virtual_ip("to");




create table ref_virtual_machine_interface_virtual_machine_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine_interface"(uuid));

create index index_virtual_machine_interface_virtual_machine_interface on ref_virtual_machine_interface_virtual_machine_interface ("from");

create table ref_virtual_machine_interface_physical_interface (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "physical_interface"(uuid));

create index index_virtual_machine_interface_physical_interface on ref_virtual_machine_interface_physical_interface ("from");

create table ref_virtual_machine_interface_service_health_check (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "service_health_check"(uuid));

create index index_virtual_machine_interface_service_health_check on ref_virtual_machine_interface_service_health_check ("from");

create table ref_virtual_machine_interface_service_endpoint (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "service_endpoint"(uuid));

create index index_virtual_machine_interface_service_endpoint on ref_virtual_machine_interface_service_endpoint ("from");

create table ref_virtual_machine_interface_bgp_router (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "bgp_router"(uuid));

create index index_virtual_machine_interface_bgp_router on ref_virtual_machine_interface_bgp_router ("from");

create table ref_virtual_machine_interface_qos_config (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "qos_config"(uuid));

create index index_virtual_machine_interface_qos_config on ref_virtual_machine_interface_qos_config ("from");

create table ref_virtual_machine_interface_virtual_network (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_network"(uuid));

create index index_virtual_machine_interface_virtual_network on ref_virtual_machine_interface_virtual_network ("from");

create table ref_virtual_machine_interface_security_logging_object (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "security_logging_object"(uuid));

create index index_virtual_machine_interface_security_logging_object on ref_virtual_machine_interface_security_logging_object ("from");

create table ref_virtual_machine_interface_interface_route_table (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "interface_route_table"(uuid));

create index index_virtual_machine_interface_interface_route_table on ref_virtual_machine_interface_interface_route_table ("from");

create table ref_virtual_machine_interface_bridge_domain (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "vlan_tag" int,
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "bridge_domain"(uuid));

create index index_virtual_machine_interface_bridge_domain on ref_virtual_machine_interface_bridge_domain ("from");

create table ref_virtual_machine_interface_virtual_machine (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine"(uuid));

create index index_virtual_machine_interface_virtual_machine on ref_virtual_machine_interface_virtual_machine ("from");

create table ref_virtual_machine_interface_routing_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "vlan_tag" int,
    "src_mac" varchar(255),
    "service_chain_address" varchar(255),
    "dst_mac" varchar(255),
    "protocol" varchar(255),
    "ipv6_service_chain_address" varchar(255),
    "direction" varchar(255),
    "mpls_label" int,
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "routing_instance"(uuid));

create index index_virtual_machine_interface_routing_instance on ref_virtual_machine_interface_routing_instance ("from");

create table ref_virtual_machine_interface_port_tuple (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "port_tuple"(uuid));

create index index_virtual_machine_interface_port_tuple on ref_virtual_machine_interface_port_tuple ("from");

create table ref_virtual_machine_interface_security_group (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine_interface"(uuid) on delete cascade, 
    foreign key ("to") references "security_group"(uuid));

create index index_virtual_machine_interface_security_group on ref_virtual_machine_interface_security_group ("from");


create table tenant_share_virtual_machine_interface (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_machine_interface"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_machine_interface_id on tenant_share_virtual_machine_interface("uuid");
create index index_t_virtual_machine_interface_to on tenant_share_virtual_machine_interface("to");

create table domain_share_virtual_machine_interface (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_machine_interface"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_machine_interface_id on domain_share_virtual_machine_interface("uuid");
create index index_d_virtual_machine_interface_to on domain_share_virtual_machine_interface("to");




create table ref_virtual_machine_service_instance (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_machine"(uuid) on delete cascade, 
    foreign key ("to") references "service_instance"(uuid));

create index index_virtual_machine_service_instance on ref_virtual_machine_service_instance ("from");


create table tenant_share_virtual_machine (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_machine"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_machine_id on tenant_share_virtual_machine("uuid");
create index index_t_virtual_machine_to on tenant_share_virtual_machine("to");

create table domain_share_virtual_machine (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_machine"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_machine_id on domain_share_virtual_machine("uuid");
create index index_d_virtual_machine_to on domain_share_virtual_machine("to");




create table ref_virtual_network_network_policy (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "off_interval" varchar(255),
    "on_interval" varchar(255),
    "end_time" varchar(255),
    "start_time" varchar(255),
    "major" int,
    "minor" int,
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "network_policy"(uuid));

create index index_virtual_network_network_policy on ref_virtual_network_network_policy ("from");

create table ref_virtual_network_qos_config (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "qos_config"(uuid));

create index index_virtual_network_qos_config on ref_virtual_network_qos_config ("from");

create table ref_virtual_network_route_table (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "route_table"(uuid));

create index index_virtual_network_route_table on ref_virtual_network_route_table ("from");

create table ref_virtual_network_virtual_network (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_network"(uuid));

create index index_virtual_network_virtual_network on ref_virtual_network_virtual_network ("from");

create table ref_virtual_network_bgpvpn (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "bgpvpn"(uuid));

create index index_virtual_network_bgpvpn on ref_virtual_network_bgpvpn ("from");

create table ref_virtual_network_network_ipam (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "ipam_subnets" json,
    "route" json,
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "network_ipam"(uuid));

create index index_virtual_network_network_ipam on ref_virtual_network_network_ipam ("from");

create table ref_virtual_network_security_logging_object (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_network"(uuid) on delete cascade, 
    foreign key ("to") references "security_logging_object"(uuid));

create index index_virtual_network_security_logging_object on ref_virtual_network_security_logging_object ("from");


create table tenant_share_virtual_network (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_network"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_network_id on tenant_share_virtual_network("uuid");
create index index_t_virtual_network_to on tenant_share_virtual_network("to");

create table domain_share_virtual_network (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_network"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_network_id on domain_share_virtual_network("uuid");
create index index_d_virtual_network_to on domain_share_virtual_network("to");




create table ref_virtual_router_virtual_machine (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "virtual_router"(uuid) on delete cascade, 
    foreign key ("to") references "virtual_machine"(uuid));

create index index_virtual_router_virtual_machine on ref_virtual_router_virtual_machine ("from");

create table ref_virtual_router_network_ipam (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    "subnet" json,
    "allocation_pools" json,
     foreign key ("from") references "virtual_router"(uuid) on delete cascade, 
    foreign key ("to") references "network_ipam"(uuid));

create index index_virtual_router_network_ipam on ref_virtual_router_network_ipam ("from");


create table tenant_share_virtual_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_router"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_virtual_router_id on tenant_share_virtual_router("uuid");
create index index_t_virtual_router_to on tenant_share_virtual_router("to");

create table domain_share_virtual_router (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "virtual_router"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_virtual_router_id on domain_share_virtual_router("uuid");
create index index_d_virtual_router_to on domain_share_virtual_router("to");




create table ref_appformix_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "appformix_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_appformix_node_node on ref_appformix_node_node ("from");


create table tenant_share_appformix_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "appformix_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_appformix_node_id on tenant_share_appformix_node("uuid");
create index index_t_appformix_node_to on tenant_share_appformix_node("to");

create table domain_share_appformix_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "appformix_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_appformix_node_id on domain_share_appformix_node("uuid");
create index index_d_appformix_node_to on domain_share_appformix_node("to");





create table tenant_share_baremetal_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "baremetal_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_baremetal_node_id on tenant_share_baremetal_node("uuid");
create index index_t_baremetal_node_to on tenant_share_baremetal_node("to");

create table domain_share_baremetal_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "baremetal_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_baremetal_node_id on domain_share_baremetal_node("uuid");
create index index_d_baremetal_node_to on domain_share_baremetal_node("to");





create table tenant_share_baremetal_port (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "baremetal_port"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_baremetal_port_id on tenant_share_baremetal_port("uuid");
create index index_t_baremetal_port_to on tenant_share_baremetal_port("to");

create table domain_share_baremetal_port (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "baremetal_port"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_baremetal_port_id on domain_share_baremetal_port("uuid");
create index index_d_baremetal_port_to on domain_share_baremetal_port("to");




create table ref_contrail_analytics_database_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_analytics_database_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_analytics_database_node_node on ref_contrail_analytics_database_node_node ("from");


create table tenant_share_contrail_analytics_database_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_analytics_database_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_analytics_database_node_id on tenant_share_contrail_analytics_database_node("uuid");
create index index_t_contrail_analytics_database_node_to on tenant_share_contrail_analytics_database_node("to");

create table domain_share_contrail_analytics_database_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_analytics_database_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_analytics_database_node_id on domain_share_contrail_analytics_database_node("uuid");
create index index_d_contrail_analytics_database_node_to on domain_share_contrail_analytics_database_node("to");




create table ref_contrail_analytics_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_analytics_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_analytics_node_node on ref_contrail_analytics_node_node ("from");


create table tenant_share_contrail_analytics_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_analytics_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_analytics_node_id on tenant_share_contrail_analytics_node("uuid");
create index index_t_contrail_analytics_node_to on tenant_share_contrail_analytics_node("to");

create table domain_share_contrail_analytics_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_analytics_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_analytics_node_id on domain_share_contrail_analytics_node("uuid");
create index index_d_contrail_analytics_node_to on domain_share_contrail_analytics_node("to");





create table tenant_share_contrail_cluster (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_cluster"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_cluster_id on tenant_share_contrail_cluster("uuid");
create index index_t_contrail_cluster_to on tenant_share_contrail_cluster("to");

create table domain_share_contrail_cluster (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_cluster"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_cluster_id on domain_share_contrail_cluster("uuid");
create index index_d_contrail_cluster_to on domain_share_contrail_cluster("to");




create table ref_contrail_config_database_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_config_database_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_config_database_node_node on ref_contrail_config_database_node_node ("from");


create table tenant_share_contrail_config_database_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_config_database_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_config_database_node_id on tenant_share_contrail_config_database_node("uuid");
create index index_t_contrail_config_database_node_to on tenant_share_contrail_config_database_node("to");

create table domain_share_contrail_config_database_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_config_database_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_config_database_node_id on domain_share_contrail_config_database_node("uuid");
create index index_d_contrail_config_database_node_to on domain_share_contrail_config_database_node("to");




create table ref_contrail_config_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_config_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_config_node_node on ref_contrail_config_node_node ("from");


create table tenant_share_contrail_config_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_config_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_config_node_id on tenant_share_contrail_config_node("uuid");
create index index_t_contrail_config_node_to on tenant_share_contrail_config_node("to");

create table domain_share_contrail_config_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_config_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_config_node_id on domain_share_contrail_config_node("uuid");
create index index_d_contrail_config_node_to on domain_share_contrail_config_node("to");




create table ref_contrail_control_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_control_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_control_node_node on ref_contrail_control_node_node ("from");


create table tenant_share_contrail_control_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_control_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_control_node_id on tenant_share_contrail_control_node("uuid");
create index index_t_contrail_control_node_to on tenant_share_contrail_control_node("to");

create table domain_share_contrail_control_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_control_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_control_node_id on domain_share_contrail_control_node("uuid");
create index index_d_contrail_control_node_to on domain_share_contrail_control_node("to");




create table ref_contrail_storage_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_storage_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_storage_node_node on ref_contrail_storage_node_node ("from");


create table tenant_share_contrail_storage_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_storage_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_storage_node_id on tenant_share_contrail_storage_node("uuid");
create index index_t_contrail_storage_node_to on tenant_share_contrail_storage_node("to");

create table domain_share_contrail_storage_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_storage_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_storage_node_id on domain_share_contrail_storage_node("uuid");
create index index_d_contrail_storage_node_to on domain_share_contrail_storage_node("to");




create table ref_contrail_vrouter_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_vrouter_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_vrouter_node_node on ref_contrail_vrouter_node_node ("from");


create table tenant_share_contrail_vrouter_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_vrouter_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_vrouter_node_id on tenant_share_contrail_vrouter_node("uuid");
create index index_t_contrail_vrouter_node_to on tenant_share_contrail_vrouter_node("to");

create table domain_share_contrail_vrouter_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_vrouter_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_vrouter_node_id on domain_share_contrail_vrouter_node("uuid");
create index index_d_contrail_vrouter_node_to on domain_share_contrail_vrouter_node("to");




create table ref_contrail_webui_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "contrail_webui_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_contrail_webui_node_node on ref_contrail_webui_node_node ("from");


create table tenant_share_contrail_webui_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_webui_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_contrail_webui_node_id on tenant_share_contrail_webui_node("uuid");
create index index_t_contrail_webui_node_to on tenant_share_contrail_webui_node("to");

create table domain_share_contrail_webui_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "contrail_webui_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_contrail_webui_node_id on domain_share_contrail_webui_node("uuid");
create index index_d_contrail_webui_node_to on domain_share_contrail_webui_node("to");





create table tenant_share_dashboard (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "dashboard"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_dashboard_id on tenant_share_dashboard("uuid");
create index index_t_dashboard_to on tenant_share_dashboard("to");

create table domain_share_dashboard (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "dashboard"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_dashboard_id on domain_share_dashboard("uuid");
create index index_d_dashboard_to on domain_share_dashboard("to");





create table tenant_share_flavor (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "flavor"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_flavor_id on tenant_share_flavor("uuid");
create index index_t_flavor_to on tenant_share_flavor("to");

create table domain_share_flavor (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "flavor"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_flavor_id on domain_share_flavor("uuid");
create index index_d_flavor_to on domain_share_flavor("to");





create table tenant_share_os_image (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "os_image"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_os_image_id on tenant_share_os_image("uuid");
create index index_t_os_image_to on tenant_share_os_image("to");

create table domain_share_os_image (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "os_image"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_os_image_id on domain_share_os_image("uuid");
create index index_d_os_image_to on domain_share_os_image("to");





create table tenant_share_keypair (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "keypair"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_keypair_id on tenant_share_keypair("uuid");
create index index_t_keypair_to on tenant_share_keypair("to");

create table domain_share_keypair (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "keypair"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_keypair_id on domain_share_keypair("uuid");
create index index_d_keypair_to on domain_share_keypair("to");




create table ref_kubernetes_master_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "kubernetes_master_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_kubernetes_master_node_node on ref_kubernetes_master_node_node ("from");


create table tenant_share_kubernetes_master_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "kubernetes_master_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_kubernetes_master_node_id on tenant_share_kubernetes_master_node("uuid");
create index index_t_kubernetes_master_node_to on tenant_share_kubernetes_master_node("to");

create table domain_share_kubernetes_master_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "kubernetes_master_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_kubernetes_master_node_id on domain_share_kubernetes_master_node("uuid");
create index index_d_kubernetes_master_node_to on domain_share_kubernetes_master_node("to");




create table ref_kubernetes_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "kubernetes_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_kubernetes_node_node on ref_kubernetes_node_node ("from");


create table tenant_share_kubernetes_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "kubernetes_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_kubernetes_node_id on tenant_share_kubernetes_node("uuid");
create index index_t_kubernetes_node_to on tenant_share_kubernetes_node("to");

create table domain_share_kubernetes_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "kubernetes_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_kubernetes_node_id on domain_share_kubernetes_node("uuid");
create index index_d_kubernetes_node_to on domain_share_kubernetes_node("to");





create table tenant_share_location (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "location"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_location_id on tenant_share_location("uuid");
create index index_t_location_to on tenant_share_location("to");

create table domain_share_location (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "location"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_location_id on domain_share_location("uuid");
create index index_d_location_to on domain_share_location("to");




create table ref_node_keypair (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "node"(uuid) on delete cascade, 
    foreign key ("to") references "keypair"(uuid));

create index index_node_keypair on ref_node_keypair ("from");


create table tenant_share_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_node_id on tenant_share_node("uuid");
create index index_t_node_to on tenant_share_node("to");

create table domain_share_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_node_id on domain_share_node("uuid");
create index index_d_node_to on domain_share_node("to");




create table ref_openstack_compute_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "openstack_compute_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_openstack_compute_node_node on ref_openstack_compute_node_node ("from");


create table tenant_share_openstack_compute_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_compute_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_openstack_compute_node_id on tenant_share_openstack_compute_node("uuid");
create index index_t_openstack_compute_node_to on tenant_share_openstack_compute_node("to");

create table domain_share_openstack_compute_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_compute_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_openstack_compute_node_id on domain_share_openstack_compute_node("uuid");
create index index_d_openstack_compute_node_to on domain_share_openstack_compute_node("to");




create table ref_openstack_control_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "openstack_control_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_openstack_control_node_node on ref_openstack_control_node_node ("from");


create table tenant_share_openstack_control_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_control_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_openstack_control_node_id on tenant_share_openstack_control_node("uuid");
create index index_t_openstack_control_node_to on tenant_share_openstack_control_node("to");

create table domain_share_openstack_control_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_control_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_openstack_control_node_id on domain_share_openstack_control_node("uuid");
create index index_d_openstack_control_node_to on domain_share_openstack_control_node("to");




create table ref_openstack_monitoring_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "openstack_monitoring_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_openstack_monitoring_node_node on ref_openstack_monitoring_node_node ("from");


create table tenant_share_openstack_monitoring_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_monitoring_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_openstack_monitoring_node_id on tenant_share_openstack_monitoring_node("uuid");
create index index_t_openstack_monitoring_node_to on tenant_share_openstack_monitoring_node("to");

create table domain_share_openstack_monitoring_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_monitoring_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_openstack_monitoring_node_id on domain_share_openstack_monitoring_node("uuid");
create index index_d_openstack_monitoring_node_to on domain_share_openstack_monitoring_node("to");




create table ref_openstack_network_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "openstack_network_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_openstack_network_node_node on ref_openstack_network_node_node ("from");


create table tenant_share_openstack_network_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_network_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_openstack_network_node_id on tenant_share_openstack_network_node("uuid");
create index index_t_openstack_network_node_to on tenant_share_openstack_network_node("to");

create table domain_share_openstack_network_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_network_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_openstack_network_node_id on domain_share_openstack_network_node("uuid");
create index index_d_openstack_network_node_to on domain_share_openstack_network_node("to");




create table ref_openstack_storage_node_node (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "openstack_storage_node"(uuid) on delete cascade, 
    foreign key ("to") references "node"(uuid));

create index index_openstack_storage_node_node on ref_openstack_storage_node_node ("from");


create table tenant_share_openstack_storage_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_storage_node"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_openstack_storage_node_id on tenant_share_openstack_storage_node("uuid");
create index index_t_openstack_storage_node_to on tenant_share_openstack_storage_node("to");

create table domain_share_openstack_storage_node (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "openstack_storage_node"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_openstack_storage_node_id on domain_share_openstack_storage_node("uuid");
create index index_d_openstack_storage_node_to on domain_share_openstack_storage_node("to");





create table tenant_share_port (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "port"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_port_id on tenant_share_port("uuid");
create index index_t_port_to on tenant_share_port("to");

create table domain_share_port (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "port"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_port_id on domain_share_port("uuid");
create index index_d_port_to on domain_share_port("to");





create table tenant_share_server (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "server"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_server_id on tenant_share_server("uuid");
create index index_t_server_to on tenant_share_server("to");

create table domain_share_server (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "server"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_server_id on domain_share_server("uuid");
create index index_d_server_to on domain_share_server("to");




create table ref_vpn_group_location (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
     foreign key ("from") references "vpn_group"(uuid) on delete cascade, 
    foreign key ("to") references "location"(uuid));

create index index_vpn_group_location on ref_vpn_group_location ("from");


create table tenant_share_vpn_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "vpn_group"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_vpn_group_id on tenant_share_vpn_group("uuid");
create index index_t_vpn_group_to on tenant_share_vpn_group("to");

create table domain_share_vpn_group (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "vpn_group"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_vpn_group_id on domain_share_vpn_group("uuid");
create index index_d_vpn_group_to on domain_share_vpn_group("to");





create table tenant_share_widget (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "widget"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_widget_id on tenant_share_widget("uuid");
create index index_t_widget_to on tenant_share_widget("to");

create table domain_share_widget (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "widget"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_widget_id on domain_share_widget("uuid");
create index index_d_widget_to on domain_share_widget("to");


