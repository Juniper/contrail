SET FOREIGN_KEY_CHECKS=0;





truncate table access_control_list;




truncate table parent_access_control_list_security_group;

truncate table parent_access_control_list_virtual_network;




truncate table address_group;




truncate table parent_address_group_project;

truncate table parent_address_group_policy_management;




truncate table alarm;




truncate table parent_alarm_global_system_config;

truncate table parent_alarm_project;




truncate table alias_ip_pool;




truncate table parent_alias_ip_pool_virtual_network;




truncate table alias_ip;


truncate table ref_alias_ip_project;

truncate table ref_alias_ip_virtual_machine_interface;



truncate table parent_alias_ip_alias_ip_pool;




truncate table analytics_node;




truncate table parent_analytics_node_global_system_config;




truncate table api_access_list;




truncate table parent_api_access_list_project;

truncate table parent_api_access_list_global_system_config;

truncate table parent_api_access_list_domain;




truncate table application_policy_set;


truncate table ref_application_policy_set_firewall_policy;

truncate table ref_application_policy_set_global_vrouter_config;



truncate table parent_application_policy_set_policy_management;

truncate table parent_application_policy_set_project;




truncate table bgp_as_a_service;


truncate table ref_bgp_as_a_service_virtual_machine_interface;

truncate table ref_bgp_as_a_service_service_health_check;



truncate table parent_bgp_as_a_service_project;




truncate table bgp_router;







truncate table bgpvpn;




truncate table parent_bgpvpn_project;




truncate table bridge_domain;




truncate table parent_bridge_domain_virtual_network;




truncate table config_node;




truncate table parent_config_node_global_system_config;




truncate table config_root;


truncate table ref_config_root_tag;






truncate table customer_attachment;


truncate table ref_customer_attachment_virtual_machine_interface;

truncate table ref_customer_attachment_floating_ip;






truncate table database_node;




truncate table parent_database_node_global_system_config;




truncate table discovery_service_assignment;







truncate table domain;




truncate table parent_domain_config_root;




truncate table dsa_rule;




truncate table parent_dsa_rule_discovery_service_assignment;




truncate table e2_service_provider;


truncate table ref_e2_service_provider_physical_router;

truncate table ref_e2_service_provider_peering_policy;






truncate table firewall_policy;


truncate table ref_firewall_policy_firewall_rule;

truncate table ref_firewall_policy_security_logging_object;



truncate table parent_firewall_policy_project;

truncate table parent_firewall_policy_policy_management;




truncate table firewall_rule;


truncate table ref_firewall_rule_address_group;

truncate table ref_firewall_rule_security_logging_object;

truncate table ref_firewall_rule_virtual_network;

truncate table ref_firewall_rule_service_group;



truncate table parent_firewall_rule_project;

truncate table parent_firewall_rule_policy_management;




truncate table floating_ip_pool;




truncate table parent_floating_ip_pool_virtual_network;




truncate table floating_ip;


truncate table ref_floating_ip_virtual_machine_interface;

truncate table ref_floating_ip_project;



truncate table parent_floating_ip_floating_ip_pool;

truncate table parent_floating_ip_instance_ip;




truncate table forwarding_class;


truncate table ref_forwarding_class_qos_queue;



truncate table parent_forwarding_class_global_qos_config;




truncate table global_qos_config;




truncate table parent_global_qos_config_global_system_config;




truncate table global_system_config;


truncate table ref_global_system_config_bgp_router;



truncate table parent_global_system_config_config_root;




truncate table global_vrouter_config;




truncate table parent_global_vrouter_config_global_system_config;




truncate table instance_ip;


truncate table ref_instance_ip_network_ipam;

truncate table ref_instance_ip_virtual_network;

truncate table ref_instance_ip_virtual_machine_interface;

truncate table ref_instance_ip_physical_router;

truncate table ref_instance_ip_virtual_router;






truncate table interface_route_table;


truncate table ref_interface_route_table_service_instance;



