package api

import "github.com/Juniper/contrail/pkg/common"

var APIs = []common.RESTAPI{

	&AccessControlListRESTAPI{},

	&AddressGroupRESTAPI{},

	&AlarmRESTAPI{},

	&AliasIPPoolRESTAPI{},

	&AliasIPRESTAPI{},

	&AnalyticsNodeRESTAPI{},

	&APIAccessListRESTAPI{},

	&ApplicationPolicySetRESTAPI{},

	&BGPAsAServiceRESTAPI{},

	&BGPRouterRESTAPI{},

	&BGPVPNRESTAPI{},

	&BridgeDomainRESTAPI{},

	&ConfigNodeRESTAPI{},

	&ConfigRootRESTAPI{},

	&CustomerAttachmentRESTAPI{},

	&DatabaseNodeRESTAPI{},

	&DiscoveryServiceAssignmentRESTAPI{},

	&DomainRESTAPI{},

	&DsaRuleRESTAPI{},

	&E2ServiceProviderRESTAPI{},

	&FirewallPolicyRESTAPI{},

	&FirewallRuleRESTAPI{},

	&FloatingIPPoolRESTAPI{},

	&FloatingIPRESTAPI{},

	&ForwardingClassRESTAPI{},

	&GlobalQosConfigRESTAPI{},

	&GlobalSystemConfigRESTAPI{},

	&GlobalVrouterConfigRESTAPI{},

	&InstanceIPRESTAPI{},

	&InterfaceRouteTableRESTAPI{},

	&LoadbalancerHealthmonitorRESTAPI{},

	&LoadbalancerListenerRESTAPI{},

	&LoadbalancerMemberRESTAPI{},

	&LoadbalancerPoolRESTAPI{},

	&LoadbalancerRESTAPI{},

	&LogicalInterfaceRESTAPI{},

	&LogicalRouterRESTAPI{},

	&NamespaceRESTAPI{},

	&NetworkDeviceConfigRESTAPI{},

	&NetworkIpamRESTAPI{},

	&NetworkPolicyRESTAPI{},

	&PeeringPolicyRESTAPI{},

	&PhysicalInterfaceRESTAPI{},

	&PhysicalRouterRESTAPI{},

	&PolicyManagementRESTAPI{},

	&PortTupleRESTAPI{},

	&ProjectRESTAPI{},

	&ProviderAttachmentRESTAPI{},

	&QosConfigRESTAPI{},

	&QosQueueRESTAPI{},

	&RouteAggregateRESTAPI{},

	&RouteTableRESTAPI{},

	&RouteTargetRESTAPI{},

	&RoutingInstanceRESTAPI{},

	&RoutingPolicyRESTAPI{},

	&SecurityGroupRESTAPI{},

	&SecurityLoggingObjectRESTAPI{},

	&ServiceApplianceRESTAPI{},

	&ServiceApplianceSetRESTAPI{},

	&ServiceConnectionModuleRESTAPI{},

	&ServiceEndpointRESTAPI{},

	&ServiceGroupRESTAPI{},

	&ServiceHealthCheckRESTAPI{},

	&ServiceInstanceRESTAPI{},

	&ServiceObjectRESTAPI{},

	&ServiceTemplateRESTAPI{},

	&SubnetRESTAPI{},

	&TagRESTAPI{},

	&TagTypeRESTAPI{},

	&UserRESTAPI{},

	&VirtualDNSRecordRESTAPI{},

	&VirtualDNSRESTAPI{},

	&VirtualIPRESTAPI{},

	&VirtualMachineInterfaceRESTAPI{},

	&VirtualMachineRESTAPI{},

	&VirtualNetworkRESTAPI{},

	&VirtualRouterRESTAPI{},

	&AppformixNodeRoleRESTAPI{},

	&ContrailAnalyticsDatabaseNodeRoleRESTAPI{},

	&ContrailAnalyticsNodeRESTAPI{},

	&ContrailClusterRESTAPI{},

	&ContrailControllerNodeRoleRESTAPI{},

	&ControllerNodeRoleRESTAPI{},

	&DashboardRESTAPI{},

	&KubernetesClusterRESTAPI{},

	&KubernetesNodeRESTAPI{},

	&LocationRESTAPI{},

	&NodeRESTAPI{},

	&OpenstackClusterRESTAPI{},

	&OpenstackComputeNodeRoleRESTAPI{},

	&OpenstackStorageNodeRoleRESTAPI{},

	&VPNGroupRESTAPI{},

	&WidgetRESTAPI{},
}
