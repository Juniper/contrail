-- Code generated by contrailschema tool from template sql_cleanup_psql.tmpl; DO NOT EDIT.

DROP PUBLICATION IF EXISTS "syncpub";

TRUNCATE TABLE metadata, int_pool, ipaddress_pool CASCADE;


TRUNCATE TABLE access_control_list CASCADE;

TRUNCATE TABLE address_group CASCADE;

TRUNCATE TABLE alarm CASCADE;

TRUNCATE TABLE alias_ip_pool CASCADE;

TRUNCATE TABLE alias_ip CASCADE;

TRUNCATE TABLE analytics_node CASCADE;

TRUNCATE TABLE api_access_list CASCADE;

TRUNCATE TABLE application_policy_set CASCADE;

TRUNCATE TABLE bgp_as_a_service CASCADE;

TRUNCATE TABLE bgp_router CASCADE;

TRUNCATE TABLE bgpvpn CASCADE;

TRUNCATE TABLE bridge_domain CASCADE;

TRUNCATE TABLE card CASCADE;

TRUNCATE TABLE config_node CASCADE;

TRUNCATE TABLE config_root CASCADE;

TRUNCATE TABLE customer_attachment CASCADE;

TRUNCATE TABLE database_node CASCADE;

TRUNCATE TABLE device_image CASCADE;

TRUNCATE TABLE discovery_service_assignment CASCADE;

TRUNCATE TABLE domain CASCADE;

TRUNCATE TABLE dsa_rule CASCADE;

TRUNCATE TABLE e2_service_provider CASCADE;

TRUNCATE TABLE fabric_namespace CASCADE;

TRUNCATE TABLE fabric CASCADE;

TRUNCATE TABLE firewall_policy CASCADE;

TRUNCATE TABLE firewall_rule CASCADE;

TRUNCATE TABLE floating_ip_pool CASCADE;

TRUNCATE TABLE floating_ip CASCADE;

TRUNCATE TABLE forwarding_class CASCADE;

TRUNCATE TABLE global_analytics_config CASCADE;

TRUNCATE TABLE global_qos_config CASCADE;

TRUNCATE TABLE global_system_config CASCADE;

TRUNCATE TABLE global_vrouter_config CASCADE;

TRUNCATE TABLE hardware CASCADE;

TRUNCATE TABLE instance_ip CASCADE;

TRUNCATE TABLE interface_route_table CASCADE;

TRUNCATE TABLE job_template CASCADE;

TRUNCATE TABLE link_aggregation_group CASCADE;

TRUNCATE TABLE loadbalancer_healthmonitor CASCADE;

TRUNCATE TABLE loadbalancer_listener CASCADE;

TRUNCATE TABLE loadbalancer_member CASCADE;

TRUNCATE TABLE loadbalancer_pool CASCADE;

TRUNCATE TABLE loadbalancer CASCADE;

TRUNCATE TABLE logical_interface CASCADE;

TRUNCATE TABLE logical_router CASCADE;

TRUNCATE TABLE multicast_group_address CASCADE;

TRUNCATE TABLE multicast_group CASCADE;

TRUNCATE TABLE multicast_policy CASCADE;

TRUNCATE TABLE namespace CASCADE;

TRUNCATE TABLE network_device_config CASCADE;

TRUNCATE TABLE network_ipam CASCADE;

TRUNCATE TABLE network_policy CASCADE;

TRUNCATE TABLE node_profile CASCADE;

TRUNCATE TABLE peering_policy CASCADE;

TRUNCATE TABLE physical_interface CASCADE;

TRUNCATE TABLE physical_router CASCADE;

TRUNCATE TABLE policy_management CASCADE;

TRUNCATE TABLE port_tuple CASCADE;

TRUNCATE TABLE project CASCADE;

TRUNCATE TABLE provider_attachment CASCADE;

TRUNCATE TABLE qos_config CASCADE;

TRUNCATE TABLE qos_queue CASCADE;

TRUNCATE TABLE role_config CASCADE;

TRUNCATE TABLE route_aggregate CASCADE;

TRUNCATE TABLE route_table CASCADE;

TRUNCATE TABLE route_target CASCADE;

TRUNCATE TABLE routing_instance CASCADE;

TRUNCATE TABLE routing_policy CASCADE;

TRUNCATE TABLE security_group CASCADE;

TRUNCATE TABLE security_logging_object CASCADE;

TRUNCATE TABLE service_appliance CASCADE;

TRUNCATE TABLE service_appliance_set CASCADE;

TRUNCATE TABLE service_connection_module CASCADE;

TRUNCATE TABLE service_endpoint CASCADE;

TRUNCATE TABLE service_group CASCADE;

