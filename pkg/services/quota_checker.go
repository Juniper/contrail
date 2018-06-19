package services

import (
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/models"
)

// QuotaCheckerService implements counting logic for specific resources in given project.
type QuotaCheckerService struct {
	BaseQuotaCheckerService
}

// NewQuotaCheckerService creates QuotaCheckerService.
func NewQuotaCheckerService(db Service) QuotaCheckerService {
	return QuotaCheckerService{
		BaseQuotaCheckerService: BaseQuotaCheckerService{
			db: db,
		},
	}
}

// CountAccessControlList counts all AccessControlList resources in given project.
func (svc *QuotaCheckerService) CountAccessControlList(ctx context.Context, obj *models.AccessControlList,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountBGPRouter counts all BGPRouter resources in given project.
func (svc *QuotaCheckerService) CountBGPRouter(ctx context.Context, obj *models.BGPRouter,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountFloatingIPPool counts all FloatingIPPool resources in given project.
func (svc *QuotaCheckerService) CountFloatingIPPool(ctx context.Context, obj *models.FloatingIPPool,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountFloatingIP counts all FloatingIP resources in given project.
func (svc *QuotaCheckerService) CountFloatingIP(ctx context.Context, obj *models.FloatingIP,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountGlobalVrouterConfig counts all GlobalVrouterConfig resources in given project.
func (svc *QuotaCheckerService) CountGlobalVrouterConfig(ctx context.Context, obj *models.GlobalVrouterConfig,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountInstanceIP counts all InstanceIP resources in given project.
func (svc *QuotaCheckerService) CountInstanceIP(ctx context.Context, obj *models.InstanceIP,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountLoadbalancerHealthmonitor counts all LoadbalancerHealthmonitor resources in given project.
func (svc *QuotaCheckerService) CountLoadbalancerHealthmonitor(ctx context.Context,
	obj *models.LoadbalancerHealthmonitor, prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountLoadbalancerMember counts all LoadbalancerMember resources in given project.
func (svc *QuotaCheckerService) CountLoadbalancerMember(ctx context.Context, obj *models.LoadbalancerMember,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountLoadbalancerPool counts all LoadbalancerPool resources in given project.
func (svc *QuotaCheckerService) CountLoadbalancerPool(ctx context.Context, obj *models.LoadbalancerPool,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountLogicalRouter counts all LogicalRouter resources in given project.
func (svc *QuotaCheckerService) CountLogicalRouter(ctx context.Context, obj *models.LogicalRouter,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

/*
func (svc *QuotaCheckerService) CreateNetworkIpam(ctx context.Context, request *models.CreateNetworkIpamRequest) (
*models.CreateNetworkIpamResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateNetworkIpam(ctx, request)
}
*/

// CountNetworkIpam counts all NetworkIpam resources in given project.
func (svc *QuotaCheckerService) CountNetworkIpam(ctx context.Context, obj *models.NetworkIpam,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

/*
func (svc *QuotaCheckerService) CreateNetworkPolicy(ctx context.Context, request *models.CreateNetworkPolicyRequest) (
*models.CreateNetworkPolicyResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateNetworkPolicy(ctx, request)
}
*/

// CountNetworkPolicy counts all NetworkPolicy resources in given project.
func (svc *QuotaCheckerService) CountNetworkPolicy(ctx context.Context, obj *models.NetworkPolicy,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountRouteTable counts all RouteTable resources in given project.
func (svc *QuotaCheckerService) CountRouteTable(ctx context.Context, obj *models.RouteTable,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountSecurityGroup counts all SecurityGroup resources in given project.
func (svc *QuotaCheckerService) CountSecurityGroup(ctx context.Context, obj *models.SecurityGroup,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountSecurityLoggingObject counts all SecurityLoggingObject resources in given project.
func (svc *QuotaCheckerService) CountSecurityLoggingObject(ctx context.Context, obj *models.SecurityLoggingObject,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountServiceInstance counts all ServiceInstance resources in given project.
func (svc *QuotaCheckerService) CountServiceInstance(ctx context.Context, obj *models.ServiceInstance,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountServiceTemplate counts all ServiceTemplate resources in given project.
func (svc *QuotaCheckerService) CountServiceTemplate(ctx context.Context, obj *models.ServiceTemplate,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountSubnet counts all Subnet resources in given project.
func (svc *QuotaCheckerService) CountSubnet(ctx context.Context, obj *models.Subnet,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualDNSRecord counts all VirtualDNSRecord resources in given project.
func (svc *QuotaCheckerService) CountVirtualDNSRecord(ctx context.Context, obj *models.VirtualDNSRecord,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualDNS counts all VirtualDNS resources in given project.
func (svc *QuotaCheckerService) CountVirtualDNS(ctx context.Context, obj *models.VirtualDNS,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualIP counts all VirtualIP resources in given project.
func (svc *QuotaCheckerService) CountVirtualIP(ctx context.Context, obj *models.VirtualIP,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualMachineInterface counts all VirtualMachineInterface resources in given project.
func (svc *QuotaCheckerService) CountVirtualMachineInterface(ctx context.Context, obj *models.VirtualMachineInterface,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}

// CountVirtualNetwork counts all VirtualNetwork resources in given project.
func (svc *QuotaCheckerService) CountVirtualNetwork(_ context.Context, _ *models.VirtualNetwork,
	prj *models.Project) (int64, error) {
	log.Debugf("Counting VirtualNetworks (VNs) in Project %v - %+v", prj.GetUUID(), prj.VirtualNetworks)
	count := int64(len(prj.GetVirtualNetworks()))
	return count, nil
}

// CountVirtualRouter counts all VirtualRouter resources in given project.
func (svc *QuotaCheckerService) CountVirtualRouter(ctx context.Context, obj *models.VirtualRouter,
	prj *models.Project) (int64, error) {
	// TODO will be implemented in separate patches (k.renczewski)
	return 0, nil
}
