package models

import (
	"context"
)

type Service interface {
	Next() Service
	SetNext(Service)

	CreateAccessControlList(context.Context, *CreateAccessControlListRequest) (*CreateAccessControlListResponse, error)
	UpdateAccessControlList(context.Context, *UpdateAccessControlListRequest) (*UpdateAccessControlListResponse, error)
	DeleteAccessControlList(context.Context, *DeleteAccessControlListRequest) (*DeleteAccessControlListResponse, error)
	GetAccessControlList(context.Context, *GetAccessControlListRequest) (*GetAccessControlListResponse, error)
	ListAccessControlList(context.Context, *ListAccessControlListRequest) (*ListAccessControlListResponse, error)

	CreateAddressGroup(context.Context, *CreateAddressGroupRequest) (*CreateAddressGroupResponse, error)
	UpdateAddressGroup(context.Context, *UpdateAddressGroupRequest) (*UpdateAddressGroupResponse, error)
	DeleteAddressGroup(context.Context, *DeleteAddressGroupRequest) (*DeleteAddressGroupResponse, error)
	GetAddressGroup(context.Context, *GetAddressGroupRequest) (*GetAddressGroupResponse, error)
	ListAddressGroup(context.Context, *ListAddressGroupRequest) (*ListAddressGroupResponse, error)

	CreateAlarm(context.Context, *CreateAlarmRequest) (*CreateAlarmResponse, error)
	UpdateAlarm(context.Context, *UpdateAlarmRequest) (*UpdateAlarmResponse, error)
	DeleteAlarm(context.Context, *DeleteAlarmRequest) (*DeleteAlarmResponse, error)
	GetAlarm(context.Context, *GetAlarmRequest) (*GetAlarmResponse, error)
	ListAlarm(context.Context, *ListAlarmRequest) (*ListAlarmResponse, error)

	CreateAliasIPPool(context.Context, *CreateAliasIPPoolRequest) (*CreateAliasIPPoolResponse, error)
	UpdateAliasIPPool(context.Context, *UpdateAliasIPPoolRequest) (*UpdateAliasIPPoolResponse, error)
	DeleteAliasIPPool(context.Context, *DeleteAliasIPPoolRequest) (*DeleteAliasIPPoolResponse, error)
	GetAliasIPPool(context.Context, *GetAliasIPPoolRequest) (*GetAliasIPPoolResponse, error)
	ListAliasIPPool(context.Context, *ListAliasIPPoolRequest) (*ListAliasIPPoolResponse, error)

	CreateAliasIP(context.Context, *CreateAliasIPRequest) (*CreateAliasIPResponse, error)
	UpdateAliasIP(context.Context, *UpdateAliasIPRequest) (*UpdateAliasIPResponse, error)
	DeleteAliasIP(context.Context, *DeleteAliasIPRequest) (*DeleteAliasIPResponse, error)
	GetAliasIP(context.Context, *GetAliasIPRequest) (*GetAliasIPResponse, error)
	ListAliasIP(context.Context, *ListAliasIPRequest) (*ListAliasIPResponse, error)

	CreateAnalyticsNode(context.Context, *CreateAnalyticsNodeRequest) (*CreateAnalyticsNodeResponse, error)
	UpdateAnalyticsNode(context.Context, *UpdateAnalyticsNodeRequest) (*UpdateAnalyticsNodeResponse, error)
	DeleteAnalyticsNode(context.Context, *DeleteAnalyticsNodeRequest) (*DeleteAnalyticsNodeResponse, error)
	GetAnalyticsNode(context.Context, *GetAnalyticsNodeRequest) (*GetAnalyticsNodeResponse, error)
	ListAnalyticsNode(context.Context, *ListAnalyticsNodeRequest) (*ListAnalyticsNodeResponse, error)

	CreateAPIAccessList(context.Context, *CreateAPIAccessListRequest) (*CreateAPIAccessListResponse, error)
	UpdateAPIAccessList(context.Context, *UpdateAPIAccessListRequest) (*UpdateAPIAccessListResponse, error)
	DeleteAPIAccessList(context.Context, *DeleteAPIAccessListRequest) (*DeleteAPIAccessListResponse, error)
	GetAPIAccessList(context.Context, *GetAPIAccessListRequest) (*GetAPIAccessListResponse, error)
	ListAPIAccessList(context.Context, *ListAPIAccessListRequest) (*ListAPIAccessListResponse, error)

	CreateApplicationPolicySet(context.Context, *CreateApplicationPolicySetRequest) (*CreateApplicationPolicySetResponse, error)
	UpdateApplicationPolicySet(context.Context, *UpdateApplicationPolicySetRequest) (*UpdateApplicationPolicySetResponse, error)
	DeleteApplicationPolicySet(context.Context, *DeleteApplicationPolicySetRequest) (*DeleteApplicationPolicySetResponse, error)
	GetApplicationPolicySet(context.Context, *GetApplicationPolicySetRequest) (*GetApplicationPolicySetResponse, error)
	ListApplicationPolicySet(context.Context, *ListApplicationPolicySetRequest) (*ListApplicationPolicySetResponse, error)

	CreateBGPAsAService(context.Context, *CreateBGPAsAServiceRequest) (*CreateBGPAsAServiceResponse, error)
	UpdateBGPAsAService(context.Context, *UpdateBGPAsAServiceRequest) (*UpdateBGPAsAServiceResponse, error)
	DeleteBGPAsAService(context.Context, *DeleteBGPAsAServiceRequest) (*DeleteBGPAsAServiceResponse, error)
	GetBGPAsAService(context.Context, *GetBGPAsAServiceRequest) (*GetBGPAsAServiceResponse, error)
	ListBGPAsAService(context.Context, *ListBGPAsAServiceRequest) (*ListBGPAsAServiceResponse, error)

	CreateBGPRouter(context.Context, *CreateBGPRouterRequest) (*CreateBGPRouterResponse, error)
	UpdateBGPRouter(context.Context, *UpdateBGPRouterRequest) (*UpdateBGPRouterResponse, error)
	DeleteBGPRouter(context.Context, *DeleteBGPRouterRequest) (*DeleteBGPRouterResponse, error)
	GetBGPRouter(context.Context, *GetBGPRouterRequest) (*GetBGPRouterResponse, error)
	ListBGPRouter(context.Context, *ListBGPRouterRequest) (*ListBGPRouterResponse, error)

	CreateBGPVPN(context.Context, *CreateBGPVPNRequest) (*CreateBGPVPNResponse, error)
	UpdateBGPVPN(context.Context, *UpdateBGPVPNRequest) (*UpdateBGPVPNResponse, error)
	DeleteBGPVPN(context.Context, *DeleteBGPVPNRequest) (*DeleteBGPVPNResponse, error)
	GetBGPVPN(context.Context, *GetBGPVPNRequest) (*GetBGPVPNResponse, error)
	ListBGPVPN(context.Context, *ListBGPVPNRequest) (*ListBGPVPNResponse, error)

	CreateBridgeDomain(context.Context, *CreateBridgeDomainRequest) (*CreateBridgeDomainResponse, error)
	UpdateBridgeDomain(context.Context, *UpdateBridgeDomainRequest) (*UpdateBridgeDomainResponse, error)
	DeleteBridgeDomain(context.Context, *DeleteBridgeDomainRequest) (*DeleteBridgeDomainResponse, error)
	GetBridgeDomain(context.Context, *GetBridgeDomainRequest) (*GetBridgeDomainResponse, error)
	ListBridgeDomain(context.Context, *ListBridgeDomainRequest) (*ListBridgeDomainResponse, error)

	CreateConfigNode(context.Context, *CreateConfigNodeRequest) (*CreateConfigNodeResponse, error)
	UpdateConfigNode(context.Context, *UpdateConfigNodeRequest) (*UpdateConfigNodeResponse, error)
	DeleteConfigNode(context.Context, *DeleteConfigNodeRequest) (*DeleteConfigNodeResponse, error)
	GetConfigNode(context.Context, *GetConfigNodeRequest) (*GetConfigNodeResponse, error)
	ListConfigNode(context.Context, *ListConfigNodeRequest) (*ListConfigNodeResponse, error)

	CreateConfigRoot(context.Context, *CreateConfigRootRequest) (*CreateConfigRootResponse, error)
	UpdateConfigRoot(context.Context, *UpdateConfigRootRequest) (*UpdateConfigRootResponse, error)
	DeleteConfigRoot(context.Context, *DeleteConfigRootRequest) (*DeleteConfigRootResponse, error)
	GetConfigRoot(context.Context, *GetConfigRootRequest) (*GetConfigRootResponse, error)
	ListConfigRoot(context.Context, *ListConfigRootRequest) (*ListConfigRootResponse, error)

	CreateCustomerAttachment(context.Context, *CreateCustomerAttachmentRequest) (*CreateCustomerAttachmentResponse, error)
	UpdateCustomerAttachment(context.Context, *UpdateCustomerAttachmentRequest) (*UpdateCustomerAttachmentResponse, error)
	DeleteCustomerAttachment(context.Context, *DeleteCustomerAttachmentRequest) (*DeleteCustomerAttachmentResponse, error)
	GetCustomerAttachment(context.Context, *GetCustomerAttachmentRequest) (*GetCustomerAttachmentResponse, error)
	ListCustomerAttachment(context.Context, *ListCustomerAttachmentRequest) (*ListCustomerAttachmentResponse, error)

	CreateDatabaseNode(context.Context, *CreateDatabaseNodeRequest) (*CreateDatabaseNodeResponse, error)
	UpdateDatabaseNode(context.Context, *UpdateDatabaseNodeRequest) (*UpdateDatabaseNodeResponse, error)
	DeleteDatabaseNode(context.Context, *DeleteDatabaseNodeRequest) (*DeleteDatabaseNodeResponse, error)
	GetDatabaseNode(context.Context, *GetDatabaseNodeRequest) (*GetDatabaseNodeResponse, error)
	ListDatabaseNode(context.Context, *ListDatabaseNodeRequest) (*ListDatabaseNodeResponse, error)

	CreateDiscoveryServiceAssignment(context.Context, *CreateDiscoveryServiceAssignmentRequest) (*CreateDiscoveryServiceAssignmentResponse, error)
	UpdateDiscoveryServiceAssignment(context.Context, *UpdateDiscoveryServiceAssignmentRequest) (*UpdateDiscoveryServiceAssignmentResponse, error)
	DeleteDiscoveryServiceAssignment(context.Context, *DeleteDiscoveryServiceAssignmentRequest) (*DeleteDiscoveryServiceAssignmentResponse, error)
	GetDiscoveryServiceAssignment(context.Context, *GetDiscoveryServiceAssignmentRequest) (*GetDiscoveryServiceAssignmentResponse, error)
	ListDiscoveryServiceAssignment(context.Context, *ListDiscoveryServiceAssignmentRequest) (*ListDiscoveryServiceAssignmentResponse, error)

	CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error)
	UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error)
	DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainResponse, error)
	GetDomain(context.Context, *GetDomainRequest) (*GetDomainResponse, error)
	ListDomain(context.Context, *ListDomainRequest) (*ListDomainResponse, error)

	CreateDsaRule(context.Context, *CreateDsaRuleRequest) (*CreateDsaRuleResponse, error)
	UpdateDsaRule(context.Context, *UpdateDsaRuleRequest) (*UpdateDsaRuleResponse, error)
	DeleteDsaRule(context.Context, *DeleteDsaRuleRequest) (*DeleteDsaRuleResponse, error)
	GetDsaRule(context.Context, *GetDsaRuleRequest) (*GetDsaRuleResponse, error)
	ListDsaRule(context.Context, *ListDsaRuleRequest) (*ListDsaRuleResponse, error)

	CreateE2ServiceProvider(context.Context, *CreateE2ServiceProviderRequest) (*CreateE2ServiceProviderResponse, error)
	UpdateE2ServiceProvider(context.Context, *UpdateE2ServiceProviderRequest) (*UpdateE2ServiceProviderResponse, error)
	DeleteE2ServiceProvider(context.Context, *DeleteE2ServiceProviderRequest) (*DeleteE2ServiceProviderResponse, error)
	GetE2ServiceProvider(context.Context, *GetE2ServiceProviderRequest) (*GetE2ServiceProviderResponse, error)
	ListE2ServiceProvider(context.Context, *ListE2ServiceProviderRequest) (*ListE2ServiceProviderResponse, error)

	CreateFirewallPolicy(context.Context, *CreateFirewallPolicyRequest) (*CreateFirewallPolicyResponse, error)
	UpdateFirewallPolicy(context.Context, *UpdateFirewallPolicyRequest) (*UpdateFirewallPolicyResponse, error)
	DeleteFirewallPolicy(context.Context, *DeleteFirewallPolicyRequest) (*DeleteFirewallPolicyResponse, error)
	GetFirewallPolicy(context.Context, *GetFirewallPolicyRequest) (*GetFirewallPolicyResponse, error)
	ListFirewallPolicy(context.Context, *ListFirewallPolicyRequest) (*ListFirewallPolicyResponse, error)

	CreateFirewallRule(context.Context, *CreateFirewallRuleRequest) (*CreateFirewallRuleResponse, error)
	UpdateFirewallRule(context.Context, *UpdateFirewallRuleRequest) (*UpdateFirewallRuleResponse, error)
	DeleteFirewallRule(context.Context, *DeleteFirewallRuleRequest) (*DeleteFirewallRuleResponse, error)
	GetFirewallRule(context.Context, *GetFirewallRuleRequest) (*GetFirewallRuleResponse, error)
	ListFirewallRule(context.Context, *ListFirewallRuleRequest) (*ListFirewallRuleResponse, error)

	CreateFloatingIPPool(context.Context, *CreateFloatingIPPoolRequest) (*CreateFloatingIPPoolResponse, error)
	UpdateFloatingIPPool(context.Context, *UpdateFloatingIPPoolRequest) (*UpdateFloatingIPPoolResponse, error)
	DeleteFloatingIPPool(context.Context, *DeleteFloatingIPPoolRequest) (*DeleteFloatingIPPoolResponse, error)
	GetFloatingIPPool(context.Context, *GetFloatingIPPoolRequest) (*GetFloatingIPPoolResponse, error)
	ListFloatingIPPool(context.Context, *ListFloatingIPPoolRequest) (*ListFloatingIPPoolResponse, error)

	CreateFloatingIP(context.Context, *CreateFloatingIPRequest) (*CreateFloatingIPResponse, error)
	UpdateFloatingIP(context.Context, *UpdateFloatingIPRequest) (*UpdateFloatingIPResponse, error)
	DeleteFloatingIP(context.Context, *DeleteFloatingIPRequest) (*DeleteFloatingIPResponse, error)
	GetFloatingIP(context.Context, *GetFloatingIPRequest) (*GetFloatingIPResponse, error)
	ListFloatingIP(context.Context, *ListFloatingIPRequest) (*ListFloatingIPResponse, error)

	CreateForwardingClass(context.Context, *CreateForwardingClassRequest) (*CreateForwardingClassResponse, error)
	UpdateForwardingClass(context.Context, *UpdateForwardingClassRequest) (*UpdateForwardingClassResponse, error)
	DeleteForwardingClass(context.Context, *DeleteForwardingClassRequest) (*DeleteForwardingClassResponse, error)
	GetForwardingClass(context.Context, *GetForwardingClassRequest) (*GetForwardingClassResponse, error)
	ListForwardingClass(context.Context, *ListForwardingClassRequest) (*ListForwardingClassResponse, error)

	CreateGlobalQosConfig(context.Context, *CreateGlobalQosConfigRequest) (*CreateGlobalQosConfigResponse, error)
	UpdateGlobalQosConfig(context.Context, *UpdateGlobalQosConfigRequest) (*UpdateGlobalQosConfigResponse, error)
	DeleteGlobalQosConfig(context.Context, *DeleteGlobalQosConfigRequest) (*DeleteGlobalQosConfigResponse, error)
	GetGlobalQosConfig(context.Context, *GetGlobalQosConfigRequest) (*GetGlobalQosConfigResponse, error)
	ListGlobalQosConfig(context.Context, *ListGlobalQosConfigRequest) (*ListGlobalQosConfigResponse, error)

	CreateGlobalSystemConfig(context.Context, *CreateGlobalSystemConfigRequest) (*CreateGlobalSystemConfigResponse, error)
	UpdateGlobalSystemConfig(context.Context, *UpdateGlobalSystemConfigRequest) (*UpdateGlobalSystemConfigResponse, error)
	DeleteGlobalSystemConfig(context.Context, *DeleteGlobalSystemConfigRequest) (*DeleteGlobalSystemConfigResponse, error)
	GetGlobalSystemConfig(context.Context, *GetGlobalSystemConfigRequest) (*GetGlobalSystemConfigResponse, error)
	ListGlobalSystemConfig(context.Context, *ListGlobalSystemConfigRequest) (*ListGlobalSystemConfigResponse, error)

	CreateGlobalVrouterConfig(context.Context, *CreateGlobalVrouterConfigRequest) (*CreateGlobalVrouterConfigResponse, error)
	UpdateGlobalVrouterConfig(context.Context, *UpdateGlobalVrouterConfigRequest) (*UpdateGlobalVrouterConfigResponse, error)
	DeleteGlobalVrouterConfig(context.Context, *DeleteGlobalVrouterConfigRequest) (*DeleteGlobalVrouterConfigResponse, error)
	GetGlobalVrouterConfig(context.Context, *GetGlobalVrouterConfigRequest) (*GetGlobalVrouterConfigResponse, error)
	ListGlobalVrouterConfig(context.Context, *ListGlobalVrouterConfigRequest) (*ListGlobalVrouterConfigResponse, error)

	CreateInstanceIP(context.Context, *CreateInstanceIPRequest) (*CreateInstanceIPResponse, error)
	UpdateInstanceIP(context.Context, *UpdateInstanceIPRequest) (*UpdateInstanceIPResponse, error)
	DeleteInstanceIP(context.Context, *DeleteInstanceIPRequest) (*DeleteInstanceIPResponse, error)
	GetInstanceIP(context.Context, *GetInstanceIPRequest) (*GetInstanceIPResponse, error)
	ListInstanceIP(context.Context, *ListInstanceIPRequest) (*ListInstanceIPResponse, error)

	CreateInterfaceRouteTable(context.Context, *CreateInterfaceRouteTableRequest) (*CreateInterfaceRouteTableResponse, error)
	UpdateInterfaceRouteTable(context.Context, *UpdateInterfaceRouteTableRequest) (*UpdateInterfaceRouteTableResponse, error)
	DeleteInterfaceRouteTable(context.Context, *DeleteInterfaceRouteTableRequest) (*DeleteInterfaceRouteTableResponse, error)
	GetInterfaceRouteTable(context.Context, *GetInterfaceRouteTableRequest) (*GetInterfaceRouteTableResponse, error)
	ListInterfaceRouteTable(context.Context, *ListInterfaceRouteTableRequest) (*ListInterfaceRouteTableResponse, error)

	CreateLoadbalancerHealthmonitor(context.Context, *CreateLoadbalancerHealthmonitorRequest) (*CreateLoadbalancerHealthmonitorResponse, error)
	UpdateLoadbalancerHealthmonitor(context.Context, *UpdateLoadbalancerHealthmonitorRequest) (*UpdateLoadbalancerHealthmonitorResponse, error)
	DeleteLoadbalancerHealthmonitor(context.Context, *DeleteLoadbalancerHealthmonitorRequest) (*DeleteLoadbalancerHealthmonitorResponse, error)
	GetLoadbalancerHealthmonitor(context.Context, *GetLoadbalancerHealthmonitorRequest) (*GetLoadbalancerHealthmonitorResponse, error)
	ListLoadbalancerHealthmonitor(context.Context, *ListLoadbalancerHealthmonitorRequest) (*ListLoadbalancerHealthmonitorResponse, error)

	CreateLoadbalancerListener(context.Context, *CreateLoadbalancerListenerRequest) (*CreateLoadbalancerListenerResponse, error)
	UpdateLoadbalancerListener(context.Context, *UpdateLoadbalancerListenerRequest) (*UpdateLoadbalancerListenerResponse, error)
	DeleteLoadbalancerListener(context.Context, *DeleteLoadbalancerListenerRequest) (*DeleteLoadbalancerListenerResponse, error)
	GetLoadbalancerListener(context.Context, *GetLoadbalancerListenerRequest) (*GetLoadbalancerListenerResponse, error)
	ListLoadbalancerListener(context.Context, *ListLoadbalancerListenerRequest) (*ListLoadbalancerListenerResponse, error)

	CreateLoadbalancerMember(context.Context, *CreateLoadbalancerMemberRequest) (*CreateLoadbalancerMemberResponse, error)
	UpdateLoadbalancerMember(context.Context, *UpdateLoadbalancerMemberRequest) (*UpdateLoadbalancerMemberResponse, error)
	DeleteLoadbalancerMember(context.Context, *DeleteLoadbalancerMemberRequest) (*DeleteLoadbalancerMemberResponse, error)
	GetLoadbalancerMember(context.Context, *GetLoadbalancerMemberRequest) (*GetLoadbalancerMemberResponse, error)
	ListLoadbalancerMember(context.Context, *ListLoadbalancerMemberRequest) (*ListLoadbalancerMemberResponse, error)

	CreateLoadbalancerPool(context.Context, *CreateLoadbalancerPoolRequest) (*CreateLoadbalancerPoolResponse, error)
	UpdateLoadbalancerPool(context.Context, *UpdateLoadbalancerPoolRequest) (*UpdateLoadbalancerPoolResponse, error)
	DeleteLoadbalancerPool(context.Context, *DeleteLoadbalancerPoolRequest) (*DeleteLoadbalancerPoolResponse, error)
	GetLoadbalancerPool(context.Context, *GetLoadbalancerPoolRequest) (*GetLoadbalancerPoolResponse, error)
	ListLoadbalancerPool(context.Context, *ListLoadbalancerPoolRequest) (*ListLoadbalancerPoolResponse, error)

	CreateLoadbalancer(context.Context, *CreateLoadbalancerRequest) (*CreateLoadbalancerResponse, error)
	UpdateLoadbalancer(context.Context, *UpdateLoadbalancerRequest) (*UpdateLoadbalancerResponse, error)
	DeleteLoadbalancer(context.Context, *DeleteLoadbalancerRequest) (*DeleteLoadbalancerResponse, error)
	GetLoadbalancer(context.Context, *GetLoadbalancerRequest) (*GetLoadbalancerResponse, error)
	ListLoadbalancer(context.Context, *ListLoadbalancerRequest) (*ListLoadbalancerResponse, error)

	CreateLogicalInterface(context.Context, *CreateLogicalInterfaceRequest) (*CreateLogicalInterfaceResponse, error)
	UpdateLogicalInterface(context.Context, *UpdateLogicalInterfaceRequest) (*UpdateLogicalInterfaceResponse, error)
	DeleteLogicalInterface(context.Context, *DeleteLogicalInterfaceRequest) (*DeleteLogicalInterfaceResponse, error)
	GetLogicalInterface(context.Context, *GetLogicalInterfaceRequest) (*GetLogicalInterfaceResponse, error)
	ListLogicalInterface(context.Context, *ListLogicalInterfaceRequest) (*ListLogicalInterfaceResponse, error)

	CreateLogicalRouter(context.Context, *CreateLogicalRouterRequest) (*CreateLogicalRouterResponse, error)
	UpdateLogicalRouter(context.Context, *UpdateLogicalRouterRequest) (*UpdateLogicalRouterResponse, error)
	DeleteLogicalRouter(context.Context, *DeleteLogicalRouterRequest) (*DeleteLogicalRouterResponse, error)
	GetLogicalRouter(context.Context, *GetLogicalRouterRequest) (*GetLogicalRouterResponse, error)
	ListLogicalRouter(context.Context, *ListLogicalRouterRequest) (*ListLogicalRouterResponse, error)

	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	ListNamespace(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error)

	CreateNetworkDeviceConfig(context.Context, *CreateNetworkDeviceConfigRequest) (*CreateNetworkDeviceConfigResponse, error)
	UpdateNetworkDeviceConfig(context.Context, *UpdateNetworkDeviceConfigRequest) (*UpdateNetworkDeviceConfigResponse, error)
	DeleteNetworkDeviceConfig(context.Context, *DeleteNetworkDeviceConfigRequest) (*DeleteNetworkDeviceConfigResponse, error)
	GetNetworkDeviceConfig(context.Context, *GetNetworkDeviceConfigRequest) (*GetNetworkDeviceConfigResponse, error)
	ListNetworkDeviceConfig(context.Context, *ListNetworkDeviceConfigRequest) (*ListNetworkDeviceConfigResponse, error)

	CreateNetworkIpam(context.Context, *CreateNetworkIpamRequest) (*CreateNetworkIpamResponse, error)
	UpdateNetworkIpam(context.Context, *UpdateNetworkIpamRequest) (*UpdateNetworkIpamResponse, error)
	DeleteNetworkIpam(context.Context, *DeleteNetworkIpamRequest) (*DeleteNetworkIpamResponse, error)
	GetNetworkIpam(context.Context, *GetNetworkIpamRequest) (*GetNetworkIpamResponse, error)
	ListNetworkIpam(context.Context, *ListNetworkIpamRequest) (*ListNetworkIpamResponse, error)

	CreateNetworkPolicy(context.Context, *CreateNetworkPolicyRequest) (*CreateNetworkPolicyResponse, error)
	UpdateNetworkPolicy(context.Context, *UpdateNetworkPolicyRequest) (*UpdateNetworkPolicyResponse, error)
	DeleteNetworkPolicy(context.Context, *DeleteNetworkPolicyRequest) (*DeleteNetworkPolicyResponse, error)
	GetNetworkPolicy(context.Context, *GetNetworkPolicyRequest) (*GetNetworkPolicyResponse, error)
	ListNetworkPolicy(context.Context, *ListNetworkPolicyRequest) (*ListNetworkPolicyResponse, error)

	CreatePeeringPolicy(context.Context, *CreatePeeringPolicyRequest) (*CreatePeeringPolicyResponse, error)
	UpdatePeeringPolicy(context.Context, *UpdatePeeringPolicyRequest) (*UpdatePeeringPolicyResponse, error)
	DeletePeeringPolicy(context.Context, *DeletePeeringPolicyRequest) (*DeletePeeringPolicyResponse, error)
	GetPeeringPolicy(context.Context, *GetPeeringPolicyRequest) (*GetPeeringPolicyResponse, error)
	ListPeeringPolicy(context.Context, *ListPeeringPolicyRequest) (*ListPeeringPolicyResponse, error)

	CreatePhysicalInterface(context.Context, *CreatePhysicalInterfaceRequest) (*CreatePhysicalInterfaceResponse, error)
	UpdatePhysicalInterface(context.Context, *UpdatePhysicalInterfaceRequest) (*UpdatePhysicalInterfaceResponse, error)
	DeletePhysicalInterface(context.Context, *DeletePhysicalInterfaceRequest) (*DeletePhysicalInterfaceResponse, error)
	GetPhysicalInterface(context.Context, *GetPhysicalInterfaceRequest) (*GetPhysicalInterfaceResponse, error)
	ListPhysicalInterface(context.Context, *ListPhysicalInterfaceRequest) (*ListPhysicalInterfaceResponse, error)

	CreatePhysicalRouter(context.Context, *CreatePhysicalRouterRequest) (*CreatePhysicalRouterResponse, error)
	UpdatePhysicalRouter(context.Context, *UpdatePhysicalRouterRequest) (*UpdatePhysicalRouterResponse, error)
	DeletePhysicalRouter(context.Context, *DeletePhysicalRouterRequest) (*DeletePhysicalRouterResponse, error)
	GetPhysicalRouter(context.Context, *GetPhysicalRouterRequest) (*GetPhysicalRouterResponse, error)
	ListPhysicalRouter(context.Context, *ListPhysicalRouterRequest) (*ListPhysicalRouterResponse, error)

	CreatePolicyManagement(context.Context, *CreatePolicyManagementRequest) (*CreatePolicyManagementResponse, error)
	UpdatePolicyManagement(context.Context, *UpdatePolicyManagementRequest) (*UpdatePolicyManagementResponse, error)
	DeletePolicyManagement(context.Context, *DeletePolicyManagementRequest) (*DeletePolicyManagementResponse, error)
	GetPolicyManagement(context.Context, *GetPolicyManagementRequest) (*GetPolicyManagementResponse, error)
	ListPolicyManagement(context.Context, *ListPolicyManagementRequest) (*ListPolicyManagementResponse, error)

	CreatePortTuple(context.Context, *CreatePortTupleRequest) (*CreatePortTupleResponse, error)
	UpdatePortTuple(context.Context, *UpdatePortTupleRequest) (*UpdatePortTupleResponse, error)
	DeletePortTuple(context.Context, *DeletePortTupleRequest) (*DeletePortTupleResponse, error)
	GetPortTuple(context.Context, *GetPortTupleRequest) (*GetPortTupleResponse, error)
	ListPortTuple(context.Context, *ListPortTupleRequest) (*ListPortTupleResponse, error)

	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectResponse, error)

	CreateProviderAttachment(context.Context, *CreateProviderAttachmentRequest) (*CreateProviderAttachmentResponse, error)
	UpdateProviderAttachment(context.Context, *UpdateProviderAttachmentRequest) (*UpdateProviderAttachmentResponse, error)
	DeleteProviderAttachment(context.Context, *DeleteProviderAttachmentRequest) (*DeleteProviderAttachmentResponse, error)
	GetProviderAttachment(context.Context, *GetProviderAttachmentRequest) (*GetProviderAttachmentResponse, error)
	ListProviderAttachment(context.Context, *ListProviderAttachmentRequest) (*ListProviderAttachmentResponse, error)

	CreateQosConfig(context.Context, *CreateQosConfigRequest) (*CreateQosConfigResponse, error)
	UpdateQosConfig(context.Context, *UpdateQosConfigRequest) (*UpdateQosConfigResponse, error)
	DeleteQosConfig(context.Context, *DeleteQosConfigRequest) (*DeleteQosConfigResponse, error)
	GetQosConfig(context.Context, *GetQosConfigRequest) (*GetQosConfigResponse, error)
	ListQosConfig(context.Context, *ListQosConfigRequest) (*ListQosConfigResponse, error)

	CreateQosQueue(context.Context, *CreateQosQueueRequest) (*CreateQosQueueResponse, error)
	UpdateQosQueue(context.Context, *UpdateQosQueueRequest) (*UpdateQosQueueResponse, error)
	DeleteQosQueue(context.Context, *DeleteQosQueueRequest) (*DeleteQosQueueResponse, error)
	GetQosQueue(context.Context, *GetQosQueueRequest) (*GetQosQueueResponse, error)
	ListQosQueue(context.Context, *ListQosQueueRequest) (*ListQosQueueResponse, error)

	CreateRouteAggregate(context.Context, *CreateRouteAggregateRequest) (*CreateRouteAggregateResponse, error)
	UpdateRouteAggregate(context.Context, *UpdateRouteAggregateRequest) (*UpdateRouteAggregateResponse, error)
	DeleteRouteAggregate(context.Context, *DeleteRouteAggregateRequest) (*DeleteRouteAggregateResponse, error)
	GetRouteAggregate(context.Context, *GetRouteAggregateRequest) (*GetRouteAggregateResponse, error)
	ListRouteAggregate(context.Context, *ListRouteAggregateRequest) (*ListRouteAggregateResponse, error)

	CreateRouteTable(context.Context, *CreateRouteTableRequest) (*CreateRouteTableResponse, error)
	UpdateRouteTable(context.Context, *UpdateRouteTableRequest) (*UpdateRouteTableResponse, error)
	DeleteRouteTable(context.Context, *DeleteRouteTableRequest) (*DeleteRouteTableResponse, error)
	GetRouteTable(context.Context, *GetRouteTableRequest) (*GetRouteTableResponse, error)
	ListRouteTable(context.Context, *ListRouteTableRequest) (*ListRouteTableResponse, error)

	CreateRouteTarget(context.Context, *CreateRouteTargetRequest) (*CreateRouteTargetResponse, error)
	UpdateRouteTarget(context.Context, *UpdateRouteTargetRequest) (*UpdateRouteTargetResponse, error)
	DeleteRouteTarget(context.Context, *DeleteRouteTargetRequest) (*DeleteRouteTargetResponse, error)
	GetRouteTarget(context.Context, *GetRouteTargetRequest) (*GetRouteTargetResponse, error)
	ListRouteTarget(context.Context, *ListRouteTargetRequest) (*ListRouteTargetResponse, error)

	CreateRoutingInstance(context.Context, *CreateRoutingInstanceRequest) (*CreateRoutingInstanceResponse, error)
	UpdateRoutingInstance(context.Context, *UpdateRoutingInstanceRequest) (*UpdateRoutingInstanceResponse, error)
	DeleteRoutingInstance(context.Context, *DeleteRoutingInstanceRequest) (*DeleteRoutingInstanceResponse, error)
	GetRoutingInstance(context.Context, *GetRoutingInstanceRequest) (*GetRoutingInstanceResponse, error)
	ListRoutingInstance(context.Context, *ListRoutingInstanceRequest) (*ListRoutingInstanceResponse, error)

	CreateRoutingPolicy(context.Context, *CreateRoutingPolicyRequest) (*CreateRoutingPolicyResponse, error)
	UpdateRoutingPolicy(context.Context, *UpdateRoutingPolicyRequest) (*UpdateRoutingPolicyResponse, error)
	DeleteRoutingPolicy(context.Context, *DeleteRoutingPolicyRequest) (*DeleteRoutingPolicyResponse, error)
	GetRoutingPolicy(context.Context, *GetRoutingPolicyRequest) (*GetRoutingPolicyResponse, error)
	ListRoutingPolicy(context.Context, *ListRoutingPolicyRequest) (*ListRoutingPolicyResponse, error)

	CreateSecurityGroup(context.Context, *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error)
	UpdateSecurityGroup(context.Context, *UpdateSecurityGroupRequest) (*UpdateSecurityGroupResponse, error)
	DeleteSecurityGroup(context.Context, *DeleteSecurityGroupRequest) (*DeleteSecurityGroupResponse, error)
	GetSecurityGroup(context.Context, *GetSecurityGroupRequest) (*GetSecurityGroupResponse, error)
	ListSecurityGroup(context.Context, *ListSecurityGroupRequest) (*ListSecurityGroupResponse, error)

	CreateSecurityLoggingObject(context.Context, *CreateSecurityLoggingObjectRequest) (*CreateSecurityLoggingObjectResponse, error)
	UpdateSecurityLoggingObject(context.Context, *UpdateSecurityLoggingObjectRequest) (*UpdateSecurityLoggingObjectResponse, error)
	DeleteSecurityLoggingObject(context.Context, *DeleteSecurityLoggingObjectRequest) (*DeleteSecurityLoggingObjectResponse, error)
	GetSecurityLoggingObject(context.Context, *GetSecurityLoggingObjectRequest) (*GetSecurityLoggingObjectResponse, error)
	ListSecurityLoggingObject(context.Context, *ListSecurityLoggingObjectRequest) (*ListSecurityLoggingObjectResponse, error)

	CreateServiceAppliance(context.Context, *CreateServiceApplianceRequest) (*CreateServiceApplianceResponse, error)
	UpdateServiceAppliance(context.Context, *UpdateServiceApplianceRequest) (*UpdateServiceApplianceResponse, error)
	DeleteServiceAppliance(context.Context, *DeleteServiceApplianceRequest) (*DeleteServiceApplianceResponse, error)
	GetServiceAppliance(context.Context, *GetServiceApplianceRequest) (*GetServiceApplianceResponse, error)
	ListServiceAppliance(context.Context, *ListServiceApplianceRequest) (*ListServiceApplianceResponse, error)

	CreateServiceApplianceSet(context.Context, *CreateServiceApplianceSetRequest) (*CreateServiceApplianceSetResponse, error)
	UpdateServiceApplianceSet(context.Context, *UpdateServiceApplianceSetRequest) (*UpdateServiceApplianceSetResponse, error)
	DeleteServiceApplianceSet(context.Context, *DeleteServiceApplianceSetRequest) (*DeleteServiceApplianceSetResponse, error)
	GetServiceApplianceSet(context.Context, *GetServiceApplianceSetRequest) (*GetServiceApplianceSetResponse, error)
	ListServiceApplianceSet(context.Context, *ListServiceApplianceSetRequest) (*ListServiceApplianceSetResponse, error)

	CreateServiceConnectionModule(context.Context, *CreateServiceConnectionModuleRequest) (*CreateServiceConnectionModuleResponse, error)
	UpdateServiceConnectionModule(context.Context, *UpdateServiceConnectionModuleRequest) (*UpdateServiceConnectionModuleResponse, error)
	DeleteServiceConnectionModule(context.Context, *DeleteServiceConnectionModuleRequest) (*DeleteServiceConnectionModuleResponse, error)
	GetServiceConnectionModule(context.Context, *GetServiceConnectionModuleRequest) (*GetServiceConnectionModuleResponse, error)
	ListServiceConnectionModule(context.Context, *ListServiceConnectionModuleRequest) (*ListServiceConnectionModuleResponse, error)

	CreateServiceEndpoint(context.Context, *CreateServiceEndpointRequest) (*CreateServiceEndpointResponse, error)
	UpdateServiceEndpoint(context.Context, *UpdateServiceEndpointRequest) (*UpdateServiceEndpointResponse, error)
	DeleteServiceEndpoint(context.Context, *DeleteServiceEndpointRequest) (*DeleteServiceEndpointResponse, error)
	GetServiceEndpoint(context.Context, *GetServiceEndpointRequest) (*GetServiceEndpointResponse, error)
	ListServiceEndpoint(context.Context, *ListServiceEndpointRequest) (*ListServiceEndpointResponse, error)

	CreateServiceGroup(context.Context, *CreateServiceGroupRequest) (*CreateServiceGroupResponse, error)
	UpdateServiceGroup(context.Context, *UpdateServiceGroupRequest) (*UpdateServiceGroupResponse, error)
	DeleteServiceGroup(context.Context, *DeleteServiceGroupRequest) (*DeleteServiceGroupResponse, error)
	GetServiceGroup(context.Context, *GetServiceGroupRequest) (*GetServiceGroupResponse, error)
	ListServiceGroup(context.Context, *ListServiceGroupRequest) (*ListServiceGroupResponse, error)

	CreateServiceHealthCheck(context.Context, *CreateServiceHealthCheckRequest) (*CreateServiceHealthCheckResponse, error)
	UpdateServiceHealthCheck(context.Context, *UpdateServiceHealthCheckRequest) (*UpdateServiceHealthCheckResponse, error)
	DeleteServiceHealthCheck(context.Context, *DeleteServiceHealthCheckRequest) (*DeleteServiceHealthCheckResponse, error)
	GetServiceHealthCheck(context.Context, *GetServiceHealthCheckRequest) (*GetServiceHealthCheckResponse, error)
	ListServiceHealthCheck(context.Context, *ListServiceHealthCheckRequest) (*ListServiceHealthCheckResponse, error)

	CreateServiceInstance(context.Context, *CreateServiceInstanceRequest) (*CreateServiceInstanceResponse, error)
	UpdateServiceInstance(context.Context, *UpdateServiceInstanceRequest) (*UpdateServiceInstanceResponse, error)
	DeleteServiceInstance(context.Context, *DeleteServiceInstanceRequest) (*DeleteServiceInstanceResponse, error)
	GetServiceInstance(context.Context, *GetServiceInstanceRequest) (*GetServiceInstanceResponse, error)
	ListServiceInstance(context.Context, *ListServiceInstanceRequest) (*ListServiceInstanceResponse, error)

	CreateServiceObject(context.Context, *CreateServiceObjectRequest) (*CreateServiceObjectResponse, error)
	UpdateServiceObject(context.Context, *UpdateServiceObjectRequest) (*UpdateServiceObjectResponse, error)
	DeleteServiceObject(context.Context, *DeleteServiceObjectRequest) (*DeleteServiceObjectResponse, error)
	GetServiceObject(context.Context, *GetServiceObjectRequest) (*GetServiceObjectResponse, error)
	ListServiceObject(context.Context, *ListServiceObjectRequest) (*ListServiceObjectResponse, error)

	CreateServiceTemplate(context.Context, *CreateServiceTemplateRequest) (*CreateServiceTemplateResponse, error)
	UpdateServiceTemplate(context.Context, *UpdateServiceTemplateRequest) (*UpdateServiceTemplateResponse, error)
	DeleteServiceTemplate(context.Context, *DeleteServiceTemplateRequest) (*DeleteServiceTemplateResponse, error)
	GetServiceTemplate(context.Context, *GetServiceTemplateRequest) (*GetServiceTemplateResponse, error)
	ListServiceTemplate(context.Context, *ListServiceTemplateRequest) (*ListServiceTemplateResponse, error)

	CreateSubnet(context.Context, *CreateSubnetRequest) (*CreateSubnetResponse, error)
	UpdateSubnet(context.Context, *UpdateSubnetRequest) (*UpdateSubnetResponse, error)
	DeleteSubnet(context.Context, *DeleteSubnetRequest) (*DeleteSubnetResponse, error)
	GetSubnet(context.Context, *GetSubnetRequest) (*GetSubnetResponse, error)
	ListSubnet(context.Context, *ListSubnetRequest) (*ListSubnetResponse, error)

	CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error)
	GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error)
	ListTag(context.Context, *ListTagRequest) (*ListTagResponse, error)

	CreateTagType(context.Context, *CreateTagTypeRequest) (*CreateTagTypeResponse, error)
	UpdateTagType(context.Context, *UpdateTagTypeRequest) (*UpdateTagTypeResponse, error)
	DeleteTagType(context.Context, *DeleteTagTypeRequest) (*DeleteTagTypeResponse, error)
	GetTagType(context.Context, *GetTagTypeRequest) (*GetTagTypeResponse, error)
	ListTagType(context.Context, *ListTagTypeRequest) (*ListTagTypeResponse, error)

	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)

	CreateVirtualDNSRecord(context.Context, *CreateVirtualDNSRecordRequest) (*CreateVirtualDNSRecordResponse, error)
	UpdateVirtualDNSRecord(context.Context, *UpdateVirtualDNSRecordRequest) (*UpdateVirtualDNSRecordResponse, error)
	DeleteVirtualDNSRecord(context.Context, *DeleteVirtualDNSRecordRequest) (*DeleteVirtualDNSRecordResponse, error)
	GetVirtualDNSRecord(context.Context, *GetVirtualDNSRecordRequest) (*GetVirtualDNSRecordResponse, error)
	ListVirtualDNSRecord(context.Context, *ListVirtualDNSRecordRequest) (*ListVirtualDNSRecordResponse, error)

	CreateVirtualDNS(context.Context, *CreateVirtualDNSRequest) (*CreateVirtualDNSResponse, error)
	UpdateVirtualDNS(context.Context, *UpdateVirtualDNSRequest) (*UpdateVirtualDNSResponse, error)
	DeleteVirtualDNS(context.Context, *DeleteVirtualDNSRequest) (*DeleteVirtualDNSResponse, error)
	GetVirtualDNS(context.Context, *GetVirtualDNSRequest) (*GetVirtualDNSResponse, error)
	ListVirtualDNS(context.Context, *ListVirtualDNSRequest) (*ListVirtualDNSResponse, error)

	CreateVirtualIP(context.Context, *CreateVirtualIPRequest) (*CreateVirtualIPResponse, error)
	UpdateVirtualIP(context.Context, *UpdateVirtualIPRequest) (*UpdateVirtualIPResponse, error)
	DeleteVirtualIP(context.Context, *DeleteVirtualIPRequest) (*DeleteVirtualIPResponse, error)
	GetVirtualIP(context.Context, *GetVirtualIPRequest) (*GetVirtualIPResponse, error)
	ListVirtualIP(context.Context, *ListVirtualIPRequest) (*ListVirtualIPResponse, error)

	CreateVirtualMachineInterface(context.Context, *CreateVirtualMachineInterfaceRequest) (*CreateVirtualMachineInterfaceResponse, error)
	UpdateVirtualMachineInterface(context.Context, *UpdateVirtualMachineInterfaceRequest) (*UpdateVirtualMachineInterfaceResponse, error)
	DeleteVirtualMachineInterface(context.Context, *DeleteVirtualMachineInterfaceRequest) (*DeleteVirtualMachineInterfaceResponse, error)
	GetVirtualMachineInterface(context.Context, *GetVirtualMachineInterfaceRequest) (*GetVirtualMachineInterfaceResponse, error)
	ListVirtualMachineInterface(context.Context, *ListVirtualMachineInterfaceRequest) (*ListVirtualMachineInterfaceResponse, error)

	CreateVirtualMachine(context.Context, *CreateVirtualMachineRequest) (*CreateVirtualMachineResponse, error)
	UpdateVirtualMachine(context.Context, *UpdateVirtualMachineRequest) (*UpdateVirtualMachineResponse, error)
	DeleteVirtualMachine(context.Context, *DeleteVirtualMachineRequest) (*DeleteVirtualMachineResponse, error)
	GetVirtualMachine(context.Context, *GetVirtualMachineRequest) (*GetVirtualMachineResponse, error)
	ListVirtualMachine(context.Context, *ListVirtualMachineRequest) (*ListVirtualMachineResponse, error)

	CreateVirtualNetwork(context.Context, *CreateVirtualNetworkRequest) (*CreateVirtualNetworkResponse, error)
	UpdateVirtualNetwork(context.Context, *UpdateVirtualNetworkRequest) (*UpdateVirtualNetworkResponse, error)
	DeleteVirtualNetwork(context.Context, *DeleteVirtualNetworkRequest) (*DeleteVirtualNetworkResponse, error)
	GetVirtualNetwork(context.Context, *GetVirtualNetworkRequest) (*GetVirtualNetworkResponse, error)
	ListVirtualNetwork(context.Context, *ListVirtualNetworkRequest) (*ListVirtualNetworkResponse, error)

	CreateVirtualRouter(context.Context, *CreateVirtualRouterRequest) (*CreateVirtualRouterResponse, error)
	UpdateVirtualRouter(context.Context, *UpdateVirtualRouterRequest) (*UpdateVirtualRouterResponse, error)
	DeleteVirtualRouter(context.Context, *DeleteVirtualRouterRequest) (*DeleteVirtualRouterResponse, error)
	GetVirtualRouter(context.Context, *GetVirtualRouterRequest) (*GetVirtualRouterResponse, error)
	ListVirtualRouter(context.Context, *ListVirtualRouterRequest) (*ListVirtualRouterResponse, error)

	CreateAppformixNode(context.Context, *CreateAppformixNodeRequest) (*CreateAppformixNodeResponse, error)
	UpdateAppformixNode(context.Context, *UpdateAppformixNodeRequest) (*UpdateAppformixNodeResponse, error)
	DeleteAppformixNode(context.Context, *DeleteAppformixNodeRequest) (*DeleteAppformixNodeResponse, error)
	GetAppformixNode(context.Context, *GetAppformixNodeRequest) (*GetAppformixNodeResponse, error)
	ListAppformixNode(context.Context, *ListAppformixNodeRequest) (*ListAppformixNodeResponse, error)

	CreateBaremetalNode(context.Context, *CreateBaremetalNodeRequest) (*CreateBaremetalNodeResponse, error)
	UpdateBaremetalNode(context.Context, *UpdateBaremetalNodeRequest) (*UpdateBaremetalNodeResponse, error)
	DeleteBaremetalNode(context.Context, *DeleteBaremetalNodeRequest) (*DeleteBaremetalNodeResponse, error)
	GetBaremetalNode(context.Context, *GetBaremetalNodeRequest) (*GetBaremetalNodeResponse, error)
	ListBaremetalNode(context.Context, *ListBaremetalNodeRequest) (*ListBaremetalNodeResponse, error)

	CreateBaremetalPort(context.Context, *CreateBaremetalPortRequest) (*CreateBaremetalPortResponse, error)
	UpdateBaremetalPort(context.Context, *UpdateBaremetalPortRequest) (*UpdateBaremetalPortResponse, error)
	DeleteBaremetalPort(context.Context, *DeleteBaremetalPortRequest) (*DeleteBaremetalPortResponse, error)
	GetBaremetalPort(context.Context, *GetBaremetalPortRequest) (*GetBaremetalPortResponse, error)
	ListBaremetalPort(context.Context, *ListBaremetalPortRequest) (*ListBaremetalPortResponse, error)

	CreateContrailAnalyticsDatabaseNode(context.Context, *CreateContrailAnalyticsDatabaseNodeRequest) (*CreateContrailAnalyticsDatabaseNodeResponse, error)
	UpdateContrailAnalyticsDatabaseNode(context.Context, *UpdateContrailAnalyticsDatabaseNodeRequest) (*UpdateContrailAnalyticsDatabaseNodeResponse, error)
	DeleteContrailAnalyticsDatabaseNode(context.Context, *DeleteContrailAnalyticsDatabaseNodeRequest) (*DeleteContrailAnalyticsDatabaseNodeResponse, error)
	GetContrailAnalyticsDatabaseNode(context.Context, *GetContrailAnalyticsDatabaseNodeRequest) (*GetContrailAnalyticsDatabaseNodeResponse, error)
	ListContrailAnalyticsDatabaseNode(context.Context, *ListContrailAnalyticsDatabaseNodeRequest) (*ListContrailAnalyticsDatabaseNodeResponse, error)

	CreateContrailAnalyticsNode(context.Context, *CreateContrailAnalyticsNodeRequest) (*CreateContrailAnalyticsNodeResponse, error)
	UpdateContrailAnalyticsNode(context.Context, *UpdateContrailAnalyticsNodeRequest) (*UpdateContrailAnalyticsNodeResponse, error)
	DeleteContrailAnalyticsNode(context.Context, *DeleteContrailAnalyticsNodeRequest) (*DeleteContrailAnalyticsNodeResponse, error)
	GetContrailAnalyticsNode(context.Context, *GetContrailAnalyticsNodeRequest) (*GetContrailAnalyticsNodeResponse, error)
	ListContrailAnalyticsNode(context.Context, *ListContrailAnalyticsNodeRequest) (*ListContrailAnalyticsNodeResponse, error)

	CreateContrailCluster(context.Context, *CreateContrailClusterRequest) (*CreateContrailClusterResponse, error)
	UpdateContrailCluster(context.Context, *UpdateContrailClusterRequest) (*UpdateContrailClusterResponse, error)
	DeleteContrailCluster(context.Context, *DeleteContrailClusterRequest) (*DeleteContrailClusterResponse, error)
	GetContrailCluster(context.Context, *GetContrailClusterRequest) (*GetContrailClusterResponse, error)
	ListContrailCluster(context.Context, *ListContrailClusterRequest) (*ListContrailClusterResponse, error)

	CreateContrailConfigDatabaseNode(context.Context, *CreateContrailConfigDatabaseNodeRequest) (*CreateContrailConfigDatabaseNodeResponse, error)
	UpdateContrailConfigDatabaseNode(context.Context, *UpdateContrailConfigDatabaseNodeRequest) (*UpdateContrailConfigDatabaseNodeResponse, error)
	DeleteContrailConfigDatabaseNode(context.Context, *DeleteContrailConfigDatabaseNodeRequest) (*DeleteContrailConfigDatabaseNodeResponse, error)
	GetContrailConfigDatabaseNode(context.Context, *GetContrailConfigDatabaseNodeRequest) (*GetContrailConfigDatabaseNodeResponse, error)
	ListContrailConfigDatabaseNode(context.Context, *ListContrailConfigDatabaseNodeRequest) (*ListContrailConfigDatabaseNodeResponse, error)

	CreateContrailConfigNode(context.Context, *CreateContrailConfigNodeRequest) (*CreateContrailConfigNodeResponse, error)
	UpdateContrailConfigNode(context.Context, *UpdateContrailConfigNodeRequest) (*UpdateContrailConfigNodeResponse, error)
	DeleteContrailConfigNode(context.Context, *DeleteContrailConfigNodeRequest) (*DeleteContrailConfigNodeResponse, error)
	GetContrailConfigNode(context.Context, *GetContrailConfigNodeRequest) (*GetContrailConfigNodeResponse, error)
	ListContrailConfigNode(context.Context, *ListContrailConfigNodeRequest) (*ListContrailConfigNodeResponse, error)

	CreateContrailControlNode(context.Context, *CreateContrailControlNodeRequest) (*CreateContrailControlNodeResponse, error)
	UpdateContrailControlNode(context.Context, *UpdateContrailControlNodeRequest) (*UpdateContrailControlNodeResponse, error)
	DeleteContrailControlNode(context.Context, *DeleteContrailControlNodeRequest) (*DeleteContrailControlNodeResponse, error)
	GetContrailControlNode(context.Context, *GetContrailControlNodeRequest) (*GetContrailControlNodeResponse, error)
	ListContrailControlNode(context.Context, *ListContrailControlNodeRequest) (*ListContrailControlNodeResponse, error)

	CreateContrailStorageNode(context.Context, *CreateContrailStorageNodeRequest) (*CreateContrailStorageNodeResponse, error)
	UpdateContrailStorageNode(context.Context, *UpdateContrailStorageNodeRequest) (*UpdateContrailStorageNodeResponse, error)
	DeleteContrailStorageNode(context.Context, *DeleteContrailStorageNodeRequest) (*DeleteContrailStorageNodeResponse, error)
	GetContrailStorageNode(context.Context, *GetContrailStorageNodeRequest) (*GetContrailStorageNodeResponse, error)
	ListContrailStorageNode(context.Context, *ListContrailStorageNodeRequest) (*ListContrailStorageNodeResponse, error)

	CreateContrailVrouterNode(context.Context, *CreateContrailVrouterNodeRequest) (*CreateContrailVrouterNodeResponse, error)
	UpdateContrailVrouterNode(context.Context, *UpdateContrailVrouterNodeRequest) (*UpdateContrailVrouterNodeResponse, error)
	DeleteContrailVrouterNode(context.Context, *DeleteContrailVrouterNodeRequest) (*DeleteContrailVrouterNodeResponse, error)
	GetContrailVrouterNode(context.Context, *GetContrailVrouterNodeRequest) (*GetContrailVrouterNodeResponse, error)
	ListContrailVrouterNode(context.Context, *ListContrailVrouterNodeRequest) (*ListContrailVrouterNodeResponse, error)

	CreateContrailControllerNode(context.Context, *CreateContrailControllerNodeRequest) (*CreateContrailControllerNodeResponse, error)
	UpdateContrailControllerNode(context.Context, *UpdateContrailControllerNodeRequest) (*UpdateContrailControllerNodeResponse, error)
	DeleteContrailControllerNode(context.Context, *DeleteContrailControllerNodeRequest) (*DeleteContrailControllerNodeResponse, error)
	GetContrailControllerNode(context.Context, *GetContrailControllerNodeRequest) (*GetContrailControllerNodeResponse, error)
	ListContrailControllerNode(context.Context, *ListContrailControllerNodeRequest) (*ListContrailControllerNodeResponse, error)

	CreateDashboard(context.Context, *CreateDashboardRequest) (*CreateDashboardResponse, error)
	UpdateDashboard(context.Context, *UpdateDashboardRequest) (*UpdateDashboardResponse, error)
	DeleteDashboard(context.Context, *DeleteDashboardRequest) (*DeleteDashboardResponse, error)
	GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error)
	ListDashboard(context.Context, *ListDashboardRequest) (*ListDashboardResponse, error)

	CreateFlavor(context.Context, *CreateFlavorRequest) (*CreateFlavorResponse, error)
	UpdateFlavor(context.Context, *UpdateFlavorRequest) (*UpdateFlavorResponse, error)
	DeleteFlavor(context.Context, *DeleteFlavorRequest) (*DeleteFlavorResponse, error)
	GetFlavor(context.Context, *GetFlavorRequest) (*GetFlavorResponse, error)
	ListFlavor(context.Context, *ListFlavorRequest) (*ListFlavorResponse, error)

	CreateOsImage(context.Context, *CreateOsImageRequest) (*CreateOsImageResponse, error)
	UpdateOsImage(context.Context, *UpdateOsImageRequest) (*UpdateOsImageResponse, error)
	DeleteOsImage(context.Context, *DeleteOsImageRequest) (*DeleteOsImageResponse, error)
	GetOsImage(context.Context, *GetOsImageRequest) (*GetOsImageResponse, error)
	ListOsImage(context.Context, *ListOsImageRequest) (*ListOsImageResponse, error)

	CreateKeypair(context.Context, *CreateKeypairRequest) (*CreateKeypairResponse, error)
	UpdateKeypair(context.Context, *UpdateKeypairRequest) (*UpdateKeypairResponse, error)
	DeleteKeypair(context.Context, *DeleteKeypairRequest) (*DeleteKeypairResponse, error)
	GetKeypair(context.Context, *GetKeypairRequest) (*GetKeypairResponse, error)
	ListKeypair(context.Context, *ListKeypairRequest) (*ListKeypairResponse, error)

	CreateKubernetesMasterNode(context.Context, *CreateKubernetesMasterNodeRequest) (*CreateKubernetesMasterNodeResponse, error)
	UpdateKubernetesMasterNode(context.Context, *UpdateKubernetesMasterNodeRequest) (*UpdateKubernetesMasterNodeResponse, error)
	DeleteKubernetesMasterNode(context.Context, *DeleteKubernetesMasterNodeRequest) (*DeleteKubernetesMasterNodeResponse, error)
	GetKubernetesMasterNode(context.Context, *GetKubernetesMasterNodeRequest) (*GetKubernetesMasterNodeResponse, error)
	ListKubernetesMasterNode(context.Context, *ListKubernetesMasterNodeRequest) (*ListKubernetesMasterNodeResponse, error)

	CreateKubernetesNode(context.Context, *CreateKubernetesNodeRequest) (*CreateKubernetesNodeResponse, error)
	UpdateKubernetesNode(context.Context, *UpdateKubernetesNodeRequest) (*UpdateKubernetesNodeResponse, error)
	DeleteKubernetesNode(context.Context, *DeleteKubernetesNodeRequest) (*DeleteKubernetesNodeResponse, error)
	GetKubernetesNode(context.Context, *GetKubernetesNodeRequest) (*GetKubernetesNodeResponse, error)
	ListKubernetesNode(context.Context, *ListKubernetesNodeRequest) (*ListKubernetesNodeResponse, error)

	CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error)
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
	ListLocation(context.Context, *ListLocationRequest) (*ListLocationResponse, error)

	CreateNode(context.Context, *CreateNodeRequest) (*CreateNodeResponse, error)
	UpdateNode(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
	DeleteNode(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	ListNode(context.Context, *ListNodeRequest) (*ListNodeResponse, error)

	CreateServer(context.Context, *CreateServerRequest) (*CreateServerResponse, error)
	UpdateServer(context.Context, *UpdateServerRequest) (*UpdateServerResponse, error)
	DeleteServer(context.Context, *DeleteServerRequest) (*DeleteServerResponse, error)
	GetServer(context.Context, *GetServerRequest) (*GetServerResponse, error)
	ListServer(context.Context, *ListServerRequest) (*ListServerResponse, error)

	CreateVPNGroup(context.Context, *CreateVPNGroupRequest) (*CreateVPNGroupResponse, error)
	UpdateVPNGroup(context.Context, *UpdateVPNGroupRequest) (*UpdateVPNGroupResponse, error)
	DeleteVPNGroup(context.Context, *DeleteVPNGroupRequest) (*DeleteVPNGroupResponse, error)
	GetVPNGroup(context.Context, *GetVPNGroupRequest) (*GetVPNGroupResponse, error)
	ListVPNGroup(context.Context, *ListVPNGroupRequest) (*ListVPNGroupResponse, error)

	CreateWidget(context.Context, *CreateWidgetRequest) (*CreateWidgetResponse, error)
	UpdateWidget(context.Context, *UpdateWidgetRequest) (*UpdateWidgetResponse, error)
	DeleteWidget(context.Context, *DeleteWidgetRequest) (*DeleteWidgetResponse, error)
	GetWidget(context.Context, *GetWidgetRequest) (*GetWidgetResponse, error)
	ListWidget(context.Context, *ListWidgetRequest) (*ListWidgetResponse, error)
}

//Chain setup chain of services.
func Chain(services []Service) {
	if len(services) < 2 {
		return
	}
	previous := services[0]
	for i := 1; i < len(services); i++ {
		current := services[i]
		previous.SetNext(current)
		previous = current
	}
}

type BaseService struct {
	next Service
}

func (service *BaseService) Next() Service {
	return service.next
}

func (service *BaseService) SetNext(next Service) {
	service.next = next
}

func (service *BaseService) CreateAccessControlList(ctx context.Context, request *CreateAccessControlListRequest) (*CreateAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAccessControlList(ctx, request)
}
func (service *BaseService) UpdateAccessControlList(ctx context.Context, request *UpdateAccessControlListRequest) (*UpdateAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAccessControlList(ctx, request)
}
func (service *BaseService) DeleteAccessControlList(ctx context.Context, request *DeleteAccessControlListRequest) (*DeleteAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAccessControlList(ctx, request)
}
func (service *BaseService) GetAccessControlList(ctx context.Context, request *GetAccessControlListRequest) (*GetAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAccessControlList(ctx, request)
}
func (service *BaseService) ListAccessControlList(ctx context.Context, request *ListAccessControlListRequest) (*ListAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAccessControlList(ctx, request)
}

func (service *BaseService) CreateAddressGroup(ctx context.Context, request *CreateAddressGroupRequest) (*CreateAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAddressGroup(ctx, request)
}
func (service *BaseService) UpdateAddressGroup(ctx context.Context, request *UpdateAddressGroupRequest) (*UpdateAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAddressGroup(ctx, request)
}
func (service *BaseService) DeleteAddressGroup(ctx context.Context, request *DeleteAddressGroupRequest) (*DeleteAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAddressGroup(ctx, request)
}
func (service *BaseService) GetAddressGroup(ctx context.Context, request *GetAddressGroupRequest) (*GetAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAddressGroup(ctx, request)
}
func (service *BaseService) ListAddressGroup(ctx context.Context, request *ListAddressGroupRequest) (*ListAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAddressGroup(ctx, request)
}

func (service *BaseService) CreateAlarm(ctx context.Context, request *CreateAlarmRequest) (*CreateAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAlarm(ctx, request)
}
func (service *BaseService) UpdateAlarm(ctx context.Context, request *UpdateAlarmRequest) (*UpdateAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAlarm(ctx, request)
}
func (service *BaseService) DeleteAlarm(ctx context.Context, request *DeleteAlarmRequest) (*DeleteAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAlarm(ctx, request)
}
func (service *BaseService) GetAlarm(ctx context.Context, request *GetAlarmRequest) (*GetAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAlarm(ctx, request)
}
func (service *BaseService) ListAlarm(ctx context.Context, request *ListAlarmRequest) (*ListAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAlarm(ctx, request)
}

func (service *BaseService) CreateAliasIPPool(ctx context.Context, request *CreateAliasIPPoolRequest) (*CreateAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAliasIPPool(ctx, request)
}
func (service *BaseService) UpdateAliasIPPool(ctx context.Context, request *UpdateAliasIPPoolRequest) (*UpdateAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAliasIPPool(ctx, request)
}
func (service *BaseService) DeleteAliasIPPool(ctx context.Context, request *DeleteAliasIPPoolRequest) (*DeleteAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAliasIPPool(ctx, request)
}
func (service *BaseService) GetAliasIPPool(ctx context.Context, request *GetAliasIPPoolRequest) (*GetAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAliasIPPool(ctx, request)
}
func (service *BaseService) ListAliasIPPool(ctx context.Context, request *ListAliasIPPoolRequest) (*ListAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAliasIPPool(ctx, request)
}

func (service *BaseService) CreateAliasIP(ctx context.Context, request *CreateAliasIPRequest) (*CreateAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAliasIP(ctx, request)
}
func (service *BaseService) UpdateAliasIP(ctx context.Context, request *UpdateAliasIPRequest) (*UpdateAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAliasIP(ctx, request)
}
func (service *BaseService) DeleteAliasIP(ctx context.Context, request *DeleteAliasIPRequest) (*DeleteAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAliasIP(ctx, request)
}
func (service *BaseService) GetAliasIP(ctx context.Context, request *GetAliasIPRequest) (*GetAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAliasIP(ctx, request)
}
func (service *BaseService) ListAliasIP(ctx context.Context, request *ListAliasIPRequest) (*ListAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAliasIP(ctx, request)
}

func (service *BaseService) CreateAnalyticsNode(ctx context.Context, request *CreateAnalyticsNodeRequest) (*CreateAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAnalyticsNode(ctx, request)
}
func (service *BaseService) UpdateAnalyticsNode(ctx context.Context, request *UpdateAnalyticsNodeRequest) (*UpdateAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAnalyticsNode(ctx, request)
}
func (service *BaseService) DeleteAnalyticsNode(ctx context.Context, request *DeleteAnalyticsNodeRequest) (*DeleteAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAnalyticsNode(ctx, request)
}
func (service *BaseService) GetAnalyticsNode(ctx context.Context, request *GetAnalyticsNodeRequest) (*GetAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAnalyticsNode(ctx, request)
}
func (service *BaseService) ListAnalyticsNode(ctx context.Context, request *ListAnalyticsNodeRequest) (*ListAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAnalyticsNode(ctx, request)
}

func (service *BaseService) CreateAPIAccessList(ctx context.Context, request *CreateAPIAccessListRequest) (*CreateAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAPIAccessList(ctx, request)
}
func (service *BaseService) UpdateAPIAccessList(ctx context.Context, request *UpdateAPIAccessListRequest) (*UpdateAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAPIAccessList(ctx, request)
}
func (service *BaseService) DeleteAPIAccessList(ctx context.Context, request *DeleteAPIAccessListRequest) (*DeleteAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAPIAccessList(ctx, request)
}
func (service *BaseService) GetAPIAccessList(ctx context.Context, request *GetAPIAccessListRequest) (*GetAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAPIAccessList(ctx, request)
}
func (service *BaseService) ListAPIAccessList(ctx context.Context, request *ListAPIAccessListRequest) (*ListAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAPIAccessList(ctx, request)
}

func (service *BaseService) CreateApplicationPolicySet(ctx context.Context, request *CreateApplicationPolicySetRequest) (*CreateApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateApplicationPolicySet(ctx, request)
}
func (service *BaseService) UpdateApplicationPolicySet(ctx context.Context, request *UpdateApplicationPolicySetRequest) (*UpdateApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateApplicationPolicySet(ctx, request)
}
func (service *BaseService) DeleteApplicationPolicySet(ctx context.Context, request *DeleteApplicationPolicySetRequest) (*DeleteApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteApplicationPolicySet(ctx, request)
}
func (service *BaseService) GetApplicationPolicySet(ctx context.Context, request *GetApplicationPolicySetRequest) (*GetApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetApplicationPolicySet(ctx, request)
}
func (service *BaseService) ListApplicationPolicySet(ctx context.Context, request *ListApplicationPolicySetRequest) (*ListApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListApplicationPolicySet(ctx, request)
}

func (service *BaseService) CreateBGPAsAService(ctx context.Context, request *CreateBGPAsAServiceRequest) (*CreateBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBGPAsAService(ctx, request)
}
func (service *BaseService) UpdateBGPAsAService(ctx context.Context, request *UpdateBGPAsAServiceRequest) (*UpdateBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPAsAService(ctx, request)
}
func (service *BaseService) DeleteBGPAsAService(ctx context.Context, request *DeleteBGPAsAServiceRequest) (*DeleteBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPAsAService(ctx, request)
}
func (service *BaseService) GetBGPAsAService(ctx context.Context, request *GetBGPAsAServiceRequest) (*GetBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBGPAsAService(ctx, request)
}
func (service *BaseService) ListBGPAsAService(ctx context.Context, request *ListBGPAsAServiceRequest) (*ListBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBGPAsAService(ctx, request)
}

func (service *BaseService) CreateBGPRouter(ctx context.Context, request *CreateBGPRouterRequest) (*CreateBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBGPRouter(ctx, request)
}
func (service *BaseService) UpdateBGPRouter(ctx context.Context, request *UpdateBGPRouterRequest) (*UpdateBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPRouter(ctx, request)
}
func (service *BaseService) DeleteBGPRouter(ctx context.Context, request *DeleteBGPRouterRequest) (*DeleteBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPRouter(ctx, request)
}
func (service *BaseService) GetBGPRouter(ctx context.Context, request *GetBGPRouterRequest) (*GetBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBGPRouter(ctx, request)
}
func (service *BaseService) ListBGPRouter(ctx context.Context, request *ListBGPRouterRequest) (*ListBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBGPRouter(ctx, request)
}

func (service *BaseService) CreateBGPVPN(ctx context.Context, request *CreateBGPVPNRequest) (*CreateBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBGPVPN(ctx, request)
}
func (service *BaseService) UpdateBGPVPN(ctx context.Context, request *UpdateBGPVPNRequest) (*UpdateBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPVPN(ctx, request)
}
func (service *BaseService) DeleteBGPVPN(ctx context.Context, request *DeleteBGPVPNRequest) (*DeleteBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPVPN(ctx, request)
}
func (service *BaseService) GetBGPVPN(ctx context.Context, request *GetBGPVPNRequest) (*GetBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBGPVPN(ctx, request)
}
func (service *BaseService) ListBGPVPN(ctx context.Context, request *ListBGPVPNRequest) (*ListBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBGPVPN(ctx, request)
}

func (service *BaseService) CreateBridgeDomain(ctx context.Context, request *CreateBridgeDomainRequest) (*CreateBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBridgeDomain(ctx, request)
}
func (service *BaseService) UpdateBridgeDomain(ctx context.Context, request *UpdateBridgeDomainRequest) (*UpdateBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBridgeDomain(ctx, request)
}
func (service *BaseService) DeleteBridgeDomain(ctx context.Context, request *DeleteBridgeDomainRequest) (*DeleteBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBridgeDomain(ctx, request)
}
func (service *BaseService) GetBridgeDomain(ctx context.Context, request *GetBridgeDomainRequest) (*GetBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBridgeDomain(ctx, request)
}
func (service *BaseService) ListBridgeDomain(ctx context.Context, request *ListBridgeDomainRequest) (*ListBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBridgeDomain(ctx, request)
}

func (service *BaseService) CreateConfigNode(ctx context.Context, request *CreateConfigNodeRequest) (*CreateConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateConfigNode(ctx, request)
}
func (service *BaseService) UpdateConfigNode(ctx context.Context, request *UpdateConfigNodeRequest) (*UpdateConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateConfigNode(ctx, request)
}
func (service *BaseService) DeleteConfigNode(ctx context.Context, request *DeleteConfigNodeRequest) (*DeleteConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteConfigNode(ctx, request)
}
func (service *BaseService) GetConfigNode(ctx context.Context, request *GetConfigNodeRequest) (*GetConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetConfigNode(ctx, request)
}
func (service *BaseService) ListConfigNode(ctx context.Context, request *ListConfigNodeRequest) (*ListConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListConfigNode(ctx, request)
}

func (service *BaseService) CreateConfigRoot(ctx context.Context, request *CreateConfigRootRequest) (*CreateConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateConfigRoot(ctx, request)
}
func (service *BaseService) UpdateConfigRoot(ctx context.Context, request *UpdateConfigRootRequest) (*UpdateConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateConfigRoot(ctx, request)
}
func (service *BaseService) DeleteConfigRoot(ctx context.Context, request *DeleteConfigRootRequest) (*DeleteConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteConfigRoot(ctx, request)
}
func (service *BaseService) GetConfigRoot(ctx context.Context, request *GetConfigRootRequest) (*GetConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetConfigRoot(ctx, request)
}
func (service *BaseService) ListConfigRoot(ctx context.Context, request *ListConfigRootRequest) (*ListConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListConfigRoot(ctx, request)
}

func (service *BaseService) CreateCustomerAttachment(ctx context.Context, request *CreateCustomerAttachmentRequest) (*CreateCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateCustomerAttachment(ctx, request)
}
func (service *BaseService) UpdateCustomerAttachment(ctx context.Context, request *UpdateCustomerAttachmentRequest) (*UpdateCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateCustomerAttachment(ctx, request)
}
func (service *BaseService) DeleteCustomerAttachment(ctx context.Context, request *DeleteCustomerAttachmentRequest) (*DeleteCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteCustomerAttachment(ctx, request)
}
func (service *BaseService) GetCustomerAttachment(ctx context.Context, request *GetCustomerAttachmentRequest) (*GetCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetCustomerAttachment(ctx, request)
}
func (service *BaseService) ListCustomerAttachment(ctx context.Context, request *ListCustomerAttachmentRequest) (*ListCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListCustomerAttachment(ctx, request)
}

func (service *BaseService) CreateDatabaseNode(ctx context.Context, request *CreateDatabaseNodeRequest) (*CreateDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDatabaseNode(ctx, request)
}
func (service *BaseService) UpdateDatabaseNode(ctx context.Context, request *UpdateDatabaseNodeRequest) (*UpdateDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDatabaseNode(ctx, request)
}
func (service *BaseService) DeleteDatabaseNode(ctx context.Context, request *DeleteDatabaseNodeRequest) (*DeleteDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDatabaseNode(ctx, request)
}
func (service *BaseService) GetDatabaseNode(ctx context.Context, request *GetDatabaseNodeRequest) (*GetDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDatabaseNode(ctx, request)
}
func (service *BaseService) ListDatabaseNode(ctx context.Context, request *ListDatabaseNodeRequest) (*ListDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDatabaseNode(ctx, request)
}

func (service *BaseService) CreateDiscoveryServiceAssignment(ctx context.Context, request *CreateDiscoveryServiceAssignmentRequest) (*CreateDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) UpdateDiscoveryServiceAssignment(ctx context.Context, request *UpdateDiscoveryServiceAssignmentRequest) (*UpdateDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) DeleteDiscoveryServiceAssignment(ctx context.Context, request *DeleteDiscoveryServiceAssignmentRequest) (*DeleteDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) GetDiscoveryServiceAssignment(ctx context.Context, request *GetDiscoveryServiceAssignmentRequest) (*GetDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) ListDiscoveryServiceAssignment(ctx context.Context, request *ListDiscoveryServiceAssignmentRequest) (*ListDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDiscoveryServiceAssignment(ctx, request)
}

func (service *BaseService) CreateDomain(ctx context.Context, request *CreateDomainRequest) (*CreateDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDomain(ctx, request)
}
func (service *BaseService) UpdateDomain(ctx context.Context, request *UpdateDomainRequest) (*UpdateDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDomain(ctx, request)
}
func (service *BaseService) DeleteDomain(ctx context.Context, request *DeleteDomainRequest) (*DeleteDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDomain(ctx, request)
}
func (service *BaseService) GetDomain(ctx context.Context, request *GetDomainRequest) (*GetDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDomain(ctx, request)
}
func (service *BaseService) ListDomain(ctx context.Context, request *ListDomainRequest) (*ListDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDomain(ctx, request)
}

func (service *BaseService) CreateDsaRule(ctx context.Context, request *CreateDsaRuleRequest) (*CreateDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDsaRule(ctx, request)
}
func (service *BaseService) UpdateDsaRule(ctx context.Context, request *UpdateDsaRuleRequest) (*UpdateDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDsaRule(ctx, request)
}
func (service *BaseService) DeleteDsaRule(ctx context.Context, request *DeleteDsaRuleRequest) (*DeleteDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDsaRule(ctx, request)
}
func (service *BaseService) GetDsaRule(ctx context.Context, request *GetDsaRuleRequest) (*GetDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDsaRule(ctx, request)
}
func (service *BaseService) ListDsaRule(ctx context.Context, request *ListDsaRuleRequest) (*ListDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDsaRule(ctx, request)
}

func (service *BaseService) CreateE2ServiceProvider(ctx context.Context, request *CreateE2ServiceProviderRequest) (*CreateE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateE2ServiceProvider(ctx, request)
}
func (service *BaseService) UpdateE2ServiceProvider(ctx context.Context, request *UpdateE2ServiceProviderRequest) (*UpdateE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateE2ServiceProvider(ctx, request)
}
func (service *BaseService) DeleteE2ServiceProvider(ctx context.Context, request *DeleteE2ServiceProviderRequest) (*DeleteE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteE2ServiceProvider(ctx, request)
}
func (service *BaseService) GetE2ServiceProvider(ctx context.Context, request *GetE2ServiceProviderRequest) (*GetE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetE2ServiceProvider(ctx, request)
}
func (service *BaseService) ListE2ServiceProvider(ctx context.Context, request *ListE2ServiceProviderRequest) (*ListE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListE2ServiceProvider(ctx, request)
}

func (service *BaseService) CreateFirewallPolicy(ctx context.Context, request *CreateFirewallPolicyRequest) (*CreateFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFirewallPolicy(ctx, request)
}
func (service *BaseService) UpdateFirewallPolicy(ctx context.Context, request *UpdateFirewallPolicyRequest) (*UpdateFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFirewallPolicy(ctx, request)
}
func (service *BaseService) DeleteFirewallPolicy(ctx context.Context, request *DeleteFirewallPolicyRequest) (*DeleteFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFirewallPolicy(ctx, request)
}
func (service *BaseService) GetFirewallPolicy(ctx context.Context, request *GetFirewallPolicyRequest) (*GetFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFirewallPolicy(ctx, request)
}
func (service *BaseService) ListFirewallPolicy(ctx context.Context, request *ListFirewallPolicyRequest) (*ListFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFirewallPolicy(ctx, request)
}

func (service *BaseService) CreateFirewallRule(ctx context.Context, request *CreateFirewallRuleRequest) (*CreateFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFirewallRule(ctx, request)
}
func (service *BaseService) UpdateFirewallRule(ctx context.Context, request *UpdateFirewallRuleRequest) (*UpdateFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFirewallRule(ctx, request)
}
func (service *BaseService) DeleteFirewallRule(ctx context.Context, request *DeleteFirewallRuleRequest) (*DeleteFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFirewallRule(ctx, request)
}
func (service *BaseService) GetFirewallRule(ctx context.Context, request *GetFirewallRuleRequest) (*GetFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFirewallRule(ctx, request)
}
func (service *BaseService) ListFirewallRule(ctx context.Context, request *ListFirewallRuleRequest) (*ListFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFirewallRule(ctx, request)
}

func (service *BaseService) CreateFloatingIPPool(ctx context.Context, request *CreateFloatingIPPoolRequest) (*CreateFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFloatingIPPool(ctx, request)
}
func (service *BaseService) UpdateFloatingIPPool(ctx context.Context, request *UpdateFloatingIPPoolRequest) (*UpdateFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFloatingIPPool(ctx, request)
}
func (service *BaseService) DeleteFloatingIPPool(ctx context.Context, request *DeleteFloatingIPPoolRequest) (*DeleteFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFloatingIPPool(ctx, request)
}
func (service *BaseService) GetFloatingIPPool(ctx context.Context, request *GetFloatingIPPoolRequest) (*GetFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFloatingIPPool(ctx, request)
}
func (service *BaseService) ListFloatingIPPool(ctx context.Context, request *ListFloatingIPPoolRequest) (*ListFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFloatingIPPool(ctx, request)
}

func (service *BaseService) CreateFloatingIP(ctx context.Context, request *CreateFloatingIPRequest) (*CreateFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFloatingIP(ctx, request)
}
func (service *BaseService) UpdateFloatingIP(ctx context.Context, request *UpdateFloatingIPRequest) (*UpdateFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFloatingIP(ctx, request)
}
func (service *BaseService) DeleteFloatingIP(ctx context.Context, request *DeleteFloatingIPRequest) (*DeleteFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFloatingIP(ctx, request)
}
func (service *BaseService) GetFloatingIP(ctx context.Context, request *GetFloatingIPRequest) (*GetFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFloatingIP(ctx, request)
}
func (service *BaseService) ListFloatingIP(ctx context.Context, request *ListFloatingIPRequest) (*ListFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFloatingIP(ctx, request)
}

func (service *BaseService) CreateForwardingClass(ctx context.Context, request *CreateForwardingClassRequest) (*CreateForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateForwardingClass(ctx, request)
}
func (service *BaseService) UpdateForwardingClass(ctx context.Context, request *UpdateForwardingClassRequest) (*UpdateForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateForwardingClass(ctx, request)
}
func (service *BaseService) DeleteForwardingClass(ctx context.Context, request *DeleteForwardingClassRequest) (*DeleteForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteForwardingClass(ctx, request)
}
func (service *BaseService) GetForwardingClass(ctx context.Context, request *GetForwardingClassRequest) (*GetForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetForwardingClass(ctx, request)
}
func (service *BaseService) ListForwardingClass(ctx context.Context, request *ListForwardingClassRequest) (*ListForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListForwardingClass(ctx, request)
}

func (service *BaseService) CreateGlobalQosConfig(ctx context.Context, request *CreateGlobalQosConfigRequest) (*CreateGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalQosConfig(ctx, request)
}
func (service *BaseService) UpdateGlobalQosConfig(ctx context.Context, request *UpdateGlobalQosConfigRequest) (*UpdateGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalQosConfig(ctx, request)
}
func (service *BaseService) DeleteGlobalQosConfig(ctx context.Context, request *DeleteGlobalQosConfigRequest) (*DeleteGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalQosConfig(ctx, request)
}
func (service *BaseService) GetGlobalQosConfig(ctx context.Context, request *GetGlobalQosConfigRequest) (*GetGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetGlobalQosConfig(ctx, request)
}
func (service *BaseService) ListGlobalQosConfig(ctx context.Context, request *ListGlobalQosConfigRequest) (*ListGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListGlobalQosConfig(ctx, request)
}

func (service *BaseService) CreateGlobalSystemConfig(ctx context.Context, request *CreateGlobalSystemConfigRequest) (*CreateGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalSystemConfig(ctx, request)
}
func (service *BaseService) UpdateGlobalSystemConfig(ctx context.Context, request *UpdateGlobalSystemConfigRequest) (*UpdateGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalSystemConfig(ctx, request)
}
func (service *BaseService) DeleteGlobalSystemConfig(ctx context.Context, request *DeleteGlobalSystemConfigRequest) (*DeleteGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalSystemConfig(ctx, request)
}
func (service *BaseService) GetGlobalSystemConfig(ctx context.Context, request *GetGlobalSystemConfigRequest) (*GetGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetGlobalSystemConfig(ctx, request)
}
func (service *BaseService) ListGlobalSystemConfig(ctx context.Context, request *ListGlobalSystemConfigRequest) (*ListGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListGlobalSystemConfig(ctx, request)
}

func (service *BaseService) CreateGlobalVrouterConfig(ctx context.Context, request *CreateGlobalVrouterConfigRequest) (*CreateGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) UpdateGlobalVrouterConfig(ctx context.Context, request *UpdateGlobalVrouterConfigRequest) (*UpdateGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) DeleteGlobalVrouterConfig(ctx context.Context, request *DeleteGlobalVrouterConfigRequest) (*DeleteGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) GetGlobalVrouterConfig(ctx context.Context, request *GetGlobalVrouterConfigRequest) (*GetGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) ListGlobalVrouterConfig(ctx context.Context, request *ListGlobalVrouterConfigRequest) (*ListGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListGlobalVrouterConfig(ctx, request)
}

func (service *BaseService) CreateInstanceIP(ctx context.Context, request *CreateInstanceIPRequest) (*CreateInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateInstanceIP(ctx, request)
}
func (service *BaseService) UpdateInstanceIP(ctx context.Context, request *UpdateInstanceIPRequest) (*UpdateInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateInstanceIP(ctx, request)
}
func (service *BaseService) DeleteInstanceIP(ctx context.Context, request *DeleteInstanceIPRequest) (*DeleteInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteInstanceIP(ctx, request)
}
func (service *BaseService) GetInstanceIP(ctx context.Context, request *GetInstanceIPRequest) (*GetInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetInstanceIP(ctx, request)
}
func (service *BaseService) ListInstanceIP(ctx context.Context, request *ListInstanceIPRequest) (*ListInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListInstanceIP(ctx, request)
}

func (service *BaseService) CreateInterfaceRouteTable(ctx context.Context, request *CreateInterfaceRouteTableRequest) (*CreateInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateInterfaceRouteTable(ctx, request)
}
func (service *BaseService) UpdateInterfaceRouteTable(ctx context.Context, request *UpdateInterfaceRouteTableRequest) (*UpdateInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateInterfaceRouteTable(ctx, request)
}
func (service *BaseService) DeleteInterfaceRouteTable(ctx context.Context, request *DeleteInterfaceRouteTableRequest) (*DeleteInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteInterfaceRouteTable(ctx, request)
}
func (service *BaseService) GetInterfaceRouteTable(ctx context.Context, request *GetInterfaceRouteTableRequest) (*GetInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetInterfaceRouteTable(ctx, request)
}
func (service *BaseService) ListInterfaceRouteTable(ctx context.Context, request *ListInterfaceRouteTableRequest) (*ListInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListInterfaceRouteTable(ctx, request)
}

func (service *BaseService) CreateLoadbalancerHealthmonitor(ctx context.Context, request *CreateLoadbalancerHealthmonitorRequest) (*CreateLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerHealthmonitor(ctx context.Context, request *UpdateLoadbalancerHealthmonitorRequest) (*UpdateLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerHealthmonitor(ctx context.Context, request *DeleteLoadbalancerHealthmonitorRequest) (*DeleteLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) GetLoadbalancerHealthmonitor(ctx context.Context, request *GetLoadbalancerHealthmonitorRequest) (*GetLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) ListLoadbalancerHealthmonitor(ctx context.Context, request *ListLoadbalancerHealthmonitorRequest) (*ListLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerHealthmonitor(ctx, request)
}

func (service *BaseService) CreateLoadbalancerListener(ctx context.Context, request *CreateLoadbalancerListenerRequest) (*CreateLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerListener(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerListener(ctx context.Context, request *UpdateLoadbalancerListenerRequest) (*UpdateLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerListener(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerListener(ctx context.Context, request *DeleteLoadbalancerListenerRequest) (*DeleteLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerListener(ctx, request)
}
func (service *BaseService) GetLoadbalancerListener(ctx context.Context, request *GetLoadbalancerListenerRequest) (*GetLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerListener(ctx, request)
}
func (service *BaseService) ListLoadbalancerListener(ctx context.Context, request *ListLoadbalancerListenerRequest) (*ListLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerListener(ctx, request)
}

func (service *BaseService) CreateLoadbalancerMember(ctx context.Context, request *CreateLoadbalancerMemberRequest) (*CreateLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerMember(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerMember(ctx context.Context, request *UpdateLoadbalancerMemberRequest) (*UpdateLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerMember(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerMember(ctx context.Context, request *DeleteLoadbalancerMemberRequest) (*DeleteLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerMember(ctx, request)
}
func (service *BaseService) GetLoadbalancerMember(ctx context.Context, request *GetLoadbalancerMemberRequest) (*GetLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerMember(ctx, request)
}
func (service *BaseService) ListLoadbalancerMember(ctx context.Context, request *ListLoadbalancerMemberRequest) (*ListLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerMember(ctx, request)
}

func (service *BaseService) CreateLoadbalancerPool(ctx context.Context, request *CreateLoadbalancerPoolRequest) (*CreateLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerPool(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerPool(ctx context.Context, request *UpdateLoadbalancerPoolRequest) (*UpdateLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerPool(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerPool(ctx context.Context, request *DeleteLoadbalancerPoolRequest) (*DeleteLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerPool(ctx, request)
}
func (service *BaseService) GetLoadbalancerPool(ctx context.Context, request *GetLoadbalancerPoolRequest) (*GetLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerPool(ctx, request)
}
func (service *BaseService) ListLoadbalancerPool(ctx context.Context, request *ListLoadbalancerPoolRequest) (*ListLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerPool(ctx, request)
}

func (service *BaseService) CreateLoadbalancer(ctx context.Context, request *CreateLoadbalancerRequest) (*CreateLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancer(ctx, request)
}
func (service *BaseService) UpdateLoadbalancer(ctx context.Context, request *UpdateLoadbalancerRequest) (*UpdateLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancer(ctx, request)
}
func (service *BaseService) DeleteLoadbalancer(ctx context.Context, request *DeleteLoadbalancerRequest) (*DeleteLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancer(ctx, request)
}
func (service *BaseService) GetLoadbalancer(ctx context.Context, request *GetLoadbalancerRequest) (*GetLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancer(ctx, request)
}
func (service *BaseService) ListLoadbalancer(ctx context.Context, request *ListLoadbalancerRequest) (*ListLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancer(ctx, request)
}

func (service *BaseService) CreateLogicalInterface(ctx context.Context, request *CreateLogicalInterfaceRequest) (*CreateLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLogicalInterface(ctx, request)
}
func (service *BaseService) UpdateLogicalInterface(ctx context.Context, request *UpdateLogicalInterfaceRequest) (*UpdateLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLogicalInterface(ctx, request)
}
func (service *BaseService) DeleteLogicalInterface(ctx context.Context, request *DeleteLogicalInterfaceRequest) (*DeleteLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLogicalInterface(ctx, request)
}
func (service *BaseService) GetLogicalInterface(ctx context.Context, request *GetLogicalInterfaceRequest) (*GetLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLogicalInterface(ctx, request)
}
func (service *BaseService) ListLogicalInterface(ctx context.Context, request *ListLogicalInterfaceRequest) (*ListLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLogicalInterface(ctx, request)
}

func (service *BaseService) CreateLogicalRouter(ctx context.Context, request *CreateLogicalRouterRequest) (*CreateLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLogicalRouter(ctx, request)
}
func (service *BaseService) UpdateLogicalRouter(ctx context.Context, request *UpdateLogicalRouterRequest) (*UpdateLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLogicalRouter(ctx, request)
}
func (service *BaseService) DeleteLogicalRouter(ctx context.Context, request *DeleteLogicalRouterRequest) (*DeleteLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLogicalRouter(ctx, request)
}
func (service *BaseService) GetLogicalRouter(ctx context.Context, request *GetLogicalRouterRequest) (*GetLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLogicalRouter(ctx, request)
}
func (service *BaseService) ListLogicalRouter(ctx context.Context, request *ListLogicalRouterRequest) (*ListLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLogicalRouter(ctx, request)
}

func (service *BaseService) CreateNamespace(ctx context.Context, request *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNamespace(ctx, request)
}
func (service *BaseService) UpdateNamespace(ctx context.Context, request *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNamespace(ctx, request)
}
func (service *BaseService) DeleteNamespace(ctx context.Context, request *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNamespace(ctx, request)
}
func (service *BaseService) GetNamespace(ctx context.Context, request *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNamespace(ctx, request)
}
func (service *BaseService) ListNamespace(ctx context.Context, request *ListNamespaceRequest) (*ListNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNamespace(ctx, request)
}

func (service *BaseService) CreateNetworkDeviceConfig(ctx context.Context, request *CreateNetworkDeviceConfigRequest) (*CreateNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) UpdateNetworkDeviceConfig(ctx context.Context, request *UpdateNetworkDeviceConfigRequest) (*UpdateNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) DeleteNetworkDeviceConfig(ctx context.Context, request *DeleteNetworkDeviceConfigRequest) (*DeleteNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) GetNetworkDeviceConfig(ctx context.Context, request *GetNetworkDeviceConfigRequest) (*GetNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) ListNetworkDeviceConfig(ctx context.Context, request *ListNetworkDeviceConfigRequest) (*ListNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNetworkDeviceConfig(ctx, request)
}

func (service *BaseService) CreateNetworkIpam(ctx context.Context, request *CreateNetworkIpamRequest) (*CreateNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkIpam(ctx, request)
}
func (service *BaseService) UpdateNetworkIpam(ctx context.Context, request *UpdateNetworkIpamRequest) (*UpdateNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkIpam(ctx, request)
}
func (service *BaseService) DeleteNetworkIpam(ctx context.Context, request *DeleteNetworkIpamRequest) (*DeleteNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkIpam(ctx, request)
}
func (service *BaseService) GetNetworkIpam(ctx context.Context, request *GetNetworkIpamRequest) (*GetNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNetworkIpam(ctx, request)
}
func (service *BaseService) ListNetworkIpam(ctx context.Context, request *ListNetworkIpamRequest) (*ListNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNetworkIpam(ctx, request)
}

func (service *BaseService) CreateNetworkPolicy(ctx context.Context, request *CreateNetworkPolicyRequest) (*CreateNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkPolicy(ctx, request)
}
func (service *BaseService) UpdateNetworkPolicy(ctx context.Context, request *UpdateNetworkPolicyRequest) (*UpdateNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkPolicy(ctx, request)
}
func (service *BaseService) DeleteNetworkPolicy(ctx context.Context, request *DeleteNetworkPolicyRequest) (*DeleteNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkPolicy(ctx, request)
}
func (service *BaseService) GetNetworkPolicy(ctx context.Context, request *GetNetworkPolicyRequest) (*GetNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNetworkPolicy(ctx, request)
}
func (service *BaseService) ListNetworkPolicy(ctx context.Context, request *ListNetworkPolicyRequest) (*ListNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNetworkPolicy(ctx, request)
}

func (service *BaseService) CreatePeeringPolicy(ctx context.Context, request *CreatePeeringPolicyRequest) (*CreatePeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePeeringPolicy(ctx, request)
}
func (service *BaseService) UpdatePeeringPolicy(ctx context.Context, request *UpdatePeeringPolicyRequest) (*UpdatePeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePeeringPolicy(ctx, request)
}
func (service *BaseService) DeletePeeringPolicy(ctx context.Context, request *DeletePeeringPolicyRequest) (*DeletePeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePeeringPolicy(ctx, request)
}
func (service *BaseService) GetPeeringPolicy(ctx context.Context, request *GetPeeringPolicyRequest) (*GetPeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPeeringPolicy(ctx, request)
}
func (service *BaseService) ListPeeringPolicy(ctx context.Context, request *ListPeeringPolicyRequest) (*ListPeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPeeringPolicy(ctx, request)
}

func (service *BaseService) CreatePhysicalInterface(ctx context.Context, request *CreatePhysicalInterfaceRequest) (*CreatePhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePhysicalInterface(ctx, request)
}
func (service *BaseService) UpdatePhysicalInterface(ctx context.Context, request *UpdatePhysicalInterfaceRequest) (*UpdatePhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePhysicalInterface(ctx, request)
}
func (service *BaseService) DeletePhysicalInterface(ctx context.Context, request *DeletePhysicalInterfaceRequest) (*DeletePhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePhysicalInterface(ctx, request)
}
func (service *BaseService) GetPhysicalInterface(ctx context.Context, request *GetPhysicalInterfaceRequest) (*GetPhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPhysicalInterface(ctx, request)
}
func (service *BaseService) ListPhysicalInterface(ctx context.Context, request *ListPhysicalInterfaceRequest) (*ListPhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPhysicalInterface(ctx, request)
}

func (service *BaseService) CreatePhysicalRouter(ctx context.Context, request *CreatePhysicalRouterRequest) (*CreatePhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePhysicalRouter(ctx, request)
}
func (service *BaseService) UpdatePhysicalRouter(ctx context.Context, request *UpdatePhysicalRouterRequest) (*UpdatePhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePhysicalRouter(ctx, request)
}
func (service *BaseService) DeletePhysicalRouter(ctx context.Context, request *DeletePhysicalRouterRequest) (*DeletePhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePhysicalRouter(ctx, request)
}
func (service *BaseService) GetPhysicalRouter(ctx context.Context, request *GetPhysicalRouterRequest) (*GetPhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPhysicalRouter(ctx, request)
}
func (service *BaseService) ListPhysicalRouter(ctx context.Context, request *ListPhysicalRouterRequest) (*ListPhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPhysicalRouter(ctx, request)
}

func (service *BaseService) CreatePolicyManagement(ctx context.Context, request *CreatePolicyManagementRequest) (*CreatePolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePolicyManagement(ctx, request)
}
func (service *BaseService) UpdatePolicyManagement(ctx context.Context, request *UpdatePolicyManagementRequest) (*UpdatePolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePolicyManagement(ctx, request)
}
func (service *BaseService) DeletePolicyManagement(ctx context.Context, request *DeletePolicyManagementRequest) (*DeletePolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePolicyManagement(ctx, request)
}
func (service *BaseService) GetPolicyManagement(ctx context.Context, request *GetPolicyManagementRequest) (*GetPolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPolicyManagement(ctx, request)
}
func (service *BaseService) ListPolicyManagement(ctx context.Context, request *ListPolicyManagementRequest) (*ListPolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPolicyManagement(ctx, request)
}

func (service *BaseService) CreatePortTuple(ctx context.Context, request *CreatePortTupleRequest) (*CreatePortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePortTuple(ctx, request)
}
func (service *BaseService) UpdatePortTuple(ctx context.Context, request *UpdatePortTupleRequest) (*UpdatePortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePortTuple(ctx, request)
}
func (service *BaseService) DeletePortTuple(ctx context.Context, request *DeletePortTupleRequest) (*DeletePortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePortTuple(ctx, request)
}
func (service *BaseService) GetPortTuple(ctx context.Context, request *GetPortTupleRequest) (*GetPortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPortTuple(ctx, request)
}
func (service *BaseService) ListPortTuple(ctx context.Context, request *ListPortTupleRequest) (*ListPortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPortTuple(ctx, request)
}

func (service *BaseService) CreateProject(ctx context.Context, request *CreateProjectRequest) (*CreateProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateProject(ctx, request)
}
func (service *BaseService) UpdateProject(ctx context.Context, request *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateProject(ctx, request)
}
func (service *BaseService) DeleteProject(ctx context.Context, request *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteProject(ctx, request)
}
func (service *BaseService) GetProject(ctx context.Context, request *GetProjectRequest) (*GetProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetProject(ctx, request)
}
func (service *BaseService) ListProject(ctx context.Context, request *ListProjectRequest) (*ListProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListProject(ctx, request)
}

func (service *BaseService) CreateProviderAttachment(ctx context.Context, request *CreateProviderAttachmentRequest) (*CreateProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateProviderAttachment(ctx, request)
}
func (service *BaseService) UpdateProviderAttachment(ctx context.Context, request *UpdateProviderAttachmentRequest) (*UpdateProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateProviderAttachment(ctx, request)
}
func (service *BaseService) DeleteProviderAttachment(ctx context.Context, request *DeleteProviderAttachmentRequest) (*DeleteProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteProviderAttachment(ctx, request)
}
func (service *BaseService) GetProviderAttachment(ctx context.Context, request *GetProviderAttachmentRequest) (*GetProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetProviderAttachment(ctx, request)
}
func (service *BaseService) ListProviderAttachment(ctx context.Context, request *ListProviderAttachmentRequest) (*ListProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListProviderAttachment(ctx, request)
}

func (service *BaseService) CreateQosConfig(ctx context.Context, request *CreateQosConfigRequest) (*CreateQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateQosConfig(ctx, request)
}
func (service *BaseService) UpdateQosConfig(ctx context.Context, request *UpdateQosConfigRequest) (*UpdateQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateQosConfig(ctx, request)
}
func (service *BaseService) DeleteQosConfig(ctx context.Context, request *DeleteQosConfigRequest) (*DeleteQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteQosConfig(ctx, request)
}
func (service *BaseService) GetQosConfig(ctx context.Context, request *GetQosConfigRequest) (*GetQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetQosConfig(ctx, request)
}
func (service *BaseService) ListQosConfig(ctx context.Context, request *ListQosConfigRequest) (*ListQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListQosConfig(ctx, request)
}

func (service *BaseService) CreateQosQueue(ctx context.Context, request *CreateQosQueueRequest) (*CreateQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateQosQueue(ctx, request)
}
func (service *BaseService) UpdateQosQueue(ctx context.Context, request *UpdateQosQueueRequest) (*UpdateQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateQosQueue(ctx, request)
}
func (service *BaseService) DeleteQosQueue(ctx context.Context, request *DeleteQosQueueRequest) (*DeleteQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteQosQueue(ctx, request)
}
func (service *BaseService) GetQosQueue(ctx context.Context, request *GetQosQueueRequest) (*GetQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetQosQueue(ctx, request)
}
func (service *BaseService) ListQosQueue(ctx context.Context, request *ListQosQueueRequest) (*ListQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListQosQueue(ctx, request)
}

func (service *BaseService) CreateRouteAggregate(ctx context.Context, request *CreateRouteAggregateRequest) (*CreateRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRouteAggregate(ctx, request)
}
func (service *BaseService) UpdateRouteAggregate(ctx context.Context, request *UpdateRouteAggregateRequest) (*UpdateRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteAggregate(ctx, request)
}
func (service *BaseService) DeleteRouteAggregate(ctx context.Context, request *DeleteRouteAggregateRequest) (*DeleteRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteAggregate(ctx, request)
}
func (service *BaseService) GetRouteAggregate(ctx context.Context, request *GetRouteAggregateRequest) (*GetRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRouteAggregate(ctx, request)
}
func (service *BaseService) ListRouteAggregate(ctx context.Context, request *ListRouteAggregateRequest) (*ListRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRouteAggregate(ctx, request)
}

func (service *BaseService) CreateRouteTable(ctx context.Context, request *CreateRouteTableRequest) (*CreateRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRouteTable(ctx, request)
}
func (service *BaseService) UpdateRouteTable(ctx context.Context, request *UpdateRouteTableRequest) (*UpdateRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteTable(ctx, request)
}
func (service *BaseService) DeleteRouteTable(ctx context.Context, request *DeleteRouteTableRequest) (*DeleteRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteTable(ctx, request)
}
func (service *BaseService) GetRouteTable(ctx context.Context, request *GetRouteTableRequest) (*GetRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRouteTable(ctx, request)
}
func (service *BaseService) ListRouteTable(ctx context.Context, request *ListRouteTableRequest) (*ListRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRouteTable(ctx, request)
}

func (service *BaseService) CreateRouteTarget(ctx context.Context, request *CreateRouteTargetRequest) (*CreateRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRouteTarget(ctx, request)
}
func (service *BaseService) UpdateRouteTarget(ctx context.Context, request *UpdateRouteTargetRequest) (*UpdateRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteTarget(ctx, request)
}
func (service *BaseService) DeleteRouteTarget(ctx context.Context, request *DeleteRouteTargetRequest) (*DeleteRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteTarget(ctx, request)
}
func (service *BaseService) GetRouteTarget(ctx context.Context, request *GetRouteTargetRequest) (*GetRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRouteTarget(ctx, request)
}
func (service *BaseService) ListRouteTarget(ctx context.Context, request *ListRouteTargetRequest) (*ListRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRouteTarget(ctx, request)
}

func (service *BaseService) CreateRoutingInstance(ctx context.Context, request *CreateRoutingInstanceRequest) (*CreateRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRoutingInstance(ctx, request)
}
func (service *BaseService) UpdateRoutingInstance(ctx context.Context, request *UpdateRoutingInstanceRequest) (*UpdateRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRoutingInstance(ctx, request)
}
func (service *BaseService) DeleteRoutingInstance(ctx context.Context, request *DeleteRoutingInstanceRequest) (*DeleteRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRoutingInstance(ctx, request)
}
func (service *BaseService) GetRoutingInstance(ctx context.Context, request *GetRoutingInstanceRequest) (*GetRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRoutingInstance(ctx, request)
}
func (service *BaseService) ListRoutingInstance(ctx context.Context, request *ListRoutingInstanceRequest) (*ListRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRoutingInstance(ctx, request)
}

func (service *BaseService) CreateRoutingPolicy(ctx context.Context, request *CreateRoutingPolicyRequest) (*CreateRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRoutingPolicy(ctx, request)
}
func (service *BaseService) UpdateRoutingPolicy(ctx context.Context, request *UpdateRoutingPolicyRequest) (*UpdateRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRoutingPolicy(ctx, request)
}
func (service *BaseService) DeleteRoutingPolicy(ctx context.Context, request *DeleteRoutingPolicyRequest) (*DeleteRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRoutingPolicy(ctx, request)
}
func (service *BaseService) GetRoutingPolicy(ctx context.Context, request *GetRoutingPolicyRequest) (*GetRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRoutingPolicy(ctx, request)
}
func (service *BaseService) ListRoutingPolicy(ctx context.Context, request *ListRoutingPolicyRequest) (*ListRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRoutingPolicy(ctx, request)
}

func (service *BaseService) CreateSecurityGroup(ctx context.Context, request *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateSecurityGroup(ctx, request)
}
func (service *BaseService) UpdateSecurityGroup(ctx context.Context, request *UpdateSecurityGroupRequest) (*UpdateSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateSecurityGroup(ctx, request)
}
func (service *BaseService) DeleteSecurityGroup(ctx context.Context, request *DeleteSecurityGroupRequest) (*DeleteSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteSecurityGroup(ctx, request)
}
func (service *BaseService) GetSecurityGroup(ctx context.Context, request *GetSecurityGroupRequest) (*GetSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetSecurityGroup(ctx, request)
}
func (service *BaseService) ListSecurityGroup(ctx context.Context, request *ListSecurityGroupRequest) (*ListSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListSecurityGroup(ctx, request)
}

func (service *BaseService) CreateSecurityLoggingObject(ctx context.Context, request *CreateSecurityLoggingObjectRequest) (*CreateSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateSecurityLoggingObject(ctx, request)
}
func (service *BaseService) UpdateSecurityLoggingObject(ctx context.Context, request *UpdateSecurityLoggingObjectRequest) (*UpdateSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateSecurityLoggingObject(ctx, request)
}
func (service *BaseService) DeleteSecurityLoggingObject(ctx context.Context, request *DeleteSecurityLoggingObjectRequest) (*DeleteSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteSecurityLoggingObject(ctx, request)
}
func (service *BaseService) GetSecurityLoggingObject(ctx context.Context, request *GetSecurityLoggingObjectRequest) (*GetSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetSecurityLoggingObject(ctx, request)
}
func (service *BaseService) ListSecurityLoggingObject(ctx context.Context, request *ListSecurityLoggingObjectRequest) (*ListSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListSecurityLoggingObject(ctx, request)
}

func (service *BaseService) CreateServiceAppliance(ctx context.Context, request *CreateServiceApplianceRequest) (*CreateServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceAppliance(ctx, request)
}
func (service *BaseService) UpdateServiceAppliance(ctx context.Context, request *UpdateServiceApplianceRequest) (*UpdateServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceAppliance(ctx, request)
}
func (service *BaseService) DeleteServiceAppliance(ctx context.Context, request *DeleteServiceApplianceRequest) (*DeleteServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceAppliance(ctx, request)
}
func (service *BaseService) GetServiceAppliance(ctx context.Context, request *GetServiceApplianceRequest) (*GetServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceAppliance(ctx, request)
}
func (service *BaseService) ListServiceAppliance(ctx context.Context, request *ListServiceApplianceRequest) (*ListServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceAppliance(ctx, request)
}

func (service *BaseService) CreateServiceApplianceSet(ctx context.Context, request *CreateServiceApplianceSetRequest) (*CreateServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceApplianceSet(ctx, request)
}
func (service *BaseService) UpdateServiceApplianceSet(ctx context.Context, request *UpdateServiceApplianceSetRequest) (*UpdateServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceApplianceSet(ctx, request)
}
func (service *BaseService) DeleteServiceApplianceSet(ctx context.Context, request *DeleteServiceApplianceSetRequest) (*DeleteServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceApplianceSet(ctx, request)
}
func (service *BaseService) GetServiceApplianceSet(ctx context.Context, request *GetServiceApplianceSetRequest) (*GetServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceApplianceSet(ctx, request)
}
func (service *BaseService) ListServiceApplianceSet(ctx context.Context, request *ListServiceApplianceSetRequest) (*ListServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceApplianceSet(ctx, request)
}

func (service *BaseService) CreateServiceConnectionModule(ctx context.Context, request *CreateServiceConnectionModuleRequest) (*CreateServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceConnectionModule(ctx, request)
}
func (service *BaseService) UpdateServiceConnectionModule(ctx context.Context, request *UpdateServiceConnectionModuleRequest) (*UpdateServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceConnectionModule(ctx, request)
}
func (service *BaseService) DeleteServiceConnectionModule(ctx context.Context, request *DeleteServiceConnectionModuleRequest) (*DeleteServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceConnectionModule(ctx, request)
}
func (service *BaseService) GetServiceConnectionModule(ctx context.Context, request *GetServiceConnectionModuleRequest) (*GetServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceConnectionModule(ctx, request)
}
func (service *BaseService) ListServiceConnectionModule(ctx context.Context, request *ListServiceConnectionModuleRequest) (*ListServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceConnectionModule(ctx, request)
}

func (service *BaseService) CreateServiceEndpoint(ctx context.Context, request *CreateServiceEndpointRequest) (*CreateServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceEndpoint(ctx, request)
}
func (service *BaseService) UpdateServiceEndpoint(ctx context.Context, request *UpdateServiceEndpointRequest) (*UpdateServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceEndpoint(ctx, request)
}
func (service *BaseService) DeleteServiceEndpoint(ctx context.Context, request *DeleteServiceEndpointRequest) (*DeleteServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceEndpoint(ctx, request)
}
func (service *BaseService) GetServiceEndpoint(ctx context.Context, request *GetServiceEndpointRequest) (*GetServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceEndpoint(ctx, request)
}
func (service *BaseService) ListServiceEndpoint(ctx context.Context, request *ListServiceEndpointRequest) (*ListServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceEndpoint(ctx, request)
}

func (service *BaseService) CreateServiceGroup(ctx context.Context, request *CreateServiceGroupRequest) (*CreateServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceGroup(ctx, request)
}
func (service *BaseService) UpdateServiceGroup(ctx context.Context, request *UpdateServiceGroupRequest) (*UpdateServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceGroup(ctx, request)
}
func (service *BaseService) DeleteServiceGroup(ctx context.Context, request *DeleteServiceGroupRequest) (*DeleteServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceGroup(ctx, request)
}
func (service *BaseService) GetServiceGroup(ctx context.Context, request *GetServiceGroupRequest) (*GetServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceGroup(ctx, request)
}
func (service *BaseService) ListServiceGroup(ctx context.Context, request *ListServiceGroupRequest) (*ListServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceGroup(ctx, request)
}

func (service *BaseService) CreateServiceHealthCheck(ctx context.Context, request *CreateServiceHealthCheckRequest) (*CreateServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceHealthCheck(ctx, request)
}
func (service *BaseService) UpdateServiceHealthCheck(ctx context.Context, request *UpdateServiceHealthCheckRequest) (*UpdateServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceHealthCheck(ctx, request)
}
func (service *BaseService) DeleteServiceHealthCheck(ctx context.Context, request *DeleteServiceHealthCheckRequest) (*DeleteServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceHealthCheck(ctx, request)
}
func (service *BaseService) GetServiceHealthCheck(ctx context.Context, request *GetServiceHealthCheckRequest) (*GetServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceHealthCheck(ctx, request)
}
func (service *BaseService) ListServiceHealthCheck(ctx context.Context, request *ListServiceHealthCheckRequest) (*ListServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceHealthCheck(ctx, request)
}

func (service *BaseService) CreateServiceInstance(ctx context.Context, request *CreateServiceInstanceRequest) (*CreateServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceInstance(ctx, request)
}
func (service *BaseService) UpdateServiceInstance(ctx context.Context, request *UpdateServiceInstanceRequest) (*UpdateServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceInstance(ctx, request)
}
func (service *BaseService) DeleteServiceInstance(ctx context.Context, request *DeleteServiceInstanceRequest) (*DeleteServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceInstance(ctx, request)
}
func (service *BaseService) GetServiceInstance(ctx context.Context, request *GetServiceInstanceRequest) (*GetServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceInstance(ctx, request)
}
func (service *BaseService) ListServiceInstance(ctx context.Context, request *ListServiceInstanceRequest) (*ListServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceInstance(ctx, request)
}

func (service *BaseService) CreateServiceObject(ctx context.Context, request *CreateServiceObjectRequest) (*CreateServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceObject(ctx, request)
}
func (service *BaseService) UpdateServiceObject(ctx context.Context, request *UpdateServiceObjectRequest) (*UpdateServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceObject(ctx, request)
}
func (service *BaseService) DeleteServiceObject(ctx context.Context, request *DeleteServiceObjectRequest) (*DeleteServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceObject(ctx, request)
}
func (service *BaseService) GetServiceObject(ctx context.Context, request *GetServiceObjectRequest) (*GetServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceObject(ctx, request)
}
func (service *BaseService) ListServiceObject(ctx context.Context, request *ListServiceObjectRequest) (*ListServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceObject(ctx, request)
}

func (service *BaseService) CreateServiceTemplate(ctx context.Context, request *CreateServiceTemplateRequest) (*CreateServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceTemplate(ctx, request)
}
func (service *BaseService) UpdateServiceTemplate(ctx context.Context, request *UpdateServiceTemplateRequest) (*UpdateServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceTemplate(ctx, request)
}
func (service *BaseService) DeleteServiceTemplate(ctx context.Context, request *DeleteServiceTemplateRequest) (*DeleteServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceTemplate(ctx, request)
}
func (service *BaseService) GetServiceTemplate(ctx context.Context, request *GetServiceTemplateRequest) (*GetServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceTemplate(ctx, request)
}
func (service *BaseService) ListServiceTemplate(ctx context.Context, request *ListServiceTemplateRequest) (*ListServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceTemplate(ctx, request)
}

func (service *BaseService) CreateSubnet(ctx context.Context, request *CreateSubnetRequest) (*CreateSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateSubnet(ctx, request)
}
func (service *BaseService) UpdateSubnet(ctx context.Context, request *UpdateSubnetRequest) (*UpdateSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateSubnet(ctx, request)
}
func (service *BaseService) DeleteSubnet(ctx context.Context, request *DeleteSubnetRequest) (*DeleteSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteSubnet(ctx, request)
}
func (service *BaseService) GetSubnet(ctx context.Context, request *GetSubnetRequest) (*GetSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetSubnet(ctx, request)
}
func (service *BaseService) ListSubnet(ctx context.Context, request *ListSubnetRequest) (*ListSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListSubnet(ctx, request)
}

func (service *BaseService) CreateTag(ctx context.Context, request *CreateTagRequest) (*CreateTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateTag(ctx, request)
}
func (service *BaseService) UpdateTag(ctx context.Context, request *UpdateTagRequest) (*UpdateTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateTag(ctx, request)
}
func (service *BaseService) DeleteTag(ctx context.Context, request *DeleteTagRequest) (*DeleteTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteTag(ctx, request)
}
func (service *BaseService) GetTag(ctx context.Context, request *GetTagRequest) (*GetTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetTag(ctx, request)
}
func (service *BaseService) ListTag(ctx context.Context, request *ListTagRequest) (*ListTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListTag(ctx, request)
}

func (service *BaseService) CreateTagType(ctx context.Context, request *CreateTagTypeRequest) (*CreateTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateTagType(ctx, request)
}
func (service *BaseService) UpdateTagType(ctx context.Context, request *UpdateTagTypeRequest) (*UpdateTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateTagType(ctx, request)
}
func (service *BaseService) DeleteTagType(ctx context.Context, request *DeleteTagTypeRequest) (*DeleteTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteTagType(ctx, request)
}
func (service *BaseService) GetTagType(ctx context.Context, request *GetTagTypeRequest) (*GetTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetTagType(ctx, request)
}
func (service *BaseService) ListTagType(ctx context.Context, request *ListTagTypeRequest) (*ListTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListTagType(ctx, request)
}

func (service *BaseService) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateUser(ctx, request)
}
func (service *BaseService) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateUser(ctx, request)
}
func (service *BaseService) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteUser(ctx, request)
}
func (service *BaseService) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetUser(ctx, request)
}
func (service *BaseService) ListUser(ctx context.Context, request *ListUserRequest) (*ListUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListUser(ctx, request)
}

func (service *BaseService) CreateVirtualDNSRecord(ctx context.Context, request *CreateVirtualDNSRecordRequest) (*CreateVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualDNSRecord(ctx, request)
}
func (service *BaseService) UpdateVirtualDNSRecord(ctx context.Context, request *UpdateVirtualDNSRecordRequest) (*UpdateVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualDNSRecord(ctx, request)
}
func (service *BaseService) DeleteVirtualDNSRecord(ctx context.Context, request *DeleteVirtualDNSRecordRequest) (*DeleteVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualDNSRecord(ctx, request)
}
func (service *BaseService) GetVirtualDNSRecord(ctx context.Context, request *GetVirtualDNSRecordRequest) (*GetVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualDNSRecord(ctx, request)
}
func (service *BaseService) ListVirtualDNSRecord(ctx context.Context, request *ListVirtualDNSRecordRequest) (*ListVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualDNSRecord(ctx, request)
}

func (service *BaseService) CreateVirtualDNS(ctx context.Context, request *CreateVirtualDNSRequest) (*CreateVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualDNS(ctx, request)
}
func (service *BaseService) UpdateVirtualDNS(ctx context.Context, request *UpdateVirtualDNSRequest) (*UpdateVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualDNS(ctx, request)
}
func (service *BaseService) DeleteVirtualDNS(ctx context.Context, request *DeleteVirtualDNSRequest) (*DeleteVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualDNS(ctx, request)
}
func (service *BaseService) GetVirtualDNS(ctx context.Context, request *GetVirtualDNSRequest) (*GetVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualDNS(ctx, request)
}
func (service *BaseService) ListVirtualDNS(ctx context.Context, request *ListVirtualDNSRequest) (*ListVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualDNS(ctx, request)
}

func (service *BaseService) CreateVirtualIP(ctx context.Context, request *CreateVirtualIPRequest) (*CreateVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualIP(ctx, request)
}
func (service *BaseService) UpdateVirtualIP(ctx context.Context, request *UpdateVirtualIPRequest) (*UpdateVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualIP(ctx, request)
}
func (service *BaseService) DeleteVirtualIP(ctx context.Context, request *DeleteVirtualIPRequest) (*DeleteVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualIP(ctx, request)
}
func (service *BaseService) GetVirtualIP(ctx context.Context, request *GetVirtualIPRequest) (*GetVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualIP(ctx, request)
}
func (service *BaseService) ListVirtualIP(ctx context.Context, request *ListVirtualIPRequest) (*ListVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualIP(ctx, request)
}

func (service *BaseService) CreateVirtualMachineInterface(ctx context.Context, request *CreateVirtualMachineInterfaceRequest) (*CreateVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualMachineInterface(ctx, request)
}
func (service *BaseService) UpdateVirtualMachineInterface(ctx context.Context, request *UpdateVirtualMachineInterfaceRequest) (*UpdateVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualMachineInterface(ctx, request)
}
func (service *BaseService) DeleteVirtualMachineInterface(ctx context.Context, request *DeleteVirtualMachineInterfaceRequest) (*DeleteVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualMachineInterface(ctx, request)
}
func (service *BaseService) GetVirtualMachineInterface(ctx context.Context, request *GetVirtualMachineInterfaceRequest) (*GetVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualMachineInterface(ctx, request)
}
func (service *BaseService) ListVirtualMachineInterface(ctx context.Context, request *ListVirtualMachineInterfaceRequest) (*ListVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualMachineInterface(ctx, request)
}

func (service *BaseService) CreateVirtualMachine(ctx context.Context, request *CreateVirtualMachineRequest) (*CreateVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualMachine(ctx, request)
}
func (service *BaseService) UpdateVirtualMachine(ctx context.Context, request *UpdateVirtualMachineRequest) (*UpdateVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualMachine(ctx, request)
}
func (service *BaseService) DeleteVirtualMachine(ctx context.Context, request *DeleteVirtualMachineRequest) (*DeleteVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualMachine(ctx, request)
}
func (service *BaseService) GetVirtualMachine(ctx context.Context, request *GetVirtualMachineRequest) (*GetVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualMachine(ctx, request)
}
func (service *BaseService) ListVirtualMachine(ctx context.Context, request *ListVirtualMachineRequest) (*ListVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualMachine(ctx, request)
}

func (service *BaseService) CreateVirtualNetwork(ctx context.Context, request *CreateVirtualNetworkRequest) (*CreateVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualNetwork(ctx, request)
}
func (service *BaseService) UpdateVirtualNetwork(ctx context.Context, request *UpdateVirtualNetworkRequest) (*UpdateVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualNetwork(ctx, request)
}
func (service *BaseService) DeleteVirtualNetwork(ctx context.Context, request *DeleteVirtualNetworkRequest) (*DeleteVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualNetwork(ctx, request)
}
func (service *BaseService) GetVirtualNetwork(ctx context.Context, request *GetVirtualNetworkRequest) (*GetVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualNetwork(ctx, request)
}
func (service *BaseService) ListVirtualNetwork(ctx context.Context, request *ListVirtualNetworkRequest) (*ListVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualNetwork(ctx, request)
}

func (service *BaseService) CreateVirtualRouter(ctx context.Context, request *CreateVirtualRouterRequest) (*CreateVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualRouter(ctx, request)
}
func (service *BaseService) UpdateVirtualRouter(ctx context.Context, request *UpdateVirtualRouterRequest) (*UpdateVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualRouter(ctx, request)
}
func (service *BaseService) DeleteVirtualRouter(ctx context.Context, request *DeleteVirtualRouterRequest) (*DeleteVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualRouter(ctx, request)
}
func (service *BaseService) GetVirtualRouter(ctx context.Context, request *GetVirtualRouterRequest) (*GetVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualRouter(ctx, request)
}
func (service *BaseService) ListVirtualRouter(ctx context.Context, request *ListVirtualRouterRequest) (*ListVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualRouter(ctx, request)
}

func (service *BaseService) CreateAppformixNode(ctx context.Context, request *CreateAppformixNodeRequest) (*CreateAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAppformixNode(ctx, request)
}
func (service *BaseService) UpdateAppformixNode(ctx context.Context, request *UpdateAppformixNodeRequest) (*UpdateAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAppformixNode(ctx, request)
}
func (service *BaseService) DeleteAppformixNode(ctx context.Context, request *DeleteAppformixNodeRequest) (*DeleteAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAppformixNode(ctx, request)
}
func (service *BaseService) GetAppformixNode(ctx context.Context, request *GetAppformixNodeRequest) (*GetAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAppformixNode(ctx, request)
}
func (service *BaseService) ListAppformixNode(ctx context.Context, request *ListAppformixNodeRequest) (*ListAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAppformixNode(ctx, request)
}

func (service *BaseService) CreateBaremetalNode(ctx context.Context, request *CreateBaremetalNodeRequest) (*CreateBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBaremetalNode(ctx, request)
}
func (service *BaseService) UpdateBaremetalNode(ctx context.Context, request *UpdateBaremetalNodeRequest) (*UpdateBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBaremetalNode(ctx, request)
}
func (service *BaseService) DeleteBaremetalNode(ctx context.Context, request *DeleteBaremetalNodeRequest) (*DeleteBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBaremetalNode(ctx, request)
}
func (service *BaseService) GetBaremetalNode(ctx context.Context, request *GetBaremetalNodeRequest) (*GetBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBaremetalNode(ctx, request)
}
func (service *BaseService) ListBaremetalNode(ctx context.Context, request *ListBaremetalNodeRequest) (*ListBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBaremetalNode(ctx, request)
}

func (service *BaseService) CreateBaremetalPort(ctx context.Context, request *CreateBaremetalPortRequest) (*CreateBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBaremetalPort(ctx, request)
}
func (service *BaseService) UpdateBaremetalPort(ctx context.Context, request *UpdateBaremetalPortRequest) (*UpdateBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBaremetalPort(ctx, request)
}
func (service *BaseService) DeleteBaremetalPort(ctx context.Context, request *DeleteBaremetalPortRequest) (*DeleteBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBaremetalPort(ctx, request)
}
func (service *BaseService) GetBaremetalPort(ctx context.Context, request *GetBaremetalPortRequest) (*GetBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBaremetalPort(ctx, request)
}
func (service *BaseService) ListBaremetalPort(ctx context.Context, request *ListBaremetalPortRequest) (*ListBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBaremetalPort(ctx, request)
}

func (service *BaseService) CreateContrailAnalyticsDatabaseNode(ctx context.Context, request *CreateContrailAnalyticsDatabaseNodeRequest) (*CreateContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) UpdateContrailAnalyticsDatabaseNode(ctx context.Context, request *UpdateContrailAnalyticsDatabaseNodeRequest) (*UpdateContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) DeleteContrailAnalyticsDatabaseNode(ctx context.Context, request *DeleteContrailAnalyticsDatabaseNodeRequest) (*DeleteContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) GetContrailAnalyticsDatabaseNode(ctx context.Context, request *GetContrailAnalyticsDatabaseNodeRequest) (*GetContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) ListContrailAnalyticsDatabaseNode(ctx context.Context, request *ListContrailAnalyticsDatabaseNodeRequest) (*ListContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailAnalyticsDatabaseNode(ctx, request)
}

func (service *BaseService) CreateContrailAnalyticsNode(ctx context.Context, request *CreateContrailAnalyticsNodeRequest) (*CreateContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) UpdateContrailAnalyticsNode(ctx context.Context, request *UpdateContrailAnalyticsNodeRequest) (*UpdateContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) DeleteContrailAnalyticsNode(ctx context.Context, request *DeleteContrailAnalyticsNodeRequest) (*DeleteContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) GetContrailAnalyticsNode(ctx context.Context, request *GetContrailAnalyticsNodeRequest) (*GetContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) ListContrailAnalyticsNode(ctx context.Context, request *ListContrailAnalyticsNodeRequest) (*ListContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailAnalyticsNode(ctx, request)
}

func (service *BaseService) CreateContrailCluster(ctx context.Context, request *CreateContrailClusterRequest) (*CreateContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailCluster(ctx, request)
}
func (service *BaseService) UpdateContrailCluster(ctx context.Context, request *UpdateContrailClusterRequest) (*UpdateContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailCluster(ctx, request)
}
func (service *BaseService) DeleteContrailCluster(ctx context.Context, request *DeleteContrailClusterRequest) (*DeleteContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailCluster(ctx, request)
}
func (service *BaseService) GetContrailCluster(ctx context.Context, request *GetContrailClusterRequest) (*GetContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailCluster(ctx, request)
}
func (service *BaseService) ListContrailCluster(ctx context.Context, request *ListContrailClusterRequest) (*ListContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailCluster(ctx, request)
}

func (service *BaseService) CreateContrailConfigDatabaseNode(ctx context.Context, request *CreateContrailConfigDatabaseNodeRequest) (*CreateContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) UpdateContrailConfigDatabaseNode(ctx context.Context, request *UpdateContrailConfigDatabaseNodeRequest) (*UpdateContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) DeleteContrailConfigDatabaseNode(ctx context.Context, request *DeleteContrailConfigDatabaseNodeRequest) (*DeleteContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) GetContrailConfigDatabaseNode(ctx context.Context, request *GetContrailConfigDatabaseNodeRequest) (*GetContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) ListContrailConfigDatabaseNode(ctx context.Context, request *ListContrailConfigDatabaseNodeRequest) (*ListContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailConfigDatabaseNode(ctx, request)
}

func (service *BaseService) CreateContrailConfigNode(ctx context.Context, request *CreateContrailConfigNodeRequest) (*CreateContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailConfigNode(ctx, request)
}
func (service *BaseService) UpdateContrailConfigNode(ctx context.Context, request *UpdateContrailConfigNodeRequest) (*UpdateContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailConfigNode(ctx, request)
}
func (service *BaseService) DeleteContrailConfigNode(ctx context.Context, request *DeleteContrailConfigNodeRequest) (*DeleteContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailConfigNode(ctx, request)
}
func (service *BaseService) GetContrailConfigNode(ctx context.Context, request *GetContrailConfigNodeRequest) (*GetContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailConfigNode(ctx, request)
}
func (service *BaseService) ListContrailConfigNode(ctx context.Context, request *ListContrailConfigNodeRequest) (*ListContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailConfigNode(ctx, request)
}

func (service *BaseService) CreateContrailControlNode(ctx context.Context, request *CreateContrailControlNodeRequest) (*CreateContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailControlNode(ctx, request)
}
func (service *BaseService) UpdateContrailControlNode(ctx context.Context, request *UpdateContrailControlNodeRequest) (*UpdateContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailControlNode(ctx, request)
}
func (service *BaseService) DeleteContrailControlNode(ctx context.Context, request *DeleteContrailControlNodeRequest) (*DeleteContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailControlNode(ctx, request)
}
func (service *BaseService) GetContrailControlNode(ctx context.Context, request *GetContrailControlNodeRequest) (*GetContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailControlNode(ctx, request)
}
func (service *BaseService) ListContrailControlNode(ctx context.Context, request *ListContrailControlNodeRequest) (*ListContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailControlNode(ctx, request)
}

func (service *BaseService) CreateContrailStorageNode(ctx context.Context, request *CreateContrailStorageNodeRequest) (*CreateContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailStorageNode(ctx, request)
}
func (service *BaseService) UpdateContrailStorageNode(ctx context.Context, request *UpdateContrailStorageNodeRequest) (*UpdateContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailStorageNode(ctx, request)
}
func (service *BaseService) DeleteContrailStorageNode(ctx context.Context, request *DeleteContrailStorageNodeRequest) (*DeleteContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailStorageNode(ctx, request)
}
func (service *BaseService) GetContrailStorageNode(ctx context.Context, request *GetContrailStorageNodeRequest) (*GetContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailStorageNode(ctx, request)
}
func (service *BaseService) ListContrailStorageNode(ctx context.Context, request *ListContrailStorageNodeRequest) (*ListContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailStorageNode(ctx, request)
}

func (service *BaseService) CreateContrailVrouterNode(ctx context.Context, request *CreateContrailVrouterNodeRequest) (*CreateContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailVrouterNode(ctx, request)
}
func (service *BaseService) UpdateContrailVrouterNode(ctx context.Context, request *UpdateContrailVrouterNodeRequest) (*UpdateContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailVrouterNode(ctx, request)
}
func (service *BaseService) DeleteContrailVrouterNode(ctx context.Context, request *DeleteContrailVrouterNodeRequest) (*DeleteContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailVrouterNode(ctx, request)
}
func (service *BaseService) GetContrailVrouterNode(ctx context.Context, request *GetContrailVrouterNodeRequest) (*GetContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailVrouterNode(ctx, request)
}
func (service *BaseService) ListContrailVrouterNode(ctx context.Context, request *ListContrailVrouterNodeRequest) (*ListContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailVrouterNode(ctx, request)
}

func (service *BaseService) CreateContrailControllerNode(ctx context.Context, request *CreateContrailControllerNodeRequest) (*CreateContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailControllerNode(ctx, request)
}
func (service *BaseService) UpdateContrailControllerNode(ctx context.Context, request *UpdateContrailControllerNodeRequest) (*UpdateContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailControllerNode(ctx, request)
}
func (service *BaseService) DeleteContrailControllerNode(ctx context.Context, request *DeleteContrailControllerNodeRequest) (*DeleteContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailControllerNode(ctx, request)
}
func (service *BaseService) GetContrailControllerNode(ctx context.Context, request *GetContrailControllerNodeRequest) (*GetContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailControllerNode(ctx, request)
}
func (service *BaseService) ListContrailControllerNode(ctx context.Context, request *ListContrailControllerNodeRequest) (*ListContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailControllerNode(ctx, request)
}

func (service *BaseService) CreateDashboard(ctx context.Context, request *CreateDashboardRequest) (*CreateDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDashboard(ctx, request)
}
func (service *BaseService) UpdateDashboard(ctx context.Context, request *UpdateDashboardRequest) (*UpdateDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDashboard(ctx, request)
}
func (service *BaseService) DeleteDashboard(ctx context.Context, request *DeleteDashboardRequest) (*DeleteDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDashboard(ctx, request)
}
func (service *BaseService) GetDashboard(ctx context.Context, request *GetDashboardRequest) (*GetDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDashboard(ctx, request)
}
func (service *BaseService) ListDashboard(ctx context.Context, request *ListDashboardRequest) (*ListDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDashboard(ctx, request)
}

func (service *BaseService) CreateFlavor(ctx context.Context, request *CreateFlavorRequest) (*CreateFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFlavor(ctx, request)
}
func (service *BaseService) UpdateFlavor(ctx context.Context, request *UpdateFlavorRequest) (*UpdateFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFlavor(ctx, request)
}
func (service *BaseService) DeleteFlavor(ctx context.Context, request *DeleteFlavorRequest) (*DeleteFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFlavor(ctx, request)
}
func (service *BaseService) GetFlavor(ctx context.Context, request *GetFlavorRequest) (*GetFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFlavor(ctx, request)
}
func (service *BaseService) ListFlavor(ctx context.Context, request *ListFlavorRequest) (*ListFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFlavor(ctx, request)
}

func (service *BaseService) CreateOsImage(ctx context.Context, request *CreateOsImageRequest) (*CreateOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateOsImage(ctx, request)
}
func (service *BaseService) UpdateOsImage(ctx context.Context, request *UpdateOsImageRequest) (*UpdateOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateOsImage(ctx, request)
}
func (service *BaseService) DeleteOsImage(ctx context.Context, request *DeleteOsImageRequest) (*DeleteOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteOsImage(ctx, request)
}
func (service *BaseService) GetOsImage(ctx context.Context, request *GetOsImageRequest) (*GetOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetOsImage(ctx, request)
}
func (service *BaseService) ListOsImage(ctx context.Context, request *ListOsImageRequest) (*ListOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListOsImage(ctx, request)
}

func (service *BaseService) CreateKeypair(ctx context.Context, request *CreateKeypairRequest) (*CreateKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateKeypair(ctx, request)
}
func (service *BaseService) UpdateKeypair(ctx context.Context, request *UpdateKeypairRequest) (*UpdateKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateKeypair(ctx, request)
}
func (service *BaseService) DeleteKeypair(ctx context.Context, request *DeleteKeypairRequest) (*DeleteKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteKeypair(ctx, request)
}
func (service *BaseService) GetKeypair(ctx context.Context, request *GetKeypairRequest) (*GetKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetKeypair(ctx, request)
}
func (service *BaseService) ListKeypair(ctx context.Context, request *ListKeypairRequest) (*ListKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListKeypair(ctx, request)
}

func (service *BaseService) CreateKubernetesMasterNode(ctx context.Context, request *CreateKubernetesMasterNodeRequest) (*CreateKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesMasterNode(ctx, request)
}
func (service *BaseService) UpdateKubernetesMasterNode(ctx context.Context, request *UpdateKubernetesMasterNodeRequest) (*UpdateKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesMasterNode(ctx, request)
}
func (service *BaseService) DeleteKubernetesMasterNode(ctx context.Context, request *DeleteKubernetesMasterNodeRequest) (*DeleteKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesMasterNode(ctx, request)
}
func (service *BaseService) GetKubernetesMasterNode(ctx context.Context, request *GetKubernetesMasterNodeRequest) (*GetKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetKubernetesMasterNode(ctx, request)
}
func (service *BaseService) ListKubernetesMasterNode(ctx context.Context, request *ListKubernetesMasterNodeRequest) (*ListKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListKubernetesMasterNode(ctx, request)
}

func (service *BaseService) CreateKubernetesNode(ctx context.Context, request *CreateKubernetesNodeRequest) (*CreateKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesNode(ctx, request)
}
func (service *BaseService) UpdateKubernetesNode(ctx context.Context, request *UpdateKubernetesNodeRequest) (*UpdateKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesNode(ctx, request)
}
func (service *BaseService) DeleteKubernetesNode(ctx context.Context, request *DeleteKubernetesNodeRequest) (*DeleteKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesNode(ctx, request)
}
func (service *BaseService) GetKubernetesNode(ctx context.Context, request *GetKubernetesNodeRequest) (*GetKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetKubernetesNode(ctx, request)
}
func (service *BaseService) ListKubernetesNode(ctx context.Context, request *ListKubernetesNodeRequest) (*ListKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListKubernetesNode(ctx, request)
}

func (service *BaseService) CreateLocation(ctx context.Context, request *CreateLocationRequest) (*CreateLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLocation(ctx, request)
}
func (service *BaseService) UpdateLocation(ctx context.Context, request *UpdateLocationRequest) (*UpdateLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLocation(ctx, request)
}
func (service *BaseService) DeleteLocation(ctx context.Context, request *DeleteLocationRequest) (*DeleteLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLocation(ctx, request)
}
func (service *BaseService) GetLocation(ctx context.Context, request *GetLocationRequest) (*GetLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLocation(ctx, request)
}
func (service *BaseService) ListLocation(ctx context.Context, request *ListLocationRequest) (*ListLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLocation(ctx, request)
}

func (service *BaseService) CreateNode(ctx context.Context, request *CreateNodeRequest) (*CreateNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNode(ctx, request)
}
func (service *BaseService) UpdateNode(ctx context.Context, request *UpdateNodeRequest) (*UpdateNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNode(ctx, request)
}
func (service *BaseService) DeleteNode(ctx context.Context, request *DeleteNodeRequest) (*DeleteNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNode(ctx, request)
}
func (service *BaseService) GetNode(ctx context.Context, request *GetNodeRequest) (*GetNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNode(ctx, request)
}
func (service *BaseService) ListNode(ctx context.Context, request *ListNodeRequest) (*ListNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNode(ctx, request)
}

func (service *BaseService) CreateServer(ctx context.Context, request *CreateServerRequest) (*CreateServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServer(ctx, request)
}
func (service *BaseService) UpdateServer(ctx context.Context, request *UpdateServerRequest) (*UpdateServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServer(ctx, request)
}
func (service *BaseService) DeleteServer(ctx context.Context, request *DeleteServerRequest) (*DeleteServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServer(ctx, request)
}
func (service *BaseService) GetServer(ctx context.Context, request *GetServerRequest) (*GetServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServer(ctx, request)
}
func (service *BaseService) ListServer(ctx context.Context, request *ListServerRequest) (*ListServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServer(ctx, request)
}

func (service *BaseService) CreateVPNGroup(ctx context.Context, request *CreateVPNGroupRequest) (*CreateVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVPNGroup(ctx, request)
}
func (service *BaseService) UpdateVPNGroup(ctx context.Context, request *UpdateVPNGroupRequest) (*UpdateVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVPNGroup(ctx, request)
}
func (service *BaseService) DeleteVPNGroup(ctx context.Context, request *DeleteVPNGroupRequest) (*DeleteVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVPNGroup(ctx, request)
}
func (service *BaseService) GetVPNGroup(ctx context.Context, request *GetVPNGroupRequest) (*GetVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVPNGroup(ctx, request)
}
func (service *BaseService) ListVPNGroup(ctx context.Context, request *ListVPNGroupRequest) (*ListVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVPNGroup(ctx, request)
}

func (service *BaseService) CreateWidget(ctx context.Context, request *CreateWidgetRequest) (*CreateWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateWidget(ctx, request)
}
func (service *BaseService) UpdateWidget(ctx context.Context, request *UpdateWidgetRequest) (*UpdateWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateWidget(ctx, request)
}
func (service *BaseService) DeleteWidget(ctx context.Context, request *DeleteWidgetRequest) (*DeleteWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteWidget(ctx, request)
}
func (service *BaseService) GetWidget(ctx context.Context, request *GetWidgetRequest) (*GetWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetWidget(ctx, request)
}
func (service *BaseService) ListWidget(ctx context.Context, request *ListWidgetRequest) (*ListWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListWidget(ctx, request)
}
