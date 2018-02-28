package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
)

type Service interface {
	Next() Service
	SetNext(Service)

	CreateAccessControlList(context.Context, *models.CreateAccessControlListRequest) (*models.CreateAccessControlListResponse, error)
	UpdateAccessControlList(context.Context, *models.UpdateAccessControlListRequest) (*models.UpdateAccessControlListResponse, error)
	DeleteAccessControlList(context.Context, *models.DeleteAccessControlListRequest) (*models.DeleteAccessControlListResponse, error)
	GetAccessControlList(context.Context, *models.GetAccessControlListRequest) (*models.GetAccessControlListResponse, error)
	ListAccessControlList(context.Context, *models.ListAccessControlListRequest) (*models.ListAccessControlListResponse, error)

	CreateAddressGroup(context.Context, *models.CreateAddressGroupRequest) (*models.CreateAddressGroupResponse, error)
	UpdateAddressGroup(context.Context, *models.UpdateAddressGroupRequest) (*models.UpdateAddressGroupResponse, error)
	DeleteAddressGroup(context.Context, *models.DeleteAddressGroupRequest) (*models.DeleteAddressGroupResponse, error)
	GetAddressGroup(context.Context, *models.GetAddressGroupRequest) (*models.GetAddressGroupResponse, error)
	ListAddressGroup(context.Context, *models.ListAddressGroupRequest) (*models.ListAddressGroupResponse, error)

	CreateAlarm(context.Context, *models.CreateAlarmRequest) (*models.CreateAlarmResponse, error)
	UpdateAlarm(context.Context, *models.UpdateAlarmRequest) (*models.UpdateAlarmResponse, error)
	DeleteAlarm(context.Context, *models.DeleteAlarmRequest) (*models.DeleteAlarmResponse, error)
	GetAlarm(context.Context, *models.GetAlarmRequest) (*models.GetAlarmResponse, error)
	ListAlarm(context.Context, *models.ListAlarmRequest) (*models.ListAlarmResponse, error)

	CreateAliasIPPool(context.Context, *models.CreateAliasIPPoolRequest) (*models.CreateAliasIPPoolResponse, error)
	UpdateAliasIPPool(context.Context, *models.UpdateAliasIPPoolRequest) (*models.UpdateAliasIPPoolResponse, error)
	DeleteAliasIPPool(context.Context, *models.DeleteAliasIPPoolRequest) (*models.DeleteAliasIPPoolResponse, error)
	GetAliasIPPool(context.Context, *models.GetAliasIPPoolRequest) (*models.GetAliasIPPoolResponse, error)
	ListAliasIPPool(context.Context, *models.ListAliasIPPoolRequest) (*models.ListAliasIPPoolResponse, error)

	CreateAliasIP(context.Context, *models.CreateAliasIPRequest) (*models.CreateAliasIPResponse, error)
	UpdateAliasIP(context.Context, *models.UpdateAliasIPRequest) (*models.UpdateAliasIPResponse, error)
	DeleteAliasIP(context.Context, *models.DeleteAliasIPRequest) (*models.DeleteAliasIPResponse, error)
	GetAliasIP(context.Context, *models.GetAliasIPRequest) (*models.GetAliasIPResponse, error)
	ListAliasIP(context.Context, *models.ListAliasIPRequest) (*models.ListAliasIPResponse, error)

	CreateAnalyticsNode(context.Context, *models.CreateAnalyticsNodeRequest) (*models.CreateAnalyticsNodeResponse, error)
	UpdateAnalyticsNode(context.Context, *models.UpdateAnalyticsNodeRequest) (*models.UpdateAnalyticsNodeResponse, error)
	DeleteAnalyticsNode(context.Context, *models.DeleteAnalyticsNodeRequest) (*models.DeleteAnalyticsNodeResponse, error)
	GetAnalyticsNode(context.Context, *models.GetAnalyticsNodeRequest) (*models.GetAnalyticsNodeResponse, error)
	ListAnalyticsNode(context.Context, *models.ListAnalyticsNodeRequest) (*models.ListAnalyticsNodeResponse, error)

	CreateAPIAccessList(context.Context, *models.CreateAPIAccessListRequest) (*models.CreateAPIAccessListResponse, error)
	UpdateAPIAccessList(context.Context, *models.UpdateAPIAccessListRequest) (*models.UpdateAPIAccessListResponse, error)
	DeleteAPIAccessList(context.Context, *models.DeleteAPIAccessListRequest) (*models.DeleteAPIAccessListResponse, error)
	GetAPIAccessList(context.Context, *models.GetAPIAccessListRequest) (*models.GetAPIAccessListResponse, error)
	ListAPIAccessList(context.Context, *models.ListAPIAccessListRequest) (*models.ListAPIAccessListResponse, error)

	CreateApplicationPolicySet(context.Context, *models.CreateApplicationPolicySetRequest) (*models.CreateApplicationPolicySetResponse, error)
	UpdateApplicationPolicySet(context.Context, *models.UpdateApplicationPolicySetRequest) (*models.UpdateApplicationPolicySetResponse, error)
	DeleteApplicationPolicySet(context.Context, *models.DeleteApplicationPolicySetRequest) (*models.DeleteApplicationPolicySetResponse, error)
	GetApplicationPolicySet(context.Context, *models.GetApplicationPolicySetRequest) (*models.GetApplicationPolicySetResponse, error)
	ListApplicationPolicySet(context.Context, *models.ListApplicationPolicySetRequest) (*models.ListApplicationPolicySetResponse, error)

	CreateBGPAsAService(context.Context, *models.CreateBGPAsAServiceRequest) (*models.CreateBGPAsAServiceResponse, error)
	UpdateBGPAsAService(context.Context, *models.UpdateBGPAsAServiceRequest) (*models.UpdateBGPAsAServiceResponse, error)
	DeleteBGPAsAService(context.Context, *models.DeleteBGPAsAServiceRequest) (*models.DeleteBGPAsAServiceResponse, error)
	GetBGPAsAService(context.Context, *models.GetBGPAsAServiceRequest) (*models.GetBGPAsAServiceResponse, error)
	ListBGPAsAService(context.Context, *models.ListBGPAsAServiceRequest) (*models.ListBGPAsAServiceResponse, error)

	CreateBGPRouter(context.Context, *models.CreateBGPRouterRequest) (*models.CreateBGPRouterResponse, error)
	UpdateBGPRouter(context.Context, *models.UpdateBGPRouterRequest) (*models.UpdateBGPRouterResponse, error)
	DeleteBGPRouter(context.Context, *models.DeleteBGPRouterRequest) (*models.DeleteBGPRouterResponse, error)
	GetBGPRouter(context.Context, *models.GetBGPRouterRequest) (*models.GetBGPRouterResponse, error)
	ListBGPRouter(context.Context, *models.ListBGPRouterRequest) (*models.ListBGPRouterResponse, error)

	CreateBGPVPN(context.Context, *models.CreateBGPVPNRequest) (*models.CreateBGPVPNResponse, error)
	UpdateBGPVPN(context.Context, *models.UpdateBGPVPNRequest) (*models.UpdateBGPVPNResponse, error)
	DeleteBGPVPN(context.Context, *models.DeleteBGPVPNRequest) (*models.DeleteBGPVPNResponse, error)
	GetBGPVPN(context.Context, *models.GetBGPVPNRequest) (*models.GetBGPVPNResponse, error)
	ListBGPVPN(context.Context, *models.ListBGPVPNRequest) (*models.ListBGPVPNResponse, error)

	CreateBridgeDomain(context.Context, *models.CreateBridgeDomainRequest) (*models.CreateBridgeDomainResponse, error)
	UpdateBridgeDomain(context.Context, *models.UpdateBridgeDomainRequest) (*models.UpdateBridgeDomainResponse, error)
	DeleteBridgeDomain(context.Context, *models.DeleteBridgeDomainRequest) (*models.DeleteBridgeDomainResponse, error)
	GetBridgeDomain(context.Context, *models.GetBridgeDomainRequest) (*models.GetBridgeDomainResponse, error)
	ListBridgeDomain(context.Context, *models.ListBridgeDomainRequest) (*models.ListBridgeDomainResponse, error)

	CreateConfigNode(context.Context, *models.CreateConfigNodeRequest) (*models.CreateConfigNodeResponse, error)
	UpdateConfigNode(context.Context, *models.UpdateConfigNodeRequest) (*models.UpdateConfigNodeResponse, error)
	DeleteConfigNode(context.Context, *models.DeleteConfigNodeRequest) (*models.DeleteConfigNodeResponse, error)
	GetConfigNode(context.Context, *models.GetConfigNodeRequest) (*models.GetConfigNodeResponse, error)
	ListConfigNode(context.Context, *models.ListConfigNodeRequest) (*models.ListConfigNodeResponse, error)

	CreateConfigRoot(context.Context, *models.CreateConfigRootRequest) (*models.CreateConfigRootResponse, error)
	UpdateConfigRoot(context.Context, *models.UpdateConfigRootRequest) (*models.UpdateConfigRootResponse, error)
	DeleteConfigRoot(context.Context, *models.DeleteConfigRootRequest) (*models.DeleteConfigRootResponse, error)
	GetConfigRoot(context.Context, *models.GetConfigRootRequest) (*models.GetConfigRootResponse, error)
	ListConfigRoot(context.Context, *models.ListConfigRootRequest) (*models.ListConfigRootResponse, error)

	CreateCustomerAttachment(context.Context, *models.CreateCustomerAttachmentRequest) (*models.CreateCustomerAttachmentResponse, error)
	UpdateCustomerAttachment(context.Context, *models.UpdateCustomerAttachmentRequest) (*models.UpdateCustomerAttachmentResponse, error)
	DeleteCustomerAttachment(context.Context, *models.DeleteCustomerAttachmentRequest) (*models.DeleteCustomerAttachmentResponse, error)
	GetCustomerAttachment(context.Context, *models.GetCustomerAttachmentRequest) (*models.GetCustomerAttachmentResponse, error)
	ListCustomerAttachment(context.Context, *models.ListCustomerAttachmentRequest) (*models.ListCustomerAttachmentResponse, error)

	CreateDatabaseNode(context.Context, *models.CreateDatabaseNodeRequest) (*models.CreateDatabaseNodeResponse, error)
	UpdateDatabaseNode(context.Context, *models.UpdateDatabaseNodeRequest) (*models.UpdateDatabaseNodeResponse, error)
	DeleteDatabaseNode(context.Context, *models.DeleteDatabaseNodeRequest) (*models.DeleteDatabaseNodeResponse, error)
	GetDatabaseNode(context.Context, *models.GetDatabaseNodeRequest) (*models.GetDatabaseNodeResponse, error)
	ListDatabaseNode(context.Context, *models.ListDatabaseNodeRequest) (*models.ListDatabaseNodeResponse, error)

	CreateDiscoveryServiceAssignment(context.Context, *models.CreateDiscoveryServiceAssignmentRequest) (*models.CreateDiscoveryServiceAssignmentResponse, error)
	UpdateDiscoveryServiceAssignment(context.Context, *models.UpdateDiscoveryServiceAssignmentRequest) (*models.UpdateDiscoveryServiceAssignmentResponse, error)
	DeleteDiscoveryServiceAssignment(context.Context, *models.DeleteDiscoveryServiceAssignmentRequest) (*models.DeleteDiscoveryServiceAssignmentResponse, error)
	GetDiscoveryServiceAssignment(context.Context, *models.GetDiscoveryServiceAssignmentRequest) (*models.GetDiscoveryServiceAssignmentResponse, error)
	ListDiscoveryServiceAssignment(context.Context, *models.ListDiscoveryServiceAssignmentRequest) (*models.ListDiscoveryServiceAssignmentResponse, error)

	CreateDomain(context.Context, *models.CreateDomainRequest) (*models.CreateDomainResponse, error)
	UpdateDomain(context.Context, *models.UpdateDomainRequest) (*models.UpdateDomainResponse, error)
	DeleteDomain(context.Context, *models.DeleteDomainRequest) (*models.DeleteDomainResponse, error)
	GetDomain(context.Context, *models.GetDomainRequest) (*models.GetDomainResponse, error)
	ListDomain(context.Context, *models.ListDomainRequest) (*models.ListDomainResponse, error)

	CreateDsaRule(context.Context, *models.CreateDsaRuleRequest) (*models.CreateDsaRuleResponse, error)
	UpdateDsaRule(context.Context, *models.UpdateDsaRuleRequest) (*models.UpdateDsaRuleResponse, error)
	DeleteDsaRule(context.Context, *models.DeleteDsaRuleRequest) (*models.DeleteDsaRuleResponse, error)
	GetDsaRule(context.Context, *models.GetDsaRuleRequest) (*models.GetDsaRuleResponse, error)
	ListDsaRule(context.Context, *models.ListDsaRuleRequest) (*models.ListDsaRuleResponse, error)

	CreateE2ServiceProvider(context.Context, *models.CreateE2ServiceProviderRequest) (*models.CreateE2ServiceProviderResponse, error)
	UpdateE2ServiceProvider(context.Context, *models.UpdateE2ServiceProviderRequest) (*models.UpdateE2ServiceProviderResponse, error)
	DeleteE2ServiceProvider(context.Context, *models.DeleteE2ServiceProviderRequest) (*models.DeleteE2ServiceProviderResponse, error)
	GetE2ServiceProvider(context.Context, *models.GetE2ServiceProviderRequest) (*models.GetE2ServiceProviderResponse, error)
	ListE2ServiceProvider(context.Context, *models.ListE2ServiceProviderRequest) (*models.ListE2ServiceProviderResponse, error)

	CreateFirewallPolicy(context.Context, *models.CreateFirewallPolicyRequest) (*models.CreateFirewallPolicyResponse, error)
	UpdateFirewallPolicy(context.Context, *models.UpdateFirewallPolicyRequest) (*models.UpdateFirewallPolicyResponse, error)
	DeleteFirewallPolicy(context.Context, *models.DeleteFirewallPolicyRequest) (*models.DeleteFirewallPolicyResponse, error)
	GetFirewallPolicy(context.Context, *models.GetFirewallPolicyRequest) (*models.GetFirewallPolicyResponse, error)
	ListFirewallPolicy(context.Context, *models.ListFirewallPolicyRequest) (*models.ListFirewallPolicyResponse, error)

	CreateFirewallRule(context.Context, *models.CreateFirewallRuleRequest) (*models.CreateFirewallRuleResponse, error)
	UpdateFirewallRule(context.Context, *models.UpdateFirewallRuleRequest) (*models.UpdateFirewallRuleResponse, error)
	DeleteFirewallRule(context.Context, *models.DeleteFirewallRuleRequest) (*models.DeleteFirewallRuleResponse, error)
	GetFirewallRule(context.Context, *models.GetFirewallRuleRequest) (*models.GetFirewallRuleResponse, error)
	ListFirewallRule(context.Context, *models.ListFirewallRuleRequest) (*models.ListFirewallRuleResponse, error)

	CreateFloatingIPPool(context.Context, *models.CreateFloatingIPPoolRequest) (*models.CreateFloatingIPPoolResponse, error)
	UpdateFloatingIPPool(context.Context, *models.UpdateFloatingIPPoolRequest) (*models.UpdateFloatingIPPoolResponse, error)
	DeleteFloatingIPPool(context.Context, *models.DeleteFloatingIPPoolRequest) (*models.DeleteFloatingIPPoolResponse, error)
	GetFloatingIPPool(context.Context, *models.GetFloatingIPPoolRequest) (*models.GetFloatingIPPoolResponse, error)
	ListFloatingIPPool(context.Context, *models.ListFloatingIPPoolRequest) (*models.ListFloatingIPPoolResponse, error)

	CreateFloatingIP(context.Context, *models.CreateFloatingIPRequest) (*models.CreateFloatingIPResponse, error)
	UpdateFloatingIP(context.Context, *models.UpdateFloatingIPRequest) (*models.UpdateFloatingIPResponse, error)
	DeleteFloatingIP(context.Context, *models.DeleteFloatingIPRequest) (*models.DeleteFloatingIPResponse, error)
	GetFloatingIP(context.Context, *models.GetFloatingIPRequest) (*models.GetFloatingIPResponse, error)
	ListFloatingIP(context.Context, *models.ListFloatingIPRequest) (*models.ListFloatingIPResponse, error)

	CreateForwardingClass(context.Context, *models.CreateForwardingClassRequest) (*models.CreateForwardingClassResponse, error)
	UpdateForwardingClass(context.Context, *models.UpdateForwardingClassRequest) (*models.UpdateForwardingClassResponse, error)
	DeleteForwardingClass(context.Context, *models.DeleteForwardingClassRequest) (*models.DeleteForwardingClassResponse, error)
	GetForwardingClass(context.Context, *models.GetForwardingClassRequest) (*models.GetForwardingClassResponse, error)
	ListForwardingClass(context.Context, *models.ListForwardingClassRequest) (*models.ListForwardingClassResponse, error)

	CreateGlobalQosConfig(context.Context, *models.CreateGlobalQosConfigRequest) (*models.CreateGlobalQosConfigResponse, error)
	UpdateGlobalQosConfig(context.Context, *models.UpdateGlobalQosConfigRequest) (*models.UpdateGlobalQosConfigResponse, error)
	DeleteGlobalQosConfig(context.Context, *models.DeleteGlobalQosConfigRequest) (*models.DeleteGlobalQosConfigResponse, error)
	GetGlobalQosConfig(context.Context, *models.GetGlobalQosConfigRequest) (*models.GetGlobalQosConfigResponse, error)
	ListGlobalQosConfig(context.Context, *models.ListGlobalQosConfigRequest) (*models.ListGlobalQosConfigResponse, error)

	CreateGlobalSystemConfig(context.Context, *models.CreateGlobalSystemConfigRequest) (*models.CreateGlobalSystemConfigResponse, error)
	UpdateGlobalSystemConfig(context.Context, *models.UpdateGlobalSystemConfigRequest) (*models.UpdateGlobalSystemConfigResponse, error)
	DeleteGlobalSystemConfig(context.Context, *models.DeleteGlobalSystemConfigRequest) (*models.DeleteGlobalSystemConfigResponse, error)
	GetGlobalSystemConfig(context.Context, *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error)
	ListGlobalSystemConfig(context.Context, *models.ListGlobalSystemConfigRequest) (*models.ListGlobalSystemConfigResponse, error)

	CreateGlobalVrouterConfig(context.Context, *models.CreateGlobalVrouterConfigRequest) (*models.CreateGlobalVrouterConfigResponse, error)
	UpdateGlobalVrouterConfig(context.Context, *models.UpdateGlobalVrouterConfigRequest) (*models.UpdateGlobalVrouterConfigResponse, error)
	DeleteGlobalVrouterConfig(context.Context, *models.DeleteGlobalVrouterConfigRequest) (*models.DeleteGlobalVrouterConfigResponse, error)
	GetGlobalVrouterConfig(context.Context, *models.GetGlobalVrouterConfigRequest) (*models.GetGlobalVrouterConfigResponse, error)
	ListGlobalVrouterConfig(context.Context, *models.ListGlobalVrouterConfigRequest) (*models.ListGlobalVrouterConfigResponse, error)

	CreateInstanceIP(context.Context, *models.CreateInstanceIPRequest) (*models.CreateInstanceIPResponse, error)
	UpdateInstanceIP(context.Context, *models.UpdateInstanceIPRequest) (*models.UpdateInstanceIPResponse, error)
	DeleteInstanceIP(context.Context, *models.DeleteInstanceIPRequest) (*models.DeleteInstanceIPResponse, error)
	GetInstanceIP(context.Context, *models.GetInstanceIPRequest) (*models.GetInstanceIPResponse, error)
	ListInstanceIP(context.Context, *models.ListInstanceIPRequest) (*models.ListInstanceIPResponse, error)

	CreateInterfaceRouteTable(context.Context, *models.CreateInterfaceRouteTableRequest) (*models.CreateInterfaceRouteTableResponse, error)
	UpdateInterfaceRouteTable(context.Context, *models.UpdateInterfaceRouteTableRequest) (*models.UpdateInterfaceRouteTableResponse, error)
	DeleteInterfaceRouteTable(context.Context, *models.DeleteInterfaceRouteTableRequest) (*models.DeleteInterfaceRouteTableResponse, error)
	GetInterfaceRouteTable(context.Context, *models.GetInterfaceRouteTableRequest) (*models.GetInterfaceRouteTableResponse, error)
	ListInterfaceRouteTable(context.Context, *models.ListInterfaceRouteTableRequest) (*models.ListInterfaceRouteTableResponse, error)

	CreateLoadbalancerHealthmonitor(context.Context, *models.CreateLoadbalancerHealthmonitorRequest) (*models.CreateLoadbalancerHealthmonitorResponse, error)
	UpdateLoadbalancerHealthmonitor(context.Context, *models.UpdateLoadbalancerHealthmonitorRequest) (*models.UpdateLoadbalancerHealthmonitorResponse, error)
	DeleteLoadbalancerHealthmonitor(context.Context, *models.DeleteLoadbalancerHealthmonitorRequest) (*models.DeleteLoadbalancerHealthmonitorResponse, error)
	GetLoadbalancerHealthmonitor(context.Context, *models.GetLoadbalancerHealthmonitorRequest) (*models.GetLoadbalancerHealthmonitorResponse, error)
	ListLoadbalancerHealthmonitor(context.Context, *models.ListLoadbalancerHealthmonitorRequest) (*models.ListLoadbalancerHealthmonitorResponse, error)

	CreateLoadbalancerListener(context.Context, *models.CreateLoadbalancerListenerRequest) (*models.CreateLoadbalancerListenerResponse, error)
	UpdateLoadbalancerListener(context.Context, *models.UpdateLoadbalancerListenerRequest) (*models.UpdateLoadbalancerListenerResponse, error)
	DeleteLoadbalancerListener(context.Context, *models.DeleteLoadbalancerListenerRequest) (*models.DeleteLoadbalancerListenerResponse, error)
	GetLoadbalancerListener(context.Context, *models.GetLoadbalancerListenerRequest) (*models.GetLoadbalancerListenerResponse, error)
	ListLoadbalancerListener(context.Context, *models.ListLoadbalancerListenerRequest) (*models.ListLoadbalancerListenerResponse, error)

	CreateLoadbalancerMember(context.Context, *models.CreateLoadbalancerMemberRequest) (*models.CreateLoadbalancerMemberResponse, error)
	UpdateLoadbalancerMember(context.Context, *models.UpdateLoadbalancerMemberRequest) (*models.UpdateLoadbalancerMemberResponse, error)
	DeleteLoadbalancerMember(context.Context, *models.DeleteLoadbalancerMemberRequest) (*models.DeleteLoadbalancerMemberResponse, error)
	GetLoadbalancerMember(context.Context, *models.GetLoadbalancerMemberRequest) (*models.GetLoadbalancerMemberResponse, error)
	ListLoadbalancerMember(context.Context, *models.ListLoadbalancerMemberRequest) (*models.ListLoadbalancerMemberResponse, error)

	CreateLoadbalancerPool(context.Context, *models.CreateLoadbalancerPoolRequest) (*models.CreateLoadbalancerPoolResponse, error)
	UpdateLoadbalancerPool(context.Context, *models.UpdateLoadbalancerPoolRequest) (*models.UpdateLoadbalancerPoolResponse, error)
	DeleteLoadbalancerPool(context.Context, *models.DeleteLoadbalancerPoolRequest) (*models.DeleteLoadbalancerPoolResponse, error)
	GetLoadbalancerPool(context.Context, *models.GetLoadbalancerPoolRequest) (*models.GetLoadbalancerPoolResponse, error)
	ListLoadbalancerPool(context.Context, *models.ListLoadbalancerPoolRequest) (*models.ListLoadbalancerPoolResponse, error)

	CreateLoadbalancer(context.Context, *models.CreateLoadbalancerRequest) (*models.CreateLoadbalancerResponse, error)
	UpdateLoadbalancer(context.Context, *models.UpdateLoadbalancerRequest) (*models.UpdateLoadbalancerResponse, error)
	DeleteLoadbalancer(context.Context, *models.DeleteLoadbalancerRequest) (*models.DeleteLoadbalancerResponse, error)
	GetLoadbalancer(context.Context, *models.GetLoadbalancerRequest) (*models.GetLoadbalancerResponse, error)
	ListLoadbalancer(context.Context, *models.ListLoadbalancerRequest) (*models.ListLoadbalancerResponse, error)

	CreateLogicalInterface(context.Context, *models.CreateLogicalInterfaceRequest) (*models.CreateLogicalInterfaceResponse, error)
	UpdateLogicalInterface(context.Context, *models.UpdateLogicalInterfaceRequest) (*models.UpdateLogicalInterfaceResponse, error)
	DeleteLogicalInterface(context.Context, *models.DeleteLogicalInterfaceRequest) (*models.DeleteLogicalInterfaceResponse, error)
	GetLogicalInterface(context.Context, *models.GetLogicalInterfaceRequest) (*models.GetLogicalInterfaceResponse, error)
	ListLogicalInterface(context.Context, *models.ListLogicalInterfaceRequest) (*models.ListLogicalInterfaceResponse, error)

	CreateLogicalRouter(context.Context, *models.CreateLogicalRouterRequest) (*models.CreateLogicalRouterResponse, error)
	UpdateLogicalRouter(context.Context, *models.UpdateLogicalRouterRequest) (*models.UpdateLogicalRouterResponse, error)
	DeleteLogicalRouter(context.Context, *models.DeleteLogicalRouterRequest) (*models.DeleteLogicalRouterResponse, error)
	GetLogicalRouter(context.Context, *models.GetLogicalRouterRequest) (*models.GetLogicalRouterResponse, error)
	ListLogicalRouter(context.Context, *models.ListLogicalRouterRequest) (*models.ListLogicalRouterResponse, error)

	CreateNamespace(context.Context, *models.CreateNamespaceRequest) (*models.CreateNamespaceResponse, error)
	UpdateNamespace(context.Context, *models.UpdateNamespaceRequest) (*models.UpdateNamespaceResponse, error)
	DeleteNamespace(context.Context, *models.DeleteNamespaceRequest) (*models.DeleteNamespaceResponse, error)
	GetNamespace(context.Context, *models.GetNamespaceRequest) (*models.GetNamespaceResponse, error)
	ListNamespace(context.Context, *models.ListNamespaceRequest) (*models.ListNamespaceResponse, error)

	CreateNetworkDeviceConfig(context.Context, *models.CreateNetworkDeviceConfigRequest) (*models.CreateNetworkDeviceConfigResponse, error)
	UpdateNetworkDeviceConfig(context.Context, *models.UpdateNetworkDeviceConfigRequest) (*models.UpdateNetworkDeviceConfigResponse, error)
	DeleteNetworkDeviceConfig(context.Context, *models.DeleteNetworkDeviceConfigRequest) (*models.DeleteNetworkDeviceConfigResponse, error)
	GetNetworkDeviceConfig(context.Context, *models.GetNetworkDeviceConfigRequest) (*models.GetNetworkDeviceConfigResponse, error)
	ListNetworkDeviceConfig(context.Context, *models.ListNetworkDeviceConfigRequest) (*models.ListNetworkDeviceConfigResponse, error)

	CreateNetworkIpam(context.Context, *models.CreateNetworkIpamRequest) (*models.CreateNetworkIpamResponse, error)
	UpdateNetworkIpam(context.Context, *models.UpdateNetworkIpamRequest) (*models.UpdateNetworkIpamResponse, error)
	DeleteNetworkIpam(context.Context, *models.DeleteNetworkIpamRequest) (*models.DeleteNetworkIpamResponse, error)
	GetNetworkIpam(context.Context, *models.GetNetworkIpamRequest) (*models.GetNetworkIpamResponse, error)
	ListNetworkIpam(context.Context, *models.ListNetworkIpamRequest) (*models.ListNetworkIpamResponse, error)

	CreateNetworkPolicy(context.Context, *models.CreateNetworkPolicyRequest) (*models.CreateNetworkPolicyResponse, error)
	UpdateNetworkPolicy(context.Context, *models.UpdateNetworkPolicyRequest) (*models.UpdateNetworkPolicyResponse, error)
	DeleteNetworkPolicy(context.Context, *models.DeleteNetworkPolicyRequest) (*models.DeleteNetworkPolicyResponse, error)
	GetNetworkPolicy(context.Context, *models.GetNetworkPolicyRequest) (*models.GetNetworkPolicyResponse, error)
	ListNetworkPolicy(context.Context, *models.ListNetworkPolicyRequest) (*models.ListNetworkPolicyResponse, error)

	CreatePeeringPolicy(context.Context, *models.CreatePeeringPolicyRequest) (*models.CreatePeeringPolicyResponse, error)
	UpdatePeeringPolicy(context.Context, *models.UpdatePeeringPolicyRequest) (*models.UpdatePeeringPolicyResponse, error)
	DeletePeeringPolicy(context.Context, *models.DeletePeeringPolicyRequest) (*models.DeletePeeringPolicyResponse, error)
	GetPeeringPolicy(context.Context, *models.GetPeeringPolicyRequest) (*models.GetPeeringPolicyResponse, error)
	ListPeeringPolicy(context.Context, *models.ListPeeringPolicyRequest) (*models.ListPeeringPolicyResponse, error)

	CreatePhysicalInterface(context.Context, *models.CreatePhysicalInterfaceRequest) (*models.CreatePhysicalInterfaceResponse, error)
	UpdatePhysicalInterface(context.Context, *models.UpdatePhysicalInterfaceRequest) (*models.UpdatePhysicalInterfaceResponse, error)
	DeletePhysicalInterface(context.Context, *models.DeletePhysicalInterfaceRequest) (*models.DeletePhysicalInterfaceResponse, error)
	GetPhysicalInterface(context.Context, *models.GetPhysicalInterfaceRequest) (*models.GetPhysicalInterfaceResponse, error)
	ListPhysicalInterface(context.Context, *models.ListPhysicalInterfaceRequest) (*models.ListPhysicalInterfaceResponse, error)

	CreatePhysicalRouter(context.Context, *models.CreatePhysicalRouterRequest) (*models.CreatePhysicalRouterResponse, error)
	UpdatePhysicalRouter(context.Context, *models.UpdatePhysicalRouterRequest) (*models.UpdatePhysicalRouterResponse, error)
	DeletePhysicalRouter(context.Context, *models.DeletePhysicalRouterRequest) (*models.DeletePhysicalRouterResponse, error)
	GetPhysicalRouter(context.Context, *models.GetPhysicalRouterRequest) (*models.GetPhysicalRouterResponse, error)
	ListPhysicalRouter(context.Context, *models.ListPhysicalRouterRequest) (*models.ListPhysicalRouterResponse, error)

	CreatePolicyManagement(context.Context, *models.CreatePolicyManagementRequest) (*models.CreatePolicyManagementResponse, error)
	UpdatePolicyManagement(context.Context, *models.UpdatePolicyManagementRequest) (*models.UpdatePolicyManagementResponse, error)
	DeletePolicyManagement(context.Context, *models.DeletePolicyManagementRequest) (*models.DeletePolicyManagementResponse, error)
	GetPolicyManagement(context.Context, *models.GetPolicyManagementRequest) (*models.GetPolicyManagementResponse, error)
	ListPolicyManagement(context.Context, *models.ListPolicyManagementRequest) (*models.ListPolicyManagementResponse, error)

	CreatePortTuple(context.Context, *models.CreatePortTupleRequest) (*models.CreatePortTupleResponse, error)
	UpdatePortTuple(context.Context, *models.UpdatePortTupleRequest) (*models.UpdatePortTupleResponse, error)
	DeletePortTuple(context.Context, *models.DeletePortTupleRequest) (*models.DeletePortTupleResponse, error)
	GetPortTuple(context.Context, *models.GetPortTupleRequest) (*models.GetPortTupleResponse, error)
	ListPortTuple(context.Context, *models.ListPortTupleRequest) (*models.ListPortTupleResponse, error)

	CreateProject(context.Context, *models.CreateProjectRequest) (*models.CreateProjectResponse, error)
	UpdateProject(context.Context, *models.UpdateProjectRequest) (*models.UpdateProjectResponse, error)
	DeleteProject(context.Context, *models.DeleteProjectRequest) (*models.DeleteProjectResponse, error)
	GetProject(context.Context, *models.GetProjectRequest) (*models.GetProjectResponse, error)
	ListProject(context.Context, *models.ListProjectRequest) (*models.ListProjectResponse, error)

	CreateProviderAttachment(context.Context, *models.CreateProviderAttachmentRequest) (*models.CreateProviderAttachmentResponse, error)
	UpdateProviderAttachment(context.Context, *models.UpdateProviderAttachmentRequest) (*models.UpdateProviderAttachmentResponse, error)
	DeleteProviderAttachment(context.Context, *models.DeleteProviderAttachmentRequest) (*models.DeleteProviderAttachmentResponse, error)
	GetProviderAttachment(context.Context, *models.GetProviderAttachmentRequest) (*models.GetProviderAttachmentResponse, error)
	ListProviderAttachment(context.Context, *models.ListProviderAttachmentRequest) (*models.ListProviderAttachmentResponse, error)

	CreateQosConfig(context.Context, *models.CreateQosConfigRequest) (*models.CreateQosConfigResponse, error)
	UpdateQosConfig(context.Context, *models.UpdateQosConfigRequest) (*models.UpdateQosConfigResponse, error)
	DeleteQosConfig(context.Context, *models.DeleteQosConfigRequest) (*models.DeleteQosConfigResponse, error)
	GetQosConfig(context.Context, *models.GetQosConfigRequest) (*models.GetQosConfigResponse, error)
	ListQosConfig(context.Context, *models.ListQosConfigRequest) (*models.ListQosConfigResponse, error)

	CreateQosQueue(context.Context, *models.CreateQosQueueRequest) (*models.CreateQosQueueResponse, error)
	UpdateQosQueue(context.Context, *models.UpdateQosQueueRequest) (*models.UpdateQosQueueResponse, error)
	DeleteQosQueue(context.Context, *models.DeleteQosQueueRequest) (*models.DeleteQosQueueResponse, error)
	GetQosQueue(context.Context, *models.GetQosQueueRequest) (*models.GetQosQueueResponse, error)
	ListQosQueue(context.Context, *models.ListQosQueueRequest) (*models.ListQosQueueResponse, error)

	CreateRouteAggregate(context.Context, *models.CreateRouteAggregateRequest) (*models.CreateRouteAggregateResponse, error)
	UpdateRouteAggregate(context.Context, *models.UpdateRouteAggregateRequest) (*models.UpdateRouteAggregateResponse, error)
	DeleteRouteAggregate(context.Context, *models.DeleteRouteAggregateRequest) (*models.DeleteRouteAggregateResponse, error)
	GetRouteAggregate(context.Context, *models.GetRouteAggregateRequest) (*models.GetRouteAggregateResponse, error)
	ListRouteAggregate(context.Context, *models.ListRouteAggregateRequest) (*models.ListRouteAggregateResponse, error)

	CreateRouteTable(context.Context, *models.CreateRouteTableRequest) (*models.CreateRouteTableResponse, error)
	UpdateRouteTable(context.Context, *models.UpdateRouteTableRequest) (*models.UpdateRouteTableResponse, error)
	DeleteRouteTable(context.Context, *models.DeleteRouteTableRequest) (*models.DeleteRouteTableResponse, error)
	GetRouteTable(context.Context, *models.GetRouteTableRequest) (*models.GetRouteTableResponse, error)
	ListRouteTable(context.Context, *models.ListRouteTableRequest) (*models.ListRouteTableResponse, error)

	CreateRouteTarget(context.Context, *models.CreateRouteTargetRequest) (*models.CreateRouteTargetResponse, error)
	UpdateRouteTarget(context.Context, *models.UpdateRouteTargetRequest) (*models.UpdateRouteTargetResponse, error)
	DeleteRouteTarget(context.Context, *models.DeleteRouteTargetRequest) (*models.DeleteRouteTargetResponse, error)
	GetRouteTarget(context.Context, *models.GetRouteTargetRequest) (*models.GetRouteTargetResponse, error)
	ListRouteTarget(context.Context, *models.ListRouteTargetRequest) (*models.ListRouteTargetResponse, error)

	CreateRoutingInstance(context.Context, *models.CreateRoutingInstanceRequest) (*models.CreateRoutingInstanceResponse, error)
	UpdateRoutingInstance(context.Context, *models.UpdateRoutingInstanceRequest) (*models.UpdateRoutingInstanceResponse, error)
	DeleteRoutingInstance(context.Context, *models.DeleteRoutingInstanceRequest) (*models.DeleteRoutingInstanceResponse, error)
	GetRoutingInstance(context.Context, *models.GetRoutingInstanceRequest) (*models.GetRoutingInstanceResponse, error)
	ListRoutingInstance(context.Context, *models.ListRoutingInstanceRequest) (*models.ListRoutingInstanceResponse, error)

	CreateRoutingPolicy(context.Context, *models.CreateRoutingPolicyRequest) (*models.CreateRoutingPolicyResponse, error)
	UpdateRoutingPolicy(context.Context, *models.UpdateRoutingPolicyRequest) (*models.UpdateRoutingPolicyResponse, error)
	DeleteRoutingPolicy(context.Context, *models.DeleteRoutingPolicyRequest) (*models.DeleteRoutingPolicyResponse, error)
	GetRoutingPolicy(context.Context, *models.GetRoutingPolicyRequest) (*models.GetRoutingPolicyResponse, error)
	ListRoutingPolicy(context.Context, *models.ListRoutingPolicyRequest) (*models.ListRoutingPolicyResponse, error)

	CreateSecurityGroup(context.Context, *models.CreateSecurityGroupRequest) (*models.CreateSecurityGroupResponse, error)
	UpdateSecurityGroup(context.Context, *models.UpdateSecurityGroupRequest) (*models.UpdateSecurityGroupResponse, error)
	DeleteSecurityGroup(context.Context, *models.DeleteSecurityGroupRequest) (*models.DeleteSecurityGroupResponse, error)
	GetSecurityGroup(context.Context, *models.GetSecurityGroupRequest) (*models.GetSecurityGroupResponse, error)
	ListSecurityGroup(context.Context, *models.ListSecurityGroupRequest) (*models.ListSecurityGroupResponse, error)

	CreateSecurityLoggingObject(context.Context, *models.CreateSecurityLoggingObjectRequest) (*models.CreateSecurityLoggingObjectResponse, error)
	UpdateSecurityLoggingObject(context.Context, *models.UpdateSecurityLoggingObjectRequest) (*models.UpdateSecurityLoggingObjectResponse, error)
	DeleteSecurityLoggingObject(context.Context, *models.DeleteSecurityLoggingObjectRequest) (*models.DeleteSecurityLoggingObjectResponse, error)
	GetSecurityLoggingObject(context.Context, *models.GetSecurityLoggingObjectRequest) (*models.GetSecurityLoggingObjectResponse, error)
	ListSecurityLoggingObject(context.Context, *models.ListSecurityLoggingObjectRequest) (*models.ListSecurityLoggingObjectResponse, error)

	CreateServiceAppliance(context.Context, *models.CreateServiceApplianceRequest) (*models.CreateServiceApplianceResponse, error)
	UpdateServiceAppliance(context.Context, *models.UpdateServiceApplianceRequest) (*models.UpdateServiceApplianceResponse, error)
	DeleteServiceAppliance(context.Context, *models.DeleteServiceApplianceRequest) (*models.DeleteServiceApplianceResponse, error)
	GetServiceAppliance(context.Context, *models.GetServiceApplianceRequest) (*models.GetServiceApplianceResponse, error)
	ListServiceAppliance(context.Context, *models.ListServiceApplianceRequest) (*models.ListServiceApplianceResponse, error)

	CreateServiceApplianceSet(context.Context, *models.CreateServiceApplianceSetRequest) (*models.CreateServiceApplianceSetResponse, error)
	UpdateServiceApplianceSet(context.Context, *models.UpdateServiceApplianceSetRequest) (*models.UpdateServiceApplianceSetResponse, error)
	DeleteServiceApplianceSet(context.Context, *models.DeleteServiceApplianceSetRequest) (*models.DeleteServiceApplianceSetResponse, error)
	GetServiceApplianceSet(context.Context, *models.GetServiceApplianceSetRequest) (*models.GetServiceApplianceSetResponse, error)
	ListServiceApplianceSet(context.Context, *models.ListServiceApplianceSetRequest) (*models.ListServiceApplianceSetResponse, error)

	CreateServiceConnectionModule(context.Context, *models.CreateServiceConnectionModuleRequest) (*models.CreateServiceConnectionModuleResponse, error)
	UpdateServiceConnectionModule(context.Context, *models.UpdateServiceConnectionModuleRequest) (*models.UpdateServiceConnectionModuleResponse, error)
	DeleteServiceConnectionModule(context.Context, *models.DeleteServiceConnectionModuleRequest) (*models.DeleteServiceConnectionModuleResponse, error)
	GetServiceConnectionModule(context.Context, *models.GetServiceConnectionModuleRequest) (*models.GetServiceConnectionModuleResponse, error)
	ListServiceConnectionModule(context.Context, *models.ListServiceConnectionModuleRequest) (*models.ListServiceConnectionModuleResponse, error)

	CreateServiceEndpoint(context.Context, *models.CreateServiceEndpointRequest) (*models.CreateServiceEndpointResponse, error)
	UpdateServiceEndpoint(context.Context, *models.UpdateServiceEndpointRequest) (*models.UpdateServiceEndpointResponse, error)
	DeleteServiceEndpoint(context.Context, *models.DeleteServiceEndpointRequest) (*models.DeleteServiceEndpointResponse, error)
	GetServiceEndpoint(context.Context, *models.GetServiceEndpointRequest) (*models.GetServiceEndpointResponse, error)
	ListServiceEndpoint(context.Context, *models.ListServiceEndpointRequest) (*models.ListServiceEndpointResponse, error)

	CreateServiceGroup(context.Context, *models.CreateServiceGroupRequest) (*models.CreateServiceGroupResponse, error)
	UpdateServiceGroup(context.Context, *models.UpdateServiceGroupRequest) (*models.UpdateServiceGroupResponse, error)
	DeleteServiceGroup(context.Context, *models.DeleteServiceGroupRequest) (*models.DeleteServiceGroupResponse, error)
	GetServiceGroup(context.Context, *models.GetServiceGroupRequest) (*models.GetServiceGroupResponse, error)
	ListServiceGroup(context.Context, *models.ListServiceGroupRequest) (*models.ListServiceGroupResponse, error)

	CreateServiceHealthCheck(context.Context, *models.CreateServiceHealthCheckRequest) (*models.CreateServiceHealthCheckResponse, error)
	UpdateServiceHealthCheck(context.Context, *models.UpdateServiceHealthCheckRequest) (*models.UpdateServiceHealthCheckResponse, error)
	DeleteServiceHealthCheck(context.Context, *models.DeleteServiceHealthCheckRequest) (*models.DeleteServiceHealthCheckResponse, error)
	GetServiceHealthCheck(context.Context, *models.GetServiceHealthCheckRequest) (*models.GetServiceHealthCheckResponse, error)
	ListServiceHealthCheck(context.Context, *models.ListServiceHealthCheckRequest) (*models.ListServiceHealthCheckResponse, error)

	CreateServiceInstance(context.Context, *models.CreateServiceInstanceRequest) (*models.CreateServiceInstanceResponse, error)
	UpdateServiceInstance(context.Context, *models.UpdateServiceInstanceRequest) (*models.UpdateServiceInstanceResponse, error)
	DeleteServiceInstance(context.Context, *models.DeleteServiceInstanceRequest) (*models.DeleteServiceInstanceResponse, error)
	GetServiceInstance(context.Context, *models.GetServiceInstanceRequest) (*models.GetServiceInstanceResponse, error)
	ListServiceInstance(context.Context, *models.ListServiceInstanceRequest) (*models.ListServiceInstanceResponse, error)

	CreateServiceObject(context.Context, *models.CreateServiceObjectRequest) (*models.CreateServiceObjectResponse, error)
	UpdateServiceObject(context.Context, *models.UpdateServiceObjectRequest) (*models.UpdateServiceObjectResponse, error)
	DeleteServiceObject(context.Context, *models.DeleteServiceObjectRequest) (*models.DeleteServiceObjectResponse, error)
	GetServiceObject(context.Context, *models.GetServiceObjectRequest) (*models.GetServiceObjectResponse, error)
	ListServiceObject(context.Context, *models.ListServiceObjectRequest) (*models.ListServiceObjectResponse, error)

	CreateServiceTemplate(context.Context, *models.CreateServiceTemplateRequest) (*models.CreateServiceTemplateResponse, error)
	UpdateServiceTemplate(context.Context, *models.UpdateServiceTemplateRequest) (*models.UpdateServiceTemplateResponse, error)
	DeleteServiceTemplate(context.Context, *models.DeleteServiceTemplateRequest) (*models.DeleteServiceTemplateResponse, error)
	GetServiceTemplate(context.Context, *models.GetServiceTemplateRequest) (*models.GetServiceTemplateResponse, error)
	ListServiceTemplate(context.Context, *models.ListServiceTemplateRequest) (*models.ListServiceTemplateResponse, error)

	CreateSubnet(context.Context, *models.CreateSubnetRequest) (*models.CreateSubnetResponse, error)
	UpdateSubnet(context.Context, *models.UpdateSubnetRequest) (*models.UpdateSubnetResponse, error)
	DeleteSubnet(context.Context, *models.DeleteSubnetRequest) (*models.DeleteSubnetResponse, error)
	GetSubnet(context.Context, *models.GetSubnetRequest) (*models.GetSubnetResponse, error)
	ListSubnet(context.Context, *models.ListSubnetRequest) (*models.ListSubnetResponse, error)

	CreateTag(context.Context, *models.CreateTagRequest) (*models.CreateTagResponse, error)
	UpdateTag(context.Context, *models.UpdateTagRequest) (*models.UpdateTagResponse, error)
	DeleteTag(context.Context, *models.DeleteTagRequest) (*models.DeleteTagResponse, error)
	GetTag(context.Context, *models.GetTagRequest) (*models.GetTagResponse, error)
	ListTag(context.Context, *models.ListTagRequest) (*models.ListTagResponse, error)

	CreateTagType(context.Context, *models.CreateTagTypeRequest) (*models.CreateTagTypeResponse, error)
	UpdateTagType(context.Context, *models.UpdateTagTypeRequest) (*models.UpdateTagTypeResponse, error)
	DeleteTagType(context.Context, *models.DeleteTagTypeRequest) (*models.DeleteTagTypeResponse, error)
	GetTagType(context.Context, *models.GetTagTypeRequest) (*models.GetTagTypeResponse, error)
	ListTagType(context.Context, *models.ListTagTypeRequest) (*models.ListTagTypeResponse, error)

	CreateUser(context.Context, *models.CreateUserRequest) (*models.CreateUserResponse, error)
	UpdateUser(context.Context, *models.UpdateUserRequest) (*models.UpdateUserResponse, error)
	DeleteUser(context.Context, *models.DeleteUserRequest) (*models.DeleteUserResponse, error)
	GetUser(context.Context, *models.GetUserRequest) (*models.GetUserResponse, error)
	ListUser(context.Context, *models.ListUserRequest) (*models.ListUserResponse, error)

	CreateVirtualDNSRecord(context.Context, *models.CreateVirtualDNSRecordRequest) (*models.CreateVirtualDNSRecordResponse, error)
	UpdateVirtualDNSRecord(context.Context, *models.UpdateVirtualDNSRecordRequest) (*models.UpdateVirtualDNSRecordResponse, error)
	DeleteVirtualDNSRecord(context.Context, *models.DeleteVirtualDNSRecordRequest) (*models.DeleteVirtualDNSRecordResponse, error)
	GetVirtualDNSRecord(context.Context, *models.GetVirtualDNSRecordRequest) (*models.GetVirtualDNSRecordResponse, error)
	ListVirtualDNSRecord(context.Context, *models.ListVirtualDNSRecordRequest) (*models.ListVirtualDNSRecordResponse, error)

	CreateVirtualDNS(context.Context, *models.CreateVirtualDNSRequest) (*models.CreateVirtualDNSResponse, error)
	UpdateVirtualDNS(context.Context, *models.UpdateVirtualDNSRequest) (*models.UpdateVirtualDNSResponse, error)
	DeleteVirtualDNS(context.Context, *models.DeleteVirtualDNSRequest) (*models.DeleteVirtualDNSResponse, error)
	GetVirtualDNS(context.Context, *models.GetVirtualDNSRequest) (*models.GetVirtualDNSResponse, error)
	ListVirtualDNS(context.Context, *models.ListVirtualDNSRequest) (*models.ListVirtualDNSResponse, error)

	CreateVirtualIP(context.Context, *models.CreateVirtualIPRequest) (*models.CreateVirtualIPResponse, error)
	UpdateVirtualIP(context.Context, *models.UpdateVirtualIPRequest) (*models.UpdateVirtualIPResponse, error)
	DeleteVirtualIP(context.Context, *models.DeleteVirtualIPRequest) (*models.DeleteVirtualIPResponse, error)
	GetVirtualIP(context.Context, *models.GetVirtualIPRequest) (*models.GetVirtualIPResponse, error)
	ListVirtualIP(context.Context, *models.ListVirtualIPRequest) (*models.ListVirtualIPResponse, error)

	CreateVirtualMachineInterface(context.Context, *models.CreateVirtualMachineInterfaceRequest) (*models.CreateVirtualMachineInterfaceResponse, error)
	UpdateVirtualMachineInterface(context.Context, *models.UpdateVirtualMachineInterfaceRequest) (*models.UpdateVirtualMachineInterfaceResponse, error)
	DeleteVirtualMachineInterface(context.Context, *models.DeleteVirtualMachineInterfaceRequest) (*models.DeleteVirtualMachineInterfaceResponse, error)
	GetVirtualMachineInterface(context.Context, *models.GetVirtualMachineInterfaceRequest) (*models.GetVirtualMachineInterfaceResponse, error)
	ListVirtualMachineInterface(context.Context, *models.ListVirtualMachineInterfaceRequest) (*models.ListVirtualMachineInterfaceResponse, error)

	CreateVirtualMachine(context.Context, *models.CreateVirtualMachineRequest) (*models.CreateVirtualMachineResponse, error)
	UpdateVirtualMachine(context.Context, *models.UpdateVirtualMachineRequest) (*models.UpdateVirtualMachineResponse, error)
	DeleteVirtualMachine(context.Context, *models.DeleteVirtualMachineRequest) (*models.DeleteVirtualMachineResponse, error)
	GetVirtualMachine(context.Context, *models.GetVirtualMachineRequest) (*models.GetVirtualMachineResponse, error)
	ListVirtualMachine(context.Context, *models.ListVirtualMachineRequest) (*models.ListVirtualMachineResponse, error)

	CreateVirtualNetwork(context.Context, *models.CreateVirtualNetworkRequest) (*models.CreateVirtualNetworkResponse, error)
	UpdateVirtualNetwork(context.Context, *models.UpdateVirtualNetworkRequest) (*models.UpdateVirtualNetworkResponse, error)
	DeleteVirtualNetwork(context.Context, *models.DeleteVirtualNetworkRequest) (*models.DeleteVirtualNetworkResponse, error)
	GetVirtualNetwork(context.Context, *models.GetVirtualNetworkRequest) (*models.GetVirtualNetworkResponse, error)
	ListVirtualNetwork(context.Context, *models.ListVirtualNetworkRequest) (*models.ListVirtualNetworkResponse, error)

	CreateVirtualRouter(context.Context, *models.CreateVirtualRouterRequest) (*models.CreateVirtualRouterResponse, error)
	UpdateVirtualRouter(context.Context, *models.UpdateVirtualRouterRequest) (*models.UpdateVirtualRouterResponse, error)
	DeleteVirtualRouter(context.Context, *models.DeleteVirtualRouterRequest) (*models.DeleteVirtualRouterResponse, error)
	GetVirtualRouter(context.Context, *models.GetVirtualRouterRequest) (*models.GetVirtualRouterResponse, error)
	ListVirtualRouter(context.Context, *models.ListVirtualRouterRequest) (*models.ListVirtualRouterResponse, error)

	CreateAppformixNode(context.Context, *models.CreateAppformixNodeRequest) (*models.CreateAppformixNodeResponse, error)
	UpdateAppformixNode(context.Context, *models.UpdateAppformixNodeRequest) (*models.UpdateAppformixNodeResponse, error)
	DeleteAppformixNode(context.Context, *models.DeleteAppformixNodeRequest) (*models.DeleteAppformixNodeResponse, error)
	GetAppformixNode(context.Context, *models.GetAppformixNodeRequest) (*models.GetAppformixNodeResponse, error)
	ListAppformixNode(context.Context, *models.ListAppformixNodeRequest) (*models.ListAppformixNodeResponse, error)

	CreateBaremetalNode(context.Context, *models.CreateBaremetalNodeRequest) (*models.CreateBaremetalNodeResponse, error)
	UpdateBaremetalNode(context.Context, *models.UpdateBaremetalNodeRequest) (*models.UpdateBaremetalNodeResponse, error)
	DeleteBaremetalNode(context.Context, *models.DeleteBaremetalNodeRequest) (*models.DeleteBaremetalNodeResponse, error)
	GetBaremetalNode(context.Context, *models.GetBaremetalNodeRequest) (*models.GetBaremetalNodeResponse, error)
	ListBaremetalNode(context.Context, *models.ListBaremetalNodeRequest) (*models.ListBaremetalNodeResponse, error)

	CreateBaremetalPort(context.Context, *models.CreateBaremetalPortRequest) (*models.CreateBaremetalPortResponse, error)
	UpdateBaremetalPort(context.Context, *models.UpdateBaremetalPortRequest) (*models.UpdateBaremetalPortResponse, error)
	DeleteBaremetalPort(context.Context, *models.DeleteBaremetalPortRequest) (*models.DeleteBaremetalPortResponse, error)
	GetBaremetalPort(context.Context, *models.GetBaremetalPortRequest) (*models.GetBaremetalPortResponse, error)
	ListBaremetalPort(context.Context, *models.ListBaremetalPortRequest) (*models.ListBaremetalPortResponse, error)

	CreateContrailAnalyticsDatabaseNode(context.Context, *models.CreateContrailAnalyticsDatabaseNodeRequest) (*models.CreateContrailAnalyticsDatabaseNodeResponse, error)
	UpdateContrailAnalyticsDatabaseNode(context.Context, *models.UpdateContrailAnalyticsDatabaseNodeRequest) (*models.UpdateContrailAnalyticsDatabaseNodeResponse, error)
	DeleteContrailAnalyticsDatabaseNode(context.Context, *models.DeleteContrailAnalyticsDatabaseNodeRequest) (*models.DeleteContrailAnalyticsDatabaseNodeResponse, error)
	GetContrailAnalyticsDatabaseNode(context.Context, *models.GetContrailAnalyticsDatabaseNodeRequest) (*models.GetContrailAnalyticsDatabaseNodeResponse, error)
	ListContrailAnalyticsDatabaseNode(context.Context, *models.ListContrailAnalyticsDatabaseNodeRequest) (*models.ListContrailAnalyticsDatabaseNodeResponse, error)

	CreateContrailAnalyticsNode(context.Context, *models.CreateContrailAnalyticsNodeRequest) (*models.CreateContrailAnalyticsNodeResponse, error)
	UpdateContrailAnalyticsNode(context.Context, *models.UpdateContrailAnalyticsNodeRequest) (*models.UpdateContrailAnalyticsNodeResponse, error)
	DeleteContrailAnalyticsNode(context.Context, *models.DeleteContrailAnalyticsNodeRequest) (*models.DeleteContrailAnalyticsNodeResponse, error)
	GetContrailAnalyticsNode(context.Context, *models.GetContrailAnalyticsNodeRequest) (*models.GetContrailAnalyticsNodeResponse, error)
	ListContrailAnalyticsNode(context.Context, *models.ListContrailAnalyticsNodeRequest) (*models.ListContrailAnalyticsNodeResponse, error)

	CreateContrailCluster(context.Context, *models.CreateContrailClusterRequest) (*models.CreateContrailClusterResponse, error)
	UpdateContrailCluster(context.Context, *models.UpdateContrailClusterRequest) (*models.UpdateContrailClusterResponse, error)
	DeleteContrailCluster(context.Context, *models.DeleteContrailClusterRequest) (*models.DeleteContrailClusterResponse, error)
	GetContrailCluster(context.Context, *models.GetContrailClusterRequest) (*models.GetContrailClusterResponse, error)
	ListContrailCluster(context.Context, *models.ListContrailClusterRequest) (*models.ListContrailClusterResponse, error)

	CreateContrailConfigDatabaseNode(context.Context, *models.CreateContrailConfigDatabaseNodeRequest) (*models.CreateContrailConfigDatabaseNodeResponse, error)
	UpdateContrailConfigDatabaseNode(context.Context, *models.UpdateContrailConfigDatabaseNodeRequest) (*models.UpdateContrailConfigDatabaseNodeResponse, error)
	DeleteContrailConfigDatabaseNode(context.Context, *models.DeleteContrailConfigDatabaseNodeRequest) (*models.DeleteContrailConfigDatabaseNodeResponse, error)
	GetContrailConfigDatabaseNode(context.Context, *models.GetContrailConfigDatabaseNodeRequest) (*models.GetContrailConfigDatabaseNodeResponse, error)
	ListContrailConfigDatabaseNode(context.Context, *models.ListContrailConfigDatabaseNodeRequest) (*models.ListContrailConfigDatabaseNodeResponse, error)

	CreateContrailConfigNode(context.Context, *models.CreateContrailConfigNodeRequest) (*models.CreateContrailConfigNodeResponse, error)
	UpdateContrailConfigNode(context.Context, *models.UpdateContrailConfigNodeRequest) (*models.UpdateContrailConfigNodeResponse, error)
	DeleteContrailConfigNode(context.Context, *models.DeleteContrailConfigNodeRequest) (*models.DeleteContrailConfigNodeResponse, error)
	GetContrailConfigNode(context.Context, *models.GetContrailConfigNodeRequest) (*models.GetContrailConfigNodeResponse, error)
	ListContrailConfigNode(context.Context, *models.ListContrailConfigNodeRequest) (*models.ListContrailConfigNodeResponse, error)

	CreateContrailControlNode(context.Context, *models.CreateContrailControlNodeRequest) (*models.CreateContrailControlNodeResponse, error)
	UpdateContrailControlNode(context.Context, *models.UpdateContrailControlNodeRequest) (*models.UpdateContrailControlNodeResponse, error)
	DeleteContrailControlNode(context.Context, *models.DeleteContrailControlNodeRequest) (*models.DeleteContrailControlNodeResponse, error)
	GetContrailControlNode(context.Context, *models.GetContrailControlNodeRequest) (*models.GetContrailControlNodeResponse, error)
	ListContrailControlNode(context.Context, *models.ListContrailControlNodeRequest) (*models.ListContrailControlNodeResponse, error)

	CreateContrailStorageNode(context.Context, *models.CreateContrailStorageNodeRequest) (*models.CreateContrailStorageNodeResponse, error)
	UpdateContrailStorageNode(context.Context, *models.UpdateContrailStorageNodeRequest) (*models.UpdateContrailStorageNodeResponse, error)
	DeleteContrailStorageNode(context.Context, *models.DeleteContrailStorageNodeRequest) (*models.DeleteContrailStorageNodeResponse, error)
	GetContrailStorageNode(context.Context, *models.GetContrailStorageNodeRequest) (*models.GetContrailStorageNodeResponse, error)
	ListContrailStorageNode(context.Context, *models.ListContrailStorageNodeRequest) (*models.ListContrailStorageNodeResponse, error)

	CreateContrailVrouterNode(context.Context, *models.CreateContrailVrouterNodeRequest) (*models.CreateContrailVrouterNodeResponse, error)
	UpdateContrailVrouterNode(context.Context, *models.UpdateContrailVrouterNodeRequest) (*models.UpdateContrailVrouterNodeResponse, error)
	DeleteContrailVrouterNode(context.Context, *models.DeleteContrailVrouterNodeRequest) (*models.DeleteContrailVrouterNodeResponse, error)
	GetContrailVrouterNode(context.Context, *models.GetContrailVrouterNodeRequest) (*models.GetContrailVrouterNodeResponse, error)
	ListContrailVrouterNode(context.Context, *models.ListContrailVrouterNodeRequest) (*models.ListContrailVrouterNodeResponse, error)

	CreateContrailControllerNode(context.Context, *models.CreateContrailControllerNodeRequest) (*models.CreateContrailControllerNodeResponse, error)
	UpdateContrailControllerNode(context.Context, *models.UpdateContrailControllerNodeRequest) (*models.UpdateContrailControllerNodeResponse, error)
	DeleteContrailControllerNode(context.Context, *models.DeleteContrailControllerNodeRequest) (*models.DeleteContrailControllerNodeResponse, error)
	GetContrailControllerNode(context.Context, *models.GetContrailControllerNodeRequest) (*models.GetContrailControllerNodeResponse, error)
	ListContrailControllerNode(context.Context, *models.ListContrailControllerNodeRequest) (*models.ListContrailControllerNodeResponse, error)

	CreateDashboard(context.Context, *models.CreateDashboardRequest) (*models.CreateDashboardResponse, error)
	UpdateDashboard(context.Context, *models.UpdateDashboardRequest) (*models.UpdateDashboardResponse, error)
	DeleteDashboard(context.Context, *models.DeleteDashboardRequest) (*models.DeleteDashboardResponse, error)
	GetDashboard(context.Context, *models.GetDashboardRequest) (*models.GetDashboardResponse, error)
	ListDashboard(context.Context, *models.ListDashboardRequest) (*models.ListDashboardResponse, error)

	CreateFlavor(context.Context, *models.CreateFlavorRequest) (*models.CreateFlavorResponse, error)
	UpdateFlavor(context.Context, *models.UpdateFlavorRequest) (*models.UpdateFlavorResponse, error)
	DeleteFlavor(context.Context, *models.DeleteFlavorRequest) (*models.DeleteFlavorResponse, error)
	GetFlavor(context.Context, *models.GetFlavorRequest) (*models.GetFlavorResponse, error)
	ListFlavor(context.Context, *models.ListFlavorRequest) (*models.ListFlavorResponse, error)

	CreateOsImage(context.Context, *models.CreateOsImageRequest) (*models.CreateOsImageResponse, error)
	UpdateOsImage(context.Context, *models.UpdateOsImageRequest) (*models.UpdateOsImageResponse, error)
	DeleteOsImage(context.Context, *models.DeleteOsImageRequest) (*models.DeleteOsImageResponse, error)
	GetOsImage(context.Context, *models.GetOsImageRequest) (*models.GetOsImageResponse, error)
	ListOsImage(context.Context, *models.ListOsImageRequest) (*models.ListOsImageResponse, error)

	CreateKeypair(context.Context, *models.CreateKeypairRequest) (*models.CreateKeypairResponse, error)
	UpdateKeypair(context.Context, *models.UpdateKeypairRequest) (*models.UpdateKeypairResponse, error)
	DeleteKeypair(context.Context, *models.DeleteKeypairRequest) (*models.DeleteKeypairResponse, error)
	GetKeypair(context.Context, *models.GetKeypairRequest) (*models.GetKeypairResponse, error)
	ListKeypair(context.Context, *models.ListKeypairRequest) (*models.ListKeypairResponse, error)

	CreateKubernetesMasterNode(context.Context, *models.CreateKubernetesMasterNodeRequest) (*models.CreateKubernetesMasterNodeResponse, error)
	UpdateKubernetesMasterNode(context.Context, *models.UpdateKubernetesMasterNodeRequest) (*models.UpdateKubernetesMasterNodeResponse, error)
	DeleteKubernetesMasterNode(context.Context, *models.DeleteKubernetesMasterNodeRequest) (*models.DeleteKubernetesMasterNodeResponse, error)
	GetKubernetesMasterNode(context.Context, *models.GetKubernetesMasterNodeRequest) (*models.GetKubernetesMasterNodeResponse, error)
	ListKubernetesMasterNode(context.Context, *models.ListKubernetesMasterNodeRequest) (*models.ListKubernetesMasterNodeResponse, error)

	CreateKubernetesNode(context.Context, *models.CreateKubernetesNodeRequest) (*models.CreateKubernetesNodeResponse, error)
	UpdateKubernetesNode(context.Context, *models.UpdateKubernetesNodeRequest) (*models.UpdateKubernetesNodeResponse, error)
	DeleteKubernetesNode(context.Context, *models.DeleteKubernetesNodeRequest) (*models.DeleteKubernetesNodeResponse, error)
	GetKubernetesNode(context.Context, *models.GetKubernetesNodeRequest) (*models.GetKubernetesNodeResponse, error)
	ListKubernetesNode(context.Context, *models.ListKubernetesNodeRequest) (*models.ListKubernetesNodeResponse, error)

	CreateLocation(context.Context, *models.CreateLocationRequest) (*models.CreateLocationResponse, error)
	UpdateLocation(context.Context, *models.UpdateLocationRequest) (*models.UpdateLocationResponse, error)
	DeleteLocation(context.Context, *models.DeleteLocationRequest) (*models.DeleteLocationResponse, error)
	GetLocation(context.Context, *models.GetLocationRequest) (*models.GetLocationResponse, error)
	ListLocation(context.Context, *models.ListLocationRequest) (*models.ListLocationResponse, error)

	CreateNode(context.Context, *models.CreateNodeRequest) (*models.CreateNodeResponse, error)
	UpdateNode(context.Context, *models.UpdateNodeRequest) (*models.UpdateNodeResponse, error)
	DeleteNode(context.Context, *models.DeleteNodeRequest) (*models.DeleteNodeResponse, error)
	GetNode(context.Context, *models.GetNodeRequest) (*models.GetNodeResponse, error)
	ListNode(context.Context, *models.ListNodeRequest) (*models.ListNodeResponse, error)

	CreateServer(context.Context, *models.CreateServerRequest) (*models.CreateServerResponse, error)
	UpdateServer(context.Context, *models.UpdateServerRequest) (*models.UpdateServerResponse, error)
	DeleteServer(context.Context, *models.DeleteServerRequest) (*models.DeleteServerResponse, error)
	GetServer(context.Context, *models.GetServerRequest) (*models.GetServerResponse, error)
	ListServer(context.Context, *models.ListServerRequest) (*models.ListServerResponse, error)

	CreateVPNGroup(context.Context, *models.CreateVPNGroupRequest) (*models.CreateVPNGroupResponse, error)
	UpdateVPNGroup(context.Context, *models.UpdateVPNGroupRequest) (*models.UpdateVPNGroupResponse, error)
	DeleteVPNGroup(context.Context, *models.DeleteVPNGroupRequest) (*models.DeleteVPNGroupResponse, error)
	GetVPNGroup(context.Context, *models.GetVPNGroupRequest) (*models.GetVPNGroupResponse, error)
	ListVPNGroup(context.Context, *models.ListVPNGroupRequest) (*models.ListVPNGroupResponse, error)

	CreateWidget(context.Context, *models.CreateWidgetRequest) (*models.CreateWidgetResponse, error)
	UpdateWidget(context.Context, *models.UpdateWidgetRequest) (*models.UpdateWidgetResponse, error)
	DeleteWidget(context.Context, *models.DeleteWidgetRequest) (*models.DeleteWidgetResponse, error)
	GetWidget(context.Context, *models.GetWidgetRequest) (*models.GetWidgetResponse, error)
	ListWidget(context.Context, *models.ListWidgetRequest) (*models.ListWidgetResponse, error)
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

func (service *BaseService) CreateAccessControlList(ctx context.Context, request *models.CreateAccessControlListRequest) (*models.CreateAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAccessControlList(ctx, request)
}
func (service *BaseService) UpdateAccessControlList(ctx context.Context, request *models.UpdateAccessControlListRequest) (*models.UpdateAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAccessControlList(ctx, request)
}
func (service *BaseService) DeleteAccessControlList(ctx context.Context, request *models.DeleteAccessControlListRequest) (*models.DeleteAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAccessControlList(ctx, request)
}
func (service *BaseService) GetAccessControlList(ctx context.Context, request *models.GetAccessControlListRequest) (*models.GetAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAccessControlList(ctx, request)
}
func (service *BaseService) ListAccessControlList(ctx context.Context, request *models.ListAccessControlListRequest) (*models.ListAccessControlListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAccessControlList(ctx, request)
}

func (service *BaseService) CreateAddressGroup(ctx context.Context, request *models.CreateAddressGroupRequest) (*models.CreateAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAddressGroup(ctx, request)
}
func (service *BaseService) UpdateAddressGroup(ctx context.Context, request *models.UpdateAddressGroupRequest) (*models.UpdateAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAddressGroup(ctx, request)
}
func (service *BaseService) DeleteAddressGroup(ctx context.Context, request *models.DeleteAddressGroupRequest) (*models.DeleteAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAddressGroup(ctx, request)
}
func (service *BaseService) GetAddressGroup(ctx context.Context, request *models.GetAddressGroupRequest) (*models.GetAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAddressGroup(ctx, request)
}
func (service *BaseService) ListAddressGroup(ctx context.Context, request *models.ListAddressGroupRequest) (*models.ListAddressGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAddressGroup(ctx, request)
}

func (service *BaseService) CreateAlarm(ctx context.Context, request *models.CreateAlarmRequest) (*models.CreateAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAlarm(ctx, request)
}
func (service *BaseService) UpdateAlarm(ctx context.Context, request *models.UpdateAlarmRequest) (*models.UpdateAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAlarm(ctx, request)
}
func (service *BaseService) DeleteAlarm(ctx context.Context, request *models.DeleteAlarmRequest) (*models.DeleteAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAlarm(ctx, request)
}
func (service *BaseService) GetAlarm(ctx context.Context, request *models.GetAlarmRequest) (*models.GetAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAlarm(ctx, request)
}
func (service *BaseService) ListAlarm(ctx context.Context, request *models.ListAlarmRequest) (*models.ListAlarmResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAlarm(ctx, request)
}

func (service *BaseService) CreateAliasIPPool(ctx context.Context, request *models.CreateAliasIPPoolRequest) (*models.CreateAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAliasIPPool(ctx, request)
}
func (service *BaseService) UpdateAliasIPPool(ctx context.Context, request *models.UpdateAliasIPPoolRequest) (*models.UpdateAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAliasIPPool(ctx, request)
}
func (service *BaseService) DeleteAliasIPPool(ctx context.Context, request *models.DeleteAliasIPPoolRequest) (*models.DeleteAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAliasIPPool(ctx, request)
}
func (service *BaseService) GetAliasIPPool(ctx context.Context, request *models.GetAliasIPPoolRequest) (*models.GetAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAliasIPPool(ctx, request)
}
func (service *BaseService) ListAliasIPPool(ctx context.Context, request *models.ListAliasIPPoolRequest) (*models.ListAliasIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAliasIPPool(ctx, request)
}

func (service *BaseService) CreateAliasIP(ctx context.Context, request *models.CreateAliasIPRequest) (*models.CreateAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAliasIP(ctx, request)
}
func (service *BaseService) UpdateAliasIP(ctx context.Context, request *models.UpdateAliasIPRequest) (*models.UpdateAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAliasIP(ctx, request)
}
func (service *BaseService) DeleteAliasIP(ctx context.Context, request *models.DeleteAliasIPRequest) (*models.DeleteAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAliasIP(ctx, request)
}
func (service *BaseService) GetAliasIP(ctx context.Context, request *models.GetAliasIPRequest) (*models.GetAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAliasIP(ctx, request)
}
func (service *BaseService) ListAliasIP(ctx context.Context, request *models.ListAliasIPRequest) (*models.ListAliasIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAliasIP(ctx, request)
}

func (service *BaseService) CreateAnalyticsNode(ctx context.Context, request *models.CreateAnalyticsNodeRequest) (*models.CreateAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAnalyticsNode(ctx, request)
}
func (service *BaseService) UpdateAnalyticsNode(ctx context.Context, request *models.UpdateAnalyticsNodeRequest) (*models.UpdateAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAnalyticsNode(ctx, request)
}
func (service *BaseService) DeleteAnalyticsNode(ctx context.Context, request *models.DeleteAnalyticsNodeRequest) (*models.DeleteAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAnalyticsNode(ctx, request)
}
func (service *BaseService) GetAnalyticsNode(ctx context.Context, request *models.GetAnalyticsNodeRequest) (*models.GetAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAnalyticsNode(ctx, request)
}
func (service *BaseService) ListAnalyticsNode(ctx context.Context, request *models.ListAnalyticsNodeRequest) (*models.ListAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAnalyticsNode(ctx, request)
}

func (service *BaseService) CreateAPIAccessList(ctx context.Context, request *models.CreateAPIAccessListRequest) (*models.CreateAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAPIAccessList(ctx, request)
}
func (service *BaseService) UpdateAPIAccessList(ctx context.Context, request *models.UpdateAPIAccessListRequest) (*models.UpdateAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAPIAccessList(ctx, request)
}
func (service *BaseService) DeleteAPIAccessList(ctx context.Context, request *models.DeleteAPIAccessListRequest) (*models.DeleteAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAPIAccessList(ctx, request)
}
func (service *BaseService) GetAPIAccessList(ctx context.Context, request *models.GetAPIAccessListRequest) (*models.GetAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAPIAccessList(ctx, request)
}
func (service *BaseService) ListAPIAccessList(ctx context.Context, request *models.ListAPIAccessListRequest) (*models.ListAPIAccessListResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAPIAccessList(ctx, request)
}

func (service *BaseService) CreateApplicationPolicySet(ctx context.Context, request *models.CreateApplicationPolicySetRequest) (*models.CreateApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateApplicationPolicySet(ctx, request)
}
func (service *BaseService) UpdateApplicationPolicySet(ctx context.Context, request *models.UpdateApplicationPolicySetRequest) (*models.UpdateApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateApplicationPolicySet(ctx, request)
}
func (service *BaseService) DeleteApplicationPolicySet(ctx context.Context, request *models.DeleteApplicationPolicySetRequest) (*models.DeleteApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteApplicationPolicySet(ctx, request)
}
func (service *BaseService) GetApplicationPolicySet(ctx context.Context, request *models.GetApplicationPolicySetRequest) (*models.GetApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetApplicationPolicySet(ctx, request)
}
func (service *BaseService) ListApplicationPolicySet(ctx context.Context, request *models.ListApplicationPolicySetRequest) (*models.ListApplicationPolicySetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListApplicationPolicySet(ctx, request)
}

func (service *BaseService) CreateBGPAsAService(ctx context.Context, request *models.CreateBGPAsAServiceRequest) (*models.CreateBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBGPAsAService(ctx, request)
}
func (service *BaseService) UpdateBGPAsAService(ctx context.Context, request *models.UpdateBGPAsAServiceRequest) (*models.UpdateBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPAsAService(ctx, request)
}
func (service *BaseService) DeleteBGPAsAService(ctx context.Context, request *models.DeleteBGPAsAServiceRequest) (*models.DeleteBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPAsAService(ctx, request)
}
func (service *BaseService) GetBGPAsAService(ctx context.Context, request *models.GetBGPAsAServiceRequest) (*models.GetBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBGPAsAService(ctx, request)
}
func (service *BaseService) ListBGPAsAService(ctx context.Context, request *models.ListBGPAsAServiceRequest) (*models.ListBGPAsAServiceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBGPAsAService(ctx, request)
}

func (service *BaseService) CreateBGPRouter(ctx context.Context, request *models.CreateBGPRouterRequest) (*models.CreateBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBGPRouter(ctx, request)
}
func (service *BaseService) UpdateBGPRouter(ctx context.Context, request *models.UpdateBGPRouterRequest) (*models.UpdateBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPRouter(ctx, request)
}
func (service *BaseService) DeleteBGPRouter(ctx context.Context, request *models.DeleteBGPRouterRequest) (*models.DeleteBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPRouter(ctx, request)
}
func (service *BaseService) GetBGPRouter(ctx context.Context, request *models.GetBGPRouterRequest) (*models.GetBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBGPRouter(ctx, request)
}
func (service *BaseService) ListBGPRouter(ctx context.Context, request *models.ListBGPRouterRequest) (*models.ListBGPRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBGPRouter(ctx, request)
}

func (service *BaseService) CreateBGPVPN(ctx context.Context, request *models.CreateBGPVPNRequest) (*models.CreateBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBGPVPN(ctx, request)
}
func (service *BaseService) UpdateBGPVPN(ctx context.Context, request *models.UpdateBGPVPNRequest) (*models.UpdateBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPVPN(ctx, request)
}
func (service *BaseService) DeleteBGPVPN(ctx context.Context, request *models.DeleteBGPVPNRequest) (*models.DeleteBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPVPN(ctx, request)
}
func (service *BaseService) GetBGPVPN(ctx context.Context, request *models.GetBGPVPNRequest) (*models.GetBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBGPVPN(ctx, request)
}
func (service *BaseService) ListBGPVPN(ctx context.Context, request *models.ListBGPVPNRequest) (*models.ListBGPVPNResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBGPVPN(ctx, request)
}

func (service *BaseService) CreateBridgeDomain(ctx context.Context, request *models.CreateBridgeDomainRequest) (*models.CreateBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBridgeDomain(ctx, request)
}
func (service *BaseService) UpdateBridgeDomain(ctx context.Context, request *models.UpdateBridgeDomainRequest) (*models.UpdateBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBridgeDomain(ctx, request)
}
func (service *BaseService) DeleteBridgeDomain(ctx context.Context, request *models.DeleteBridgeDomainRequest) (*models.DeleteBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBridgeDomain(ctx, request)
}
func (service *BaseService) GetBridgeDomain(ctx context.Context, request *models.GetBridgeDomainRequest) (*models.GetBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBridgeDomain(ctx, request)
}
func (service *BaseService) ListBridgeDomain(ctx context.Context, request *models.ListBridgeDomainRequest) (*models.ListBridgeDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBridgeDomain(ctx, request)
}

func (service *BaseService) CreateConfigNode(ctx context.Context, request *models.CreateConfigNodeRequest) (*models.CreateConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateConfigNode(ctx, request)
}
func (service *BaseService) UpdateConfigNode(ctx context.Context, request *models.UpdateConfigNodeRequest) (*models.UpdateConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateConfigNode(ctx, request)
}
func (service *BaseService) DeleteConfigNode(ctx context.Context, request *models.DeleteConfigNodeRequest) (*models.DeleteConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteConfigNode(ctx, request)
}
func (service *BaseService) GetConfigNode(ctx context.Context, request *models.GetConfigNodeRequest) (*models.GetConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetConfigNode(ctx, request)
}
func (service *BaseService) ListConfigNode(ctx context.Context, request *models.ListConfigNodeRequest) (*models.ListConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListConfigNode(ctx, request)
}

func (service *BaseService) CreateConfigRoot(ctx context.Context, request *models.CreateConfigRootRequest) (*models.CreateConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateConfigRoot(ctx, request)
}
func (service *BaseService) UpdateConfigRoot(ctx context.Context, request *models.UpdateConfigRootRequest) (*models.UpdateConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateConfigRoot(ctx, request)
}
func (service *BaseService) DeleteConfigRoot(ctx context.Context, request *models.DeleteConfigRootRequest) (*models.DeleteConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteConfigRoot(ctx, request)
}
func (service *BaseService) GetConfigRoot(ctx context.Context, request *models.GetConfigRootRequest) (*models.GetConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetConfigRoot(ctx, request)
}
func (service *BaseService) ListConfigRoot(ctx context.Context, request *models.ListConfigRootRequest) (*models.ListConfigRootResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListConfigRoot(ctx, request)
}

func (service *BaseService) CreateCustomerAttachment(ctx context.Context, request *models.CreateCustomerAttachmentRequest) (*models.CreateCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateCustomerAttachment(ctx, request)
}
func (service *BaseService) UpdateCustomerAttachment(ctx context.Context, request *models.UpdateCustomerAttachmentRequest) (*models.UpdateCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateCustomerAttachment(ctx, request)
}
func (service *BaseService) DeleteCustomerAttachment(ctx context.Context, request *models.DeleteCustomerAttachmentRequest) (*models.DeleteCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteCustomerAttachment(ctx, request)
}
func (service *BaseService) GetCustomerAttachment(ctx context.Context, request *models.GetCustomerAttachmentRequest) (*models.GetCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetCustomerAttachment(ctx, request)
}
func (service *BaseService) ListCustomerAttachment(ctx context.Context, request *models.ListCustomerAttachmentRequest) (*models.ListCustomerAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListCustomerAttachment(ctx, request)
}

func (service *BaseService) CreateDatabaseNode(ctx context.Context, request *models.CreateDatabaseNodeRequest) (*models.CreateDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDatabaseNode(ctx, request)
}
func (service *BaseService) UpdateDatabaseNode(ctx context.Context, request *models.UpdateDatabaseNodeRequest) (*models.UpdateDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDatabaseNode(ctx, request)
}
func (service *BaseService) DeleteDatabaseNode(ctx context.Context, request *models.DeleteDatabaseNodeRequest) (*models.DeleteDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDatabaseNode(ctx, request)
}
func (service *BaseService) GetDatabaseNode(ctx context.Context, request *models.GetDatabaseNodeRequest) (*models.GetDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDatabaseNode(ctx, request)
}
func (service *BaseService) ListDatabaseNode(ctx context.Context, request *models.ListDatabaseNodeRequest) (*models.ListDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDatabaseNode(ctx, request)
}

func (service *BaseService) CreateDiscoveryServiceAssignment(ctx context.Context, request *models.CreateDiscoveryServiceAssignmentRequest) (*models.CreateDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) UpdateDiscoveryServiceAssignment(ctx context.Context, request *models.UpdateDiscoveryServiceAssignmentRequest) (*models.UpdateDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) DeleteDiscoveryServiceAssignment(ctx context.Context, request *models.DeleteDiscoveryServiceAssignmentRequest) (*models.DeleteDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) GetDiscoveryServiceAssignment(ctx context.Context, request *models.GetDiscoveryServiceAssignmentRequest) (*models.GetDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDiscoveryServiceAssignment(ctx, request)
}
func (service *BaseService) ListDiscoveryServiceAssignment(ctx context.Context, request *models.ListDiscoveryServiceAssignmentRequest) (*models.ListDiscoveryServiceAssignmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDiscoveryServiceAssignment(ctx, request)
}

func (service *BaseService) CreateDomain(ctx context.Context, request *models.CreateDomainRequest) (*models.CreateDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDomain(ctx, request)
}
func (service *BaseService) UpdateDomain(ctx context.Context, request *models.UpdateDomainRequest) (*models.UpdateDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDomain(ctx, request)
}
func (service *BaseService) DeleteDomain(ctx context.Context, request *models.DeleteDomainRequest) (*models.DeleteDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDomain(ctx, request)
}
func (service *BaseService) GetDomain(ctx context.Context, request *models.GetDomainRequest) (*models.GetDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDomain(ctx, request)
}
func (service *BaseService) ListDomain(ctx context.Context, request *models.ListDomainRequest) (*models.ListDomainResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDomain(ctx, request)
}

func (service *BaseService) CreateDsaRule(ctx context.Context, request *models.CreateDsaRuleRequest) (*models.CreateDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDsaRule(ctx, request)
}
func (service *BaseService) UpdateDsaRule(ctx context.Context, request *models.UpdateDsaRuleRequest) (*models.UpdateDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDsaRule(ctx, request)
}
func (service *BaseService) DeleteDsaRule(ctx context.Context, request *models.DeleteDsaRuleRequest) (*models.DeleteDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDsaRule(ctx, request)
}
func (service *BaseService) GetDsaRule(ctx context.Context, request *models.GetDsaRuleRequest) (*models.GetDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDsaRule(ctx, request)
}
func (service *BaseService) ListDsaRule(ctx context.Context, request *models.ListDsaRuleRequest) (*models.ListDsaRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDsaRule(ctx, request)
}

func (service *BaseService) CreateE2ServiceProvider(ctx context.Context, request *models.CreateE2ServiceProviderRequest) (*models.CreateE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateE2ServiceProvider(ctx, request)
}
func (service *BaseService) UpdateE2ServiceProvider(ctx context.Context, request *models.UpdateE2ServiceProviderRequest) (*models.UpdateE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateE2ServiceProvider(ctx, request)
}
func (service *BaseService) DeleteE2ServiceProvider(ctx context.Context, request *models.DeleteE2ServiceProviderRequest) (*models.DeleteE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteE2ServiceProvider(ctx, request)
}
func (service *BaseService) GetE2ServiceProvider(ctx context.Context, request *models.GetE2ServiceProviderRequest) (*models.GetE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetE2ServiceProvider(ctx, request)
}
func (service *BaseService) ListE2ServiceProvider(ctx context.Context, request *models.ListE2ServiceProviderRequest) (*models.ListE2ServiceProviderResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListE2ServiceProvider(ctx, request)
}

func (service *BaseService) CreateFirewallPolicy(ctx context.Context, request *models.CreateFirewallPolicyRequest) (*models.CreateFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFirewallPolicy(ctx, request)
}
func (service *BaseService) UpdateFirewallPolicy(ctx context.Context, request *models.UpdateFirewallPolicyRequest) (*models.UpdateFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFirewallPolicy(ctx, request)
}
func (service *BaseService) DeleteFirewallPolicy(ctx context.Context, request *models.DeleteFirewallPolicyRequest) (*models.DeleteFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFirewallPolicy(ctx, request)
}
func (service *BaseService) GetFirewallPolicy(ctx context.Context, request *models.GetFirewallPolicyRequest) (*models.GetFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFirewallPolicy(ctx, request)
}
func (service *BaseService) ListFirewallPolicy(ctx context.Context, request *models.ListFirewallPolicyRequest) (*models.ListFirewallPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFirewallPolicy(ctx, request)
}

func (service *BaseService) CreateFirewallRule(ctx context.Context, request *models.CreateFirewallRuleRequest) (*models.CreateFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFirewallRule(ctx, request)
}
func (service *BaseService) UpdateFirewallRule(ctx context.Context, request *models.UpdateFirewallRuleRequest) (*models.UpdateFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFirewallRule(ctx, request)
}
func (service *BaseService) DeleteFirewallRule(ctx context.Context, request *models.DeleteFirewallRuleRequest) (*models.DeleteFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFirewallRule(ctx, request)
}
func (service *BaseService) GetFirewallRule(ctx context.Context, request *models.GetFirewallRuleRequest) (*models.GetFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFirewallRule(ctx, request)
}
func (service *BaseService) ListFirewallRule(ctx context.Context, request *models.ListFirewallRuleRequest) (*models.ListFirewallRuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFirewallRule(ctx, request)
}

func (service *BaseService) CreateFloatingIPPool(ctx context.Context, request *models.CreateFloatingIPPoolRequest) (*models.CreateFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFloatingIPPool(ctx, request)
}
func (service *BaseService) UpdateFloatingIPPool(ctx context.Context, request *models.UpdateFloatingIPPoolRequest) (*models.UpdateFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFloatingIPPool(ctx, request)
}
func (service *BaseService) DeleteFloatingIPPool(ctx context.Context, request *models.DeleteFloatingIPPoolRequest) (*models.DeleteFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFloatingIPPool(ctx, request)
}
func (service *BaseService) GetFloatingIPPool(ctx context.Context, request *models.GetFloatingIPPoolRequest) (*models.GetFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFloatingIPPool(ctx, request)
}
func (service *BaseService) ListFloatingIPPool(ctx context.Context, request *models.ListFloatingIPPoolRequest) (*models.ListFloatingIPPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFloatingIPPool(ctx, request)
}

func (service *BaseService) CreateFloatingIP(ctx context.Context, request *models.CreateFloatingIPRequest) (*models.CreateFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFloatingIP(ctx, request)
}
func (service *BaseService) UpdateFloatingIP(ctx context.Context, request *models.UpdateFloatingIPRequest) (*models.UpdateFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFloatingIP(ctx, request)
}
func (service *BaseService) DeleteFloatingIP(ctx context.Context, request *models.DeleteFloatingIPRequest) (*models.DeleteFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFloatingIP(ctx, request)
}
func (service *BaseService) GetFloatingIP(ctx context.Context, request *models.GetFloatingIPRequest) (*models.GetFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFloatingIP(ctx, request)
}
func (service *BaseService) ListFloatingIP(ctx context.Context, request *models.ListFloatingIPRequest) (*models.ListFloatingIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFloatingIP(ctx, request)
}

func (service *BaseService) CreateForwardingClass(ctx context.Context, request *models.CreateForwardingClassRequest) (*models.CreateForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateForwardingClass(ctx, request)
}
func (service *BaseService) UpdateForwardingClass(ctx context.Context, request *models.UpdateForwardingClassRequest) (*models.UpdateForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateForwardingClass(ctx, request)
}
func (service *BaseService) DeleteForwardingClass(ctx context.Context, request *models.DeleteForwardingClassRequest) (*models.DeleteForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteForwardingClass(ctx, request)
}
func (service *BaseService) GetForwardingClass(ctx context.Context, request *models.GetForwardingClassRequest) (*models.GetForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetForwardingClass(ctx, request)
}
func (service *BaseService) ListForwardingClass(ctx context.Context, request *models.ListForwardingClassRequest) (*models.ListForwardingClassResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListForwardingClass(ctx, request)
}

func (service *BaseService) CreateGlobalQosConfig(ctx context.Context, request *models.CreateGlobalQosConfigRequest) (*models.CreateGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalQosConfig(ctx, request)
}
func (service *BaseService) UpdateGlobalQosConfig(ctx context.Context, request *models.UpdateGlobalQosConfigRequest) (*models.UpdateGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalQosConfig(ctx, request)
}
func (service *BaseService) DeleteGlobalQosConfig(ctx context.Context, request *models.DeleteGlobalQosConfigRequest) (*models.DeleteGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalQosConfig(ctx, request)
}
func (service *BaseService) GetGlobalQosConfig(ctx context.Context, request *models.GetGlobalQosConfigRequest) (*models.GetGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetGlobalQosConfig(ctx, request)
}
func (service *BaseService) ListGlobalQosConfig(ctx context.Context, request *models.ListGlobalQosConfigRequest) (*models.ListGlobalQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListGlobalQosConfig(ctx, request)
}

func (service *BaseService) CreateGlobalSystemConfig(ctx context.Context, request *models.CreateGlobalSystemConfigRequest) (*models.CreateGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalSystemConfig(ctx, request)
}
func (service *BaseService) UpdateGlobalSystemConfig(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest) (*models.UpdateGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalSystemConfig(ctx, request)
}
func (service *BaseService) DeleteGlobalSystemConfig(ctx context.Context, request *models.DeleteGlobalSystemConfigRequest) (*models.DeleteGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalSystemConfig(ctx, request)
}
func (service *BaseService) GetGlobalSystemConfig(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetGlobalSystemConfig(ctx, request)
}
func (service *BaseService) ListGlobalSystemConfig(ctx context.Context, request *models.ListGlobalSystemConfigRequest) (*models.ListGlobalSystemConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListGlobalSystemConfig(ctx, request)
}

func (service *BaseService) CreateGlobalVrouterConfig(ctx context.Context, request *models.CreateGlobalVrouterConfigRequest) (*models.CreateGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) UpdateGlobalVrouterConfig(ctx context.Context, request *models.UpdateGlobalVrouterConfigRequest) (*models.UpdateGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) DeleteGlobalVrouterConfig(ctx context.Context, request *models.DeleteGlobalVrouterConfigRequest) (*models.DeleteGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) GetGlobalVrouterConfig(ctx context.Context, request *models.GetGlobalVrouterConfigRequest) (*models.GetGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetGlobalVrouterConfig(ctx, request)
}
func (service *BaseService) ListGlobalVrouterConfig(ctx context.Context, request *models.ListGlobalVrouterConfigRequest) (*models.ListGlobalVrouterConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListGlobalVrouterConfig(ctx, request)
}

func (service *BaseService) CreateInstanceIP(ctx context.Context, request *models.CreateInstanceIPRequest) (*models.CreateInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateInstanceIP(ctx, request)
}
func (service *BaseService) UpdateInstanceIP(ctx context.Context, request *models.UpdateInstanceIPRequest) (*models.UpdateInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateInstanceIP(ctx, request)
}
func (service *BaseService) DeleteInstanceIP(ctx context.Context, request *models.DeleteInstanceIPRequest) (*models.DeleteInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteInstanceIP(ctx, request)
}
func (service *BaseService) GetInstanceIP(ctx context.Context, request *models.GetInstanceIPRequest) (*models.GetInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetInstanceIP(ctx, request)
}
func (service *BaseService) ListInstanceIP(ctx context.Context, request *models.ListInstanceIPRequest) (*models.ListInstanceIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListInstanceIP(ctx, request)
}

func (service *BaseService) CreateInterfaceRouteTable(ctx context.Context, request *models.CreateInterfaceRouteTableRequest) (*models.CreateInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateInterfaceRouteTable(ctx, request)
}
func (service *BaseService) UpdateInterfaceRouteTable(ctx context.Context, request *models.UpdateInterfaceRouteTableRequest) (*models.UpdateInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateInterfaceRouteTable(ctx, request)
}
func (service *BaseService) DeleteInterfaceRouteTable(ctx context.Context, request *models.DeleteInterfaceRouteTableRequest) (*models.DeleteInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteInterfaceRouteTable(ctx, request)
}
func (service *BaseService) GetInterfaceRouteTable(ctx context.Context, request *models.GetInterfaceRouteTableRequest) (*models.GetInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetInterfaceRouteTable(ctx, request)
}
func (service *BaseService) ListInterfaceRouteTable(ctx context.Context, request *models.ListInterfaceRouteTableRequest) (*models.ListInterfaceRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListInterfaceRouteTable(ctx, request)
}

func (service *BaseService) CreateLoadbalancerHealthmonitor(ctx context.Context, request *models.CreateLoadbalancerHealthmonitorRequest) (*models.CreateLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerHealthmonitor(ctx context.Context, request *models.UpdateLoadbalancerHealthmonitorRequest) (*models.UpdateLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerHealthmonitor(ctx context.Context, request *models.DeleteLoadbalancerHealthmonitorRequest) (*models.DeleteLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) GetLoadbalancerHealthmonitor(ctx context.Context, request *models.GetLoadbalancerHealthmonitorRequest) (*models.GetLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerHealthmonitor(ctx, request)
}
func (service *BaseService) ListLoadbalancerHealthmonitor(ctx context.Context, request *models.ListLoadbalancerHealthmonitorRequest) (*models.ListLoadbalancerHealthmonitorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerHealthmonitor(ctx, request)
}

func (service *BaseService) CreateLoadbalancerListener(ctx context.Context, request *models.CreateLoadbalancerListenerRequest) (*models.CreateLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerListener(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerListener(ctx context.Context, request *models.UpdateLoadbalancerListenerRequest) (*models.UpdateLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerListener(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerListener(ctx context.Context, request *models.DeleteLoadbalancerListenerRequest) (*models.DeleteLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerListener(ctx, request)
}
func (service *BaseService) GetLoadbalancerListener(ctx context.Context, request *models.GetLoadbalancerListenerRequest) (*models.GetLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerListener(ctx, request)
}
func (service *BaseService) ListLoadbalancerListener(ctx context.Context, request *models.ListLoadbalancerListenerRequest) (*models.ListLoadbalancerListenerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerListener(ctx, request)
}

func (service *BaseService) CreateLoadbalancerMember(ctx context.Context, request *models.CreateLoadbalancerMemberRequest) (*models.CreateLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerMember(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerMember(ctx context.Context, request *models.UpdateLoadbalancerMemberRequest) (*models.UpdateLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerMember(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerMember(ctx context.Context, request *models.DeleteLoadbalancerMemberRequest) (*models.DeleteLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerMember(ctx, request)
}
func (service *BaseService) GetLoadbalancerMember(ctx context.Context, request *models.GetLoadbalancerMemberRequest) (*models.GetLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerMember(ctx, request)
}
func (service *BaseService) ListLoadbalancerMember(ctx context.Context, request *models.ListLoadbalancerMemberRequest) (*models.ListLoadbalancerMemberResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerMember(ctx, request)
}

func (service *BaseService) CreateLoadbalancerPool(ctx context.Context, request *models.CreateLoadbalancerPoolRequest) (*models.CreateLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerPool(ctx, request)
}
func (service *BaseService) UpdateLoadbalancerPool(ctx context.Context, request *models.UpdateLoadbalancerPoolRequest) (*models.UpdateLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerPool(ctx, request)
}
func (service *BaseService) DeleteLoadbalancerPool(ctx context.Context, request *models.DeleteLoadbalancerPoolRequest) (*models.DeleteLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerPool(ctx, request)
}
func (service *BaseService) GetLoadbalancerPool(ctx context.Context, request *models.GetLoadbalancerPoolRequest) (*models.GetLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancerPool(ctx, request)
}
func (service *BaseService) ListLoadbalancerPool(ctx context.Context, request *models.ListLoadbalancerPoolRequest) (*models.ListLoadbalancerPoolResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancerPool(ctx, request)
}

func (service *BaseService) CreateLoadbalancer(ctx context.Context, request *models.CreateLoadbalancerRequest) (*models.CreateLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancer(ctx, request)
}
func (service *BaseService) UpdateLoadbalancer(ctx context.Context, request *models.UpdateLoadbalancerRequest) (*models.UpdateLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancer(ctx, request)
}
func (service *BaseService) DeleteLoadbalancer(ctx context.Context, request *models.DeleteLoadbalancerRequest) (*models.DeleteLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancer(ctx, request)
}
func (service *BaseService) GetLoadbalancer(ctx context.Context, request *models.GetLoadbalancerRequest) (*models.GetLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLoadbalancer(ctx, request)
}
func (service *BaseService) ListLoadbalancer(ctx context.Context, request *models.ListLoadbalancerRequest) (*models.ListLoadbalancerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLoadbalancer(ctx, request)
}

func (service *BaseService) CreateLogicalInterface(ctx context.Context, request *models.CreateLogicalInterfaceRequest) (*models.CreateLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLogicalInterface(ctx, request)
}
func (service *BaseService) UpdateLogicalInterface(ctx context.Context, request *models.UpdateLogicalInterfaceRequest) (*models.UpdateLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLogicalInterface(ctx, request)
}
func (service *BaseService) DeleteLogicalInterface(ctx context.Context, request *models.DeleteLogicalInterfaceRequest) (*models.DeleteLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLogicalInterface(ctx, request)
}
func (service *BaseService) GetLogicalInterface(ctx context.Context, request *models.GetLogicalInterfaceRequest) (*models.GetLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLogicalInterface(ctx, request)
}
func (service *BaseService) ListLogicalInterface(ctx context.Context, request *models.ListLogicalInterfaceRequest) (*models.ListLogicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLogicalInterface(ctx, request)
}

func (service *BaseService) CreateLogicalRouter(ctx context.Context, request *models.CreateLogicalRouterRequest) (*models.CreateLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLogicalRouter(ctx, request)
}
func (service *BaseService) UpdateLogicalRouter(ctx context.Context, request *models.UpdateLogicalRouterRequest) (*models.UpdateLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLogicalRouter(ctx, request)
}
func (service *BaseService) DeleteLogicalRouter(ctx context.Context, request *models.DeleteLogicalRouterRequest) (*models.DeleteLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLogicalRouter(ctx, request)
}
func (service *BaseService) GetLogicalRouter(ctx context.Context, request *models.GetLogicalRouterRequest) (*models.GetLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLogicalRouter(ctx, request)
}
func (service *BaseService) ListLogicalRouter(ctx context.Context, request *models.ListLogicalRouterRequest) (*models.ListLogicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLogicalRouter(ctx, request)
}

func (service *BaseService) CreateNamespace(ctx context.Context, request *models.CreateNamespaceRequest) (*models.CreateNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNamespace(ctx, request)
}
func (service *BaseService) UpdateNamespace(ctx context.Context, request *models.UpdateNamespaceRequest) (*models.UpdateNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNamespace(ctx, request)
}
func (service *BaseService) DeleteNamespace(ctx context.Context, request *models.DeleteNamespaceRequest) (*models.DeleteNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNamespace(ctx, request)
}
func (service *BaseService) GetNamespace(ctx context.Context, request *models.GetNamespaceRequest) (*models.GetNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNamespace(ctx, request)
}
func (service *BaseService) ListNamespace(ctx context.Context, request *models.ListNamespaceRequest) (*models.ListNamespaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNamespace(ctx, request)
}

func (service *BaseService) CreateNetworkDeviceConfig(ctx context.Context, request *models.CreateNetworkDeviceConfigRequest) (*models.CreateNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) UpdateNetworkDeviceConfig(ctx context.Context, request *models.UpdateNetworkDeviceConfigRequest) (*models.UpdateNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) DeleteNetworkDeviceConfig(ctx context.Context, request *models.DeleteNetworkDeviceConfigRequest) (*models.DeleteNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) GetNetworkDeviceConfig(ctx context.Context, request *models.GetNetworkDeviceConfigRequest) (*models.GetNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNetworkDeviceConfig(ctx, request)
}
func (service *BaseService) ListNetworkDeviceConfig(ctx context.Context, request *models.ListNetworkDeviceConfigRequest) (*models.ListNetworkDeviceConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNetworkDeviceConfig(ctx, request)
}

func (service *BaseService) CreateNetworkIpam(ctx context.Context, request *models.CreateNetworkIpamRequest) (*models.CreateNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkIpam(ctx, request)
}
func (service *BaseService) UpdateNetworkIpam(ctx context.Context, request *models.UpdateNetworkIpamRequest) (*models.UpdateNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkIpam(ctx, request)
}
func (service *BaseService) DeleteNetworkIpam(ctx context.Context, request *models.DeleteNetworkIpamRequest) (*models.DeleteNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkIpam(ctx, request)
}
func (service *BaseService) GetNetworkIpam(ctx context.Context, request *models.GetNetworkIpamRequest) (*models.GetNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNetworkIpam(ctx, request)
}
func (service *BaseService) ListNetworkIpam(ctx context.Context, request *models.ListNetworkIpamRequest) (*models.ListNetworkIpamResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNetworkIpam(ctx, request)
}

func (service *BaseService) CreateNetworkPolicy(ctx context.Context, request *models.CreateNetworkPolicyRequest) (*models.CreateNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkPolicy(ctx, request)
}
func (service *BaseService) UpdateNetworkPolicy(ctx context.Context, request *models.UpdateNetworkPolicyRequest) (*models.UpdateNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkPolicy(ctx, request)
}
func (service *BaseService) DeleteNetworkPolicy(ctx context.Context, request *models.DeleteNetworkPolicyRequest) (*models.DeleteNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkPolicy(ctx, request)
}
func (service *BaseService) GetNetworkPolicy(ctx context.Context, request *models.GetNetworkPolicyRequest) (*models.GetNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNetworkPolicy(ctx, request)
}
func (service *BaseService) ListNetworkPolicy(ctx context.Context, request *models.ListNetworkPolicyRequest) (*models.ListNetworkPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNetworkPolicy(ctx, request)
}

func (service *BaseService) CreatePeeringPolicy(ctx context.Context, request *models.CreatePeeringPolicyRequest) (*models.CreatePeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePeeringPolicy(ctx, request)
}
func (service *BaseService) UpdatePeeringPolicy(ctx context.Context, request *models.UpdatePeeringPolicyRequest) (*models.UpdatePeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePeeringPolicy(ctx, request)
}
func (service *BaseService) DeletePeeringPolicy(ctx context.Context, request *models.DeletePeeringPolicyRequest) (*models.DeletePeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePeeringPolicy(ctx, request)
}
func (service *BaseService) GetPeeringPolicy(ctx context.Context, request *models.GetPeeringPolicyRequest) (*models.GetPeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPeeringPolicy(ctx, request)
}
func (service *BaseService) ListPeeringPolicy(ctx context.Context, request *models.ListPeeringPolicyRequest) (*models.ListPeeringPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPeeringPolicy(ctx, request)
}

func (service *BaseService) CreatePhysicalInterface(ctx context.Context, request *models.CreatePhysicalInterfaceRequest) (*models.CreatePhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePhysicalInterface(ctx, request)
}
func (service *BaseService) UpdatePhysicalInterface(ctx context.Context, request *models.UpdatePhysicalInterfaceRequest) (*models.UpdatePhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePhysicalInterface(ctx, request)
}
func (service *BaseService) DeletePhysicalInterface(ctx context.Context, request *models.DeletePhysicalInterfaceRequest) (*models.DeletePhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePhysicalInterface(ctx, request)
}
func (service *BaseService) GetPhysicalInterface(ctx context.Context, request *models.GetPhysicalInterfaceRequest) (*models.GetPhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPhysicalInterface(ctx, request)
}
func (service *BaseService) ListPhysicalInterface(ctx context.Context, request *models.ListPhysicalInterfaceRequest) (*models.ListPhysicalInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPhysicalInterface(ctx, request)
}

func (service *BaseService) CreatePhysicalRouter(ctx context.Context, request *models.CreatePhysicalRouterRequest) (*models.CreatePhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePhysicalRouter(ctx, request)
}
func (service *BaseService) UpdatePhysicalRouter(ctx context.Context, request *models.UpdatePhysicalRouterRequest) (*models.UpdatePhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePhysicalRouter(ctx, request)
}
func (service *BaseService) DeletePhysicalRouter(ctx context.Context, request *models.DeletePhysicalRouterRequest) (*models.DeletePhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePhysicalRouter(ctx, request)
}
func (service *BaseService) GetPhysicalRouter(ctx context.Context, request *models.GetPhysicalRouterRequest) (*models.GetPhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPhysicalRouter(ctx, request)
}
func (service *BaseService) ListPhysicalRouter(ctx context.Context, request *models.ListPhysicalRouterRequest) (*models.ListPhysicalRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPhysicalRouter(ctx, request)
}

func (service *BaseService) CreatePolicyManagement(ctx context.Context, request *models.CreatePolicyManagementRequest) (*models.CreatePolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePolicyManagement(ctx, request)
}
func (service *BaseService) UpdatePolicyManagement(ctx context.Context, request *models.UpdatePolicyManagementRequest) (*models.UpdatePolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePolicyManagement(ctx, request)
}
func (service *BaseService) DeletePolicyManagement(ctx context.Context, request *models.DeletePolicyManagementRequest) (*models.DeletePolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePolicyManagement(ctx, request)
}
func (service *BaseService) GetPolicyManagement(ctx context.Context, request *models.GetPolicyManagementRequest) (*models.GetPolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPolicyManagement(ctx, request)
}
func (service *BaseService) ListPolicyManagement(ctx context.Context, request *models.ListPolicyManagementRequest) (*models.ListPolicyManagementResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPolicyManagement(ctx, request)
}

func (service *BaseService) CreatePortTuple(ctx context.Context, request *models.CreatePortTupleRequest) (*models.CreatePortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreatePortTuple(ctx, request)
}
func (service *BaseService) UpdatePortTuple(ctx context.Context, request *models.UpdatePortTupleRequest) (*models.UpdatePortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdatePortTuple(ctx, request)
}
func (service *BaseService) DeletePortTuple(ctx context.Context, request *models.DeletePortTupleRequest) (*models.DeletePortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeletePortTuple(ctx, request)
}
func (service *BaseService) GetPortTuple(ctx context.Context, request *models.GetPortTupleRequest) (*models.GetPortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetPortTuple(ctx, request)
}
func (service *BaseService) ListPortTuple(ctx context.Context, request *models.ListPortTupleRequest) (*models.ListPortTupleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListPortTuple(ctx, request)
}

func (service *BaseService) CreateProject(ctx context.Context, request *models.CreateProjectRequest) (*models.CreateProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateProject(ctx, request)
}
func (service *BaseService) UpdateProject(ctx context.Context, request *models.UpdateProjectRequest) (*models.UpdateProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateProject(ctx, request)
}
func (service *BaseService) DeleteProject(ctx context.Context, request *models.DeleteProjectRequest) (*models.DeleteProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteProject(ctx, request)
}
func (service *BaseService) GetProject(ctx context.Context, request *models.GetProjectRequest) (*models.GetProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetProject(ctx, request)
}
func (service *BaseService) ListProject(ctx context.Context, request *models.ListProjectRequest) (*models.ListProjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListProject(ctx, request)
}

func (service *BaseService) CreateProviderAttachment(ctx context.Context, request *models.CreateProviderAttachmentRequest) (*models.CreateProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateProviderAttachment(ctx, request)
}
func (service *BaseService) UpdateProviderAttachment(ctx context.Context, request *models.UpdateProviderAttachmentRequest) (*models.UpdateProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateProviderAttachment(ctx, request)
}
func (service *BaseService) DeleteProviderAttachment(ctx context.Context, request *models.DeleteProviderAttachmentRequest) (*models.DeleteProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteProviderAttachment(ctx, request)
}
func (service *BaseService) GetProviderAttachment(ctx context.Context, request *models.GetProviderAttachmentRequest) (*models.GetProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetProviderAttachment(ctx, request)
}
func (service *BaseService) ListProviderAttachment(ctx context.Context, request *models.ListProviderAttachmentRequest) (*models.ListProviderAttachmentResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListProviderAttachment(ctx, request)
}

func (service *BaseService) CreateQosConfig(ctx context.Context, request *models.CreateQosConfigRequest) (*models.CreateQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateQosConfig(ctx, request)
}
func (service *BaseService) UpdateQosConfig(ctx context.Context, request *models.UpdateQosConfigRequest) (*models.UpdateQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateQosConfig(ctx, request)
}
func (service *BaseService) DeleteQosConfig(ctx context.Context, request *models.DeleteQosConfigRequest) (*models.DeleteQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteQosConfig(ctx, request)
}
func (service *BaseService) GetQosConfig(ctx context.Context, request *models.GetQosConfigRequest) (*models.GetQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetQosConfig(ctx, request)
}
func (service *BaseService) ListQosConfig(ctx context.Context, request *models.ListQosConfigRequest) (*models.ListQosConfigResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListQosConfig(ctx, request)
}

func (service *BaseService) CreateQosQueue(ctx context.Context, request *models.CreateQosQueueRequest) (*models.CreateQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateQosQueue(ctx, request)
}
func (service *BaseService) UpdateQosQueue(ctx context.Context, request *models.UpdateQosQueueRequest) (*models.UpdateQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateQosQueue(ctx, request)
}
func (service *BaseService) DeleteQosQueue(ctx context.Context, request *models.DeleteQosQueueRequest) (*models.DeleteQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteQosQueue(ctx, request)
}
func (service *BaseService) GetQosQueue(ctx context.Context, request *models.GetQosQueueRequest) (*models.GetQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetQosQueue(ctx, request)
}
func (service *BaseService) ListQosQueue(ctx context.Context, request *models.ListQosQueueRequest) (*models.ListQosQueueResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListQosQueue(ctx, request)
}

func (service *BaseService) CreateRouteAggregate(ctx context.Context, request *models.CreateRouteAggregateRequest) (*models.CreateRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRouteAggregate(ctx, request)
}
func (service *BaseService) UpdateRouteAggregate(ctx context.Context, request *models.UpdateRouteAggregateRequest) (*models.UpdateRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteAggregate(ctx, request)
}
func (service *BaseService) DeleteRouteAggregate(ctx context.Context, request *models.DeleteRouteAggregateRequest) (*models.DeleteRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteAggregate(ctx, request)
}
func (service *BaseService) GetRouteAggregate(ctx context.Context, request *models.GetRouteAggregateRequest) (*models.GetRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRouteAggregate(ctx, request)
}
func (service *BaseService) ListRouteAggregate(ctx context.Context, request *models.ListRouteAggregateRequest) (*models.ListRouteAggregateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRouteAggregate(ctx, request)
}

func (service *BaseService) CreateRouteTable(ctx context.Context, request *models.CreateRouteTableRequest) (*models.CreateRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRouteTable(ctx, request)
}
func (service *BaseService) UpdateRouteTable(ctx context.Context, request *models.UpdateRouteTableRequest) (*models.UpdateRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteTable(ctx, request)
}
func (service *BaseService) DeleteRouteTable(ctx context.Context, request *models.DeleteRouteTableRequest) (*models.DeleteRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteTable(ctx, request)
}
func (service *BaseService) GetRouteTable(ctx context.Context, request *models.GetRouteTableRequest) (*models.GetRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRouteTable(ctx, request)
}
func (service *BaseService) ListRouteTable(ctx context.Context, request *models.ListRouteTableRequest) (*models.ListRouteTableResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRouteTable(ctx, request)
}

func (service *BaseService) CreateRouteTarget(ctx context.Context, request *models.CreateRouteTargetRequest) (*models.CreateRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRouteTarget(ctx, request)
}
func (service *BaseService) UpdateRouteTarget(ctx context.Context, request *models.UpdateRouteTargetRequest) (*models.UpdateRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteTarget(ctx, request)
}
func (service *BaseService) DeleteRouteTarget(ctx context.Context, request *models.DeleteRouteTargetRequest) (*models.DeleteRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteTarget(ctx, request)
}
func (service *BaseService) GetRouteTarget(ctx context.Context, request *models.GetRouteTargetRequest) (*models.GetRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRouteTarget(ctx, request)
}
func (service *BaseService) ListRouteTarget(ctx context.Context, request *models.ListRouteTargetRequest) (*models.ListRouteTargetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRouteTarget(ctx, request)
}

func (service *BaseService) CreateRoutingInstance(ctx context.Context, request *models.CreateRoutingInstanceRequest) (*models.CreateRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRoutingInstance(ctx, request)
}
func (service *BaseService) UpdateRoutingInstance(ctx context.Context, request *models.UpdateRoutingInstanceRequest) (*models.UpdateRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRoutingInstance(ctx, request)
}
func (service *BaseService) DeleteRoutingInstance(ctx context.Context, request *models.DeleteRoutingInstanceRequest) (*models.DeleteRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRoutingInstance(ctx, request)
}
func (service *BaseService) GetRoutingInstance(ctx context.Context, request *models.GetRoutingInstanceRequest) (*models.GetRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRoutingInstance(ctx, request)
}
func (service *BaseService) ListRoutingInstance(ctx context.Context, request *models.ListRoutingInstanceRequest) (*models.ListRoutingInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRoutingInstance(ctx, request)
}

func (service *BaseService) CreateRoutingPolicy(ctx context.Context, request *models.CreateRoutingPolicyRequest) (*models.CreateRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateRoutingPolicy(ctx, request)
}
func (service *BaseService) UpdateRoutingPolicy(ctx context.Context, request *models.UpdateRoutingPolicyRequest) (*models.UpdateRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateRoutingPolicy(ctx, request)
}
func (service *BaseService) DeleteRoutingPolicy(ctx context.Context, request *models.DeleteRoutingPolicyRequest) (*models.DeleteRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteRoutingPolicy(ctx, request)
}
func (service *BaseService) GetRoutingPolicy(ctx context.Context, request *models.GetRoutingPolicyRequest) (*models.GetRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetRoutingPolicy(ctx, request)
}
func (service *BaseService) ListRoutingPolicy(ctx context.Context, request *models.ListRoutingPolicyRequest) (*models.ListRoutingPolicyResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListRoutingPolicy(ctx, request)
}

func (service *BaseService) CreateSecurityGroup(ctx context.Context, request *models.CreateSecurityGroupRequest) (*models.CreateSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateSecurityGroup(ctx, request)
}
func (service *BaseService) UpdateSecurityGroup(ctx context.Context, request *models.UpdateSecurityGroupRequest) (*models.UpdateSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateSecurityGroup(ctx, request)
}
func (service *BaseService) DeleteSecurityGroup(ctx context.Context, request *models.DeleteSecurityGroupRequest) (*models.DeleteSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteSecurityGroup(ctx, request)
}
func (service *BaseService) GetSecurityGroup(ctx context.Context, request *models.GetSecurityGroupRequest) (*models.GetSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetSecurityGroup(ctx, request)
}
func (service *BaseService) ListSecurityGroup(ctx context.Context, request *models.ListSecurityGroupRequest) (*models.ListSecurityGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListSecurityGroup(ctx, request)
}

func (service *BaseService) CreateSecurityLoggingObject(ctx context.Context, request *models.CreateSecurityLoggingObjectRequest) (*models.CreateSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateSecurityLoggingObject(ctx, request)
}
func (service *BaseService) UpdateSecurityLoggingObject(ctx context.Context, request *models.UpdateSecurityLoggingObjectRequest) (*models.UpdateSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateSecurityLoggingObject(ctx, request)
}
func (service *BaseService) DeleteSecurityLoggingObject(ctx context.Context, request *models.DeleteSecurityLoggingObjectRequest) (*models.DeleteSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteSecurityLoggingObject(ctx, request)
}
func (service *BaseService) GetSecurityLoggingObject(ctx context.Context, request *models.GetSecurityLoggingObjectRequest) (*models.GetSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetSecurityLoggingObject(ctx, request)
}
func (service *BaseService) ListSecurityLoggingObject(ctx context.Context, request *models.ListSecurityLoggingObjectRequest) (*models.ListSecurityLoggingObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListSecurityLoggingObject(ctx, request)
}

func (service *BaseService) CreateServiceAppliance(ctx context.Context, request *models.CreateServiceApplianceRequest) (*models.CreateServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceAppliance(ctx, request)
}
func (service *BaseService) UpdateServiceAppliance(ctx context.Context, request *models.UpdateServiceApplianceRequest) (*models.UpdateServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceAppliance(ctx, request)
}
func (service *BaseService) DeleteServiceAppliance(ctx context.Context, request *models.DeleteServiceApplianceRequest) (*models.DeleteServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceAppliance(ctx, request)
}
func (service *BaseService) GetServiceAppliance(ctx context.Context, request *models.GetServiceApplianceRequest) (*models.GetServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceAppliance(ctx, request)
}
func (service *BaseService) ListServiceAppliance(ctx context.Context, request *models.ListServiceApplianceRequest) (*models.ListServiceApplianceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceAppliance(ctx, request)
}

func (service *BaseService) CreateServiceApplianceSet(ctx context.Context, request *models.CreateServiceApplianceSetRequest) (*models.CreateServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceApplianceSet(ctx, request)
}
func (service *BaseService) UpdateServiceApplianceSet(ctx context.Context, request *models.UpdateServiceApplianceSetRequest) (*models.UpdateServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceApplianceSet(ctx, request)
}
func (service *BaseService) DeleteServiceApplianceSet(ctx context.Context, request *models.DeleteServiceApplianceSetRequest) (*models.DeleteServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceApplianceSet(ctx, request)
}
func (service *BaseService) GetServiceApplianceSet(ctx context.Context, request *models.GetServiceApplianceSetRequest) (*models.GetServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceApplianceSet(ctx, request)
}
func (service *BaseService) ListServiceApplianceSet(ctx context.Context, request *models.ListServiceApplianceSetRequest) (*models.ListServiceApplianceSetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceApplianceSet(ctx, request)
}

func (service *BaseService) CreateServiceConnectionModule(ctx context.Context, request *models.CreateServiceConnectionModuleRequest) (*models.CreateServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceConnectionModule(ctx, request)
}
func (service *BaseService) UpdateServiceConnectionModule(ctx context.Context, request *models.UpdateServiceConnectionModuleRequest) (*models.UpdateServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceConnectionModule(ctx, request)
}
func (service *BaseService) DeleteServiceConnectionModule(ctx context.Context, request *models.DeleteServiceConnectionModuleRequest) (*models.DeleteServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceConnectionModule(ctx, request)
}
func (service *BaseService) GetServiceConnectionModule(ctx context.Context, request *models.GetServiceConnectionModuleRequest) (*models.GetServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceConnectionModule(ctx, request)
}
func (service *BaseService) ListServiceConnectionModule(ctx context.Context, request *models.ListServiceConnectionModuleRequest) (*models.ListServiceConnectionModuleResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceConnectionModule(ctx, request)
}

func (service *BaseService) CreateServiceEndpoint(ctx context.Context, request *models.CreateServiceEndpointRequest) (*models.CreateServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceEndpoint(ctx, request)
}
func (service *BaseService) UpdateServiceEndpoint(ctx context.Context, request *models.UpdateServiceEndpointRequest) (*models.UpdateServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceEndpoint(ctx, request)
}
func (service *BaseService) DeleteServiceEndpoint(ctx context.Context, request *models.DeleteServiceEndpointRequest) (*models.DeleteServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceEndpoint(ctx, request)
}
func (service *BaseService) GetServiceEndpoint(ctx context.Context, request *models.GetServiceEndpointRequest) (*models.GetServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceEndpoint(ctx, request)
}
func (service *BaseService) ListServiceEndpoint(ctx context.Context, request *models.ListServiceEndpointRequest) (*models.ListServiceEndpointResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceEndpoint(ctx, request)
}

func (service *BaseService) CreateServiceGroup(ctx context.Context, request *models.CreateServiceGroupRequest) (*models.CreateServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceGroup(ctx, request)
}
func (service *BaseService) UpdateServiceGroup(ctx context.Context, request *models.UpdateServiceGroupRequest) (*models.UpdateServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceGroup(ctx, request)
}
func (service *BaseService) DeleteServiceGroup(ctx context.Context, request *models.DeleteServiceGroupRequest) (*models.DeleteServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceGroup(ctx, request)
}
func (service *BaseService) GetServiceGroup(ctx context.Context, request *models.GetServiceGroupRequest) (*models.GetServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceGroup(ctx, request)
}
func (service *BaseService) ListServiceGroup(ctx context.Context, request *models.ListServiceGroupRequest) (*models.ListServiceGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceGroup(ctx, request)
}

func (service *BaseService) CreateServiceHealthCheck(ctx context.Context, request *models.CreateServiceHealthCheckRequest) (*models.CreateServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceHealthCheck(ctx, request)
}
func (service *BaseService) UpdateServiceHealthCheck(ctx context.Context, request *models.UpdateServiceHealthCheckRequest) (*models.UpdateServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceHealthCheck(ctx, request)
}
func (service *BaseService) DeleteServiceHealthCheck(ctx context.Context, request *models.DeleteServiceHealthCheckRequest) (*models.DeleteServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceHealthCheck(ctx, request)
}
func (service *BaseService) GetServiceHealthCheck(ctx context.Context, request *models.GetServiceHealthCheckRequest) (*models.GetServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceHealthCheck(ctx, request)
}
func (service *BaseService) ListServiceHealthCheck(ctx context.Context, request *models.ListServiceHealthCheckRequest) (*models.ListServiceHealthCheckResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceHealthCheck(ctx, request)
}

func (service *BaseService) CreateServiceInstance(ctx context.Context, request *models.CreateServiceInstanceRequest) (*models.CreateServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceInstance(ctx, request)
}
func (service *BaseService) UpdateServiceInstance(ctx context.Context, request *models.UpdateServiceInstanceRequest) (*models.UpdateServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceInstance(ctx, request)
}
func (service *BaseService) DeleteServiceInstance(ctx context.Context, request *models.DeleteServiceInstanceRequest) (*models.DeleteServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceInstance(ctx, request)
}
func (service *BaseService) GetServiceInstance(ctx context.Context, request *models.GetServiceInstanceRequest) (*models.GetServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceInstance(ctx, request)
}
func (service *BaseService) ListServiceInstance(ctx context.Context, request *models.ListServiceInstanceRequest) (*models.ListServiceInstanceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceInstance(ctx, request)
}

func (service *BaseService) CreateServiceObject(ctx context.Context, request *models.CreateServiceObjectRequest) (*models.CreateServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceObject(ctx, request)
}
func (service *BaseService) UpdateServiceObject(ctx context.Context, request *models.UpdateServiceObjectRequest) (*models.UpdateServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceObject(ctx, request)
}
func (service *BaseService) DeleteServiceObject(ctx context.Context, request *models.DeleteServiceObjectRequest) (*models.DeleteServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceObject(ctx, request)
}
func (service *BaseService) GetServiceObject(ctx context.Context, request *models.GetServiceObjectRequest) (*models.GetServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceObject(ctx, request)
}
func (service *BaseService) ListServiceObject(ctx context.Context, request *models.ListServiceObjectRequest) (*models.ListServiceObjectResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceObject(ctx, request)
}

func (service *BaseService) CreateServiceTemplate(ctx context.Context, request *models.CreateServiceTemplateRequest) (*models.CreateServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServiceTemplate(ctx, request)
}
func (service *BaseService) UpdateServiceTemplate(ctx context.Context, request *models.UpdateServiceTemplateRequest) (*models.UpdateServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceTemplate(ctx, request)
}
func (service *BaseService) DeleteServiceTemplate(ctx context.Context, request *models.DeleteServiceTemplateRequest) (*models.DeleteServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceTemplate(ctx, request)
}
func (service *BaseService) GetServiceTemplate(ctx context.Context, request *models.GetServiceTemplateRequest) (*models.GetServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServiceTemplate(ctx, request)
}
func (service *BaseService) ListServiceTemplate(ctx context.Context, request *models.ListServiceTemplateRequest) (*models.ListServiceTemplateResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServiceTemplate(ctx, request)
}

func (service *BaseService) CreateSubnet(ctx context.Context, request *models.CreateSubnetRequest) (*models.CreateSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateSubnet(ctx, request)
}
func (service *BaseService) UpdateSubnet(ctx context.Context, request *models.UpdateSubnetRequest) (*models.UpdateSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateSubnet(ctx, request)
}
func (service *BaseService) DeleteSubnet(ctx context.Context, request *models.DeleteSubnetRequest) (*models.DeleteSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteSubnet(ctx, request)
}
func (service *BaseService) GetSubnet(ctx context.Context, request *models.GetSubnetRequest) (*models.GetSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetSubnet(ctx, request)
}
func (service *BaseService) ListSubnet(ctx context.Context, request *models.ListSubnetRequest) (*models.ListSubnetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListSubnet(ctx, request)
}

func (service *BaseService) CreateTag(ctx context.Context, request *models.CreateTagRequest) (*models.CreateTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateTag(ctx, request)
}
func (service *BaseService) UpdateTag(ctx context.Context, request *models.UpdateTagRequest) (*models.UpdateTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateTag(ctx, request)
}
func (service *BaseService) DeleteTag(ctx context.Context, request *models.DeleteTagRequest) (*models.DeleteTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteTag(ctx, request)
}
func (service *BaseService) GetTag(ctx context.Context, request *models.GetTagRequest) (*models.GetTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetTag(ctx, request)
}
func (service *BaseService) ListTag(ctx context.Context, request *models.ListTagRequest) (*models.ListTagResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListTag(ctx, request)
}

func (service *BaseService) CreateTagType(ctx context.Context, request *models.CreateTagTypeRequest) (*models.CreateTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateTagType(ctx, request)
}
func (service *BaseService) UpdateTagType(ctx context.Context, request *models.UpdateTagTypeRequest) (*models.UpdateTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateTagType(ctx, request)
}
func (service *BaseService) DeleteTagType(ctx context.Context, request *models.DeleteTagTypeRequest) (*models.DeleteTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteTagType(ctx, request)
}
func (service *BaseService) GetTagType(ctx context.Context, request *models.GetTagTypeRequest) (*models.GetTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetTagType(ctx, request)
}
func (service *BaseService) ListTagType(ctx context.Context, request *models.ListTagTypeRequest) (*models.ListTagTypeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListTagType(ctx, request)
}

func (service *BaseService) CreateUser(ctx context.Context, request *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateUser(ctx, request)
}
func (service *BaseService) UpdateUser(ctx context.Context, request *models.UpdateUserRequest) (*models.UpdateUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateUser(ctx, request)
}
func (service *BaseService) DeleteUser(ctx context.Context, request *models.DeleteUserRequest) (*models.DeleteUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteUser(ctx, request)
}
func (service *BaseService) GetUser(ctx context.Context, request *models.GetUserRequest) (*models.GetUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetUser(ctx, request)
}
func (service *BaseService) ListUser(ctx context.Context, request *models.ListUserRequest) (*models.ListUserResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListUser(ctx, request)
}

func (service *BaseService) CreateVirtualDNSRecord(ctx context.Context, request *models.CreateVirtualDNSRecordRequest) (*models.CreateVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualDNSRecord(ctx, request)
}
func (service *BaseService) UpdateVirtualDNSRecord(ctx context.Context, request *models.UpdateVirtualDNSRecordRequest) (*models.UpdateVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualDNSRecord(ctx, request)
}
func (service *BaseService) DeleteVirtualDNSRecord(ctx context.Context, request *models.DeleteVirtualDNSRecordRequest) (*models.DeleteVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualDNSRecord(ctx, request)
}
func (service *BaseService) GetVirtualDNSRecord(ctx context.Context, request *models.GetVirtualDNSRecordRequest) (*models.GetVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualDNSRecord(ctx, request)
}
func (service *BaseService) ListVirtualDNSRecord(ctx context.Context, request *models.ListVirtualDNSRecordRequest) (*models.ListVirtualDNSRecordResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualDNSRecord(ctx, request)
}

func (service *BaseService) CreateVirtualDNS(ctx context.Context, request *models.CreateVirtualDNSRequest) (*models.CreateVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualDNS(ctx, request)
}
func (service *BaseService) UpdateVirtualDNS(ctx context.Context, request *models.UpdateVirtualDNSRequest) (*models.UpdateVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualDNS(ctx, request)
}
func (service *BaseService) DeleteVirtualDNS(ctx context.Context, request *models.DeleteVirtualDNSRequest) (*models.DeleteVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualDNS(ctx, request)
}
func (service *BaseService) GetVirtualDNS(ctx context.Context, request *models.GetVirtualDNSRequest) (*models.GetVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualDNS(ctx, request)
}
func (service *BaseService) ListVirtualDNS(ctx context.Context, request *models.ListVirtualDNSRequest) (*models.ListVirtualDNSResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualDNS(ctx, request)
}

func (service *BaseService) CreateVirtualIP(ctx context.Context, request *models.CreateVirtualIPRequest) (*models.CreateVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualIP(ctx, request)
}
func (service *BaseService) UpdateVirtualIP(ctx context.Context, request *models.UpdateVirtualIPRequest) (*models.UpdateVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualIP(ctx, request)
}
func (service *BaseService) DeleteVirtualIP(ctx context.Context, request *models.DeleteVirtualIPRequest) (*models.DeleteVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualIP(ctx, request)
}
func (service *BaseService) GetVirtualIP(ctx context.Context, request *models.GetVirtualIPRequest) (*models.GetVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualIP(ctx, request)
}
func (service *BaseService) ListVirtualIP(ctx context.Context, request *models.ListVirtualIPRequest) (*models.ListVirtualIPResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualIP(ctx, request)
}

func (service *BaseService) CreateVirtualMachineInterface(ctx context.Context, request *models.CreateVirtualMachineInterfaceRequest) (*models.CreateVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualMachineInterface(ctx, request)
}
func (service *BaseService) UpdateVirtualMachineInterface(ctx context.Context, request *models.UpdateVirtualMachineInterfaceRequest) (*models.UpdateVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualMachineInterface(ctx, request)
}
func (service *BaseService) DeleteVirtualMachineInterface(ctx context.Context, request *models.DeleteVirtualMachineInterfaceRequest) (*models.DeleteVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualMachineInterface(ctx, request)
}
func (service *BaseService) GetVirtualMachineInterface(ctx context.Context, request *models.GetVirtualMachineInterfaceRequest) (*models.GetVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualMachineInterface(ctx, request)
}
func (service *BaseService) ListVirtualMachineInterface(ctx context.Context, request *models.ListVirtualMachineInterfaceRequest) (*models.ListVirtualMachineInterfaceResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualMachineInterface(ctx, request)
}

func (service *BaseService) CreateVirtualMachine(ctx context.Context, request *models.CreateVirtualMachineRequest) (*models.CreateVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualMachine(ctx, request)
}
func (service *BaseService) UpdateVirtualMachine(ctx context.Context, request *models.UpdateVirtualMachineRequest) (*models.UpdateVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualMachine(ctx, request)
}
func (service *BaseService) DeleteVirtualMachine(ctx context.Context, request *models.DeleteVirtualMachineRequest) (*models.DeleteVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualMachine(ctx, request)
}
func (service *BaseService) GetVirtualMachine(ctx context.Context, request *models.GetVirtualMachineRequest) (*models.GetVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualMachine(ctx, request)
}
func (service *BaseService) ListVirtualMachine(ctx context.Context, request *models.ListVirtualMachineRequest) (*models.ListVirtualMachineResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualMachine(ctx, request)
}

func (service *BaseService) CreateVirtualNetwork(ctx context.Context, request *models.CreateVirtualNetworkRequest) (*models.CreateVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualNetwork(ctx, request)
}
func (service *BaseService) UpdateVirtualNetwork(ctx context.Context, request *models.UpdateVirtualNetworkRequest) (*models.UpdateVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualNetwork(ctx, request)
}
func (service *BaseService) DeleteVirtualNetwork(ctx context.Context, request *models.DeleteVirtualNetworkRequest) (*models.DeleteVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualNetwork(ctx, request)
}
func (service *BaseService) GetVirtualNetwork(ctx context.Context, request *models.GetVirtualNetworkRequest) (*models.GetVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualNetwork(ctx, request)
}
func (service *BaseService) ListVirtualNetwork(ctx context.Context, request *models.ListVirtualNetworkRequest) (*models.ListVirtualNetworkResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualNetwork(ctx, request)
}

func (service *BaseService) CreateVirtualRouter(ctx context.Context, request *models.CreateVirtualRouterRequest) (*models.CreateVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualRouter(ctx, request)
}
func (service *BaseService) UpdateVirtualRouter(ctx context.Context, request *models.UpdateVirtualRouterRequest) (*models.UpdateVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualRouter(ctx, request)
}
func (service *BaseService) DeleteVirtualRouter(ctx context.Context, request *models.DeleteVirtualRouterRequest) (*models.DeleteVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualRouter(ctx, request)
}
func (service *BaseService) GetVirtualRouter(ctx context.Context, request *models.GetVirtualRouterRequest) (*models.GetVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVirtualRouter(ctx, request)
}
func (service *BaseService) ListVirtualRouter(ctx context.Context, request *models.ListVirtualRouterRequest) (*models.ListVirtualRouterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVirtualRouter(ctx, request)
}

func (service *BaseService) CreateAppformixNode(ctx context.Context, request *models.CreateAppformixNodeRequest) (*models.CreateAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateAppformixNode(ctx, request)
}
func (service *BaseService) UpdateAppformixNode(ctx context.Context, request *models.UpdateAppformixNodeRequest) (*models.UpdateAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateAppformixNode(ctx, request)
}
func (service *BaseService) DeleteAppformixNode(ctx context.Context, request *models.DeleteAppformixNodeRequest) (*models.DeleteAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteAppformixNode(ctx, request)
}
func (service *BaseService) GetAppformixNode(ctx context.Context, request *models.GetAppformixNodeRequest) (*models.GetAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetAppformixNode(ctx, request)
}
func (service *BaseService) ListAppformixNode(ctx context.Context, request *models.ListAppformixNodeRequest) (*models.ListAppformixNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListAppformixNode(ctx, request)
}

func (service *BaseService) CreateBaremetalNode(ctx context.Context, request *models.CreateBaremetalNodeRequest) (*models.CreateBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBaremetalNode(ctx, request)
}
func (service *BaseService) UpdateBaremetalNode(ctx context.Context, request *models.UpdateBaremetalNodeRequest) (*models.UpdateBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBaremetalNode(ctx, request)
}
func (service *BaseService) DeleteBaremetalNode(ctx context.Context, request *models.DeleteBaremetalNodeRequest) (*models.DeleteBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBaremetalNode(ctx, request)
}
func (service *BaseService) GetBaremetalNode(ctx context.Context, request *models.GetBaremetalNodeRequest) (*models.GetBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBaremetalNode(ctx, request)
}
func (service *BaseService) ListBaremetalNode(ctx context.Context, request *models.ListBaremetalNodeRequest) (*models.ListBaremetalNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBaremetalNode(ctx, request)
}

func (service *BaseService) CreateBaremetalPort(ctx context.Context, request *models.CreateBaremetalPortRequest) (*models.CreateBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateBaremetalPort(ctx, request)
}
func (service *BaseService) UpdateBaremetalPort(ctx context.Context, request *models.UpdateBaremetalPortRequest) (*models.UpdateBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateBaremetalPort(ctx, request)
}
func (service *BaseService) DeleteBaremetalPort(ctx context.Context, request *models.DeleteBaremetalPortRequest) (*models.DeleteBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteBaremetalPort(ctx, request)
}
func (service *BaseService) GetBaremetalPort(ctx context.Context, request *models.GetBaremetalPortRequest) (*models.GetBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetBaremetalPort(ctx, request)
}
func (service *BaseService) ListBaremetalPort(ctx context.Context, request *models.ListBaremetalPortRequest) (*models.ListBaremetalPortResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListBaremetalPort(ctx, request)
}

func (service *BaseService) CreateContrailAnalyticsDatabaseNode(ctx context.Context, request *models.CreateContrailAnalyticsDatabaseNodeRequest) (*models.CreateContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) UpdateContrailAnalyticsDatabaseNode(ctx context.Context, request *models.UpdateContrailAnalyticsDatabaseNodeRequest) (*models.UpdateContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) DeleteContrailAnalyticsDatabaseNode(ctx context.Context, request *models.DeleteContrailAnalyticsDatabaseNodeRequest) (*models.DeleteContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) GetContrailAnalyticsDatabaseNode(ctx context.Context, request *models.GetContrailAnalyticsDatabaseNodeRequest) (*models.GetContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailAnalyticsDatabaseNode(ctx, request)
}
func (service *BaseService) ListContrailAnalyticsDatabaseNode(ctx context.Context, request *models.ListContrailAnalyticsDatabaseNodeRequest) (*models.ListContrailAnalyticsDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailAnalyticsDatabaseNode(ctx, request)
}

func (service *BaseService) CreateContrailAnalyticsNode(ctx context.Context, request *models.CreateContrailAnalyticsNodeRequest) (*models.CreateContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) UpdateContrailAnalyticsNode(ctx context.Context, request *models.UpdateContrailAnalyticsNodeRequest) (*models.UpdateContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) DeleteContrailAnalyticsNode(ctx context.Context, request *models.DeleteContrailAnalyticsNodeRequest) (*models.DeleteContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) GetContrailAnalyticsNode(ctx context.Context, request *models.GetContrailAnalyticsNodeRequest) (*models.GetContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailAnalyticsNode(ctx, request)
}
func (service *BaseService) ListContrailAnalyticsNode(ctx context.Context, request *models.ListContrailAnalyticsNodeRequest) (*models.ListContrailAnalyticsNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailAnalyticsNode(ctx, request)
}

func (service *BaseService) CreateContrailCluster(ctx context.Context, request *models.CreateContrailClusterRequest) (*models.CreateContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailCluster(ctx, request)
}
func (service *BaseService) UpdateContrailCluster(ctx context.Context, request *models.UpdateContrailClusterRequest) (*models.UpdateContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailCluster(ctx, request)
}
func (service *BaseService) DeleteContrailCluster(ctx context.Context, request *models.DeleteContrailClusterRequest) (*models.DeleteContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailCluster(ctx, request)
}
func (service *BaseService) GetContrailCluster(ctx context.Context, request *models.GetContrailClusterRequest) (*models.GetContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailCluster(ctx, request)
}
func (service *BaseService) ListContrailCluster(ctx context.Context, request *models.ListContrailClusterRequest) (*models.ListContrailClusterResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailCluster(ctx, request)
}

func (service *BaseService) CreateContrailConfigDatabaseNode(ctx context.Context, request *models.CreateContrailConfigDatabaseNodeRequest) (*models.CreateContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) UpdateContrailConfigDatabaseNode(ctx context.Context, request *models.UpdateContrailConfigDatabaseNodeRequest) (*models.UpdateContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) DeleteContrailConfigDatabaseNode(ctx context.Context, request *models.DeleteContrailConfigDatabaseNodeRequest) (*models.DeleteContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) GetContrailConfigDatabaseNode(ctx context.Context, request *models.GetContrailConfigDatabaseNodeRequest) (*models.GetContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailConfigDatabaseNode(ctx, request)
}
func (service *BaseService) ListContrailConfigDatabaseNode(ctx context.Context, request *models.ListContrailConfigDatabaseNodeRequest) (*models.ListContrailConfigDatabaseNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailConfigDatabaseNode(ctx, request)
}

func (service *BaseService) CreateContrailConfigNode(ctx context.Context, request *models.CreateContrailConfigNodeRequest) (*models.CreateContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailConfigNode(ctx, request)
}
func (service *BaseService) UpdateContrailConfigNode(ctx context.Context, request *models.UpdateContrailConfigNodeRequest) (*models.UpdateContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailConfigNode(ctx, request)
}
func (service *BaseService) DeleteContrailConfigNode(ctx context.Context, request *models.DeleteContrailConfigNodeRequest) (*models.DeleteContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailConfigNode(ctx, request)
}
func (service *BaseService) GetContrailConfigNode(ctx context.Context, request *models.GetContrailConfigNodeRequest) (*models.GetContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailConfigNode(ctx, request)
}
func (service *BaseService) ListContrailConfigNode(ctx context.Context, request *models.ListContrailConfigNodeRequest) (*models.ListContrailConfigNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailConfigNode(ctx, request)
}

func (service *BaseService) CreateContrailControlNode(ctx context.Context, request *models.CreateContrailControlNodeRequest) (*models.CreateContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailControlNode(ctx, request)
}
func (service *BaseService) UpdateContrailControlNode(ctx context.Context, request *models.UpdateContrailControlNodeRequest) (*models.UpdateContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailControlNode(ctx, request)
}
func (service *BaseService) DeleteContrailControlNode(ctx context.Context, request *models.DeleteContrailControlNodeRequest) (*models.DeleteContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailControlNode(ctx, request)
}
func (service *BaseService) GetContrailControlNode(ctx context.Context, request *models.GetContrailControlNodeRequest) (*models.GetContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailControlNode(ctx, request)
}
func (service *BaseService) ListContrailControlNode(ctx context.Context, request *models.ListContrailControlNodeRequest) (*models.ListContrailControlNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailControlNode(ctx, request)
}

func (service *BaseService) CreateContrailStorageNode(ctx context.Context, request *models.CreateContrailStorageNodeRequest) (*models.CreateContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailStorageNode(ctx, request)
}
func (service *BaseService) UpdateContrailStorageNode(ctx context.Context, request *models.UpdateContrailStorageNodeRequest) (*models.UpdateContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailStorageNode(ctx, request)
}
func (service *BaseService) DeleteContrailStorageNode(ctx context.Context, request *models.DeleteContrailStorageNodeRequest) (*models.DeleteContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailStorageNode(ctx, request)
}
func (service *BaseService) GetContrailStorageNode(ctx context.Context, request *models.GetContrailStorageNodeRequest) (*models.GetContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailStorageNode(ctx, request)
}
func (service *BaseService) ListContrailStorageNode(ctx context.Context, request *models.ListContrailStorageNodeRequest) (*models.ListContrailStorageNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailStorageNode(ctx, request)
}

func (service *BaseService) CreateContrailVrouterNode(ctx context.Context, request *models.CreateContrailVrouterNodeRequest) (*models.CreateContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailVrouterNode(ctx, request)
}
func (service *BaseService) UpdateContrailVrouterNode(ctx context.Context, request *models.UpdateContrailVrouterNodeRequest) (*models.UpdateContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailVrouterNode(ctx, request)
}
func (service *BaseService) DeleteContrailVrouterNode(ctx context.Context, request *models.DeleteContrailVrouterNodeRequest) (*models.DeleteContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailVrouterNode(ctx, request)
}
func (service *BaseService) GetContrailVrouterNode(ctx context.Context, request *models.GetContrailVrouterNodeRequest) (*models.GetContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailVrouterNode(ctx, request)
}
func (service *BaseService) ListContrailVrouterNode(ctx context.Context, request *models.ListContrailVrouterNodeRequest) (*models.ListContrailVrouterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailVrouterNode(ctx, request)
}

func (service *BaseService) CreateContrailControllerNode(ctx context.Context, request *models.CreateContrailControllerNodeRequest) (*models.CreateContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateContrailControllerNode(ctx, request)
}
func (service *BaseService) UpdateContrailControllerNode(ctx context.Context, request *models.UpdateContrailControllerNodeRequest) (*models.UpdateContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailControllerNode(ctx, request)
}
func (service *BaseService) DeleteContrailControllerNode(ctx context.Context, request *models.DeleteContrailControllerNodeRequest) (*models.DeleteContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailControllerNode(ctx, request)
}
func (service *BaseService) GetContrailControllerNode(ctx context.Context, request *models.GetContrailControllerNodeRequest) (*models.GetContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetContrailControllerNode(ctx, request)
}
func (service *BaseService) ListContrailControllerNode(ctx context.Context, request *models.ListContrailControllerNodeRequest) (*models.ListContrailControllerNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListContrailControllerNode(ctx, request)
}

func (service *BaseService) CreateDashboard(ctx context.Context, request *models.CreateDashboardRequest) (*models.CreateDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateDashboard(ctx, request)
}
func (service *BaseService) UpdateDashboard(ctx context.Context, request *models.UpdateDashboardRequest) (*models.UpdateDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateDashboard(ctx, request)
}
func (service *BaseService) DeleteDashboard(ctx context.Context, request *models.DeleteDashboardRequest) (*models.DeleteDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteDashboard(ctx, request)
}
func (service *BaseService) GetDashboard(ctx context.Context, request *models.GetDashboardRequest) (*models.GetDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetDashboard(ctx, request)
}
func (service *BaseService) ListDashboard(ctx context.Context, request *models.ListDashboardRequest) (*models.ListDashboardResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListDashboard(ctx, request)
}

func (service *BaseService) CreateFlavor(ctx context.Context, request *models.CreateFlavorRequest) (*models.CreateFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateFlavor(ctx, request)
}
func (service *BaseService) UpdateFlavor(ctx context.Context, request *models.UpdateFlavorRequest) (*models.UpdateFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateFlavor(ctx, request)
}
func (service *BaseService) DeleteFlavor(ctx context.Context, request *models.DeleteFlavorRequest) (*models.DeleteFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteFlavor(ctx, request)
}
func (service *BaseService) GetFlavor(ctx context.Context, request *models.GetFlavorRequest) (*models.GetFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetFlavor(ctx, request)
}
func (service *BaseService) ListFlavor(ctx context.Context, request *models.ListFlavorRequest) (*models.ListFlavorResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListFlavor(ctx, request)
}

func (service *BaseService) CreateOsImage(ctx context.Context, request *models.CreateOsImageRequest) (*models.CreateOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateOsImage(ctx, request)
}
func (service *BaseService) UpdateOsImage(ctx context.Context, request *models.UpdateOsImageRequest) (*models.UpdateOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateOsImage(ctx, request)
}
func (service *BaseService) DeleteOsImage(ctx context.Context, request *models.DeleteOsImageRequest) (*models.DeleteOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteOsImage(ctx, request)
}
func (service *BaseService) GetOsImage(ctx context.Context, request *models.GetOsImageRequest) (*models.GetOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetOsImage(ctx, request)
}
func (service *BaseService) ListOsImage(ctx context.Context, request *models.ListOsImageRequest) (*models.ListOsImageResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListOsImage(ctx, request)
}

func (service *BaseService) CreateKeypair(ctx context.Context, request *models.CreateKeypairRequest) (*models.CreateKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateKeypair(ctx, request)
}
func (service *BaseService) UpdateKeypair(ctx context.Context, request *models.UpdateKeypairRequest) (*models.UpdateKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateKeypair(ctx, request)
}
func (service *BaseService) DeleteKeypair(ctx context.Context, request *models.DeleteKeypairRequest) (*models.DeleteKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteKeypair(ctx, request)
}
func (service *BaseService) GetKeypair(ctx context.Context, request *models.GetKeypairRequest) (*models.GetKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetKeypair(ctx, request)
}
func (service *BaseService) ListKeypair(ctx context.Context, request *models.ListKeypairRequest) (*models.ListKeypairResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListKeypair(ctx, request)
}

func (service *BaseService) CreateKubernetesMasterNode(ctx context.Context, request *models.CreateKubernetesMasterNodeRequest) (*models.CreateKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesMasterNode(ctx, request)
}
func (service *BaseService) UpdateKubernetesMasterNode(ctx context.Context, request *models.UpdateKubernetesMasterNodeRequest) (*models.UpdateKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesMasterNode(ctx, request)
}
func (service *BaseService) DeleteKubernetesMasterNode(ctx context.Context, request *models.DeleteKubernetesMasterNodeRequest) (*models.DeleteKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesMasterNode(ctx, request)
}
func (service *BaseService) GetKubernetesMasterNode(ctx context.Context, request *models.GetKubernetesMasterNodeRequest) (*models.GetKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetKubernetesMasterNode(ctx, request)
}
func (service *BaseService) ListKubernetesMasterNode(ctx context.Context, request *models.ListKubernetesMasterNodeRequest) (*models.ListKubernetesMasterNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListKubernetesMasterNode(ctx, request)
}

func (service *BaseService) CreateKubernetesNode(ctx context.Context, request *models.CreateKubernetesNodeRequest) (*models.CreateKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesNode(ctx, request)
}
func (service *BaseService) UpdateKubernetesNode(ctx context.Context, request *models.UpdateKubernetesNodeRequest) (*models.UpdateKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesNode(ctx, request)
}
func (service *BaseService) DeleteKubernetesNode(ctx context.Context, request *models.DeleteKubernetesNodeRequest) (*models.DeleteKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesNode(ctx, request)
}
func (service *BaseService) GetKubernetesNode(ctx context.Context, request *models.GetKubernetesNodeRequest) (*models.GetKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetKubernetesNode(ctx, request)
}
func (service *BaseService) ListKubernetesNode(ctx context.Context, request *models.ListKubernetesNodeRequest) (*models.ListKubernetesNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListKubernetesNode(ctx, request)
}

func (service *BaseService) CreateLocation(ctx context.Context, request *models.CreateLocationRequest) (*models.CreateLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateLocation(ctx, request)
}
func (service *BaseService) UpdateLocation(ctx context.Context, request *models.UpdateLocationRequest) (*models.UpdateLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateLocation(ctx, request)
}
func (service *BaseService) DeleteLocation(ctx context.Context, request *models.DeleteLocationRequest) (*models.DeleteLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteLocation(ctx, request)
}
func (service *BaseService) GetLocation(ctx context.Context, request *models.GetLocationRequest) (*models.GetLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetLocation(ctx, request)
}
func (service *BaseService) ListLocation(ctx context.Context, request *models.ListLocationRequest) (*models.ListLocationResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListLocation(ctx, request)
}

func (service *BaseService) CreateNode(ctx context.Context, request *models.CreateNodeRequest) (*models.CreateNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateNode(ctx, request)
}
func (service *BaseService) UpdateNode(ctx context.Context, request *models.UpdateNodeRequest) (*models.UpdateNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateNode(ctx, request)
}
func (service *BaseService) DeleteNode(ctx context.Context, request *models.DeleteNodeRequest) (*models.DeleteNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteNode(ctx, request)
}
func (service *BaseService) GetNode(ctx context.Context, request *models.GetNodeRequest) (*models.GetNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetNode(ctx, request)
}
func (service *BaseService) ListNode(ctx context.Context, request *models.ListNodeRequest) (*models.ListNodeResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListNode(ctx, request)
}

func (service *BaseService) CreateServer(ctx context.Context, request *models.CreateServerRequest) (*models.CreateServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateServer(ctx, request)
}
func (service *BaseService) UpdateServer(ctx context.Context, request *models.UpdateServerRequest) (*models.UpdateServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateServer(ctx, request)
}
func (service *BaseService) DeleteServer(ctx context.Context, request *models.DeleteServerRequest) (*models.DeleteServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteServer(ctx, request)
}
func (service *BaseService) GetServer(ctx context.Context, request *models.GetServerRequest) (*models.GetServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetServer(ctx, request)
}
func (service *BaseService) ListServer(ctx context.Context, request *models.ListServerRequest) (*models.ListServerResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListServer(ctx, request)
}

func (service *BaseService) CreateVPNGroup(ctx context.Context, request *models.CreateVPNGroupRequest) (*models.CreateVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateVPNGroup(ctx, request)
}
func (service *BaseService) UpdateVPNGroup(ctx context.Context, request *models.UpdateVPNGroupRequest) (*models.UpdateVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateVPNGroup(ctx, request)
}
func (service *BaseService) DeleteVPNGroup(ctx context.Context, request *models.DeleteVPNGroupRequest) (*models.DeleteVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteVPNGroup(ctx, request)
}
func (service *BaseService) GetVPNGroup(ctx context.Context, request *models.GetVPNGroupRequest) (*models.GetVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetVPNGroup(ctx, request)
}
func (service *BaseService) ListVPNGroup(ctx context.Context, request *models.ListVPNGroupRequest) (*models.ListVPNGroupResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListVPNGroup(ctx, request)
}

func (service *BaseService) CreateWidget(ctx context.Context, request *models.CreateWidgetRequest) (*models.CreateWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().CreateWidget(ctx, request)
}
func (service *BaseService) UpdateWidget(ctx context.Context, request *models.UpdateWidgetRequest) (*models.UpdateWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().UpdateWidget(ctx, request)
}
func (service *BaseService) DeleteWidget(ctx context.Context, request *models.DeleteWidgetRequest) (*models.DeleteWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().DeleteWidget(ctx, request)
}
func (service *BaseService) GetWidget(ctx context.Context, request *models.GetWidgetRequest) (*models.GetWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().GetWidget(ctx, request)
}
func (service *BaseService) ListWidget(ctx context.Context, request *models.ListWidgetRequest) (*models.ListWidgetResponse, error) {
	if service.next == nil {
		return nil, nil
	}
	return service.Next().ListWidget(ctx, request)
}
