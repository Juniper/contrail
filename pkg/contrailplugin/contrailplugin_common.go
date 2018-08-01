// nolint
package contrailplugin

import (
	"sync"

	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/services"

	log "github.com/sirupsen/logrus"
)

// NewPluginService instantiates a plugin service.
func NewPluginService(name string) *PluginService {
	service := &PluginService{
		BaseService: services.BaseService{},
		Name:        name,
	}
	service.Init()
	return service
}

// PluginService
type PluginService struct {
	services.BaseService
	Name string
}

// Init initializes the plugin service
func (service *PluginService) Init() {
	// Write Plugin initialization code here
	log.Printf("PluginService(%s): initialized\n", service.Name)
	return
}

func (service *PluginService) Debug(objStr string, objValue *sync.Map) {
	oMap := make(map[interface{}]interface{})
	objValue.Range(func(k, v interface{}) bool {
		oMap[k] = v
		return true
	})
	log.Printf("%s: %v", objStr, oMap)
}

// ObjCallback to Handle Objects CRUD
type ObjCallback func(obj interface{})

var resource_functions = map[string]ObjCallback{

	"AccessControlList": EvaluateAccessControlList,

	"AddressGroup": EvaluateAddressGroup,

	"Alarm": EvaluateAlarm,

	"AliasIPPool": EvaluateAliasIPPool,

	"AliasIP": EvaluateAliasIP,

	"AnalyticsNode": EvaluateAnalyticsNode,

	"APIAccessList": EvaluateAPIAccessList,

	"ApplicationPolicySet": EvaluateApplicationPolicySet,

	"BGPAsAService": EvaluateBGPAsAService,

	"BGPRouter": EvaluateBGPRouter,

	"BGPVPN": EvaluateBGPVPN,

	"BridgeDomain": EvaluateBridgeDomain,

	"ConfigNode": EvaluateConfigNode,

	"ConfigRoot": EvaluateConfigRoot,

	"CustomerAttachment": EvaluateCustomerAttachment,

	"DatabaseNode": EvaluateDatabaseNode,

	"DeviceImage": EvaluateDeviceImage,

	"DiscoveryServiceAssignment": EvaluateDiscoveryServiceAssignment,

	"Domain": EvaluateDomain,

	"DsaRule": EvaluateDsaRule,

	"E2ServiceProvider": EvaluateE2ServiceProvider,

	"FabricNamespace": EvaluateFabricNamespace,

	"Fabric": EvaluateFabric,

	"FirewallPolicy": EvaluateFirewallPolicy,

	"FirewallRule": EvaluateFirewallRule,

	"FloatingIPPool": EvaluateFloatingIPPool,

	"FloatingIP": EvaluateFloatingIP,

	"ForwardingClass": EvaluateForwardingClass,

	"GlobalAnalyticsConfig": EvaluateGlobalAnalyticsConfig,

	"GlobalQosConfig": EvaluateGlobalQosConfig,

	"GlobalSystemConfig": EvaluateGlobalSystemConfig,

	"GlobalVrouterConfig": EvaluateGlobalVrouterConfig,

	"InstanceIP": EvaluateInstanceIP,

	"InterfaceRouteTable": EvaluateInterfaceRouteTable,

	"JobTemplate": EvaluateJobTemplate,

	"LinkAggregationGroup": EvaluateLinkAggregationGroup,

	"LoadbalancerHealthmonitor": EvaluateLoadbalancerHealthmonitor,

	"LoadbalancerListener": EvaluateLoadbalancerListener,

	"LoadbalancerMember": EvaluateLoadbalancerMember,

	"LoadbalancerPool": EvaluateLoadbalancerPool,

	"Loadbalancer": EvaluateLoadbalancer,

	"LogicalInterface": EvaluateLogicalInterface,

	"LogicalRouter": EvaluateLogicalRouter,

	"MulticastGroupAddress": EvaluateMulticastGroupAddress,

	"MulticastGroup": EvaluateMulticastGroup,

	"Namespace": EvaluateNamespace,

	"NetworkDeviceConfig": EvaluateNetworkDeviceConfig,

	"NetworkIpam": EvaluateNetworkIpam,

	"NetworkPolicy": EvaluateNetworkPolicy,

	"PeeringPolicy": EvaluatePeeringPolicy,

	"PhysicalInterface": EvaluatePhysicalInterface,

	"PhysicalRouter": EvaluatePhysicalRouter,

	"PolicyManagement": EvaluatePolicyManagement,

	"PortTuple": EvaluatePortTuple,

	"Project": EvaluateProject,

	"ProviderAttachment": EvaluateProviderAttachment,

	"QosConfig": EvaluateQosConfig,

	"QosQueue": EvaluateQosQueue,

	"RouteAggregate": EvaluateRouteAggregate,

	"RouteTable": EvaluateRouteTable,

	"RouteTarget": EvaluateRouteTarget,

	"RoutingInstance": EvaluateRoutingInstance,

	"RoutingPolicy": EvaluateRoutingPolicy,

	"SecurityGroup": EvaluateSecurityGroup,

	"SecurityLoggingObject": EvaluateSecurityLoggingObject,

	"ServiceAppliance": EvaluateServiceAppliance,

	"ServiceApplianceSet": EvaluateServiceApplianceSet,

	"ServiceConnectionModule": EvaluateServiceConnectionModule,

	"ServiceEndpoint": EvaluateServiceEndpoint,

	"ServiceGroup": EvaluateServiceGroup,

	"ServiceHealthCheck": EvaluateServiceHealthCheck,

	"ServiceInstance": EvaluateServiceInstance,

	"ServiceObject": EvaluateServiceObject,

	"ServiceTemplate": EvaluateServiceTemplate,

	"StructuredSyslogApplicationRecord": EvaluateStructuredSyslogApplicationRecord,

	"StructuredSyslogConfig": EvaluateStructuredSyslogConfig,

	"StructuredSyslogHostnameRecord": EvaluateStructuredSyslogHostnameRecord,

	"StructuredSyslogMessage": EvaluateStructuredSyslogMessage,

	"StructuredSyslogSLAProfile": EvaluateStructuredSyslogSLAProfile,

	"SubCluster": EvaluateSubCluster,

	"Subnet": EvaluateSubnet,

	"Tag": EvaluateTag,

	"TagType": EvaluateTagType,

	"VirtualDNSRecord": EvaluateVirtualDNSRecord,

	"VirtualDNS": EvaluateVirtualDNS,

	"VirtualIP": EvaluateVirtualIP,

	"VirtualMachineInterface": EvaluateVirtualMachineInterface,

	"VirtualMachine": EvaluateVirtualMachine,

	"VirtualNetwork": EvaluateVirtualNetwork,

	"VirtualRouter": EvaluateVirtualRouter,

	"AppformixNode": EvaluateAppformixNode,

	"BaremetalNode": EvaluateBaremetalNode,

	"BaremetalPort": EvaluateBaremetalPort,

	"ContrailAnalyticsDatabaseNode": EvaluateContrailAnalyticsDatabaseNode,

	"ContrailAnalyticsNode": EvaluateContrailAnalyticsNode,

	"ContrailCluster": EvaluateContrailCluster,

	"ContrailConfigDatabaseNode": EvaluateContrailConfigDatabaseNode,

	"ContrailConfigNode": EvaluateContrailConfigNode,

	"ContrailControlNode": EvaluateContrailControlNode,

	"ContrailServiceNode": EvaluateContrailServiceNode,

	"ContrailStorageNode": EvaluateContrailStorageNode,

	"ContrailVrouterNode": EvaluateContrailVrouterNode,

	"ContrailWebuiNode": EvaluateContrailWebuiNode,

	"Credential": EvaluateCredential,

	"Dashboard": EvaluateDashboard,

	"Endpoint": EvaluateEndpoint,

	"Flavor": EvaluateFlavor,

	"OsImage": EvaluateOsImage,

	"Keypair": EvaluateKeypair,

	"KubernetesCluster": EvaluateKubernetesCluster,

	"KubernetesMasterNode": EvaluateKubernetesMasterNode,

	"KubernetesNode": EvaluateKubernetesNode,

	"Location": EvaluateLocation,

	"Node": EvaluateNode,

	"OpenstackCluster": EvaluateOpenstackCluster,

	"OpenstackComputeNode": EvaluateOpenstackComputeNode,

	"OpenstackControlNode": EvaluateOpenstackControlNode,

	"OpenstackMonitoringNode": EvaluateOpenstackMonitoringNode,

	"OpenstackNetworkNode": EvaluateOpenstackNetworkNode,

	"OpenstackStorageNode": EvaluateOpenstackStorageNode,

	"Port": EvaluatePort,

	"PortGroup": EvaluatePortGroup,

	"Server": EvaluateServer,

	"VPNGroup": EvaluateVPNGroup,

	"Widget": EvaluateWidget,
}

// EvaluateDependencies evaluates the dependencies upon object change
func EvaluateDependencies(obj interface{}, resource_name string) {
	log.Printf("EvaluateDependencies called for (%s): \n", resource_name)
	d := dependencies.NewDependencyProcessor(compilationif.ObjsCache)
	d.Evaluate(obj, resource_name, "Self")
	objMap := d.GetResources()

	for objTypeKey, objList := range objMap {
		log.Printf("Processing ObjType[%s] \n", objTypeKey)
		for objUuid, objVal := range objList {
			log.Printf("Processing ObjUuid[%s] \n", objUuid)
			log.Printf("Processing Object[%v] \n", objVal)

			// Look up Intent object by objUuid
			objMap, ok := compilationif.ObjsCache.Load(objTypeKey + "Intent")
			log.Printf("ObjMap [%v] \n", objMap)
			if ok {
				if intentObj, ok := objMap.(*sync.Map).Load(objUuid); ok {
					log.Printf("Intent Object[%v] \n", intentObj)
					resource_functions[objTypeKey](intentObj)
				}
			}
		}
	}
}
