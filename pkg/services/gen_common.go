package services

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// nolint
type ContrailService struct {
	BaseService
}

// nolint
type RESTResource struct {
	Kind string      `json:"kind"`
	Data interface{} `json:"data"`
}

// nolint
type RESTSyncRequest struct {
	Resources []*RESTResource `json:"resources"`
}

//RESTSync handle a bluk Create REST service.
func (service *ContrailService) RESTSync(c echo.Context) error {
	requestData := &RESTSyncRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	responses := []interface{}{}
	for _, resource := range requestData.Resources {
		switch resource.Kind {

		case "access_control_list":
			request := &models.CreateAccessControlListRequest{
				AccessControlList: models.InterfaceToAccessControlList(resource.Data),
			}
			response, err := service.CreateAccessControlList(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.AccessControlList,
			})

		case "address_group":
			request := &models.CreateAddressGroupRequest{
				AddressGroup: models.InterfaceToAddressGroup(resource.Data),
			}
			response, err := service.CreateAddressGroup(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.AddressGroup,
			})

		case "alarm":
			request := &models.CreateAlarmRequest{
				Alarm: models.InterfaceToAlarm(resource.Data),
			}
			response, err := service.CreateAlarm(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Alarm,
			})

		case "alias_ip_pool":
			request := &models.CreateAliasIPPoolRequest{
				AliasIPPool: models.InterfaceToAliasIPPool(resource.Data),
			}
			response, err := service.CreateAliasIPPool(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.AliasIPPool,
			})

		case "alias_ip":
			request := &models.CreateAliasIPRequest{
				AliasIP: models.InterfaceToAliasIP(resource.Data),
			}
			response, err := service.CreateAliasIP(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.AliasIP,
			})

		case "analytics_node":
			request := &models.CreateAnalyticsNodeRequest{
				AnalyticsNode: models.InterfaceToAnalyticsNode(resource.Data),
			}
			response, err := service.CreateAnalyticsNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.AnalyticsNode,
			})

		case "api_access_list":
			request := &models.CreateAPIAccessListRequest{
				APIAccessList: models.InterfaceToAPIAccessList(resource.Data),
			}
			response, err := service.CreateAPIAccessList(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.APIAccessList,
			})

		case "application_policy_set":
			request := &models.CreateApplicationPolicySetRequest{
				ApplicationPolicySet: models.InterfaceToApplicationPolicySet(resource.Data),
			}
			response, err := service.CreateApplicationPolicySet(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ApplicationPolicySet,
			})

		case "bgp_as_a_service":
			request := &models.CreateBGPAsAServiceRequest{
				BGPAsAService: models.InterfaceToBGPAsAService(resource.Data),
			}
			response, err := service.CreateBGPAsAService(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.BGPAsAService,
			})

		case "bgp_router":
			request := &models.CreateBGPRouterRequest{
				BGPRouter: models.InterfaceToBGPRouter(resource.Data),
			}
			response, err := service.CreateBGPRouter(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.BGPRouter,
			})

		case "bgpvpn":
			request := &models.CreateBGPVPNRequest{
				BGPVPN: models.InterfaceToBGPVPN(resource.Data),
			}
			response, err := service.CreateBGPVPN(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.BGPVPN,
			})

		case "bridge_domain":
			request := &models.CreateBridgeDomainRequest{
				BridgeDomain: models.InterfaceToBridgeDomain(resource.Data),
			}
			response, err := service.CreateBridgeDomain(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.BridgeDomain,
			})

		case "config_node":
			request := &models.CreateConfigNodeRequest{
				ConfigNode: models.InterfaceToConfigNode(resource.Data),
			}
			response, err := service.CreateConfigNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ConfigNode,
			})

		case "config_root":
			request := &models.CreateConfigRootRequest{
				ConfigRoot: models.InterfaceToConfigRoot(resource.Data),
			}
			response, err := service.CreateConfigRoot(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ConfigRoot,
			})

		case "customer_attachment":
			request := &models.CreateCustomerAttachmentRequest{
				CustomerAttachment: models.InterfaceToCustomerAttachment(resource.Data),
			}
			response, err := service.CreateCustomerAttachment(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.CustomerAttachment,
			})

		case "database_node":
			request := &models.CreateDatabaseNodeRequest{
				DatabaseNode: models.InterfaceToDatabaseNode(resource.Data),
			}
			response, err := service.CreateDatabaseNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.DatabaseNode,
			})

		case "discovery_service_assignment":
			request := &models.CreateDiscoveryServiceAssignmentRequest{
				DiscoveryServiceAssignment: models.InterfaceToDiscoveryServiceAssignment(resource.Data),
			}
			response, err := service.CreateDiscoveryServiceAssignment(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.DiscoveryServiceAssignment,
			})

		case "domain":
			request := &models.CreateDomainRequest{
				Domain: models.InterfaceToDomain(resource.Data),
			}
			response, err := service.CreateDomain(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Domain,
			})

		case "dsa_rule":
			request := &models.CreateDsaRuleRequest{
				DsaRule: models.InterfaceToDsaRule(resource.Data),
			}
			response, err := service.CreateDsaRule(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.DsaRule,
			})

		case "e2_service_provider":
			request := &models.CreateE2ServiceProviderRequest{
				E2ServiceProvider: models.InterfaceToE2ServiceProvider(resource.Data),
			}
			response, err := service.CreateE2ServiceProvider(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.E2ServiceProvider,
			})

		case "firewall_policy":
			request := &models.CreateFirewallPolicyRequest{
				FirewallPolicy: models.InterfaceToFirewallPolicy(resource.Data),
			}
			response, err := service.CreateFirewallPolicy(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.FirewallPolicy,
			})

		case "firewall_rule":
			request := &models.CreateFirewallRuleRequest{
				FirewallRule: models.InterfaceToFirewallRule(resource.Data),
			}
			response, err := service.CreateFirewallRule(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.FirewallRule,
			})

		case "floating_ip_pool":
			request := &models.CreateFloatingIPPoolRequest{
				FloatingIPPool: models.InterfaceToFloatingIPPool(resource.Data),
			}
			response, err := service.CreateFloatingIPPool(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.FloatingIPPool,
			})

		case "floating_ip":
			request := &models.CreateFloatingIPRequest{
				FloatingIP: models.InterfaceToFloatingIP(resource.Data),
			}
			response, err := service.CreateFloatingIP(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.FloatingIP,
			})

		case "forwarding_class":
			request := &models.CreateForwardingClassRequest{
				ForwardingClass: models.InterfaceToForwardingClass(resource.Data),
			}
			response, err := service.CreateForwardingClass(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ForwardingClass,
			})

		case "global_qos_config":
			request := &models.CreateGlobalQosConfigRequest{
				GlobalQosConfig: models.InterfaceToGlobalQosConfig(resource.Data),
			}
			response, err := service.CreateGlobalQosConfig(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.GlobalQosConfig,
			})

		case "global_system_config":
			request := &models.CreateGlobalSystemConfigRequest{
				GlobalSystemConfig: models.InterfaceToGlobalSystemConfig(resource.Data),
			}
			response, err := service.CreateGlobalSystemConfig(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.GlobalSystemConfig,
			})

		case "global_vrouter_config":
			request := &models.CreateGlobalVrouterConfigRequest{
				GlobalVrouterConfig: models.InterfaceToGlobalVrouterConfig(resource.Data),
			}
			response, err := service.CreateGlobalVrouterConfig(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.GlobalVrouterConfig,
			})

		case "instance_ip":
			request := &models.CreateInstanceIPRequest{
				InstanceIP: models.InterfaceToInstanceIP(resource.Data),
			}
			response, err := service.CreateInstanceIP(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.InstanceIP,
			})

		case "interface_route_table":
			request := &models.CreateInterfaceRouteTableRequest{
				InterfaceRouteTable: models.InterfaceToInterfaceRouteTable(resource.Data),
			}
			response, err := service.CreateInterfaceRouteTable(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.InterfaceRouteTable,
			})

		case "loadbalancer_healthmonitor":
			request := &models.CreateLoadbalancerHealthmonitorRequest{
				LoadbalancerHealthmonitor: models.InterfaceToLoadbalancerHealthmonitor(resource.Data),
			}
			response, err := service.CreateLoadbalancerHealthmonitor(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.LoadbalancerHealthmonitor,
			})

		case "loadbalancer_listener":
			request := &models.CreateLoadbalancerListenerRequest{
				LoadbalancerListener: models.InterfaceToLoadbalancerListener(resource.Data),
			}
			response, err := service.CreateLoadbalancerListener(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.LoadbalancerListener,
			})

		case "loadbalancer_member":
			request := &models.CreateLoadbalancerMemberRequest{
				LoadbalancerMember: models.InterfaceToLoadbalancerMember(resource.Data),
			}
			response, err := service.CreateLoadbalancerMember(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.LoadbalancerMember,
			})

		case "loadbalancer_pool":
			request := &models.CreateLoadbalancerPoolRequest{
				LoadbalancerPool: models.InterfaceToLoadbalancerPool(resource.Data),
			}
			response, err := service.CreateLoadbalancerPool(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.LoadbalancerPool,
			})

		case "loadbalancer":
			request := &models.CreateLoadbalancerRequest{
				Loadbalancer: models.InterfaceToLoadbalancer(resource.Data),
			}
			response, err := service.CreateLoadbalancer(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Loadbalancer,
			})

		case "logical_interface":
			request := &models.CreateLogicalInterfaceRequest{
				LogicalInterface: models.InterfaceToLogicalInterface(resource.Data),
			}
			response, err := service.CreateLogicalInterface(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.LogicalInterface,
			})

		case "logical_router":
			request := &models.CreateLogicalRouterRequest{
				LogicalRouter: models.InterfaceToLogicalRouter(resource.Data),
			}
			response, err := service.CreateLogicalRouter(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.LogicalRouter,
			})

		case "namespace":
			request := &models.CreateNamespaceRequest{
				Namespace: models.InterfaceToNamespace(resource.Data),
			}
			response, err := service.CreateNamespace(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Namespace,
			})

		case "network_device_config":
			request := &models.CreateNetworkDeviceConfigRequest{
				NetworkDeviceConfig: models.InterfaceToNetworkDeviceConfig(resource.Data),
			}
			response, err := service.CreateNetworkDeviceConfig(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.NetworkDeviceConfig,
			})

		case "network_ipam":
			request := &models.CreateNetworkIpamRequest{
				NetworkIpam: models.InterfaceToNetworkIpam(resource.Data),
			}
			response, err := service.CreateNetworkIpam(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.NetworkIpam,
			})

		case "network_policy":
			request := &models.CreateNetworkPolicyRequest{
				NetworkPolicy: models.InterfaceToNetworkPolicy(resource.Data),
			}
			response, err := service.CreateNetworkPolicy(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.NetworkPolicy,
			})

		case "peering_policy":
			request := &models.CreatePeeringPolicyRequest{
				PeeringPolicy: models.InterfaceToPeeringPolicy(resource.Data),
			}
			response, err := service.CreatePeeringPolicy(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.PeeringPolicy,
			})

		case "physical_interface":
			request := &models.CreatePhysicalInterfaceRequest{
				PhysicalInterface: models.InterfaceToPhysicalInterface(resource.Data),
			}
			response, err := service.CreatePhysicalInterface(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.PhysicalInterface,
			})

		case "physical_router":
			request := &models.CreatePhysicalRouterRequest{
				PhysicalRouter: models.InterfaceToPhysicalRouter(resource.Data),
			}
			response, err := service.CreatePhysicalRouter(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.PhysicalRouter,
			})

		case "policy_management":
			request := &models.CreatePolicyManagementRequest{
				PolicyManagement: models.InterfaceToPolicyManagement(resource.Data),
			}
			response, err := service.CreatePolicyManagement(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.PolicyManagement,
			})

		case "port_tuple":
			request := &models.CreatePortTupleRequest{
				PortTuple: models.InterfaceToPortTuple(resource.Data),
			}
			response, err := service.CreatePortTuple(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.PortTuple,
			})

		case "project":
			request := &models.CreateProjectRequest{
				Project: models.InterfaceToProject(resource.Data),
			}
			response, err := service.CreateProject(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Project,
			})

		case "provider_attachment":
			request := &models.CreateProviderAttachmentRequest{
				ProviderAttachment: models.InterfaceToProviderAttachment(resource.Data),
			}
			response, err := service.CreateProviderAttachment(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ProviderAttachment,
			})

		case "qos_config":
			request := &models.CreateQosConfigRequest{
				QosConfig: models.InterfaceToQosConfig(resource.Data),
			}
			response, err := service.CreateQosConfig(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.QosConfig,
			})

		case "qos_queue":
			request := &models.CreateQosQueueRequest{
				QosQueue: models.InterfaceToQosQueue(resource.Data),
			}
			response, err := service.CreateQosQueue(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.QosQueue,
			})

		case "route_aggregate":
			request := &models.CreateRouteAggregateRequest{
				RouteAggregate: models.InterfaceToRouteAggregate(resource.Data),
			}
			response, err := service.CreateRouteAggregate(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.RouteAggregate,
			})

		case "route_table":
			request := &models.CreateRouteTableRequest{
				RouteTable: models.InterfaceToRouteTable(resource.Data),
			}
			response, err := service.CreateRouteTable(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.RouteTable,
			})

		case "route_target":
			request := &models.CreateRouteTargetRequest{
				RouteTarget: models.InterfaceToRouteTarget(resource.Data),
			}
			response, err := service.CreateRouteTarget(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.RouteTarget,
			})

		case "routing_instance":
			request := &models.CreateRoutingInstanceRequest{
				RoutingInstance: models.InterfaceToRoutingInstance(resource.Data),
			}
			response, err := service.CreateRoutingInstance(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.RoutingInstance,
			})

		case "routing_policy":
			request := &models.CreateRoutingPolicyRequest{
				RoutingPolicy: models.InterfaceToRoutingPolicy(resource.Data),
			}
			response, err := service.CreateRoutingPolicy(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.RoutingPolicy,
			})

		case "security_group":
			request := &models.CreateSecurityGroupRequest{
				SecurityGroup: models.InterfaceToSecurityGroup(resource.Data),
			}
			response, err := service.CreateSecurityGroup(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.SecurityGroup,
			})

		case "security_logging_object":
			request := &models.CreateSecurityLoggingObjectRequest{
				SecurityLoggingObject: models.InterfaceToSecurityLoggingObject(resource.Data),
			}
			response, err := service.CreateSecurityLoggingObject(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.SecurityLoggingObject,
			})

		case "service_appliance":
			request := &models.CreateServiceApplianceRequest{
				ServiceAppliance: models.InterfaceToServiceAppliance(resource.Data),
			}
			response, err := service.CreateServiceAppliance(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceAppliance,
			})

		case "service_appliance_set":
			request := &models.CreateServiceApplianceSetRequest{
				ServiceApplianceSet: models.InterfaceToServiceApplianceSet(resource.Data),
			}
			response, err := service.CreateServiceApplianceSet(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceApplianceSet,
			})

		case "service_connection_module":
			request := &models.CreateServiceConnectionModuleRequest{
				ServiceConnectionModule: models.InterfaceToServiceConnectionModule(resource.Data),
			}
			response, err := service.CreateServiceConnectionModule(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceConnectionModule,
			})

		case "service_endpoint":
			request := &models.CreateServiceEndpointRequest{
				ServiceEndpoint: models.InterfaceToServiceEndpoint(resource.Data),
			}
			response, err := service.CreateServiceEndpoint(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceEndpoint,
			})

		case "service_group":
			request := &models.CreateServiceGroupRequest{
				ServiceGroup: models.InterfaceToServiceGroup(resource.Data),
			}
			response, err := service.CreateServiceGroup(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceGroup,
			})

		case "service_health_check":
			request := &models.CreateServiceHealthCheckRequest{
				ServiceHealthCheck: models.InterfaceToServiceHealthCheck(resource.Data),
			}
			response, err := service.CreateServiceHealthCheck(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceHealthCheck,
			})

		case "service_instance":
			request := &models.CreateServiceInstanceRequest{
				ServiceInstance: models.InterfaceToServiceInstance(resource.Data),
			}
			response, err := service.CreateServiceInstance(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceInstance,
			})

		case "service_object":
			request := &models.CreateServiceObjectRequest{
				ServiceObject: models.InterfaceToServiceObject(resource.Data),
			}
			response, err := service.CreateServiceObject(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceObject,
			})

		case "service_template":
			request := &models.CreateServiceTemplateRequest{
				ServiceTemplate: models.InterfaceToServiceTemplate(resource.Data),
			}
			response, err := service.CreateServiceTemplate(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ServiceTemplate,
			})

		case "subnet":
			request := &models.CreateSubnetRequest{
				Subnet: models.InterfaceToSubnet(resource.Data),
			}
			response, err := service.CreateSubnet(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Subnet,
			})

		case "tag":
			request := &models.CreateTagRequest{
				Tag: models.InterfaceToTag(resource.Data),
			}
			response, err := service.CreateTag(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Tag,
			})

		case "tag_type":
			request := &models.CreateTagTypeRequest{
				TagType: models.InterfaceToTagType(resource.Data),
			}
			response, err := service.CreateTagType(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.TagType,
			})

		case "user":
			request := &models.CreateUserRequest{
				User: models.InterfaceToUser(resource.Data),
			}
			response, err := service.CreateUser(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.User,
			})

		case "virtual_DNS_record":
			request := &models.CreateVirtualDNSRecordRequest{
				VirtualDNSRecord: models.InterfaceToVirtualDNSRecord(resource.Data),
			}
			response, err := service.CreateVirtualDNSRecord(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualDNSRecord,
			})

		case "virtual_DNS":
			request := &models.CreateVirtualDNSRequest{
				VirtualDNS: models.InterfaceToVirtualDNS(resource.Data),
			}
			response, err := service.CreateVirtualDNS(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualDNS,
			})

		case "virtual_ip":
			request := &models.CreateVirtualIPRequest{
				VirtualIP: models.InterfaceToVirtualIP(resource.Data),
			}
			response, err := service.CreateVirtualIP(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualIP,
			})

		case "virtual_machine_interface":
			request := &models.CreateVirtualMachineInterfaceRequest{
				VirtualMachineInterface: models.InterfaceToVirtualMachineInterface(resource.Data),
			}
			response, err := service.CreateVirtualMachineInterface(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualMachineInterface,
			})

		case "virtual_machine":
			request := &models.CreateVirtualMachineRequest{
				VirtualMachine: models.InterfaceToVirtualMachine(resource.Data),
			}
			response, err := service.CreateVirtualMachine(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualMachine,
			})

		case "virtual_network":
			request := &models.CreateVirtualNetworkRequest{
				VirtualNetwork: models.InterfaceToVirtualNetwork(resource.Data),
			}
			response, err := service.CreateVirtualNetwork(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualNetwork,
			})

		case "virtual_router":
			request := &models.CreateVirtualRouterRequest{
				VirtualRouter: models.InterfaceToVirtualRouter(resource.Data),
			}
			response, err := service.CreateVirtualRouter(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VirtualRouter,
			})

		case "appformix_node":
			request := &models.CreateAppformixNodeRequest{
				AppformixNode: models.InterfaceToAppformixNode(resource.Data),
			}
			response, err := service.CreateAppformixNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.AppformixNode,
			})

		case "baremetal_node":
			request := &models.CreateBaremetalNodeRequest{
				BaremetalNode: models.InterfaceToBaremetalNode(resource.Data),
			}
			response, err := service.CreateBaremetalNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.BaremetalNode,
			})

		case "baremetal_port":
			request := &models.CreateBaremetalPortRequest{
				BaremetalPort: models.InterfaceToBaremetalPort(resource.Data),
			}
			response, err := service.CreateBaremetalPort(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.BaremetalPort,
			})

		case "contrail_analytics_database_node":
			request := &models.CreateContrailAnalyticsDatabaseNodeRequest{
				ContrailAnalyticsDatabaseNode: models.InterfaceToContrailAnalyticsDatabaseNode(resource.Data),
			}
			response, err := service.CreateContrailAnalyticsDatabaseNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailAnalyticsDatabaseNode,
			})

		case "contrail_analytics_node":
			request := &models.CreateContrailAnalyticsNodeRequest{
				ContrailAnalyticsNode: models.InterfaceToContrailAnalyticsNode(resource.Data),
			}
			response, err := service.CreateContrailAnalyticsNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailAnalyticsNode,
			})

		case "contrail_cluster":
			request := &models.CreateContrailClusterRequest{
				ContrailCluster: models.InterfaceToContrailCluster(resource.Data),
			}
			response, err := service.CreateContrailCluster(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailCluster,
			})

		case "contrail_config_database_node":
			request := &models.CreateContrailConfigDatabaseNodeRequest{
				ContrailConfigDatabaseNode: models.InterfaceToContrailConfigDatabaseNode(resource.Data),
			}
			response, err := service.CreateContrailConfigDatabaseNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailConfigDatabaseNode,
			})

		case "contrail_config_node":
			request := &models.CreateContrailConfigNodeRequest{
				ContrailConfigNode: models.InterfaceToContrailConfigNode(resource.Data),
			}
			response, err := service.CreateContrailConfigNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailConfigNode,
			})

		case "contrail_control_node":
			request := &models.CreateContrailControlNodeRequest{
				ContrailControlNode: models.InterfaceToContrailControlNode(resource.Data),
			}
			response, err := service.CreateContrailControlNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailControlNode,
			})

		case "contrail_storage_node":
			request := &models.CreateContrailStorageNodeRequest{
				ContrailStorageNode: models.InterfaceToContrailStorageNode(resource.Data),
			}
			response, err := service.CreateContrailStorageNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailStorageNode,
			})

		case "contrail_vrouter_node":
			request := &models.CreateContrailVrouterNodeRequest{
				ContrailVrouterNode: models.InterfaceToContrailVrouterNode(resource.Data),
			}
			response, err := service.CreateContrailVrouterNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailVrouterNode,
			})

		case "contrail_controller_node":
			request := &models.CreateContrailControllerNodeRequest{
				ContrailControllerNode: models.InterfaceToContrailControllerNode(resource.Data),
			}
			response, err := service.CreateContrailControllerNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.ContrailControllerNode,
			})

		case "dashboard":
			request := &models.CreateDashboardRequest{
				Dashboard: models.InterfaceToDashboard(resource.Data),
			}
			response, err := service.CreateDashboard(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Dashboard,
			})

		case "flavor":
			request := &models.CreateFlavorRequest{
				Flavor: models.InterfaceToFlavor(resource.Data),
			}
			response, err := service.CreateFlavor(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Flavor,
			})

		case "os_image":
			request := &models.CreateOsImageRequest{
				OsImage: models.InterfaceToOsImage(resource.Data),
			}
			response, err := service.CreateOsImage(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.OsImage,
			})

		case "keypair":
			request := &models.CreateKeypairRequest{
				Keypair: models.InterfaceToKeypair(resource.Data),
			}
			response, err := service.CreateKeypair(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Keypair,
			})

		case "kubernetes_master_node":
			request := &models.CreateKubernetesMasterNodeRequest{
				KubernetesMasterNode: models.InterfaceToKubernetesMasterNode(resource.Data),
			}
			response, err := service.CreateKubernetesMasterNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.KubernetesMasterNode,
			})

		case "kubernetes_node":
			request := &models.CreateKubernetesNodeRequest{
				KubernetesNode: models.InterfaceToKubernetesNode(resource.Data),
			}
			response, err := service.CreateKubernetesNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.KubernetesNode,
			})

		case "location":
			request := &models.CreateLocationRequest{
				Location: models.InterfaceToLocation(resource.Data),
			}
			response, err := service.CreateLocation(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Location,
			})

		case "node":
			request := &models.CreateNodeRequest{
				Node: models.InterfaceToNode(resource.Data),
			}
			response, err := service.CreateNode(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Node,
			})

		case "server":
			request := &models.CreateServerRequest{
				Server: models.InterfaceToServer(resource.Data),
			}
			response, err := service.CreateServer(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Server,
			})

		case "vpn_group":
			request := &models.CreateVPNGroupRequest{
				VPNGroup: models.InterfaceToVPNGroup(resource.Data),
			}
			response, err := service.CreateVPNGroup(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.VPNGroup,
			})

		case "widget":
			request := &models.CreateWidgetRequest{
				Widget: models.InterfaceToWidget(resource.Data),
			}
			response, err := service.CreateWidget(ctx, request)
			if err != nil {
				return common.ToHTTPError(err)
			}
			responses = append(responses, &RESTResource{
				Kind: resource.Kind,
				Data: response.Widget,
			})

		}
	}
	return c.JSON(http.StatusCreated, responses)
}

