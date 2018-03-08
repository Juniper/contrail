# Protocol Documentation
<a name="top"/>

## Table of Contents

- [github.com/Juniper/contrail/pkg/models/generated.proto](#github.com/Juniper/contrail/pkg/models/generated.proto)
    - [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList)
    - [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList)
    - [AclEntriesType](#github.com.Juniper.contrail.pkg.models.AclEntriesType)
    - [AclRuleType](#github.com.Juniper.contrail.pkg.models.AclRuleType)
    - [ActionListType](#github.com.Juniper.contrail.pkg.models.ActionListType)
    - [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup)
    - [AddressType](#github.com.Juniper.contrail.pkg.models.AddressType)
    - [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm)
    - [AlarmAndList](#github.com.Juniper.contrail.pkg.models.AlarmAndList)
    - [AlarmExpression](#github.com.Juniper.contrail.pkg.models.AlarmExpression)
    - [AlarmOperand2](#github.com.Juniper.contrail.pkg.models.AlarmOperand2)
    - [AlarmOrList](#github.com.Juniper.contrail.pkg.models.AlarmOrList)
    - [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP)
    - [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool)
    - [AliasIPProjectRef](#github.com.Juniper.contrail.pkg.models.AliasIPProjectRef)
    - [AliasIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.AliasIPVirtualMachineInterfaceRef)
    - [AllocationPoolType](#github.com.Juniper.contrail.pkg.models.AllocationPoolType)
    - [AllowedAddressPair](#github.com.Juniper.contrail.pkg.models.AllowedAddressPair)
    - [AllowedAddressPairs](#github.com.Juniper.contrail.pkg.models.AllowedAddressPairs)
    - [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode)
    - [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode)
    - [AppformixNodeNodeRef](#github.com.Juniper.contrail.pkg.models.AppformixNodeNodeRef)
    - [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet)
    - [ApplicationPolicySetFirewallPolicyRef](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySetFirewallPolicyRef)
    - [ApplicationPolicySetGlobalVrouterConfigRef](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySetGlobalVrouterConfigRef)
    - [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService)
    - [BGPAsAServiceServiceHealthCheckRef](#github.com.Juniper.contrail.pkg.models.BGPAsAServiceServiceHealthCheckRef)
    - [BGPAsAServiceVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.BGPAsAServiceVirtualMachineInterfaceRef)
    - [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter)
    - [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN)
    - [BGPaaServiceParametersType](#github.com.Juniper.contrail.pkg.models.BGPaaServiceParametersType)
    - [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode)
    - [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort)
    - [BaremetalProperties](#github.com.Juniper.contrail.pkg.models.BaremetalProperties)
    - [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain)
    - [BridgeDomainMembershipType](#github.com.Juniper.contrail.pkg.models.BridgeDomainMembershipType)
    - [CommunityAttributes](#github.com.Juniper.contrail.pkg.models.CommunityAttributes)
    - [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode)
    - [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot)
    - [ConfigRootTagRef](#github.com.Juniper.contrail.pkg.models.ConfigRootTagRef)
    - [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode)
    - [ContrailAnalyticsDatabaseNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNodeNodeRef)
    - [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode)
    - [ContrailAnalyticsNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNodeNodeRef)
    - [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster)
    - [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode)
    - [ContrailConfigDatabaseNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNodeNodeRef)
    - [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode)
    - [ContrailConfigNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailConfigNodeNodeRef)
    - [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode)
    - [ContrailControlNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailControlNodeNodeRef)
    - [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode)
    - [ContrailStorageNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailStorageNodeNodeRef)
    - [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode)
    - [ContrailVrouterNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNodeNodeRef)
    - [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode)
    - [ContrailWebuiNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNodeNodeRef)
    - [ControlTrafficDscpType](#github.com.Juniper.contrail.pkg.models.ControlTrafficDscpType)
    - [CreateAPIAccessListRequest](#github.com.Juniper.contrail.pkg.models.CreateAPIAccessListRequest)
    - [CreateAPIAccessListResponse](#github.com.Juniper.contrail.pkg.models.CreateAPIAccessListResponse)
    - [CreateAccessControlListRequest](#github.com.Juniper.contrail.pkg.models.CreateAccessControlListRequest)
    - [CreateAccessControlListResponse](#github.com.Juniper.contrail.pkg.models.CreateAccessControlListResponse)
    - [CreateAddressGroupRequest](#github.com.Juniper.contrail.pkg.models.CreateAddressGroupRequest)
    - [CreateAddressGroupResponse](#github.com.Juniper.contrail.pkg.models.CreateAddressGroupResponse)
    - [CreateAlarmRequest](#github.com.Juniper.contrail.pkg.models.CreateAlarmRequest)
    - [CreateAlarmResponse](#github.com.Juniper.contrail.pkg.models.CreateAlarmResponse)
    - [CreateAliasIPPoolRequest](#github.com.Juniper.contrail.pkg.models.CreateAliasIPPoolRequest)
    - [CreateAliasIPPoolResponse](#github.com.Juniper.contrail.pkg.models.CreateAliasIPPoolResponse)
    - [CreateAliasIPRequest](#github.com.Juniper.contrail.pkg.models.CreateAliasIPRequest)
    - [CreateAliasIPResponse](#github.com.Juniper.contrail.pkg.models.CreateAliasIPResponse)
    - [CreateAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateAnalyticsNodeRequest)
    - [CreateAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateAnalyticsNodeResponse)
    - [CreateAppformixNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateAppformixNodeRequest)
    - [CreateAppformixNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateAppformixNodeResponse)
    - [CreateApplicationPolicySetRequest](#github.com.Juniper.contrail.pkg.models.CreateApplicationPolicySetRequest)
    - [CreateApplicationPolicySetResponse](#github.com.Juniper.contrail.pkg.models.CreateApplicationPolicySetResponse)
    - [CreateBGPAsAServiceRequest](#github.com.Juniper.contrail.pkg.models.CreateBGPAsAServiceRequest)
    - [CreateBGPAsAServiceResponse](#github.com.Juniper.contrail.pkg.models.CreateBGPAsAServiceResponse)
    - [CreateBGPRouterRequest](#github.com.Juniper.contrail.pkg.models.CreateBGPRouterRequest)
    - [CreateBGPRouterResponse](#github.com.Juniper.contrail.pkg.models.CreateBGPRouterResponse)
    - [CreateBGPVPNRequest](#github.com.Juniper.contrail.pkg.models.CreateBGPVPNRequest)
    - [CreateBGPVPNResponse](#github.com.Juniper.contrail.pkg.models.CreateBGPVPNResponse)
    - [CreateBaremetalNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateBaremetalNodeRequest)
    - [CreateBaremetalNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateBaremetalNodeResponse)
    - [CreateBaremetalPortRequest](#github.com.Juniper.contrail.pkg.models.CreateBaremetalPortRequest)
    - [CreateBaremetalPortResponse](#github.com.Juniper.contrail.pkg.models.CreateBaremetalPortResponse)
    - [CreateBridgeDomainRequest](#github.com.Juniper.contrail.pkg.models.CreateBridgeDomainRequest)
    - [CreateBridgeDomainResponse](#github.com.Juniper.contrail.pkg.models.CreateBridgeDomainResponse)
    - [CreateConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateConfigNodeRequest)
    - [CreateConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateConfigNodeResponse)
    - [CreateConfigRootRequest](#github.com.Juniper.contrail.pkg.models.CreateConfigRootRequest)
    - [CreateConfigRootResponse](#github.com.Juniper.contrail.pkg.models.CreateConfigRootResponse)
    - [CreateContrailAnalyticsDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsDatabaseNodeRequest)
    - [CreateContrailAnalyticsDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsDatabaseNodeResponse)
    - [CreateContrailAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsNodeRequest)
    - [CreateContrailAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsNodeResponse)
    - [CreateContrailClusterRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailClusterRequest)
    - [CreateContrailClusterResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailClusterResponse)
    - [CreateContrailConfigDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailConfigDatabaseNodeRequest)
    - [CreateContrailConfigDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailConfigDatabaseNodeResponse)
    - [CreateContrailConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailConfigNodeRequest)
    - [CreateContrailConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailConfigNodeResponse)
    - [CreateContrailControlNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailControlNodeRequest)
    - [CreateContrailControlNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailControlNodeResponse)
    - [CreateContrailStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailStorageNodeRequest)
    - [CreateContrailStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailStorageNodeResponse)
    - [CreateContrailVrouterNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailVrouterNodeRequest)
    - [CreateContrailVrouterNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailVrouterNodeResponse)
    - [CreateContrailWebuiNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateContrailWebuiNodeRequest)
    - [CreateContrailWebuiNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateContrailWebuiNodeResponse)
    - [CreateCustomerAttachmentRequest](#github.com.Juniper.contrail.pkg.models.CreateCustomerAttachmentRequest)
    - [CreateCustomerAttachmentResponse](#github.com.Juniper.contrail.pkg.models.CreateCustomerAttachmentResponse)
    - [CreateDashboardRequest](#github.com.Juniper.contrail.pkg.models.CreateDashboardRequest)
    - [CreateDashboardResponse](#github.com.Juniper.contrail.pkg.models.CreateDashboardResponse)
    - [CreateDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateDatabaseNodeRequest)
    - [CreateDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateDatabaseNodeResponse)
    - [CreateDiscoveryServiceAssignmentRequest](#github.com.Juniper.contrail.pkg.models.CreateDiscoveryServiceAssignmentRequest)
    - [CreateDiscoveryServiceAssignmentResponse](#github.com.Juniper.contrail.pkg.models.CreateDiscoveryServiceAssignmentResponse)
    - [CreateDomainRequest](#github.com.Juniper.contrail.pkg.models.CreateDomainRequest)
    - [CreateDomainResponse](#github.com.Juniper.contrail.pkg.models.CreateDomainResponse)
    - [CreateDsaRuleRequest](#github.com.Juniper.contrail.pkg.models.CreateDsaRuleRequest)
    - [CreateDsaRuleResponse](#github.com.Juniper.contrail.pkg.models.CreateDsaRuleResponse)
    - [CreateE2ServiceProviderRequest](#github.com.Juniper.contrail.pkg.models.CreateE2ServiceProviderRequest)
    - [CreateE2ServiceProviderResponse](#github.com.Juniper.contrail.pkg.models.CreateE2ServiceProviderResponse)
    - [CreateFirewallPolicyRequest](#github.com.Juniper.contrail.pkg.models.CreateFirewallPolicyRequest)
    - [CreateFirewallPolicyResponse](#github.com.Juniper.contrail.pkg.models.CreateFirewallPolicyResponse)
    - [CreateFirewallRuleRequest](#github.com.Juniper.contrail.pkg.models.CreateFirewallRuleRequest)
    - [CreateFirewallRuleResponse](#github.com.Juniper.contrail.pkg.models.CreateFirewallRuleResponse)
    - [CreateFlavorRequest](#github.com.Juniper.contrail.pkg.models.CreateFlavorRequest)
    - [CreateFlavorResponse](#github.com.Juniper.contrail.pkg.models.CreateFlavorResponse)
    - [CreateFloatingIPPoolRequest](#github.com.Juniper.contrail.pkg.models.CreateFloatingIPPoolRequest)
    - [CreateFloatingIPPoolResponse](#github.com.Juniper.contrail.pkg.models.CreateFloatingIPPoolResponse)
    - [CreateFloatingIPRequest](#github.com.Juniper.contrail.pkg.models.CreateFloatingIPRequest)
    - [CreateFloatingIPResponse](#github.com.Juniper.contrail.pkg.models.CreateFloatingIPResponse)
    - [CreateForwardingClassRequest](#github.com.Juniper.contrail.pkg.models.CreateForwardingClassRequest)
    - [CreateForwardingClassResponse](#github.com.Juniper.contrail.pkg.models.CreateForwardingClassResponse)
    - [CreateGlobalQosConfigRequest](#github.com.Juniper.contrail.pkg.models.CreateGlobalQosConfigRequest)
    - [CreateGlobalQosConfigResponse](#github.com.Juniper.contrail.pkg.models.CreateGlobalQosConfigResponse)
    - [CreateGlobalSystemConfigRequest](#github.com.Juniper.contrail.pkg.models.CreateGlobalSystemConfigRequest)
    - [CreateGlobalSystemConfigResponse](#github.com.Juniper.contrail.pkg.models.CreateGlobalSystemConfigResponse)
    - [CreateGlobalVrouterConfigRequest](#github.com.Juniper.contrail.pkg.models.CreateGlobalVrouterConfigRequest)
    - [CreateGlobalVrouterConfigResponse](#github.com.Juniper.contrail.pkg.models.CreateGlobalVrouterConfigResponse)
    - [CreateInstanceIPRequest](#github.com.Juniper.contrail.pkg.models.CreateInstanceIPRequest)
    - [CreateInstanceIPResponse](#github.com.Juniper.contrail.pkg.models.CreateInstanceIPResponse)
    - [CreateInterfaceRouteTableRequest](#github.com.Juniper.contrail.pkg.models.CreateInterfaceRouteTableRequest)
    - [CreateInterfaceRouteTableResponse](#github.com.Juniper.contrail.pkg.models.CreateInterfaceRouteTableResponse)
    - [CreateKeypairRequest](#github.com.Juniper.contrail.pkg.models.CreateKeypairRequest)
    - [CreateKeypairResponse](#github.com.Juniper.contrail.pkg.models.CreateKeypairResponse)
    - [CreateKubernetesMasterNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateKubernetesMasterNodeRequest)
    - [CreateKubernetesMasterNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateKubernetesMasterNodeResponse)
    - [CreateKubernetesNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateKubernetesNodeRequest)
    - [CreateKubernetesNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateKubernetesNodeResponse)
    - [CreateLoadbalancerHealthmonitorRequest](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerHealthmonitorRequest)
    - [CreateLoadbalancerHealthmonitorResponse](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerHealthmonitorResponse)
    - [CreateLoadbalancerListenerRequest](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerListenerRequest)
    - [CreateLoadbalancerListenerResponse](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerListenerResponse)
    - [CreateLoadbalancerMemberRequest](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerMemberRequest)
    - [CreateLoadbalancerMemberResponse](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerMemberResponse)
    - [CreateLoadbalancerPoolRequest](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerPoolRequest)
    - [CreateLoadbalancerPoolResponse](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerPoolResponse)
    - [CreateLoadbalancerRequest](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerRequest)
    - [CreateLoadbalancerResponse](#github.com.Juniper.contrail.pkg.models.CreateLoadbalancerResponse)
    - [CreateLocationRequest](#github.com.Juniper.contrail.pkg.models.CreateLocationRequest)
    - [CreateLocationResponse](#github.com.Juniper.contrail.pkg.models.CreateLocationResponse)
    - [CreateLogicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.CreateLogicalInterfaceRequest)
    - [CreateLogicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.CreateLogicalInterfaceResponse)
    - [CreateLogicalRouterRequest](#github.com.Juniper.contrail.pkg.models.CreateLogicalRouterRequest)
    - [CreateLogicalRouterResponse](#github.com.Juniper.contrail.pkg.models.CreateLogicalRouterResponse)
    - [CreateNamespaceRequest](#github.com.Juniper.contrail.pkg.models.CreateNamespaceRequest)
    - [CreateNamespaceResponse](#github.com.Juniper.contrail.pkg.models.CreateNamespaceResponse)
    - [CreateNetworkDeviceConfigRequest](#github.com.Juniper.contrail.pkg.models.CreateNetworkDeviceConfigRequest)
    - [CreateNetworkDeviceConfigResponse](#github.com.Juniper.contrail.pkg.models.CreateNetworkDeviceConfigResponse)
    - [CreateNetworkIpamRequest](#github.com.Juniper.contrail.pkg.models.CreateNetworkIpamRequest)
    - [CreateNetworkIpamResponse](#github.com.Juniper.contrail.pkg.models.CreateNetworkIpamResponse)
    - [CreateNetworkPolicyRequest](#github.com.Juniper.contrail.pkg.models.CreateNetworkPolicyRequest)
    - [CreateNetworkPolicyResponse](#github.com.Juniper.contrail.pkg.models.CreateNetworkPolicyResponse)
    - [CreateNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateNodeRequest)
    - [CreateNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateNodeResponse)
    - [CreateOpenstackComputeNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateOpenstackComputeNodeRequest)
    - [CreateOpenstackComputeNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateOpenstackComputeNodeResponse)
    - [CreateOpenstackControlNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateOpenstackControlNodeRequest)
    - [CreateOpenstackControlNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateOpenstackControlNodeResponse)
    - [CreateOpenstackMonitoringNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateOpenstackMonitoringNodeRequest)
    - [CreateOpenstackMonitoringNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateOpenstackMonitoringNodeResponse)
    - [CreateOpenstackNetworkNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateOpenstackNetworkNodeRequest)
    - [CreateOpenstackNetworkNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateOpenstackNetworkNodeResponse)
    - [CreateOpenstackStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.CreateOpenstackStorageNodeRequest)
    - [CreateOpenstackStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.CreateOpenstackStorageNodeResponse)
    - [CreateOsImageRequest](#github.com.Juniper.contrail.pkg.models.CreateOsImageRequest)
    - [CreateOsImageResponse](#github.com.Juniper.contrail.pkg.models.CreateOsImageResponse)
    - [CreatePeeringPolicyRequest](#github.com.Juniper.contrail.pkg.models.CreatePeeringPolicyRequest)
    - [CreatePeeringPolicyResponse](#github.com.Juniper.contrail.pkg.models.CreatePeeringPolicyResponse)
    - [CreatePhysicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.CreatePhysicalInterfaceRequest)
    - [CreatePhysicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.CreatePhysicalInterfaceResponse)
    - [CreatePhysicalRouterRequest](#github.com.Juniper.contrail.pkg.models.CreatePhysicalRouterRequest)
    - [CreatePhysicalRouterResponse](#github.com.Juniper.contrail.pkg.models.CreatePhysicalRouterResponse)
    - [CreatePolicyManagementRequest](#github.com.Juniper.contrail.pkg.models.CreatePolicyManagementRequest)
    - [CreatePolicyManagementResponse](#github.com.Juniper.contrail.pkg.models.CreatePolicyManagementResponse)
    - [CreatePortRequest](#github.com.Juniper.contrail.pkg.models.CreatePortRequest)
    - [CreatePortResponse](#github.com.Juniper.contrail.pkg.models.CreatePortResponse)
    - [CreatePortTupleRequest](#github.com.Juniper.contrail.pkg.models.CreatePortTupleRequest)
    - [CreatePortTupleResponse](#github.com.Juniper.contrail.pkg.models.CreatePortTupleResponse)
    - [CreateProjectRequest](#github.com.Juniper.contrail.pkg.models.CreateProjectRequest)
    - [CreateProjectResponse](#github.com.Juniper.contrail.pkg.models.CreateProjectResponse)
    - [CreateProviderAttachmentRequest](#github.com.Juniper.contrail.pkg.models.CreateProviderAttachmentRequest)
    - [CreateProviderAttachmentResponse](#github.com.Juniper.contrail.pkg.models.CreateProviderAttachmentResponse)
    - [CreateQosConfigRequest](#github.com.Juniper.contrail.pkg.models.CreateQosConfigRequest)
    - [CreateQosConfigResponse](#github.com.Juniper.contrail.pkg.models.CreateQosConfigResponse)
    - [CreateQosQueueRequest](#github.com.Juniper.contrail.pkg.models.CreateQosQueueRequest)
    - [CreateQosQueueResponse](#github.com.Juniper.contrail.pkg.models.CreateQosQueueResponse)
    - [CreateRouteAggregateRequest](#github.com.Juniper.contrail.pkg.models.CreateRouteAggregateRequest)
    - [CreateRouteAggregateResponse](#github.com.Juniper.contrail.pkg.models.CreateRouteAggregateResponse)
    - [CreateRouteTableRequest](#github.com.Juniper.contrail.pkg.models.CreateRouteTableRequest)
    - [CreateRouteTableResponse](#github.com.Juniper.contrail.pkg.models.CreateRouteTableResponse)
    - [CreateRouteTargetRequest](#github.com.Juniper.contrail.pkg.models.CreateRouteTargetRequest)
    - [CreateRouteTargetResponse](#github.com.Juniper.contrail.pkg.models.CreateRouteTargetResponse)
    - [CreateRoutingInstanceRequest](#github.com.Juniper.contrail.pkg.models.CreateRoutingInstanceRequest)
    - [CreateRoutingInstanceResponse](#github.com.Juniper.contrail.pkg.models.CreateRoutingInstanceResponse)
    - [CreateRoutingPolicyRequest](#github.com.Juniper.contrail.pkg.models.CreateRoutingPolicyRequest)
    - [CreateRoutingPolicyResponse](#github.com.Juniper.contrail.pkg.models.CreateRoutingPolicyResponse)
    - [CreateSecurityGroupRequest](#github.com.Juniper.contrail.pkg.models.CreateSecurityGroupRequest)
    - [CreateSecurityGroupResponse](#github.com.Juniper.contrail.pkg.models.CreateSecurityGroupResponse)
    - [CreateSecurityLoggingObjectRequest](#github.com.Juniper.contrail.pkg.models.CreateSecurityLoggingObjectRequest)
    - [CreateSecurityLoggingObjectResponse](#github.com.Juniper.contrail.pkg.models.CreateSecurityLoggingObjectResponse)
    - [CreateServerRequest](#github.com.Juniper.contrail.pkg.models.CreateServerRequest)
    - [CreateServerResponse](#github.com.Juniper.contrail.pkg.models.CreateServerResponse)
    - [CreateServiceApplianceRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceApplianceRequest)
    - [CreateServiceApplianceResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceApplianceResponse)
    - [CreateServiceApplianceSetRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceApplianceSetRequest)
    - [CreateServiceApplianceSetResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceApplianceSetResponse)
    - [CreateServiceConnectionModuleRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceConnectionModuleRequest)
    - [CreateServiceConnectionModuleResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceConnectionModuleResponse)
    - [CreateServiceEndpointRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceEndpointRequest)
    - [CreateServiceEndpointResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceEndpointResponse)
    - [CreateServiceGroupRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceGroupRequest)
    - [CreateServiceGroupResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceGroupResponse)
    - [CreateServiceHealthCheckRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceHealthCheckRequest)
    - [CreateServiceHealthCheckResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceHealthCheckResponse)
    - [CreateServiceInstanceRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceInstanceRequest)
    - [CreateServiceInstanceResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceInstanceResponse)
    - [CreateServiceObjectRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceObjectRequest)
    - [CreateServiceObjectResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceObjectResponse)
    - [CreateServiceTemplateRequest](#github.com.Juniper.contrail.pkg.models.CreateServiceTemplateRequest)
    - [CreateServiceTemplateResponse](#github.com.Juniper.contrail.pkg.models.CreateServiceTemplateResponse)
    - [CreateSubnetRequest](#github.com.Juniper.contrail.pkg.models.CreateSubnetRequest)
    - [CreateSubnetResponse](#github.com.Juniper.contrail.pkg.models.CreateSubnetResponse)
    - [CreateTagRequest](#github.com.Juniper.contrail.pkg.models.CreateTagRequest)
    - [CreateTagResponse](#github.com.Juniper.contrail.pkg.models.CreateTagResponse)
    - [CreateTagTypeRequest](#github.com.Juniper.contrail.pkg.models.CreateTagTypeRequest)
    - [CreateTagTypeResponse](#github.com.Juniper.contrail.pkg.models.CreateTagTypeResponse)
    - [CreateUserRequest](#github.com.Juniper.contrail.pkg.models.CreateUserRequest)
    - [CreateUserResponse](#github.com.Juniper.contrail.pkg.models.CreateUserResponse)
    - [CreateVPNGroupRequest](#github.com.Juniper.contrail.pkg.models.CreateVPNGroupRequest)
    - [CreateVPNGroupResponse](#github.com.Juniper.contrail.pkg.models.CreateVPNGroupResponse)
    - [CreateVirtualDNSRecordRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualDNSRecordRequest)
    - [CreateVirtualDNSRecordResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualDNSRecordResponse)
    - [CreateVirtualDNSRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualDNSRequest)
    - [CreateVirtualDNSResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualDNSResponse)
    - [CreateVirtualIPRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualIPRequest)
    - [CreateVirtualIPResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualIPResponse)
    - [CreateVirtualMachineInterfaceRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualMachineInterfaceRequest)
    - [CreateVirtualMachineInterfaceResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualMachineInterfaceResponse)
    - [CreateVirtualMachineRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualMachineRequest)
    - [CreateVirtualMachineResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualMachineResponse)
    - [CreateVirtualNetworkRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualNetworkRequest)
    - [CreateVirtualNetworkResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualNetworkResponse)
    - [CreateVirtualRouterRequest](#github.com.Juniper.contrail.pkg.models.CreateVirtualRouterRequest)
    - [CreateVirtualRouterResponse](#github.com.Juniper.contrail.pkg.models.CreateVirtualRouterResponse)
    - [CreateWidgetRequest](#github.com.Juniper.contrail.pkg.models.CreateWidgetRequest)
    - [CreateWidgetResponse](#github.com.Juniper.contrail.pkg.models.CreateWidgetResponse)
    - [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment)
    - [CustomerAttachmentFloatingIPRef](#github.com.Juniper.contrail.pkg.models.CustomerAttachmentFloatingIPRef)
    - [CustomerAttachmentVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.CustomerAttachmentVirtualMachineInterfaceRef)
    - [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard)
    - [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode)
    - [DeleteAPIAccessListRequest](#github.com.Juniper.contrail.pkg.models.DeleteAPIAccessListRequest)
    - [DeleteAPIAccessListResponse](#github.com.Juniper.contrail.pkg.models.DeleteAPIAccessListResponse)
    - [DeleteAccessControlListRequest](#github.com.Juniper.contrail.pkg.models.DeleteAccessControlListRequest)
    - [DeleteAccessControlListResponse](#github.com.Juniper.contrail.pkg.models.DeleteAccessControlListResponse)
    - [DeleteAddressGroupRequest](#github.com.Juniper.contrail.pkg.models.DeleteAddressGroupRequest)
    - [DeleteAddressGroupResponse](#github.com.Juniper.contrail.pkg.models.DeleteAddressGroupResponse)
    - [DeleteAlarmRequest](#github.com.Juniper.contrail.pkg.models.DeleteAlarmRequest)
    - [DeleteAlarmResponse](#github.com.Juniper.contrail.pkg.models.DeleteAlarmResponse)
    - [DeleteAliasIPPoolRequest](#github.com.Juniper.contrail.pkg.models.DeleteAliasIPPoolRequest)
    - [DeleteAliasIPPoolResponse](#github.com.Juniper.contrail.pkg.models.DeleteAliasIPPoolResponse)
    - [DeleteAliasIPRequest](#github.com.Juniper.contrail.pkg.models.DeleteAliasIPRequest)
    - [DeleteAliasIPResponse](#github.com.Juniper.contrail.pkg.models.DeleteAliasIPResponse)
    - [DeleteAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteAnalyticsNodeRequest)
    - [DeleteAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteAnalyticsNodeResponse)
    - [DeleteAppformixNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteAppformixNodeRequest)
    - [DeleteAppformixNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteAppformixNodeResponse)
    - [DeleteApplicationPolicySetRequest](#github.com.Juniper.contrail.pkg.models.DeleteApplicationPolicySetRequest)
    - [DeleteApplicationPolicySetResponse](#github.com.Juniper.contrail.pkg.models.DeleteApplicationPolicySetResponse)
    - [DeleteBGPAsAServiceRequest](#github.com.Juniper.contrail.pkg.models.DeleteBGPAsAServiceRequest)
    - [DeleteBGPAsAServiceResponse](#github.com.Juniper.contrail.pkg.models.DeleteBGPAsAServiceResponse)
    - [DeleteBGPRouterRequest](#github.com.Juniper.contrail.pkg.models.DeleteBGPRouterRequest)
    - [DeleteBGPRouterResponse](#github.com.Juniper.contrail.pkg.models.DeleteBGPRouterResponse)
    - [DeleteBGPVPNRequest](#github.com.Juniper.contrail.pkg.models.DeleteBGPVPNRequest)
    - [DeleteBGPVPNResponse](#github.com.Juniper.contrail.pkg.models.DeleteBGPVPNResponse)
    - [DeleteBaremetalNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteBaremetalNodeRequest)
    - [DeleteBaremetalNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteBaremetalNodeResponse)
    - [DeleteBaremetalPortRequest](#github.com.Juniper.contrail.pkg.models.DeleteBaremetalPortRequest)
    - [DeleteBaremetalPortResponse](#github.com.Juniper.contrail.pkg.models.DeleteBaremetalPortResponse)
    - [DeleteBridgeDomainRequest](#github.com.Juniper.contrail.pkg.models.DeleteBridgeDomainRequest)
    - [DeleteBridgeDomainResponse](#github.com.Juniper.contrail.pkg.models.DeleteBridgeDomainResponse)
    - [DeleteConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteConfigNodeRequest)
    - [DeleteConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteConfigNodeResponse)
    - [DeleteConfigRootRequest](#github.com.Juniper.contrail.pkg.models.DeleteConfigRootRequest)
    - [DeleteConfigRootResponse](#github.com.Juniper.contrail.pkg.models.DeleteConfigRootResponse)
    - [DeleteContrailAnalyticsDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsDatabaseNodeRequest)
    - [DeleteContrailAnalyticsDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsDatabaseNodeResponse)
    - [DeleteContrailAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsNodeRequest)
    - [DeleteContrailAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsNodeResponse)
    - [DeleteContrailClusterRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailClusterRequest)
    - [DeleteContrailClusterResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailClusterResponse)
    - [DeleteContrailConfigDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailConfigDatabaseNodeRequest)
    - [DeleteContrailConfigDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailConfigDatabaseNodeResponse)
    - [DeleteContrailConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailConfigNodeRequest)
    - [DeleteContrailConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailConfigNodeResponse)
    - [DeleteContrailControlNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailControlNodeRequest)
    - [DeleteContrailControlNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailControlNodeResponse)
    - [DeleteContrailStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailStorageNodeRequest)
    - [DeleteContrailStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailStorageNodeResponse)
    - [DeleteContrailVrouterNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailVrouterNodeRequest)
    - [DeleteContrailVrouterNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailVrouterNodeResponse)
    - [DeleteContrailWebuiNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteContrailWebuiNodeRequest)
    - [DeleteContrailWebuiNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteContrailWebuiNodeResponse)
    - [DeleteCustomerAttachmentRequest](#github.com.Juniper.contrail.pkg.models.DeleteCustomerAttachmentRequest)
    - [DeleteCustomerAttachmentResponse](#github.com.Juniper.contrail.pkg.models.DeleteCustomerAttachmentResponse)
    - [DeleteDashboardRequest](#github.com.Juniper.contrail.pkg.models.DeleteDashboardRequest)
    - [DeleteDashboardResponse](#github.com.Juniper.contrail.pkg.models.DeleteDashboardResponse)
    - [DeleteDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteDatabaseNodeRequest)
    - [DeleteDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteDatabaseNodeResponse)
    - [DeleteDiscoveryServiceAssignmentRequest](#github.com.Juniper.contrail.pkg.models.DeleteDiscoveryServiceAssignmentRequest)
    - [DeleteDiscoveryServiceAssignmentResponse](#github.com.Juniper.contrail.pkg.models.DeleteDiscoveryServiceAssignmentResponse)
    - [DeleteDomainRequest](#github.com.Juniper.contrail.pkg.models.DeleteDomainRequest)
    - [DeleteDomainResponse](#github.com.Juniper.contrail.pkg.models.DeleteDomainResponse)
    - [DeleteDsaRuleRequest](#github.com.Juniper.contrail.pkg.models.DeleteDsaRuleRequest)
    - [DeleteDsaRuleResponse](#github.com.Juniper.contrail.pkg.models.DeleteDsaRuleResponse)
    - [DeleteE2ServiceProviderRequest](#github.com.Juniper.contrail.pkg.models.DeleteE2ServiceProviderRequest)
    - [DeleteE2ServiceProviderResponse](#github.com.Juniper.contrail.pkg.models.DeleteE2ServiceProviderResponse)
    - [DeleteFirewallPolicyRequest](#github.com.Juniper.contrail.pkg.models.DeleteFirewallPolicyRequest)
    - [DeleteFirewallPolicyResponse](#github.com.Juniper.contrail.pkg.models.DeleteFirewallPolicyResponse)
    - [DeleteFirewallRuleRequest](#github.com.Juniper.contrail.pkg.models.DeleteFirewallRuleRequest)
    - [DeleteFirewallRuleResponse](#github.com.Juniper.contrail.pkg.models.DeleteFirewallRuleResponse)
    - [DeleteFlavorRequest](#github.com.Juniper.contrail.pkg.models.DeleteFlavorRequest)
    - [DeleteFlavorResponse](#github.com.Juniper.contrail.pkg.models.DeleteFlavorResponse)
    - [DeleteFloatingIPPoolRequest](#github.com.Juniper.contrail.pkg.models.DeleteFloatingIPPoolRequest)
    - [DeleteFloatingIPPoolResponse](#github.com.Juniper.contrail.pkg.models.DeleteFloatingIPPoolResponse)
    - [DeleteFloatingIPRequest](#github.com.Juniper.contrail.pkg.models.DeleteFloatingIPRequest)
    - [DeleteFloatingIPResponse](#github.com.Juniper.contrail.pkg.models.DeleteFloatingIPResponse)
    - [DeleteForwardingClassRequest](#github.com.Juniper.contrail.pkg.models.DeleteForwardingClassRequest)
    - [DeleteForwardingClassResponse](#github.com.Juniper.contrail.pkg.models.DeleteForwardingClassResponse)
    - [DeleteGlobalQosConfigRequest](#github.com.Juniper.contrail.pkg.models.DeleteGlobalQosConfigRequest)
    - [DeleteGlobalQosConfigResponse](#github.com.Juniper.contrail.pkg.models.DeleteGlobalQosConfigResponse)
    - [DeleteGlobalSystemConfigRequest](#github.com.Juniper.contrail.pkg.models.DeleteGlobalSystemConfigRequest)
    - [DeleteGlobalSystemConfigResponse](#github.com.Juniper.contrail.pkg.models.DeleteGlobalSystemConfigResponse)
    - [DeleteGlobalVrouterConfigRequest](#github.com.Juniper.contrail.pkg.models.DeleteGlobalVrouterConfigRequest)
    - [DeleteGlobalVrouterConfigResponse](#github.com.Juniper.contrail.pkg.models.DeleteGlobalVrouterConfigResponse)
    - [DeleteInstanceIPRequest](#github.com.Juniper.contrail.pkg.models.DeleteInstanceIPRequest)
    - [DeleteInstanceIPResponse](#github.com.Juniper.contrail.pkg.models.DeleteInstanceIPResponse)
    - [DeleteInterfaceRouteTableRequest](#github.com.Juniper.contrail.pkg.models.DeleteInterfaceRouteTableRequest)
    - [DeleteInterfaceRouteTableResponse](#github.com.Juniper.contrail.pkg.models.DeleteInterfaceRouteTableResponse)
    - [DeleteKeypairRequest](#github.com.Juniper.contrail.pkg.models.DeleteKeypairRequest)
    - [DeleteKeypairResponse](#github.com.Juniper.contrail.pkg.models.DeleteKeypairResponse)
    - [DeleteKubernetesMasterNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteKubernetesMasterNodeRequest)
    - [DeleteKubernetesMasterNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteKubernetesMasterNodeResponse)
    - [DeleteKubernetesNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteKubernetesNodeRequest)
    - [DeleteKubernetesNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteKubernetesNodeResponse)
    - [DeleteLoadbalancerHealthmonitorRequest](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerHealthmonitorRequest)
    - [DeleteLoadbalancerHealthmonitorResponse](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerHealthmonitorResponse)
    - [DeleteLoadbalancerListenerRequest](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerListenerRequest)
    - [DeleteLoadbalancerListenerResponse](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerListenerResponse)
    - [DeleteLoadbalancerMemberRequest](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerMemberRequest)
    - [DeleteLoadbalancerMemberResponse](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerMemberResponse)
    - [DeleteLoadbalancerPoolRequest](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerPoolRequest)
    - [DeleteLoadbalancerPoolResponse](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerPoolResponse)
    - [DeleteLoadbalancerRequest](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerRequest)
    - [DeleteLoadbalancerResponse](#github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerResponse)
    - [DeleteLocationRequest](#github.com.Juniper.contrail.pkg.models.DeleteLocationRequest)
    - [DeleteLocationResponse](#github.com.Juniper.contrail.pkg.models.DeleteLocationResponse)
    - [DeleteLogicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.DeleteLogicalInterfaceRequest)
    - [DeleteLogicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.DeleteLogicalInterfaceResponse)
    - [DeleteLogicalRouterRequest](#github.com.Juniper.contrail.pkg.models.DeleteLogicalRouterRequest)
    - [DeleteLogicalRouterResponse](#github.com.Juniper.contrail.pkg.models.DeleteLogicalRouterResponse)
    - [DeleteNamespaceRequest](#github.com.Juniper.contrail.pkg.models.DeleteNamespaceRequest)
    - [DeleteNamespaceResponse](#github.com.Juniper.contrail.pkg.models.DeleteNamespaceResponse)
    - [DeleteNetworkDeviceConfigRequest](#github.com.Juniper.contrail.pkg.models.DeleteNetworkDeviceConfigRequest)
    - [DeleteNetworkDeviceConfigResponse](#github.com.Juniper.contrail.pkg.models.DeleteNetworkDeviceConfigResponse)
    - [DeleteNetworkIpamRequest](#github.com.Juniper.contrail.pkg.models.DeleteNetworkIpamRequest)
    - [DeleteNetworkIpamResponse](#github.com.Juniper.contrail.pkg.models.DeleteNetworkIpamResponse)
    - [DeleteNetworkPolicyRequest](#github.com.Juniper.contrail.pkg.models.DeleteNetworkPolicyRequest)
    - [DeleteNetworkPolicyResponse](#github.com.Juniper.contrail.pkg.models.DeleteNetworkPolicyResponse)
    - [DeleteNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteNodeRequest)
    - [DeleteNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteNodeResponse)
    - [DeleteOpenstackComputeNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackComputeNodeRequest)
    - [DeleteOpenstackComputeNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackComputeNodeResponse)
    - [DeleteOpenstackControlNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackControlNodeRequest)
    - [DeleteOpenstackControlNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackControlNodeResponse)
    - [DeleteOpenstackMonitoringNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackMonitoringNodeRequest)
    - [DeleteOpenstackMonitoringNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackMonitoringNodeResponse)
    - [DeleteOpenstackNetworkNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackNetworkNodeRequest)
    - [DeleteOpenstackNetworkNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackNetworkNodeResponse)
    - [DeleteOpenstackStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackStorageNodeRequest)
    - [DeleteOpenstackStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.DeleteOpenstackStorageNodeResponse)
    - [DeleteOsImageRequest](#github.com.Juniper.contrail.pkg.models.DeleteOsImageRequest)
    - [DeleteOsImageResponse](#github.com.Juniper.contrail.pkg.models.DeleteOsImageResponse)
    - [DeletePeeringPolicyRequest](#github.com.Juniper.contrail.pkg.models.DeletePeeringPolicyRequest)
    - [DeletePeeringPolicyResponse](#github.com.Juniper.contrail.pkg.models.DeletePeeringPolicyResponse)
    - [DeletePhysicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.DeletePhysicalInterfaceRequest)
    - [DeletePhysicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.DeletePhysicalInterfaceResponse)
    - [DeletePhysicalRouterRequest](#github.com.Juniper.contrail.pkg.models.DeletePhysicalRouterRequest)
    - [DeletePhysicalRouterResponse](#github.com.Juniper.contrail.pkg.models.DeletePhysicalRouterResponse)
    - [DeletePolicyManagementRequest](#github.com.Juniper.contrail.pkg.models.DeletePolicyManagementRequest)
    - [DeletePolicyManagementResponse](#github.com.Juniper.contrail.pkg.models.DeletePolicyManagementResponse)
    - [DeletePortRequest](#github.com.Juniper.contrail.pkg.models.DeletePortRequest)
    - [DeletePortResponse](#github.com.Juniper.contrail.pkg.models.DeletePortResponse)
    - [DeletePortTupleRequest](#github.com.Juniper.contrail.pkg.models.DeletePortTupleRequest)
    - [DeletePortTupleResponse](#github.com.Juniper.contrail.pkg.models.DeletePortTupleResponse)
    - [DeleteProjectRequest](#github.com.Juniper.contrail.pkg.models.DeleteProjectRequest)
    - [DeleteProjectResponse](#github.com.Juniper.contrail.pkg.models.DeleteProjectResponse)
    - [DeleteProviderAttachmentRequest](#github.com.Juniper.contrail.pkg.models.DeleteProviderAttachmentRequest)
    - [DeleteProviderAttachmentResponse](#github.com.Juniper.contrail.pkg.models.DeleteProviderAttachmentResponse)
    - [DeleteQosConfigRequest](#github.com.Juniper.contrail.pkg.models.DeleteQosConfigRequest)
    - [DeleteQosConfigResponse](#github.com.Juniper.contrail.pkg.models.DeleteQosConfigResponse)
    - [DeleteQosQueueRequest](#github.com.Juniper.contrail.pkg.models.DeleteQosQueueRequest)
    - [DeleteQosQueueResponse](#github.com.Juniper.contrail.pkg.models.DeleteQosQueueResponse)
    - [DeleteRouteAggregateRequest](#github.com.Juniper.contrail.pkg.models.DeleteRouteAggregateRequest)
    - [DeleteRouteAggregateResponse](#github.com.Juniper.contrail.pkg.models.DeleteRouteAggregateResponse)
    - [DeleteRouteTableRequest](#github.com.Juniper.contrail.pkg.models.DeleteRouteTableRequest)
    - [DeleteRouteTableResponse](#github.com.Juniper.contrail.pkg.models.DeleteRouteTableResponse)
    - [DeleteRouteTargetRequest](#github.com.Juniper.contrail.pkg.models.DeleteRouteTargetRequest)
    - [DeleteRouteTargetResponse](#github.com.Juniper.contrail.pkg.models.DeleteRouteTargetResponse)
    - [DeleteRoutingInstanceRequest](#github.com.Juniper.contrail.pkg.models.DeleteRoutingInstanceRequest)
    - [DeleteRoutingInstanceResponse](#github.com.Juniper.contrail.pkg.models.DeleteRoutingInstanceResponse)
    - [DeleteRoutingPolicyRequest](#github.com.Juniper.contrail.pkg.models.DeleteRoutingPolicyRequest)
    - [DeleteRoutingPolicyResponse](#github.com.Juniper.contrail.pkg.models.DeleteRoutingPolicyResponse)
    - [DeleteSecurityGroupRequest](#github.com.Juniper.contrail.pkg.models.DeleteSecurityGroupRequest)
    - [DeleteSecurityGroupResponse](#github.com.Juniper.contrail.pkg.models.DeleteSecurityGroupResponse)
    - [DeleteSecurityLoggingObjectRequest](#github.com.Juniper.contrail.pkg.models.DeleteSecurityLoggingObjectRequest)
    - [DeleteSecurityLoggingObjectResponse](#github.com.Juniper.contrail.pkg.models.DeleteSecurityLoggingObjectResponse)
    - [DeleteServerRequest](#github.com.Juniper.contrail.pkg.models.DeleteServerRequest)
    - [DeleteServerResponse](#github.com.Juniper.contrail.pkg.models.DeleteServerResponse)
    - [DeleteServiceApplianceRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceRequest)
    - [DeleteServiceApplianceResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceResponse)
    - [DeleteServiceApplianceSetRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceSetRequest)
    - [DeleteServiceApplianceSetResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceSetResponse)
    - [DeleteServiceConnectionModuleRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceConnectionModuleRequest)
    - [DeleteServiceConnectionModuleResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceConnectionModuleResponse)
    - [DeleteServiceEndpointRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceEndpointRequest)
    - [DeleteServiceEndpointResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceEndpointResponse)
    - [DeleteServiceGroupRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceGroupRequest)
    - [DeleteServiceGroupResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceGroupResponse)
    - [DeleteServiceHealthCheckRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceHealthCheckRequest)
    - [DeleteServiceHealthCheckResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceHealthCheckResponse)
    - [DeleteServiceInstanceRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceInstanceRequest)
    - [DeleteServiceInstanceResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceInstanceResponse)
    - [DeleteServiceObjectRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceObjectRequest)
    - [DeleteServiceObjectResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceObjectResponse)
    - [DeleteServiceTemplateRequest](#github.com.Juniper.contrail.pkg.models.DeleteServiceTemplateRequest)
    - [DeleteServiceTemplateResponse](#github.com.Juniper.contrail.pkg.models.DeleteServiceTemplateResponse)
    - [DeleteSubnetRequest](#github.com.Juniper.contrail.pkg.models.DeleteSubnetRequest)
    - [DeleteSubnetResponse](#github.com.Juniper.contrail.pkg.models.DeleteSubnetResponse)
    - [DeleteTagRequest](#github.com.Juniper.contrail.pkg.models.DeleteTagRequest)
    - [DeleteTagResponse](#github.com.Juniper.contrail.pkg.models.DeleteTagResponse)
    - [DeleteTagTypeRequest](#github.com.Juniper.contrail.pkg.models.DeleteTagTypeRequest)
    - [DeleteTagTypeResponse](#github.com.Juniper.contrail.pkg.models.DeleteTagTypeResponse)
    - [DeleteUserRequest](#github.com.Juniper.contrail.pkg.models.DeleteUserRequest)
    - [DeleteUserResponse](#github.com.Juniper.contrail.pkg.models.DeleteUserResponse)
    - [DeleteVPNGroupRequest](#github.com.Juniper.contrail.pkg.models.DeleteVPNGroupRequest)
    - [DeleteVPNGroupResponse](#github.com.Juniper.contrail.pkg.models.DeleteVPNGroupResponse)
    - [DeleteVirtualDNSRecordRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSRecordRequest)
    - [DeleteVirtualDNSRecordResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSRecordResponse)
    - [DeleteVirtualDNSRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSRequest)
    - [DeleteVirtualDNSResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSResponse)
    - [DeleteVirtualIPRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualIPRequest)
    - [DeleteVirtualIPResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualIPResponse)
    - [DeleteVirtualMachineInterfaceRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineInterfaceRequest)
    - [DeleteVirtualMachineInterfaceResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineInterfaceResponse)
    - [DeleteVirtualMachineRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineRequest)
    - [DeleteVirtualMachineResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineResponse)
    - [DeleteVirtualNetworkRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualNetworkRequest)
    - [DeleteVirtualNetworkResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualNetworkResponse)
    - [DeleteVirtualRouterRequest](#github.com.Juniper.contrail.pkg.models.DeleteVirtualRouterRequest)
    - [DeleteVirtualRouterResponse](#github.com.Juniper.contrail.pkg.models.DeleteVirtualRouterResponse)
    - [DeleteWidgetRequest](#github.com.Juniper.contrail.pkg.models.DeleteWidgetRequest)
    - [DeleteWidgetResponse](#github.com.Juniper.contrail.pkg.models.DeleteWidgetResponse)
    - [DhcpOptionType](#github.com.Juniper.contrail.pkg.models.DhcpOptionType)
    - [DhcpOptionsListType](#github.com.Juniper.contrail.pkg.models.DhcpOptionsListType)
    - [DiscoveryPubSubEndPointType](#github.com.Juniper.contrail.pkg.models.DiscoveryPubSubEndPointType)
    - [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment)
    - [DiscoveryServiceAssignmentType](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignmentType)
    - [Domain](#github.com.Juniper.contrail.pkg.models.Domain)
    - [DomainLimitsType](#github.com.Juniper.contrail.pkg.models.DomainLimitsType)
    - [DriverInfo](#github.com.Juniper.contrail.pkg.models.DriverInfo)
    - [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule)
    - [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider)
    - [E2ServiceProviderPeeringPolicyRef](#github.com.Juniper.contrail.pkg.models.E2ServiceProviderPeeringPolicyRef)
    - [E2ServiceProviderPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.E2ServiceProviderPhysicalRouterRef)
    - [EcmpHashingIncludeFields](#github.com.Juniper.contrail.pkg.models.EcmpHashingIncludeFields)
    - [EncapsulationPrioritiesType](#github.com.Juniper.contrail.pkg.models.EncapsulationPrioritiesType)
    - [FatFlowProtocols](#github.com.Juniper.contrail.pkg.models.FatFlowProtocols)
    - [Filter](#github.com.Juniper.contrail.pkg.models.Filter)
    - [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy)
    - [FirewallPolicyFirewallRuleRef](#github.com.Juniper.contrail.pkg.models.FirewallPolicyFirewallRuleRef)
    - [FirewallPolicySecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.FirewallPolicySecurityLoggingObjectRef)
    - [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule)
    - [FirewallRuleAddressGroupRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleAddressGroupRef)
    - [FirewallRuleEndpointType](#github.com.Juniper.contrail.pkg.models.FirewallRuleEndpointType)
    - [FirewallRuleMatchTagsType](#github.com.Juniper.contrail.pkg.models.FirewallRuleMatchTagsType)
    - [FirewallRuleMatchTagsTypeIdList](#github.com.Juniper.contrail.pkg.models.FirewallRuleMatchTagsTypeIdList)
    - [FirewallRuleSecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleSecurityLoggingObjectRef)
    - [FirewallRuleServiceGroupRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleServiceGroupRef)
    - [FirewallRuleVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleVirtualNetworkRef)
    - [FirewallSequence](#github.com.Juniper.contrail.pkg.models.FirewallSequence)
    - [FirewallServiceGroupType](#github.com.Juniper.contrail.pkg.models.FirewallServiceGroupType)
    - [FirewallServiceType](#github.com.Juniper.contrail.pkg.models.FirewallServiceType)
    - [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor)
    - [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP)
    - [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool)
    - [FloatingIPProjectRef](#github.com.Juniper.contrail.pkg.models.FloatingIPProjectRef)
    - [FloatingIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.FloatingIPVirtualMachineInterfaceRef)
    - [FloatingIpPoolSubnetType](#github.com.Juniper.contrail.pkg.models.FloatingIpPoolSubnetType)
    - [FlowAgingTimeout](#github.com.Juniper.contrail.pkg.models.FlowAgingTimeout)
    - [FlowAgingTimeoutList](#github.com.Juniper.contrail.pkg.models.FlowAgingTimeoutList)
    - [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass)
    - [ForwardingClassQosQueueRef](#github.com.Juniper.contrail.pkg.models.ForwardingClassQosQueueRef)
    - [GetAPIAccessListRequest](#github.com.Juniper.contrail.pkg.models.GetAPIAccessListRequest)
    - [GetAPIAccessListResponse](#github.com.Juniper.contrail.pkg.models.GetAPIAccessListResponse)
    - [GetAccessControlListRequest](#github.com.Juniper.contrail.pkg.models.GetAccessControlListRequest)
    - [GetAccessControlListResponse](#github.com.Juniper.contrail.pkg.models.GetAccessControlListResponse)
    - [GetAddressGroupRequest](#github.com.Juniper.contrail.pkg.models.GetAddressGroupRequest)
    - [GetAddressGroupResponse](#github.com.Juniper.contrail.pkg.models.GetAddressGroupResponse)
    - [GetAlarmRequest](#github.com.Juniper.contrail.pkg.models.GetAlarmRequest)
    - [GetAlarmResponse](#github.com.Juniper.contrail.pkg.models.GetAlarmResponse)
    - [GetAliasIPPoolRequest](#github.com.Juniper.contrail.pkg.models.GetAliasIPPoolRequest)
    - [GetAliasIPPoolResponse](#github.com.Juniper.contrail.pkg.models.GetAliasIPPoolResponse)
    - [GetAliasIPRequest](#github.com.Juniper.contrail.pkg.models.GetAliasIPRequest)
    - [GetAliasIPResponse](#github.com.Juniper.contrail.pkg.models.GetAliasIPResponse)
    - [GetAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.GetAnalyticsNodeRequest)
    - [GetAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.GetAnalyticsNodeResponse)
    - [GetAppformixNodeRequest](#github.com.Juniper.contrail.pkg.models.GetAppformixNodeRequest)
    - [GetAppformixNodeResponse](#github.com.Juniper.contrail.pkg.models.GetAppformixNodeResponse)
    - [GetApplicationPolicySetRequest](#github.com.Juniper.contrail.pkg.models.GetApplicationPolicySetRequest)
    - [GetApplicationPolicySetResponse](#github.com.Juniper.contrail.pkg.models.GetApplicationPolicySetResponse)
    - [GetBGPAsAServiceRequest](#github.com.Juniper.contrail.pkg.models.GetBGPAsAServiceRequest)
    - [GetBGPAsAServiceResponse](#github.com.Juniper.contrail.pkg.models.GetBGPAsAServiceResponse)
    - [GetBGPRouterRequest](#github.com.Juniper.contrail.pkg.models.GetBGPRouterRequest)
    - [GetBGPRouterResponse](#github.com.Juniper.contrail.pkg.models.GetBGPRouterResponse)
    - [GetBGPVPNRequest](#github.com.Juniper.contrail.pkg.models.GetBGPVPNRequest)
    - [GetBGPVPNResponse](#github.com.Juniper.contrail.pkg.models.GetBGPVPNResponse)
    - [GetBaremetalNodeRequest](#github.com.Juniper.contrail.pkg.models.GetBaremetalNodeRequest)
    - [GetBaremetalNodeResponse](#github.com.Juniper.contrail.pkg.models.GetBaremetalNodeResponse)
    - [GetBaremetalPortRequest](#github.com.Juniper.contrail.pkg.models.GetBaremetalPortRequest)
    - [GetBaremetalPortResponse](#github.com.Juniper.contrail.pkg.models.GetBaremetalPortResponse)
    - [GetBridgeDomainRequest](#github.com.Juniper.contrail.pkg.models.GetBridgeDomainRequest)
    - [GetBridgeDomainResponse](#github.com.Juniper.contrail.pkg.models.GetBridgeDomainResponse)
    - [GetConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.GetConfigNodeRequest)
    - [GetConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.GetConfigNodeResponse)
    - [GetConfigRootRequest](#github.com.Juniper.contrail.pkg.models.GetConfigRootRequest)
    - [GetConfigRootResponse](#github.com.Juniper.contrail.pkg.models.GetConfigRootResponse)
    - [GetContrailAnalyticsDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsDatabaseNodeRequest)
    - [GetContrailAnalyticsDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsDatabaseNodeResponse)
    - [GetContrailAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsNodeRequest)
    - [GetContrailAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsNodeResponse)
    - [GetContrailClusterRequest](#github.com.Juniper.contrail.pkg.models.GetContrailClusterRequest)
    - [GetContrailClusterResponse](#github.com.Juniper.contrail.pkg.models.GetContrailClusterResponse)
    - [GetContrailConfigDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailConfigDatabaseNodeRequest)
    - [GetContrailConfigDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailConfigDatabaseNodeResponse)
    - [GetContrailConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailConfigNodeRequest)
    - [GetContrailConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailConfigNodeResponse)
    - [GetContrailControlNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailControlNodeRequest)
    - [GetContrailControlNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailControlNodeResponse)
    - [GetContrailStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailStorageNodeRequest)
    - [GetContrailStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailStorageNodeResponse)
    - [GetContrailVrouterNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailVrouterNodeRequest)
    - [GetContrailVrouterNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailVrouterNodeResponse)
    - [GetContrailWebuiNodeRequest](#github.com.Juniper.contrail.pkg.models.GetContrailWebuiNodeRequest)
    - [GetContrailWebuiNodeResponse](#github.com.Juniper.contrail.pkg.models.GetContrailWebuiNodeResponse)
    - [GetCustomerAttachmentRequest](#github.com.Juniper.contrail.pkg.models.GetCustomerAttachmentRequest)
    - [GetCustomerAttachmentResponse](#github.com.Juniper.contrail.pkg.models.GetCustomerAttachmentResponse)
    - [GetDashboardRequest](#github.com.Juniper.contrail.pkg.models.GetDashboardRequest)
    - [GetDashboardResponse](#github.com.Juniper.contrail.pkg.models.GetDashboardResponse)
    - [GetDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.GetDatabaseNodeRequest)
    - [GetDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.GetDatabaseNodeResponse)
    - [GetDiscoveryServiceAssignmentRequest](#github.com.Juniper.contrail.pkg.models.GetDiscoveryServiceAssignmentRequest)
    - [GetDiscoveryServiceAssignmentResponse](#github.com.Juniper.contrail.pkg.models.GetDiscoveryServiceAssignmentResponse)
    - [GetDomainRequest](#github.com.Juniper.contrail.pkg.models.GetDomainRequest)
    - [GetDomainResponse](#github.com.Juniper.contrail.pkg.models.GetDomainResponse)
    - [GetDsaRuleRequest](#github.com.Juniper.contrail.pkg.models.GetDsaRuleRequest)
    - [GetDsaRuleResponse](#github.com.Juniper.contrail.pkg.models.GetDsaRuleResponse)
    - [GetE2ServiceProviderRequest](#github.com.Juniper.contrail.pkg.models.GetE2ServiceProviderRequest)
    - [GetE2ServiceProviderResponse](#github.com.Juniper.contrail.pkg.models.GetE2ServiceProviderResponse)
    - [GetFirewallPolicyRequest](#github.com.Juniper.contrail.pkg.models.GetFirewallPolicyRequest)
    - [GetFirewallPolicyResponse](#github.com.Juniper.contrail.pkg.models.GetFirewallPolicyResponse)
    - [GetFirewallRuleRequest](#github.com.Juniper.contrail.pkg.models.GetFirewallRuleRequest)
    - [GetFirewallRuleResponse](#github.com.Juniper.contrail.pkg.models.GetFirewallRuleResponse)
    - [GetFlavorRequest](#github.com.Juniper.contrail.pkg.models.GetFlavorRequest)
    - [GetFlavorResponse](#github.com.Juniper.contrail.pkg.models.GetFlavorResponse)
    - [GetFloatingIPPoolRequest](#github.com.Juniper.contrail.pkg.models.GetFloatingIPPoolRequest)
    - [GetFloatingIPPoolResponse](#github.com.Juniper.contrail.pkg.models.GetFloatingIPPoolResponse)
    - [GetFloatingIPRequest](#github.com.Juniper.contrail.pkg.models.GetFloatingIPRequest)
    - [GetFloatingIPResponse](#github.com.Juniper.contrail.pkg.models.GetFloatingIPResponse)
    - [GetForwardingClassRequest](#github.com.Juniper.contrail.pkg.models.GetForwardingClassRequest)
    - [GetForwardingClassResponse](#github.com.Juniper.contrail.pkg.models.GetForwardingClassResponse)
    - [GetGlobalQosConfigRequest](#github.com.Juniper.contrail.pkg.models.GetGlobalQosConfigRequest)
    - [GetGlobalQosConfigResponse](#github.com.Juniper.contrail.pkg.models.GetGlobalQosConfigResponse)
    - [GetGlobalSystemConfigRequest](#github.com.Juniper.contrail.pkg.models.GetGlobalSystemConfigRequest)
    - [GetGlobalSystemConfigResponse](#github.com.Juniper.contrail.pkg.models.GetGlobalSystemConfigResponse)
    - [GetGlobalVrouterConfigRequest](#github.com.Juniper.contrail.pkg.models.GetGlobalVrouterConfigRequest)
    - [GetGlobalVrouterConfigResponse](#github.com.Juniper.contrail.pkg.models.GetGlobalVrouterConfigResponse)
    - [GetInstanceIPRequest](#github.com.Juniper.contrail.pkg.models.GetInstanceIPRequest)
    - [GetInstanceIPResponse](#github.com.Juniper.contrail.pkg.models.GetInstanceIPResponse)
    - [GetInterfaceRouteTableRequest](#github.com.Juniper.contrail.pkg.models.GetInterfaceRouteTableRequest)
    - [GetInterfaceRouteTableResponse](#github.com.Juniper.contrail.pkg.models.GetInterfaceRouteTableResponse)
    - [GetKeypairRequest](#github.com.Juniper.contrail.pkg.models.GetKeypairRequest)
    - [GetKeypairResponse](#github.com.Juniper.contrail.pkg.models.GetKeypairResponse)
    - [GetKubernetesMasterNodeRequest](#github.com.Juniper.contrail.pkg.models.GetKubernetesMasterNodeRequest)
    - [GetKubernetesMasterNodeResponse](#github.com.Juniper.contrail.pkg.models.GetKubernetesMasterNodeResponse)
    - [GetKubernetesNodeRequest](#github.com.Juniper.contrail.pkg.models.GetKubernetesNodeRequest)
    - [GetKubernetesNodeResponse](#github.com.Juniper.contrail.pkg.models.GetKubernetesNodeResponse)
    - [GetLoadbalancerHealthmonitorRequest](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerHealthmonitorRequest)
    - [GetLoadbalancerHealthmonitorResponse](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerHealthmonitorResponse)
    - [GetLoadbalancerListenerRequest](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerListenerRequest)
    - [GetLoadbalancerListenerResponse](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerListenerResponse)
    - [GetLoadbalancerMemberRequest](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerMemberRequest)
    - [GetLoadbalancerMemberResponse](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerMemberResponse)
    - [GetLoadbalancerPoolRequest](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerPoolRequest)
    - [GetLoadbalancerPoolResponse](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerPoolResponse)
    - [GetLoadbalancerRequest](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerRequest)
    - [GetLoadbalancerResponse](#github.com.Juniper.contrail.pkg.models.GetLoadbalancerResponse)
    - [GetLocationRequest](#github.com.Juniper.contrail.pkg.models.GetLocationRequest)
    - [GetLocationResponse](#github.com.Juniper.contrail.pkg.models.GetLocationResponse)
    - [GetLogicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.GetLogicalInterfaceRequest)
    - [GetLogicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.GetLogicalInterfaceResponse)
    - [GetLogicalRouterRequest](#github.com.Juniper.contrail.pkg.models.GetLogicalRouterRequest)
    - [GetLogicalRouterResponse](#github.com.Juniper.contrail.pkg.models.GetLogicalRouterResponse)
    - [GetNamespaceRequest](#github.com.Juniper.contrail.pkg.models.GetNamespaceRequest)
    - [GetNamespaceResponse](#github.com.Juniper.contrail.pkg.models.GetNamespaceResponse)
    - [GetNetworkDeviceConfigRequest](#github.com.Juniper.contrail.pkg.models.GetNetworkDeviceConfigRequest)
    - [GetNetworkDeviceConfigResponse](#github.com.Juniper.contrail.pkg.models.GetNetworkDeviceConfigResponse)
    - [GetNetworkIpamRequest](#github.com.Juniper.contrail.pkg.models.GetNetworkIpamRequest)
    - [GetNetworkIpamResponse](#github.com.Juniper.contrail.pkg.models.GetNetworkIpamResponse)
    - [GetNetworkPolicyRequest](#github.com.Juniper.contrail.pkg.models.GetNetworkPolicyRequest)
    - [GetNetworkPolicyResponse](#github.com.Juniper.contrail.pkg.models.GetNetworkPolicyResponse)
    - [GetNodeRequest](#github.com.Juniper.contrail.pkg.models.GetNodeRequest)
    - [GetNodeResponse](#github.com.Juniper.contrail.pkg.models.GetNodeResponse)
    - [GetOpenstackComputeNodeRequest](#github.com.Juniper.contrail.pkg.models.GetOpenstackComputeNodeRequest)
    - [GetOpenstackComputeNodeResponse](#github.com.Juniper.contrail.pkg.models.GetOpenstackComputeNodeResponse)
    - [GetOpenstackControlNodeRequest](#github.com.Juniper.contrail.pkg.models.GetOpenstackControlNodeRequest)
    - [GetOpenstackControlNodeResponse](#github.com.Juniper.contrail.pkg.models.GetOpenstackControlNodeResponse)
    - [GetOpenstackMonitoringNodeRequest](#github.com.Juniper.contrail.pkg.models.GetOpenstackMonitoringNodeRequest)
    - [GetOpenstackMonitoringNodeResponse](#github.com.Juniper.contrail.pkg.models.GetOpenstackMonitoringNodeResponse)
    - [GetOpenstackNetworkNodeRequest](#github.com.Juniper.contrail.pkg.models.GetOpenstackNetworkNodeRequest)
    - [GetOpenstackNetworkNodeResponse](#github.com.Juniper.contrail.pkg.models.GetOpenstackNetworkNodeResponse)
    - [GetOpenstackStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.GetOpenstackStorageNodeRequest)
    - [GetOpenstackStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.GetOpenstackStorageNodeResponse)
    - [GetOsImageRequest](#github.com.Juniper.contrail.pkg.models.GetOsImageRequest)
    - [GetOsImageResponse](#github.com.Juniper.contrail.pkg.models.GetOsImageResponse)
    - [GetPeeringPolicyRequest](#github.com.Juniper.contrail.pkg.models.GetPeeringPolicyRequest)
    - [GetPeeringPolicyResponse](#github.com.Juniper.contrail.pkg.models.GetPeeringPolicyResponse)
    - [GetPhysicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.GetPhysicalInterfaceRequest)
    - [GetPhysicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.GetPhysicalInterfaceResponse)
    - [GetPhysicalRouterRequest](#github.com.Juniper.contrail.pkg.models.GetPhysicalRouterRequest)
    - [GetPhysicalRouterResponse](#github.com.Juniper.contrail.pkg.models.GetPhysicalRouterResponse)
    - [GetPolicyManagementRequest](#github.com.Juniper.contrail.pkg.models.GetPolicyManagementRequest)
    - [GetPolicyManagementResponse](#github.com.Juniper.contrail.pkg.models.GetPolicyManagementResponse)
    - [GetPortRequest](#github.com.Juniper.contrail.pkg.models.GetPortRequest)
    - [GetPortResponse](#github.com.Juniper.contrail.pkg.models.GetPortResponse)
    - [GetPortTupleRequest](#github.com.Juniper.contrail.pkg.models.GetPortTupleRequest)
    - [GetPortTupleResponse](#github.com.Juniper.contrail.pkg.models.GetPortTupleResponse)
    - [GetProjectRequest](#github.com.Juniper.contrail.pkg.models.GetProjectRequest)
    - [GetProjectResponse](#github.com.Juniper.contrail.pkg.models.GetProjectResponse)
    - [GetProviderAttachmentRequest](#github.com.Juniper.contrail.pkg.models.GetProviderAttachmentRequest)
    - [GetProviderAttachmentResponse](#github.com.Juniper.contrail.pkg.models.GetProviderAttachmentResponse)
    - [GetQosConfigRequest](#github.com.Juniper.contrail.pkg.models.GetQosConfigRequest)
    - [GetQosConfigResponse](#github.com.Juniper.contrail.pkg.models.GetQosConfigResponse)
    - [GetQosQueueRequest](#github.com.Juniper.contrail.pkg.models.GetQosQueueRequest)
    - [GetQosQueueResponse](#github.com.Juniper.contrail.pkg.models.GetQosQueueResponse)
    - [GetRouteAggregateRequest](#github.com.Juniper.contrail.pkg.models.GetRouteAggregateRequest)
    - [GetRouteAggregateResponse](#github.com.Juniper.contrail.pkg.models.GetRouteAggregateResponse)
    - [GetRouteTableRequest](#github.com.Juniper.contrail.pkg.models.GetRouteTableRequest)
    - [GetRouteTableResponse](#github.com.Juniper.contrail.pkg.models.GetRouteTableResponse)
    - [GetRouteTargetRequest](#github.com.Juniper.contrail.pkg.models.GetRouteTargetRequest)
    - [GetRouteTargetResponse](#github.com.Juniper.contrail.pkg.models.GetRouteTargetResponse)
    - [GetRoutingInstanceRequest](#github.com.Juniper.contrail.pkg.models.GetRoutingInstanceRequest)
    - [GetRoutingInstanceResponse](#github.com.Juniper.contrail.pkg.models.GetRoutingInstanceResponse)
    - [GetRoutingPolicyRequest](#github.com.Juniper.contrail.pkg.models.GetRoutingPolicyRequest)
    - [GetRoutingPolicyResponse](#github.com.Juniper.contrail.pkg.models.GetRoutingPolicyResponse)
    - [GetSecurityGroupRequest](#github.com.Juniper.contrail.pkg.models.GetSecurityGroupRequest)
    - [GetSecurityGroupResponse](#github.com.Juniper.contrail.pkg.models.GetSecurityGroupResponse)
    - [GetSecurityLoggingObjectRequest](#github.com.Juniper.contrail.pkg.models.GetSecurityLoggingObjectRequest)
    - [GetSecurityLoggingObjectResponse](#github.com.Juniper.contrail.pkg.models.GetSecurityLoggingObjectResponse)
    - [GetServerRequest](#github.com.Juniper.contrail.pkg.models.GetServerRequest)
    - [GetServerResponse](#github.com.Juniper.contrail.pkg.models.GetServerResponse)
    - [GetServiceApplianceRequest](#github.com.Juniper.contrail.pkg.models.GetServiceApplianceRequest)
    - [GetServiceApplianceResponse](#github.com.Juniper.contrail.pkg.models.GetServiceApplianceResponse)
    - [GetServiceApplianceSetRequest](#github.com.Juniper.contrail.pkg.models.GetServiceApplianceSetRequest)
    - [GetServiceApplianceSetResponse](#github.com.Juniper.contrail.pkg.models.GetServiceApplianceSetResponse)
    - [GetServiceConnectionModuleRequest](#github.com.Juniper.contrail.pkg.models.GetServiceConnectionModuleRequest)
    - [GetServiceConnectionModuleResponse](#github.com.Juniper.contrail.pkg.models.GetServiceConnectionModuleResponse)
    - [GetServiceEndpointRequest](#github.com.Juniper.contrail.pkg.models.GetServiceEndpointRequest)
    - [GetServiceEndpointResponse](#github.com.Juniper.contrail.pkg.models.GetServiceEndpointResponse)
    - [GetServiceGroupRequest](#github.com.Juniper.contrail.pkg.models.GetServiceGroupRequest)
    - [GetServiceGroupResponse](#github.com.Juniper.contrail.pkg.models.GetServiceGroupResponse)
    - [GetServiceHealthCheckRequest](#github.com.Juniper.contrail.pkg.models.GetServiceHealthCheckRequest)
    - [GetServiceHealthCheckResponse](#github.com.Juniper.contrail.pkg.models.GetServiceHealthCheckResponse)
    - [GetServiceInstanceRequest](#github.com.Juniper.contrail.pkg.models.GetServiceInstanceRequest)
    - [GetServiceInstanceResponse](#github.com.Juniper.contrail.pkg.models.GetServiceInstanceResponse)
    - [GetServiceObjectRequest](#github.com.Juniper.contrail.pkg.models.GetServiceObjectRequest)
    - [GetServiceObjectResponse](#github.com.Juniper.contrail.pkg.models.GetServiceObjectResponse)
    - [GetServiceTemplateRequest](#github.com.Juniper.contrail.pkg.models.GetServiceTemplateRequest)
    - [GetServiceTemplateResponse](#github.com.Juniper.contrail.pkg.models.GetServiceTemplateResponse)
    - [GetSubnetRequest](#github.com.Juniper.contrail.pkg.models.GetSubnetRequest)
    - [GetSubnetResponse](#github.com.Juniper.contrail.pkg.models.GetSubnetResponse)
    - [GetTagRequest](#github.com.Juniper.contrail.pkg.models.GetTagRequest)
    - [GetTagResponse](#github.com.Juniper.contrail.pkg.models.GetTagResponse)
    - [GetTagTypeRequest](#github.com.Juniper.contrail.pkg.models.GetTagTypeRequest)
    - [GetTagTypeResponse](#github.com.Juniper.contrail.pkg.models.GetTagTypeResponse)
    - [GetUserRequest](#github.com.Juniper.contrail.pkg.models.GetUserRequest)
    - [GetUserResponse](#github.com.Juniper.contrail.pkg.models.GetUserResponse)
    - [GetVPNGroupRequest](#github.com.Juniper.contrail.pkg.models.GetVPNGroupRequest)
    - [GetVPNGroupResponse](#github.com.Juniper.contrail.pkg.models.GetVPNGroupResponse)
    - [GetVirtualDNSRecordRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualDNSRecordRequest)
    - [GetVirtualDNSRecordResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualDNSRecordResponse)
    - [GetVirtualDNSRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualDNSRequest)
    - [GetVirtualDNSResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualDNSResponse)
    - [GetVirtualIPRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualIPRequest)
    - [GetVirtualIPResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualIPResponse)
    - [GetVirtualMachineInterfaceRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualMachineInterfaceRequest)
    - [GetVirtualMachineInterfaceResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualMachineInterfaceResponse)
    - [GetVirtualMachineRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualMachineRequest)
    - [GetVirtualMachineResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualMachineResponse)
    - [GetVirtualNetworkRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualNetworkRequest)
    - [GetVirtualNetworkResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualNetworkResponse)
    - [GetVirtualRouterRequest](#github.com.Juniper.contrail.pkg.models.GetVirtualRouterRequest)
    - [GetVirtualRouterResponse](#github.com.Juniper.contrail.pkg.models.GetVirtualRouterResponse)
    - [GetWidgetRequest](#github.com.Juniper.contrail.pkg.models.GetWidgetRequest)
    - [GetWidgetResponse](#github.com.Juniper.contrail.pkg.models.GetWidgetResponse)
    - [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig)
    - [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig)
    - [GlobalSystemConfigBGPRouterRef](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfigBGPRouterRef)
    - [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig)
    - [GracefulRestartParametersType](#github.com.Juniper.contrail.pkg.models.GracefulRestartParametersType)
    - [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType)
    - [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP)
    - [InstanceIPNetworkIpamRef](#github.com.Juniper.contrail.pkg.models.InstanceIPNetworkIpamRef)
    - [InstanceIPPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.InstanceIPPhysicalRouterRef)
    - [InstanceIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.InstanceIPVirtualMachineInterfaceRef)
    - [InstanceIPVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.InstanceIPVirtualNetworkRef)
    - [InstanceIPVirtualRouterRef](#github.com.Juniper.contrail.pkg.models.InstanceIPVirtualRouterRef)
    - [InstanceInfo](#github.com.Juniper.contrail.pkg.models.InstanceInfo)
    - [InterfaceMirrorType](#github.com.Juniper.contrail.pkg.models.InterfaceMirrorType)
    - [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable)
    - [InterfaceRouteTableServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTableServiceInstanceRef)
    - [IpAddressesType](#github.com.Juniper.contrail.pkg.models.IpAddressesType)
    - [IpamDnsAddressType](#github.com.Juniper.contrail.pkg.models.IpamDnsAddressType)
    - [IpamSubnetType](#github.com.Juniper.contrail.pkg.models.IpamSubnetType)
    - [IpamSubnets](#github.com.Juniper.contrail.pkg.models.IpamSubnets)
    - [IpamType](#github.com.Juniper.contrail.pkg.models.IpamType)
    - [JunosServicePorts](#github.com.Juniper.contrail.pkg.models.JunosServicePorts)
    - [KeyValuePair](#github.com.Juniper.contrail.pkg.models.KeyValuePair)
    - [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs)
    - [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair)
    - [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode)
    - [KubernetesMasterNodeNodeRef](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNodeNodeRef)
    - [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode)
    - [KubernetesNodeNodeRef](#github.com.Juniper.contrail.pkg.models.KubernetesNodeNodeRef)
    - [LinklocalServiceEntryType](#github.com.Juniper.contrail.pkg.models.LinklocalServiceEntryType)
    - [LinklocalServicesTypes](#github.com.Juniper.contrail.pkg.models.LinklocalServicesTypes)
    - [ListAPIAccessListRequest](#github.com.Juniper.contrail.pkg.models.ListAPIAccessListRequest)
    - [ListAPIAccessListResponse](#github.com.Juniper.contrail.pkg.models.ListAPIAccessListResponse)
    - [ListAccessControlListRequest](#github.com.Juniper.contrail.pkg.models.ListAccessControlListRequest)
    - [ListAccessControlListResponse](#github.com.Juniper.contrail.pkg.models.ListAccessControlListResponse)
    - [ListAddressGroupRequest](#github.com.Juniper.contrail.pkg.models.ListAddressGroupRequest)
    - [ListAddressGroupResponse](#github.com.Juniper.contrail.pkg.models.ListAddressGroupResponse)
    - [ListAlarmRequest](#github.com.Juniper.contrail.pkg.models.ListAlarmRequest)
    - [ListAlarmResponse](#github.com.Juniper.contrail.pkg.models.ListAlarmResponse)
    - [ListAliasIPPoolRequest](#github.com.Juniper.contrail.pkg.models.ListAliasIPPoolRequest)
    - [ListAliasIPPoolResponse](#github.com.Juniper.contrail.pkg.models.ListAliasIPPoolResponse)
    - [ListAliasIPRequest](#github.com.Juniper.contrail.pkg.models.ListAliasIPRequest)
    - [ListAliasIPResponse](#github.com.Juniper.contrail.pkg.models.ListAliasIPResponse)
    - [ListAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.ListAnalyticsNodeRequest)
    - [ListAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.ListAnalyticsNodeResponse)
    - [ListAppformixNodeRequest](#github.com.Juniper.contrail.pkg.models.ListAppformixNodeRequest)
    - [ListAppformixNodeResponse](#github.com.Juniper.contrail.pkg.models.ListAppformixNodeResponse)
    - [ListApplicationPolicySetRequest](#github.com.Juniper.contrail.pkg.models.ListApplicationPolicySetRequest)
    - [ListApplicationPolicySetResponse](#github.com.Juniper.contrail.pkg.models.ListApplicationPolicySetResponse)
    - [ListBGPAsAServiceRequest](#github.com.Juniper.contrail.pkg.models.ListBGPAsAServiceRequest)
    - [ListBGPAsAServiceResponse](#github.com.Juniper.contrail.pkg.models.ListBGPAsAServiceResponse)
    - [ListBGPRouterRequest](#github.com.Juniper.contrail.pkg.models.ListBGPRouterRequest)
    - [ListBGPRouterResponse](#github.com.Juniper.contrail.pkg.models.ListBGPRouterResponse)
    - [ListBGPVPNRequest](#github.com.Juniper.contrail.pkg.models.ListBGPVPNRequest)
    - [ListBGPVPNResponse](#github.com.Juniper.contrail.pkg.models.ListBGPVPNResponse)
    - [ListBaremetalNodeRequest](#github.com.Juniper.contrail.pkg.models.ListBaremetalNodeRequest)
    - [ListBaremetalNodeResponse](#github.com.Juniper.contrail.pkg.models.ListBaremetalNodeResponse)
    - [ListBaremetalPortRequest](#github.com.Juniper.contrail.pkg.models.ListBaremetalPortRequest)
    - [ListBaremetalPortResponse](#github.com.Juniper.contrail.pkg.models.ListBaremetalPortResponse)
    - [ListBridgeDomainRequest](#github.com.Juniper.contrail.pkg.models.ListBridgeDomainRequest)
    - [ListBridgeDomainResponse](#github.com.Juniper.contrail.pkg.models.ListBridgeDomainResponse)
    - [ListConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.ListConfigNodeRequest)
    - [ListConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.ListConfigNodeResponse)
    - [ListConfigRootRequest](#github.com.Juniper.contrail.pkg.models.ListConfigRootRequest)
    - [ListConfigRootResponse](#github.com.Juniper.contrail.pkg.models.ListConfigRootResponse)
    - [ListContrailAnalyticsDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsDatabaseNodeRequest)
    - [ListContrailAnalyticsDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsDatabaseNodeResponse)
    - [ListContrailAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsNodeRequest)
    - [ListContrailAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsNodeResponse)
    - [ListContrailClusterRequest](#github.com.Juniper.contrail.pkg.models.ListContrailClusterRequest)
    - [ListContrailClusterResponse](#github.com.Juniper.contrail.pkg.models.ListContrailClusterResponse)
    - [ListContrailConfigDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailConfigDatabaseNodeRequest)
    - [ListContrailConfigDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailConfigDatabaseNodeResponse)
    - [ListContrailConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailConfigNodeRequest)
    - [ListContrailConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailConfigNodeResponse)
    - [ListContrailControlNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailControlNodeRequest)
    - [ListContrailControlNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailControlNodeResponse)
    - [ListContrailStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailStorageNodeRequest)
    - [ListContrailStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailStorageNodeResponse)
    - [ListContrailVrouterNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailVrouterNodeRequest)
    - [ListContrailVrouterNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailVrouterNodeResponse)
    - [ListContrailWebuiNodeRequest](#github.com.Juniper.contrail.pkg.models.ListContrailWebuiNodeRequest)
    - [ListContrailWebuiNodeResponse](#github.com.Juniper.contrail.pkg.models.ListContrailWebuiNodeResponse)
    - [ListCustomerAttachmentRequest](#github.com.Juniper.contrail.pkg.models.ListCustomerAttachmentRequest)
    - [ListCustomerAttachmentResponse](#github.com.Juniper.contrail.pkg.models.ListCustomerAttachmentResponse)
    - [ListDashboardRequest](#github.com.Juniper.contrail.pkg.models.ListDashboardRequest)
    - [ListDashboardResponse](#github.com.Juniper.contrail.pkg.models.ListDashboardResponse)
    - [ListDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.ListDatabaseNodeRequest)
    - [ListDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.ListDatabaseNodeResponse)
    - [ListDiscoveryServiceAssignmentRequest](#github.com.Juniper.contrail.pkg.models.ListDiscoveryServiceAssignmentRequest)
    - [ListDiscoveryServiceAssignmentResponse](#github.com.Juniper.contrail.pkg.models.ListDiscoveryServiceAssignmentResponse)
    - [ListDomainRequest](#github.com.Juniper.contrail.pkg.models.ListDomainRequest)
    - [ListDomainResponse](#github.com.Juniper.contrail.pkg.models.ListDomainResponse)
    - [ListDsaRuleRequest](#github.com.Juniper.contrail.pkg.models.ListDsaRuleRequest)
    - [ListDsaRuleResponse](#github.com.Juniper.contrail.pkg.models.ListDsaRuleResponse)
    - [ListE2ServiceProviderRequest](#github.com.Juniper.contrail.pkg.models.ListE2ServiceProviderRequest)
    - [ListE2ServiceProviderResponse](#github.com.Juniper.contrail.pkg.models.ListE2ServiceProviderResponse)
    - [ListFirewallPolicyRequest](#github.com.Juniper.contrail.pkg.models.ListFirewallPolicyRequest)
    - [ListFirewallPolicyResponse](#github.com.Juniper.contrail.pkg.models.ListFirewallPolicyResponse)
    - [ListFirewallRuleRequest](#github.com.Juniper.contrail.pkg.models.ListFirewallRuleRequest)
    - [ListFirewallRuleResponse](#github.com.Juniper.contrail.pkg.models.ListFirewallRuleResponse)
    - [ListFlavorRequest](#github.com.Juniper.contrail.pkg.models.ListFlavorRequest)
    - [ListFlavorResponse](#github.com.Juniper.contrail.pkg.models.ListFlavorResponse)
    - [ListFloatingIPPoolRequest](#github.com.Juniper.contrail.pkg.models.ListFloatingIPPoolRequest)
    - [ListFloatingIPPoolResponse](#github.com.Juniper.contrail.pkg.models.ListFloatingIPPoolResponse)
    - [ListFloatingIPRequest](#github.com.Juniper.contrail.pkg.models.ListFloatingIPRequest)
    - [ListFloatingIPResponse](#github.com.Juniper.contrail.pkg.models.ListFloatingIPResponse)
    - [ListForwardingClassRequest](#github.com.Juniper.contrail.pkg.models.ListForwardingClassRequest)
    - [ListForwardingClassResponse](#github.com.Juniper.contrail.pkg.models.ListForwardingClassResponse)
    - [ListGlobalQosConfigRequest](#github.com.Juniper.contrail.pkg.models.ListGlobalQosConfigRequest)
    - [ListGlobalQosConfigResponse](#github.com.Juniper.contrail.pkg.models.ListGlobalQosConfigResponse)
    - [ListGlobalSystemConfigRequest](#github.com.Juniper.contrail.pkg.models.ListGlobalSystemConfigRequest)
    - [ListGlobalSystemConfigResponse](#github.com.Juniper.contrail.pkg.models.ListGlobalSystemConfigResponse)
    - [ListGlobalVrouterConfigRequest](#github.com.Juniper.contrail.pkg.models.ListGlobalVrouterConfigRequest)
    - [ListGlobalVrouterConfigResponse](#github.com.Juniper.contrail.pkg.models.ListGlobalVrouterConfigResponse)
    - [ListInstanceIPRequest](#github.com.Juniper.contrail.pkg.models.ListInstanceIPRequest)
    - [ListInstanceIPResponse](#github.com.Juniper.contrail.pkg.models.ListInstanceIPResponse)
    - [ListInterfaceRouteTableRequest](#github.com.Juniper.contrail.pkg.models.ListInterfaceRouteTableRequest)
    - [ListInterfaceRouteTableResponse](#github.com.Juniper.contrail.pkg.models.ListInterfaceRouteTableResponse)
    - [ListKeypairRequest](#github.com.Juniper.contrail.pkg.models.ListKeypairRequest)
    - [ListKeypairResponse](#github.com.Juniper.contrail.pkg.models.ListKeypairResponse)
    - [ListKubernetesMasterNodeRequest](#github.com.Juniper.contrail.pkg.models.ListKubernetesMasterNodeRequest)
    - [ListKubernetesMasterNodeResponse](#github.com.Juniper.contrail.pkg.models.ListKubernetesMasterNodeResponse)
    - [ListKubernetesNodeRequest](#github.com.Juniper.contrail.pkg.models.ListKubernetesNodeRequest)
    - [ListKubernetesNodeResponse](#github.com.Juniper.contrail.pkg.models.ListKubernetesNodeResponse)
    - [ListLoadbalancerHealthmonitorRequest](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerHealthmonitorRequest)
    - [ListLoadbalancerHealthmonitorResponse](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerHealthmonitorResponse)
    - [ListLoadbalancerListenerRequest](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerListenerRequest)
    - [ListLoadbalancerListenerResponse](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerListenerResponse)
    - [ListLoadbalancerMemberRequest](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerMemberRequest)
    - [ListLoadbalancerMemberResponse](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerMemberResponse)
    - [ListLoadbalancerPoolRequest](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerPoolRequest)
    - [ListLoadbalancerPoolResponse](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerPoolResponse)
    - [ListLoadbalancerRequest](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerRequest)
    - [ListLoadbalancerResponse](#github.com.Juniper.contrail.pkg.models.ListLoadbalancerResponse)
    - [ListLocationRequest](#github.com.Juniper.contrail.pkg.models.ListLocationRequest)
    - [ListLocationResponse](#github.com.Juniper.contrail.pkg.models.ListLocationResponse)
    - [ListLogicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.ListLogicalInterfaceRequest)
    - [ListLogicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.ListLogicalInterfaceResponse)
    - [ListLogicalRouterRequest](#github.com.Juniper.contrail.pkg.models.ListLogicalRouterRequest)
    - [ListLogicalRouterResponse](#github.com.Juniper.contrail.pkg.models.ListLogicalRouterResponse)
    - [ListNamespaceRequest](#github.com.Juniper.contrail.pkg.models.ListNamespaceRequest)
    - [ListNamespaceResponse](#github.com.Juniper.contrail.pkg.models.ListNamespaceResponse)
    - [ListNetworkDeviceConfigRequest](#github.com.Juniper.contrail.pkg.models.ListNetworkDeviceConfigRequest)
    - [ListNetworkDeviceConfigResponse](#github.com.Juniper.contrail.pkg.models.ListNetworkDeviceConfigResponse)
    - [ListNetworkIpamRequest](#github.com.Juniper.contrail.pkg.models.ListNetworkIpamRequest)
    - [ListNetworkIpamResponse](#github.com.Juniper.contrail.pkg.models.ListNetworkIpamResponse)
    - [ListNetworkPolicyRequest](#github.com.Juniper.contrail.pkg.models.ListNetworkPolicyRequest)
    - [ListNetworkPolicyResponse](#github.com.Juniper.contrail.pkg.models.ListNetworkPolicyResponse)
    - [ListNodeRequest](#github.com.Juniper.contrail.pkg.models.ListNodeRequest)
    - [ListNodeResponse](#github.com.Juniper.contrail.pkg.models.ListNodeResponse)
    - [ListOpenstackComputeNodeRequest](#github.com.Juniper.contrail.pkg.models.ListOpenstackComputeNodeRequest)
    - [ListOpenstackComputeNodeResponse](#github.com.Juniper.contrail.pkg.models.ListOpenstackComputeNodeResponse)
    - [ListOpenstackControlNodeRequest](#github.com.Juniper.contrail.pkg.models.ListOpenstackControlNodeRequest)
    - [ListOpenstackControlNodeResponse](#github.com.Juniper.contrail.pkg.models.ListOpenstackControlNodeResponse)
    - [ListOpenstackMonitoringNodeRequest](#github.com.Juniper.contrail.pkg.models.ListOpenstackMonitoringNodeRequest)
    - [ListOpenstackMonitoringNodeResponse](#github.com.Juniper.contrail.pkg.models.ListOpenstackMonitoringNodeResponse)
    - [ListOpenstackNetworkNodeRequest](#github.com.Juniper.contrail.pkg.models.ListOpenstackNetworkNodeRequest)
    - [ListOpenstackNetworkNodeResponse](#github.com.Juniper.contrail.pkg.models.ListOpenstackNetworkNodeResponse)
    - [ListOpenstackStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.ListOpenstackStorageNodeRequest)
    - [ListOpenstackStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.ListOpenstackStorageNodeResponse)
    - [ListOsImageRequest](#github.com.Juniper.contrail.pkg.models.ListOsImageRequest)
    - [ListOsImageResponse](#github.com.Juniper.contrail.pkg.models.ListOsImageResponse)
    - [ListPeeringPolicyRequest](#github.com.Juniper.contrail.pkg.models.ListPeeringPolicyRequest)
    - [ListPeeringPolicyResponse](#github.com.Juniper.contrail.pkg.models.ListPeeringPolicyResponse)
    - [ListPhysicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.ListPhysicalInterfaceRequest)
    - [ListPhysicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.ListPhysicalInterfaceResponse)
    - [ListPhysicalRouterRequest](#github.com.Juniper.contrail.pkg.models.ListPhysicalRouterRequest)
    - [ListPhysicalRouterResponse](#github.com.Juniper.contrail.pkg.models.ListPhysicalRouterResponse)
    - [ListPolicyManagementRequest](#github.com.Juniper.contrail.pkg.models.ListPolicyManagementRequest)
    - [ListPolicyManagementResponse](#github.com.Juniper.contrail.pkg.models.ListPolicyManagementResponse)
    - [ListPortRequest](#github.com.Juniper.contrail.pkg.models.ListPortRequest)
    - [ListPortResponse](#github.com.Juniper.contrail.pkg.models.ListPortResponse)
    - [ListPortTupleRequest](#github.com.Juniper.contrail.pkg.models.ListPortTupleRequest)
    - [ListPortTupleResponse](#github.com.Juniper.contrail.pkg.models.ListPortTupleResponse)
    - [ListProjectRequest](#github.com.Juniper.contrail.pkg.models.ListProjectRequest)
    - [ListProjectResponse](#github.com.Juniper.contrail.pkg.models.ListProjectResponse)
    - [ListProviderAttachmentRequest](#github.com.Juniper.contrail.pkg.models.ListProviderAttachmentRequest)
    - [ListProviderAttachmentResponse](#github.com.Juniper.contrail.pkg.models.ListProviderAttachmentResponse)
    - [ListQosConfigRequest](#github.com.Juniper.contrail.pkg.models.ListQosConfigRequest)
    - [ListQosConfigResponse](#github.com.Juniper.contrail.pkg.models.ListQosConfigResponse)
    - [ListQosQueueRequest](#github.com.Juniper.contrail.pkg.models.ListQosQueueRequest)
    - [ListQosQueueResponse](#github.com.Juniper.contrail.pkg.models.ListQosQueueResponse)
    - [ListRouteAggregateRequest](#github.com.Juniper.contrail.pkg.models.ListRouteAggregateRequest)
    - [ListRouteAggregateResponse](#github.com.Juniper.contrail.pkg.models.ListRouteAggregateResponse)
    - [ListRouteTableRequest](#github.com.Juniper.contrail.pkg.models.ListRouteTableRequest)
    - [ListRouteTableResponse](#github.com.Juniper.contrail.pkg.models.ListRouteTableResponse)
    - [ListRouteTargetRequest](#github.com.Juniper.contrail.pkg.models.ListRouteTargetRequest)
    - [ListRouteTargetResponse](#github.com.Juniper.contrail.pkg.models.ListRouteTargetResponse)
    - [ListRoutingInstanceRequest](#github.com.Juniper.contrail.pkg.models.ListRoutingInstanceRequest)
    - [ListRoutingInstanceResponse](#github.com.Juniper.contrail.pkg.models.ListRoutingInstanceResponse)
    - [ListRoutingPolicyRequest](#github.com.Juniper.contrail.pkg.models.ListRoutingPolicyRequest)
    - [ListRoutingPolicyResponse](#github.com.Juniper.contrail.pkg.models.ListRoutingPolicyResponse)
    - [ListSecurityGroupRequest](#github.com.Juniper.contrail.pkg.models.ListSecurityGroupRequest)
    - [ListSecurityGroupResponse](#github.com.Juniper.contrail.pkg.models.ListSecurityGroupResponse)
    - [ListSecurityLoggingObjectRequest](#github.com.Juniper.contrail.pkg.models.ListSecurityLoggingObjectRequest)
    - [ListSecurityLoggingObjectResponse](#github.com.Juniper.contrail.pkg.models.ListSecurityLoggingObjectResponse)
    - [ListServerRequest](#github.com.Juniper.contrail.pkg.models.ListServerRequest)
    - [ListServerResponse](#github.com.Juniper.contrail.pkg.models.ListServerResponse)
    - [ListServiceApplianceRequest](#github.com.Juniper.contrail.pkg.models.ListServiceApplianceRequest)
    - [ListServiceApplianceResponse](#github.com.Juniper.contrail.pkg.models.ListServiceApplianceResponse)
    - [ListServiceApplianceSetRequest](#github.com.Juniper.contrail.pkg.models.ListServiceApplianceSetRequest)
    - [ListServiceApplianceSetResponse](#github.com.Juniper.contrail.pkg.models.ListServiceApplianceSetResponse)
    - [ListServiceConnectionModuleRequest](#github.com.Juniper.contrail.pkg.models.ListServiceConnectionModuleRequest)
    - [ListServiceConnectionModuleResponse](#github.com.Juniper.contrail.pkg.models.ListServiceConnectionModuleResponse)
    - [ListServiceEndpointRequest](#github.com.Juniper.contrail.pkg.models.ListServiceEndpointRequest)
    - [ListServiceEndpointResponse](#github.com.Juniper.contrail.pkg.models.ListServiceEndpointResponse)
    - [ListServiceGroupRequest](#github.com.Juniper.contrail.pkg.models.ListServiceGroupRequest)
    - [ListServiceGroupResponse](#github.com.Juniper.contrail.pkg.models.ListServiceGroupResponse)
    - [ListServiceHealthCheckRequest](#github.com.Juniper.contrail.pkg.models.ListServiceHealthCheckRequest)
    - [ListServiceHealthCheckResponse](#github.com.Juniper.contrail.pkg.models.ListServiceHealthCheckResponse)
    - [ListServiceInstanceRequest](#github.com.Juniper.contrail.pkg.models.ListServiceInstanceRequest)
    - [ListServiceInstanceResponse](#github.com.Juniper.contrail.pkg.models.ListServiceInstanceResponse)
    - [ListServiceObjectRequest](#github.com.Juniper.contrail.pkg.models.ListServiceObjectRequest)
    - [ListServiceObjectResponse](#github.com.Juniper.contrail.pkg.models.ListServiceObjectResponse)
    - [ListServiceTemplateRequest](#github.com.Juniper.contrail.pkg.models.ListServiceTemplateRequest)
    - [ListServiceTemplateResponse](#github.com.Juniper.contrail.pkg.models.ListServiceTemplateResponse)
    - [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec)
    - [ListSubnetRequest](#github.com.Juniper.contrail.pkg.models.ListSubnetRequest)
    - [ListSubnetResponse](#github.com.Juniper.contrail.pkg.models.ListSubnetResponse)
    - [ListTagRequest](#github.com.Juniper.contrail.pkg.models.ListTagRequest)
    - [ListTagResponse](#github.com.Juniper.contrail.pkg.models.ListTagResponse)
    - [ListTagTypeRequest](#github.com.Juniper.contrail.pkg.models.ListTagTypeRequest)
    - [ListTagTypeResponse](#github.com.Juniper.contrail.pkg.models.ListTagTypeResponse)
    - [ListUserRequest](#github.com.Juniper.contrail.pkg.models.ListUserRequest)
    - [ListUserResponse](#github.com.Juniper.contrail.pkg.models.ListUserResponse)
    - [ListVPNGroupRequest](#github.com.Juniper.contrail.pkg.models.ListVPNGroupRequest)
    - [ListVPNGroupResponse](#github.com.Juniper.contrail.pkg.models.ListVPNGroupResponse)
    - [ListVirtualDNSRecordRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualDNSRecordRequest)
    - [ListVirtualDNSRecordResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualDNSRecordResponse)
    - [ListVirtualDNSRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualDNSRequest)
    - [ListVirtualDNSResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualDNSResponse)
    - [ListVirtualIPRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualIPRequest)
    - [ListVirtualIPResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualIPResponse)
    - [ListVirtualMachineInterfaceRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualMachineInterfaceRequest)
    - [ListVirtualMachineInterfaceResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualMachineInterfaceResponse)
    - [ListVirtualMachineRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualMachineRequest)
    - [ListVirtualMachineResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualMachineResponse)
    - [ListVirtualNetworkRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualNetworkRequest)
    - [ListVirtualNetworkResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualNetworkResponse)
    - [ListVirtualRouterRequest](#github.com.Juniper.contrail.pkg.models.ListVirtualRouterRequest)
    - [ListVirtualRouterResponse](#github.com.Juniper.contrail.pkg.models.ListVirtualRouterResponse)
    - [ListWidgetRequest](#github.com.Juniper.contrail.pkg.models.ListWidgetRequest)
    - [ListWidgetResponse](#github.com.Juniper.contrail.pkg.models.ListWidgetResponse)
    - [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer)
    - [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor)
    - [LoadbalancerHealthmonitorType](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitorType)
    - [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener)
    - [LoadbalancerListenerLoadbalancerRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerListenerLoadbalancerRef)
    - [LoadbalancerListenerType](#github.com.Juniper.contrail.pkg.models.LoadbalancerListenerType)
    - [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember)
    - [LoadbalancerMemberType](#github.com.Juniper.contrail.pkg.models.LoadbalancerMemberType)
    - [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool)
    - [LoadbalancerPoolLoadbalancerHealthmonitorRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolLoadbalancerHealthmonitorRef)
    - [LoadbalancerPoolLoadbalancerListenerRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolLoadbalancerListenerRef)
    - [LoadbalancerPoolServiceApplianceSetRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolServiceApplianceSetRef)
    - [LoadbalancerPoolServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolServiceInstanceRef)
    - [LoadbalancerPoolType](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolType)
    - [LoadbalancerPoolVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolVirtualMachineInterfaceRef)
    - [LoadbalancerServiceApplianceSetRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerServiceApplianceSetRef)
    - [LoadbalancerServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerServiceInstanceRef)
    - [LoadbalancerType](#github.com.Juniper.contrail.pkg.models.LoadbalancerType)
    - [LoadbalancerVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerVirtualMachineInterfaceRef)
    - [LocalLinkConnection](#github.com.Juniper.contrail.pkg.models.LocalLinkConnection)
    - [Location](#github.com.Juniper.contrail.pkg.models.Location)
    - [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface)
    - [LogicalInterfaceVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LogicalInterfaceVirtualMachineInterfaceRef)
    - [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter)
    - [LogicalRouterBGPVPNRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterBGPVPNRef)
    - [LogicalRouterPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterPhysicalRouterRef)
    - [LogicalRouterRouteTableRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterRouteTableRef)
    - [LogicalRouterRouteTargetRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterRouteTargetRef)
    - [LogicalRouterServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterServiceInstanceRef)
    - [LogicalRouterVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterVirtualMachineInterfaceRef)
    - [LogicalRouterVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterVirtualNetworkRef)
    - [MACLimitControlType](#github.com.Juniper.contrail.pkg.models.MACLimitControlType)
    - [MACMoveLimitControlType](#github.com.Juniper.contrail.pkg.models.MACMoveLimitControlType)
    - [MacAddressesType](#github.com.Juniper.contrail.pkg.models.MacAddressesType)
    - [MatchConditionType](#github.com.Juniper.contrail.pkg.models.MatchConditionType)
    - [MemberType](#github.com.Juniper.contrail.pkg.models.MemberType)
    - [MirrorActionType](#github.com.Juniper.contrail.pkg.models.MirrorActionType)
    - [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace)
    - [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig)
    - [NetworkDeviceConfigPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfigPhysicalRouterRef)
    - [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam)
    - [NetworkIpamVirtualDNSRef](#github.com.Juniper.contrail.pkg.models.NetworkIpamVirtualDNSRef)
    - [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy)
    - [Node](#github.com.Juniper.contrail.pkg.models.Node)
    - [NodeKeypairRef](#github.com.Juniper.contrail.pkg.models.NodeKeypairRef)
    - [OpenStackAddress](#github.com.Juniper.contrail.pkg.models.OpenStackAddress)
    - [OpenStackFlavorProperty](#github.com.Juniper.contrail.pkg.models.OpenStackFlavorProperty)
    - [OpenStackImageProperty](#github.com.Juniper.contrail.pkg.models.OpenStackImageProperty)
    - [OpenStackLink](#github.com.Juniper.contrail.pkg.models.OpenStackLink)
    - [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode)
    - [OpenstackComputeNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNodeNodeRef)
    - [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode)
    - [OpenstackControlNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackControlNodeNodeRef)
    - [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode)
    - [OpenstackMonitoringNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNodeNodeRef)
    - [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode)
    - [OpenstackNetworkNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNodeNodeRef)
    - [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode)
    - [OpenstackStorageNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNodeNodeRef)
    - [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage)
    - [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy)
    - [PermType](#github.com.Juniper.contrail.pkg.models.PermType)
    - [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2)
    - [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface)
    - [PhysicalInterfacePhysicalInterfaceRef](#github.com.Juniper.contrail.pkg.models.PhysicalInterfacePhysicalInterfaceRef)
    - [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter)
    - [PhysicalRouterBGPRouterRef](#github.com.Juniper.contrail.pkg.models.PhysicalRouterBGPRouterRef)
    - [PhysicalRouterVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.PhysicalRouterVirtualNetworkRef)
    - [PhysicalRouterVirtualRouterRef](#github.com.Juniper.contrail.pkg.models.PhysicalRouterVirtualRouterRef)
    - [PluginProperties](#github.com.Juniper.contrail.pkg.models.PluginProperties)
    - [PluginProperty](#github.com.Juniper.contrail.pkg.models.PluginProperty)
    - [PolicyBasedForwardingRuleType](#github.com.Juniper.contrail.pkg.models.PolicyBasedForwardingRuleType)
    - [PolicyEntriesType](#github.com.Juniper.contrail.pkg.models.PolicyEntriesType)
    - [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement)
    - [PolicyRuleType](#github.com.Juniper.contrail.pkg.models.PolicyRuleType)
    - [Port](#github.com.Juniper.contrail.pkg.models.Port)
    - [PortMap](#github.com.Juniper.contrail.pkg.models.PortMap)
    - [PortMappings](#github.com.Juniper.contrail.pkg.models.PortMappings)
    - [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple)
    - [PortType](#github.com.Juniper.contrail.pkg.models.PortType)
    - [Project](#github.com.Juniper.contrail.pkg.models.Project)
    - [ProjectAliasIPPoolRef](#github.com.Juniper.contrail.pkg.models.ProjectAliasIPPoolRef)
    - [ProjectApplicationPolicySetRef](#github.com.Juniper.contrail.pkg.models.ProjectApplicationPolicySetRef)
    - [ProjectFloatingIPPoolRef](#github.com.Juniper.contrail.pkg.models.ProjectFloatingIPPoolRef)
    - [ProjectNamespaceRef](#github.com.Juniper.contrail.pkg.models.ProjectNamespaceRef)
    - [ProtocolType](#github.com.Juniper.contrail.pkg.models.ProtocolType)
    - [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment)
    - [ProviderAttachmentVirtualRouterRef](#github.com.Juniper.contrail.pkg.models.ProviderAttachmentVirtualRouterRef)
    - [ProviderDetails](#github.com.Juniper.contrail.pkg.models.ProviderDetails)
    - [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig)
    - [QosConfigGlobalSystemConfigRef](#github.com.Juniper.contrail.pkg.models.QosConfigGlobalSystemConfigRef)
    - [QosIdForwardingClassPair](#github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPair)
    - [QosIdForwardingClassPairs](#github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPairs)
    - [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue)
    - [QuotaType](#github.com.Juniper.contrail.pkg.models.QuotaType)
    - [RbacPermType](#github.com.Juniper.contrail.pkg.models.RbacPermType)
    - [RbacRuleEntriesType](#github.com.Juniper.contrail.pkg.models.RbacRuleEntriesType)
    - [RbacRuleType](#github.com.Juniper.contrail.pkg.models.RbacRuleType)
    - [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate)
    - [RouteAggregateServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.RouteAggregateServiceInstanceRef)
    - [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable)
    - [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType)
    - [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget)
    - [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList)
    - [RouteType](#github.com.Juniper.contrail.pkg.models.RouteType)
    - [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance)
    - [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy)
    - [RoutingPolicyServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.RoutingPolicyServiceInstanceRef)
    - [RoutingPolicyServiceInstanceType](#github.com.Juniper.contrail.pkg.models.RoutingPolicyServiceInstanceType)
    - [SNMPCredentials](#github.com.Juniper.contrail.pkg.models.SNMPCredentials)
    - [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup)
    - [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject)
    - [SecurityLoggingObjectNetworkPolicyRef](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectNetworkPolicyRef)
    - [SecurityLoggingObjectRuleEntryType](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleEntryType)
    - [SecurityLoggingObjectRuleListType](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleListType)
    - [SecurityLoggingObjectSecurityGroupRef](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectSecurityGroupRef)
    - [SequenceType](#github.com.Juniper.contrail.pkg.models.SequenceType)
    - [Server](#github.com.Juniper.contrail.pkg.models.Server)
    - [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance)
    - [ServiceApplianceInterfaceType](#github.com.Juniper.contrail.pkg.models.ServiceApplianceInterfaceType)
    - [ServiceAppliancePhysicalInterfaceRef](#github.com.Juniper.contrail.pkg.models.ServiceAppliancePhysicalInterfaceRef)
    - [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet)
    - [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule)
    - [ServiceConnectionModuleServiceObjectRef](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModuleServiceObjectRef)
    - [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint)
    - [ServiceEndpointPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.ServiceEndpointPhysicalRouterRef)
    - [ServiceEndpointServiceConnectionModuleRef](#github.com.Juniper.contrail.pkg.models.ServiceEndpointServiceConnectionModuleRef)
    - [ServiceEndpointServiceObjectRef](#github.com.Juniper.contrail.pkg.models.ServiceEndpointServiceObjectRef)
    - [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup)
    - [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck)
    - [ServiceHealthCheckServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheckServiceInstanceRef)
    - [ServiceHealthCheckType](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheckType)
    - [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance)
    - [ServiceInstanceInstanceIPRef](#github.com.Juniper.contrail.pkg.models.ServiceInstanceInstanceIPRef)
    - [ServiceInstanceInterfaceType](#github.com.Juniper.contrail.pkg.models.ServiceInstanceInterfaceType)
    - [ServiceInstanceServiceTemplateRef](#github.com.Juniper.contrail.pkg.models.ServiceInstanceServiceTemplateRef)
    - [ServiceInstanceType](#github.com.Juniper.contrail.pkg.models.ServiceInstanceType)
    - [ServiceInterfaceTag](#github.com.Juniper.contrail.pkg.models.ServiceInterfaceTag)
    - [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject)
    - [ServiceScaleOutType](#github.com.Juniper.contrail.pkg.models.ServiceScaleOutType)
    - [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate)
    - [ServiceTemplateInterfaceType](#github.com.Juniper.contrail.pkg.models.ServiceTemplateInterfaceType)
    - [ServiceTemplateServiceApplianceSetRef](#github.com.Juniper.contrail.pkg.models.ServiceTemplateServiceApplianceSetRef)
    - [ServiceTemplateType](#github.com.Juniper.contrail.pkg.models.ServiceTemplateType)
    - [ShareType](#github.com.Juniper.contrail.pkg.models.ShareType)
    - [StaticMirrorNhType](#github.com.Juniper.contrail.pkg.models.StaticMirrorNhType)
    - [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet)
    - [SubnetListType](#github.com.Juniper.contrail.pkg.models.SubnetListType)
    - [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType)
    - [SubnetVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.SubnetVirtualMachineInterfaceRef)
    - [Tag](#github.com.Juniper.contrail.pkg.models.Tag)
    - [TagTagTypeRef](#github.com.Juniper.contrail.pkg.models.TagTagTypeRef)
    - [TagType](#github.com.Juniper.contrail.pkg.models.TagType)
    - [TelemetryResourceInfo](#github.com.Juniper.contrail.pkg.models.TelemetryResourceInfo)
    - [TelemetryStateInfo](#github.com.Juniper.contrail.pkg.models.TelemetryStateInfo)
    - [TimerType](#github.com.Juniper.contrail.pkg.models.TimerType)
    - [UpdateAPIAccessListRequest](#github.com.Juniper.contrail.pkg.models.UpdateAPIAccessListRequest)
    - [UpdateAPIAccessListResponse](#github.com.Juniper.contrail.pkg.models.UpdateAPIAccessListResponse)
    - [UpdateAccessControlListRequest](#github.com.Juniper.contrail.pkg.models.UpdateAccessControlListRequest)
    - [UpdateAccessControlListResponse](#github.com.Juniper.contrail.pkg.models.UpdateAccessControlListResponse)
    - [UpdateAddressGroupRequest](#github.com.Juniper.contrail.pkg.models.UpdateAddressGroupRequest)
    - [UpdateAddressGroupResponse](#github.com.Juniper.contrail.pkg.models.UpdateAddressGroupResponse)
    - [UpdateAlarmRequest](#github.com.Juniper.contrail.pkg.models.UpdateAlarmRequest)
    - [UpdateAlarmResponse](#github.com.Juniper.contrail.pkg.models.UpdateAlarmResponse)
    - [UpdateAliasIPPoolRequest](#github.com.Juniper.contrail.pkg.models.UpdateAliasIPPoolRequest)
    - [UpdateAliasIPPoolResponse](#github.com.Juniper.contrail.pkg.models.UpdateAliasIPPoolResponse)
    - [UpdateAliasIPRequest](#github.com.Juniper.contrail.pkg.models.UpdateAliasIPRequest)
    - [UpdateAliasIPResponse](#github.com.Juniper.contrail.pkg.models.UpdateAliasIPResponse)
    - [UpdateAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateAnalyticsNodeRequest)
    - [UpdateAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateAnalyticsNodeResponse)
    - [UpdateAppformixNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateAppformixNodeRequest)
    - [UpdateAppformixNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateAppformixNodeResponse)
    - [UpdateApplicationPolicySetRequest](#github.com.Juniper.contrail.pkg.models.UpdateApplicationPolicySetRequest)
    - [UpdateApplicationPolicySetResponse](#github.com.Juniper.contrail.pkg.models.UpdateApplicationPolicySetResponse)
    - [UpdateBGPAsAServiceRequest](#github.com.Juniper.contrail.pkg.models.UpdateBGPAsAServiceRequest)
    - [UpdateBGPAsAServiceResponse](#github.com.Juniper.contrail.pkg.models.UpdateBGPAsAServiceResponse)
    - [UpdateBGPRouterRequest](#github.com.Juniper.contrail.pkg.models.UpdateBGPRouterRequest)
    - [UpdateBGPRouterResponse](#github.com.Juniper.contrail.pkg.models.UpdateBGPRouterResponse)
    - [UpdateBGPVPNRequest](#github.com.Juniper.contrail.pkg.models.UpdateBGPVPNRequest)
    - [UpdateBGPVPNResponse](#github.com.Juniper.contrail.pkg.models.UpdateBGPVPNResponse)
    - [UpdateBaremetalNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateBaremetalNodeRequest)
    - [UpdateBaremetalNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateBaremetalNodeResponse)
    - [UpdateBaremetalPortRequest](#github.com.Juniper.contrail.pkg.models.UpdateBaremetalPortRequest)
    - [UpdateBaremetalPortResponse](#github.com.Juniper.contrail.pkg.models.UpdateBaremetalPortResponse)
    - [UpdateBridgeDomainRequest](#github.com.Juniper.contrail.pkg.models.UpdateBridgeDomainRequest)
    - [UpdateBridgeDomainResponse](#github.com.Juniper.contrail.pkg.models.UpdateBridgeDomainResponse)
    - [UpdateConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateConfigNodeRequest)
    - [UpdateConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateConfigNodeResponse)
    - [UpdateConfigRootRequest](#github.com.Juniper.contrail.pkg.models.UpdateConfigRootRequest)
    - [UpdateConfigRootResponse](#github.com.Juniper.contrail.pkg.models.UpdateConfigRootResponse)
    - [UpdateContrailAnalyticsDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsDatabaseNodeRequest)
    - [UpdateContrailAnalyticsDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsDatabaseNodeResponse)
    - [UpdateContrailAnalyticsNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsNodeRequest)
    - [UpdateContrailAnalyticsNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsNodeResponse)
    - [UpdateContrailClusterRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailClusterRequest)
    - [UpdateContrailClusterResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailClusterResponse)
    - [UpdateContrailConfigDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailConfigDatabaseNodeRequest)
    - [UpdateContrailConfigDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailConfigDatabaseNodeResponse)
    - [UpdateContrailConfigNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailConfigNodeRequest)
    - [UpdateContrailConfigNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailConfigNodeResponse)
    - [UpdateContrailControlNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailControlNodeRequest)
    - [UpdateContrailControlNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailControlNodeResponse)
    - [UpdateContrailStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailStorageNodeRequest)
    - [UpdateContrailStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailStorageNodeResponse)
    - [UpdateContrailVrouterNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailVrouterNodeRequest)
    - [UpdateContrailVrouterNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailVrouterNodeResponse)
    - [UpdateContrailWebuiNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateContrailWebuiNodeRequest)
    - [UpdateContrailWebuiNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateContrailWebuiNodeResponse)
    - [UpdateCustomerAttachmentRequest](#github.com.Juniper.contrail.pkg.models.UpdateCustomerAttachmentRequest)
    - [UpdateCustomerAttachmentResponse](#github.com.Juniper.contrail.pkg.models.UpdateCustomerAttachmentResponse)
    - [UpdateDashboardRequest](#github.com.Juniper.contrail.pkg.models.UpdateDashboardRequest)
    - [UpdateDashboardResponse](#github.com.Juniper.contrail.pkg.models.UpdateDashboardResponse)
    - [UpdateDatabaseNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateDatabaseNodeRequest)
    - [UpdateDatabaseNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateDatabaseNodeResponse)
    - [UpdateDiscoveryServiceAssignmentRequest](#github.com.Juniper.contrail.pkg.models.UpdateDiscoveryServiceAssignmentRequest)
    - [UpdateDiscoveryServiceAssignmentResponse](#github.com.Juniper.contrail.pkg.models.UpdateDiscoveryServiceAssignmentResponse)
    - [UpdateDomainRequest](#github.com.Juniper.contrail.pkg.models.UpdateDomainRequest)
    - [UpdateDomainResponse](#github.com.Juniper.contrail.pkg.models.UpdateDomainResponse)
    - [UpdateDsaRuleRequest](#github.com.Juniper.contrail.pkg.models.UpdateDsaRuleRequest)
    - [UpdateDsaRuleResponse](#github.com.Juniper.contrail.pkg.models.UpdateDsaRuleResponse)
    - [UpdateE2ServiceProviderRequest](#github.com.Juniper.contrail.pkg.models.UpdateE2ServiceProviderRequest)
    - [UpdateE2ServiceProviderResponse](#github.com.Juniper.contrail.pkg.models.UpdateE2ServiceProviderResponse)
    - [UpdateFirewallPolicyRequest](#github.com.Juniper.contrail.pkg.models.UpdateFirewallPolicyRequest)
    - [UpdateFirewallPolicyResponse](#github.com.Juniper.contrail.pkg.models.UpdateFirewallPolicyResponse)
    - [UpdateFirewallRuleRequest](#github.com.Juniper.contrail.pkg.models.UpdateFirewallRuleRequest)
    - [UpdateFirewallRuleResponse](#github.com.Juniper.contrail.pkg.models.UpdateFirewallRuleResponse)
    - [UpdateFlavorRequest](#github.com.Juniper.contrail.pkg.models.UpdateFlavorRequest)
    - [UpdateFlavorResponse](#github.com.Juniper.contrail.pkg.models.UpdateFlavorResponse)
    - [UpdateFloatingIPPoolRequest](#github.com.Juniper.contrail.pkg.models.UpdateFloatingIPPoolRequest)
    - [UpdateFloatingIPPoolResponse](#github.com.Juniper.contrail.pkg.models.UpdateFloatingIPPoolResponse)
    - [UpdateFloatingIPRequest](#github.com.Juniper.contrail.pkg.models.UpdateFloatingIPRequest)
    - [UpdateFloatingIPResponse](#github.com.Juniper.contrail.pkg.models.UpdateFloatingIPResponse)
    - [UpdateForwardingClassRequest](#github.com.Juniper.contrail.pkg.models.UpdateForwardingClassRequest)
    - [UpdateForwardingClassResponse](#github.com.Juniper.contrail.pkg.models.UpdateForwardingClassResponse)
    - [UpdateGlobalQosConfigRequest](#github.com.Juniper.contrail.pkg.models.UpdateGlobalQosConfigRequest)
    - [UpdateGlobalQosConfigResponse](#github.com.Juniper.contrail.pkg.models.UpdateGlobalQosConfigResponse)
    - [UpdateGlobalSystemConfigRequest](#github.com.Juniper.contrail.pkg.models.UpdateGlobalSystemConfigRequest)
    - [UpdateGlobalSystemConfigResponse](#github.com.Juniper.contrail.pkg.models.UpdateGlobalSystemConfigResponse)
    - [UpdateGlobalVrouterConfigRequest](#github.com.Juniper.contrail.pkg.models.UpdateGlobalVrouterConfigRequest)
    - [UpdateGlobalVrouterConfigResponse](#github.com.Juniper.contrail.pkg.models.UpdateGlobalVrouterConfigResponse)
    - [UpdateInstanceIPRequest](#github.com.Juniper.contrail.pkg.models.UpdateInstanceIPRequest)
    - [UpdateInstanceIPResponse](#github.com.Juniper.contrail.pkg.models.UpdateInstanceIPResponse)
    - [UpdateInterfaceRouteTableRequest](#github.com.Juniper.contrail.pkg.models.UpdateInterfaceRouteTableRequest)
    - [UpdateInterfaceRouteTableResponse](#github.com.Juniper.contrail.pkg.models.UpdateInterfaceRouteTableResponse)
    - [UpdateKeypairRequest](#github.com.Juniper.contrail.pkg.models.UpdateKeypairRequest)
    - [UpdateKeypairResponse](#github.com.Juniper.contrail.pkg.models.UpdateKeypairResponse)
    - [UpdateKubernetesMasterNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateKubernetesMasterNodeRequest)
    - [UpdateKubernetesMasterNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateKubernetesMasterNodeResponse)
    - [UpdateKubernetesNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateKubernetesNodeRequest)
    - [UpdateKubernetesNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateKubernetesNodeResponse)
    - [UpdateLoadbalancerHealthmonitorRequest](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerHealthmonitorRequest)
    - [UpdateLoadbalancerHealthmonitorResponse](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerHealthmonitorResponse)
    - [UpdateLoadbalancerListenerRequest](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerListenerRequest)
    - [UpdateLoadbalancerListenerResponse](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerListenerResponse)
    - [UpdateLoadbalancerMemberRequest](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerMemberRequest)
    - [UpdateLoadbalancerMemberResponse](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerMemberResponse)
    - [UpdateLoadbalancerPoolRequest](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerPoolRequest)
    - [UpdateLoadbalancerPoolResponse](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerPoolResponse)
    - [UpdateLoadbalancerRequest](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerRequest)
    - [UpdateLoadbalancerResponse](#github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerResponse)
    - [UpdateLocationRequest](#github.com.Juniper.contrail.pkg.models.UpdateLocationRequest)
    - [UpdateLocationResponse](#github.com.Juniper.contrail.pkg.models.UpdateLocationResponse)
    - [UpdateLogicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.UpdateLogicalInterfaceRequest)
    - [UpdateLogicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.UpdateLogicalInterfaceResponse)
    - [UpdateLogicalRouterRequest](#github.com.Juniper.contrail.pkg.models.UpdateLogicalRouterRequest)
    - [UpdateLogicalRouterResponse](#github.com.Juniper.contrail.pkg.models.UpdateLogicalRouterResponse)
    - [UpdateNamespaceRequest](#github.com.Juniper.contrail.pkg.models.UpdateNamespaceRequest)
    - [UpdateNamespaceResponse](#github.com.Juniper.contrail.pkg.models.UpdateNamespaceResponse)
    - [UpdateNetworkDeviceConfigRequest](#github.com.Juniper.contrail.pkg.models.UpdateNetworkDeviceConfigRequest)
    - [UpdateNetworkDeviceConfigResponse](#github.com.Juniper.contrail.pkg.models.UpdateNetworkDeviceConfigResponse)
    - [UpdateNetworkIpamRequest](#github.com.Juniper.contrail.pkg.models.UpdateNetworkIpamRequest)
    - [UpdateNetworkIpamResponse](#github.com.Juniper.contrail.pkg.models.UpdateNetworkIpamResponse)
    - [UpdateNetworkPolicyRequest](#github.com.Juniper.contrail.pkg.models.UpdateNetworkPolicyRequest)
    - [UpdateNetworkPolicyResponse](#github.com.Juniper.contrail.pkg.models.UpdateNetworkPolicyResponse)
    - [UpdateNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateNodeRequest)
    - [UpdateNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateNodeResponse)
    - [UpdateOpenstackComputeNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackComputeNodeRequest)
    - [UpdateOpenstackComputeNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackComputeNodeResponse)
    - [UpdateOpenstackControlNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackControlNodeRequest)
    - [UpdateOpenstackControlNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackControlNodeResponse)
    - [UpdateOpenstackMonitoringNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackMonitoringNodeRequest)
    - [UpdateOpenstackMonitoringNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackMonitoringNodeResponse)
    - [UpdateOpenstackNetworkNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackNetworkNodeRequest)
    - [UpdateOpenstackNetworkNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackNetworkNodeResponse)
    - [UpdateOpenstackStorageNodeRequest](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackStorageNodeRequest)
    - [UpdateOpenstackStorageNodeResponse](#github.com.Juniper.contrail.pkg.models.UpdateOpenstackStorageNodeResponse)
    - [UpdateOsImageRequest](#github.com.Juniper.contrail.pkg.models.UpdateOsImageRequest)
    - [UpdateOsImageResponse](#github.com.Juniper.contrail.pkg.models.UpdateOsImageResponse)
    - [UpdatePeeringPolicyRequest](#github.com.Juniper.contrail.pkg.models.UpdatePeeringPolicyRequest)
    - [UpdatePeeringPolicyResponse](#github.com.Juniper.contrail.pkg.models.UpdatePeeringPolicyResponse)
    - [UpdatePhysicalInterfaceRequest](#github.com.Juniper.contrail.pkg.models.UpdatePhysicalInterfaceRequest)
    - [UpdatePhysicalInterfaceResponse](#github.com.Juniper.contrail.pkg.models.UpdatePhysicalInterfaceResponse)
    - [UpdatePhysicalRouterRequest](#github.com.Juniper.contrail.pkg.models.UpdatePhysicalRouterRequest)
    - [UpdatePhysicalRouterResponse](#github.com.Juniper.contrail.pkg.models.UpdatePhysicalRouterResponse)
    - [UpdatePolicyManagementRequest](#github.com.Juniper.contrail.pkg.models.UpdatePolicyManagementRequest)
    - [UpdatePolicyManagementResponse](#github.com.Juniper.contrail.pkg.models.UpdatePolicyManagementResponse)
    - [UpdatePortRequest](#github.com.Juniper.contrail.pkg.models.UpdatePortRequest)
    - [UpdatePortResponse](#github.com.Juniper.contrail.pkg.models.UpdatePortResponse)
    - [UpdatePortTupleRequest](#github.com.Juniper.contrail.pkg.models.UpdatePortTupleRequest)
    - [UpdatePortTupleResponse](#github.com.Juniper.contrail.pkg.models.UpdatePortTupleResponse)
    - [UpdateProjectRequest](#github.com.Juniper.contrail.pkg.models.UpdateProjectRequest)
    - [UpdateProjectResponse](#github.com.Juniper.contrail.pkg.models.UpdateProjectResponse)
    - [UpdateProviderAttachmentRequest](#github.com.Juniper.contrail.pkg.models.UpdateProviderAttachmentRequest)
    - [UpdateProviderAttachmentResponse](#github.com.Juniper.contrail.pkg.models.UpdateProviderAttachmentResponse)
    - [UpdateQosConfigRequest](#github.com.Juniper.contrail.pkg.models.UpdateQosConfigRequest)
    - [UpdateQosConfigResponse](#github.com.Juniper.contrail.pkg.models.UpdateQosConfigResponse)
    - [UpdateQosQueueRequest](#github.com.Juniper.contrail.pkg.models.UpdateQosQueueRequest)
    - [UpdateQosQueueResponse](#github.com.Juniper.contrail.pkg.models.UpdateQosQueueResponse)
    - [UpdateRouteAggregateRequest](#github.com.Juniper.contrail.pkg.models.UpdateRouteAggregateRequest)
    - [UpdateRouteAggregateResponse](#github.com.Juniper.contrail.pkg.models.UpdateRouteAggregateResponse)
    - [UpdateRouteTableRequest](#github.com.Juniper.contrail.pkg.models.UpdateRouteTableRequest)
    - [UpdateRouteTableResponse](#github.com.Juniper.contrail.pkg.models.UpdateRouteTableResponse)
    - [UpdateRouteTargetRequest](#github.com.Juniper.contrail.pkg.models.UpdateRouteTargetRequest)
    - [UpdateRouteTargetResponse](#github.com.Juniper.contrail.pkg.models.UpdateRouteTargetResponse)
    - [UpdateRoutingInstanceRequest](#github.com.Juniper.contrail.pkg.models.UpdateRoutingInstanceRequest)
    - [UpdateRoutingInstanceResponse](#github.com.Juniper.contrail.pkg.models.UpdateRoutingInstanceResponse)
    - [UpdateRoutingPolicyRequest](#github.com.Juniper.contrail.pkg.models.UpdateRoutingPolicyRequest)
    - [UpdateRoutingPolicyResponse](#github.com.Juniper.contrail.pkg.models.UpdateRoutingPolicyResponse)
    - [UpdateSecurityGroupRequest](#github.com.Juniper.contrail.pkg.models.UpdateSecurityGroupRequest)
    - [UpdateSecurityGroupResponse](#github.com.Juniper.contrail.pkg.models.UpdateSecurityGroupResponse)
    - [UpdateSecurityLoggingObjectRequest](#github.com.Juniper.contrail.pkg.models.UpdateSecurityLoggingObjectRequest)
    - [UpdateSecurityLoggingObjectResponse](#github.com.Juniper.contrail.pkg.models.UpdateSecurityLoggingObjectResponse)
    - [UpdateServerRequest](#github.com.Juniper.contrail.pkg.models.UpdateServerRequest)
    - [UpdateServerResponse](#github.com.Juniper.contrail.pkg.models.UpdateServerResponse)
    - [UpdateServiceApplianceRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceRequest)
    - [UpdateServiceApplianceResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceResponse)
    - [UpdateServiceApplianceSetRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceSetRequest)
    - [UpdateServiceApplianceSetResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceSetResponse)
    - [UpdateServiceConnectionModuleRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceConnectionModuleRequest)
    - [UpdateServiceConnectionModuleResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceConnectionModuleResponse)
    - [UpdateServiceEndpointRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceEndpointRequest)
    - [UpdateServiceEndpointResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceEndpointResponse)
    - [UpdateServiceGroupRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceGroupRequest)
    - [UpdateServiceGroupResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceGroupResponse)
    - [UpdateServiceHealthCheckRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceHealthCheckRequest)
    - [UpdateServiceHealthCheckResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceHealthCheckResponse)
    - [UpdateServiceInstanceRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceInstanceRequest)
    - [UpdateServiceInstanceResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceInstanceResponse)
    - [UpdateServiceObjectRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceObjectRequest)
    - [UpdateServiceObjectResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceObjectResponse)
    - [UpdateServiceTemplateRequest](#github.com.Juniper.contrail.pkg.models.UpdateServiceTemplateRequest)
    - [UpdateServiceTemplateResponse](#github.com.Juniper.contrail.pkg.models.UpdateServiceTemplateResponse)
    - [UpdateSubnetRequest](#github.com.Juniper.contrail.pkg.models.UpdateSubnetRequest)
    - [UpdateSubnetResponse](#github.com.Juniper.contrail.pkg.models.UpdateSubnetResponse)
    - [UpdateTagRequest](#github.com.Juniper.contrail.pkg.models.UpdateTagRequest)
    - [UpdateTagResponse](#github.com.Juniper.contrail.pkg.models.UpdateTagResponse)
    - [UpdateTagTypeRequest](#github.com.Juniper.contrail.pkg.models.UpdateTagTypeRequest)
    - [UpdateTagTypeResponse](#github.com.Juniper.contrail.pkg.models.UpdateTagTypeResponse)
    - [UpdateUserRequest](#github.com.Juniper.contrail.pkg.models.UpdateUserRequest)
    - [UpdateUserResponse](#github.com.Juniper.contrail.pkg.models.UpdateUserResponse)
    - [UpdateVPNGroupRequest](#github.com.Juniper.contrail.pkg.models.UpdateVPNGroupRequest)
    - [UpdateVPNGroupResponse](#github.com.Juniper.contrail.pkg.models.UpdateVPNGroupResponse)
    - [UpdateVirtualDNSRecordRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSRecordRequest)
    - [UpdateVirtualDNSRecordResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSRecordResponse)
    - [UpdateVirtualDNSRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSRequest)
    - [UpdateVirtualDNSResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSResponse)
    - [UpdateVirtualIPRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualIPRequest)
    - [UpdateVirtualIPResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualIPResponse)
    - [UpdateVirtualMachineInterfaceRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineInterfaceRequest)
    - [UpdateVirtualMachineInterfaceResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineInterfaceResponse)
    - [UpdateVirtualMachineRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineRequest)
    - [UpdateVirtualMachineResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineResponse)
    - [UpdateVirtualNetworkRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualNetworkRequest)
    - [UpdateVirtualNetworkResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualNetworkResponse)
    - [UpdateVirtualRouterRequest](#github.com.Juniper.contrail.pkg.models.UpdateVirtualRouterRequest)
    - [UpdateVirtualRouterResponse](#github.com.Juniper.contrail.pkg.models.UpdateVirtualRouterResponse)
    - [UpdateWidgetRequest](#github.com.Juniper.contrail.pkg.models.UpdateWidgetRequest)
    - [UpdateWidgetResponse](#github.com.Juniper.contrail.pkg.models.UpdateWidgetResponse)
    - [User](#github.com.Juniper.contrail.pkg.models.User)
    - [UserCredentials](#github.com.Juniper.contrail.pkg.models.UserCredentials)
    - [UserDefinedLogStat](#github.com.Juniper.contrail.pkg.models.UserDefinedLogStat)
    - [UserDefinedLogStatList](#github.com.Juniper.contrail.pkg.models.UserDefinedLogStatList)
    - [UveKeysType](#github.com.Juniper.contrail.pkg.models.UveKeysType)
    - [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup)
    - [VPNGroupLocationRef](#github.com.Juniper.contrail.pkg.models.VPNGroupLocationRef)
    - [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS)
    - [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord)
    - [VirtualDnsRecordType](#github.com.Juniper.contrail.pkg.models.VirtualDnsRecordType)
    - [VirtualDnsType](#github.com.Juniper.contrail.pkg.models.VirtualDnsType)
    - [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP)
    - [VirtualIPLoadbalancerPoolRef](#github.com.Juniper.contrail.pkg.models.VirtualIPLoadbalancerPoolRef)
    - [VirtualIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.VirtualIPVirtualMachineInterfaceRef)
    - [VirtualIpType](#github.com.Juniper.contrail.pkg.models.VirtualIpType)
    - [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine)
    - [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface)
    - [VirtualMachineInterfaceBGPRouterRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceBGPRouterRef)
    - [VirtualMachineInterfaceBridgeDomainRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceBridgeDomainRef)
    - [VirtualMachineInterfaceInterfaceRouteTableRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceInterfaceRouteTableRef)
    - [VirtualMachineInterfacePhysicalInterfaceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePhysicalInterfaceRef)
    - [VirtualMachineInterfacePortTupleRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePortTupleRef)
    - [VirtualMachineInterfacePropertiesType](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePropertiesType)
    - [VirtualMachineInterfaceQosConfigRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceQosConfigRef)
    - [VirtualMachineInterfaceRoutingInstanceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceRoutingInstanceRef)
    - [VirtualMachineInterfaceSecurityGroupRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceSecurityGroupRef)
    - [VirtualMachineInterfaceSecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceSecurityLoggingObjectRef)
    - [VirtualMachineInterfaceServiceEndpointRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceServiceEndpointRef)
    - [VirtualMachineInterfaceServiceHealthCheckRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceServiceHealthCheckRef)
    - [VirtualMachineInterfaceVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualMachineInterfaceRef)
    - [VirtualMachineInterfaceVirtualMachineRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualMachineRef)
    - [VirtualMachineInterfaceVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualNetworkRef)
    - [VirtualMachineServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineServiceInstanceRef)
    - [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork)
    - [VirtualNetworkBGPVPNRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkBGPVPNRef)
    - [VirtualNetworkNetworkIpamRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkNetworkIpamRef)
    - [VirtualNetworkNetworkPolicyRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkNetworkPolicyRef)
    - [VirtualNetworkPolicyType](#github.com.Juniper.contrail.pkg.models.VirtualNetworkPolicyType)
    - [VirtualNetworkQosConfigRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkQosConfigRef)
    - [VirtualNetworkRouteTableRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkRouteTableRef)
    - [VirtualNetworkSecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkSecurityLoggingObjectRef)
    - [VirtualNetworkType](#github.com.Juniper.contrail.pkg.models.VirtualNetworkType)
    - [VirtualNetworkVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkVirtualNetworkRef)
    - [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter)
    - [VirtualRouterNetworkIpamRef](#github.com.Juniper.contrail.pkg.models.VirtualRouterNetworkIpamRef)
    - [VirtualRouterNetworkIpamType](#github.com.Juniper.contrail.pkg.models.VirtualRouterNetworkIpamType)
    - [VirtualRouterVirtualMachineRef](#github.com.Juniper.contrail.pkg.models.VirtualRouterVirtualMachineRef)
    - [VnSubnetsType](#github.com.Juniper.contrail.pkg.models.VnSubnetsType)
    - [VrfAssignRuleType](#github.com.Juniper.contrail.pkg.models.VrfAssignRuleType)
    - [VrfAssignTableType](#github.com.Juniper.contrail.pkg.models.VrfAssignTableType)
    - [Widget](#github.com.Juniper.contrail.pkg.models.Widget)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="github.com/Juniper/contrail/pkg/models/generated.proto"/>
<p align="right"><a href="#top">Top</a></p>

## github.com/Juniper/contrail/pkg/models/generated.proto



<a name="github.com.Juniper.contrail.pkg.models.APIAccessList"/>

### APIAccessList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| api_access_list_entries | [RbacRuleEntriesType](#github.com.Juniper.contrail.pkg.models.RbacRuleEntriesType) |  | List of rules e.g network.* =&amp;gt; admin:CRUD (admin can perform all ops on networks). |






<a name="github.com.Juniper.contrail.pkg.models.AccessControlList"/>

### AccessControlList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| access_control_list_hash | [int64](#int64) |  | A hash value of all the access-control-list-entries in this ACL objects automatically generated by system. |
| access_control_list_entries | [AclEntriesType](#github.com.Juniper.contrail.pkg.models.AclEntriesType) |  | Automatically generated by system based on security groups or network policies. |






<a name="github.com.Juniper.contrail.pkg.models.AclEntriesType"/>

### AclEntriesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic | [bool](#bool) |  | For Internal use only |
| acl_rule | [AclRuleType](#github.com.Juniper.contrail.pkg.models.AclRuleType) | repeated | For Internal use only |






<a name="github.com.Juniper.contrail.pkg.models.AclRuleType"/>

### AclRuleType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule_uuid | [string](#string) |  | Rule UUID is identifier used in flow records to identify rule |
| match_condition | [MatchConditionType](#github.com.Juniper.contrail.pkg.models.MatchConditionType) |  | Match condition for packets |
| direction | [string](#string) |  | Direction in the rule |
| action_list | [ActionListType](#github.com.Juniper.contrail.pkg.models.ActionListType) |  | Actions to be performed if packets match condition |






<a name="github.com.Juniper.contrail.pkg.models.ActionListType"/>

### ActionListType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gateway_name | [string](#string) |  | For internal use only |
| log | [bool](#bool) |  | Flow records for traffic matching this rule are sent at higher priority |
| alert | [bool](#bool) |  | For internal use only |
| qos_action | [string](#string) |  | FQN of Qos configuration object for QoS marking |
| assign_routing_instance | [string](#string) |  | For internal use only |
| mirror_to | [MirrorActionType](#github.com.Juniper.contrail.pkg.models.MirrorActionType) |  | Mirror traffic matching this rule |
| simple_action | [string](#string) |  | Simple allow(pass) or deny action for traffic matching this rule |
| apply_service | [string](#string) | repeated | Ordered list of service instances in service chain applied to traffic matching the rule |






<a name="github.com.Juniper.contrail.pkg.models.AddressGroup"/>

### AddressGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| address_group_prefix | [SubnetListType](#github.com.Juniper.contrail.pkg.models.SubnetListType) |  | List of IP prefix |






<a name="github.com.Juniper.contrail.pkg.models.AddressType"/>

### AddressType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_group | [string](#string) |  | Any address that belongs to interface with this security-group |
| subnet | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | Any address that belongs to this subnet |
| network_policy | [string](#string) |  | Any address that belongs to virtual network which has this policy attached |
| subnet_list | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) | repeated | Any address that belongs to any one of subnet in this list |
| virtual_network | [string](#string) |  | Any address that belongs to this virtual network |






<a name="github.com.Juniper.contrail.pkg.models.Alarm"/>

### Alarm



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| alarm_rules | [AlarmOrList](#github.com.Juniper.contrail.pkg.models.AlarmOrList) |  | Rules based on the UVE attributes specified as OR-of-ANDs of AlarmExpression template. Example: &amp;quot;alarm_rules&amp;quot;: {&amp;quot;or_list&amp;quot;: [{&amp;quot;and_list&amp;quot;: [{AlarmExpression1}, {AlarmExpression2}, ...]}, {&amp;quot;and_list&amp;quot;: [{AlarmExpression3}, {AlarmExpression4}, ...]}]} |
| uve_keys | [UveKeysType](#github.com.Juniper.contrail.pkg.models.UveKeysType) |  | List of UVE tables or UVE objects where this alarm config should be applied. For example, rules based on NodeStatus UVE can be applied to multiple object types or specific uve objects such as analytics-node, config-node, control-node:&amp;lt;hostname&amp;gt;, etc., |
| alarm_severity | [int64](#int64) |  | Severity level for the alarm. |






<a name="github.com.Juniper.contrail.pkg.models.AlarmAndList"/>

### AlarmAndList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| and_list | [AlarmExpression](#github.com.Juniper.contrail.pkg.models.AlarmExpression) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.AlarmExpression"/>

### AlarmExpression



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operations | [string](#string) |  | operation to compare operand1 and operand2 |
| operand1 | [string](#string) |  | UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state |
| variables | [string](#string) | repeated | List of UVE attributes that would be useful when the alarm is raised. For example, user may want to raise an alarm if the NodeStatus.process_info.process_state != PROCESS_STATE_RUNNING. But, it would be useful to know the process_name whose state != PROCESS_STATE_RUNNING. This UVE attribute which is neither part of operand1 nor operand2 may be specified in variables |
| operand2 | [AlarmOperand2](#github.com.Juniper.contrail.pkg.models.AlarmOperand2) |  | UVE attribute or a json value to compare with the UVE attribute in operand1 |






<a name="github.com.Juniper.contrail.pkg.models.AlarmOperand2"/>

### AlarmOperand2



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uve_attribute | [string](#string) |  | UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state |
| json_value | [string](#string) |  | json value as string |






<a name="github.com.Juniper.contrail.pkg.models.AlarmOrList"/>

### AlarmOrList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| or_list | [AlarmAndList](#github.com.Juniper.contrail.pkg.models.AlarmAndList) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.AliasIP"/>

### AliasIP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| alias_ip_address | [string](#string) |  | Alias ip address. |
| alias_ip_address_family | [string](#string) |  | Ip address family of the alias ip, IpV4 or IpV6 |
| project_refs | [AliasIPProjectRef](#github.com.Juniper.contrail.pkg.models.AliasIPProjectRef) | repeated | Reference to project from which this alias ip was allocated. |
| virtual_machine_interface_refs | [AliasIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.AliasIPVirtualMachineInterfaceRef) | repeated | Reference to virtual machine interface to which this alias ip is attached. |






<a name="github.com.Juniper.contrail.pkg.models.AliasIPPool"/>

### AliasIPPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| alias_ips | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) | repeated | alias ip is a ip that can be assigned to virtual-machine-interface(VMI), By doing so VMI can now be part of the alias ip network. packets originating with alias-ip as the source-ip belongs to alias-ip-network |






<a name="github.com.Juniper.contrail.pkg.models.AliasIPProjectRef"/>

### AliasIPProjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.AliasIPVirtualMachineInterfaceRef"/>

### AliasIPVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.AllocationPoolType"/>

### AllocationPoolType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vrouter_specific_pool | [bool](#bool) |  |  |
| start | [string](#string) |  |  |
| end | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.AllowedAddressPair"/>

### AllowedAddressPair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  |  |
| mac | [string](#string) |  |  |
| address_mode | [string](#string) |  | Address-mode active-backup is used for VRRP address. Address-mode active-active is used for ECMP. |






<a name="github.com.Juniper.contrail.pkg.models.AllowedAddressPairs"/>

### AllowedAddressPairs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowed_address_pair | [AllowedAddressPair](#github.com.Juniper.contrail.pkg.models.AllowedAddressPair) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.AnalyticsNode"/>

### AnalyticsNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| analytics_node_ip_address | [string](#string) |  | Ip address of the analytics node, set while provisioning. |






<a name="github.com.Juniper.contrail.pkg.models.AppformixNode"/>

### AppformixNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [AppformixNodeNodeRef](#github.com.Juniper.contrail.pkg.models.AppformixNodeNodeRef) | repeated | Reference to node object for this appformix node. |






<a name="github.com.Juniper.contrail.pkg.models.AppformixNodeNodeRef"/>

### AppformixNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ApplicationPolicySet"/>

### ApplicationPolicySet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| all_applications | [bool](#bool) |  | If set, indicates application policy set to be applied to all application tags |
| firewall_policy_refs | [ApplicationPolicySetFirewallPolicyRef](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySetFirewallPolicyRef) | repeated | Reference to firewall-policy attached to this application-policy |
| global_vrouter_config_refs | [ApplicationPolicySetGlobalVrouterConfigRef](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySetGlobalVrouterConfigRef) | repeated | Reference to global-vrouter-config is automatically created by system for global application policy sets |






<a name="github.com.Juniper.contrail.pkg.models.ApplicationPolicySetFirewallPolicyRef"/>

### ApplicationPolicySetFirewallPolicyRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [FirewallSequence](#github.com.Juniper.contrail.pkg.models.FirewallSequence) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ApplicationPolicySetGlobalVrouterConfigRef"/>

### ApplicationPolicySetGlobalVrouterConfigRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.BGPAsAService"/>

### BGPAsAService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| bgpaas_shared | [bool](#bool) |  | True if only one BGP router needs to be created. Otherwise, one BGP router is created for each VMI |
| bgpaas_session_attributes | [string](#string) |  | BGP peering session attributes. |
| bgpaas_suppress_route_advertisement | [bool](#bool) |  | True when server should not advertise any routes to the client i.e. the client has static routes (typically a default) configured. |
| bgpaas_ipv4_mapped_ipv6_nexthop | [bool](#bool) |  | True when client bgp implementation expects to receive a ipv4-mapped ipv6 address (as opposed to regular ipv6 address) as the bgp nexthop for ipv6 routes. |
| bgpaas_ip_address | [string](#string) |  | Ip address of the BGP peer. |
| autonomous_system | [int64](#int64) |  | 16 bit BGP Autonomous System number for the cluster. |
| virtual_machine_interface_refs | [BGPAsAServiceVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.BGPAsAServiceVirtualMachineInterfaceRef) | repeated | Reference to VMI on which BGPaaS BGP peering will happen. |
| service_health_check_refs | [BGPAsAServiceServiceHealthCheckRef](#github.com.Juniper.contrail.pkg.models.BGPAsAServiceServiceHealthCheckRef) | repeated | Reference to health check object attached to BGPaaS object, used to enable BFD health check over active BGPaaS VMI. |






<a name="github.com.Juniper.contrail.pkg.models.BGPAsAServiceServiceHealthCheckRef"/>

### BGPAsAServiceServiceHealthCheckRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.BGPAsAServiceVirtualMachineInterfaceRef"/>

### BGPAsAServiceVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.BGPRouter"/>

### BGPRouter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |






<a name="github.com.Juniper.contrail.pkg.models.BGPVPN"/>

### BGPVPN



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that are used as both import and export for this virtual network. |
| import_route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that are used as import for this virtual network. |
| export_route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that are used as export for this virtual network. |
| bgpvpn_type | [string](#string) |  | BGP VPN type selection between IP VPN (l3) and Ethernet VPN (l2) (default: l3). |






<a name="github.com.Juniper.contrail.pkg.models.BGPaaServiceParametersType"/>

### BGPaaServiceParametersType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_start | [int64](#int64) |  |  |
| port_end | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.BaremetalNode"/>

### BaremetalNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| name | [string](#string) |  | Name of the nodename to easily identify Baremetal |
| driver_info | [DriverInfo](#github.com.Juniper.contrail.pkg.models.DriverInfo) |  | Details of the driver for power management |
| bm_properties | [BaremetalProperties](#github.com.Juniper.contrail.pkg.models.BaremetalProperties) |  | Details of baremetal hardware for scheduler |
| instance_uuid | [string](#string) |  | UUID of the Nova instance associated with this Node |
| instance_info | [InstanceInfo](#github.com.Juniper.contrail.pkg.models.InstanceInfo) |  | Details of Instance launched on the baremetal |
| maintenance | [bool](#bool) |  | Whether or not this Node is currently in maintenance mode |
| maintenance_reason | [string](#string) |  | Reason why this Node was placed into maintenance mode |
| power_state | [string](#string) |  | The current power state of this Node |
| target_power_state | [string](#string) |  | If a power state transition has been requested, this field represents the requested state |
| provision_state | [string](#string) |  | The current provisioning state of this Node |
| target_provision_state | [string](#string) |  | If a provisioning action has been requested, this field represents the requested state |
| console_enabled | [bool](#bool) |  | Indicates whether console access is enabled or disabled on this node |
| created_at | [string](#string) |  | The UTC date and time when the resource was created, ISO 8601 format |
| updated_at | [string](#string) |  | The UTC date and time when the resource was created, ISO 8601 format |
| last_error | [string](#string) |  | Any error from the most recent (last) transaction that started but failed to finish. |






<a name="github.com.Juniper.contrail.pkg.models.BaremetalPort"/>

### BaremetalPort



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| mac_address | [string](#string) |  | Mac Address of the NIC of Baremetal |
| created_at | [string](#string) |  | The UTC date and time when the resource was created, ISO 8601 format |
| updated_at | [string](#string) |  | The UTC date and time when the resource was created, ISO 8601 format |
| node | [string](#string) |  | UUID of the node where this port is connected |
| pxe_enabled | [bool](#bool) |  | Indicates whether PXE is enabled or disabled on the Port. |
| local_link_connection | [LocalLinkConnection](#github.com.Juniper.contrail.pkg.models.LocalLinkConnection) |  | The Port binding profile |






<a name="github.com.Juniper.contrail.pkg.models.BaremetalProperties"/>

### BaremetalProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cpu_count | [int64](#int64) |  | Number of CPU cores on the baremetal |
| cpu_arch | [string](#string) |  | Architecture of the baremetal server |
| disk_gb | [int64](#int64) |  | Disk size of root device (in GB) |
| memory_mb | [int64](#int64) |  | RAM of the Baremetal server (in MB) |






<a name="github.com.Juniper.contrail.pkg.models.BridgeDomain"/>

### BridgeDomain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| mac_aging_time | [int64](#int64) |  | MAC aging time on the network |
| isid | [int64](#int64) |  | i-sid value |
| mac_learning_enabled | [bool](#bool) |  | Enable MAC learning on the network |
| mac_move_control | [MACMoveLimitControlType](#github.com.Juniper.contrail.pkg.models.MACMoveLimitControlType) |  | MAC move control on the network |
| mac_limit_control | [MACLimitControlType](#github.com.Juniper.contrail.pkg.models.MACLimitControlType) |  | MAC limit control on the network |






<a name="github.com.Juniper.contrail.pkg.models.BridgeDomainMembershipType"/>

### BridgeDomainMembershipType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vlan_tag | [int64](#int64) |  | VLAN tag of the incoming packet that maps the virtual-machine-interface to bridge domain |






<a name="github.com.Juniper.contrail.pkg.models.CommunityAttributes"/>

### CommunityAttributes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| community_attribute | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ConfigNode"/>

### ConfigNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| config_node_ip_address | [string](#string) |  | Ip address of the config node, set while provisioning. |






<a name="github.com.Juniper.contrail.pkg.models.ConfigRoot"/>

### ConfigRoot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| tag_refs | [ConfigRootTagRef](#github.com.Juniper.contrail.pkg.models.ConfigRootTagRef) | repeated | Attribute attached to objects - has a type and value |
| domains | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) | repeated | Domain is authentication namespace, a collection of projects. |
| global_system_configs | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) | repeated | Global system config is object where all global system configuration is present. |
| tags | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) | repeated | Attribute attached to objects - has a type and value |






<a name="github.com.Juniper.contrail.pkg.models.ConfigRootTagRef"/>

### ConfigRootTagRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode"/>

### ContrailAnalyticsDatabaseNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [ContrailAnalyticsDatabaseNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNodeNodeRef) | repeated | Reference to node object for this analyticsdb node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNodeNodeRef"/>

### ContrailAnalyticsDatabaseNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode"/>

### ContrailAnalyticsNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [ContrailAnalyticsNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNodeNodeRef) | repeated | Reference to node object for this analytics node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNodeNodeRef"/>

### ContrailAnalyticsNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailCluster"/>

### ContrailCluster



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| container_registry | [string](#string) |  | Container registry to pull contrail containers |
| contrail_version | [string](#string) |  | Version of contrail containers to be used. |
| rabbitmq_port | [string](#string) |  | RabbitMQ port used to connect to RabbitMQ. |
| provisioner_type | [string](#string) |  | Tool to be used to provision the cluster |
| orchestrator | [string](#string) |  | Orchestrator to use with contrrail |
| config_audit_ttl | [string](#string) |  | Configuration Audit Retention Time in hours |
| default_gateway | [string](#string) |  | Default Gateway |
| default_vrouter_bond_interface | [string](#string) |  | vRouter Bond Interface |
| default_vrouter_bond_interface_members | [string](#string) |  | vRouter Bond Interface Members |
| flow_ttl | [string](#string) |  | Flow Data Retention Time in hours |
| statistics_ttl | [string](#string) |  | Statistics Data Retention Time in hours |
| openstack_internal_vip | [string](#string) |  | VIP for openstack internal network |
| openstack_external_vip | [string](#string) |  | VIP for openstack external network |
| openstack_internal_vip_interface | [string](#string) |  | VIP for openstack internal network |
| openstack_external_vip_interface | [string](#string) |  | Interface to configure VIP for openstack external network |
| openstack_enable_haproxy | [string](#string) |  | Configure haproxy in openstack control node |
| appformix_nodes | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) | repeated | Parent of this appformix node. |
| contrail_analytics_database_nodes | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) | repeated | Parent of this analyticsdb node. |
| contrail_analytics_nodes | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) | repeated | Parent of this analytics node. |
| contrail_config_database_nodes | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) | repeated | Parent of this configdb node. |
| contrail_config_nodes | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) | repeated | Parent of this config node. |
| contrail_control_nodes | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) | repeated | Parent of this control node. |
| contrail_storage_nodes | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) | repeated | Parent of this storage node. |
| contrail_vrouter_nodes | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) | repeated | Parent of this vrouter node. |
| contrail_webui_nodes | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) | repeated | Parent of this webui node. |
| kubernetes_master_nodes | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) | repeated | Parent of this kubernetes master node. |
| kubernetes_nodes | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) | repeated | Parent of this kubernetes node. |
| openstack_compute_nodes | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) | repeated | Parent of this openstack_compute node. |
| openstack_control_nodes | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) | repeated | Parent of this openstack_control node. |
| openstack_monitoring_nodes | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) | repeated | Parent of this openstack_monitoring node. |
| openstack_network_nodes | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) | repeated | Parent of this openstack_network node. |
| openstack_storage_nodes | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) | repeated | Parent of this openstack_storage node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode"/>

### ContrailConfigDatabaseNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [ContrailConfigDatabaseNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNodeNodeRef) | repeated | Reference to node object for this configdb node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNodeNodeRef"/>

### ContrailConfigDatabaseNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailConfigNode"/>

### ContrailConfigNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [ContrailConfigNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailConfigNodeNodeRef) | repeated | Reference to node object for this config node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailConfigNodeNodeRef"/>

### ContrailConfigNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailControlNode"/>

### ContrailControlNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [ContrailControlNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailControlNodeNodeRef) | repeated | Reference to node object for this control node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailControlNodeNodeRef"/>

### ContrailControlNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailStorageNode"/>

### ContrailStorageNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| journal_drives | [string](#string) |  | SSD Drives to use for journals |
| osd_drives | [string](#string) |  | Drives to use for cloud storage |
| storage_access_bond_interface_members | [string](#string) |  | Storage Management Bond Interface Members |
| storage_backend_bond_interface_members | [string](#string) |  | Storage Backend Bond Interface Members |
| node_refs | [ContrailStorageNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailStorageNodeNodeRef) | repeated | Reference to node object for this storage node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailStorageNodeNodeRef"/>

### ContrailStorageNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailVrouterNode"/>

### ContrailVrouterNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| default_gateway | [string](#string) |  | Default Gateway |
| vrouter_bond_interface | [string](#string) |  | vRouter Bond Interface |
| vrouter_bond_interface_members | [string](#string) |  | vRouter Bond Interface Members |
| vrouter_type | [string](#string) |  |  |
| node_refs | [ContrailVrouterNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNodeNodeRef) | repeated | Reference to node object for this vrouter node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailVrouterNodeNodeRef"/>

### ContrailVrouterNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ContrailWebuiNode"/>

### ContrailWebuiNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [ContrailWebuiNodeNodeRef](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNodeNodeRef) | repeated | Reference to node object for this webui node. |






<a name="github.com.Juniper.contrail.pkg.models.ContrailWebuiNodeNodeRef"/>

### ContrailWebuiNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ControlTrafficDscpType"/>

### ControlTrafficDscpType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| control | [int64](#int64) |  | DSCP value for control protocols traffic |
| analytics | [int64](#int64) |  | DSCP value for traffic towards analytics |
| dns | [int64](#int64) |  | DSCP value for DNS traffic |






<a name="github.com.Juniper.contrail.pkg.models.CreateAPIAccessListRequest"/>

### CreateAPIAccessListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_access_list | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAPIAccessListResponse"/>

### CreateAPIAccessListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_access_list | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAccessControlListRequest"/>

### CreateAccessControlListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_control_list | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAccessControlListResponse"/>

### CreateAccessControlListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_control_list | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAddressGroupRequest"/>

### CreateAddressGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_group | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAddressGroupResponse"/>

### CreateAddressGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_group | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAlarmRequest"/>

### CreateAlarmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alarm | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAlarmResponse"/>

### CreateAlarmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alarm | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAliasIPPoolRequest"/>

### CreateAliasIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip_pool | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAliasIPPoolResponse"/>

### CreateAliasIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip_pool | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAliasIPRequest"/>

### CreateAliasIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAliasIPResponse"/>

### CreateAliasIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAnalyticsNodeRequest"/>

### CreateAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analytics_node | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAnalyticsNodeResponse"/>

### CreateAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analytics_node | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAppformixNodeRequest"/>

### CreateAppformixNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appformix_node | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateAppformixNodeResponse"/>

### CreateAppformixNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appformix_node | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateApplicationPolicySetRequest"/>

### CreateApplicationPolicySetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_policy_set | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateApplicationPolicySetResponse"/>

### CreateApplicationPolicySetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_policy_set | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBGPAsAServiceRequest"/>

### CreateBGPAsAServiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_as_a_service | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBGPAsAServiceResponse"/>

### CreateBGPAsAServiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_as_a_service | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBGPRouterRequest"/>

### CreateBGPRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_router | [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBGPRouterResponse"/>

### CreateBGPRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_router | [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBGPVPNRequest"/>

### CreateBGPVPNRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgpvpn | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBGPVPNResponse"/>

### CreateBGPVPNResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgpvpn | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBaremetalNodeRequest"/>

### CreateBaremetalNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_node | [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBaremetalNodeResponse"/>

### CreateBaremetalNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_node | [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBaremetalPortRequest"/>

### CreateBaremetalPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_port | [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBaremetalPortResponse"/>

### CreateBaremetalPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_port | [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBridgeDomainRequest"/>

### CreateBridgeDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bridge_domain | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateBridgeDomainResponse"/>

### CreateBridgeDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bridge_domain | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateConfigNodeRequest"/>

### CreateConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_node | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateConfigNodeResponse"/>

### CreateConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_node | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateConfigRootRequest"/>

### CreateConfigRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_root | [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateConfigRootResponse"/>

### CreateConfigRootResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_root | [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsDatabaseNodeRequest"/>

### CreateContrailAnalyticsDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_database_node | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsDatabaseNodeResponse"/>

### CreateContrailAnalyticsDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_database_node | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsNodeRequest"/>

### CreateContrailAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_node | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailAnalyticsNodeResponse"/>

### CreateContrailAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_node | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailClusterRequest"/>

### CreateContrailClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_cluster | [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailClusterResponse"/>

### CreateContrailClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_cluster | [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailConfigDatabaseNodeRequest"/>

### CreateContrailConfigDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_database_node | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailConfigDatabaseNodeResponse"/>

### CreateContrailConfigDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_database_node | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailConfigNodeRequest"/>

### CreateContrailConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_node | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailConfigNodeResponse"/>

### CreateContrailConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_node | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailControlNodeRequest"/>

### CreateContrailControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_control_node | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailControlNodeResponse"/>

### CreateContrailControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_control_node | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailStorageNodeRequest"/>

### CreateContrailStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_storage_node | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailStorageNodeResponse"/>

### CreateContrailStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_storage_node | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailVrouterNodeRequest"/>

### CreateContrailVrouterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_vrouter_node | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailVrouterNodeResponse"/>

### CreateContrailVrouterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_vrouter_node | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailWebuiNodeRequest"/>

### CreateContrailWebuiNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_webui_node | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateContrailWebuiNodeResponse"/>

### CreateContrailWebuiNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_webui_node | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateCustomerAttachmentRequest"/>

### CreateCustomerAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_attachment | [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateCustomerAttachmentResponse"/>

### CreateCustomerAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_attachment | [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDashboardRequest"/>

### CreateDashboardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dashboard | [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDashboardResponse"/>

### CreateDashboardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dashboard | [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDatabaseNodeRequest"/>

### CreateDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_node | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDatabaseNodeResponse"/>

### CreateDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_node | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDiscoveryServiceAssignmentRequest"/>

### CreateDiscoveryServiceAssignmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| discovery_service_assignment | [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDiscoveryServiceAssignmentResponse"/>

### CreateDiscoveryServiceAssignmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| discovery_service_assignment | [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDomainRequest"/>

### CreateDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDomainResponse"/>

### CreateDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDsaRuleRequest"/>

### CreateDsaRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dsa_rule | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateDsaRuleResponse"/>

### CreateDsaRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dsa_rule | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateE2ServiceProviderRequest"/>

### CreateE2ServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_service_provider | [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateE2ServiceProviderResponse"/>

### CreateE2ServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_service_provider | [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFirewallPolicyRequest"/>

### CreateFirewallPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_policy | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFirewallPolicyResponse"/>

### CreateFirewallPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_policy | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFirewallRuleRequest"/>

### CreateFirewallRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_rule | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFirewallRuleResponse"/>

### CreateFirewallRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_rule | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFlavorRequest"/>

### CreateFlavorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flavor | [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFlavorResponse"/>

### CreateFlavorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flavor | [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFloatingIPPoolRequest"/>

### CreateFloatingIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_pool | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFloatingIPPoolResponse"/>

### CreateFloatingIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_pool | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFloatingIPRequest"/>

### CreateFloatingIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateFloatingIPResponse"/>

### CreateFloatingIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateForwardingClassRequest"/>

### CreateForwardingClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_class | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateForwardingClassResponse"/>

### CreateForwardingClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_class | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateGlobalQosConfigRequest"/>

### CreateGlobalQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_qos_config | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateGlobalQosConfigResponse"/>

### CreateGlobalQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_qos_config | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateGlobalSystemConfigRequest"/>

### CreateGlobalSystemConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_system_config | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateGlobalSystemConfigResponse"/>

### CreateGlobalSystemConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_system_config | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateGlobalVrouterConfigRequest"/>

### CreateGlobalVrouterConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_vrouter_config | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateGlobalVrouterConfigResponse"/>

### CreateGlobalVrouterConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_vrouter_config | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateInstanceIPRequest"/>

### CreateInstanceIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_ip | [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateInstanceIPResponse"/>

### CreateInstanceIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_ip | [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateInterfaceRouteTableRequest"/>

### CreateInterfaceRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_route_table | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateInterfaceRouteTableResponse"/>

### CreateInterfaceRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_route_table | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateKeypairRequest"/>

### CreateKeypairRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keypair | [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateKeypairResponse"/>

### CreateKeypairResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keypair | [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateKubernetesMasterNodeRequest"/>

### CreateKubernetesMasterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_master_node | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateKubernetesMasterNodeResponse"/>

### CreateKubernetesMasterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_master_node | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateKubernetesNodeRequest"/>

### CreateKubernetesNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_node | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateKubernetesNodeResponse"/>

### CreateKubernetesNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_node | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerHealthmonitorRequest"/>

### CreateLoadbalancerHealthmonitorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_healthmonitor | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerHealthmonitorResponse"/>

### CreateLoadbalancerHealthmonitorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_healthmonitor | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerListenerRequest"/>

### CreateLoadbalancerListenerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_listener | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerListenerResponse"/>

### CreateLoadbalancerListenerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_listener | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerMemberRequest"/>

### CreateLoadbalancerMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_member | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerMemberResponse"/>

### CreateLoadbalancerMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_member | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerPoolRequest"/>

### CreateLoadbalancerPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_pool | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerPoolResponse"/>

### CreateLoadbalancerPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_pool | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerRequest"/>

### CreateLoadbalancerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLoadbalancerResponse"/>

### CreateLoadbalancerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLocationRequest"/>

### CreateLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [Location](#github.com.Juniper.contrail.pkg.models.Location) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLocationResponse"/>

### CreateLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [Location](#github.com.Juniper.contrail.pkg.models.Location) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLogicalInterfaceRequest"/>

### CreateLogicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_interface | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLogicalInterfaceResponse"/>

### CreateLogicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_interface | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLogicalRouterRequest"/>

### CreateLogicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_router | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateLogicalRouterResponse"/>

### CreateLogicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_router | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNamespaceRequest"/>

### CreateNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNamespaceResponse"/>

### CreateNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNetworkDeviceConfigRequest"/>

### CreateNetworkDeviceConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_device_config | [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNetworkDeviceConfigResponse"/>

### CreateNetworkDeviceConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_device_config | [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNetworkIpamRequest"/>

### CreateNetworkIpamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_ipam | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNetworkIpamResponse"/>

### CreateNetworkIpamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_ipam | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNetworkPolicyRequest"/>

### CreateNetworkPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_policy | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNetworkPolicyResponse"/>

### CreateNetworkPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_policy | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNodeRequest"/>

### CreateNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#github.com.Juniper.contrail.pkg.models.Node) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateNodeResponse"/>

### CreateNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#github.com.Juniper.contrail.pkg.models.Node) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackComputeNodeRequest"/>

### CreateOpenstackComputeNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_compute_node | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackComputeNodeResponse"/>

### CreateOpenstackComputeNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_compute_node | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackControlNodeRequest"/>

### CreateOpenstackControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_control_node | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackControlNodeResponse"/>

### CreateOpenstackControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_control_node | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackMonitoringNodeRequest"/>

### CreateOpenstackMonitoringNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_monitoring_node | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackMonitoringNodeResponse"/>

### CreateOpenstackMonitoringNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_monitoring_node | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackNetworkNodeRequest"/>

### CreateOpenstackNetworkNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_network_node | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackNetworkNodeResponse"/>

### CreateOpenstackNetworkNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_network_node | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackStorageNodeRequest"/>

### CreateOpenstackStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_storage_node | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOpenstackStorageNodeResponse"/>

### CreateOpenstackStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_storage_node | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOsImageRequest"/>

### CreateOsImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| os_image | [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateOsImageResponse"/>

### CreateOsImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| os_image | [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePeeringPolicyRequest"/>

### CreatePeeringPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peering_policy | [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePeeringPolicyResponse"/>

### CreatePeeringPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peering_policy | [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePhysicalInterfaceRequest"/>

### CreatePhysicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_interface | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePhysicalInterfaceResponse"/>

### CreatePhysicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_interface | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePhysicalRouterRequest"/>

### CreatePhysicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_router | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePhysicalRouterResponse"/>

### CreatePhysicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_router | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePolicyManagementRequest"/>

### CreatePolicyManagementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_management | [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePolicyManagementResponse"/>

### CreatePolicyManagementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_management | [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePortRequest"/>

### CreatePortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [Port](#github.com.Juniper.contrail.pkg.models.Port) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePortResponse"/>

### CreatePortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [Port](#github.com.Juniper.contrail.pkg.models.Port) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePortTupleRequest"/>

### CreatePortTupleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_tuple | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreatePortTupleResponse"/>

### CreatePortTupleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_tuple | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateProjectRequest"/>

### CreateProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#github.com.Juniper.contrail.pkg.models.Project) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateProjectResponse"/>

### CreateProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#github.com.Juniper.contrail.pkg.models.Project) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateProviderAttachmentRequest"/>

### CreateProviderAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_attachment | [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateProviderAttachmentResponse"/>

### CreateProviderAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_attachment | [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateQosConfigRequest"/>

### CreateQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_config | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateQosConfigResponse"/>

### CreateQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_config | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateQosQueueRequest"/>

### CreateQosQueueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_queue | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateQosQueueResponse"/>

### CreateQosQueueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_queue | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRouteAggregateRequest"/>

### CreateRouteAggregateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_aggregate | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRouteAggregateResponse"/>

### CreateRouteAggregateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_aggregate | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRouteTableRequest"/>

### CreateRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_table | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRouteTableResponse"/>

### CreateRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_table | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRouteTargetRequest"/>

### CreateRouteTargetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_target | [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRouteTargetResponse"/>

### CreateRouteTargetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_target | [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRoutingInstanceRequest"/>

### CreateRoutingInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instance | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRoutingInstanceResponse"/>

### CreateRoutingInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instance | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRoutingPolicyRequest"/>

### CreateRoutingPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_policy | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateRoutingPolicyResponse"/>

### CreateRoutingPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_policy | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateSecurityGroupRequest"/>

### CreateSecurityGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_group | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateSecurityGroupResponse"/>

### CreateSecurityGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_group | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateSecurityLoggingObjectRequest"/>

### CreateSecurityLoggingObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_logging_object | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateSecurityLoggingObjectResponse"/>

### CreateSecurityLoggingObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_logging_object | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServerRequest"/>

### CreateServerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [Server](#github.com.Juniper.contrail.pkg.models.Server) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServerResponse"/>

### CreateServerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [Server](#github.com.Juniper.contrail.pkg.models.Server) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceApplianceRequest"/>

### CreateServiceApplianceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceApplianceResponse"/>

### CreateServiceApplianceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceApplianceSetRequest"/>

### CreateServiceApplianceSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance_set | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceApplianceSetResponse"/>

### CreateServiceApplianceSetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance_set | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceConnectionModuleRequest"/>

### CreateServiceConnectionModuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_connection_module | [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceConnectionModuleResponse"/>

### CreateServiceConnectionModuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_connection_module | [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceEndpointRequest"/>

### CreateServiceEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_endpoint | [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceEndpointResponse"/>

### CreateServiceEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_endpoint | [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceGroupRequest"/>

### CreateServiceGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_group | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceGroupResponse"/>

### CreateServiceGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_group | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceHealthCheckRequest"/>

### CreateServiceHealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_health_check | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceHealthCheckResponse"/>

### CreateServiceHealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_health_check | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceInstanceRequest"/>

### CreateServiceInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_instance | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceInstanceResponse"/>

### CreateServiceInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_instance | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceObjectRequest"/>

### CreateServiceObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_object | [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceObjectResponse"/>

### CreateServiceObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_object | [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceTemplateRequest"/>

### CreateServiceTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_template | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateServiceTemplateResponse"/>

### CreateServiceTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_template | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateSubnetRequest"/>

### CreateSubnetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateSubnetResponse"/>

### CreateSubnetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateTagRequest"/>

### CreateTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateTagResponse"/>

### CreateTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateTagTypeRequest"/>

### CreateTagTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_type | [TagType](#github.com.Juniper.contrail.pkg.models.TagType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateTagTypeResponse"/>

### CreateTagTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_type | [TagType](#github.com.Juniper.contrail.pkg.models.TagType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateUserRequest"/>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#github.com.Juniper.contrail.pkg.models.User) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateUserResponse"/>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#github.com.Juniper.contrail.pkg.models.User) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVPNGroupRequest"/>

### CreateVPNGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpn_group | [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVPNGroupResponse"/>

### CreateVPNGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpn_group | [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualDNSRecordRequest"/>

### CreateVirtualDNSRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS_record | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualDNSRecordResponse"/>

### CreateVirtualDNSRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS_record | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualDNSRequest"/>

### CreateVirtualDNSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualDNSResponse"/>

### CreateVirtualDNSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualIPRequest"/>

### CreateVirtualIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_ip | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualIPResponse"/>

### CreateVirtualIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_ip | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualMachineInterfaceRequest"/>

### CreateVirtualMachineInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine_interface | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualMachineInterfaceResponse"/>

### CreateVirtualMachineInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine_interface | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualMachineRequest"/>

### CreateVirtualMachineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine | [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualMachineResponse"/>

### CreateVirtualMachineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine | [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualNetworkRequest"/>

### CreateVirtualNetworkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_network | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualNetworkResponse"/>

### CreateVirtualNetworkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_network | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualRouterRequest"/>

### CreateVirtualRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_router | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateVirtualRouterResponse"/>

### CreateVirtualRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_router | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateWidgetRequest"/>

### CreateWidgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| widget | [Widget](#github.com.Juniper.contrail.pkg.models.Widget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CreateWidgetResponse"/>

### CreateWidgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| widget | [Widget](#github.com.Juniper.contrail.pkg.models.Widget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.CustomerAttachment"/>

### CustomerAttachment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_machine_interface_refs | [CustomerAttachmentVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.CustomerAttachmentVirtualMachineInterfaceRef) | repeated | Not in Use. |
| floating_ip_refs | [CustomerAttachmentFloatingIPRef](#github.com.Juniper.contrail.pkg.models.CustomerAttachmentFloatingIPRef) | repeated | Not in Use. |






<a name="github.com.Juniper.contrail.pkg.models.CustomerAttachmentFloatingIPRef"/>

### CustomerAttachmentFloatingIPRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.CustomerAttachmentVirtualMachineInterfaceRef"/>

### CustomerAttachmentVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.Dashboard"/>

### Dashboard



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| container_config | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DatabaseNode"/>

### DatabaseNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| database_node_ip_address | [string](#string) |  | Ip address of the database node, set while provisioning. |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAPIAccessListRequest"/>

### DeleteAPIAccessListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAPIAccessListResponse"/>

### DeleteAPIAccessListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAccessControlListRequest"/>

### DeleteAccessControlListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAccessControlListResponse"/>

### DeleteAccessControlListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAddressGroupRequest"/>

### DeleteAddressGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAddressGroupResponse"/>

### DeleteAddressGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAlarmRequest"/>

### DeleteAlarmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAlarmResponse"/>

### DeleteAlarmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAliasIPPoolRequest"/>

### DeleteAliasIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAliasIPPoolResponse"/>

### DeleteAliasIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAliasIPRequest"/>

### DeleteAliasIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAliasIPResponse"/>

### DeleteAliasIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAnalyticsNodeRequest"/>

### DeleteAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAnalyticsNodeResponse"/>

### DeleteAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAppformixNodeRequest"/>

### DeleteAppformixNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteAppformixNodeResponse"/>

### DeleteAppformixNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteApplicationPolicySetRequest"/>

### DeleteApplicationPolicySetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteApplicationPolicySetResponse"/>

### DeleteApplicationPolicySetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBGPAsAServiceRequest"/>

### DeleteBGPAsAServiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBGPAsAServiceResponse"/>

### DeleteBGPAsAServiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBGPRouterRequest"/>

### DeleteBGPRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBGPRouterResponse"/>

### DeleteBGPRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBGPVPNRequest"/>

### DeleteBGPVPNRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBGPVPNResponse"/>

### DeleteBGPVPNResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBaremetalNodeRequest"/>

### DeleteBaremetalNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBaremetalNodeResponse"/>

### DeleteBaremetalNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBaremetalPortRequest"/>

### DeleteBaremetalPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBaremetalPortResponse"/>

### DeleteBaremetalPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBridgeDomainRequest"/>

### DeleteBridgeDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteBridgeDomainResponse"/>

### DeleteBridgeDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteConfigNodeRequest"/>

### DeleteConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteConfigNodeResponse"/>

### DeleteConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteConfigRootRequest"/>

### DeleteConfigRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteConfigRootResponse"/>

### DeleteConfigRootResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsDatabaseNodeRequest"/>

### DeleteContrailAnalyticsDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsDatabaseNodeResponse"/>

### DeleteContrailAnalyticsDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsNodeRequest"/>

### DeleteContrailAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailAnalyticsNodeResponse"/>

### DeleteContrailAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailClusterRequest"/>

### DeleteContrailClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailClusterResponse"/>

### DeleteContrailClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailConfigDatabaseNodeRequest"/>

### DeleteContrailConfigDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailConfigDatabaseNodeResponse"/>

### DeleteContrailConfigDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailConfigNodeRequest"/>

### DeleteContrailConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailConfigNodeResponse"/>

### DeleteContrailConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailControlNodeRequest"/>

### DeleteContrailControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailControlNodeResponse"/>

### DeleteContrailControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailStorageNodeRequest"/>

### DeleteContrailStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailStorageNodeResponse"/>

### DeleteContrailStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailVrouterNodeRequest"/>

### DeleteContrailVrouterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailVrouterNodeResponse"/>

### DeleteContrailVrouterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailWebuiNodeRequest"/>

### DeleteContrailWebuiNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteContrailWebuiNodeResponse"/>

### DeleteContrailWebuiNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteCustomerAttachmentRequest"/>

### DeleteCustomerAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteCustomerAttachmentResponse"/>

### DeleteCustomerAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDashboardRequest"/>

### DeleteDashboardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDashboardResponse"/>

### DeleteDashboardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDatabaseNodeRequest"/>

### DeleteDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDatabaseNodeResponse"/>

### DeleteDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDiscoveryServiceAssignmentRequest"/>

### DeleteDiscoveryServiceAssignmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDiscoveryServiceAssignmentResponse"/>

### DeleteDiscoveryServiceAssignmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDomainRequest"/>

### DeleteDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDomainResponse"/>

### DeleteDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDsaRuleRequest"/>

### DeleteDsaRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteDsaRuleResponse"/>

### DeleteDsaRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteE2ServiceProviderRequest"/>

### DeleteE2ServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteE2ServiceProviderResponse"/>

### DeleteE2ServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFirewallPolicyRequest"/>

### DeleteFirewallPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFirewallPolicyResponse"/>

### DeleteFirewallPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFirewallRuleRequest"/>

### DeleteFirewallRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFirewallRuleResponse"/>

### DeleteFirewallRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFlavorRequest"/>

### DeleteFlavorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFlavorResponse"/>

### DeleteFlavorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFloatingIPPoolRequest"/>

### DeleteFloatingIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFloatingIPPoolResponse"/>

### DeleteFloatingIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFloatingIPRequest"/>

### DeleteFloatingIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteFloatingIPResponse"/>

### DeleteFloatingIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteForwardingClassRequest"/>

### DeleteForwardingClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteForwardingClassResponse"/>

### DeleteForwardingClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteGlobalQosConfigRequest"/>

### DeleteGlobalQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteGlobalQosConfigResponse"/>

### DeleteGlobalQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteGlobalSystemConfigRequest"/>

### DeleteGlobalSystemConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteGlobalSystemConfigResponse"/>

### DeleteGlobalSystemConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteGlobalVrouterConfigRequest"/>

### DeleteGlobalVrouterConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteGlobalVrouterConfigResponse"/>

### DeleteGlobalVrouterConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteInstanceIPRequest"/>

### DeleteInstanceIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteInstanceIPResponse"/>

### DeleteInstanceIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteInterfaceRouteTableRequest"/>

### DeleteInterfaceRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteInterfaceRouteTableResponse"/>

### DeleteInterfaceRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteKeypairRequest"/>

### DeleteKeypairRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteKeypairResponse"/>

### DeleteKeypairResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteKubernetesMasterNodeRequest"/>

### DeleteKubernetesMasterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteKubernetesMasterNodeResponse"/>

### DeleteKubernetesMasterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteKubernetesNodeRequest"/>

### DeleteKubernetesNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteKubernetesNodeResponse"/>

### DeleteKubernetesNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerHealthmonitorRequest"/>

### DeleteLoadbalancerHealthmonitorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerHealthmonitorResponse"/>

### DeleteLoadbalancerHealthmonitorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerListenerRequest"/>

### DeleteLoadbalancerListenerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerListenerResponse"/>

### DeleteLoadbalancerListenerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerMemberRequest"/>

### DeleteLoadbalancerMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerMemberResponse"/>

### DeleteLoadbalancerMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerPoolRequest"/>

### DeleteLoadbalancerPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerPoolResponse"/>

### DeleteLoadbalancerPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerRequest"/>

### DeleteLoadbalancerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLoadbalancerResponse"/>

### DeleteLoadbalancerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLocationRequest"/>

### DeleteLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLocationResponse"/>

### DeleteLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLogicalInterfaceRequest"/>

### DeleteLogicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLogicalInterfaceResponse"/>

### DeleteLogicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLogicalRouterRequest"/>

### DeleteLogicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteLogicalRouterResponse"/>

### DeleteLogicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNamespaceRequest"/>

### DeleteNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNamespaceResponse"/>

### DeleteNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNetworkDeviceConfigRequest"/>

### DeleteNetworkDeviceConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNetworkDeviceConfigResponse"/>

### DeleteNetworkDeviceConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNetworkIpamRequest"/>

### DeleteNetworkIpamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNetworkIpamResponse"/>

### DeleteNetworkIpamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNetworkPolicyRequest"/>

### DeleteNetworkPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNetworkPolicyResponse"/>

### DeleteNetworkPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNodeRequest"/>

### DeleteNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteNodeResponse"/>

### DeleteNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackComputeNodeRequest"/>

### DeleteOpenstackComputeNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackComputeNodeResponse"/>

### DeleteOpenstackComputeNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackControlNodeRequest"/>

### DeleteOpenstackControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackControlNodeResponse"/>

### DeleteOpenstackControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackMonitoringNodeRequest"/>

### DeleteOpenstackMonitoringNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackMonitoringNodeResponse"/>

### DeleteOpenstackMonitoringNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackNetworkNodeRequest"/>

### DeleteOpenstackNetworkNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackNetworkNodeResponse"/>

### DeleteOpenstackNetworkNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackStorageNodeRequest"/>

### DeleteOpenstackStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOpenstackStorageNodeResponse"/>

### DeleteOpenstackStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOsImageRequest"/>

### DeleteOsImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteOsImageResponse"/>

### DeleteOsImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePeeringPolicyRequest"/>

### DeletePeeringPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePeeringPolicyResponse"/>

### DeletePeeringPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePhysicalInterfaceRequest"/>

### DeletePhysicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePhysicalInterfaceResponse"/>

### DeletePhysicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePhysicalRouterRequest"/>

### DeletePhysicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePhysicalRouterResponse"/>

### DeletePhysicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePolicyManagementRequest"/>

### DeletePolicyManagementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePolicyManagementResponse"/>

### DeletePolicyManagementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePortRequest"/>

### DeletePortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePortResponse"/>

### DeletePortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePortTupleRequest"/>

### DeletePortTupleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeletePortTupleResponse"/>

### DeletePortTupleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteProjectRequest"/>

### DeleteProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteProjectResponse"/>

### DeleteProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteProviderAttachmentRequest"/>

### DeleteProviderAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteProviderAttachmentResponse"/>

### DeleteProviderAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteQosConfigRequest"/>

### DeleteQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteQosConfigResponse"/>

### DeleteQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteQosQueueRequest"/>

### DeleteQosQueueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteQosQueueResponse"/>

### DeleteQosQueueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRouteAggregateRequest"/>

### DeleteRouteAggregateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRouteAggregateResponse"/>

### DeleteRouteAggregateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRouteTableRequest"/>

### DeleteRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRouteTableResponse"/>

### DeleteRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRouteTargetRequest"/>

### DeleteRouteTargetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRouteTargetResponse"/>

### DeleteRouteTargetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRoutingInstanceRequest"/>

### DeleteRoutingInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRoutingInstanceResponse"/>

### DeleteRoutingInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRoutingPolicyRequest"/>

### DeleteRoutingPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteRoutingPolicyResponse"/>

### DeleteRoutingPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteSecurityGroupRequest"/>

### DeleteSecurityGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteSecurityGroupResponse"/>

### DeleteSecurityGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteSecurityLoggingObjectRequest"/>

### DeleteSecurityLoggingObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteSecurityLoggingObjectResponse"/>

### DeleteSecurityLoggingObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServerRequest"/>

### DeleteServerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServerResponse"/>

### DeleteServerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceRequest"/>

### DeleteServiceApplianceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceResponse"/>

### DeleteServiceApplianceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceSetRequest"/>

### DeleteServiceApplianceSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceApplianceSetResponse"/>

### DeleteServiceApplianceSetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceConnectionModuleRequest"/>

### DeleteServiceConnectionModuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceConnectionModuleResponse"/>

### DeleteServiceConnectionModuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceEndpointRequest"/>

### DeleteServiceEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceEndpointResponse"/>

### DeleteServiceEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceGroupRequest"/>

### DeleteServiceGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceGroupResponse"/>

### DeleteServiceGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceHealthCheckRequest"/>

### DeleteServiceHealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceHealthCheckResponse"/>

### DeleteServiceHealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceInstanceRequest"/>

### DeleteServiceInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceInstanceResponse"/>

### DeleteServiceInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceObjectRequest"/>

### DeleteServiceObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceObjectResponse"/>

### DeleteServiceObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceTemplateRequest"/>

### DeleteServiceTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteServiceTemplateResponse"/>

### DeleteServiceTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteSubnetRequest"/>

### DeleteSubnetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteSubnetResponse"/>

### DeleteSubnetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteTagRequest"/>

### DeleteTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteTagResponse"/>

### DeleteTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteTagTypeRequest"/>

### DeleteTagTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteTagTypeResponse"/>

### DeleteTagTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteUserRequest"/>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteUserResponse"/>

### DeleteUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVPNGroupRequest"/>

### DeleteVPNGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVPNGroupResponse"/>

### DeleteVPNGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSRecordRequest"/>

### DeleteVirtualDNSRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSRecordResponse"/>

### DeleteVirtualDNSRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSRequest"/>

### DeleteVirtualDNSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualDNSResponse"/>

### DeleteVirtualDNSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualIPRequest"/>

### DeleteVirtualIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualIPResponse"/>

### DeleteVirtualIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineInterfaceRequest"/>

### DeleteVirtualMachineInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineInterfaceResponse"/>

### DeleteVirtualMachineInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineRequest"/>

### DeleteVirtualMachineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualMachineResponse"/>

### DeleteVirtualMachineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualNetworkRequest"/>

### DeleteVirtualNetworkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualNetworkResponse"/>

### DeleteVirtualNetworkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualRouterRequest"/>

### DeleteVirtualRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteVirtualRouterResponse"/>

### DeleteVirtualRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteWidgetRequest"/>

### DeleteWidgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DeleteWidgetResponse"/>

### DeleteWidgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.DhcpOptionType"/>

### DhcpOptionType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dhcp_option_value | [string](#string) |  | Encoded DHCP option value (decimal) |
| dhcp_option_value_bytes | [string](#string) |  | Value of the DHCP option to be copied byte by byte |
| dhcp_option_name | [string](#string) |  | Name of the DHCP option |






<a name="github.com.Juniper.contrail.pkg.models.DhcpOptionsListType"/>

### DhcpOptionsListType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dhcp_option | [DhcpOptionType](#github.com.Juniper.contrail.pkg.models.DhcpOptionType) | repeated | List of DHCP options |






<a name="github.com.Juniper.contrail.pkg.models.DiscoveryPubSubEndPointType"/>

### DiscoveryPubSubEndPointType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ep_version | [string](#string) |  | All servers or clients whose version match this version |
| ep_id | [string](#string) |  | Specific service or client which is set of one. |
| ep_type | [string](#string) |  | Type of service or client |
| ep_prefix | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | All servers or clients whose ip match this prefix |






<a name="github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment"/>

### DiscoveryServiceAssignment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| dsa_rules | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) | repeated | Discovery service rule for assigning subscriber to publisher. (set of subscriber) can be assigned to (set of publisher). |






<a name="github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignmentType"/>

### DiscoveryServiceAssignmentType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscriber | [DiscoveryPubSubEndPointType](#github.com.Juniper.contrail.pkg.models.DiscoveryPubSubEndPointType) | repeated | subscriber set |
| publisher | [DiscoveryPubSubEndPointType](#github.com.Juniper.contrail.pkg.models.DiscoveryPubSubEndPointType) |  | Publisher set |






<a name="github.com.Juniper.contrail.pkg.models.Domain"/>

### Domain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| domain_limits | [DomainLimitsType](#github.com.Juniper.contrail.pkg.models.DomainLimitsType) |  | Domain level quota, not currently implemented |
| api_access_lists | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) | repeated | API access list is list of rules that define role based access to each API and its properties at domain level. |
| namespaces | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) | repeated | Namespace is unique networking namespace within this domain. If namespace is not present then default namespace of default project is used. |
| projects | [Project](#github.com.Juniper.contrail.pkg.models.Project) | repeated | Project represent one instance of application or tenant. |
| service_templates | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) | repeated | Service template defines how a service may be deployed in the network. Service instance is instantiated from config in service template. |
| virtual_DNSs | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) | repeated | Virtual DNS server is DNS as service for tenants. It is inbound DNS service for virtual machines in this project. DNS requests by end points inside this project/IPAM are served by this DNS server rules. |






<a name="github.com.Juniper.contrail.pkg.models.DomainLimitsType"/>

### DomainLimitsType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_limit | [int64](#int64) |  | Maximum number of projects allowed in this domain |
| virtual_network_limit | [int64](#int64) |  | Maximum number of virtual networks allowed in this domain |
| security_group_limit | [int64](#int64) |  | Maximum number of security groups allowed in this domain |






<a name="github.com.Juniper.contrail.pkg.models.DriverInfo"/>

### DriverInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipmi_address | [string](#string) |  | IPMI address of the server to manage boot device and power management |
| ipmi_username | [string](#string) |  | Username to access IPMI |
| ipmi_password | [string](#string) |  | Password to access IPMI |
| deploy_kernel | [string](#string) |  | UUID of the deploy kernel |
| deploy_ramdisk | [string](#string) |  | UUID of the deploy initrd/ramdisk |






<a name="github.com.Juniper.contrail.pkg.models.DsaRule"/>

### DsaRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| dsa_rule_entry | [DiscoveryServiceAssignmentType](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignmentType) |  | rule entry defining publisher set and subscriber set. |






<a name="github.com.Juniper.contrail.pkg.models.E2ServiceProvider"/>

### E2ServiceProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| e2_service_provider_promiscuous | [bool](#bool) |  | This service provider is connected to all other service providers. |
| physical_router_refs | [E2ServiceProviderPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.E2ServiceProviderPhysicalRouterRef) | repeated | Links the service provider to peer routers. |
| peering_policy_refs | [E2ServiceProviderPeeringPolicyRef](#github.com.Juniper.contrail.pkg.models.E2ServiceProviderPeeringPolicyRef) | repeated | Links the service provider to a peering policy. |






<a name="github.com.Juniper.contrail.pkg.models.E2ServiceProviderPeeringPolicyRef"/>

### E2ServiceProviderPeeringPolicyRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.E2ServiceProviderPhysicalRouterRef"/>

### E2ServiceProviderPhysicalRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.EcmpHashingIncludeFields"/>

### EcmpHashingIncludeFields



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination_ip | [bool](#bool) |  | When false, do not use destination ip in the ECMP hash calculation |
| ip_protocol | [bool](#bool) |  | When false, do not use ip protocol in the ECMP hash calculation |
| source_ip | [bool](#bool) |  | When false, do not use source ip in the ECMP hash calculation |
| hashing_configured | [bool](#bool) |  | When True, Packet header fields used in calculating ECMP hash is decided by following flags |
| source_port | [bool](#bool) |  | When false, do not use source port in the ECMP hash calculation |
| destination_port | [bool](#bool) |  | When false, do not use destination port in the ECMP hash calculation |






<a name="github.com.Juniper.contrail.pkg.models.EncapsulationPrioritiesType"/>

### EncapsulationPrioritiesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| encapsulation | [string](#string) | repeated | Ordered list of encapsulation types to be used in priority |






<a name="github.com.Juniper.contrail.pkg.models.FatFlowProtocols"/>

### FatFlowProtocols



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fat_flow_protocol | [ProtocolType](#github.com.Juniper.contrail.pkg.models.ProtocolType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.Filter"/>

### Filter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Filter key |
| values | [string](#string) | repeated | Filter values |






<a name="github.com.Juniper.contrail.pkg.models.FirewallPolicy"/>

### FirewallPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| firewall_rule_refs | [FirewallPolicyFirewallRuleRef](#github.com.Juniper.contrail.pkg.models.FirewallPolicyFirewallRuleRef) | repeated | Reference to firewall-rule attached to this firewall-policy |
| security_logging_object_refs | [FirewallPolicySecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.FirewallPolicySecurityLoggingObjectRef) | repeated | Reference to security-logging-object attached to this firewall-policy |






<a name="github.com.Juniper.contrail.pkg.models.FirewallPolicyFirewallRuleRef"/>

### FirewallPolicyFirewallRuleRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [FirewallSequence](#github.com.Juniper.contrail.pkg.models.FirewallSequence) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallPolicySecurityLoggingObjectRef"/>

### FirewallPolicySecurityLoggingObjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRule"/>

### FirewallRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| endpoint_1 | [FirewallRuleEndpointType](#github.com.Juniper.contrail.pkg.models.FirewallRuleEndpointType) |  | match condition for traffic source |
| endpoint_2 | [FirewallRuleEndpointType](#github.com.Juniper.contrail.pkg.models.FirewallRuleEndpointType) |  | match condition for traffic destination |
| action_list | [ActionListType](#github.com.Juniper.contrail.pkg.models.ActionListType) |  | Actions to be performed if packets match condition |
| service | [FirewallServiceType](#github.com.Juniper.contrail.pkg.models.FirewallServiceType) |  | Service (port, protocol) for packets match condition |
| direction | [string](#string) |  | Direction in the rule |
| match_tag_types | [FirewallRuleMatchTagsTypeIdList](#github.com.Juniper.contrail.pkg.models.FirewallRuleMatchTagsTypeIdList) |  | matching tags ids for source and destination endpoints |
| match_tags | [FirewallRuleMatchTagsType](#github.com.Juniper.contrail.pkg.models.FirewallRuleMatchTagsType) |  | matching tags for source and destination endpoints |
| address_group_refs | [FirewallRuleAddressGroupRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleAddressGroupRef) | repeated | Reference to address group attached to endpoints |
| security_logging_object_refs | [FirewallRuleSecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleSecurityLoggingObjectRef) | repeated | Reference to security-logging-object attached to this firewall-rule |
| virtual_network_refs | [FirewallRuleVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleVirtualNetworkRef) | repeated | Reference to virtual network attached to endpoints |
| service_group_refs | [FirewallRuleServiceGroupRef](#github.com.Juniper.contrail.pkg.models.FirewallRuleServiceGroupRef) | repeated | Reference to service-group attached to this firewall policy |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleAddressGroupRef"/>

### FirewallRuleAddressGroupRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleEndpointType"/>

### FirewallRuleEndpointType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_group | [string](#string) |  | Any workload with interface in this address-group |
| subnet | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | Any workload that belongs to this subnet |
| tags | [string](#string) | repeated | Any workload with tags matching tags in this list |
| tag_ids | [int64](#int64) | repeated | Any workload with tags ids matching all the tags ids in this list |
| virtual_network | [string](#string) |  | Any workload that belongs to this virtual network |
| any | [bool](#bool) |  | Match any workload |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleMatchTagsType"/>

### FirewallRuleMatchTagsType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_list | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleMatchTagsTypeIdList"/>

### FirewallRuleMatchTagsTypeIdList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_type | [int64](#int64) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleSecurityLoggingObjectRef"/>

### FirewallRuleSecurityLoggingObjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleServiceGroupRef"/>

### FirewallRuleServiceGroupRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallRuleVirtualNetworkRef"/>

### FirewallRuleVirtualNetworkRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallSequence"/>

### FirewallSequence



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sequence | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallServiceGroupType"/>

### FirewallServiceGroupType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_service | [FirewallServiceType](#github.com.Juniper.contrail.pkg.models.FirewallServiceType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FirewallServiceType"/>

### FirewallServiceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protocol | [string](#string) |  | Layer 4 protocol in ip packet |
| dst_ports | [PortType](#github.com.Juniper.contrail.pkg.models.PortType) |  | Range of destination port for layer 4 protocol |
| src_ports | [PortType](#github.com.Juniper.contrail.pkg.models.PortType) |  | Range of source port for layer 4 protocol |
| protocol_id | [int64](#int64) |  | Layer 4 protocol id in ip packet |






<a name="github.com.Juniper.contrail.pkg.models.Flavor"/>

### Flavor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| name | [string](#string) |  | The display name of a flavor |
| disk | [int64](#int64) |  | The size of the root disk that will be created in GiB |
| vcpus | [int64](#int64) |  | The number of virtual CPUs that will be allocated to the server |
| ram | [int64](#int64) |  | The amount of RAM a flavor has, in MiB |
| id | [string](#string) |  | The ID of the flavor, if not provided UUID will be auto-generated |
| property | [string](#string) |  | Extra specs needed to boot the image |
| rxtx_factor | [int64](#int64) |  | The receive / transmit factor (as a float) that will be set on ports if the network backend supports the QOS extension. Otherwise it will be ignored. It defaults to 1.0. |
| swap | [int64](#int64) |  | The size of a dedicated swap disk that will be allocated, in MiB. If 0 (the default), no dedicated swap disk will be created. |
| is_public | [bool](#bool) |  | Whether the flavor is public |
| ephemeral | [int64](#int64) |  | The size of the ephemeral disk that will be created, in GiB |
| links | [OpenStackLink](#github.com.Juniper.contrail.pkg.models.OpenStackLink) |  | links for the image for server instance |






<a name="github.com.Juniper.contrail.pkg.models.FloatingIP"/>

### FloatingIP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| floating_ip_address_family | [string](#string) |  | Ip address family of the floating ip, IpV4 or IpV6 |
| floating_ip_port_mappings | [PortMappings](#github.com.Juniper.contrail.pkg.models.PortMappings) |  | List of PortMaps for this floating-ip. |
| floating_ip_is_virtual_ip | [bool](#bool) |  | This floating ip is used as virtual ip (VIP) in case of LBaaS. |
| floating_ip_address | [string](#string) |  | Floating ip address. |
| floating_ip_port_mappings_enable | [bool](#bool) |  | If it is false, floating-ip Nat is done for all Ports. If it is true, floating-ip Nat is done to the list of PortMaps. |
| floating_ip_fixed_ip_address | [string](#string) |  | This floating is tracking given fixed ip of the interface. The given fixed ip is used in 1:1 NAT. |
| floating_ip_traffic_direction | [string](#string) |  | Specifies direction of traffic for the floating-ip |
| virtual_machine_interface_refs | [FloatingIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.FloatingIPVirtualMachineInterfaceRef) | repeated | Reference to virtual machine interface to which this floating ip is attached. |
| project_refs | [FloatingIPProjectRef](#github.com.Juniper.contrail.pkg.models.FloatingIPProjectRef) | repeated | Reference to project is which this floating ip was allocated. |






<a name="github.com.Juniper.contrail.pkg.models.FloatingIPPool"/>

### FloatingIPPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| floating_ip_pool_subnets | [FloatingIpPoolSubnetType](#github.com.Juniper.contrail.pkg.models.FloatingIpPoolSubnetType) |  | Subnets that restrict floating ip allocation from the corresponding virtual network. |
| floating_ips | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) | repeated | Floating ip is a ip that can be assigned to (virtual machine interface(VMI), ip), By doing so VMI can no be part of the floating ip network and floating ip is used as one:one to NAT for doing so. |






<a name="github.com.Juniper.contrail.pkg.models.FloatingIPProjectRef"/>

### FloatingIPProjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FloatingIPVirtualMachineInterfaceRef"/>

### FloatingIPVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.FloatingIpPoolSubnetType"/>

### FloatingIpPoolSubnetType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet_uuid | [string](#string) | repeated | List of subnets associated with this floating ip pool. |






<a name="github.com.Juniper.contrail.pkg.models.FlowAgingTimeout"/>

### FlowAgingTimeout



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timeout_in_seconds | [int64](#int64) |  |  |
| protocol | [string](#string) |  |  |
| port | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.FlowAgingTimeoutList"/>

### FlowAgingTimeoutList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flow_aging_timeout | [FlowAgingTimeout](#github.com.Juniper.contrail.pkg.models.FlowAgingTimeout) | repeated | List of (ip protocol, port number, timeout in seconds) |






<a name="github.com.Juniper.contrail.pkg.models.ForwardingClass"/>

### ForwardingClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| forwarding_class_dscp | [int64](#int64) |  | DSCP value to be written on outgoing packet for this forwarding-class. |
| forwarding_class_vlan_priority | [int64](#int64) |  | 802.1p value to be written on outgoing packet for this forwarding-class. |
| forwarding_class_mpls_exp | [int64](#int64) |  | MPLS exp value to be written on outgoing packet for this forwarding-class. |
| forwarding_class_id | [int64](#int64) |  | Unique ID for this forwarding class. |
| qos_queue_refs | [ForwardingClassQosQueueRef](#github.com.Juniper.contrail.pkg.models.ForwardingClassQosQueueRef) | repeated | Qos queue to be used for this forwarding class. |






<a name="github.com.Juniper.contrail.pkg.models.ForwardingClassQosQueueRef"/>

### ForwardingClassQosQueueRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAPIAccessListRequest"/>

### GetAPIAccessListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAPIAccessListResponse"/>

### GetAPIAccessListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_access_list | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAccessControlListRequest"/>

### GetAccessControlListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAccessControlListResponse"/>

### GetAccessControlListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_control_list | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAddressGroupRequest"/>

### GetAddressGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAddressGroupResponse"/>

### GetAddressGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_group | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAlarmRequest"/>

### GetAlarmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAlarmResponse"/>

### GetAlarmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alarm | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAliasIPPoolRequest"/>

### GetAliasIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAliasIPPoolResponse"/>

### GetAliasIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip_pool | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAliasIPRequest"/>

### GetAliasIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAliasIPResponse"/>

### GetAliasIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAnalyticsNodeRequest"/>

### GetAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAnalyticsNodeResponse"/>

### GetAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analytics_node | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAppformixNodeRequest"/>

### GetAppformixNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetAppformixNodeResponse"/>

### GetAppformixNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appformix_node | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetApplicationPolicySetRequest"/>

### GetApplicationPolicySetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetApplicationPolicySetResponse"/>

### GetApplicationPolicySetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_policy_set | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBGPAsAServiceRequest"/>

### GetBGPAsAServiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBGPAsAServiceResponse"/>

### GetBGPAsAServiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_as_a_service | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBGPRouterRequest"/>

### GetBGPRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBGPRouterResponse"/>

### GetBGPRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_router | [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBGPVPNRequest"/>

### GetBGPVPNRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBGPVPNResponse"/>

### GetBGPVPNResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgpvpn | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBaremetalNodeRequest"/>

### GetBaremetalNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBaremetalNodeResponse"/>

### GetBaremetalNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_node | [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBaremetalPortRequest"/>

### GetBaremetalPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBaremetalPortResponse"/>

### GetBaremetalPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_port | [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBridgeDomainRequest"/>

### GetBridgeDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetBridgeDomainResponse"/>

### GetBridgeDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bridge_domain | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetConfigNodeRequest"/>

### GetConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetConfigNodeResponse"/>

### GetConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_node | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetConfigRootRequest"/>

### GetConfigRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetConfigRootResponse"/>

### GetConfigRootResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_root | [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsDatabaseNodeRequest"/>

### GetContrailAnalyticsDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsDatabaseNodeResponse"/>

### GetContrailAnalyticsDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_database_node | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsNodeRequest"/>

### GetContrailAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailAnalyticsNodeResponse"/>

### GetContrailAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_node | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailClusterRequest"/>

### GetContrailClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailClusterResponse"/>

### GetContrailClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_cluster | [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailConfigDatabaseNodeRequest"/>

### GetContrailConfigDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailConfigDatabaseNodeResponse"/>

### GetContrailConfigDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_database_node | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailConfigNodeRequest"/>

### GetContrailConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailConfigNodeResponse"/>

### GetContrailConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_node | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailControlNodeRequest"/>

### GetContrailControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailControlNodeResponse"/>

### GetContrailControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_control_node | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailStorageNodeRequest"/>

### GetContrailStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailStorageNodeResponse"/>

### GetContrailStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_storage_node | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailVrouterNodeRequest"/>

### GetContrailVrouterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailVrouterNodeResponse"/>

### GetContrailVrouterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_vrouter_node | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailWebuiNodeRequest"/>

### GetContrailWebuiNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetContrailWebuiNodeResponse"/>

### GetContrailWebuiNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_webui_node | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetCustomerAttachmentRequest"/>

### GetCustomerAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetCustomerAttachmentResponse"/>

### GetCustomerAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_attachment | [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDashboardRequest"/>

### GetDashboardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDashboardResponse"/>

### GetDashboardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dashboard | [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDatabaseNodeRequest"/>

### GetDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDatabaseNodeResponse"/>

### GetDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_node | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDiscoveryServiceAssignmentRequest"/>

### GetDiscoveryServiceAssignmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDiscoveryServiceAssignmentResponse"/>

### GetDiscoveryServiceAssignmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| discovery_service_assignment | [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDomainRequest"/>

### GetDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDomainResponse"/>

### GetDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDsaRuleRequest"/>

### GetDsaRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetDsaRuleResponse"/>

### GetDsaRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dsa_rule | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetE2ServiceProviderRequest"/>

### GetE2ServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetE2ServiceProviderResponse"/>

### GetE2ServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_service_provider | [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFirewallPolicyRequest"/>

### GetFirewallPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFirewallPolicyResponse"/>

### GetFirewallPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_policy | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFirewallRuleRequest"/>

### GetFirewallRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFirewallRuleResponse"/>

### GetFirewallRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_rule | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFlavorRequest"/>

### GetFlavorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFlavorResponse"/>

### GetFlavorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flavor | [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFloatingIPPoolRequest"/>

### GetFloatingIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFloatingIPPoolResponse"/>

### GetFloatingIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_pool | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFloatingIPRequest"/>

### GetFloatingIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetFloatingIPResponse"/>

### GetFloatingIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetForwardingClassRequest"/>

### GetForwardingClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetForwardingClassResponse"/>

### GetForwardingClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_class | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetGlobalQosConfigRequest"/>

### GetGlobalQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetGlobalQosConfigResponse"/>

### GetGlobalQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_qos_config | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetGlobalSystemConfigRequest"/>

### GetGlobalSystemConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetGlobalSystemConfigResponse"/>

### GetGlobalSystemConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_system_config | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetGlobalVrouterConfigRequest"/>

### GetGlobalVrouterConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetGlobalVrouterConfigResponse"/>

### GetGlobalVrouterConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_vrouter_config | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetInstanceIPRequest"/>

### GetInstanceIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetInstanceIPResponse"/>

### GetInstanceIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_ip | [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetInterfaceRouteTableRequest"/>

### GetInterfaceRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetInterfaceRouteTableResponse"/>

### GetInterfaceRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_route_table | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetKeypairRequest"/>

### GetKeypairRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetKeypairResponse"/>

### GetKeypairResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keypair | [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetKubernetesMasterNodeRequest"/>

### GetKubernetesMasterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetKubernetesMasterNodeResponse"/>

### GetKubernetesMasterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_master_node | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetKubernetesNodeRequest"/>

### GetKubernetesNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetKubernetesNodeResponse"/>

### GetKubernetesNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_node | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerHealthmonitorRequest"/>

### GetLoadbalancerHealthmonitorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerHealthmonitorResponse"/>

### GetLoadbalancerHealthmonitorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_healthmonitor | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerListenerRequest"/>

### GetLoadbalancerListenerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerListenerResponse"/>

### GetLoadbalancerListenerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_listener | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerMemberRequest"/>

### GetLoadbalancerMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerMemberResponse"/>

### GetLoadbalancerMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_member | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerPoolRequest"/>

### GetLoadbalancerPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerPoolResponse"/>

### GetLoadbalancerPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_pool | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerRequest"/>

### GetLoadbalancerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLoadbalancerResponse"/>

### GetLoadbalancerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLocationRequest"/>

### GetLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLocationResponse"/>

### GetLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [Location](#github.com.Juniper.contrail.pkg.models.Location) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLogicalInterfaceRequest"/>

### GetLogicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLogicalInterfaceResponse"/>

### GetLogicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_interface | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLogicalRouterRequest"/>

### GetLogicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetLogicalRouterResponse"/>

### GetLogicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_router | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNamespaceRequest"/>

### GetNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNamespaceResponse"/>

### GetNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNetworkDeviceConfigRequest"/>

### GetNetworkDeviceConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNetworkDeviceConfigResponse"/>

### GetNetworkDeviceConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_device_config | [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNetworkIpamRequest"/>

### GetNetworkIpamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNetworkIpamResponse"/>

### GetNetworkIpamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_ipam | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNetworkPolicyRequest"/>

### GetNetworkPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNetworkPolicyResponse"/>

### GetNetworkPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_policy | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNodeRequest"/>

### GetNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetNodeResponse"/>

### GetNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#github.com.Juniper.contrail.pkg.models.Node) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackComputeNodeRequest"/>

### GetOpenstackComputeNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackComputeNodeResponse"/>

### GetOpenstackComputeNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_compute_node | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackControlNodeRequest"/>

### GetOpenstackControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackControlNodeResponse"/>

### GetOpenstackControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_control_node | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackMonitoringNodeRequest"/>

### GetOpenstackMonitoringNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackMonitoringNodeResponse"/>

### GetOpenstackMonitoringNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_monitoring_node | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackNetworkNodeRequest"/>

### GetOpenstackNetworkNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackNetworkNodeResponse"/>

### GetOpenstackNetworkNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_network_node | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackStorageNodeRequest"/>

### GetOpenstackStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOpenstackStorageNodeResponse"/>

### GetOpenstackStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_storage_node | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOsImageRequest"/>

### GetOsImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetOsImageResponse"/>

### GetOsImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| os_image | [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPeeringPolicyRequest"/>

### GetPeeringPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPeeringPolicyResponse"/>

### GetPeeringPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peering_policy | [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPhysicalInterfaceRequest"/>

### GetPhysicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPhysicalInterfaceResponse"/>

### GetPhysicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_interface | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPhysicalRouterRequest"/>

### GetPhysicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPhysicalRouterResponse"/>

### GetPhysicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_router | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPolicyManagementRequest"/>

### GetPolicyManagementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPolicyManagementResponse"/>

### GetPolicyManagementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_management | [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPortRequest"/>

### GetPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPortResponse"/>

### GetPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [Port](#github.com.Juniper.contrail.pkg.models.Port) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPortTupleRequest"/>

### GetPortTupleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetPortTupleResponse"/>

### GetPortTupleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_tuple | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetProjectRequest"/>

### GetProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetProjectResponse"/>

### GetProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#github.com.Juniper.contrail.pkg.models.Project) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetProviderAttachmentRequest"/>

### GetProviderAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetProviderAttachmentResponse"/>

### GetProviderAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_attachment | [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetQosConfigRequest"/>

### GetQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetQosConfigResponse"/>

### GetQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_config | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetQosQueueRequest"/>

### GetQosQueueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetQosQueueResponse"/>

### GetQosQueueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_queue | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRouteAggregateRequest"/>

### GetRouteAggregateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRouteAggregateResponse"/>

### GetRouteAggregateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_aggregate | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRouteTableRequest"/>

### GetRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRouteTableResponse"/>

### GetRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_table | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRouteTargetRequest"/>

### GetRouteTargetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRouteTargetResponse"/>

### GetRouteTargetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_target | [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRoutingInstanceRequest"/>

### GetRoutingInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRoutingInstanceResponse"/>

### GetRoutingInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instance | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRoutingPolicyRequest"/>

### GetRoutingPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetRoutingPolicyResponse"/>

### GetRoutingPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_policy | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetSecurityGroupRequest"/>

### GetSecurityGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetSecurityGroupResponse"/>

### GetSecurityGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_group | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetSecurityLoggingObjectRequest"/>

### GetSecurityLoggingObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetSecurityLoggingObjectResponse"/>

### GetSecurityLoggingObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_logging_object | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServerRequest"/>

### GetServerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServerResponse"/>

### GetServerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [Server](#github.com.Juniper.contrail.pkg.models.Server) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceApplianceRequest"/>

### GetServiceApplianceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceApplianceResponse"/>

### GetServiceApplianceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceApplianceSetRequest"/>

### GetServiceApplianceSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceApplianceSetResponse"/>

### GetServiceApplianceSetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance_set | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceConnectionModuleRequest"/>

### GetServiceConnectionModuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceConnectionModuleResponse"/>

### GetServiceConnectionModuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_connection_module | [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceEndpointRequest"/>

### GetServiceEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceEndpointResponse"/>

### GetServiceEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_endpoint | [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceGroupRequest"/>

### GetServiceGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceGroupResponse"/>

### GetServiceGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_group | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceHealthCheckRequest"/>

### GetServiceHealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceHealthCheckResponse"/>

### GetServiceHealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_health_check | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceInstanceRequest"/>

### GetServiceInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceInstanceResponse"/>

### GetServiceInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_instance | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceObjectRequest"/>

### GetServiceObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceObjectResponse"/>

### GetServiceObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_object | [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceTemplateRequest"/>

### GetServiceTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetServiceTemplateResponse"/>

### GetServiceTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_template | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetSubnetRequest"/>

### GetSubnetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetSubnetResponse"/>

### GetSubnetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetTagRequest"/>

### GetTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetTagResponse"/>

### GetTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetTagTypeRequest"/>

### GetTagTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetTagTypeResponse"/>

### GetTagTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_type | [TagType](#github.com.Juniper.contrail.pkg.models.TagType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetUserRequest"/>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetUserResponse"/>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#github.com.Juniper.contrail.pkg.models.User) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVPNGroupRequest"/>

### GetVPNGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVPNGroupResponse"/>

### GetVPNGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpn_group | [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualDNSRecordRequest"/>

### GetVirtualDNSRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualDNSRecordResponse"/>

### GetVirtualDNSRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS_record | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualDNSRequest"/>

### GetVirtualDNSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualDNSResponse"/>

### GetVirtualDNSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualIPRequest"/>

### GetVirtualIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualIPResponse"/>

### GetVirtualIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_ip | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualMachineInterfaceRequest"/>

### GetVirtualMachineInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualMachineInterfaceResponse"/>

### GetVirtualMachineInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine_interface | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualMachineRequest"/>

### GetVirtualMachineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualMachineResponse"/>

### GetVirtualMachineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine | [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualNetworkRequest"/>

### GetVirtualNetworkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualNetworkResponse"/>

### GetVirtualNetworkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_network | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualRouterRequest"/>

### GetVirtualRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetVirtualRouterResponse"/>

### GetVirtualRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_router | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetWidgetRequest"/>

### GetWidgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GetWidgetResponse"/>

### GetWidgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| widget | [Widget](#github.com.Juniper.contrail.pkg.models.Widget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.GlobalQosConfig"/>

### GlobalQosConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| control_traffic_dscp | [ControlTrafficDscpType](#github.com.Juniper.contrail.pkg.models.ControlTrafficDscpType) |  | DSCP value of IP header for control traffic |
| forwarding_classs | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) | repeated | Link to global-qos config. |
| qos_configs | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) | repeated | Global system QoS config for vhost and fabric traffic . |
| qos_queues | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) | repeated | QOS queue config object in this project. |






<a name="github.com.Juniper.contrail.pkg.models.GlobalSystemConfig"/>

### GlobalSystemConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| config_version | [string](#string) |  | Version of OpenContrail software that generated this config. |
| bgpaas_parameters | [BGPaaServiceParametersType](#github.com.Juniper.contrail.pkg.models.BGPaaServiceParametersType) |  | BGP As A Service Parameters configuration |
| alarm_enable | [bool](#bool) |  | Flag to enable/disable alarms configured under global-system-config. True, if not set. |
| mac_move_control | [MACMoveLimitControlType](#github.com.Juniper.contrail.pkg.models.MACMoveLimitControlType) |  | MAC move control on the network |
| plugin_tuning | [PluginProperties](#github.com.Juniper.contrail.pkg.models.PluginProperties) |  | Various Orchestration system plugin(interface) parameters, like Openstack Neutron plugin. |
| ibgp_auto_mesh | [bool](#bool) |  | When true, system will automatically create BGP peering mesh with all control-nodes that have same BGP AS number as global AS number. |
| mac_aging_time | [int64](#int64) |  | MAC aging time on the network |
| bgp_always_compare_med | [bool](#bool) |  | Always compare MED even if paths are received from different ASes. |
| user_defined_log_statistics | [UserDefinedLogStatList](#github.com.Juniper.contrail.pkg.models.UserDefinedLogStatList) |  | stats name and patterns |
| graceful_restart_parameters | [GracefulRestartParametersType](#github.com.Juniper.contrail.pkg.models.GracefulRestartParametersType) |  | Graceful Restart parameters |
| ip_fabric_subnets | [SubnetListType](#github.com.Juniper.contrail.pkg.models.SubnetListType) |  | List of all subnets in which vrouter ip address exist. Used by Device manager to configure dynamic GRE tunnels on the SDN gateway. |
| autonomous_system | [int64](#int64) |  | 16 bit BGP Autonomous System number for the cluster. |
| mac_limit_control | [MACLimitControlType](#github.com.Juniper.contrail.pkg.models.MACLimitControlType) |  | MAC limit control on the network |
| bgp_router_refs | [GlobalSystemConfigBGPRouterRef](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfigBGPRouterRef) | repeated | List of references to all bgp routers in systems. |
| alarms | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) | repeated | List of alarms that are applicable to objects anchored under global-system-config. |
| analytics_nodes | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) | repeated | Analytics node is object representing a logical node in system which serves operational API and analytics collector. |
| api_access_lists | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) | repeated | Global API access list applicable to all domain and projects |
| config_nodes | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) | repeated | Config node is object representing a logical node in system which serves config API. |
| database_nodes | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) | repeated | Database node is object representing a logical node in system which host Cassandra DB and Zookeeper. |
| global_qos_configs | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) | repeated | Global QoS system config is object where all global system QoS configuration is present. |
| global_vrouter_configs | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) | repeated | Global vrouter config is object where all global vrouter config is present. |
| physical_routers | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) | repeated | Physical router object represent any physical device that participates in virtual networking, like routers, switches, servers, firewalls etc. |
| service_appliance_sets | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) | repeated |  |
| virtual_routers | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) | repeated | Virtual router is packet forwarding system on devices such as compute nodes(servers), TOR(s), routers. |






<a name="github.com.Juniper.contrail.pkg.models.GlobalSystemConfigBGPRouterRef"/>

### GlobalSystemConfigBGPRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig"/>

### GlobalVrouterConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| ecmp_hashing_include_fields | [EcmpHashingIncludeFields](#github.com.Juniper.contrail.pkg.models.EcmpHashingIncludeFields) |  | ECMP hashing config at global level. |
| flow_aging_timeout_list | [FlowAgingTimeoutList](#github.com.Juniper.contrail.pkg.models.FlowAgingTimeoutList) |  | Flow aging timeout per application (protocol, port) list. |
| forwarding_mode | [string](#string) |  | Packet forwarding mode for this system L2-only, L3-only OR L2-L3. L2-L3 is default. |
| flow_export_rate | [int64](#int64) |  | Flow export rate is global config, rate at which each vrouter will sample and export flow records to analytics |
| linklocal_services | [LinklocalServicesTypes](#github.com.Juniper.contrail.pkg.models.LinklocalServicesTypes) |  | Global services provided on link local subnet to the virtual machines. |
| encapsulation_priorities | [EncapsulationPrioritiesType](#github.com.Juniper.contrail.pkg.models.EncapsulationPrioritiesType) |  | Ordered list of encapsulations that vrouter will use in priority order. |
| vxlan_network_identifier_mode | [string](#string) |  |  |
| enable_security_logging | [bool](#bool) |  | Enable or disable security-logging in the system |
| security_logging_objects | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) | repeated | Reference to security logging object for global-vrouter-config. |






<a name="github.com.Juniper.contrail.pkg.models.GracefulRestartParametersType"/>

### GracefulRestartParametersType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enable | [bool](#bool) |  | Enable/Disable knob for all GR parameters to take effect |
| end_of_rib_timeout | [int64](#int64) |  | Maximum time (in seconds) to wait for EndOfRib reception/transmission |
| bgp_helper_enable | [bool](#bool) |  | Enable GR Helper mode for BGP peers in contrail-control |
| xmpp_helper_enable | [bool](#bool) |  | Enable GR Helper mode for XMPP peers (agents) in contrail-control |
| restart_time | [int64](#int64) |  | Time (in seconds) taken by the restarting speaker to get back to stable state |
| long_lived_restart_time | [int64](#int64) |  | Extended Time (in seconds) taken by the restarting speaker after restart-time to get back to stable state |






<a name="github.com.Juniper.contrail.pkg.models.IdPermsType"/>

### IdPermsType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enable | [bool](#bool) |  | Administratively Enable/Disable this object |
| description | [string](#string) |  | User provided text |
| created | [string](#string) |  | Time when this object was created |
| creator | [string](#string) |  | Id of tenant who created this object |
| user_visible | [bool](#bool) |  | System created internal objects will have this flag set and will not be visible |
| last_modified | [string](#string) |  | Time when this object was created |
| permissions | [PermType](#github.com.Juniper.contrail.pkg.models.PermType) |  | No longer used, will be removed |






<a name="github.com.Juniper.contrail.pkg.models.InstanceIP"/>

### InstanceIP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_health_check_ip | [bool](#bool) |  | This instance ip is used as service health check source ip |
| secondary_ip_tracking_ip | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | When this instance ip is secondary ip, it can track activeness of another ip. |
| instance_ip_address | [string](#string) |  | Ip address value for instance ip. |
| instance_ip_mode | [string](#string) |  | Ip address HA mode in case this instance ip is used in more than one interface, active-Active or active-Standby. |
| subnet_uuid | [string](#string) |  | This instance ip was allocated from this Subnet(UUID). |
| instance_ip_family | [string](#string) |  | Ip address family for instance ip, IPv4(v4) or IPv6(v6). |
| service_instance_ip | [bool](#bool) |  | This instance ip is used as service chain next hop |
| instance_ip_local_ip | [bool](#bool) |  | This instance ip is local to compute and will not be exported to other nodes. |
| instance_ip_secondary | [bool](#bool) |  | This instance ip is secondary ip of the interface. |
| virtual_machine_interface_refs | [InstanceIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.InstanceIPVirtualMachineInterfaceRef) | repeated | Reference to virtual machine interface to which this instance ip is attached. |
| physical_router_refs | [InstanceIPPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.InstanceIPPhysicalRouterRef) | repeated | This instance ip is used as IRB address on the referenced physical router (e.g.MX), In case of OVSDB TOR usecase this address will be used as default gateway for Host behind the TOR. |
| virtual_router_refs | [InstanceIPVirtualRouterRef](#github.com.Juniper.contrail.pkg.models.InstanceIPVirtualRouterRef) | repeated | Reference to virtual router of this instance ip. |
| network_ipam_refs | [InstanceIPNetworkIpamRef](#github.com.Juniper.contrail.pkg.models.InstanceIPNetworkIpamRef) | repeated | Reference to network ipam of this instance ip. |
| virtual_network_refs | [InstanceIPVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.InstanceIPVirtualNetworkRef) | repeated | Reference to virtual network of this instance ip. |
| floating_ips | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) | repeated | floating-ip can be child of instance-ip. By doing so instance-ip can be used as floating-ip. |






<a name="github.com.Juniper.contrail.pkg.models.InstanceIPNetworkIpamRef"/>

### InstanceIPNetworkIpamRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.InstanceIPPhysicalRouterRef"/>

### InstanceIPPhysicalRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.InstanceIPVirtualMachineInterfaceRef"/>

### InstanceIPVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.InstanceIPVirtualNetworkRef"/>

### InstanceIPVirtualNetworkRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.InstanceIPVirtualRouterRef"/>

### InstanceIPVirtualRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.InstanceInfo"/>

### InstanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| display_name | [string](#string) |  | Name of the nova instance |
| image_source | [string](#string) |  | UUID of the image for instance |
| local_gb | [string](#string) |  |  |
| memory_mb | [string](#string) |  |  |
| nova_host_id | [string](#string) |  |  |
| root_gb | [string](#string) |  |  |
| swap_mb | [string](#string) |  |  |
| vcpus | [string](#string) |  |  |
| capabilities | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.InterfaceMirrorType"/>

### InterfaceMirrorType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| traffic_direction | [string](#string) |  | Specifies direction of traffic to mirror, Ingress, Egress or both |
| mirror_to | [MirrorActionType](#github.com.Juniper.contrail.pkg.models.MirrorActionType) |  | Mirror destination configuration |






<a name="github.com.Juniper.contrail.pkg.models.InterfaceRouteTable"/>

### InterfaceRouteTable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| interface_route_table_routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  | Interface route table used same structure as route table, however the next hop field is irrelevant. |
| service_instance_refs | [InterfaceRouteTableServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTableServiceInstanceRef) | repeated | Reference to interface route table attached to (service instance, interface), This is used to add interface static routes to service instance interface. |






<a name="github.com.Juniper.contrail.pkg.models.InterfaceRouteTableServiceInstanceRef"/>

### InterfaceRouteTableServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [ServiceInterfaceTag](#github.com.Juniper.contrail.pkg.models.ServiceInterfaceTag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.IpAddressesType"/>

### IpAddressesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip_address | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.IpamDnsAddressType"/>

### IpamDnsAddressType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_dns_server_address | [IpAddressesType](#github.com.Juniper.contrail.pkg.models.IpAddressesType) |  | In case of tenant DNS server method, Ip address of DNS server. This will be given in DHCP |
| virtual_dns_server_name | [string](#string) |  | In case of virtual DNS server, name of virtual DNS server |






<a name="github.com.Juniper.contrail.pkg.models.IpamSubnetType"/>

### IpamSubnetType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | ip prefix and length for the subnet |
| addr_from_start | [bool](#bool) |  | Start address allocation from start or from end of address range. |
| enable_dhcp | [bool](#bool) |  | Enable DHCP for the VM(s) in this subnet |
| default_gateway | [string](#string) |  | default-gateway ip address in the subnet, if not provided one is auto generated by the system. |
| alloc_unit | [int64](#int64) |  | allocation unit for this subnet to allocate bulk ip addresses |
| created | [string](#string) |  | timestamp when subnet object gets created |
| dns_nameservers | [string](#string) | repeated | Tenant DNS servers ip address in tenant DNS method |
| dhcp_option_list | [DhcpOptionsListType](#github.com.Juniper.contrail.pkg.models.DhcpOptionsListType) |  | DHCP options list to be sent via DHCP for VM(s) in this subnet |
| subnet_uuid | [string](#string) |  | Subnet UUID is auto generated by the system |
| allocation_pools | [AllocationPoolType](#github.com.Juniper.contrail.pkg.models.AllocationPoolType) | repeated | List of ranges of ip address within the subnet from which to allocate ip address. default is entire prefix |
| last_modified | [string](#string) |  | timestamp when subnet object gets updated |
| host_routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  | Host routes to be sent via DHCP for VM(s) in this subnet, Next hop for these routes is always default gateway |
| dns_server_address | [string](#string) |  | DNS server ip address in the subnet, if not provided one is auto generated by the system. |
| subnet_name | [string](#string) |  | User provided name for this subnet |






<a name="github.com.Juniper.contrail.pkg.models.IpamSubnets"/>

### IpamSubnets



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnets | [IpamSubnetType](#github.com.Juniper.contrail.pkg.models.IpamSubnetType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.IpamType"/>

### IpamType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipam_method | [string](#string) |  |  |
| ipam_dns_method | [string](#string) |  |  |
| ipam_dns_server | [IpamDnsAddressType](#github.com.Juniper.contrail.pkg.models.IpamDnsAddressType) |  |  |
| dhcp_option_list | [DhcpOptionsListType](#github.com.Juniper.contrail.pkg.models.DhcpOptionsListType) |  |  |
| host_routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  |  |
| cidr_block | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.JunosServicePorts"/>

### JunosServicePorts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_port | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.KeyValuePair"/>

### KeyValuePair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |
| key | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.KeyValuePairs"/>

### KeyValuePairs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key_value_pair | [KeyValuePair](#github.com.Juniper.contrail.pkg.models.KeyValuePair) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.Keypair"/>

### Keypair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| name | [string](#string) |  | A name for the keypair which will be used to reference it later |
| public_key | [string](#string) |  | SSH Public Key |






<a name="github.com.Juniper.contrail.pkg.models.KubernetesMasterNode"/>

### KubernetesMasterNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [KubernetesMasterNodeNodeRef](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNodeNodeRef) | repeated | Reference to node object for this kubernetes master node. |






<a name="github.com.Juniper.contrail.pkg.models.KubernetesMasterNodeNodeRef"/>

### KubernetesMasterNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.KubernetesNode"/>

### KubernetesNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [KubernetesNodeNodeRef](#github.com.Juniper.contrail.pkg.models.KubernetesNodeNodeRef) | repeated | Reference to node object for this kubernetes node. |






<a name="github.com.Juniper.contrail.pkg.models.KubernetesNodeNodeRef"/>

### KubernetesNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LinklocalServiceEntryType"/>

### LinklocalServiceEntryType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip_fabric_service_ip | [string](#string) | repeated | Destination ip address to which link local traffic will forwarded |
| linklocal_service_name | [string](#string) |  | Name of the link local service. VM name resolution of this name will result in link local ip address |
| linklocal_service_ip | [string](#string) |  | ip address of the link local service. |
| ip_fabric_service_port | [int64](#int64) |  | Destination TCP port number to which link local traffic will forwarded |
| ip_fabric_DNS_service_name | [string](#string) |  | DNS name to which link local service will be proxied |
| linklocal_service_port | [int64](#int64) |  | Destination TCP port number of link local service |






<a name="github.com.Juniper.contrail.pkg.models.LinklocalServicesTypes"/>

### LinklocalServicesTypes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| linklocal_service_entry | [LinklocalServiceEntryType](#github.com.Juniper.contrail.pkg.models.LinklocalServiceEntryType) | repeated | List of link local services |






<a name="github.com.Juniper.contrail.pkg.models.ListAPIAccessListRequest"/>

### ListAPIAccessListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAPIAccessListResponse"/>

### ListAPIAccessListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_access_lists | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAccessControlListRequest"/>

### ListAccessControlListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAccessControlListResponse"/>

### ListAccessControlListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_control_lists | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAddressGroupRequest"/>

### ListAddressGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAddressGroupResponse"/>

### ListAddressGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_groups | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAlarmRequest"/>

### ListAlarmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAlarmResponse"/>

### ListAlarmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alarms | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAliasIPPoolRequest"/>

### ListAliasIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAliasIPPoolResponse"/>

### ListAliasIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip_pools | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAliasIPRequest"/>

### ListAliasIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAliasIPResponse"/>

### ListAliasIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ips | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAnalyticsNodeRequest"/>

### ListAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAnalyticsNodeResponse"/>

### ListAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analytics_nodes | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAppformixNodeRequest"/>

### ListAppformixNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListAppformixNodeResponse"/>

### ListAppformixNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appformix_nodes | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListApplicationPolicySetRequest"/>

### ListApplicationPolicySetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListApplicationPolicySetResponse"/>

### ListApplicationPolicySetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_policy_sets | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBGPAsAServiceRequest"/>

### ListBGPAsAServiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBGPAsAServiceResponse"/>

### ListBGPAsAServiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_as_a_services | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBGPRouterRequest"/>

### ListBGPRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBGPRouterResponse"/>

### ListBGPRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_routers | [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBGPVPNRequest"/>

### ListBGPVPNRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBGPVPNResponse"/>

### ListBGPVPNResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgpvpns | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBaremetalNodeRequest"/>

### ListBaremetalNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBaremetalNodeResponse"/>

### ListBaremetalNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_nodes | [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBaremetalPortRequest"/>

### ListBaremetalPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBaremetalPortResponse"/>

### ListBaremetalPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_ports | [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBridgeDomainRequest"/>

### ListBridgeDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListBridgeDomainResponse"/>

### ListBridgeDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bridge_domains | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListConfigNodeRequest"/>

### ListConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListConfigNodeResponse"/>

### ListConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_nodes | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListConfigRootRequest"/>

### ListConfigRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListConfigRootResponse"/>

### ListConfigRootResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_roots | [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsDatabaseNodeRequest"/>

### ListContrailAnalyticsDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsDatabaseNodeResponse"/>

### ListContrailAnalyticsDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_database_nodes | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsNodeRequest"/>

### ListContrailAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailAnalyticsNodeResponse"/>

### ListContrailAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_nodes | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailClusterRequest"/>

### ListContrailClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailClusterResponse"/>

### ListContrailClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_clusters | [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailConfigDatabaseNodeRequest"/>

### ListContrailConfigDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailConfigDatabaseNodeResponse"/>

### ListContrailConfigDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_database_nodes | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailConfigNodeRequest"/>

### ListContrailConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailConfigNodeResponse"/>

### ListContrailConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_nodes | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailControlNodeRequest"/>

### ListContrailControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailControlNodeResponse"/>

### ListContrailControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_control_nodes | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailStorageNodeRequest"/>

### ListContrailStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailStorageNodeResponse"/>

### ListContrailStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_storage_nodes | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailVrouterNodeRequest"/>

### ListContrailVrouterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailVrouterNodeResponse"/>

### ListContrailVrouterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_vrouter_nodes | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailWebuiNodeRequest"/>

### ListContrailWebuiNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListContrailWebuiNodeResponse"/>

### ListContrailWebuiNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_webui_nodes | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListCustomerAttachmentRequest"/>

### ListCustomerAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListCustomerAttachmentResponse"/>

### ListCustomerAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_attachments | [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDashboardRequest"/>

### ListDashboardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDashboardResponse"/>

### ListDashboardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dashboards | [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDatabaseNodeRequest"/>

### ListDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDatabaseNodeResponse"/>

### ListDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_nodes | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDiscoveryServiceAssignmentRequest"/>

### ListDiscoveryServiceAssignmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDiscoveryServiceAssignmentResponse"/>

### ListDiscoveryServiceAssignmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| discovery_service_assignments | [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDomainRequest"/>

### ListDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDomainResponse"/>

### ListDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domains | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDsaRuleRequest"/>

### ListDsaRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListDsaRuleResponse"/>

### ListDsaRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dsa_rules | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListE2ServiceProviderRequest"/>

### ListE2ServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListE2ServiceProviderResponse"/>

### ListE2ServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_service_providers | [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFirewallPolicyRequest"/>

### ListFirewallPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFirewallPolicyResponse"/>

### ListFirewallPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_policys | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFirewallRuleRequest"/>

### ListFirewallRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFirewallRuleResponse"/>

### ListFirewallRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_rules | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFlavorRequest"/>

### ListFlavorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFlavorResponse"/>

### ListFlavorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flavors | [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFloatingIPPoolRequest"/>

### ListFloatingIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFloatingIPPoolResponse"/>

### ListFloatingIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_pools | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFloatingIPRequest"/>

### ListFloatingIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListFloatingIPResponse"/>

### ListFloatingIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ips | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListForwardingClassRequest"/>

### ListForwardingClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListForwardingClassResponse"/>

### ListForwardingClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_classs | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListGlobalQosConfigRequest"/>

### ListGlobalQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListGlobalQosConfigResponse"/>

### ListGlobalQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_qos_configs | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListGlobalSystemConfigRequest"/>

### ListGlobalSystemConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListGlobalSystemConfigResponse"/>

### ListGlobalSystemConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_system_configs | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListGlobalVrouterConfigRequest"/>

### ListGlobalVrouterConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListGlobalVrouterConfigResponse"/>

### ListGlobalVrouterConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_vrouter_configs | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListInstanceIPRequest"/>

### ListInstanceIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListInstanceIPResponse"/>

### ListInstanceIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_ips | [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListInterfaceRouteTableRequest"/>

### ListInterfaceRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListInterfaceRouteTableResponse"/>

### ListInterfaceRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_route_tables | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListKeypairRequest"/>

### ListKeypairRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListKeypairResponse"/>

### ListKeypairResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keypairs | [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListKubernetesMasterNodeRequest"/>

### ListKubernetesMasterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListKubernetesMasterNodeResponse"/>

### ListKubernetesMasterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_master_nodes | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListKubernetesNodeRequest"/>

### ListKubernetesNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListKubernetesNodeResponse"/>

### ListKubernetesNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_nodes | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerHealthmonitorRequest"/>

### ListLoadbalancerHealthmonitorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerHealthmonitorResponse"/>

### ListLoadbalancerHealthmonitorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_healthmonitors | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerListenerRequest"/>

### ListLoadbalancerListenerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerListenerResponse"/>

### ListLoadbalancerListenerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_listeners | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerMemberRequest"/>

### ListLoadbalancerMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerMemberResponse"/>

### ListLoadbalancerMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_members | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerPoolRequest"/>

### ListLoadbalancerPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerPoolResponse"/>

### ListLoadbalancerPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_pools | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerRequest"/>

### ListLoadbalancerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLoadbalancerResponse"/>

### ListLoadbalancerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancers | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLocationRequest"/>

### ListLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLocationResponse"/>

### ListLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locations | [Location](#github.com.Juniper.contrail.pkg.models.Location) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLogicalInterfaceRequest"/>

### ListLogicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLogicalInterfaceResponse"/>

### ListLogicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_interfaces | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLogicalRouterRequest"/>

### ListLogicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListLogicalRouterResponse"/>

### ListLogicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_routers | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNamespaceRequest"/>

### ListNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNamespaceResponse"/>

### ListNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaces | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNetworkDeviceConfigRequest"/>

### ListNetworkDeviceConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNetworkDeviceConfigResponse"/>

### ListNetworkDeviceConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_device_configs | [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNetworkIpamRequest"/>

### ListNetworkIpamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNetworkIpamResponse"/>

### ListNetworkIpamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_ipams | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNetworkPolicyRequest"/>

### ListNetworkPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNetworkPolicyResponse"/>

### ListNetworkPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_policys | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNodeRequest"/>

### ListNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListNodeResponse"/>

### ListNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [Node](#github.com.Juniper.contrail.pkg.models.Node) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackComputeNodeRequest"/>

### ListOpenstackComputeNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackComputeNodeResponse"/>

### ListOpenstackComputeNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_compute_nodes | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackControlNodeRequest"/>

### ListOpenstackControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackControlNodeResponse"/>

### ListOpenstackControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_control_nodes | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackMonitoringNodeRequest"/>

### ListOpenstackMonitoringNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackMonitoringNodeResponse"/>

### ListOpenstackMonitoringNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_monitoring_nodes | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackNetworkNodeRequest"/>

### ListOpenstackNetworkNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackNetworkNodeResponse"/>

### ListOpenstackNetworkNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_network_nodes | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackStorageNodeRequest"/>

### ListOpenstackStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOpenstackStorageNodeResponse"/>

### ListOpenstackStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_storage_nodes | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOsImageRequest"/>

### ListOsImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListOsImageResponse"/>

### ListOsImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| os_images | [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPeeringPolicyRequest"/>

### ListPeeringPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPeeringPolicyResponse"/>

### ListPeeringPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peering_policys | [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPhysicalInterfaceRequest"/>

### ListPhysicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPhysicalInterfaceResponse"/>

### ListPhysicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_interfaces | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPhysicalRouterRequest"/>

### ListPhysicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPhysicalRouterResponse"/>

### ListPhysicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_routers | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPolicyManagementRequest"/>

### ListPolicyManagementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPolicyManagementResponse"/>

### ListPolicyManagementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_managements | [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPortRequest"/>

### ListPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPortResponse"/>

### ListPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ports | [Port](#github.com.Juniper.contrail.pkg.models.Port) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPortTupleRequest"/>

### ListPortTupleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListPortTupleResponse"/>

### ListPortTupleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_tuples | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListProjectRequest"/>

### ListProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListProjectResponse"/>

### ListProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [Project](#github.com.Juniper.contrail.pkg.models.Project) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListProviderAttachmentRequest"/>

### ListProviderAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListProviderAttachmentResponse"/>

### ListProviderAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_attachments | [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListQosConfigRequest"/>

### ListQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListQosConfigResponse"/>

### ListQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_configs | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListQosQueueRequest"/>

### ListQosQueueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListQosQueueResponse"/>

### ListQosQueueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_queues | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRouteAggregateRequest"/>

### ListRouteAggregateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRouteAggregateResponse"/>

### ListRouteAggregateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_aggregates | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRouteTableRequest"/>

### ListRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRouteTableResponse"/>

### ListRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_tables | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRouteTargetRequest"/>

### ListRouteTargetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRouteTargetResponse"/>

### ListRouteTargetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_targets | [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRoutingInstanceRequest"/>

### ListRoutingInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRoutingInstanceResponse"/>

### ListRoutingInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instances | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRoutingPolicyRequest"/>

### ListRoutingPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListRoutingPolicyResponse"/>

### ListRoutingPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_policys | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListSecurityGroupRequest"/>

### ListSecurityGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListSecurityGroupResponse"/>

### ListSecurityGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_groups | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListSecurityLoggingObjectRequest"/>

### ListSecurityLoggingObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListSecurityLoggingObjectResponse"/>

### ListSecurityLoggingObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_logging_objects | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServerRequest"/>

### ListServerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServerResponse"/>

### ListServerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| servers | [Server](#github.com.Juniper.contrail.pkg.models.Server) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceApplianceRequest"/>

### ListServiceApplianceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceApplianceResponse"/>

### ListServiceApplianceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliances | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceApplianceSetRequest"/>

### ListServiceApplianceSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceApplianceSetResponse"/>

### ListServiceApplianceSetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance_sets | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceConnectionModuleRequest"/>

### ListServiceConnectionModuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceConnectionModuleResponse"/>

### ListServiceConnectionModuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_connection_modules | [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceEndpointRequest"/>

### ListServiceEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceEndpointResponse"/>

### ListServiceEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_endpoints | [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceGroupRequest"/>

### ListServiceGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceGroupResponse"/>

### ListServiceGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_groups | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceHealthCheckRequest"/>

### ListServiceHealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceHealthCheckResponse"/>

### ListServiceHealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_health_checks | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceInstanceRequest"/>

### ListServiceInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceInstanceResponse"/>

### ListServiceInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_instances | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceObjectRequest"/>

### ListServiceObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceObjectResponse"/>

### ListServiceObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_objects | [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceTemplateRequest"/>

### ListServiceTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListServiceTemplateResponse"/>

### ListServiceTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_templates | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListSpec"/>

### ListSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [Filter](#github.com.Juniper.contrail.pkg.models.Filter) | repeated | QueryFilter |
| limit | [int64](#int64) |  | Number of items expected to be returned |
| offset | [int64](#int64) |  | Starting offset of items |
| detail | [bool](#bool) |  | Include detail informatoin or not |
| count | [bool](#bool) |  | TBD |
| shared | [bool](#bool) |  | Include shared resources or not |
| exclude_hrefs | [bool](#bool) |  | Exclude href parameters |
| parent_fq_name | [string](#string) | repeated | Filter by parent FQ Name |
| parent_type | [string](#string) |  | Filter by parent type |
| parent_uuids | [string](#string) | repeated | Filter by parent UUIDs |
| backref_uuids | [string](#string) | repeated | Filter by backref UUIDss |
| object_uuids | [string](#string) | repeated | Filter by UUIDs |
| fields | [string](#string) | repeated | limit displayed fields |






<a name="github.com.Juniper.contrail.pkg.models.ListSubnetRequest"/>

### ListSubnetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListSubnetResponse"/>

### ListSubnetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnets | [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListTagRequest"/>

### ListTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListTagResponse"/>

### ListTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListTagTypeRequest"/>

### ListTagTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListTagTypeResponse"/>

### ListTagTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_types | [TagType](#github.com.Juniper.contrail.pkg.models.TagType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListUserRequest"/>

### ListUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListUserResponse"/>

### ListUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#github.com.Juniper.contrail.pkg.models.User) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVPNGroupRequest"/>

### ListVPNGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVPNGroupResponse"/>

### ListVPNGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpn_groups | [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualDNSRecordRequest"/>

### ListVirtualDNSRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualDNSRecordResponse"/>

### ListVirtualDNSRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS_records | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualDNSRequest"/>

### ListVirtualDNSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualDNSResponse"/>

### ListVirtualDNSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNSs | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualIPRequest"/>

### ListVirtualIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualIPResponse"/>

### ListVirtualIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_ips | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualMachineInterfaceRequest"/>

### ListVirtualMachineInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualMachineInterfaceResponse"/>

### ListVirtualMachineInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine_interfaces | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualMachineRequest"/>

### ListVirtualMachineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualMachineResponse"/>

### ListVirtualMachineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machines | [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualNetworkRequest"/>

### ListVirtualNetworkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualNetworkResponse"/>

### ListVirtualNetworkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_networks | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualRouterRequest"/>

### ListVirtualRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListVirtualRouterResponse"/>

### ListVirtualRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_routers | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ListWidgetRequest"/>

### ListWidgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ListSpec](#github.com.Juniper.contrail.pkg.models.ListSpec) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ListWidgetResponse"/>

### ListWidgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| widgets | [Widget](#github.com.Juniper.contrail.pkg.models.Widget) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.Loadbalancer"/>

### Loadbalancer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| loadbalancer_properties | [LoadbalancerType](#github.com.Juniper.contrail.pkg.models.LoadbalancerType) |  | Loadbalancer configuration like admin state, VIP, VIP subnet etc. |
| loadbalancer_provider | [string](#string) |  | Provider field selects backend provider of the LBaaS, Cloudadmin could offer different levels of service like gold, silver, bronze. Provided by HA-proxy or various HW or SW appliances in the backend. |
| service_appliance_set_refs | [LoadbalancerServiceApplianceSetRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerServiceApplianceSetRef) | repeated |  |
| virtual_machine_interface_refs | [LoadbalancerVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerVirtualMachineInterfaceRef) | repeated | Reference to the virtual machine interface for VIP, created automatically by the system. |
| service_instance_refs | [LoadbalancerServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerServiceInstanceRef) | repeated | Reference to the service instance, created automatically by the system. |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor"/>

### LoadbalancerHealthmonitor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| loadbalancer_healthmonitor_properties | [LoadbalancerHealthmonitorType](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitorType) |  | Configuration parameters for health monitor like type, method, retries etc. |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitorType"/>

### LoadbalancerHealthmonitorType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delay | [int64](#int64) |  | Time in seconds at which health check is repeated |
| expected_codes | [string](#string) |  | In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok. |
| max_retries | [int64](#int64) |  | Number of failures before declaring health bad |
| http_method | [string](#string) |  | In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc |
| admin_state | [bool](#bool) |  | Administratively up or dowm. |
| timeout | [int64](#int64) |  | Time in seconds to wait for response |
| url_path | [string](#string) |  | In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address |
| monitor_type | [string](#string) |  | Protocol used to monitor health, PING, HTTP, HTTPS or TCP |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerListener"/>

### LoadbalancerListener



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| loadbalancer_listener_properties | [LoadbalancerListenerType](#github.com.Juniper.contrail.pkg.models.LoadbalancerListenerType) |  |  |
| loadbalancer_refs | [LoadbalancerListenerLoadbalancerRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerListenerLoadbalancerRef) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerListenerLoadbalancerRef"/>

### LoadbalancerListenerLoadbalancerRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerListenerType"/>

### LoadbalancerListenerType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| default_tls_container | [string](#string) |  |  |
| protocol | [string](#string) |  |  |
| connection_limit | [int64](#int64) |  |  |
| admin_state | [bool](#bool) |  |  |
| sni_containers | [string](#string) | repeated |  |
| protocol_port | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerMember"/>

### LoadbalancerMember



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| loadbalancer_member_properties | [LoadbalancerMemberType](#github.com.Juniper.contrail.pkg.models.LoadbalancerMemberType) |  | Member configuration like ip address, destination port, weight etc. |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerMemberType"/>

### LoadbalancerMemberType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  | Operational status of the member. |
| status_description | [string](#string) |  | Operational status description of the member. |
| weight | [int64](#int64) |  | Weight for load balancing |
| admin_state | [bool](#bool) |  | Administrative up or down. |
| address | [string](#string) |  | Ip address of the member |
| protocol_port | [int64](#int64) |  | Destination port for the application on the member. |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPool"/>

### LoadbalancerPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| loadbalancer_pool_properties | [LoadbalancerPoolType](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolType) |  | Configuration for loadbalancer pool like protocol, subnet, etc. |
| loadbalancer_pool_custom_attributes | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Custom loadbalancer config, opaque to the system. Specified as list of Key:Value pairs. Applicable to LBaaS V1. |
| loadbalancer_pool_provider | [string](#string) |  | Provider field selects backend provider of the LBaaS, Cloudadmin could offer different levels of service like gold, silver, bronze. Provided by HA-proxy or various HW or SW appliances in the backend. Applicable to LBaaS V1 |
| loadbalancer_listener_refs | [LoadbalancerPoolLoadbalancerListenerRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolLoadbalancerListenerRef) | repeated | Reference to loadbalancer listener served by this pool, applicable to LBaaS V2. |
| service_instance_refs | [LoadbalancerPoolServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolServiceInstanceRef) | repeated | Reference to the service instance serving this pool, applicable to LBaaS V1. |
| loadbalancer_healthmonitor_refs | [LoadbalancerPoolLoadbalancerHealthmonitorRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolLoadbalancerHealthmonitorRef) | repeated | Reference to loadbalancer healthmonitor that this pool uses. |
| service_appliance_set_refs | [LoadbalancerPoolServiceApplianceSetRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolServiceApplianceSetRef) | repeated |  |
| virtual_machine_interface_refs | [LoadbalancerPoolVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LoadbalancerPoolVirtualMachineInterfaceRef) | repeated | Reference to the virtual machine interface reaching pool subnet, applicable to LBaaS V1. |
| loadbalancer_members | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) | repeated | Configuration object representing each member of load balancer pool. |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPoolLoadbalancerHealthmonitorRef"/>

### LoadbalancerPoolLoadbalancerHealthmonitorRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPoolLoadbalancerListenerRef"/>

### LoadbalancerPoolLoadbalancerListenerRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPoolServiceApplianceSetRef"/>

### LoadbalancerPoolServiceApplianceSetRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPoolServiceInstanceRef"/>

### LoadbalancerPoolServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPoolType"/>

### LoadbalancerPoolType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  | Operating status for this loadbalancer pool. |
| protocol | [string](#string) |  | IP protocol string like http, https or tcp. |
| subnet_id | [string](#string) |  | UUID of the subnet from where the members of the pool are reachable. |
| session_persistence | [string](#string) |  | Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE. |
| admin_state | [bool](#bool) |  | Administrative up or down |
| persistence_cookie_name | [string](#string) |  | To Be Added |
| status_description | [string](#string) |  | Operating status description for this loadbalancer pool. |
| loadbalancer_method | [string](#string) |  | Load balancing method ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerPoolVirtualMachineInterfaceRef"/>

### LoadbalancerPoolVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerServiceApplianceSetRef"/>

### LoadbalancerServiceApplianceSetRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerServiceInstanceRef"/>

### LoadbalancerServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerType"/>

### LoadbalancerType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  | Operational status of the load balancer updated by system. |
| provisioning_status | [string](#string) |  | Provisioning status of the load balancer updated by system. |
| admin_state | [bool](#bool) |  | Administrative up or down |
| vip_address | [string](#string) |  | Virtual ip for this LBaaS |
| vip_subnet_id | [string](#string) |  | Subnet UUID of the subnet of VIP, representing virtual network. |
| operating_status | [string](#string) |  | Operational status of the load balancer updated by system. |






<a name="github.com.Juniper.contrail.pkg.models.LoadbalancerVirtualMachineInterfaceRef"/>

### LoadbalancerVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LocalLinkConnection"/>

### LocalLinkConnection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| switch_id | [string](#string) |  | Switch hostname |
| port_id | [string](#string) |  | Port ID of switch where Baremetal is connected |
| switch_info | [string](#string) |  | UUID of the Physical-Port with contrail database |






<a name="github.com.Juniper.contrail.pkg.models.Location"/>

### Location



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| type | [string](#string) |  | Type of location |
| private_dns_servers | [string](#string) |  | List of DNS servers |
| private_ntp_hosts | [string](#string) |  | List of NTP sources |
| private_ospd_package_url | [string](#string) |  | Location of Contrail Networking Packages |
| private_ospd_user_name | [string](#string) |  | OSPD Non-Root User Account |
| private_ospd_user_password | [string](#string) |  | OSPD Passowrd for account |
| private_ospd_vm_disk_gb | [string](#string) |  | disk spae to assign to RH OSPD vm |
| private_ospd_vm_name | [string](#string) |  | Name of RH OSPD VM |
| private_ospd_vm_ram_mb | [string](#string) |  | ram to assign to RH OSPD vm |
| private_ospd_vm_vcpus | [string](#string) |  | vcpus to assign to RH OSPD vm |
| private_redhat_pool_id | [string](#string) |  | Repo Pool ID |
| private_redhat_subscription_key | [string](#string) |  | Subscription Activation Key |
| private_redhat_subscription_pasword | [string](#string) |  | Password for subscription account |
| private_redhat_subscription_user | [string](#string) |  | User name for RedHat subscription account |
| gcp_account_info | [string](#string) |  | copy and paste contents of account.json |
| gcp_asn | [int64](#int64) |  |  |
| gcp_region | [string](#string) |  |  |
| gcp_subnet | [string](#string) |  |  |
| aws_access_key | [string](#string) |  |  |
| aws_region | [string](#string) |  |  |
| aws_secret_key | [string](#string) |  |  |
| aws_subnet | [string](#string) |  |  |
| physical_routers | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) | repeated | Physical router location. |






<a name="github.com.Juniper.contrail.pkg.models.LogicalInterface"/>

### LogicalInterface



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| logical_interface_vlan_tag | [int64](#int64) |  | VLAN tag (.1Q) classifier for this logical interface. |
| logical_interface_type | [string](#string) |  |  |
| virtual_machine_interface_refs | [LogicalInterfaceVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LogicalInterfaceVirtualMachineInterfaceRef) | repeated | References to virtual machine interfaces that represent end points that are reachable by this logical interface. |






<a name="github.com.Juniper.contrail.pkg.models.LogicalInterfaceVirtualMachineInterfaceRef"/>

### LogicalInterfaceVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouter"/>

### LogicalRouter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| vxlan_network_identifier | [string](#string) |  | The VNI that needs to be associated with the internal VN if vxlan_routing mode is enabled. |
| configured_route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that represent this logical router, all virtual networks connected to this logical router will have this as their route target list. |
| route_target_refs | [LogicalRouterRouteTargetRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterRouteTargetRef) | repeated | Route target that represent this logical router. |
| virtual_machine_interface_refs | [LogicalRouterVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterVirtualMachineInterfaceRef) | repeated | Reference to the interface attached to this logical router. By attaching a interface to logical network all subnets in the virtual network of the interface has this router. |
| service_instance_refs | [LogicalRouterServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterServiceInstanceRef) | repeated | Reference to service instance doing SNAT functionality for external gateway. |
| route_table_refs | [LogicalRouterRouteTableRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterRouteTableRef) | repeated | Reference to the route table attached to this logical router. By attaching route table, system will create static routes with the route target only of route targets linked to this logical router |
| virtual_network_refs | [LogicalRouterVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterVirtualNetworkRef) | repeated | Reference to virtual network used as external gateway for this logical network. This link will cause a SNAT being spawned between all networks connected to logical router and external network. |
| physical_router_refs | [LogicalRouterPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterPhysicalRouterRef) | repeated | Reference to physical router, when this link is present device manager configures logical router associated route targets/interfaces on the Physical Router. |
| bgpvpn_refs | [LogicalRouterBGPVPNRef](#github.com.Juniper.contrail.pkg.models.LogicalRouterBGPVPNRef) | repeated | Back reference to logical router associated to the BGP VPN resource |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterBGPVPNRef"/>

### LogicalRouterBGPVPNRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterPhysicalRouterRef"/>

### LogicalRouterPhysicalRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterRouteTableRef"/>

### LogicalRouterRouteTableRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterRouteTargetRef"/>

### LogicalRouterRouteTargetRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterServiceInstanceRef"/>

### LogicalRouterServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterVirtualMachineInterfaceRef"/>

### LogicalRouterVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.LogicalRouterVirtualNetworkRef"/>

### LogicalRouterVirtualNetworkRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.MACLimitControlType"/>

### MACLimitControlType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mac_limit | [int64](#int64) |  | Number of MACs that can be learnt |
| mac_limit_action | [string](#string) |  | Action to be taken when MAC limit exceeds |






<a name="github.com.Juniper.contrail.pkg.models.MACMoveLimitControlType"/>

### MACMoveLimitControlType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mac_move_time_window | [int64](#int64) |  | MAC move time window |
| mac_move_limit | [int64](#int64) |  | Number of MAC moves permitted in mac move time window |
| mac_move_limit_action | [string](#string) |  | Action to be taken when MAC move limit exceeds |






<a name="github.com.Juniper.contrail.pkg.models.MacAddressesType"/>

### MacAddressesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mac_address | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.MatchConditionType"/>

### MatchConditionType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| src_port | [PortType](#github.com.Juniper.contrail.pkg.models.PortType) |  | Range of source port for layer 4 protocol |
| src_address | [AddressType](#github.com.Juniper.contrail.pkg.models.AddressType) |  | Source ip matching criteria |
| ethertype | [string](#string) |  |  |
| dst_address | [AddressType](#github.com.Juniper.contrail.pkg.models.AddressType) |  | Destination ip matching criteria |
| dst_port | [PortType](#github.com.Juniper.contrail.pkg.models.PortType) |  | Range of destination port for layer 4 protocol |
| protocol | [string](#string) |  | Layer 4 protocol in ip packet |






<a name="github.com.Juniper.contrail.pkg.models.MemberType"/>

### MemberType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [string](#string) |  | User role for the project |






<a name="github.com.Juniper.contrail.pkg.models.MirrorActionType"/>

### MirrorActionType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nic_assisted_mirroring_vlan | [int64](#int64) |  | The VLAN to be tagged on the traffic for NIC to Mirror |
| analyzer_name | [string](#string) |  | Name of service instance used as analyzer |
| nh_mode | [string](#string) |  | This mode used to determine static or dynamic nh |
| juniper_header | [bool](#bool) |  | This flag is used to determine with/without juniper-header |
| udp_port | [int64](#int64) |  | ip udp port used in contrail default encapsulation for mirroring |
| routing_instance | [string](#string) |  | Internal use only, should be set to -1 |
| static_nh_header | [StaticMirrorNhType](#github.com.Juniper.contrail.pkg.models.StaticMirrorNhType) |  | vtep details required if static nh enabled |
| analyzer_ip_address | [string](#string) |  | ip address of interface to which mirrored packets are sent |
| encapsulation | [string](#string) |  | Encapsulation for Mirrored packet, not used currently |
| analyzer_mac_address | [string](#string) |  | mac address of interface to which mirrored packets are sent |
| nic_assisted_mirroring | [bool](#bool) |  | This flag is used to select nic assisted mirroring |






<a name="github.com.Juniper.contrail.pkg.models.Namespace"/>

### Namespace



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| namespace_cidr | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | All networks in this namespace belong to this list of Prefixes. Not implemented. |






<a name="github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig"/>

### NetworkDeviceConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| physical_router_refs | [NetworkDeviceConfigPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfigPhysicalRouterRef) | repeated | Network device config of a physical router. |






<a name="github.com.Juniper.contrail.pkg.models.NetworkDeviceConfigPhysicalRouterRef"/>

### NetworkDeviceConfigPhysicalRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.NetworkIpam"/>

### NetworkIpam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| network_ipam_mgmt | [IpamType](#github.com.Juniper.contrail.pkg.models.IpamType) |  | Network IP Address Management configuration. |
| ipam_subnets | [IpamSubnets](#github.com.Juniper.contrail.pkg.models.IpamSubnets) |  | List of subnets for this ipam. |
| ipam_subnet_method | [string](#string) |  | Subnet method configuration for ipam, user can configure user-defined, flat or auto. |
| virtual_DNS_refs | [NetworkIpamVirtualDNSRef](#github.com.Juniper.contrail.pkg.models.NetworkIpamVirtualDNSRef) | repeated | Reference to virtual DNS used by this IPAM. |






<a name="github.com.Juniper.contrail.pkg.models.NetworkIpamVirtualDNSRef"/>

### NetworkIpamVirtualDNSRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.NetworkPolicy"/>

### NetworkPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| network_policy_entries | [PolicyEntriesType](#github.com.Juniper.contrail.pkg.models.PolicyEntriesType) |  | Network policy rule entries. |






<a name="github.com.Juniper.contrail.pkg.models.Node"/>

### Node



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| hostname | [string](#string) |  | Fully qualified host name |
| ip_address | [string](#string) |  | IP Address |
| mac_address | [string](#string) |  | Provisioning Interface MAC Address |
| type | [string](#string) |  | Type of machine resource |
| password | [string](#string) |  | UserPassword |
| ssh_key | [string](#string) |  | SSH Public Key |
| username | [string](#string) |  | User Name |
| aws_ami | [string](#string) |  |  |
| aws_instance_type | [string](#string) |  |  |
| gcp_image | [string](#string) |  |  |
| gcp_machine_type | [string](#string) |  |  |
| private_machine_properties | [string](#string) |  | Machine Properties from ironic |
| private_machine_state | [string](#string) |  | Machine State |
| driver_info | [DriverInfo](#github.com.Juniper.contrail.pkg.models.DriverInfo) |  | Details of the driver for power management |
| bm_properties | [BaremetalProperties](#github.com.Juniper.contrail.pkg.models.BaremetalProperties) |  | Details of baremetal hardware for scheduler |
| keypair_refs | [NodeKeypairRef](#github.com.Juniper.contrail.pkg.models.NodeKeypairRef) | repeated | Reference to keypair object to import. |
| ports | [Port](#github.com.Juniper.contrail.pkg.models.Port) | repeated | Parent of this port. |






<a name="github.com.Juniper.contrail.pkg.models.NodeKeypairRef"/>

### NodeKeypairRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.OpenStackAddress"/>

### OpenStackAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addr | [string](#string) |  | IP Address given to the instance |






<a name="github.com.Juniper.contrail.pkg.models.OpenStackFlavorProperty"/>

### OpenStackFlavorProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | UUID of the flavor used to boot server instance or empty |
| links | [OpenStackLink](#github.com.Juniper.contrail.pkg.models.OpenStackLink) |  | links for the flavor used to boot server instance |






<a name="github.com.Juniper.contrail.pkg.models.OpenStackImageProperty"/>

### OpenStackImageProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | UUID of the image for server instance |
| links | [OpenStackLink](#github.com.Juniper.contrail.pkg.models.OpenStackLink) |  | links for the image for server instance |






<a name="github.com.Juniper.contrail.pkg.models.OpenStackLink"/>

### OpenStackLink



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| href | [string](#string) |  | Link of the resource |
| rel | [string](#string) |  | Type of the link |
| type | [string](#string) |  | provides a hint as to the type of representation to expect when following the link |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackComputeNode"/>

### OpenstackComputeNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [OpenstackComputeNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNodeNodeRef) | repeated | Reference to node object for this openstack_compute node. |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackComputeNodeNodeRef"/>

### OpenstackComputeNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackControlNode"/>

### OpenstackControlNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [OpenstackControlNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackControlNodeNodeRef) | repeated | Reference to node object for this openstack_control node. |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackControlNodeNodeRef"/>

### OpenstackControlNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode"/>

### OpenstackMonitoringNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [OpenstackMonitoringNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNodeNodeRef) | repeated | Reference to node object for this openstack_monitoring node. |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNodeNodeRef"/>

### OpenstackMonitoringNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode"/>

### OpenstackNetworkNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [OpenstackNetworkNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNodeNodeRef) | repeated | Reference to node object for this openstack_network node. |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackNetworkNodeNodeRef"/>

### OpenstackNetworkNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackStorageNode"/>

### OpenstackStorageNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| node_refs | [OpenstackStorageNodeNodeRef](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNodeNodeRef) | repeated | Reference to node object for this openstack_storage node. |






<a name="github.com.Juniper.contrail.pkg.models.OpenstackStorageNodeNodeRef"/>

### OpenstackStorageNodeNodeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.OsImage"/>

### OsImage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| name | [string](#string) |  | Name of the image to be created/updated |
| owner | [string](#string) |  | An identifier for the owner of the image |
| id | [string](#string) |  | A unique, user-defined image UUID, in the format &amp;quot;nnnnnnnn-nnnn-nnnn-nnnn-nnnnnnnnnnnn&amp;quot;, Where n is a hexadecimal digit from 0 to f, or F. |
| size | [int64](#int64) |  | The size of the image data, in bytes |
| status | [string](#string) |  | The image status |
| location | [string](#string) |  | The URL to access the image file from the external store |
| file | [string](#string) |  | Abosolute path of file to be used for creating image |
| checksum | [string](#string) |  | Hash that is used over the image data |
| created_at | [string](#string) |  | The UTC date and time when the resource was created, ISO 8601 format |
| updated_at | [string](#string) |  | The UTC date and time when the resource was created, ISO 8601 format |
| container_format | [string](#string) |  |  |
| disk_format | [string](#string) |  |  |
| protected | [bool](#bool) |  |  |
| visibility | [string](#string) |  | Visibility for this image |
| property | [string](#string) |  | Additional properties of the image (name=value pair) |
| min_disk | [int64](#int64) |  | Amount of disk space in GB that is required to boot the image |
| min_ram | [int64](#int64) |  | Amount of RAM in MB that is required to boot the image |
| tags | [string](#string) |  | List of tags for this image |






<a name="github.com.Juniper.contrail.pkg.models.PeeringPolicy"/>

### PeeringPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| peering_service | [string](#string) |  | Peering policy service type. |






<a name="github.com.Juniper.contrail.pkg.models.PermType"/>

### PermType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |
| owner_access | [int64](#int64) |  |  |
| other_access | [int64](#int64) |  |  |
| group | [string](#string) |  |  |
| group_access | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.PermType2"/>

### PermType2



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | Owner tenant of the object |
| owner_access | [int64](#int64) |  | Owner permissions of the object |
| global_access | [int64](#int64) |  | Globally(others) shared object and permissions for others of the object |
| share | [ShareType](#github.com.Juniper.contrail.pkg.models.ShareType) | repeated | Selectively shared object, List of (tenant, permissions) |






<a name="github.com.Juniper.contrail.pkg.models.PhysicalInterface"/>

### PhysicalInterface



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| ethernet_segment_identifier | [string](#string) |  | Ethernet Segment Id configured for the Physical Interface. In a multihomed environment, user should configure the peer Physical interface with the same ESI. |
| physical_interface_refs | [PhysicalInterfacePhysicalInterfaceRef](#github.com.Juniper.contrail.pkg.models.PhysicalInterfacePhysicalInterfaceRef) | repeated | Reference to the other physical interface that is connected to this physical interface. |
| logical_interfaces | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) | repeated | Logical interfaces on physical interface on physical routers. |






<a name="github.com.Juniper.contrail.pkg.models.PhysicalInterfacePhysicalInterfaceRef"/>

### PhysicalInterfacePhysicalInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.PhysicalRouter"/>

### PhysicalRouter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| physical_router_management_ip | [string](#string) |  | Management ip for this physical router. It is used by the device manager to perform netconf and by SNMP collector if enabled. |
| physical_router_snmp_credentials | [SNMPCredentials](#github.com.Juniper.contrail.pkg.models.SNMPCredentials) |  | SNMP credentials for the physical router used by SNMP collector. |
| physical_router_role | [string](#string) |  | Physical router role (e.g spine or leaf), used by the device manager to provision physical router, for e.g device manager may choose to configure physical router based on its role. |
| physical_router_user_credentials | [UserCredentials](#github.com.Juniper.contrail.pkg.models.UserCredentials) |  | Username and password for netconf to the physical router by device manager. |
| physical_router_vendor_name | [string](#string) |  | Vendor name of the physical router (e.g juniper). Used by the device manager to select driver. |
| physical_router_vnc_managed | [bool](#bool) |  | This physical router is enabled to be configured by device manager. |
| physical_router_product_name | [string](#string) |  | Model name of the physical router (e.g juniper). Used by the device manager to select driver. |
| physical_router_lldp | [bool](#bool) |  | LLDP support on this router |
| physical_router_loopback_ip | [string](#string) |  | This is ip address of loopback interface of physical router. Used by the device manager to configure physical router loopback interface. |
| physical_router_image_uri | [string](#string) |  | Physical router OS image uri |
| telemetry_info | [TelemetryStateInfo](#github.com.Juniper.contrail.pkg.models.TelemetryStateInfo) |  | Telemetry info of router. Check TelemetryStateInfo |
| physical_router_snmp | [bool](#bool) |  | SNMP support on this router |
| physical_router_dataplane_ip | [string](#string) |  | This is ip address in the ip-fabric(underlay) network that can be used in data plane by physical router. Usually it is the VTEP address in VxLAN for the TOR switch. |
| physical_router_junos_service_ports | [JunosServicePorts](#github.com.Juniper.contrail.pkg.models.JunosServicePorts) |  | Juniper JUNOS specific service interfaces name to perform services like NAT. |
| virtual_network_refs | [PhysicalRouterVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.PhysicalRouterVirtualNetworkRef) | repeated | Reference to virtual network, whose VRF is present on this physical router, Applicable when only VRF is present with no physical interfaces from this physical vrouter. Generally used when using device manager and option A&#43;B for this virtual network in L3VPN use case. |
| bgp_router_refs | [PhysicalRouterBGPRouterRef](#github.com.Juniper.contrail.pkg.models.PhysicalRouterBGPRouterRef) | repeated | Reference to BGP peer representing this physical router. |
| virtual_router_refs | [PhysicalRouterVirtualRouterRef](#github.com.Juniper.contrail.pkg.models.PhysicalRouterVirtualRouterRef) | repeated | Reference to vrouter responsible for this physical router. Currently only applicable for vrouters that are TOR agents. |
| logical_interfaces | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) | repeated | Logical interfaces on physical routers. |
| physical_interfaces | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) | repeated | Physical interfaces on physical routers. |






<a name="github.com.Juniper.contrail.pkg.models.PhysicalRouterBGPRouterRef"/>

### PhysicalRouterBGPRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.PhysicalRouterVirtualNetworkRef"/>

### PhysicalRouterVirtualNetworkRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.PhysicalRouterVirtualRouterRef"/>

### PhysicalRouterVirtualRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.PluginProperties"/>

### PluginProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plugin_property | [PluginProperty](#github.com.Juniper.contrail.pkg.models.PluginProperty) | repeated | List of plugin specific properties (property, value) |






<a name="github.com.Juniper.contrail.pkg.models.PluginProperty"/>

### PluginProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| property | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.PolicyBasedForwardingRuleType"/>

### PolicyBasedForwardingRuleType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dst_mac | [string](#string) |  |  |
| protocol | [string](#string) |  |  |
| ipv6_service_chain_address | [string](#string) |  |  |
| direction | [string](#string) |  |  |
| mpls_label | [int64](#int64) |  |  |
| vlan_tag | [int64](#int64) |  |  |
| src_mac | [string](#string) |  |  |
| service_chain_address | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.PolicyEntriesType"/>

### PolicyEntriesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_rule | [PolicyRuleType](#github.com.Juniper.contrail.pkg.models.PolicyRuleType) | repeated | List of policy rules |






<a name="github.com.Juniper.contrail.pkg.models.PolicyManagement"/>

### PolicyManagement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| address_groups | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) | repeated | Address Group object |
| application_policy_sets | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) | repeated | Application-policy object defining policies to apply for a given application tag |
| firewall_policys | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) | repeated | firewall-policy object consisting of one or more firewall rules |
| firewall_rules | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) | repeated | Firewall-rule object |
| service_groups | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) | repeated | Service Group object |






<a name="github.com.Juniper.contrail.pkg.models.PolicyRuleType"/>

### PolicyRuleType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| direction | [string](#string) |  |  |
| protocol | [string](#string) |  | Layer 4 protocol in ip packet |
| dst_addresses | [AddressType](#github.com.Juniper.contrail.pkg.models.AddressType) | repeated | Destination ip matching criteria |
| action_list | [ActionListType](#github.com.Juniper.contrail.pkg.models.ActionListType) |  | Actions to be performed if packets match condition |
| created | [string](#string) |  | timestamp when security group rule object gets created |
| rule_uuid | [string](#string) |  | Rule UUID is identifier used in flow records to identify rule |
| dst_ports | [PortType](#github.com.Juniper.contrail.pkg.models.PortType) | repeated | Range of destination port for layer 4 protocol |
| application | [string](#string) | repeated | Optionally application can be specified instead of protocol and port. not currently implemented |
| last_modified | [string](#string) |  | timestamp when security group rule object gets updated |
| ethertype | [string](#string) |  |  |
| src_addresses | [AddressType](#github.com.Juniper.contrail.pkg.models.AddressType) | repeated | Source ip matching criteria |
| rule_sequence | [SequenceType](#github.com.Juniper.contrail.pkg.models.SequenceType) |  | Deprecated, Will be removed because rules themselves are already an ordered list |
| src_ports | [PortType](#github.com.Juniper.contrail.pkg.models.PortType) | repeated | Range of source port for layer 4 protocol |






<a name="github.com.Juniper.contrail.pkg.models.Port"/>

### Port



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| mac_address | [string](#string) |  | Mac Address of the NIC in Node |
| node_uuid | [string](#string) |  | UUID of the parent node where this port is connected |
| pxe_enabled | [bool](#bool) |  | Indicates whether PXE is enabled or disabled on the Port. |
| local_link_connection | [LocalLinkConnection](#github.com.Juniper.contrail.pkg.models.LocalLinkConnection) |  | The Port binding profile |






<a name="github.com.Juniper.contrail.pkg.models.PortMap"/>

### PortMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| src_port | [int64](#int64) |  |  |
| protocol | [string](#string) |  |  |
| dst_port | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.PortMappings"/>

### PortMappings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_mappings | [PortMap](#github.com.Juniper.contrail.pkg.models.PortMap) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.PortTuple"/>

### PortTuple



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |






<a name="github.com.Juniper.contrail.pkg.models.PortType"/>

### PortType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| end_port | [int64](#int64) |  |  |
| start_port | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.Project"/>

### Project



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| vxlan_routing | [bool](#bool) |  | When this knob is enabled for a project, an internal system VN (VN-Int) is created for every logical router in the project. |
| alarm_enable | [bool](#bool) |  | Flag to enable/disable alarms configured under global-system-config. True, if not set. |
| quota | [QuotaType](#github.com.Juniper.contrail.pkg.models.QuotaType) |  | Max instances limits for various objects under project. |
| floating_ip_pool_refs | [ProjectFloatingIPPoolRef](#github.com.Juniper.contrail.pkg.models.ProjectFloatingIPPoolRef) | repeated | Reference to floating ip pool in this project. |
| alias_ip_pool_refs | [ProjectAliasIPPoolRef](#github.com.Juniper.contrail.pkg.models.ProjectAliasIPPoolRef) | repeated | Reference to alias ip pool in this project. |
| namespace_refs | [ProjectNamespaceRef](#github.com.Juniper.contrail.pkg.models.ProjectNamespaceRef) | repeated | Reference to network namespace of this project. |
| application_policy_set_refs | [ProjectApplicationPolicySetRef](#github.com.Juniper.contrail.pkg.models.ProjectApplicationPolicySetRef) | repeated | Reference to default application-policy-set is automatically createdby system for default socped application policy sets. Needed by vrouter to identify default application-policy-set rules of a virtual machine interface |
| address_groups | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) | repeated | Project level address Group object |
| alarms | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) | repeated | List of alarms that are applicable to objects anchored under the project. |
| api_access_lists | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) | repeated | API access list is list of rules that define role based access to each API and its properties at project level. |
| application_policy_sets | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) | repeated | Project level application-policy object defining policies to apply for a given application tag |
| bgp_as_a_services | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) | repeated | BGP as service object represents BGP peer in the virtual network that can participate in dynamic routing with implicit default gateway of the virtual network. |
| bgpvpns | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) | repeated | BGP VPN resource contains a set of parameters for a BGP-based VPN |
| firewall_policys | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) | repeated | Project level firewall-policy object consisting of one or more firewall rules |
| firewall_rules | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) | repeated | Project level firewall-rule object |
| interface_route_tables | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) | repeated | Interface route table is mechanism to add static routes pointing to this interface. |
| loadbalancer_healthmonitors | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) | repeated | Health monitor objects is configuration to monitor health of individual pool members. |
| loadbalancer_listeners | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) | repeated | Listener represents the application(protocol, port) to be load balanced. |
| loadbalancer_pools | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) | repeated | Loadbalancer pool object represent set(pool) member servers which needs load balancing. |
| loadbalancers | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) | repeated | Loadbalancer object represents a LBaaS instance. One single Virtual IP and multiple (listeners, pools). Applicable to LBaaS V2. |
| logical_routers | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) | repeated | Logical router is a mechanism to connect multiple virtual network as they have been connected by a router. |
| network_ipams | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) | repeated | IP Address Management object that controls, ip allocation, DNS and DHCP |
| network_policys | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) | repeated | Network Policy is set access control rules that can be attached to virtual networks. Network ACL(s) and connectivity information is derived from Network policies that are attached to virtual networks. |
| qos_configs | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) | repeated |  |
| route_aggregates | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) | repeated | List of references of child routing route aggregate objects. automatically maintained by system. |
| route_tables | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) | repeated | Network route table is mechanism of adding static routes in the virtual network |
| routing_policys | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) | repeated | List of references of child routing policy objects. automatically maintained by system. |
| security_groups | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) | repeated | Security Groups are set of state full access control rules attached to interfaces.It can be used to implement microsegmentation. |
| security_logging_objects | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) | repeated | Security logging object configuration for specifying session logging criteria |
| service_groups | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) | repeated | Project level service Group object |
| service_health_checks | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) | repeated | Service health check is a keepalive mechanism for the virtual machine interface. Liveliness of the interface is determined based on configuration in the service health check. It is mainly designed for service instance interfaces. However it will work with any interface which present on contrail vrouter. |
| service_instances | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) | repeated | Service instance represents logical instance service used in the virtual world, e.g. firewall, load balancer etc. It can represent one or multiple virtual machines or physical devices. Many service instances can share a virtual machine or physical device. |
| tags | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) | repeated | Attribute attached to objects - has a type and value |
| users | [User](#github.com.Juniper.contrail.pkg.models.User) | repeated | Reference to a project of this user. |
| virtual_ips | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) | repeated | Virtual ip object application(protocol, port). Applicable only to LBaaS V1 |
| virtual_machine_interfaces | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) | repeated | Virtual machine interface represent a interface(port) into virtual network. It may or may not have corresponding virtual machine. A virtual machine interface has atleast a MAC address and Ip address. |
| virtual_networks | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) | repeated | Virtual network is collection of end points (interface or ip(s) or MAC(s)) that can talk to each other by default. It is collection of subnets connected by implicit router which default gateway in each subnet. |






<a name="github.com.Juniper.contrail.pkg.models.ProjectAliasIPPoolRef"/>

### ProjectAliasIPPoolRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ProjectApplicationPolicySetRef"/>

### ProjectApplicationPolicySetRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ProjectFloatingIPPoolRef"/>

### ProjectFloatingIPPoolRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ProjectNamespaceRef"/>

### ProjectNamespaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ProtocolType"/>

### ProtocolType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protocol | [string](#string) |  |  |
| port | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ProviderAttachment"/>

### ProviderAttachment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_router_refs | [ProviderAttachmentVirtualRouterRef](#github.com.Juniper.contrail.pkg.models.ProviderAttachmentVirtualRouterRef) | repeated | Not in Use. |






<a name="github.com.Juniper.contrail.pkg.models.ProviderAttachmentVirtualRouterRef"/>

### ProviderAttachmentVirtualRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ProviderDetails"/>

### ProviderDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| segmentation_id | [int64](#int64) |  |  |
| physical_network | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.QosConfig"/>

### QosConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| qos_config_type | [string](#string) |  | Specifies if qos-config is for vhost, fabric or for project. |
| mpls_exp_entries | [QosIdForwardingClassPairs](#github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPairs) |  | Map of MPLS EXP values to applicable forwarding class for packet. |
| vlan_priority_entries | [QosIdForwardingClassPairs](#github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPairs) |  | Map of .1p priority code to applicable forwarding class for packet. |
| default_forwarding_class_id | [int64](#int64) |  | Default forwarding class used for all non-specified QOS bits |
| dscp_entries | [QosIdForwardingClassPairs](#github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPairs) |  | Map of DSCP match condition and applicable forwarding class for packet. |
| global_system_config_refs | [QosConfigGlobalSystemConfigRef](#github.com.Juniper.contrail.pkg.models.QosConfigGlobalSystemConfigRef) | repeated | This link is internally created and may be removed in future. End users should not set this link or assume it to be there |






<a name="github.com.Juniper.contrail.pkg.models.QosConfigGlobalSystemConfigRef"/>

### QosConfigGlobalSystemConfigRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPair"/>

### QosIdForwardingClassPair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  | QoS bit value (DSCP or Vlan priority or EXP bit value |
| forwarding_class_id | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPairs"/>

### QosIdForwardingClassPairs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_id_forwarding_class_pair | [QosIdForwardingClassPair](#github.com.Juniper.contrail.pkg.models.QosIdForwardingClassPair) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.QosQueue"/>

### QosQueue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| qos_queue_identifier | [int64](#int64) |  | Unique id for this queue. |
| max_bandwidth | [int64](#int64) |  | Maximum bandwidth for this queue. |
| min_bandwidth | [int64](#int64) |  | Minimum bandwidth for this queue. |






<a name="github.com.Juniper.contrail.pkg.models.QuotaType"/>

### QuotaType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_router | [int64](#int64) |  | Maximum number of logical routers |
| network_policy | [int64](#int64) |  | Maximum number of network policies |
| loadbalancer_pool | [int64](#int64) |  | Maximum number of loadbalancer pools |
| route_table | [int64](#int64) |  | Maximum number of route tables |
| subnet | [int64](#int64) |  | Maximum number of subnets |
| network_ipam | [int64](#int64) |  | Maximum number of network IPAMs |
| virtual_DNS_record | [int64](#int64) |  | Maximum number of virtual DNS records |
| logical_router | [int64](#int64) |  | Maximum number of logical routers |
| security_group_rule | [int64](#int64) |  | Maximum number of security group rules |
| virtual_DNS | [int64](#int64) |  | Maximum number of virtual DNS servers |
| service_instance | [int64](#int64) |  | Maximum number of service instances |
| service_template | [int64](#int64) |  | Maximum number of service templates |
| bgp_router | [int64](#int64) |  | Maximum number of bgp routers |
| floating_ip | [int64](#int64) |  | Maximum number of floating ips |
| floating_ip_pool | [int64](#int64) |  | Maximum number of floating ip pools |
| loadbalancer_member | [int64](#int64) |  | Maximum number of loadbalancer member |
| access_control_list | [int64](#int64) |  | Maximum number of access control lists |
| virtual_machine_interface | [int64](#int64) |  | Maximum number of virtual machine interfaces |
| instance_ip | [int64](#int64) |  | Maximum number of instance ips |
| global_vrouter_config | [int64](#int64) |  | Maximum number of global vrouter configs |
| security_logging_object | [int64](#int64) |  | Maximum number of security logging objects |
| loadbalancer_healthmonitor | [int64](#int64) |  | Maximum number of loadbalancer health monitors |
| virtual_ip | [int64](#int64) |  | Maximum number of virtual ips |
| defaults | [int64](#int64) |  | Need to clarify |
| security_group | [int64](#int64) |  | Maximum number of security groups |
| virtual_network | [int64](#int64) |  | Maximum number of virtual networks |






<a name="github.com.Juniper.contrail.pkg.models.RbacPermType"/>

### RbacPermType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_crud | [string](#string) |  | String CRUD representing permissions for C=create, R=read, U=update, D=delete. |
| role_name | [string](#string) |  | Name of the role |






<a name="github.com.Juniper.contrail.pkg.models.RbacRuleEntriesType"/>

### RbacRuleEntriesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rbac_rule | [RbacRuleType](#github.com.Juniper.contrail.pkg.models.RbacRuleType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.RbacRuleType"/>

### RbacRuleType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule_object | [string](#string) |  | Name of the REST API (object) for this rule, * represent all objects |
| rule_perms | [RbacPermType](#github.com.Juniper.contrail.pkg.models.RbacPermType) | repeated | List of [(role, permissions),...] |
| rule_field | [string](#string) |  | Name of the level one field (property) for this object, * represent all properties |






<a name="github.com.Juniper.contrail.pkg.models.RouteAggregate"/>

### RouteAggregate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_instance_refs | [RouteAggregateServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.RouteAggregateServiceInstanceRef) | repeated | Reference to route-aggregate policy attached to (service instance, interface). |






<a name="github.com.Juniper.contrail.pkg.models.RouteAggregateServiceInstanceRef"/>

### RouteAggregateServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [ServiceInterfaceTag](#github.com.Juniper.contrail.pkg.models.ServiceInterfaceTag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.RouteTable"/>

### RouteTable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  | Routes in the route table are configured in following way. |






<a name="github.com.Juniper.contrail.pkg.models.RouteTableType"/>

### RouteTableType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [RouteType](#github.com.Juniper.contrail.pkg.models.RouteType) | repeated | List of ip routes with following fields. |






<a name="github.com.Juniper.contrail.pkg.models.RouteTarget"/>

### RouteTarget



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |






<a name="github.com.Juniper.contrail.pkg.models.RouteTargetList"/>

### RouteTargetList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_target | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.RouteType"/>

### RouteType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prefix | [string](#string) |  | Ip prefix/len format prefix |
| next_hop | [string](#string) |  | Ip address or service instance name. |
| community_attributes | [CommunityAttributes](#github.com.Juniper.contrail.pkg.models.CommunityAttributes) |  |  |
| next_hop_type | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.RoutingInstance"/>

### RoutingInstance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |






<a name="github.com.Juniper.contrail.pkg.models.RoutingPolicy"/>

### RoutingPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_instance_refs | [RoutingPolicyServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.RoutingPolicyServiceInstanceRef) | repeated | Reference to routing policy attached to (service instance, interface). |






<a name="github.com.Juniper.contrail.pkg.models.RoutingPolicyServiceInstanceRef"/>

### RoutingPolicyServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [RoutingPolicyServiceInstanceType](#github.com.Juniper.contrail.pkg.models.RoutingPolicyServiceInstanceType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.RoutingPolicyServiceInstanceType"/>

### RoutingPolicyServiceInstanceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| right_sequence | [string](#string) |  |  |
| left_sequence | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.SNMPCredentials"/>

### SNMPCredentials



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| v3_privacy_protocol | [string](#string) |  |  |
| retries | [int64](#int64) |  |  |
| v3_authentication_password | [string](#string) |  |  |
| v3_engine_time | [int64](#int64) |  |  |
| v3_engine_id | [string](#string) |  |  |
| local_port | [int64](#int64) |  |  |
| v3_security_level | [string](#string) |  |  |
| v3_context | [string](#string) |  |  |
| v3_security_name | [string](#string) |  |  |
| v3_authentication_protocol | [string](#string) |  |  |
| v2_community | [string](#string) |  |  |
| v3_security_engine_id | [string](#string) |  |  |
| v3_context_engine_id | [string](#string) |  |  |
| version | [int64](#int64) |  |  |
| timeout | [int64](#int64) |  |  |
| v3_privacy_password | [string](#string) |  |  |
| v3_engine_boots | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.SecurityGroup"/>

### SecurityGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| security_group_entries | [PolicyEntriesType](#github.com.Juniper.contrail.pkg.models.PolicyEntriesType) |  | Security group rule entries. |
| configured_security_group_id | [int64](#int64) |  | Unique 32 bit user defined ID assigned to this security group [1, 8M - 1]. |
| security_group_id | [int64](#int64) |  | Unique 32 bit ID automatically assigned to this security group [8M&#43;1, 32G]. |
| access_control_lists | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) | repeated | port access control list is automatically derived from all the security groups attached to port. |






<a name="github.com.Juniper.contrail.pkg.models.SecurityLoggingObject"/>

### SecurityLoggingObject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| security_logging_object_rules | [SecurityLoggingObjectRuleListType](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleListType) |  | Security logging object rules derived internally. |
| security_logging_object_rate | [int64](#int64) |  | Security logging object rate defining rate of session logging |
| security_group_refs | [SecurityLoggingObjectSecurityGroupRef](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectSecurityGroupRef) | repeated | Reference to security-group attached to this security-logging-object |
| network_policy_refs | [SecurityLoggingObjectNetworkPolicyRef](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectNetworkPolicyRef) | repeated | Reference to network-policy attached to this security-logging-object |






<a name="github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectNetworkPolicyRef"/>

### SecurityLoggingObjectNetworkPolicyRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [SecurityLoggingObjectRuleListType](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleListType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleEntryType"/>

### SecurityLoggingObjectRuleEntryType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule_uuid | [string](#string) |  | Rule UUID of network policy or security-group. When this is absent it implies all rules of security-group or network-policy |
| rate | [int64](#int64) |  | Rate at which sessions are logged. When rates are specified at multiple levels, the rate which specifies highest frequency is selected |






<a name="github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleListType"/>

### SecurityLoggingObjectRuleListType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule | [SecurityLoggingObjectRuleEntryType](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleEntryType) | repeated | List of rules along with logging rate for each rule. Both rule-uuid and rate are optional. When rule-uuid is absent then it means all rules of associated SG or network-policy |






<a name="github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectSecurityGroupRef"/>

### SecurityLoggingObjectSecurityGroupRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [SecurityLoggingObjectRuleListType](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObjectRuleListType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.SequenceType"/>

### SequenceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| major | [int64](#int64) |  |  |
| minor | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.Server"/>

### Server



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| created | [string](#string) |  | The date and time when the resource was created. The date and time stamp format is ISO 8601 |
| hostId | [string](#string) |  | An ID string representing the host |
| id | [string](#string) |  | The UUID of the server |
| name | [string](#string) |  | The UUID of the server |
| image | [OpenStackImageProperty](#github.com.Juniper.contrail.pkg.models.OpenStackImageProperty) |  | The UUID and links for the image for your server instance |
| flavor | [OpenStackFlavorProperty](#github.com.Juniper.contrail.pkg.models.OpenStackFlavorProperty) |  | The UUID and links for the flavor for your server instance |
| addresses | [OpenStackAddress](#github.com.Juniper.contrail.pkg.models.OpenStackAddress) |  | The addresses for the server |
| accessIPv4 | [string](#string) |  | IPv4 address that should be used to access this server |
| accessIPv6 | [string](#string) |  | IPv6 address that should be used to access this server |
| config_drive | [bool](#bool) |  | Indicates whether or not a config drive was used for this server |
| progress | [int64](#int64) |  | A percentage value of the build progress |
| status | [string](#string) |  | The server status |
| host_status | [string](#string) |  | The host status |
| tenant_id | [string](#string) |  | The UUID of the tenant in a multi-tenancy cloud |
| updated | [string](#string) |  | The date and time when the resource was updated. The date and time stamp format is ISO 8601 |
| user_id | [int64](#int64) |  | The user ID of the user who owns the server |
| locked | [bool](#bool) |  | True if the instance is locked otherwise False |






<a name="github.com.Juniper.contrail.pkg.models.ServiceAppliance"/>

### ServiceAppliance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_appliance_user_credentials | [UserCredentials](#github.com.Juniper.contrail.pkg.models.UserCredentials) |  | Authentication credentials for driver to access service appliance. |
| service_appliance_ip_address | [string](#string) |  | Management Ip address of the service-appliance. |
| service_appliance_properties | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | List of Key:Value pairs used by the provider driver of this service appliance. |
| physical_interface_refs | [ServiceAppliancePhysicalInterfaceRef](#github.com.Juniper.contrail.pkg.models.ServiceAppliancePhysicalInterfaceRef) | repeated | Reference to physical interface that can be used as (service interface type)left, right, management OR other. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceApplianceInterfaceType"/>

### ServiceApplianceInterfaceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_type | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceAppliancePhysicalInterfaceRef"/>

### ServiceAppliancePhysicalInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [ServiceApplianceInterfaceType](#github.com.Juniper.contrail.pkg.models.ServiceApplianceInterfaceType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceApplianceSet"/>

### ServiceApplianceSet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_appliance_set_properties | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | List of Key:Value pairs that are used by the provider driver and opaque to system. |
| service_appliance_ha_mode | [string](#string) |  | High availability mode for the service appliance set, active-active or active-backup. |
| service_appliance_driver | [string](#string) |  | Name of the provider driver for this service appliance set. |
| service_appliances | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) | repeated | Service appliance is a member in service appliance set (e.g. Loadbalancer, Firewall provider).By default system will create &amp;quot;ha-proxy&amp;quot; based service appliance. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceConnectionModule"/>

### ServiceConnectionModule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_type | [string](#string) |  | Type of service assigned for this object |
| e2_service | [string](#string) |  | E2 service type. |
| service_object_refs | [ServiceConnectionModuleServiceObjectRef](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModuleServiceObjectRef) | repeated | Links the service-connection-module to a service object. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceConnectionModuleServiceObjectRef"/>

### ServiceConnectionModuleServiceObjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceEndpoint"/>

### ServiceEndpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_connection_module_refs | [ServiceEndpointServiceConnectionModuleRef](#github.com.Juniper.contrail.pkg.models.ServiceEndpointServiceConnectionModuleRef) | repeated | Link the service endpoint to service connection node |
| physical_router_refs | [ServiceEndpointPhysicalRouterRef](#github.com.Juniper.contrail.pkg.models.ServiceEndpointPhysicalRouterRef) | repeated | Reference to Router for a given service. |
| service_object_refs | [ServiceEndpointServiceObjectRef](#github.com.Juniper.contrail.pkg.models.ServiceEndpointServiceObjectRef) | repeated | Links the service-endpoints to a service object. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceEndpointPhysicalRouterRef"/>

### ServiceEndpointPhysicalRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceEndpointServiceConnectionModuleRef"/>

### ServiceEndpointServiceConnectionModuleRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceEndpointServiceObjectRef"/>

### ServiceEndpointServiceObjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceGroup"/>

### ServiceGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_group_firewall_service_list | [FirewallServiceGroupType](#github.com.Juniper.contrail.pkg.models.FirewallServiceGroupType) |  | list of service objects (protocol, source port and destination port |






<a name="github.com.Juniper.contrail.pkg.models.ServiceHealthCheck"/>

### ServiceHealthCheck



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_health_check_properties | [ServiceHealthCheckType](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheckType) |  | Service health check has following fields. |
| service_instance_refs | [ServiceHealthCheckServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheckServiceInstanceRef) | repeated | Reference to service instance using this service health check. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceHealthCheckServiceInstanceRef"/>

### ServiceHealthCheckServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [ServiceInterfaceTag](#github.com.Juniper.contrail.pkg.models.ServiceInterfaceTag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceHealthCheckType"/>

### ServiceHealthCheckType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delayUsecs | [int64](#int64) |  | Time in micro seconds at which health check is repeated |
| timeoutUsecs | [int64](#int64) |  | Time in micro seconds to wait for response |
| enabled | [bool](#bool) |  | Administratively enable or disable this health check. |
| delay | [int64](#int64) |  | Time in seconds at which health check is repeated |
| expected_codes | [string](#string) |  | In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok. |
| max_retries | [int64](#int64) |  | Number of failures before declaring health bad |
| health_check_type | [string](#string) |  | Health check type, currently only link-local, end-to-end and segment are supported |
| http_method | [string](#string) |  | In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc |
| timeout | [int64](#int64) |  | Time in seconds to wait for response |
| url_path | [string](#string) |  | In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address |
| monitor_type | [string](#string) |  | Protocol used to monitor health, currently only HTTP, ICMP(ping), and BFD are supported |






<a name="github.com.Juniper.contrail.pkg.models.ServiceInstance"/>

### ServiceInstance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_instance_bindings | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Opaque key value pair for generating config for the service instance. |
| service_instance_properties | [ServiceInstanceType](#github.com.Juniper.contrail.pkg.models.ServiceInstanceType) |  | Service instance configuration parameters. |
| service_template_refs | [ServiceInstanceServiceTemplateRef](#github.com.Juniper.contrail.pkg.models.ServiceInstanceServiceTemplateRef) | repeated | Reference to the service template of this service instance. |
| instance_ip_refs | [ServiceInstanceInstanceIPRef](#github.com.Juniper.contrail.pkg.models.ServiceInstanceInstanceIPRef) | repeated | Reference to ip address, which is used as nexthop pointing to (service instance, service interface). |
| port_tuples | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceInstanceInstanceIPRef"/>

### ServiceInstanceInstanceIPRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [ServiceInterfaceTag](#github.com.Juniper.contrail.pkg.models.ServiceInterfaceTag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceInstanceInterfaceType"/>

### ServiceInstanceInterfaceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_network | [string](#string) |  | Interface belongs to this virtual network. |
| ip_address | [string](#string) |  | Shared ip for this interface (Only V1) |
| allowed_address_pairs | [AllowedAddressPairs](#github.com.Juniper.contrail.pkg.models.AllowedAddressPairs) |  | Allowed address pairs, list of (IP address, MAC) for this interface |
| static_routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  | Static routes for this interface (Only V1) |






<a name="github.com.Juniper.contrail.pkg.models.ServiceInstanceServiceTemplateRef"/>

### ServiceInstanceServiceTemplateRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceInstanceType"/>

### ServiceInstanceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| right_virtual_network | [string](#string) |  | Deprecated |
| right_ip_address | [string](#string) |  | Deprecated |
| availability_zone | [string](#string) |  | Availability zone used to spawn VM(s) for this service instance, used in version 1 (V1) only |
| management_virtual_network | [string](#string) |  | Deprecated |
| scale_out | [ServiceScaleOutType](#github.com.Juniper.contrail.pkg.models.ServiceScaleOutType) |  | Number of virtual machines in this service instance, used in version 1 (V1) only |
| ha_mode | [string](#string) |  | When scale-out is greater than one, decides if active-active or active-backup, used in version 1 (V1) only |
| virtual_router_id | [string](#string) |  | UUID of a virtual-router on which this service instance need to spawn. Used to spawn services on CPE device when Nova is not present |
| interface_list | [ServiceInstanceInterfaceType](#github.com.Juniper.contrail.pkg.models.ServiceInstanceInterfaceType) | repeated | List of service instance interface properties. Ordered list as per service template |
| left_ip_address | [string](#string) |  | Deprecated |
| left_virtual_network | [string](#string) |  | Deprecated |
| auto_policy | [bool](#bool) |  | Set when system creates internal service chains, example SNAT with router external flag in logical router |






<a name="github.com.Juniper.contrail.pkg.models.ServiceInterfaceTag"/>

### ServiceInterfaceTag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_type | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceObject"/>

### ServiceObject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceScaleOutType"/>

### ServiceScaleOutType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auto_scale | [bool](#bool) |  | Automatically change the number of virtual machines. Not implemented |
| max_instances | [int64](#int64) |  | Maximum number of scale out factor(virtual machines). can be changed dynamically |






<a name="github.com.Juniper.contrail.pkg.models.ServiceTemplate"/>

### ServiceTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_template_properties | [ServiceTemplateType](#github.com.Juniper.contrail.pkg.models.ServiceTemplateType) |  | Service template configuration parameters. |
| service_appliance_set_refs | [ServiceTemplateServiceApplianceSetRef](#github.com.Juniper.contrail.pkg.models.ServiceTemplateServiceApplianceSetRef) | repeated | Reference to the service appliance set represented by this service template. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceTemplateInterfaceType"/>

### ServiceTemplateInterfaceType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| static_route_enable | [bool](#bool) |  | Static routes configured required on this interface of service instance (Only V1) |
| shared_ip | [bool](#bool) |  | Shared ip is required on this interface when service instance is scaled out (Only V1) |
| service_interface_type | [string](#string) |  | Type of service interface supported by this template left, right or other. |






<a name="github.com.Juniper.contrail.pkg.models.ServiceTemplateServiceApplianceSetRef"/>

### ServiceTemplateServiceApplianceSetRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.ServiceTemplateType"/>

### ServiceTemplateType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| availability_zone_enable | [bool](#bool) |  | Enable availability zone for version 1 service instances |
| instance_data | [string](#string) |  | Opaque string (typically in json format) used to spawn a vrouter-instance. |
| ordered_interfaces | [bool](#bool) |  | Deprecated |
| service_virtualization_type | [string](#string) |  | Service virtualization type decides how individual service instances are instantiated |
| interface_type | [ServiceTemplateInterfaceType](#github.com.Juniper.contrail.pkg.models.ServiceTemplateInterfaceType) | repeated | List of interfaces which decided number of interfaces and type |
| image_name | [string](#string) |  | Glance image name for the service virtual machine, Version 1 only |
| service_mode | [string](#string) |  | Service instance mode decides how packets are forwarded across the service |
| version | [int64](#int64) |  |  |
| service_type | [string](#string) |  | Service instance mode decides how routing happens across the service |
| flavor | [string](#string) |  | Nova flavor used for service virtual machines, Version 1 only |
| service_scaling | [bool](#bool) |  | Enable scaling of service virtual machines, Version 1 only |
| vrouter_instance_type | [string](#string) |  | Mechanism used to spawn service instance, when vrouter is spawning instances.Allowed values libvirt-qemu, docker or netns |






<a name="github.com.Juniper.contrail.pkg.models.ShareType"/>

### ShareType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_access | [int64](#int64) |  | Allowed permissions in sharing |
| tenant | [string](#string) |  | Name of tenant with whom the object is shared |






<a name="github.com.Juniper.contrail.pkg.models.StaticMirrorNhType"/>

### StaticMirrorNhType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vtep_dst_ip_address | [string](#string) |  | ip address of destination vtep |
| vtep_dst_mac_address | [string](#string) |  | mac address of destination vtep |
| vni | [int64](#int64) |  | Vni of vtep |






<a name="github.com.Juniper.contrail.pkg.models.Subnet"/>

### Subnet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| subnet_ip_prefix | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) |  | Ip prefix/length of the subnet. |
| virtual_machine_interface_refs | [SubnetVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.SubnetVirtualMachineInterfaceRef) | repeated | Subnet belongs of the referenced virtual machine interface. This is used in CPE use case when a subnet is reachable via the interface. It also serves as dynamic DHCP pool for host on this LAN, where vrouter is DHCP server. |






<a name="github.com.Juniper.contrail.pkg.models.SubnetListType"/>

### SubnetListType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.SubnetType"/>

### SubnetType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip_prefix | [string](#string) |  |  |
| ip_prefix_len | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.SubnetVirtualMachineInterfaceRef"/>

### SubnetVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.Tag"/>

### Tag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| tag_type_name | [string](#string) |  | Tag type string representation |
| tag_id | [string](#string) |  | Internal Tag ID encapsulating tag type and value in hexadecimal fomat: 0xTTTTVVVV (T: type, V: value) |
| tag_value | [string](#string) |  | Tag value string representation |
| tag_type_refs | [TagTagTypeRef](#github.com.Juniper.contrail.pkg.models.TagTagTypeRef) | repeated | Tag type reference which is limited to one |






<a name="github.com.Juniper.contrail.pkg.models.TagTagTypeRef"/>

### TagTagTypeRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.TagType"/>

### TagType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| tag_type_id | [string](#string) |  | Internal Tag type ID coded on 16 bits where the first 255 IDs are reserved and pre-defined. Users (principally cloud admin) can define arbitrary types but its automatically shared to all project as it is a global resource. |






<a name="github.com.Juniper.contrail.pkg.models.TelemetryResourceInfo"/>

### TelemetryResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  |  |
| rate | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.TelemetryStateInfo"/>

### TelemetryStateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [TelemetryResourceInfo](#github.com.Juniper.contrail.pkg.models.TelemetryResourceInfo) | repeated |  |
| server_port | [int64](#int64) |  |  |
| server_ip | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.TimerType"/>

### TimerType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [string](#string) |  |  |
| off_interval | [string](#string) |  |  |
| on_interval | [string](#string) |  |  |
| end_time | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAPIAccessListRequest"/>

### UpdateAPIAccessListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_access_list | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAPIAccessListResponse"/>

### UpdateAPIAccessListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_access_list | [APIAccessList](#github.com.Juniper.contrail.pkg.models.APIAccessList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAccessControlListRequest"/>

### UpdateAccessControlListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_control_list | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAccessControlListResponse"/>

### UpdateAccessControlListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_control_list | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAddressGroupRequest"/>

### UpdateAddressGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_group | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAddressGroupResponse"/>

### UpdateAddressGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_group | [AddressGroup](#github.com.Juniper.contrail.pkg.models.AddressGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAlarmRequest"/>

### UpdateAlarmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alarm | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAlarmResponse"/>

### UpdateAlarmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alarm | [Alarm](#github.com.Juniper.contrail.pkg.models.Alarm) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAliasIPPoolRequest"/>

### UpdateAliasIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip_pool | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAliasIPPoolResponse"/>

### UpdateAliasIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip_pool | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAliasIPRequest"/>

### UpdateAliasIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAliasIPResponse"/>

### UpdateAliasIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alias_ip | [AliasIP](#github.com.Juniper.contrail.pkg.models.AliasIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAnalyticsNodeRequest"/>

### UpdateAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analytics_node | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAnalyticsNodeResponse"/>

### UpdateAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analytics_node | [AnalyticsNode](#github.com.Juniper.contrail.pkg.models.AnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAppformixNodeRequest"/>

### UpdateAppformixNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appformix_node | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateAppformixNodeResponse"/>

### UpdateAppformixNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appformix_node | [AppformixNode](#github.com.Juniper.contrail.pkg.models.AppformixNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateApplicationPolicySetRequest"/>

### UpdateApplicationPolicySetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_policy_set | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateApplicationPolicySetResponse"/>

### UpdateApplicationPolicySetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_policy_set | [ApplicationPolicySet](#github.com.Juniper.contrail.pkg.models.ApplicationPolicySet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBGPAsAServiceRequest"/>

### UpdateBGPAsAServiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_as_a_service | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBGPAsAServiceResponse"/>

### UpdateBGPAsAServiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_as_a_service | [BGPAsAService](#github.com.Juniper.contrail.pkg.models.BGPAsAService) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBGPRouterRequest"/>

### UpdateBGPRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_router | [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBGPRouterResponse"/>

### UpdateBGPRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgp_router | [BGPRouter](#github.com.Juniper.contrail.pkg.models.BGPRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBGPVPNRequest"/>

### UpdateBGPVPNRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgpvpn | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBGPVPNResponse"/>

### UpdateBGPVPNResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bgpvpn | [BGPVPN](#github.com.Juniper.contrail.pkg.models.BGPVPN) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBaremetalNodeRequest"/>

### UpdateBaremetalNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_node | [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBaremetalNodeResponse"/>

### UpdateBaremetalNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_node | [BaremetalNode](#github.com.Juniper.contrail.pkg.models.BaremetalNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBaremetalPortRequest"/>

### UpdateBaremetalPortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_port | [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBaremetalPortResponse"/>

### UpdateBaremetalPortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baremetal_port | [BaremetalPort](#github.com.Juniper.contrail.pkg.models.BaremetalPort) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBridgeDomainRequest"/>

### UpdateBridgeDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bridge_domain | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateBridgeDomainResponse"/>

### UpdateBridgeDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bridge_domain | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateConfigNodeRequest"/>

### UpdateConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_node | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateConfigNodeResponse"/>

### UpdateConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_node | [ConfigNode](#github.com.Juniper.contrail.pkg.models.ConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateConfigRootRequest"/>

### UpdateConfigRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_root | [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateConfigRootResponse"/>

### UpdateConfigRootResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_root | [ConfigRoot](#github.com.Juniper.contrail.pkg.models.ConfigRoot) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsDatabaseNodeRequest"/>

### UpdateContrailAnalyticsDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_database_node | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsDatabaseNodeResponse"/>

### UpdateContrailAnalyticsDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_database_node | [ContrailAnalyticsDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsNodeRequest"/>

### UpdateContrailAnalyticsNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_node | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailAnalyticsNodeResponse"/>

### UpdateContrailAnalyticsNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_analytics_node | [ContrailAnalyticsNode](#github.com.Juniper.contrail.pkg.models.ContrailAnalyticsNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailClusterRequest"/>

### UpdateContrailClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_cluster | [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailClusterResponse"/>

### UpdateContrailClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_cluster | [ContrailCluster](#github.com.Juniper.contrail.pkg.models.ContrailCluster) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailConfigDatabaseNodeRequest"/>

### UpdateContrailConfigDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_database_node | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailConfigDatabaseNodeResponse"/>

### UpdateContrailConfigDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_database_node | [ContrailConfigDatabaseNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigDatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailConfigNodeRequest"/>

### UpdateContrailConfigNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_node | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailConfigNodeResponse"/>

### UpdateContrailConfigNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_config_node | [ContrailConfigNode](#github.com.Juniper.contrail.pkg.models.ContrailConfigNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailControlNodeRequest"/>

### UpdateContrailControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_control_node | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailControlNodeResponse"/>

### UpdateContrailControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_control_node | [ContrailControlNode](#github.com.Juniper.contrail.pkg.models.ContrailControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailStorageNodeRequest"/>

### UpdateContrailStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_storage_node | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailStorageNodeResponse"/>

### UpdateContrailStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_storage_node | [ContrailStorageNode](#github.com.Juniper.contrail.pkg.models.ContrailStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailVrouterNodeRequest"/>

### UpdateContrailVrouterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_vrouter_node | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailVrouterNodeResponse"/>

### UpdateContrailVrouterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_vrouter_node | [ContrailVrouterNode](#github.com.Juniper.contrail.pkg.models.ContrailVrouterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailWebuiNodeRequest"/>

### UpdateContrailWebuiNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_webui_node | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateContrailWebuiNodeResponse"/>

### UpdateContrailWebuiNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contrail_webui_node | [ContrailWebuiNode](#github.com.Juniper.contrail.pkg.models.ContrailWebuiNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateCustomerAttachmentRequest"/>

### UpdateCustomerAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_attachment | [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateCustomerAttachmentResponse"/>

### UpdateCustomerAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_attachment | [CustomerAttachment](#github.com.Juniper.contrail.pkg.models.CustomerAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDashboardRequest"/>

### UpdateDashboardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dashboard | [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDashboardResponse"/>

### UpdateDashboardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dashboard | [Dashboard](#github.com.Juniper.contrail.pkg.models.Dashboard) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDatabaseNodeRequest"/>

### UpdateDatabaseNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_node | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDatabaseNodeResponse"/>

### UpdateDatabaseNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_node | [DatabaseNode](#github.com.Juniper.contrail.pkg.models.DatabaseNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDiscoveryServiceAssignmentRequest"/>

### UpdateDiscoveryServiceAssignmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| discovery_service_assignment | [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDiscoveryServiceAssignmentResponse"/>

### UpdateDiscoveryServiceAssignmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| discovery_service_assignment | [DiscoveryServiceAssignment](#github.com.Juniper.contrail.pkg.models.DiscoveryServiceAssignment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDomainRequest"/>

### UpdateDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDomainResponse"/>

### UpdateDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#github.com.Juniper.contrail.pkg.models.Domain) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDsaRuleRequest"/>

### UpdateDsaRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dsa_rule | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateDsaRuleResponse"/>

### UpdateDsaRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dsa_rule | [DsaRule](#github.com.Juniper.contrail.pkg.models.DsaRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateE2ServiceProviderRequest"/>

### UpdateE2ServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_service_provider | [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateE2ServiceProviderResponse"/>

### UpdateE2ServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_service_provider | [E2ServiceProvider](#github.com.Juniper.contrail.pkg.models.E2ServiceProvider) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFirewallPolicyRequest"/>

### UpdateFirewallPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_policy | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFirewallPolicyResponse"/>

### UpdateFirewallPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_policy | [FirewallPolicy](#github.com.Juniper.contrail.pkg.models.FirewallPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFirewallRuleRequest"/>

### UpdateFirewallRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_rule | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFirewallRuleResponse"/>

### UpdateFirewallRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firewall_rule | [FirewallRule](#github.com.Juniper.contrail.pkg.models.FirewallRule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFlavorRequest"/>

### UpdateFlavorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flavor | [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFlavorResponse"/>

### UpdateFlavorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flavor | [Flavor](#github.com.Juniper.contrail.pkg.models.Flavor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFloatingIPPoolRequest"/>

### UpdateFloatingIPPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_pool | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFloatingIPPoolResponse"/>

### UpdateFloatingIPPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_pool | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFloatingIPRequest"/>

### UpdateFloatingIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateFloatingIPResponse"/>

### UpdateFloatingIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip | [FloatingIP](#github.com.Juniper.contrail.pkg.models.FloatingIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateForwardingClassRequest"/>

### UpdateForwardingClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_class | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateForwardingClassResponse"/>

### UpdateForwardingClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_class | [ForwardingClass](#github.com.Juniper.contrail.pkg.models.ForwardingClass) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateGlobalQosConfigRequest"/>

### UpdateGlobalQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_qos_config | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateGlobalQosConfigResponse"/>

### UpdateGlobalQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_qos_config | [GlobalQosConfig](#github.com.Juniper.contrail.pkg.models.GlobalQosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateGlobalSystemConfigRequest"/>

### UpdateGlobalSystemConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_system_config | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateGlobalSystemConfigResponse"/>

### UpdateGlobalSystemConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_system_config | [GlobalSystemConfig](#github.com.Juniper.contrail.pkg.models.GlobalSystemConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateGlobalVrouterConfigRequest"/>

### UpdateGlobalVrouterConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_vrouter_config | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateGlobalVrouterConfigResponse"/>

### UpdateGlobalVrouterConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_vrouter_config | [GlobalVrouterConfig](#github.com.Juniper.contrail.pkg.models.GlobalVrouterConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateInstanceIPRequest"/>

### UpdateInstanceIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_ip | [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateInstanceIPResponse"/>

### UpdateInstanceIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_ip | [InstanceIP](#github.com.Juniper.contrail.pkg.models.InstanceIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateInterfaceRouteTableRequest"/>

### UpdateInterfaceRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_route_table | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateInterfaceRouteTableResponse"/>

### UpdateInterfaceRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_route_table | [InterfaceRouteTable](#github.com.Juniper.contrail.pkg.models.InterfaceRouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateKeypairRequest"/>

### UpdateKeypairRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keypair | [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateKeypairResponse"/>

### UpdateKeypairResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keypair | [Keypair](#github.com.Juniper.contrail.pkg.models.Keypair) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateKubernetesMasterNodeRequest"/>

### UpdateKubernetesMasterNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_master_node | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateKubernetesMasterNodeResponse"/>

### UpdateKubernetesMasterNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_master_node | [KubernetesMasterNode](#github.com.Juniper.contrail.pkg.models.KubernetesMasterNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateKubernetesNodeRequest"/>

### UpdateKubernetesNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_node | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateKubernetesNodeResponse"/>

### UpdateKubernetesNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes_node | [KubernetesNode](#github.com.Juniper.contrail.pkg.models.KubernetesNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerHealthmonitorRequest"/>

### UpdateLoadbalancerHealthmonitorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_healthmonitor | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerHealthmonitorResponse"/>

### UpdateLoadbalancerHealthmonitorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_healthmonitor | [LoadbalancerHealthmonitor](#github.com.Juniper.contrail.pkg.models.LoadbalancerHealthmonitor) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerListenerRequest"/>

### UpdateLoadbalancerListenerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_listener | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerListenerResponse"/>

### UpdateLoadbalancerListenerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_listener | [LoadbalancerListener](#github.com.Juniper.contrail.pkg.models.LoadbalancerListener) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerMemberRequest"/>

### UpdateLoadbalancerMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_member | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerMemberResponse"/>

### UpdateLoadbalancerMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_member | [LoadbalancerMember](#github.com.Juniper.contrail.pkg.models.LoadbalancerMember) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerPoolRequest"/>

### UpdateLoadbalancerPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_pool | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerPoolResponse"/>

### UpdateLoadbalancerPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer_pool | [LoadbalancerPool](#github.com.Juniper.contrail.pkg.models.LoadbalancerPool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerRequest"/>

### UpdateLoadbalancerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLoadbalancerResponse"/>

### UpdateLoadbalancerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| loadbalancer | [Loadbalancer](#github.com.Juniper.contrail.pkg.models.Loadbalancer) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLocationRequest"/>

### UpdateLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [Location](#github.com.Juniper.contrail.pkg.models.Location) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLocationResponse"/>

### UpdateLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [Location](#github.com.Juniper.contrail.pkg.models.Location) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLogicalInterfaceRequest"/>

### UpdateLogicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_interface | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLogicalInterfaceResponse"/>

### UpdateLogicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_interface | [LogicalInterface](#github.com.Juniper.contrail.pkg.models.LogicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLogicalRouterRequest"/>

### UpdateLogicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_router | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateLogicalRouterResponse"/>

### UpdateLogicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logical_router | [LogicalRouter](#github.com.Juniper.contrail.pkg.models.LogicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNamespaceRequest"/>

### UpdateNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNamespaceResponse"/>

### UpdateNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [Namespace](#github.com.Juniper.contrail.pkg.models.Namespace) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNetworkDeviceConfigRequest"/>

### UpdateNetworkDeviceConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_device_config | [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNetworkDeviceConfigResponse"/>

### UpdateNetworkDeviceConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_device_config | [NetworkDeviceConfig](#github.com.Juniper.contrail.pkg.models.NetworkDeviceConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNetworkIpamRequest"/>

### UpdateNetworkIpamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_ipam | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNetworkIpamResponse"/>

### UpdateNetworkIpamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_ipam | [NetworkIpam](#github.com.Juniper.contrail.pkg.models.NetworkIpam) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNetworkPolicyRequest"/>

### UpdateNetworkPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_policy | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNetworkPolicyResponse"/>

### UpdateNetworkPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network_policy | [NetworkPolicy](#github.com.Juniper.contrail.pkg.models.NetworkPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNodeRequest"/>

### UpdateNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#github.com.Juniper.contrail.pkg.models.Node) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateNodeResponse"/>

### UpdateNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#github.com.Juniper.contrail.pkg.models.Node) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackComputeNodeRequest"/>

### UpdateOpenstackComputeNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_compute_node | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackComputeNodeResponse"/>

### UpdateOpenstackComputeNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_compute_node | [OpenstackComputeNode](#github.com.Juniper.contrail.pkg.models.OpenstackComputeNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackControlNodeRequest"/>

### UpdateOpenstackControlNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_control_node | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackControlNodeResponse"/>

### UpdateOpenstackControlNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_control_node | [OpenstackControlNode](#github.com.Juniper.contrail.pkg.models.OpenstackControlNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackMonitoringNodeRequest"/>

### UpdateOpenstackMonitoringNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_monitoring_node | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackMonitoringNodeResponse"/>

### UpdateOpenstackMonitoringNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_monitoring_node | [OpenstackMonitoringNode](#github.com.Juniper.contrail.pkg.models.OpenstackMonitoringNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackNetworkNodeRequest"/>

### UpdateOpenstackNetworkNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_network_node | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackNetworkNodeResponse"/>

### UpdateOpenstackNetworkNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_network_node | [OpenstackNetworkNode](#github.com.Juniper.contrail.pkg.models.OpenstackNetworkNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackStorageNodeRequest"/>

### UpdateOpenstackStorageNodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_storage_node | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOpenstackStorageNodeResponse"/>

### UpdateOpenstackStorageNodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openstack_storage_node | [OpenstackStorageNode](#github.com.Juniper.contrail.pkg.models.OpenstackStorageNode) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOsImageRequest"/>

### UpdateOsImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| os_image | [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateOsImageResponse"/>

### UpdateOsImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| os_image | [OsImage](#github.com.Juniper.contrail.pkg.models.OsImage) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePeeringPolicyRequest"/>

### UpdatePeeringPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peering_policy | [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePeeringPolicyResponse"/>

### UpdatePeeringPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peering_policy | [PeeringPolicy](#github.com.Juniper.contrail.pkg.models.PeeringPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePhysicalInterfaceRequest"/>

### UpdatePhysicalInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_interface | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePhysicalInterfaceResponse"/>

### UpdatePhysicalInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_interface | [PhysicalInterface](#github.com.Juniper.contrail.pkg.models.PhysicalInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePhysicalRouterRequest"/>

### UpdatePhysicalRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_router | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePhysicalRouterResponse"/>

### UpdatePhysicalRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| physical_router | [PhysicalRouter](#github.com.Juniper.contrail.pkg.models.PhysicalRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePolicyManagementRequest"/>

### UpdatePolicyManagementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_management | [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePolicyManagementResponse"/>

### UpdatePolicyManagementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy_management | [PolicyManagement](#github.com.Juniper.contrail.pkg.models.PolicyManagement) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePortRequest"/>

### UpdatePortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [Port](#github.com.Juniper.contrail.pkg.models.Port) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePortResponse"/>

### UpdatePortResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [Port](#github.com.Juniper.contrail.pkg.models.Port) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePortTupleRequest"/>

### UpdatePortTupleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_tuple | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdatePortTupleResponse"/>

### UpdatePortTupleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_tuple | [PortTuple](#github.com.Juniper.contrail.pkg.models.PortTuple) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateProjectRequest"/>

### UpdateProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#github.com.Juniper.contrail.pkg.models.Project) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateProjectResponse"/>

### UpdateProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#github.com.Juniper.contrail.pkg.models.Project) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateProviderAttachmentRequest"/>

### UpdateProviderAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_attachment | [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateProviderAttachmentResponse"/>

### UpdateProviderAttachmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_attachment | [ProviderAttachment](#github.com.Juniper.contrail.pkg.models.ProviderAttachment) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateQosConfigRequest"/>

### UpdateQosConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_config | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateQosConfigResponse"/>

### UpdateQosConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_config | [QosConfig](#github.com.Juniper.contrail.pkg.models.QosConfig) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateQosQueueRequest"/>

### UpdateQosQueueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_queue | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateQosQueueResponse"/>

### UpdateQosQueueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| qos_queue | [QosQueue](#github.com.Juniper.contrail.pkg.models.QosQueue) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRouteAggregateRequest"/>

### UpdateRouteAggregateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_aggregate | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRouteAggregateResponse"/>

### UpdateRouteAggregateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_aggregate | [RouteAggregate](#github.com.Juniper.contrail.pkg.models.RouteAggregate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRouteTableRequest"/>

### UpdateRouteTableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_table | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRouteTableResponse"/>

### UpdateRouteTableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_table | [RouteTable](#github.com.Juniper.contrail.pkg.models.RouteTable) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRouteTargetRequest"/>

### UpdateRouteTargetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_target | [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRouteTargetResponse"/>

### UpdateRouteTargetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route_target | [RouteTarget](#github.com.Juniper.contrail.pkg.models.RouteTarget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRoutingInstanceRequest"/>

### UpdateRoutingInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instance | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRoutingInstanceResponse"/>

### UpdateRoutingInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instance | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRoutingPolicyRequest"/>

### UpdateRoutingPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_policy | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateRoutingPolicyResponse"/>

### UpdateRoutingPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_policy | [RoutingPolicy](#github.com.Juniper.contrail.pkg.models.RoutingPolicy) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateSecurityGroupRequest"/>

### UpdateSecurityGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_group | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateSecurityGroupResponse"/>

### UpdateSecurityGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_group | [SecurityGroup](#github.com.Juniper.contrail.pkg.models.SecurityGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateSecurityLoggingObjectRequest"/>

### UpdateSecurityLoggingObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_logging_object | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateSecurityLoggingObjectResponse"/>

### UpdateSecurityLoggingObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_logging_object | [SecurityLoggingObject](#github.com.Juniper.contrail.pkg.models.SecurityLoggingObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServerRequest"/>

### UpdateServerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [Server](#github.com.Juniper.contrail.pkg.models.Server) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServerResponse"/>

### UpdateServerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [Server](#github.com.Juniper.contrail.pkg.models.Server) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceRequest"/>

### UpdateServiceApplianceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceResponse"/>

### UpdateServiceApplianceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance | [ServiceAppliance](#github.com.Juniper.contrail.pkg.models.ServiceAppliance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceSetRequest"/>

### UpdateServiceApplianceSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance_set | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceApplianceSetResponse"/>

### UpdateServiceApplianceSetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_appliance_set | [ServiceApplianceSet](#github.com.Juniper.contrail.pkg.models.ServiceApplianceSet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceConnectionModuleRequest"/>

### UpdateServiceConnectionModuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_connection_module | [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceConnectionModuleResponse"/>

### UpdateServiceConnectionModuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_connection_module | [ServiceConnectionModule](#github.com.Juniper.contrail.pkg.models.ServiceConnectionModule) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceEndpointRequest"/>

### UpdateServiceEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_endpoint | [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceEndpointResponse"/>

### UpdateServiceEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_endpoint | [ServiceEndpoint](#github.com.Juniper.contrail.pkg.models.ServiceEndpoint) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceGroupRequest"/>

### UpdateServiceGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_group | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceGroupResponse"/>

### UpdateServiceGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_group | [ServiceGroup](#github.com.Juniper.contrail.pkg.models.ServiceGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceHealthCheckRequest"/>

### UpdateServiceHealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_health_check | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceHealthCheckResponse"/>

### UpdateServiceHealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_health_check | [ServiceHealthCheck](#github.com.Juniper.contrail.pkg.models.ServiceHealthCheck) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceInstanceRequest"/>

### UpdateServiceInstanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_instance | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceInstanceResponse"/>

### UpdateServiceInstanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_instance | [ServiceInstance](#github.com.Juniper.contrail.pkg.models.ServiceInstance) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceObjectRequest"/>

### UpdateServiceObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_object | [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceObjectResponse"/>

### UpdateServiceObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_object | [ServiceObject](#github.com.Juniper.contrail.pkg.models.ServiceObject) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceTemplateRequest"/>

### UpdateServiceTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_template | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateServiceTemplateResponse"/>

### UpdateServiceTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_template | [ServiceTemplate](#github.com.Juniper.contrail.pkg.models.ServiceTemplate) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateSubnetRequest"/>

### UpdateSubnetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateSubnetResponse"/>

### UpdateSubnetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [Subnet](#github.com.Juniper.contrail.pkg.models.Subnet) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateTagRequest"/>

### UpdateTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateTagResponse"/>

### UpdateTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#github.com.Juniper.contrail.pkg.models.Tag) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateTagTypeRequest"/>

### UpdateTagTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_type | [TagType](#github.com.Juniper.contrail.pkg.models.TagType) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateTagTypeResponse"/>

### UpdateTagTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_type | [TagType](#github.com.Juniper.contrail.pkg.models.TagType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateUserRequest"/>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#github.com.Juniper.contrail.pkg.models.User) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateUserResponse"/>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#github.com.Juniper.contrail.pkg.models.User) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVPNGroupRequest"/>

### UpdateVPNGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpn_group | [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVPNGroupResponse"/>

### UpdateVPNGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpn_group | [VPNGroup](#github.com.Juniper.contrail.pkg.models.VPNGroup) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSRecordRequest"/>

### UpdateVirtualDNSRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS_record | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSRecordResponse"/>

### UpdateVirtualDNSRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS_record | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSRequest"/>

### UpdateVirtualDNSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualDNSResponse"/>

### UpdateVirtualDNSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_DNS | [VirtualDNS](#github.com.Juniper.contrail.pkg.models.VirtualDNS) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualIPRequest"/>

### UpdateVirtualIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_ip | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualIPResponse"/>

### UpdateVirtualIPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_ip | [VirtualIP](#github.com.Juniper.contrail.pkg.models.VirtualIP) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineInterfaceRequest"/>

### UpdateVirtualMachineInterfaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine_interface | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineInterfaceResponse"/>

### UpdateVirtualMachineInterfaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine_interface | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineRequest"/>

### UpdateVirtualMachineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine | [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualMachineResponse"/>

### UpdateVirtualMachineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_machine | [VirtualMachine](#github.com.Juniper.contrail.pkg.models.VirtualMachine) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualNetworkRequest"/>

### UpdateVirtualNetworkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_network | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualNetworkResponse"/>

### UpdateVirtualNetworkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_network | [VirtualNetwork](#github.com.Juniper.contrail.pkg.models.VirtualNetwork) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualRouterRequest"/>

### UpdateVirtualRouterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_router | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateVirtualRouterResponse"/>

### UpdateVirtualRouterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtual_router | [VirtualRouter](#github.com.Juniper.contrail.pkg.models.VirtualRouter) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateWidgetRequest"/>

### UpdateWidgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| widget | [Widget](#github.com.Juniper.contrail.pkg.models.Widget) |  |  |
| field_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UpdateWidgetResponse"/>

### UpdateWidgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| widget | [Widget](#github.com.Juniper.contrail.pkg.models.Widget) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.User"/>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| password | [string](#string) |  | Domain level quota, not currently implemented |






<a name="github.com.Juniper.contrail.pkg.models.UserCredentials"/>

### UserCredentials



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.UserDefinedLogStat"/>

### UserDefinedLogStat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pattern | [string](#string) |  | Perl type regular expression pattern to match |
| name | [string](#string) |  | Name of the stat |






<a name="github.com.Juniper.contrail.pkg.models.UserDefinedLogStatList"/>

### UserDefinedLogStatList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| statlist | [UserDefinedLogStat](#github.com.Juniper.contrail.pkg.models.UserDefinedLogStat) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.UveKeysType"/>

### UveKeysType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uve_key | [string](#string) | repeated | List of UVE tables where this alarm config should be applied |






<a name="github.com.Juniper.contrail.pkg.models.VPNGroup"/>

### VPNGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provisioning_log | [string](#string) |  | Provisioning Log |
| provisioning_progress | [int64](#int64) |  | Provisioning progress 0 - 100% |
| provisioning_progress_stage | [string](#string) |  | Provisioning Progress Stage |
| provisioning_start_time | [string](#string) |  | Time provisioning started |
| provisioning_state | [string](#string) |  | Provisioning Status |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| type | [string](#string) |  | Type of VPN |
| location_refs | [VPNGroupLocationRef](#github.com.Juniper.contrail.pkg.models.VPNGroupLocationRef) | repeated | Reference to the locations |






<a name="github.com.Juniper.contrail.pkg.models.VPNGroupLocationRef"/>

### VPNGroupLocationRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualDNS"/>

### VirtualDNS



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_DNS_data | [VirtualDnsType](#github.com.Juniper.contrail.pkg.models.VirtualDnsType) |  | Virtual DNS data has configuration for virtual DNS like domain, dynamic records etc. |
| virtual_DNS_records | [VirtualDNSRecord](#github.com.Juniper.contrail.pkg.models.VirtualDNSRecord) | repeated | Static DNS records in virtual DNS server. |






<a name="github.com.Juniper.contrail.pkg.models.VirtualDNSRecord"/>

### VirtualDNSRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_DNS_record_data | [VirtualDnsRecordType](#github.com.Juniper.contrail.pkg.models.VirtualDnsRecordType) |  | DNS record data has configuration like type, name, ip address, loadbalancing etc. |






<a name="github.com.Juniper.contrail.pkg.models.VirtualDnsRecordType"/>

### VirtualDnsRecordType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record_name | [string](#string) |  | DNS name to be resolved |
| record_class | [string](#string) |  | DNS record class supported is IN |
| record_data | [string](#string) |  | DNS record data is either ip address or string depending on type |
| record_type | [string](#string) |  | DNS record type can be A, AAAA, CNAME, PTR, NS and MX |
| record_ttl_seconds | [int64](#int64) |  | Time To Live for this DNS record |
| record_mx_preference | [int64](#int64) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualDnsType"/>

### VirtualDnsType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| floating_ip_record | [string](#string) |  | Decides how floating ip records are added |
| domain_name | [string](#string) |  | Default domain name for this virtual DNS server |
| external_visible | [bool](#bool) |  | Currently this option is not supported |
| next_virtual_DNS | [string](#string) |  | Next virtual DNS server to lookup if record is not found. Default is proxy to infrastructure DNS |
| dynamic_records_from_client | [bool](#bool) |  | Allow automatic addition of records on VM launch, default is True |
| reverse_resolution | [bool](#bool) |  | Allow reverse DNS resolution, ip to name mapping |
| default_ttl_seconds | [int64](#int64) |  | Default Time To Live for DNS records |
| record_order | [string](#string) |  | Order of DNS load balancing, fixed, random, round-robin. Default is random |






<a name="github.com.Juniper.contrail.pkg.models.VirtualIP"/>

### VirtualIP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_ip_properties | [VirtualIpType](#github.com.Juniper.contrail.pkg.models.VirtualIpType) |  | Virtual ip configuration like port, protocol, subnet etc. |
| loadbalancer_pool_refs | [VirtualIPLoadbalancerPoolRef](#github.com.Juniper.contrail.pkg.models.VirtualIPLoadbalancerPoolRef) | repeated | Reference to the load balancer pool that this virtual ip represent. Applicable only to LBaaS V1 |
| virtual_machine_interface_refs | [VirtualIPVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.VirtualIPVirtualMachineInterfaceRef) | repeated | Reference to the virtual machine interface for virtual ip. Applicable only to LBaaS V1 |






<a name="github.com.Juniper.contrail.pkg.models.VirtualIPLoadbalancerPoolRef"/>

### VirtualIPLoadbalancerPoolRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualIPVirtualMachineInterfaceRef"/>

### VirtualIPVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualIpType"/>

### VirtualIpType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  | Operating status for this virtual ip. |
| status_description | [string](#string) |  | Operating status description this virtual ip. |
| protocol | [string](#string) |  | IP protocol string like http, https or tcp. |
| subnet_id | [string](#string) |  | UUID of subnet in which to allocate the Virtual IP. |
| persistence_cookie_name | [string](#string) |  | Set this string if the relation of client and server(pool member) need to persist. |
| connection_limit | [int64](#int64) |  | Maximum number of concurrent connections |
| persistence_type | [string](#string) |  | Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE. |
| admin_state | [bool](#bool) |  | Administrative up or down. |
| address | [string](#string) |  | IP address automatically allocated by system. |
| protocol_port | [int64](#int64) |  | Layer 4 protocol destination port. |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachine"/>

### VirtualMachine



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| service_instance_refs | [VirtualMachineServiceInstanceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineServiceInstanceRef) | repeated | Reference to the service instance of this virtual machine. |
| virtual_machine_interfaces | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) | repeated | References to child interfaces this virtual machine has, this is DEPRECATED. |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterface"/>

### VirtualMachineInterface



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| ecmp_hashing_include_fields | [EcmpHashingIncludeFields](#github.com.Juniper.contrail.pkg.models.EcmpHashingIncludeFields) |  | ECMP hashing config at global level. |
| virtual_machine_interface_host_routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  | List of host routes(prefixes, nexthop) that are passed to VM via DHCP. |
| virtual_machine_interface_mac_addresses | [MacAddressesType](#github.com.Juniper.contrail.pkg.models.MacAddressesType) |  | MAC address of the virtual machine interface, automatically assigned by system if not provided. |
| virtual_machine_interface_dhcp_option_list | [DhcpOptionsListType](#github.com.Juniper.contrail.pkg.models.DhcpOptionsListType) |  | DHCP options configuration specific to this interface. |
| virtual_machine_interface_bindings | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) for this interface. Neutron port bindings use this. |
| virtual_machine_interface_disable_policy | [bool](#bool) |  | When True all policy checks for ingress and egress traffic from this interface are disabled. Flow table entries are not created. Features that require policy will not work on this interface, these include security group, floating IP, service chain, linklocal services. |
| virtual_machine_interface_allowed_address_pairs | [AllowedAddressPairs](#github.com.Juniper.contrail.pkg.models.AllowedAddressPairs) |  | List of (IP address, MAC) other than instance ip on this interface. |
| virtual_machine_interface_fat_flow_protocols | [FatFlowProtocols](#github.com.Juniper.contrail.pkg.models.FatFlowProtocols) |  | List of (protocol, port number), for flows to interface with (protocol, destination port number), vrouter will ignore source port while setting up flow and ignore it as source port in reverse flow. Hence many flows will map to single flow. |
| vlan_tag_based_bridge_domain | [bool](#bool) |  | Enable VLAN tag based bridge domain classification on the port |
| virtual_machine_interface_device_owner | [string](#string) |  | For openstack compatibility, not used by system. |
| vrf_assign_table | [VrfAssignTableType](#github.com.Juniper.contrail.pkg.models.VrfAssignTableType) |  | VRF assignment policy for this interface, automatically generated by system. |
| port_security_enabled | [bool](#bool) |  | Port security status on the network |
| virtual_machine_interface_properties | [VirtualMachineInterfacePropertiesType](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePropertiesType) |  | Virtual Machine Interface miscellaneous configurations. |
| virtual_machine_interface_refs | [VirtualMachineInterfaceVirtualMachineInterfaceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualMachineInterfaceRef) | repeated | List of references to the sub interfaces of this interface. |
| virtual_machine_refs | [VirtualMachineInterfaceVirtualMachineRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualMachineRef) | repeated | This interface belongs to the referenced virtual machine. |
| service_health_check_refs | [VirtualMachineInterfaceServiceHealthCheckRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceServiceHealthCheckRef) | repeated | Reference to health check object attached to this interface. |
| interface_route_table_refs | [VirtualMachineInterfaceInterfaceRouteTableRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceInterfaceRouteTableRef) | repeated | Reference to the interface route table attached to this interface. |
| physical_interface_refs | [VirtualMachineInterfacePhysicalInterfaceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePhysicalInterfaceRef) | repeated | Reference to the physical interface of service appliance this service interface represents. |
| bridge_domain_refs | [VirtualMachineInterfaceBridgeDomainRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceBridgeDomainRef) | repeated | Virtual Machine interface maps to a bridge-domain by defaultor based on in coming 802.1Q vlan tag |
| security_group_refs | [VirtualMachineInterfaceSecurityGroupRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceSecurityGroupRef) | repeated | Interface ACL, Automatically generated by system based on security groups attached to this interface. |
| service_endpoint_refs | [VirtualMachineInterfaceServiceEndpointRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceServiceEndpointRef) | repeated | Links the access endpoint i.e virtual-machine-interface to service endpoint. |
| virtual_network_refs | [VirtualMachineInterfaceVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualNetworkRef) | repeated | This interface is member of the referenced virtual network. |
| bgp_router_refs | [VirtualMachineInterfaceBGPRouterRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceBGPRouterRef) | repeated | Reference to bgp-router from the virtual machine interface. |
| security_logging_object_refs | [VirtualMachineInterfaceSecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceSecurityLoggingObjectRef) | repeated | Reference to security logging object for this virtual machine interface |
| routing_instance_refs | [VirtualMachineInterfaceRoutingInstanceRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceRoutingInstanceRef) | repeated | Automatically generated Forwarding policy. This will be deprecated in future in favour of VRF assign rules. |
| qos_config_refs | [VirtualMachineInterfaceQosConfigRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceQosConfigRef) | repeated | Reference to QoS config for this virtual machine interface. |
| port_tuple_refs | [VirtualMachineInterfacePortTupleRef](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePortTupleRef) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceBGPRouterRef"/>

### VirtualMachineInterfaceBGPRouterRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceBridgeDomainRef"/>

### VirtualMachineInterfaceBridgeDomainRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [BridgeDomainMembershipType](#github.com.Juniper.contrail.pkg.models.BridgeDomainMembershipType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceInterfaceRouteTableRef"/>

### VirtualMachineInterfaceInterfaceRouteTableRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePhysicalInterfaceRef"/>

### VirtualMachineInterfacePhysicalInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePortTupleRef"/>

### VirtualMachineInterfacePortTupleRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfacePropertiesType"/>

### VirtualMachineInterfacePropertiesType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_interface_vlan_tag | [int64](#int64) |  | 802.1Q VLAN tag to be used if this interface is sub-interface for some other interface. |
| local_preference | [int64](#int64) |  | BGP route local preference for routes representing this interface, higher value is higher preference |
| interface_mirror | [InterfaceMirrorType](#github.com.Juniper.contrail.pkg.models.InterfaceMirrorType) |  | Interface Mirror configuration |
| service_interface_type | [string](#string) |  | This interface belongs to Service Instance and is tagged as left, right or other |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceQosConfigRef"/>

### VirtualMachineInterfaceQosConfigRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceRoutingInstanceRef"/>

### VirtualMachineInterfaceRoutingInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [PolicyBasedForwardingRuleType](#github.com.Juniper.contrail.pkg.models.PolicyBasedForwardingRuleType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceSecurityGroupRef"/>

### VirtualMachineInterfaceSecurityGroupRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceSecurityLoggingObjectRef"/>

### VirtualMachineInterfaceSecurityLoggingObjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceServiceEndpointRef"/>

### VirtualMachineInterfaceServiceEndpointRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceServiceHealthCheckRef"/>

### VirtualMachineInterfaceServiceHealthCheckRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualMachineInterfaceRef"/>

### VirtualMachineInterfaceVirtualMachineInterfaceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualMachineRef"/>

### VirtualMachineInterfaceVirtualMachineRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineInterfaceVirtualNetworkRef"/>

### VirtualMachineInterfaceVirtualNetworkRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualMachineServiceInstanceRef"/>

### VirtualMachineServiceInstanceRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetwork"/>

### VirtualNetwork



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_network_properties | [VirtualNetworkType](#github.com.Juniper.contrail.pkg.models.VirtualNetworkType) |  | Virtual network miscellaneous configurations. |
| ecmp_hashing_include_fields | [EcmpHashingIncludeFields](#github.com.Juniper.contrail.pkg.models.EcmpHashingIncludeFields) |  | ECMP hashing config at global level. |
| virtual_network_network_id | [int64](#int64) |  | System assigned unique 32 bit ID for every virtual network. |
| address_allocation_mode | [string](#string) |  | Address allocation mode for virtual network. |
| pbb_evpn_enable | [bool](#bool) |  | Enable/Disable PBB EVPN tunneling on the network |
| router_external | [bool](#bool) |  | When true, this virtual network is openstack router external network. |
| import_route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that are used as import for this virtual network. |
| mac_aging_time | [int64](#int64) |  | MAC aging time on the network |
| provider_properties | [ProviderDetails](#github.com.Juniper.contrail.pkg.models.ProviderDetails) |  | Virtual network is provider network. Specifies VLAN tag and physical network name. |
| route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that are used as both import and export for this virtual network. |
| mac_learning_enabled | [bool](#bool) |  | Enable MAC learning on the network |
| export_route_target_list | [RouteTargetList](#github.com.Juniper.contrail.pkg.models.RouteTargetList) |  | List of route targets that are used as export for this virtual network. |
| flood_unknown_unicast | [bool](#bool) |  | When true, packets with unknown unicast MAC address are flooded within the network. Default they are dropped. |
| pbb_etree_enable | [bool](#bool) |  | Enable/Disable PBB ETREE mode on the network |
| layer2_control_word | [bool](#bool) |  | Enable/Disable adding control word to the Layer 2 encapsulation |
| external_ipam | [bool](#bool) |  | IP address assignment to VM is done statically, outside of (external to) Contrail Ipam. vCenter only feature. |
| port_security_enabled | [bool](#bool) |  | Port security status on the network |
| mac_move_control | [MACMoveLimitControlType](#github.com.Juniper.contrail.pkg.models.MACMoveLimitControlType) |  | MAC move control on the network |
| multi_policy_service_chains_enabled | [bool](#bool) |  |  |
| mac_limit_control | [MACLimitControlType](#github.com.Juniper.contrail.pkg.models.MACLimitControlType) |  | MAC limit control on the network |
| is_shared | [bool](#bool) |  | When true, this virtual network is shared with all tenants. |
| network_ipam_refs | [VirtualNetworkNetworkIpamRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkNetworkIpamRef) | repeated | Reference to network-ipam this network is using. It has list of subnets that are being used as property of the reference. |
| security_logging_object_refs | [VirtualNetworkSecurityLoggingObjectRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkSecurityLoggingObjectRef) | repeated | Reference to security logging object for this virtual network. |
| network_policy_refs | [VirtualNetworkNetworkPolicyRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkNetworkPolicyRef) | repeated | Reference to network-policy attached to this network. It has sequence number to specify attachment order. |
| qos_config_refs | [VirtualNetworkQosConfigRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkQosConfigRef) | repeated | Reference to QoS configuration for this virtual network. |
| route_table_refs | [VirtualNetworkRouteTableRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkRouteTableRef) | repeated | Reference to route table attached to this virtual network. |
| virtual_network_refs | [VirtualNetworkVirtualNetworkRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkVirtualNetworkRef) | repeated | Reference to a virtual network which is the provider network for the given virtual network. Traffic forwarding happens in the routing instance of the provider network. |
| bgpvpn_refs | [VirtualNetworkBGPVPNRef](#github.com.Juniper.contrail.pkg.models.VirtualNetworkBGPVPNRef) | repeated | Back reference to virtual network associated to the BGP VPN resource |
| access_control_lists | [AccessControlList](#github.com.Juniper.contrail.pkg.models.AccessControlList) | repeated | Virtual network access control list are automatically derived from all the network policies attached to virtual network. |
| alias_ip_pools | [AliasIPPool](#github.com.Juniper.contrail.pkg.models.AliasIPPool) | repeated | Alias ip pool is set of addresses that are carved out of a given network. Ip(s) from this set can be assigned to virtual-machine-interface so that they become members of this network |
| bridge_domains | [BridgeDomain](#github.com.Juniper.contrail.pkg.models.BridgeDomain) | repeated | bridge-domains configured in a virtual network |
| floating_ip_pools | [FloatingIPPool](#github.com.Juniper.contrail.pkg.models.FloatingIPPool) | repeated | Floating ip pool is set of ip address that are carved out of a given network. Ip(s) from this set can be assigned to (virtual machine interface, ip) so that they become members of this network using one:one NAT. |
| routing_instances | [RoutingInstance](#github.com.Juniper.contrail.pkg.models.RoutingInstance) | repeated | List of references of routing instances for this virtual network, routing instances are internal to the system. |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkBGPVPNRef"/>

### VirtualNetworkBGPVPNRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkNetworkIpamRef"/>

### VirtualNetworkNetworkIpamRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [VnSubnetsType](#github.com.Juniper.contrail.pkg.models.VnSubnetsType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkNetworkPolicyRef"/>

### VirtualNetworkNetworkPolicyRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [VirtualNetworkPolicyType](#github.com.Juniper.contrail.pkg.models.VirtualNetworkPolicyType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkPolicyType"/>

### VirtualNetworkPolicyType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timer | [TimerType](#github.com.Juniper.contrail.pkg.models.TimerType) |  | Timer to specify when the policy can be active |
| sequence | [SequenceType](#github.com.Juniper.contrail.pkg.models.SequenceType) |  | Sequence number to specify order of policy attachment to network |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkQosConfigRef"/>

### VirtualNetworkQosConfigRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkRouteTableRef"/>

### VirtualNetworkRouteTableRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkSecurityLoggingObjectRef"/>

### VirtualNetworkSecurityLoggingObjectRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkType"/>

### VirtualNetworkType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forwarding_mode | [string](#string) |  | Packet forwarding mode for this virtual network |
| allow_transit | [bool](#bool) |  |  |
| network_id | [int64](#int64) |  | Not currently in used |
| mirror_destination | [bool](#bool) |  | Flag to mark the virtual network as mirror destination network |
| vxlan_network_identifier | [int64](#int64) |  | VxLAN VNI value for this network |
| rpf | [string](#string) |  | Flag used to disable Reverse Path Forwarding(RPF) check for this network |






<a name="github.com.Juniper.contrail.pkg.models.VirtualNetworkVirtualNetworkRef"/>

### VirtualNetworkVirtualNetworkRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualRouter"/>

### VirtualRouter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| virtual_router_dpdk_enabled | [bool](#bool) |  | This vrouter&amp;#39;s data path is using DPDK library, Virtual machines interfaces scheduled on this compute node will be tagged with additional flags so that they are spawned with user space virtio driver. It is only applicable for embedded vrouter. |
| virtual_router_type | [string](#string) |  | Different types of the vrouters in the system. |
| virtual_router_ip_address | [string](#string) |  | Ip address of the virtual router. |
| virtual_machine_refs | [VirtualRouterVirtualMachineRef](#github.com.Juniper.contrail.pkg.models.VirtualRouterVirtualMachineRef) | repeated | References to all virtual machines on this vrouter. This link is not present for dynamically scheduled VMs by Nova. |
| network_ipam_refs | [VirtualRouterNetworkIpamRef](#github.com.Juniper.contrail.pkg.models.VirtualRouterNetworkIpamRef) | repeated | Reference to network-ipam this virtual-router is using. It has list of virtual-router specific allocation-pools and cidrs that are being used as property of the reference. |
| virtual_machine_interfaces | [VirtualMachineInterface](#github.com.Juniper.contrail.pkg.models.VirtualMachineInterface) | repeated | An interface on a virtual-router, e.g. vhost0 |






<a name="github.com.Juniper.contrail.pkg.models.VirtualRouterNetworkIpamRef"/>

### VirtualRouterNetworkIpamRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |
| attr | [VirtualRouterNetworkIpamType](#github.com.Juniper.contrail.pkg.models.VirtualRouterNetworkIpamType) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VirtualRouterNetworkIpamType"/>

### VirtualRouterNetworkIpamType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnet | [SubnetType](#github.com.Juniper.contrail.pkg.models.SubnetType) | repeated | List of ip prefix and length for vrouter specific subnets |
| allocation_pools | [AllocationPoolType](#github.com.Juniper.contrail.pkg.models.AllocationPoolType) | repeated | List of ranges of ip address for vrouter specific allocation |






<a name="github.com.Juniper.contrail.pkg.models.VirtualRouterVirtualMachineRef"/>

### VirtualRouterVirtualMachineRef



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| to | [string](#string) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.VnSubnetsType"/>

### VnSubnetsType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipam_subnets | [IpamSubnetType](#github.com.Juniper.contrail.pkg.models.IpamSubnetType) | repeated |  |
| host_routes | [RouteTableType](#github.com.Juniper.contrail.pkg.models.RouteTableType) |  | Common host routes to be sent via DHCP for VM(s) in all the subnets, Next hop for these routes is always default gateway |






<a name="github.com.Juniper.contrail.pkg.models.VrfAssignRuleType"/>

### VrfAssignRuleType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_instance | [string](#string) |  |  |
| match_condition | [MatchConditionType](#github.com.Juniper.contrail.pkg.models.MatchConditionType) |  |  |
| vlan_tag | [int64](#int64) |  |  |
| ignore_acl | [bool](#bool) |  |  |






<a name="github.com.Juniper.contrail.pkg.models.VrfAssignTableType"/>

### VrfAssignTableType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vrf_assign_rule | [VrfAssignRuleType](#github.com.Juniper.contrail.pkg.models.VrfAssignRuleType) | repeated |  |






<a name="github.com.Juniper.contrail.pkg.models.Widget"/>

### Widget



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of the object, system automatically allocates one if not provided |
| parent_uuid | [string](#string) |  | UUID of the parent object |
| parent_type | [string](#string) |  | Parent resource type |
| fq_name | [string](#string) | repeated | FQ Name of the object |
| id_perms | [IdPermsType](#github.com.Juniper.contrail.pkg.models.IdPermsType) |  | System maintained identity, time and permissions data. |
| display_name | [string](#string) |  | Display name user configured string(name) that can be updated any time. Used as openstack name. |
| annotations | [KeyValuePairs](#github.com.Juniper.contrail.pkg.models.KeyValuePairs) |  | Dictionary of arbitrary (key, value) on a resource. |
| perms2 | [PermType2](#github.com.Juniper.contrail.pkg.models.PermType2) |  | Permissions data for role based access. |
| configuration_version | [int64](#int64) |  | Configuration Version for the object. |
| container_config | [string](#string) |  |  |
| content_config | [string](#string) |  |  |
| layout_config | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