truncate table parent_interface_route_table_project;




truncate table loadbalancer_healthmonitor;




truncate table parent_loadbalancer_healthmonitor_project;




truncate table loadbalancer_listener;


truncate table ref_loadbalancer_listener_loadbalancer;



truncate table parent_loadbalancer_listener_project;




truncate table loadbalancer_member;




truncate table parent_loadbalancer_member_loadbalancer_pool;




truncate table loadbalancer_pool;


truncate table ref_loadbalancer_pool_service_instance;

truncate table ref_loadbalancer_pool_loadbalancer_healthmonitor;

truncate table ref_loadbalancer_pool_service_appliance_set;

truncate table ref_loadbalancer_pool_virtual_machine_interface;

truncate table ref_loadbalancer_pool_loadbalancer_listener;



truncate table parent_loadbalancer_pool_project;




truncate table loadbalancer;


truncate table ref_loadbalancer_service_appliance_set;

truncate table ref_loadbalancer_virtual_machine_interface;

truncate table ref_loadbalancer_service_instance;



truncate table parent_loadbalancer_project;




truncate table logical_interface;


truncate table ref_logical_interface_virtual_machine_interface;



truncate table parent_logical_interface_physical_router;

truncate table parent_logical_interface_physical_interface;




truncate table logical_router;


truncate table ref_logical_router_service_instance;

truncate table ref_logical_router_route_table;

truncate table ref_logical_router_virtual_network;

truncate table ref_logical_router_physical_router;

truncate table ref_logical_router_bgpvpn;

truncate table ref_logical_router_route_target;

truncate table ref_logical_router_virtual_machine_interface;



truncate table parent_logical_router_project;




truncate table namespace;




truncate table parent_namespace_domain;




truncate table network_device_config;


truncate table ref_network_device_config_physical_router;






truncate table network_ipam;


truncate table ref_network_ipam_virtual_DNS;



truncate table parent_network_ipam_project;




truncate table network_policy;




truncate table parent_network_policy_project;




truncate table peering_policy;







truncate table physical_interface;


truncate table ref_physical_interface_physical_interface;



truncate table parent_physical_interface_physical_router;




truncate table physical_router;


truncate table ref_physical_router_bgp_router;

truncate table ref_physical_router_virtual_router;

truncate table ref_physical_router_virtual_network;



truncate table parent_physical_router_global_system_config;

truncate table parent_physical_router_location;




truncate table policy_management;







truncate table port_tuple;




truncate table parent_port_tuple_service_instance;




truncate table project;


truncate table ref_project_namespace;

truncate table ref_project_application_policy_set;

truncate table ref_project_floating_ip_pool;

truncate table ref_project_alias_ip_pool;



truncate table parent_project_domain;




truncate table provider_attachment;


truncate table ref_provider_attachment_virtual_router;






truncate table qos_config;


truncate table ref_qos_config_global_system_config;



truncate table parent_qos_config_project;

truncate table parent_qos_config_global_qos_config;




truncate table qos_queue;




truncate table parent_qos_queue_global_qos_config;




truncate table route_aggregate;


truncate table ref_route_aggregate_service_instance;



truncate table parent_route_aggregate_project;




truncate table route_table;




truncate table parent_route_table_project;




truncate table route_target;







truncate table routing_instance;




truncate table parent_routing_instance_virtual_network;




truncate table routing_policy;


truncate table ref_routing_policy_service_instance;



truncate table parent_routing_policy_project;




truncate table security_group;




truncate table parent_security_group_project;




truncate table security_logging_object;


truncate table ref_security_logging_object_security_group;

truncate table ref_security_logging_object_network_policy;



truncate table parent_security_logging_object_project;

truncate table parent_security_logging_object_global_vrouter_config;




truncate table service_appliance;


truncate table ref_service_appliance_physical_interface;



truncate table parent_service_appliance_service_appliance_set;




truncate table service_appliance_set;