TRUNCATE TABLE service_health_check CASCADE;

TRUNCATE TABLE service_instance CASCADE;

TRUNCATE TABLE service_object CASCADE;

TRUNCATE TABLE service_template CASCADE;

TRUNCATE TABLE structured_syslog_application_record CASCADE;

TRUNCATE TABLE structured_syslog_config CASCADE;

TRUNCATE TABLE structured_syslog_hostname_record CASCADE;

TRUNCATE TABLE structured_syslog_message CASCADE;

TRUNCATE TABLE structured_syslog_sla_profile CASCADE;

TRUNCATE TABLE sub_cluster CASCADE;

TRUNCATE TABLE subnet CASCADE;

TRUNCATE TABLE tag CASCADE;

TRUNCATE TABLE tag_type CASCADE;

TRUNCATE TABLE virtual_dns_record CASCADE;

TRUNCATE TABLE virtual_dns CASCADE;

TRUNCATE TABLE virtual_ip CASCADE;

TRUNCATE TABLE virtual_machine_interface CASCADE;

TRUNCATE TABLE virtual_machine CASCADE;

TRUNCATE TABLE virtual_network CASCADE;

TRUNCATE TABLE virtual_router CASCADE;

TRUNCATE TABLE appformix_bare_host_node CASCADE;

TRUNCATE TABLE appformix_cluster CASCADE;

TRUNCATE TABLE appformix_compute_node CASCADE;

TRUNCATE TABLE appformix_controller_node CASCADE;

TRUNCATE TABLE appformix_openstack_node CASCADE;

TRUNCATE TABLE baremetal_node CASCADE;

TRUNCATE TABLE baremetal_port CASCADE;

TRUNCATE TABLE cloud CASCADE;

TRUNCATE TABLE cloud_account CASCADE;

TRUNCATE TABLE cloud_private_subnet CASCADE;

TRUNCATE TABLE cloud_project CASCADE;

TRUNCATE TABLE cloud_region CASCADE;

TRUNCATE TABLE cloud_security_group CASCADE;

TRUNCATE TABLE cloud_security_group_rule CASCADE;

TRUNCATE TABLE cloud_user CASCADE;

TRUNCATE TABLE virtual_cloud CASCADE;

TRUNCATE TABLE contrail_analytics_database_node CASCADE;

TRUNCATE TABLE contrail_analytics_node CASCADE;

TRUNCATE TABLE contrail_cluster CASCADE;

TRUNCATE TABLE contrail_config_database_node CASCADE;

TRUNCATE TABLE contrail_config_node CASCADE;

TRUNCATE TABLE contrail_control_node CASCADE;

TRUNCATE TABLE contrail_multicloud_gw_node CASCADE;

TRUNCATE TABLE contrail_service_node CASCADE;

TRUNCATE TABLE contrail_storage_node CASCADE;

TRUNCATE TABLE contrail_vrouter_node CASCADE;

TRUNCATE TABLE contrail_webui_node CASCADE;

TRUNCATE TABLE contrail_ztp_dhcp_node CASCADE;

TRUNCATE TABLE contrail_ztp_tftp_node CASCADE;

TRUNCATE TABLE credential CASCADE;

TRUNCATE TABLE dashboard CASCADE;

TRUNCATE TABLE endpoint CASCADE;

TRUNCATE TABLE flavor CASCADE;

TRUNCATE TABLE os_image CASCADE;

TRUNCATE TABLE keypair CASCADE;

TRUNCATE TABLE kubernetes_cluster CASCADE;

TRUNCATE TABLE kubernetes_kubemanager_node CASCADE;

TRUNCATE TABLE kubernetes_master_node CASCADE;

TRUNCATE TABLE kubernetes_node CASCADE;

TRUNCATE TABLE node CASCADE;

TRUNCATE TABLE openstack_baremetal_node CASCADE;

TRUNCATE TABLE openstack_cluster CASCADE;

TRUNCATE TABLE openstack_compute_node CASCADE;

TRUNCATE TABLE openstack_control_node CASCADE;

TRUNCATE TABLE openstack_monitoring_node CASCADE;

TRUNCATE TABLE openstack_network_node CASCADE;

TRUNCATE TABLE openstack_storage_node CASCADE;

TRUNCATE TABLE port CASCADE;

TRUNCATE TABLE port_group CASCADE;

TRUNCATE TABLE server CASCADE;

TRUNCATE TABLE vcenter CASCADE;

TRUNCATE TABLE vcenter_compute CASCADE;

TRUNCATE TABLE vcenter_plugin_node CASCADE;

TRUNCATE TABLE vpn_group CASCADE;

TRUNCATE TABLE widget CASCADE;

