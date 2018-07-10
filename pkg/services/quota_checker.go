package services

import (
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/models"
)

// QuotaCounter implements counting logic for specific resources in given project.
type QuotaCounter struct {
}

// NewQuotaCheckerService creates QuotaCheckerService.
func NewQuotaCheckerService(db Service) *QuotaCheckerService {
	return &QuotaCheckerService{
		db:           db,
		quotaCounter: &QuotaCounter{},
	}
}

// CountAccessControlList counts all AccessControlList resources in given project.
func (svc *QuotaCounter) CountAccessControlList(
	ctx context.Context, obj *models.AccessControlList, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountBGPRouter counts all BGPRouter resources in given project.
func (svc *QuotaCounter) CountBGPRouter(
	ctx context.Context, obj *models.BGPRouter, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountFloatingIPPool counts all FloatingIPPool resources in given project.
func (svc *QuotaCounter) CountFloatingIPPool(
	ctx context.Context, obj *models.FloatingIPPool, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountFloatingIP counts all FloatingIP resources in given project.
func (svc *QuotaCounter) CountFloatingIP(
	ctx context.Context, obj *models.FloatingIP, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountGlobalVrouterConfig counts all GlobalVrouterConfig resources in given project.
func (svc *QuotaCounter) CountGlobalVrouterConfig(
	ctx context.Context, obj *models.GlobalVrouterConfig, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountInstanceIP counts all InstanceIP resources in given project.
func (svc *QuotaCounter) CountInstanceIP(
	ctx context.Context, obj *models.InstanceIP, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountLoadbalancerHealthmonitor counts all LoadbalancerHealthmonitor resources in given project.
func (svc *QuotaCounter) CountLoadbalancerHealthmonitor(
	ctx context.Context, obj *models.LoadbalancerHealthmonitor, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting LoadbalancerHealthmonitors in Project %v - %+v", prj.GetUUID(), prj.LoadbalancerHealthmonitors)
	count := int64(len(prj.GetLoadbalancerHealthmonitors()))
	return count, nil
}

// CountLoadbalancerMember counts all LoadbalancerMember resources in given project.
func (svc *QuotaCounter) CountLoadbalancerMember(
	ctx context.Context, obj *models.LoadbalancerMember, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountLoadbalancerPool counts all LoadbalancerPool resources in given project.
func (svc *QuotaCounter) CountLoadbalancerPool(
	ctx context.Context, obj *models.LoadbalancerPool, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting LoadbalancerPools in Project %v - %+v", prj.GetUUID(), prj.LoadbalancerPools)
	count := int64(len(prj.GetLoadbalancerPools()))
	return count, nil
}

// CountLogicalRouter counts all LogicalRouter resources in given project.
func (svc *QuotaCounter) CountLogicalRouter(
	ctx context.Context, obj *models.LogicalRouter, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting LogicalRouters in Project %v - %+v", prj.GetUUID(), prj.LogicalRouters)
	count := int64(len(prj.GetLogicalRouters()))
	return count, nil
}

/*
func (svc *QuotaCheckerService) CreateNetworkIpam(
	ctx context.Context, request *models.CreateNetworkIpamRequest,
) (*models.CreateNetworkIpamResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateNetworkIpam(ctx, request)
}
*/

// CountNetworkIpam counts all NetworkIpam resources in given project.
func (svc *QuotaCounter) CountNetworkIpam(
	ctx context.Context, obj *models.NetworkIpam, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting NetworkIpams in Project %v - %+v", prj.GetUUID(), prj.NetworkIpams)
	count := int64(len(prj.GetNetworkIpams()))
	return count, nil
}

/*
func (svc *QuotaCheckerService) CreateNetworkPolicy(
	ctx context.Context, request *models.CreateNetworkPolicyRequest,
) (*models.CreateNetworkPolicyResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateNetworkPolicy(ctx, request)
}
*/

// CountNetworkPolicy counts all NetworkPolicy resources in given project.
func (svc *QuotaCounter) CountNetworkPolicy(
	ctx context.Context, obj *models.NetworkPolicy, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting NetworkPolicys in Project %v - %+v", prj.GetUUID(), prj.NetworkPolicys)
	count := int64(len(prj.GetNetworkPolicys()))
	return count, nil
}

// CountRouteTable counts all RouteTable resources in given project.
func (svc *QuotaCounter) CountRouteTable(
	ctx context.Context, obj *models.RouteTable, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting RouteTables in Project %v - %+v", prj.GetUUID(), prj.RouteTables)
	count := int64(len(prj.GetRouteTables()))
	return count, nil
}

// CountSecurityGroup counts all SecurityGroup resources in given project.
func (svc *QuotaCounter) CountSecurityGroup(
	ctx context.Context, obj *models.SecurityGroup, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting SecurityGroups in Project %v - %+v", prj.GetUUID(), prj.SecurityGroups)
	count := int64(len(prj.GetSecurityGroups()))
	return count, nil
}

// CountSecurityLoggingObject counts all SecurityLoggingObject resources in given project.
func (svc *QuotaCounter) CountSecurityLoggingObject(
	ctx context.Context, obj *models.SecurityLoggingObject, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting SecurityLoggingObjects in Project %v - %+v", prj.GetUUID(), prj.SecurityLoggingObjects)
	count := int64(len(prj.GetSecurityLoggingObjects()))
	return count, nil
}

// CountServiceInstance counts all ServiceInstance resources in given project.
func (svc *QuotaCounter) CountServiceInstance(
	ctx context.Context, obj *models.ServiceInstance, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting ServiceInstances in Project %v - %+v", prj.GetUUID(), prj.ServiceInstances)
	count := int64(len(prj.GetServiceInstances()))
	return count, nil
}

// CountServiceTemplate counts all ServiceTemplate resources in given project.
func (svc *QuotaCounter) CountServiceTemplate(
	ctx context.Context, obj *models.ServiceTemplate, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountSubnet counts all Subnet resources in given project.
func (svc *QuotaCounter) CountSubnet(
	ctx context.Context, obj *models.Subnet, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualDNSRecord counts all VirtualDNSRecord resources in given project.
func (svc *QuotaCounter) CountVirtualDNSRecord(
	ctx context.Context, obj *models.VirtualDNSRecord, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualDNS counts all VirtualDNS resources in given project.
func (svc *QuotaCounter) CountVirtualDNS(
	ctx context.Context, obj *models.VirtualDNS, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualIP counts all VirtualIP resources in given project.
func (svc *QuotaCounter) CountVirtualIP(
	ctx context.Context, obj *models.VirtualIP, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualMachineInterface counts all VirtualMachineInterface resources in given project.
func (svc *QuotaCounter) CountVirtualMachineInterface(
	ctx context.Context, obj *models.VirtualMachineInterface, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualNetwork counts all VirtualNetwork resources in given project.
func (svc *QuotaCounter) CountVirtualNetwork(
	_ context.Context, _ *models.VirtualNetwork, prj *models.Project,
) (int64, error) {
	log.Debugf("Counting VirtualNetworks (VNs) in Project %v - %+v", prj.GetUUID(), prj.VirtualNetworks)
	count := int64(len(prj.GetVirtualNetworks()))
	return count, nil
}

// CountVirtualRouter counts all VirtualRouter resources in given project.
func (svc *QuotaCounter) CountVirtualRouter(
	ctx context.Context, obj *models.VirtualRouter, prj *models.Project,
) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}