//RegisterRESTAPI register REST API service for path.
// nolint
func (service *ContrailService) RegisterRESTAPI(e *echo.Echo) {

	e.POST("/access-control-lists", service.RESTCreateAccessControlList)
	e.GET("/access-control-lists", service.RESTListAccessControlList)
	e.PUT("/access-control-list/:id", service.RESTUpdateAccessControlList)
	e.GET("/access-control-list/:id", service.RESTGetAccessControlList)
	e.DELETE("/access-control-list/:id", service.RESTDeleteAccessControlList)

	e.POST("/address-groups", service.RESTCreateAddressGroup)
	e.GET("/address-groups", service.RESTListAddressGroup)
	e.PUT("/address-group/:id", service.RESTUpdateAddressGroup)
	e.GET("/address-group/:id", service.RESTGetAddressGroup)
	e.DELETE("/address-group/:id", service.RESTDeleteAddressGroup)

	e.POST("/alarms", service.RESTCreateAlarm)
	e.GET("/alarms", service.RESTListAlarm)
	e.PUT("/alarm/:id", service.RESTUpdateAlarm)
	e.GET("/alarm/:id", service.RESTGetAlarm)
	e.DELETE("/alarm/:id", service.RESTDeleteAlarm)

	e.POST("/alias-ip-pools", service.RESTCreateAliasIPPool)
	e.GET("/alias-ip-pools", service.RESTListAliasIPPool)
	e.PUT("/alias-ip-pool/:id", service.RESTUpdateAliasIPPool)
	e.GET("/alias-ip-pool/:id", service.RESTGetAliasIPPool)
	e.DELETE("/alias-ip-pool/:id", service.RESTDeleteAliasIPPool)

	e.POST("/alias-ips", service.RESTCreateAliasIP)
	e.GET("/alias-ips", service.RESTListAliasIP)
	e.PUT("/alias-ip/:id", service.RESTUpdateAliasIP)
	e.GET("/alias-ip/:id", service.RESTGetAliasIP)
	e.DELETE("/alias-ip/:id", service.RESTDeleteAliasIP)

	e.POST("/analytics-nodes", service.RESTCreateAnalyticsNode)
	e.GET("/analytics-nodes", service.RESTListAnalyticsNode)
	e.PUT("/analytics-node/:id", service.RESTUpdateAnalyticsNode)
	e.GET("/analytics-node/:id", service.RESTGetAnalyticsNode)
	e.DELETE("/analytics-node/:id", service.RESTDeleteAnalyticsNode)

	e.POST("/api-access-lists", service.RESTCreateAPIAccessList)
	e.GET("/api-access-lists", service.RESTListAPIAccessList)
	e.PUT("/api-access-list/:id", service.RESTUpdateAPIAccessList)
	e.GET("/api-access-list/:id", service.RESTGetAPIAccessList)
	e.DELETE("/api-access-list/:id", service.RESTDeleteAPIAccessList)

	e.POST("/application-policy-sets", service.RESTCreateApplicationPolicySet)
	e.GET("/application-policy-sets", service.RESTListApplicationPolicySet)
	e.PUT("/application-policy-set/:id", service.RESTUpdateApplicationPolicySet)
	e.GET("/application-policy-set/:id", service.RESTGetApplicationPolicySet)
	e.DELETE("/application-policy-set/:id", service.RESTDeleteApplicationPolicySet)

	e.POST("/bgp-as-a-services", service.RESTCreateBGPAsAService)
	e.GET("/bgp-as-a-services", service.RESTListBGPAsAService)
	e.PUT("/bgp-as-a-service/:id", service.RESTUpdateBGPAsAService)
	e.GET("/bgp-as-a-service/:id", service.RESTGetBGPAsAService)
	e.DELETE("/bgp-as-a-service/:id", service.RESTDeleteBGPAsAService)

	e.POST("/bgp-routers", service.RESTCreateBGPRouter)
	e.GET("/bgp-routers", service.RESTListBGPRouter)
	e.PUT("/bgp-router/:id", service.RESTUpdateBGPRouter)
	e.GET("/bgp-router/:id", service.RESTGetBGPRouter)
	e.DELETE("/bgp-router/:id", service.RESTDeleteBGPRouter)

	e.POST("/bgpvpns", service.RESTCreateBGPVPN)
	e.GET("/bgpvpns", service.RESTListBGPVPN)
	e.PUT("/bgpvpn/:id", service.RESTUpdateBGPVPN)
	e.GET("/bgpvpn/:id", service.RESTGetBGPVPN)
	e.DELETE("/bgpvpn/:id", service.RESTDeleteBGPVPN)

	e.POST("/bridge-domains", service.RESTCreateBridgeDomain)
	e.GET("/bridge-domains", service.RESTListBridgeDomain)
	e.PUT("/bridge-domain/:id", service.RESTUpdateBridgeDomain)
	e.GET("/bridge-domain/:id", service.RESTGetBridgeDomain)
	e.DELETE("/bridge-domain/:id", service.RESTDeleteBridgeDomain)

	e.POST("/config-nodes", service.RESTCreateConfigNode)
	e.GET("/config-nodes", service.RESTListConfigNode)
	e.PUT("/config-node/:id", service.RESTUpdateConfigNode)
	e.GET("/config-node/:id", service.RESTGetConfigNode)
	e.DELETE("/config-node/:id", service.RESTDeleteConfigNode)

	e.POST("/config-roots", service.RESTCreateConfigRoot)
	e.GET("/config-roots", service.RESTListConfigRoot)
	e.PUT("/config-root/:id", service.RESTUpdateConfigRoot)
	e.GET("/config-root/:id", service.RESTGetConfigRoot)
	e.DELETE("/config-root/:id", service.RESTDeleteConfigRoot)

	e.POST("/customer-attachments", service.RESTCreateCustomerAttachment)
	e.GET("/customer-attachments", service.RESTListCustomerAttachment)
	e.PUT("/customer-attachment/:id", service.RESTUpdateCustomerAttachment)
	e.GET("/customer-attachment/:id", service.RESTGetCustomerAttachment)
	e.DELETE("/customer-attachment/:id", service.RESTDeleteCustomerAttachment)

	e.POST("/database-nodes", service.RESTCreateDatabaseNode)
	e.GET("/database-nodes", service.RESTListDatabaseNode)
	e.PUT("/database-node/:id", service.RESTUpdateDatabaseNode)
	e.GET("/database-node/:id", service.RESTGetDatabaseNode)
	e.DELETE("/database-node/:id", service.RESTDeleteDatabaseNode)

	e.POST("/discovery-service-assignments", service.RESTCreateDiscoveryServiceAssignment)
	e.GET("/discovery-service-assignments", service.RESTListDiscoveryServiceAssignment)
	e.PUT("/discovery-service-assignment/:id", service.RESTUpdateDiscoveryServiceAssignment)
	e.GET("/discovery-service-assignment/:id", service.RESTGetDiscoveryServiceAssignment)
	e.DELETE("/discovery-service-assignment/:id", service.RESTDeleteDiscoveryServiceAssignment)

	e.POST("/domains", service.RESTCreateDomain)
	e.GET("/domains", service.RESTListDomain)
	e.PUT("/domain/:id", service.RESTUpdateDomain)
	e.GET("/domain/:id", service.RESTGetDomain)
	e.DELETE("/domain/:id", service.RESTDeleteDomain)

	e.POST("/dsa-rules", service.RESTCreateDsaRule)
	e.GET("/dsa-rules", service.RESTListDsaRule)
	e.PUT("/dsa-rule/:id", service.RESTUpdateDsaRule)
	e.GET("/dsa-rule/:id", service.RESTGetDsaRule)
	e.DELETE("/dsa-rule/:id", service.RESTDeleteDsaRule)

	e.POST("/e2-service-providers", service.RESTCreateE2ServiceProvider)
	e.GET("/e2-service-providers", service.RESTListE2ServiceProvider)
	e.PUT("/e2-service-provider/:id", service.RESTUpdateE2ServiceProvider)
	e.GET("/e2-service-provider/:id", service.RESTGetE2ServiceProvider)
	e.DELETE("/e2-service-provider/:id", service.RESTDeleteE2ServiceProvider)

	e.POST("/firewall-policys", service.RESTCreateFirewallPolicy)
	e.GET("/firewall-policys", service.RESTListFirewallPolicy)
	e.PUT("/firewall-policy/:id", service.RESTUpdateFirewallPolicy)
	e.GET("/firewall-policy/:id", service.RESTGetFirewallPolicy)
	e.DELETE("/firewall-policy/:id", service.RESTDeleteFirewallPolicy)

	e.POST("/firewall-rules", service.RESTCreateFirewallRule)
	e.GET("/firewall-rules", service.RESTListFirewallRule)
	e.PUT("/firewall-rule/:id", service.RESTUpdateFirewallRule)
	e.GET("/firewall-rule/:id", service.RESTGetFirewallRule)
	e.DELETE("/firewall-rule/:id", service.RESTDeleteFirewallRule)

	e.POST("/floating-ip-pools", service.RESTCreateFloatingIPPool)
	e.GET("/floating-ip-pools", service.RESTListFloatingIPPool)
	e.PUT("/floating-ip-pool/:id", service.RESTUpdateFloatingIPPool)
	e.GET("/floating-ip-pool/:id", service.RESTGetFloatingIPPool)
	e.DELETE("/floating-ip-pool/:id", service.RESTDeleteFloatingIPPool)

	e.POST("/floating-ips", service.RESTCreateFloatingIP)
	e.GET("/floating-ips", service.RESTListFloatingIP)
	e.PUT("/floating-ip/:id", service.RESTUpdateFloatingIP)
	e.GET("/floating-ip/:id", service.RESTGetFloatingIP)
	e.DELETE("/floating-ip/:id", service.RESTDeleteFloatingIP)

	e.POST("/forwarding-classs", service.RESTCreateForwardingClass)
	e.GET("/forwarding-classs", service.RESTListForwardingClass)
	e.PUT("/forwarding-class/:id", service.RESTUpdateForwardingClass)
	e.GET("/forwarding-class/:id", service.RESTGetForwardingClass)
	e.DELETE("/forwarding-class/:id", service.RESTDeleteForwardingClass)

	e.POST("/global-qos-configs", service.RESTCreateGlobalQosConfig)
	e.GET("/global-qos-configs", service.RESTListGlobalQosConfig)
	e.PUT("/global-qos-config/:id", service.RESTUpdateGlobalQosConfig)
	e.GET("/global-qos-config/:id", service.RESTGetGlobalQosConfig)
	e.DELETE("/global-qos-config/:id", service.RESTDeleteGlobalQosConfig)

	e.POST("/global-system-configs", service.RESTCreateGlobalSystemConfig)
	e.GET("/global-system-configs", service.RESTListGlobalSystemConfig)
	e.PUT("/global-system-config/:id", service.RESTUpdateGlobalSystemConfig)
	e.GET("/global-system-config/:id", service.RESTGetGlobalSystemConfig)
	e.DELETE("/global-system-config/:id", service.RESTDeleteGlobalSystemConfig)

	e.POST("/global-vrouter-configs", service.RESTCreateGlobalVrouterConfig)
	e.GET("/global-vrouter-configs", service.RESTListGlobalVrouterConfig)
	e.PUT("/global-vrouter-config/:id", service.RESTUpdateGlobalVrouterConfig)
	e.GET("/global-vrouter-config/:id", service.RESTGetGlobalVrouterConfig)
	e.DELETE("/global-vrouter-config/:id", service.RESTDeleteGlobalVrouterConfig)

	e.POST("/instance-ips", service.RESTCreateInstanceIP)
	e.GET("/instance-ips", service.RESTListInstanceIP)
	e.PUT("/instance-ip/:id", service.RESTUpdateInstanceIP)
	e.GET("/instance-ip/:id", service.RESTGetInstanceIP)
	e.DELETE("/instance-ip/:id", service.RESTDeleteInstanceIP)

	e.POST("/interface-route-tables", service.RESTCreateInterfaceRouteTable)
	e.GET("/interface-route-tables", service.RESTListInterfaceRouteTable)
	e.PUT("/interface-route-table/:id", service.RESTUpdateInterfaceRouteTable)
	e.GET("/interface-route-table/:id", service.RESTGetInterfaceRouteTable)
	e.DELETE("/interface-route-table/:id", service.RESTDeleteInterfaceRouteTable)

	e.POST("/loadbalancer-healthmonitors", service.RESTCreateLoadbalancerHealthmonitor)
	e.GET("/loadbalancer-healthmonitors", service.RESTListLoadbalancerHealthmonitor)
	e.PUT("/loadbalancer-healthmonitor/:id", service.RESTUpdateLoadbalancerHealthmonitor)
	e.GET("/loadbalancer-healthmonitor/:id", service.RESTGetLoadbalancerHealthmonitor)
	e.DELETE("/loadbalancer-healthmonitor/:id", service.RESTDeleteLoadbalancerHealthmonitor)

	e.POST("/loadbalancer-listeners", service.RESTCreateLoadbalancerListener)
	e.GET("/loadbalancer-listeners", service.RESTListLoadbalancerListener)
	e.PUT("/loadbalancer-listener/:id", service.RESTUpdateLoadbalancerListener)
	e.GET("/loadbalancer-listener/:id", service.RESTGetLoadbalancerListener)
	e.DELETE("/loadbalancer-listener/:id", service.RESTDeleteLoadbalancerListener)

	e.POST("/loadbalancer-members", service.RESTCreateLoadbalancerMember)
	e.GET("/loadbalancer-members", service.RESTListLoadbalancerMember)
	e.PUT("/loadbalancer-member/:id", service.RESTUpdateLoadbalancerMember)
	e.GET("/loadbalancer-member/:id", service.RESTGetLoadbalancerMember)
	e.DELETE("/loadbalancer-member/:id", service.RESTDeleteLoadbalancerMember)

	e.POST("/loadbalancer-pools", service.RESTCreateLoadbalancerPool)
	e.GET("/loadbalancer-pools", service.RESTListLoadbalancerPool)
	e.PUT("/loadbalancer-pool/:id", service.RESTUpdateLoadbalancerPool)
	e.GET("/loadbalancer-pool/:id", service.RESTGetLoadbalancerPool)
	e.DELETE("/loadbalancer-pool/:id", service.RESTDeleteLoadbalancerPool)

	e.POST("/loadbalancers", service.RESTCreateLoadbalancer)
	e.GET("/loadbalancers", service.RESTListLoadbalancer)
	e.PUT("/loadbalancer/:id", service.RESTUpdateLoadbalancer)
	e.GET("/loadbalancer/:id", service.RESTGetLoadbalancer)
	e.DELETE("/loadbalancer/:id", service.RESTDeleteLoadbalancer)

	e.POST("/logical-interfaces", service.RESTCreateLogicalInterface)
	e.GET("/logical-interfaces", service.RESTListLogicalInterface)
	e.PUT("/logical-interface/:id", service.RESTUpdateLogicalInterface)
	e.GET("/logical-interface/:id", service.RESTGetLogicalInterface)
	e.DELETE("/logical-interface/:id", service.RESTDeleteLogicalInterface)

	e.POST("/logical-routers", service.RESTCreateLogicalRouter)
	e.GET("/logical-routers", service.RESTListLogicalRouter)
	e.PUT("/logical-router/:id", service.RESTUpdateLogicalRouter)
	e.GET("/logical-router/:id", service.RESTGetLogicalRouter)
	e.DELETE("/logical-router/:id", service.RESTDeleteLogicalRouter)

	e.POST("/namespaces", service.RESTCreateNamespace)
	e.GET("/namespaces", service.RESTListNamespace)
	e.PUT("/namespace/:id", service.RESTUpdateNamespace)
	e.GET("/namespace/:id", service.RESTGetNamespace)
	e.DELETE("/namespace/:id", service.RESTDeleteNamespace)

	e.POST("/network-device-configs", service.RESTCreateNetworkDeviceConfig)
	e.GET("/network-device-configs", service.RESTListNetworkDeviceConfig)
	e.PUT("/network-device-config/:id", service.RESTUpdateNetworkDeviceConfig)
	e.GET("/network-device-config/:id", service.RESTGetNetworkDeviceConfig)
	e.DELETE("/network-device-config/:id", service.RESTDeleteNetworkDeviceConfig)

	e.POST("/network-ipams", service.RESTCreateNetworkIpam)
	e.GET("/network-ipams", service.RESTListNetworkIpam)
	e.PUT("/network-ipam/:id", service.RESTUpdateNetworkIpam)
	e.GET("/network-ipam/:id", service.RESTGetNetworkIpam)
	e.DELETE("/network-ipam/:id", service.RESTDeleteNetworkIpam)

	e.POST("/network-policys", service.RESTCreateNetworkPolicy)
	e.GET("/network-policys", service.RESTListNetworkPolicy)
	e.PUT("/network-policy/:id", service.RESTUpdateNetworkPolicy)
	e.GET("/network-policy/:id", service.RESTGetNetworkPolicy)
	e.DELETE("/network-policy/:id", service.RESTDeleteNetworkPolicy)

	e.POST("/peering-policys", service.RESTCreatePeeringPolicy)
	e.GET("/peering-policys", service.RESTListPeeringPolicy)
	e.PUT("/peering-policy/:id", service.RESTUpdatePeeringPolicy)
	e.GET("/peering-policy/:id", service.RESTGetPeeringPolicy)
	e.DELETE("/peering-policy/:id", service.RESTDeletePeeringPolicy)

	e.POST("/physical-interfaces", service.RESTCreatePhysicalInterface)
	e.GET("/physical-interfaces", service.RESTListPhysicalInterface)
	e.PUT("/physical-interface/:id", service.RESTUpdatePhysicalInterface)
	e.GET("/physical-interface/:id", service.RESTGetPhysicalInterface)
	e.DELETE("/physical-interface/:id", service.RESTDeletePhysicalInterface)

	e.POST("/physical-routers", service.RESTCreatePhysicalRouter)
	e.GET("/physical-routers", service.RESTListPhysicalRouter)
	e.PUT("/physical-router/:id", service.RESTUpdatePhysicalRouter)
	e.GET("/physical-router/:id", service.RESTGetPhysicalRouter)
	e.DELETE("/physical-router/:id", service.RESTDeletePhysicalRouter)

	e.POST("/policy-managements", service.RESTCreatePolicyManagement)
	e.GET("/policy-managements", service.RESTListPolicyManagement)
	e.PUT("/policy-management/:id", service.RESTUpdatePolicyManagement)
	e.GET("/policy-management/:id", service.RESTGetPolicyManagement)
	e.DELETE("/policy-management/:id", service.RESTDeletePolicyManagement)

	e.POST("/port-tuples", service.RESTCreatePortTuple)
	e.GET("/port-tuples", service.RESTListPortTuple)
	e.PUT("/port-tuple/:id", service.RESTUpdatePortTuple)
	e.GET("/port-tuple/:id", service.RESTGetPortTuple)
	e.DELETE("/port-tuple/:id", service.RESTDeletePortTuple)

	e.POST("/projects", service.RESTCreateProject)
	e.GET("/projects", service.RESTListProject)
	e.PUT("/project/:id", service.RESTUpdateProject)
	e.GET("/project/:id", service.RESTGetProject)
	e.DELETE("/project/:id", service.RESTDeleteProject)

	e.POST("/provider-attachments", service.RESTCreateProviderAttachment)
	e.GET("/provider-attachments", service.RESTListProviderAttachment)
	e.PUT("/provider-attachment/:id", service.RESTUpdateProviderAttachment)
	e.GET("/provider-attachment/:id", service.RESTGetProviderAttachment)
	e.DELETE("/provider-attachment/:id", service.RESTDeleteProviderAttachment)

	e.POST("/qos-configs", service.RESTCreateQosConfig)
	e.GET("/qos-configs", service.RESTListQosConfig)
	e.PUT("/qos-config/:id", service.RESTUpdateQosConfig)
	e.GET("/qos-config/:id", service.RESTGetQosConfig)
	e.DELETE("/qos-config/:id", service.RESTDeleteQosConfig)

	e.POST("/qos-queues", service.RESTCreateQosQueue)
	e.GET("/qos-queues", service.RESTListQosQueue)
	e.PUT("/qos-queue/:id", service.RESTUpdateQosQueue)
	e.GET("/qos-queue/:id", service.RESTGetQosQueue)
	e.DELETE("/qos-queue/:id", service.RESTDeleteQosQueue)

	e.POST("/route-aggregates", service.RESTCreateRouteAggregate)
	e.GET("/route-aggregates", service.RESTListRouteAggregate)
	e.PUT("/route-aggregate/:id", service.RESTUpdateRouteAggregate)
	e.GET("/route-aggregate/:id", service.RESTGetRouteAggregate)
	e.DELETE("/route-aggregate/:id", service.RESTDeleteRouteAggregate)

	e.POST("/route-tables", service.RESTCreateRouteTable)
	e.GET("/route-tables", service.RESTListRouteTable)
	e.PUT("/route-table/:id", service.RESTUpdateRouteTable)
	e.GET("/route-table/:id", service.RESTGetRouteTable)
	e.DELETE("/route-table/:id", service.RESTDeleteRouteTable)

	e.POST("/route-targets", service.RESTCreateRouteTarget)
	e.GET("/route-targets", service.RESTListRouteTarget)
	e.PUT("/route-target/:id", service.RESTUpdateRouteTarget)
	e.GET("/route-target/:id", service.RESTGetRouteTarget)
	e.DELETE("/route-target/:id", service.RESTDeleteRouteTarget)

	e.POST("/routing-instances", service.RESTCreateRoutingInstance)
	e.GET("/routing-instances", service.RESTListRoutingInstance)
	e.PUT("/routing-instance/:id", service.RESTUpdateRoutingInstance)
	e.GET("/routing-instance/:id", service.RESTGetRoutingInstance)
	e.DELETE("/routing-instance/:id", service.RESTDeleteRoutingInstance)

	e.POST("/routing-policys", service.RESTCreateRoutingPolicy)
	e.GET("/routing-policys", service.RESTListRoutingPolicy)
	e.PUT("/routing-policy/:id", service.RESTUpdateRoutingPolicy)
	e.GET("/routing-policy/:id", service.RESTGetRoutingPolicy)
	e.DELETE("/routing-policy/:id", service.RESTDeleteRoutingPolicy)

	e.POST("/security-groups", service.RESTCreateSecurityGroup)
	e.GET("/security-groups", service.RESTListSecurityGroup)
	e.PUT("/security-group/:id", service.RESTUpdateSecurityGroup)
	e.GET("/security-group/:id", service.RESTGetSecurityGroup)
	e.DELETE("/security-group/:id", service.RESTDeleteSecurityGroup)

	e.POST("/security-logging-objects", service.RESTCreateSecurityLoggingObject)
	e.GET("/security-logging-objects", service.RESTListSecurityLoggingObject)
	e.PUT("/security-logging-object/:id", service.RESTUpdateSecurityLoggingObject)
	e.GET("/security-logging-object/:id", service.RESTGetSecurityLoggingObject)
	e.DELETE("/security-logging-object/:id", service.RESTDeleteSecurityLoggingObject)

	e.POST("/service-appliances", service.RESTCreateServiceAppliance)
	e.GET("/service-appliances", service.RESTListServiceAppliance)
	e.PUT("/service-appliance/:id", service.RESTUpdateServiceAppliance)
	e.GET("/service-appliance/:id", service.RESTGetServiceAppliance)
	e.DELETE("/service-appliance/:id", service.RESTDeleteServiceAppliance)

	e.POST("/service-appliance-sets", service.RESTCreateServiceApplianceSet)
	e.GET("/service-appliance-sets", service.RESTListServiceApplianceSet)
	e.PUT("/service-appliance-set/:id", service.RESTUpdateServiceApplianceSet)
	e.GET("/service-appliance-set/:id", service.RESTGetServiceApplianceSet)
	e.DELETE("/service-appliance-set/:id", service.RESTDeleteServiceApplianceSet)

	e.POST("/service-connection-modules", service.RESTCreateServiceConnectionModule)
	e.GET("/service-connection-modules", service.RESTListServiceConnectionModule)
	e.PUT("/service-connection-module/:id", service.RESTUpdateServiceConnectionModule)
	e.GET("/service-connection-module/:id", service.RESTGetServiceConnectionModule)
	e.DELETE("/service-connection-module/:id", service.RESTDeleteServiceConnectionModule)

	e.POST("/service-endpoints", service.RESTCreateServiceEndpoint)
	e.GET("/service-endpoints", service.RESTListServiceEndpoint)
	e.PUT("/service-endpoint/:id", service.RESTUpdateServiceEndpoint)
	e.GET("/service-endpoint/:id", service.RESTGetServiceEndpoint)
	e.DELETE("/service-endpoint/:id", service.RESTDeleteServiceEndpoint)

	e.POST("/service-groups", service.RESTCreateServiceGroup)
	e.GET("/service-groups", service.RESTListServiceGroup)
	e.PUT("/service-group/:id", service.RESTUpdateServiceGroup)
	e.GET("/service-group/:id", service.RESTGetServiceGroup)
	e.DELETE("/service-group/:id", service.RESTDeleteServiceGroup)

	e.POST("/service-health-checks", service.RESTCreateServiceHealthCheck)
	e.GET("/service-health-checks", service.RESTListServiceHealthCheck)
	e.PUT("/service-health-check/:id", service.RESTUpdateServiceHealthCheck)
	e.GET("/service-health-check/:id", service.RESTGetServiceHealthCheck)
	e.DELETE("/service-health-check/:id", service.RESTDeleteServiceHealthCheck)

	e.POST("/service-instances", service.RESTCreateServiceInstance)
	e.GET("/service-instances", service.RESTListServiceInstance)
	e.PUT("/service-instance/:id", service.RESTUpdateServiceInstance)
	e.GET("/service-instance/:id", service.RESTGetServiceInstance)
	e.DELETE("/service-instance/:id", service.RESTDeleteServiceInstance)

	e.POST("/service-objects", service.RESTCreateServiceObject)
	e.GET("/service-objects", service.RESTListServiceObject)
	e.PUT("/service-object/:id", service.RESTUpdateServiceObject)
	e.GET("/service-object/:id", service.RESTGetServiceObject)
	e.DELETE("/service-object/:id", service.RESTDeleteServiceObject)

	e.POST("/service-templates", service.RESTCreateServiceTemplate)
	e.GET("/service-templates", service.RESTListServiceTemplate)
	e.PUT("/service-template/:id", service.RESTUpdateServiceTemplate)
	e.GET("/service-template/:id", service.RESTGetServiceTemplate)
	e.DELETE("/service-template/:id", service.RESTDeleteServiceTemplate)

	e.POST("/subnets", service.RESTCreateSubnet)
	e.GET("/subnets", service.RESTListSubnet)
	e.PUT("/subnet/:id", service.RESTUpdateSubnet)
	e.GET("/subnet/:id", service.RESTGetSubnet)
	e.DELETE("/subnet/:id", service.RESTDeleteSubnet)

	e.POST("/tags", service.RESTCreateTag)
	e.GET("/tags", service.RESTListTag)
	e.PUT("/tag/:id", service.RESTUpdateTag)
	e.GET("/tag/:id", service.RESTGetTag)
	e.DELETE("/tag/:id", service.RESTDeleteTag)

	e.POST("/tag-types", service.RESTCreateTagType)
	e.GET("/tag-types", service.RESTListTagType)
	e.PUT("/tag-type/:id", service.RESTUpdateTagType)
	e.GET("/tag-type/:id", service.RESTGetTagType)
	e.DELETE("/tag-type/:id", service.RESTDeleteTagType)

	e.POST("/users", service.RESTCreateUser)
	e.GET("/users", service.RESTListUser)
	e.PUT("/user/:id", service.RESTUpdateUser)
	e.GET("/user/:id", service.RESTGetUser)
	e.DELETE("/user/:id", service.RESTDeleteUser)

	e.POST("/virtual-DNS-records", service.RESTCreateVirtualDNSRecord)
	e.GET("/virtual-DNS-records", service.RESTListVirtualDNSRecord)
	e.PUT("/virtual-DNS-record/:id", service.RESTUpdateVirtualDNSRecord)
	e.GET("/virtual-DNS-record/:id", service.RESTGetVirtualDNSRecord)
	e.DELETE("/virtual-DNS-record/:id", service.RESTDeleteVirtualDNSRecord)

	e.POST("/virtual-DNSs", service.RESTCreateVirtualDNS)
	e.GET("/virtual-DNSs", service.RESTListVirtualDNS)
	e.PUT("/virtual-DNS/:id", service.RESTUpdateVirtualDNS)
	e.GET("/virtual-DNS/:id", service.RESTGetVirtualDNS)
	e.DELETE("/virtual-DNS/:id", service.RESTDeleteVirtualDNS)

	e.POST("/virtual-ips", service.RESTCreateVirtualIP)
	e.GET("/virtual-ips", service.RESTListVirtualIP)
	e.PUT("/virtual-ip/:id", service.RESTUpdateVirtualIP)
	e.GET("/virtual-ip/:id", service.RESTGetVirtualIP)
	e.DELETE("/virtual-ip/:id", service.RESTDeleteVirtualIP)

	e.POST("/virtual-machine-interfaces", service.RESTCreateVirtualMachineInterface)
	e.GET("/virtual-machine-interfaces", service.RESTListVirtualMachineInterface)
	e.PUT("/virtual-machine-interface/:id", service.RESTUpdateVirtualMachineInterface)
	e.GET("/virtual-machine-interface/:id", service.RESTGetVirtualMachineInterface)
	e.DELETE("/virtual-machine-interface/:id", service.RESTDeleteVirtualMachineInterface)

	e.POST("/virtual-machines", service.RESTCreateVirtualMachine)
	e.GET("/virtual-machines", service.RESTListVirtualMachine)
	e.PUT("/virtual-machine/:id", service.RESTUpdateVirtualMachine)
	e.GET("/virtual-machine/:id", service.RESTGetVirtualMachine)
	e.DELETE("/virtual-machine/:id", service.RESTDeleteVirtualMachine)

	e.POST("/virtual-networks", service.RESTCreateVirtualNetwork)
	e.GET("/virtual-networks", service.RESTListVirtualNetwork)
	e.PUT("/virtual-network/:id", service.RESTUpdateVirtualNetwork)
	e.GET("/virtual-network/:id", service.RESTGetVirtualNetwork)
	e.DELETE("/virtual-network/:id", service.RESTDeleteVirtualNetwork)

	e.POST("/virtual-routers", service.RESTCreateVirtualRouter)
	e.GET("/virtual-routers", service.RESTListVirtualRouter)
	e.PUT("/virtual-router/:id", service.RESTUpdateVirtualRouter)
	e.GET("/virtual-router/:id", service.RESTGetVirtualRouter)
	e.DELETE("/virtual-router/:id", service.RESTDeleteVirtualRouter)

	e.POST("/appformix-nodes", service.RESTCreateAppformixNode)
	e.GET("/appformix-nodes", service.RESTListAppformixNode)
	e.PUT("/appformix-node/:id", service.RESTUpdateAppformixNode)
	e.GET("/appformix-node/:id", service.RESTGetAppformixNode)
	e.DELETE("/appformix-node/:id", service.RESTDeleteAppformixNode)

	e.POST("/openstackbaremetal-nodes", service.RESTCreateBaremetalNode)
	e.GET("/openstackbaremetal-nodes", service.RESTListBaremetalNode)
	e.PUT("/openstackbaremetal-node/:id", service.RESTUpdateBaremetalNode)
	e.GET("/openstackbaremetal-node/:id", service.RESTGetBaremetalNode)
	e.DELETE("/openstackbaremetal-node/:id", service.RESTDeleteBaremetalNode)

	e.POST("/openstackbaremetal-ports", service.RESTCreateBaremetalPort)
	e.GET("/openstackbaremetal-ports", service.RESTListBaremetalPort)
	e.PUT("/openstackbaremetal-port/:id", service.RESTUpdateBaremetalPort)
	e.GET("/openstackbaremetal-port/:id", service.RESTGetBaremetalPort)
	e.DELETE("/openstackbaremetal-port/:id", service.RESTDeleteBaremetalPort)

	e.POST("/contrail-analytics-database-nodes", service.RESTCreateContrailAnalyticsDatabaseNode)
	e.GET("/contrail-analytics-database-nodes", service.RESTListContrailAnalyticsDatabaseNode)
	e.PUT("/contrail-analytics-database-node/:id", service.RESTUpdateContrailAnalyticsDatabaseNode)
	e.GET("/contrail-analytics-database-node/:id", service.RESTGetContrailAnalyticsDatabaseNode)
	e.DELETE("/contrail-analytics-database-node/:id", service.RESTDeleteContrailAnalyticsDatabaseNode)

	e.POST("/contrail-analytics-nodes", service.RESTCreateContrailAnalyticsNode)
	e.GET("/contrail-analytics-nodes", service.RESTListContrailAnalyticsNode)
	e.PUT("/contrail-analytics-node/:id", service.RESTUpdateContrailAnalyticsNode)
	e.GET("/contrail-analytics-node/:id", service.RESTGetContrailAnalyticsNode)
	e.DELETE("/contrail-analytics-node/:id", service.RESTDeleteContrailAnalyticsNode)

	e.POST("/contrail-clusters", service.RESTCreateContrailCluster)
	e.GET("/contrail-clusters", service.RESTListContrailCluster)
	e.PUT("/contrail-cluster/:id", service.RESTUpdateContrailCluster)
	e.GET("/contrail-cluster/:id", service.RESTGetContrailCluster)
	e.DELETE("/contrail-cluster/:id", service.RESTDeleteContrailCluster)

	e.POST("/contrail-config-database-nodes", service.RESTCreateContrailConfigDatabaseNode)
	e.GET("/contrail-config-database-nodes", service.RESTListContrailConfigDatabaseNode)
	e.PUT("/contrail-config-database-node/:id", service.RESTUpdateContrailConfigDatabaseNode)
	e.GET("/contrail-config-database-node/:id", service.RESTGetContrailConfigDatabaseNode)
	e.DELETE("/contrail-config-database-node/:id", service.RESTDeleteContrailConfigDatabaseNode)

	e.POST("/contrail-config-nodes", service.RESTCreateContrailConfigNode)
	e.GET("/contrail-config-nodes", service.RESTListContrailConfigNode)
	e.PUT("/contrail-config-node/:id", service.RESTUpdateContrailConfigNode)
	e.GET("/contrail-config-node/:id", service.RESTGetContrailConfigNode)
	e.DELETE("/contrail-config-node/:id", service.RESTDeleteContrailConfigNode)

	e.POST("/contrail-control-nodes", service.RESTCreateContrailControlNode)
	e.GET("/contrail-control-nodes", service.RESTListContrailControlNode)
	e.PUT("/contrail-control-node/:id", service.RESTUpdateContrailControlNode)
	e.GET("/contrail-control-node/:id", service.RESTGetContrailControlNode)
	e.DELETE("/contrail-control-node/:id", service.RESTDeleteContrailControlNode)

	e.POST("/contrail-storage-nodes", service.RESTCreateContrailStorageNode)
	e.GET("/contrail-storage-nodes", service.RESTListContrailStorageNode)
	e.PUT("/contrail-storage-node/:id", service.RESTUpdateContrailStorageNode)
	e.GET("/contrail-storage-node/:id", service.RESTGetContrailStorageNode)
	e.DELETE("/contrail-storage-node/:id", service.RESTDeleteContrailStorageNode)

	e.POST("/contrailcontrail-vrouter-nodes", service.RESTCreateContrailVrouterNode)
	e.GET("/contrailcontrail-vrouter-nodes", service.RESTListContrailVrouterNode)
	e.PUT("/contrailcontrail-vrouter-node/:id", service.RESTUpdateContrailVrouterNode)
	e.GET("/contrailcontrail-vrouter-node/:id", service.RESTGetContrailVrouterNode)
	e.DELETE("/contrailcontrail-vrouter-node/:id", service.RESTDeleteContrailVrouterNode)

	e.POST("/contrail-controller-nodes", service.RESTCreateContrailControllerNode)
	e.GET("/contrail-controller-nodes", service.RESTListContrailControllerNode)
	e.PUT("/contrail-controller-node/:id", service.RESTUpdateContrailControllerNode)
	e.GET("/contrail-controller-node/:id", service.RESTGetContrailControllerNode)
	e.DELETE("/contrail-controller-node/:id", service.RESTDeleteContrailControllerNode)

	e.POST("/dashboards", service.RESTCreateDashboard)
	e.GET("/dashboards", service.RESTListDashboard)
	e.PUT("/dashboard/:id", service.RESTUpdateDashboard)
	e.GET("/dashboard/:id", service.RESTGetDashboard)
	e.DELETE("/dashboard/:id", service.RESTDeleteDashboard)

	e.POST("/openstackflavors", service.RESTCreateFlavor)
	e.GET("/openstackflavors", service.RESTListFlavor)
	e.PUT("/openstackflavor/:id", service.RESTUpdateFlavor)
	e.GET("/openstackflavor/:id", service.RESTGetFlavor)
	e.DELETE("/openstackflavor/:id", service.RESTDeleteFlavor)

	e.POST("/openstackos-images", service.RESTCreateOsImage)
	e.GET("/openstackos-images", service.RESTListOsImage)
	e.PUT("/openstackos-image/:id", service.RESTUpdateOsImage)
	e.GET("/openstackos-image/:id", service.RESTGetOsImage)
	e.DELETE("/openstackos-image/:id", service.RESTDeleteOsImage)

	e.POST("/openstackkeypairs", service.RESTCreateKeypair)
	e.GET("/openstackkeypairs", service.RESTListKeypair)
	e.PUT("/openstackkeypair/:id", service.RESTUpdateKeypair)
	e.GET("/openstackkeypair/:id", service.RESTGetKeypair)
	e.DELETE("/openstackkeypair/:id", service.RESTDeleteKeypair)

	e.POST("/kubernetes-master-nodes", service.RESTCreateKubernetesMasterNode)
	e.GET("/kubernetes-master-nodes", service.RESTListKubernetesMasterNode)
	e.PUT("/kubernetes-master-node/:id", service.RESTUpdateKubernetesMasterNode)
	e.GET("/kubernetes-master-node/:id", service.RESTGetKubernetesMasterNode)
	e.DELETE("/kubernetes-master-node/:id", service.RESTDeleteKubernetesMasterNode)

	e.POST("/kubernetes-nodes", service.RESTCreateKubernetesNode)
	e.GET("/kubernetes-nodes", service.RESTListKubernetesNode)
	e.PUT("/kubernetes-node/:id", service.RESTUpdateKubernetesNode)
	e.GET("/kubernetes-node/:id", service.RESTGetKubernetesNode)
	e.DELETE("/kubernetes-node/:id", service.RESTDeleteKubernetesNode)

	e.POST("/locations", service.RESTCreateLocation)
	e.GET("/locations", service.RESTListLocation)
	e.PUT("/location/:id", service.RESTUpdateLocation)
	e.GET("/location/:id", service.RESTGetLocation)
	e.DELETE("/location/:id", service.RESTDeleteLocation)

	e.POST("/nodes", service.RESTCreateNode)
	e.GET("/nodes", service.RESTListNode)
	e.PUT("/node/:id", service.RESTUpdateNode)
	e.GET("/node/:id", service.RESTGetNode)
	e.DELETE("/node/:id", service.RESTDeleteNode)

	e.POST("/openserverservers", service.RESTCreateServer)
	e.GET("/openserverservers", service.RESTListServer)
	e.PUT("/openserverserver/:id", service.RESTUpdateServer)
	e.GET("/openserverserver/:id", service.RESTGetServer)
	e.DELETE("/openserverserver/:id", service.RESTDeleteServer)

	e.POST("/vpn-groups", service.RESTCreateVPNGroup)
	e.GET("/vpn-groups", service.RESTListVPNGroup)
	e.PUT("/vpn-group/:id", service.RESTUpdateVPNGroup)
	e.GET("/vpn-group/:id", service.RESTGetVPNGroup)
	e.DELETE("/vpn-group/:id", service.RESTDeleteVPNGroup)

	e.POST("/widgets", service.RESTCreateWidget)
	e.GET("/widgets", service.RESTListWidget)
	e.PUT("/widget/:id", service.RESTUpdateWidget)
	e.GET("/widget/:id", service.RESTGetWidget)
	e.DELETE("/widget/:id", service.RESTDeleteWidget)

	e.POST("sync", service.RESTSync)
}
