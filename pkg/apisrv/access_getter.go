package apisrv

import (
	"context"

	"github.com/Juniper/asf/pkg/rbac"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

type contrailAccessGetter struct {
	service *services.ContrailService
}

func (ag contrailAccessGetter) GetAPIAccessLists(ctx context.Context) []*rbac.APIAccessList {
	listRequest := &services.ListAPIAccessListRequest{
		Spec: &baseservices.ListSpec{},
	}
	if result, err := ag.service.ListAPIAccessList(ctx, listRequest); err == nil {
		return models.APIAccessLists(result.APIAccessLists).ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAccessControlListPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAccessControlListRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAccessControlList(ctx, listRequest); err == nil {
		return result.AccessControlList.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAddressGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAddressGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAddressGroup(ctx, listRequest); err == nil {
		return result.AddressGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAlarmPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAlarmRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAlarm(ctx, listRequest); err == nil {
		return result.Alarm.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAliasIPPoolPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAliasIPPoolRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAliasIPPool(ctx, listRequest); err == nil {
		return result.AliasIPPool.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAliasIPPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAliasIPRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAliasIP(ctx, listRequest); err == nil {
		return result.AliasIP.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAnalyticsAlarmNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAnalyticsAlarmNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAnalyticsAlarmNode(ctx, listRequest); err == nil {
		return result.AnalyticsAlarmNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAnalyticsNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAnalyticsNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAnalyticsNode(ctx, listRequest); err == nil {
		return result.AnalyticsNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAnalyticsSNMPNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAnalyticsSNMPNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAnalyticsSNMPNode(ctx, listRequest); err == nil {
		return result.AnalyticsSNMPNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAPIAccessListPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAPIAccessListRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAPIAccessList(ctx, listRequest); err == nil {
		return result.APIAccessList.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetApplicationPolicySetPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetApplicationPolicySetRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetApplicationPolicySet(ctx, listRequest); err == nil {
		return result.ApplicationPolicySet.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetBGPAsAServicePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetBGPAsAServiceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetBGPAsAService(ctx, listRequest); err == nil {
		return result.BGPAsAService.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetBGPRouterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetBGPRouterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetBGPRouter(ctx, listRequest); err == nil {
		return result.BGPRouter.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetBGPVPNPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetBGPVPNRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetBGPVPN(ctx, listRequest); err == nil {
		return result.BGPVPN.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetBridgeDomainPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetBridgeDomainRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetBridgeDomain(ctx, listRequest); err == nil {
		return result.BridgeDomain.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCardPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCardRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCard(ctx, listRequest); err == nil {
		return result.Card.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetConfigDatabaseNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetConfigDatabaseNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetConfigDatabaseNode(ctx, listRequest); err == nil {
		return result.ConfigDatabaseNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetConfigNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetConfigNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetConfigNode(ctx, listRequest); err == nil {
		return result.ConfigNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetConfigRootPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetConfigRootRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetConfigRoot(ctx, listRequest); err == nil {
		return result.ConfigRoot.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetControlNodeZonePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetControlNodeZoneRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetControlNodeZone(ctx, listRequest); err == nil {
		return result.ControlNodeZone.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCustomerAttachmentPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCustomerAttachmentRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCustomerAttachment(ctx, listRequest); err == nil {
		return result.CustomerAttachment.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDataCenterInterconnectPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDataCenterInterconnectRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDataCenterInterconnect(ctx, listRequest); err == nil {
		return result.DataCenterInterconnect.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDatabaseNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDatabaseNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDatabaseNode(ctx, listRequest); err == nil {
		return result.DatabaseNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDeviceImagePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDeviceImageRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDeviceImage(ctx, listRequest); err == nil {
		return result.DeviceImage.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDevicemgrNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDevicemgrNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDevicemgrNode(ctx, listRequest); err == nil {
		return result.DevicemgrNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDiscoveryServiceAssignmentPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDiscoveryServiceAssignmentRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDiscoveryServiceAssignment(ctx, listRequest); err == nil {
		return result.DiscoveryServiceAssignment.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDomainPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDomainRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDomain(ctx, listRequest); err == nil {
		return result.Domain.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDsaRulePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDsaRuleRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDsaRule(ctx, listRequest); err == nil {
		return result.DsaRule.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetE2ServiceProviderPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetE2ServiceProviderRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetE2ServiceProvider(ctx, listRequest); err == nil {
		return result.E2ServiceProvider.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFabricNamespacePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFabricNamespaceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFabricNamespace(ctx, listRequest); err == nil {
		return result.FabricNamespace.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFabricPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFabricRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFabric(ctx, listRequest); err == nil {
		return result.Fabric.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFirewallPolicyPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFirewallPolicyRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFirewallPolicy(ctx, listRequest); err == nil {
		return result.FirewallPolicy.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFirewallRulePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFirewallRuleRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFirewallRule(ctx, listRequest); err == nil {
		return result.FirewallRule.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFloatingIPPoolPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFloatingIPPoolRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFloatingIPPool(ctx, listRequest); err == nil {
		return result.FloatingIPPool.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFloatingIPPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFloatingIPRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFloatingIP(ctx, listRequest); err == nil {
		return result.FloatingIP.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetForwardingClassPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetForwardingClassRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetForwardingClass(ctx, listRequest); err == nil {
		return result.ForwardingClass.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetGlobalAnalyticsConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetGlobalAnalyticsConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetGlobalAnalyticsConfig(ctx, listRequest); err == nil {
		return result.GlobalAnalyticsConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetGlobalQosConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetGlobalQosConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetGlobalQosConfig(ctx, listRequest); err == nil {
		return result.GlobalQosConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetGlobalSystemConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetGlobalSystemConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetGlobalSystemConfig(ctx, listRequest); err == nil {
		return result.GlobalSystemConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetGlobalVrouterConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetGlobalVrouterConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetGlobalVrouterConfig(ctx, listRequest); err == nil {
		return result.GlobalVrouterConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetHardwarePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetHardwareRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetHardware(ctx, listRequest); err == nil {
		return result.Hardware.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetInstanceIPPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetInstanceIPRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetInstanceIP(ctx, listRequest); err == nil {
		return result.InstanceIP.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetInterfaceRouteTablePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetInterfaceRouteTableRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetInterfaceRouteTable(ctx, listRequest); err == nil {
		return result.InterfaceRouteTable.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetJobTemplatePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetJobTemplateRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetJobTemplate(ctx, listRequest); err == nil {
		return result.JobTemplate.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLinkAggregationGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLinkAggregationGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLinkAggregationGroup(ctx, listRequest); err == nil {
		return result.LinkAggregationGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLoadbalancerHealthmonitorPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLoadbalancerHealthmonitorRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLoadbalancerHealthmonitor(ctx, listRequest); err == nil {
		return result.LoadbalancerHealthmonitor.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLoadbalancerListenerPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLoadbalancerListenerRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLoadbalancerListener(ctx, listRequest); err == nil {
		return result.LoadbalancerListener.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLoadbalancerMemberPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLoadbalancerMemberRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLoadbalancerMember(ctx, listRequest); err == nil {
		return result.LoadbalancerMember.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLoadbalancerPoolPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLoadbalancerPoolRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLoadbalancerPool(ctx, listRequest); err == nil {
		return result.LoadbalancerPool.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLoadbalancerPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLoadbalancerRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLoadbalancer(ctx, listRequest); err == nil {
		return result.Loadbalancer.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLogicalInterfacePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLogicalInterfaceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLogicalInterface(ctx, listRequest); err == nil {
		return result.LogicalInterface.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetLogicalRouterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetLogicalRouterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetLogicalRouter(ctx, listRequest); err == nil {
		return result.LogicalRouter.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMulticastGroupAddressPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMulticastGroupAddressRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMulticastGroupAddress(ctx, listRequest); err == nil {
		return result.MulticastGroupAddress.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMulticastGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMulticastGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMulticastGroup(ctx, listRequest); err == nil {
		return result.MulticastGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMulticastPolicyPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMulticastPolicyRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMulticastPolicy(ctx, listRequest); err == nil {
		return result.MulticastPolicy.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetNamespacePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetNamespaceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetNamespace(ctx, listRequest); err == nil {
		return result.Namespace.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetNetworkDeviceConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetNetworkDeviceConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetNetworkDeviceConfig(ctx, listRequest); err == nil {
		return result.NetworkDeviceConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetNetworkIpamPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetNetworkIpamRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetNetworkIpam(ctx, listRequest); err == nil {
		return result.NetworkIpam.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetNetworkPolicyPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetNetworkPolicyRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetNetworkPolicy(ctx, listRequest); err == nil {
		return result.NetworkPolicy.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetNodeProfilePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetNodeProfileRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetNodeProfile(ctx, listRequest); err == nil {
		return result.NodeProfile.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetNode(ctx, listRequest); err == nil {
		return result.Node.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPeeringPolicyPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPeeringPolicyRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPeeringPolicy(ctx, listRequest); err == nil {
		return result.PeeringPolicy.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPhysicalInterfacePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPhysicalInterfaceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPhysicalInterface(ctx, listRequest); err == nil {
		return result.PhysicalInterface.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPhysicalRouterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPhysicalRouterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPhysicalRouter(ctx, listRequest); err == nil {
		return result.PhysicalRouter.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPolicyManagementPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPolicyManagementRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPolicyManagement(ctx, listRequest); err == nil {
		return result.PolicyManagement.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPortGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPortGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPortGroup(ctx, listRequest); err == nil {
		return result.PortGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPortPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPortRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPort(ctx, listRequest); err == nil {
		return result.Port.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetPortTuplePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetPortTupleRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetPortTuple(ctx, listRequest); err == nil {
		return result.PortTuple.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetProjectPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetProjectRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetProject(ctx, listRequest); err == nil {
		return result.Project.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetProviderAttachmentPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetProviderAttachmentRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetProviderAttachment(ctx, listRequest); err == nil {
		return result.ProviderAttachment.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetQosConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetQosConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetQosConfig(ctx, listRequest); err == nil {
		return result.QosConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetQosQueuePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetQosQueueRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetQosQueue(ctx, listRequest); err == nil {
		return result.QosQueue.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRoleConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRoleConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRoleConfig(ctx, listRequest); err == nil {
		return result.RoleConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRouteAggregatePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRouteAggregateRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRouteAggregate(ctx, listRequest); err == nil {
		return result.RouteAggregate.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRouteTablePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRouteTableRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRouteTable(ctx, listRequest); err == nil {
		return result.RouteTable.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRouteTargetPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRouteTargetRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRouteTarget(ctx, listRequest); err == nil {
		return result.RouteTarget.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRoutingInstancePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRoutingInstanceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRoutingInstance(ctx, listRequest); err == nil {
		return result.RoutingInstance.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRoutingPolicyPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRoutingPolicyRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRoutingPolicy(ctx, listRequest); err == nil {
		return result.RoutingPolicy.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetSecurityGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetSecurityGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetSecurityGroup(ctx, listRequest); err == nil {
		return result.SecurityGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetSecurityLoggingObjectPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetSecurityLoggingObjectRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetSecurityLoggingObject(ctx, listRequest); err == nil {
		return result.SecurityLoggingObject.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceAppliancePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceApplianceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceAppliance(ctx, listRequest); err == nil {
		return result.ServiceAppliance.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceApplianceSetPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceApplianceSetRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceApplianceSet(ctx, listRequest); err == nil {
		return result.ServiceApplianceSet.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceConnectionModulePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceConnectionModuleRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceConnectionModule(ctx, listRequest); err == nil {
		return result.ServiceConnectionModule.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceEndpointPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceEndpointRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceEndpoint(ctx, listRequest); err == nil {
		return result.ServiceEndpoint.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceGroup(ctx, listRequest); err == nil {
		return result.ServiceGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceHealthCheckPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceHealthCheckRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceHealthCheck(ctx, listRequest); err == nil {
		return result.ServiceHealthCheck.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceInstancePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceInstanceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceInstance(ctx, listRequest); err == nil {
		return result.ServiceInstance.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceObjectPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceObjectRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceObject(ctx, listRequest); err == nil {
		return result.ServiceObject.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServiceTemplatePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServiceTemplateRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServiceTemplate(ctx, listRequest); err == nil {
		return result.ServiceTemplate.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetStructuredSyslogApplicationRecordPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetStructuredSyslogApplicationRecordRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetStructuredSyslogApplicationRecord(ctx, listRequest); err == nil {
		return result.StructuredSyslogApplicationRecord.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetStructuredSyslogConfigPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetStructuredSyslogConfigRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetStructuredSyslogConfig(ctx, listRequest); err == nil {
		return result.StructuredSyslogConfig.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetStructuredSyslogHostnameRecordPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetStructuredSyslogHostnameRecordRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetStructuredSyslogHostnameRecord(ctx, listRequest); err == nil {
		return result.StructuredSyslogHostnameRecord.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetStructuredSyslogMessagePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetStructuredSyslogMessageRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetStructuredSyslogMessage(ctx, listRequest); err == nil {
		return result.StructuredSyslogMessage.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetStructuredSyslogSLAProfilePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetStructuredSyslogSLAProfileRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetStructuredSyslogSLAProfile(ctx, listRequest); err == nil {
		return result.StructuredSyslogSLAProfile.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetSubClusterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetSubClusterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetSubCluster(ctx, listRequest); err == nil {
		return result.SubCluster.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetSubnetPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetSubnetRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetSubnet(ctx, listRequest); err == nil {
		return result.Subnet.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetTagPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetTagRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetTag(ctx, listRequest); err == nil {
		return result.Tag.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetTagTypePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetTagTypeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetTagType(ctx, listRequest); err == nil {
		return result.TagType.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualDNSRecordPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualDNSRecordRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualDNSRecord(ctx, listRequest); err == nil {
		return result.VirtualDNSRecord.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualDNSPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualDNSRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualDNS(ctx, listRequest); err == nil {
		return result.VirtualDNS.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualIPPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualIPRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualIP(ctx, listRequest); err == nil {
		return result.VirtualIP.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualMachineInterfacePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualMachineInterfaceRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualMachineInterface(ctx, listRequest); err == nil {
		return result.VirtualMachineInterface.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualMachinePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualMachineRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualMachine(ctx, listRequest); err == nil {
		return result.VirtualMachine.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualNetworkPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualNetworkRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualNetwork(ctx, listRequest); err == nil {
		return result.VirtualNetwork.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualPortGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualPortGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualPortGroup(ctx, listRequest); err == nil {
		return result.VirtualPortGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualRouterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualRouterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualRouter(ctx, listRequest); err == nil {
		return result.VirtualRouter.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetWebuiNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetWebuiNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetWebuiNode(ctx, listRequest); err == nil {
		return result.WebuiNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixBareHostNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixBareHostNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixBareHostNode(ctx, listRequest); err == nil {
		return result.AppformixBareHostNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixClusterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixClusterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixCluster(ctx, listRequest); err == nil {
		return result.AppformixCluster.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixComputeNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixComputeNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixComputeNode(ctx, listRequest); err == nil {
		return result.AppformixComputeNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixControllerNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixControllerNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixControllerNode(ctx, listRequest); err == nil {
		return result.AppformixControllerNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixNetworkAgentsNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixNetworkAgentsNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixNetworkAgentsNode(ctx, listRequest); err == nil {
		return result.AppformixNetworkAgentsNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixOpenstackNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixOpenstackNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixOpenstackNode(ctx, listRequest); err == nil {
		return result.AppformixOpenstackNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloud(ctx, listRequest); err == nil {
		return result.Cloud.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudPrivateSubnetPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudPrivateSubnetRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloudPrivateSubnet(ctx, listRequest); err == nil {
		return result.CloudPrivateSubnet.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudProviderPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudProviderRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloudProvider(ctx, listRequest); err == nil {
		return result.CloudProvider.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudRegionPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudRegionRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloudRegion(ctx, listRequest); err == nil {
		return result.CloudRegion.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudSecurityGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudSecurityGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloudSecurityGroup(ctx, listRequest); err == nil {
		return result.CloudSecurityGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudSecurityGroupRulePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudSecurityGroupRuleRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloudSecurityGroupRule(ctx, listRequest); err == nil {
		return result.CloudSecurityGroupRule.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCloudUserPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCloudUserRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCloudUser(ctx, listRequest); err == nil {
		return result.CloudUser.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVirtualCloudPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVirtualCloudRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVirtualCloud(ctx, listRequest); err == nil {
		return result.VirtualCloud.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailAnalyticsAlarmNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailAnalyticsAlarmNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailAnalyticsAlarmNode(ctx, listRequest); err == nil {
		return result.ContrailAnalyticsAlarmNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailAnalyticsDatabaseNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailAnalyticsDatabaseNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailAnalyticsDatabaseNode(ctx, listRequest); err == nil {
		return result.ContrailAnalyticsDatabaseNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailAnalyticsNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailAnalyticsNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailAnalyticsNode(ctx, listRequest); err == nil {
		return result.ContrailAnalyticsNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailAnalyticsSNMPNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailAnalyticsSNMPNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailAnalyticsSNMPNode(ctx, listRequest); err == nil {
		return result.ContrailAnalyticsSNMPNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailClusterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailClusterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailCluster(ctx, listRequest); err == nil {
		return result.ContrailCluster.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailConfigDatabaseNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailConfigDatabaseNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailConfigDatabaseNode(ctx, listRequest); err == nil {
		return result.ContrailConfigDatabaseNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailConfigNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailConfigNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailConfigNode(ctx, listRequest); err == nil {
		return result.ContrailConfigNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailControlNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailControlNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailControlNode(ctx, listRequest); err == nil {
		return result.ContrailControlNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailMulticloudGWNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailMulticloudGWNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailMulticloudGWNode(ctx, listRequest); err == nil {
		return result.ContrailMulticloudGWNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailServiceNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailServiceNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailServiceNode(ctx, listRequest); err == nil {
		return result.ContrailServiceNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailStorageNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailStorageNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailStorageNode(ctx, listRequest); err == nil {
		return result.ContrailStorageNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailVcenterFabricManagerNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailVcenterFabricManagerNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailVcenterFabricManagerNode(ctx, listRequest); err == nil {
		return result.ContrailVcenterFabricManagerNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailVrouterNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailVrouterNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailVrouterNode(ctx, listRequest); err == nil {
		return result.ContrailVrouterNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailWebuiNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailWebuiNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailWebuiNode(ctx, listRequest); err == nil {
		return result.ContrailWebuiNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailWinCNMPluginNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailWinCNMPluginNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailWinCNMPluginNode(ctx, listRequest); err == nil {
		return result.ContrailWinCNMPluginNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailZTPDHCPNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailZTPDHCPNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailZTPDHCPNode(ctx, listRequest); err == nil {
		return result.ContrailZTPDHCPNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetContrailZTPTFTPNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetContrailZTPTFTPNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetContrailZTPTFTPNode(ctx, listRequest); err == nil {
		return result.ContrailZTPTFTPNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetCredentialPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetCredentialRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetCredential(ctx, listRequest); err == nil {
		return result.Credential.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetDashboardPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetDashboardRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetDashboard(ctx, listRequest); err == nil {
		return result.Dashboard.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetEndpointPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetEndpointRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetEndpoint(ctx, listRequest); err == nil {
		return result.Endpoint.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetFlavorPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetFlavorRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetFlavor(ctx, listRequest); err == nil {
		return result.Flavor.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOsImagePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOsImageRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOsImage(ctx, listRequest); err == nil {
		return result.OsImage.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetKeypairPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetKeypairRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetKeypair(ctx, listRequest); err == nil {
		return result.Keypair.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetKubernetesClusterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetKubernetesClusterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetKubernetesCluster(ctx, listRequest); err == nil {
		return result.KubernetesCluster.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetKubernetesKubemanagerNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetKubernetesKubemanagerNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetKubernetesKubemanagerNode(ctx, listRequest); err == nil {
		return result.KubernetesKubemanagerNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetKubernetesMasterNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetKubernetesMasterNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetKubernetesMasterNode(ctx, listRequest); err == nil {
		return result.KubernetesMasterNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetKubernetesNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetKubernetesNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetKubernetesNode(ctx, listRequest); err == nil {
		return result.KubernetesNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMesosAgentPrivateNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMesosAgentPrivateNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMesosAgentPrivateNode(ctx, listRequest); err == nil {
		return result.MesosAgentPrivateNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMesosAgentPublicNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMesosAgentPublicNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMesosAgentPublicNode(ctx, listRequest); err == nil {
		return result.MesosAgentPublicNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMesosClusterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMesosClusterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMesosCluster(ctx, listRequest); err == nil {
		return result.MesosCluster.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMesosMasterNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMesosMasterNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMesosMasterNode(ctx, listRequest); err == nil {
		return result.MesosMasterNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetMesosMesosmanagerNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetMesosMesosmanagerNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetMesosMesosmanagerNode(ctx, listRequest); err == nil {
		return result.MesosMesosmanagerNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackBaremetalNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackBaremetalNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackBaremetalNode(ctx, listRequest); err == nil {
		return result.OpenstackBaremetalNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackClusterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackClusterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackCluster(ctx, listRequest); err == nil {
		return result.OpenstackCluster.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackComputeNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackComputeNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackComputeNode(ctx, listRequest); err == nil {
		return result.OpenstackComputeNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackControlNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackControlNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackControlNode(ctx, listRequest); err == nil {
		return result.OpenstackControlNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackMonitoringNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackMonitoringNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackMonitoringNode(ctx, listRequest); err == nil {
		return result.OpenstackMonitoringNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackNetworkNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackNetworkNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackNetworkNode(ctx, listRequest); err == nil {
		return result.OpenstackNetworkNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetOpenstackStorageNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetOpenstackStorageNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetOpenstackStorageNode(ctx, listRequest); err == nil {
		return result.OpenstackStorageNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRhospdCloudManagerPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRhospdCloudManagerRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRhospdCloudManager(ctx, listRequest); err == nil {
		return result.RhospdCloudManager.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRhospdFlavorPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRhospdFlavorRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRhospdFlavor(ctx, listRequest); err == nil {
		return result.RhospdFlavor.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRhospdJumphostNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRhospdJumphostNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRhospdJumphostNode(ctx, listRequest); err == nil {
		return result.RhospdJumphostNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRhospdOvercloudNetworkPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRhospdOvercloudNetworkRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRhospdOvercloudNetwork(ctx, listRequest); err == nil {
		return result.RhospdOvercloudNetwork.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRhospdOvercloudNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRhospdOvercloudNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRhospdOvercloudNode(ctx, listRequest); err == nil {
		return result.RhospdOvercloudNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetRhospdUndercloudNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetRhospdUndercloudNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetRhospdUndercloudNode(ctx, listRequest); err == nil {
		return result.RhospdUndercloudNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetServerPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetServerRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetServer(ctx, listRequest); err == nil {
		return result.Server.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVCenterPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVCenterRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVCenter(ctx, listRequest); err == nil {
		return result.VCenter.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVCenterComputePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVCenterComputeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVCenterCompute(ctx, listRequest); err == nil {
		return result.VCenterCompute.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVCenterManagerNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVCenterManagerNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVCenterManagerNode(ctx, listRequest); err == nil {
		return result.VCenterManagerNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVCenterPluginNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVCenterPluginNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVCenterPluginNode(ctx, listRequest); err == nil {
		return result.VCenterPluginNode.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetVPNGroupPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetVPNGroupRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetVPNGroup(ctx, listRequest); err == nil {
		return result.VPNGroup.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetWidgetPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetWidgetRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetWidget(ctx, listRequest); err == nil {
		return result.Widget.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixFlowsPermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixFlowsRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixFlows(ctx, listRequest); err == nil {
		return result.AppformixFlows.GetPerms2().ToRBAC()
	}
	return nil
}

func (ag contrailAccessGetter) GetAppformixFlowsNodePermType2(ctx context.Context, uuid string) *rbac.PermType2 {
	listRequest := &services.GetAppformixFlowsNodeRequest{
		ID: uuid,
	}
	if result, err := ag.service.GetAppformixFlowsNode(ctx, listRequest); err == nil {
		return result.AppformixFlowsNode.GetPerms2().ToRBAC()
	}
	return nil
}
