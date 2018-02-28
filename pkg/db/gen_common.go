package db

func (db *DB) initQueryBuilders() {
	queryBuilders := map[string]*QueryBuilder{}
	db.queryBuilders = queryBuilders

	queryBuilders["access_control_list"] = NewQueryBuilder(db.Dialect,
		"access_control_list", AccessControlListFields,
		AccessControlListRefFields,
		AccessControlListBackRefFields)

	queryBuilders["address_group"] = NewQueryBuilder(db.Dialect,
		"address_group", AddressGroupFields,
		AddressGroupRefFields,
		AddressGroupBackRefFields)

	queryBuilders["alarm"] = NewQueryBuilder(db.Dialect,
		"alarm", AlarmFields,
		AlarmRefFields,
		AlarmBackRefFields)

	queryBuilders["alias_ip_pool"] = NewQueryBuilder(db.Dialect,
		"alias_ip_pool", AliasIPPoolFields,
		AliasIPPoolRefFields,
		AliasIPPoolBackRefFields)

	queryBuilders["alias_ip"] = NewQueryBuilder(db.Dialect,
		"alias_ip", AliasIPFields,
		AliasIPRefFields,
		AliasIPBackRefFields)

	queryBuilders["analytics_node"] = NewQueryBuilder(db.Dialect,
		"analytics_node", AnalyticsNodeFields,
		AnalyticsNodeRefFields,
		AnalyticsNodeBackRefFields)

	queryBuilders["api_access_list"] = NewQueryBuilder(db.Dialect,
		"api_access_list", APIAccessListFields,
		APIAccessListRefFields,
		APIAccessListBackRefFields)

	queryBuilders["application_policy_set"] = NewQueryBuilder(db.Dialect,
		"application_policy_set", ApplicationPolicySetFields,
		ApplicationPolicySetRefFields,
		ApplicationPolicySetBackRefFields)

	queryBuilders["bgp_as_a_service"] = NewQueryBuilder(db.Dialect,
		"bgp_as_a_service", BGPAsAServiceFields,
		BGPAsAServiceRefFields,
		BGPAsAServiceBackRefFields)

	queryBuilders["bgp_router"] = NewQueryBuilder(db.Dialect,
		"bgp_router", BGPRouterFields,
		BGPRouterRefFields,
		BGPRouterBackRefFields)

	queryBuilders["bgpvpn"] = NewQueryBuilder(db.Dialect,
		"bgpvpn", BGPVPNFields,
		BGPVPNRefFields,
		BGPVPNBackRefFields)

	queryBuilders["bridge_domain"] = NewQueryBuilder(db.Dialect,
		"bridge_domain", BridgeDomainFields,
		BridgeDomainRefFields,
		BridgeDomainBackRefFields)

	queryBuilders["config_node"] = NewQueryBuilder(db.Dialect,
		"config_node", ConfigNodeFields,
		ConfigNodeRefFields,
		ConfigNodeBackRefFields)

	queryBuilders["config_root"] = NewQueryBuilder(db.Dialect,
		"config_root", ConfigRootFields,
		ConfigRootRefFields,
		ConfigRootBackRefFields)

	queryBuilders["customer_attachment"] = NewQueryBuilder(db.Dialect,
		"customer_attachment", CustomerAttachmentFields,
		CustomerAttachmentRefFields,
		CustomerAttachmentBackRefFields)

	queryBuilders["database_node"] = NewQueryBuilder(db.Dialect,
		"database_node", DatabaseNodeFields,
		DatabaseNodeRefFields,
		DatabaseNodeBackRefFields)

	queryBuilders["discovery_service_assignment"] = NewQueryBuilder(db.Dialect,
		"discovery_service_assignment", DiscoveryServiceAssignmentFields,
		DiscoveryServiceAssignmentRefFields,
		DiscoveryServiceAssignmentBackRefFields)

	queryBuilders["domain"] = NewQueryBuilder(db.Dialect,
		"domain", DomainFields,
		DomainRefFields,
		DomainBackRefFields)

	queryBuilders["dsa_rule"] = NewQueryBuilder(db.Dialect,
		"dsa_rule", DsaRuleFields,
		DsaRuleRefFields,
		DsaRuleBackRefFields)

	queryBuilders["e2_service_provider"] = NewQueryBuilder(db.Dialect,
		"e2_service_provider", E2ServiceProviderFields,
		E2ServiceProviderRefFields,
		E2ServiceProviderBackRefFields)

	queryBuilders["firewall_policy"] = NewQueryBuilder(db.Dialect,
		"firewall_policy", FirewallPolicyFields,
		FirewallPolicyRefFields,
		FirewallPolicyBackRefFields)

	queryBuilders["firewall_rule"] = NewQueryBuilder(db.Dialect,
		"firewall_rule", FirewallRuleFields,
		FirewallRuleRefFields,
		FirewallRuleBackRefFields)

	queryBuilders["floating_ip_pool"] = NewQueryBuilder(db.Dialect,
		"floating_ip_pool", FloatingIPPoolFields,
		FloatingIPPoolRefFields,
		FloatingIPPoolBackRefFields)

	queryBuilders["floating_ip"] = NewQueryBuilder(db.Dialect,
		"floating_ip", FloatingIPFields,
		FloatingIPRefFields,
		FloatingIPBackRefFields)

	queryBuilders["forwarding_class"] = NewQueryBuilder(db.Dialect,
		"forwarding_class", ForwardingClassFields,
		ForwardingClassRefFields,
		ForwardingClassBackRefFields)

	queryBuilders["global_qos_config"] = NewQueryBuilder(db.Dialect,
		"global_qos_config", GlobalQosConfigFields,
		GlobalQosConfigRefFields,
		GlobalQosConfigBackRefFields)

	queryBuilders["global_system_config"] = NewQueryBuilder(db.Dialect,
		"global_system_config", GlobalSystemConfigFields,
		GlobalSystemConfigRefFields,
		GlobalSystemConfigBackRefFields)

	queryBuilders["global_vrouter_config"] = NewQueryBuilder(db.Dialect,
		"global_vrouter_config", GlobalVrouterConfigFields,
		GlobalVrouterConfigRefFields,
		GlobalVrouterConfigBackRefFields)

	queryBuilders["instance_ip"] = NewQueryBuilder(db.Dialect,
		"instance_ip", InstanceIPFields,
		InstanceIPRefFields,
		InstanceIPBackRefFields)

	queryBuilders["interface_route_table"] = NewQueryBuilder(db.Dialect,
		"interface_route_table", InterfaceRouteTableFields,
		InterfaceRouteTableRefFields,
		InterfaceRouteTableBackRefFields)

	queryBuilders["loadbalancer_healthmonitor"] = NewQueryBuilder(db.Dialect,
		"loadbalancer_healthmonitor", LoadbalancerHealthmonitorFields,
		LoadbalancerHealthmonitorRefFields,
		LoadbalancerHealthmonitorBackRefFields)

	queryBuilders["loadbalancer_listener"] = NewQueryBuilder(db.Dialect,
		"loadbalancer_listener", LoadbalancerListenerFields,
		LoadbalancerListenerRefFields,
		LoadbalancerListenerBackRefFields)

	queryBuilders["loadbalancer_member"] = NewQueryBuilder(db.Dialect,
		"loadbalancer_member", LoadbalancerMemberFields,
		LoadbalancerMemberRefFields,
		LoadbalancerMemberBackRefFields)

	queryBuilders["loadbalancer_pool"] = NewQueryBuilder(db.Dialect,
		"loadbalancer_pool", LoadbalancerPoolFields,
		LoadbalancerPoolRefFields,
		LoadbalancerPoolBackRefFields)

	queryBuilders["loadbalancer"] = NewQueryBuilder(db.Dialect,
		"loadbalancer", LoadbalancerFields,
		LoadbalancerRefFields,
		LoadbalancerBackRefFields)

	queryBuilders["logical_interface"] = NewQueryBuilder(db.Dialect,
		"logical_interface", LogicalInterfaceFields,
		LogicalInterfaceRefFields,
		LogicalInterfaceBackRefFields)

	queryBuilders["logical_router"] = NewQueryBuilder(db.Dialect,
		"logical_router", LogicalRouterFields,
		LogicalRouterRefFields,
		LogicalRouterBackRefFields)

	queryBuilders["namespace"] = NewQueryBuilder(db.Dialect,
		"namespace", NamespaceFields,
		NamespaceRefFields,
		NamespaceBackRefFields)

	queryBuilders["network_device_config"] = NewQueryBuilder(db.Dialect,
		"network_device_config", NetworkDeviceConfigFields,
		NetworkDeviceConfigRefFields,
		NetworkDeviceConfigBackRefFields)

	queryBuilders["network_ipam"] = NewQueryBuilder(db.Dialect,
		"network_ipam", NetworkIpamFields,
		NetworkIpamRefFields,
		NetworkIpamBackRefFields)

	queryBuilders["network_policy"] = NewQueryBuilder(db.Dialect,
		"network_policy", NetworkPolicyFields,
		NetworkPolicyRefFields,
		NetworkPolicyBackRefFields)

	queryBuilders["peering_policy"] = NewQueryBuilder(db.Dialect,
		"peering_policy", PeeringPolicyFields,
		PeeringPolicyRefFields,
		PeeringPolicyBackRefFields)

	queryBuilders["physical_interface"] = NewQueryBuilder(db.Dialect,
		"physical_interface", PhysicalInterfaceFields,
		PhysicalInterfaceRefFields,
		PhysicalInterfaceBackRefFields)

	queryBuilders["physical_router"] = NewQueryBuilder(db.Dialect,
		"physical_router", PhysicalRouterFields,
		PhysicalRouterRefFields,
		PhysicalRouterBackRefFields)

	queryBuilders["policy_management"] = NewQueryBuilder(db.Dialect,
		"policy_management", PolicyManagementFields,
		PolicyManagementRefFields,
		PolicyManagementBackRefFields)

	queryBuilders["port_tuple"] = NewQueryBuilder(db.Dialect,
		"port_tuple", PortTupleFields,
		PortTupleRefFields,
		PortTupleBackRefFields)

	queryBuilders["project"] = NewQueryBuilder(db.Dialect,
		"project", ProjectFields,
		ProjectRefFields,
		ProjectBackRefFields)

	queryBuilders["provider_attachment"] = NewQueryBuilder(db.Dialect,
		"provider_attachment", ProviderAttachmentFields,
		ProviderAttachmentRefFields,
		ProviderAttachmentBackRefFields)

	queryBuilders["qos_config"] = NewQueryBuilder(db.Dialect,
		"qos_config", QosConfigFields,
		QosConfigRefFields,
		QosConfigBackRefFields)

	queryBuilders["qos_queue"] = NewQueryBuilder(db.Dialect,
		"qos_queue", QosQueueFields,
		QosQueueRefFields,
		QosQueueBackRefFields)

	queryBuilders["route_aggregate"] = NewQueryBuilder(db.Dialect,
		"route_aggregate", RouteAggregateFields,
		RouteAggregateRefFields,
		RouteAggregateBackRefFields)

	queryBuilders["route_table"] = NewQueryBuilder(db.Dialect,
		"route_table", RouteTableFields,
		RouteTableRefFields,
		RouteTableBackRefFields)

	queryBuilders["route_target"] = NewQueryBuilder(db.Dialect,
		"route_target", RouteTargetFields,
		RouteTargetRefFields,
		RouteTargetBackRefFields)

	queryBuilders["routing_instance"] = NewQueryBuilder(db.Dialect,
		"routing_instance", RoutingInstanceFields,
		RoutingInstanceRefFields,
		RoutingInstanceBackRefFields)

	queryBuilders["routing_policy"] = NewQueryBuilder(db.Dialect,
		"routing_policy", RoutingPolicyFields,
		RoutingPolicyRefFields,
		RoutingPolicyBackRefFields)

	queryBuilders["security_group"] = NewQueryBuilder(db.Dialect,
		"security_group", SecurityGroupFields,
		SecurityGroupRefFields,
		SecurityGroupBackRefFields)

	queryBuilders["security_logging_object"] = NewQueryBuilder(db.Dialect,
		"security_logging_object", SecurityLoggingObjectFields,
		SecurityLoggingObjectRefFields,
		SecurityLoggingObjectBackRefFields)

	queryBuilders["service_appliance"] = NewQueryBuilder(db.Dialect,
		"service_appliance", ServiceApplianceFields,
		ServiceApplianceRefFields,
		ServiceApplianceBackRefFields)

	queryBuilders["service_appliance_set"] = NewQueryBuilder(db.Dialect,
		"service_appliance_set", ServiceApplianceSetFields,
		ServiceApplianceSetRefFields,
		ServiceApplianceSetBackRefFields)

	queryBuilders["service_connection_module"] = NewQueryBuilder(db.Dialect,
		"service_connection_module", ServiceConnectionModuleFields,
		ServiceConnectionModuleRefFields,
		ServiceConnectionModuleBackRefFields)

	queryBuilders["service_endpoint"] = NewQueryBuilder(db.Dialect,
		"service_endpoint", ServiceEndpointFields,
		ServiceEndpointRefFields,
		ServiceEndpointBackRefFields)

	queryBuilders["service_group"] = NewQueryBuilder(db.Dialect,
		"service_group", ServiceGroupFields,
		ServiceGroupRefFields,
		ServiceGroupBackRefFields)

	queryBuilders["service_health_check"] = NewQueryBuilder(db.Dialect,
		"service_health_check", ServiceHealthCheckFields,
		ServiceHealthCheckRefFields,
		ServiceHealthCheckBackRefFields)

	queryBuilders["service_instance"] = NewQueryBuilder(db.Dialect,
		"service_instance", ServiceInstanceFields,
		ServiceInstanceRefFields,
		ServiceInstanceBackRefFields)

	queryBuilders["service_object"] = NewQueryBuilder(db.Dialect,
		"service_object", ServiceObjectFields,
		ServiceObjectRefFields,
		ServiceObjectBackRefFields)

	queryBuilders["service_template"] = NewQueryBuilder(db.Dialect,
		"service_template", ServiceTemplateFields,
		ServiceTemplateRefFields,
		ServiceTemplateBackRefFields)

	queryBuilders["subnet"] = NewQueryBuilder(db.Dialect,
		"subnet", SubnetFields,
		SubnetRefFields,
		SubnetBackRefFields)

	queryBuilders["tag"] = NewQueryBuilder(db.Dialect,
		"tag", TagFields,
		TagRefFields,
		TagBackRefFields)

	queryBuilders["tag_type"] = NewQueryBuilder(db.Dialect,
		"tag_type", TagTypeFields,
		TagTypeRefFields,
		TagTypeBackRefFields)

	queryBuilders["user"] = NewQueryBuilder(db.Dialect,
		"user", UserFields,
		UserRefFields,
		UserBackRefFields)

	queryBuilders["virtual_DNS_record"] = NewQueryBuilder(db.Dialect,
		"virtual_DNS_record", VirtualDNSRecordFields,
		VirtualDNSRecordRefFields,
		VirtualDNSRecordBackRefFields)

	queryBuilders["virtual_DNS"] = NewQueryBuilder(db.Dialect,
		"virtual_DNS", VirtualDNSFields,
		VirtualDNSRefFields,
		VirtualDNSBackRefFields)

	queryBuilders["virtual_ip"] = NewQueryBuilder(db.Dialect,
		"virtual_ip", VirtualIPFields,
		VirtualIPRefFields,
		VirtualIPBackRefFields)

	queryBuilders["virtual_machine_interface"] = NewQueryBuilder(db.Dialect,
		"virtual_machine_interface", VirtualMachineInterfaceFields,
		VirtualMachineInterfaceRefFields,
		VirtualMachineInterfaceBackRefFields)

	queryBuilders["virtual_machine"] = NewQueryBuilder(db.Dialect,
		"virtual_machine", VirtualMachineFields,
		VirtualMachineRefFields,
		VirtualMachineBackRefFields)

	queryBuilders["virtual_network"] = NewQueryBuilder(db.Dialect,
		"virtual_network", VirtualNetworkFields,
		VirtualNetworkRefFields,
		VirtualNetworkBackRefFields)

	queryBuilders["virtual_router"] = NewQueryBuilder(db.Dialect,
		"virtual_router", VirtualRouterFields,
		VirtualRouterRefFields,
		VirtualRouterBackRefFields)

	queryBuilders["appformix_node"] = NewQueryBuilder(db.Dialect,
		"appformix_node", AppformixNodeFields,
		AppformixNodeRefFields,
		AppformixNodeBackRefFields)

	queryBuilders["baremetal_node"] = NewQueryBuilder(db.Dialect,
		"baremetal_node", BaremetalNodeFields,
		BaremetalNodeRefFields,
		BaremetalNodeBackRefFields)

	queryBuilders["baremetal_port"] = NewQueryBuilder(db.Dialect,
		"baremetal_port", BaremetalPortFields,
		BaremetalPortRefFields,
		BaremetalPortBackRefFields)

	queryBuilders["contrail_analytics_database_node"] = NewQueryBuilder(db.Dialect,
		"contrail_analytics_database_node", ContrailAnalyticsDatabaseNodeFields,
		ContrailAnalyticsDatabaseNodeRefFields,
		ContrailAnalyticsDatabaseNodeBackRefFields)

	queryBuilders["contrail_analytics_node"] = NewQueryBuilder(db.Dialect,
		"contrail_analytics_node", ContrailAnalyticsNodeFields,
		ContrailAnalyticsNodeRefFields,
		ContrailAnalyticsNodeBackRefFields)

	queryBuilders["contrail_cluster"] = NewQueryBuilder(db.Dialect,
		"contrail_cluster", ContrailClusterFields,
		ContrailClusterRefFields,
		ContrailClusterBackRefFields)

	queryBuilders["contrail_config_database_node"] = NewQueryBuilder(db.Dialect,
		"contrail_config_database_node", ContrailConfigDatabaseNodeFields,
		ContrailConfigDatabaseNodeRefFields,
		ContrailConfigDatabaseNodeBackRefFields)

	queryBuilders["contrail_config_node"] = NewQueryBuilder(db.Dialect,
		"contrail_config_node", ContrailConfigNodeFields,
		ContrailConfigNodeRefFields,
		ContrailConfigNodeBackRefFields)

	queryBuilders["contrail_control_node"] = NewQueryBuilder(db.Dialect,
		"contrail_control_node", ContrailControlNodeFields,
		ContrailControlNodeRefFields,
		ContrailControlNodeBackRefFields)

	queryBuilders["contrail_storage_node"] = NewQueryBuilder(db.Dialect,
		"contrail_storage_node", ContrailStorageNodeFields,
		ContrailStorageNodeRefFields,
		ContrailStorageNodeBackRefFields)

	queryBuilders["contrail_vrouter_node"] = NewQueryBuilder(db.Dialect,
		"contrail_vrouter_node", ContrailVrouterNodeFields,
		ContrailVrouterNodeRefFields,
		ContrailVrouterNodeBackRefFields)

	queryBuilders["contrail_webui_node"] = NewQueryBuilder(db.Dialect,
		"contrail_webui_node", ContrailWebuiNodeFields,
		ContrailWebuiNodeRefFields,
		ContrailWebuiNodeBackRefFields)

	queryBuilders["dashboard"] = NewQueryBuilder(db.Dialect,
		"dashboard", DashboardFields,
		DashboardRefFields,
		DashboardBackRefFields)

	queryBuilders["flavor"] = NewQueryBuilder(db.Dialect,
		"flavor", FlavorFields,
		FlavorRefFields,
		FlavorBackRefFields)

	queryBuilders["os_image"] = NewQueryBuilder(db.Dialect,
		"os_image", OsImageFields,
		OsImageRefFields,
		OsImageBackRefFields)

	queryBuilders["keypair"] = NewQueryBuilder(db.Dialect,
		"keypair", KeypairFields,
		KeypairRefFields,
		KeypairBackRefFields)

	queryBuilders["kubernetes_master_node"] = NewQueryBuilder(db.Dialect,
		"kubernetes_master_node", KubernetesMasterNodeFields,
		KubernetesMasterNodeRefFields,
		KubernetesMasterNodeBackRefFields)

	queryBuilders["kubernetes_node"] = NewQueryBuilder(db.Dialect,
		"kubernetes_node", KubernetesNodeFields,
		KubernetesNodeRefFields,
		KubernetesNodeBackRefFields)

	queryBuilders["location"] = NewQueryBuilder(db.Dialect,
		"location", LocationFields,
		LocationRefFields,
		LocationBackRefFields)

	queryBuilders["node"] = NewQueryBuilder(db.Dialect,
		"node", NodeFields,
		NodeRefFields,
		NodeBackRefFields)

	queryBuilders["openstack_compute_node"] = NewQueryBuilder(db.Dialect,
		"openstack_compute_node", OpenstackComputeNodeFields,
		OpenstackComputeNodeRefFields,
		OpenstackComputeNodeBackRefFields)

	queryBuilders["openstack_control_node"] = NewQueryBuilder(db.Dialect,
		"openstack_control_node", OpenstackControlNodeFields,
		OpenstackControlNodeRefFields,
		OpenstackControlNodeBackRefFields)

	queryBuilders["openstack_monitoring_node"] = NewQueryBuilder(db.Dialect,
		"openstack_monitoring_node", OpenstackMonitoringNodeFields,
		OpenstackMonitoringNodeRefFields,
		OpenstackMonitoringNodeBackRefFields)

	queryBuilders["openstack_network_node"] = NewQueryBuilder(db.Dialect,
		"openstack_network_node", OpenstackNetworkNodeFields,
		OpenstackNetworkNodeRefFields,
		OpenstackNetworkNodeBackRefFields)

	queryBuilders["openstack_storage_node"] = NewQueryBuilder(db.Dialect,
		"openstack_storage_node", OpenstackStorageNodeFields,
		OpenstackStorageNodeRefFields,
		OpenstackStorageNodeBackRefFields)

	queryBuilders["server"] = NewQueryBuilder(db.Dialect,
		"server", ServerFields,
		ServerRefFields,
		ServerBackRefFields)

	queryBuilders["vpn_group"] = NewQueryBuilder(db.Dialect,
		"vpn_group", VPNGroupFields,
		VPNGroupRefFields,
		VPNGroupBackRefFields)

	queryBuilders["widget"] = NewQueryBuilder(db.Dialect,
		"widget", WidgetFields,
		WidgetRefFields,
		WidgetBackRefFields)

}