truncate table parent_service_appliance_set_global_system_config;




truncate table service_connection_module;


truncate table ref_service_connection_module_service_object;






truncate table service_endpoint;


truncate table ref_service_endpoint_service_object;

truncate table ref_service_endpoint_service_connection_module;

truncate table ref_service_endpoint_physical_router;






truncate table service_group;




truncate table parent_service_group_project;

truncate table parent_service_group_policy_management;




truncate table service_health_check;


truncate table ref_service_health_check_service_instance;



truncate table parent_service_health_check_project;




truncate table service_instance;


truncate table ref_service_instance_service_template;

truncate table ref_service_instance_instance_ip;



truncate table parent_service_instance_project;




truncate table service_object;







truncate table service_template;


truncate table ref_service_template_service_appliance_set;



truncate table parent_service_template_domain;




truncate table subnet;


truncate table ref_subnet_virtual_machine_interface;






truncate table tag;


truncate table ref_tag_tag_type;



truncate table parent_tag_config_root;

truncate table parent_tag_project;




truncate table tag_type;







truncate table user;




truncate table parent_user_project;




truncate table virtual_DNS_record;




truncate table parent_virtual_DNS_record_virtual_DNS;




truncate table virtual_DNS;




truncate table parent_virtual_DNS_domain;




truncate table virtual_ip;


truncate table ref_virtual_ip_loadbalancer_pool;

truncate table ref_virtual_ip_virtual_machine_interface;



truncate table parent_virtual_ip_project;




truncate table virtual_machine_interface;


truncate table ref_virtual_machine_interface_virtual_machine_interface;

truncate table ref_virtual_machine_interface_interface_route_table;

truncate table ref_virtual_machine_interface_physical_interface;

truncate table ref_virtual_machine_interface_bgp_router;

truncate table ref_virtual_machine_interface_qos_config;

truncate table ref_virtual_machine_interface_port_tuple;

truncate table ref_virtual_machine_interface_routing_instance;

truncate table ref_virtual_machine_interface_service_health_check;

truncate table ref_virtual_machine_interface_security_group;

truncate table ref_virtual_machine_interface_virtual_network;

truncate table ref_virtual_machine_interface_bridge_domain;

truncate table ref_virtual_machine_interface_service_endpoint;

truncate table ref_virtual_machine_interface_virtual_machine;

truncate table ref_virtual_machine_interface_security_logging_object;



truncate table parent_virtual_machine_interface_project;

truncate table parent_virtual_machine_interface_virtual_machine;

truncate table parent_virtual_machine_interface_virtual_router;




truncate table virtual_machine;


truncate table ref_virtual_machine_service_instance;






truncate table virtual_network;


truncate table ref_virtual_network_route_table;

truncate table ref_virtual_network_virtual_network;

truncate table ref_virtual_network_bgpvpn;

truncate table ref_virtual_network_network_ipam;

truncate table ref_virtual_network_security_logging_object;

truncate table ref_virtual_network_network_policy;

truncate table ref_virtual_network_qos_config;



truncate table parent_virtual_network_project;




truncate table virtual_router;


truncate table ref_virtual_router_network_ipam;

truncate table ref_virtual_router_virtual_machine;



truncate table parent_virtual_router_global_system_config;




truncate table appformix_node_role;







truncate table baremetal_node;







truncate table baremetal_port;







truncate table contrail_analytics_database_node_role;







truncate table contrail_analytics_node;







truncate table contrail_cluster;







truncate table contrail_controller_node_role;







truncate table controller_node_role;







truncate table dashboard;







truncate table keypair;







truncate table kubernetes_cluster;







truncate table kubernetes_node;







truncate table location;







truncate table node;







truncate table openstack_cluster;







truncate table openstack_compute_node_role;







truncate table openstack_storage_node_role;







truncate table vpn_group;


truncate table ref_vpn_group_location;






truncate table widget;








SET FOREIGN_KEY_CHECKS=1;