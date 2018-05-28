package services

import (
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

// QuotaCheckerService implements counting logic for specific resources under given project
type QuotaCheckerService struct {
	BaseQuotaCheckerService
}

func NewQuotaCheckerService(db serviceif.Service) QuotaCheckerService {
	return QuotaCheckerService{
		BaseQuotaCheckerService: BaseQuotaCheckerService{
			db: db,
		},
	}
}

func (svc *QuotaCheckerService) CountAccessControlList(ctx context.Context, obj *models.AccessControlList, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountBGPRouter(ctx context.Context, obj *models.BGPRouter, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountFloatingIPPool(ctx context.Context, obj *models.FloatingIPPool, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountFloatingIP(ctx context.Context, obj *models.FloatingIP, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountGlobalVrouterConfig(ctx context.Context, obj *models.GlobalVrouterConfig, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountInstanceIP(ctx context.Context, obj *models.InstanceIP, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountLoadbalancerHealthmonitor(ctx context.Context, obj *models.LoadbalancerHealthmonitor, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountLoadbalancerMember(ctx context.Context, obj *models.LoadbalancerMember, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountLoadbalancerPool(ctx context.Context, obj *models.LoadbalancerPool, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountLogicalRouter(ctx context.Context, obj *models.LogicalRouter, prj *models.Project) (int64, error) {
	return 0, nil
}

/*
func (svc *QuotaCheckerService) CreateNetworkIpam(ctx context.Context, request *models.CreateNetworkIpamRequest) (*models.CreateNetworkIpamResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateNetworkIpam(ctx, request)
}
*/
func (svc *QuotaCheckerService) CountNetworkIpam(ctx context.Context, obj *models.NetworkIpam, prj *models.Project) (int64, error) {
	return 0, nil
}

/*
func (svc *QuotaCheckerService) CreateNetworkPolicy(ctx context.Context, request *models.CreateNetworkPolicyRequest) (*models.CreateNetworkPolicyResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateNetworkPolicy(ctx, request)
}
*/
func (svc *QuotaCheckerService) CountNetworkPolicy(ctx context.Context, obj *models.NetworkPolicy, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountRouteTable(ctx context.Context, obj *models.RouteTable, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountSecurityGroup(ctx context.Context, obj *models.SecurityGroup, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountSecurityLoggingObject(ctx context.Context, obj *models.SecurityLoggingObject, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountServiceInstance(ctx context.Context, obj *models.ServiceInstance, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountServiceTemplate(ctx context.Context, obj *models.ServiceTemplate, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountSubnet(ctx context.Context, obj *models.Subnet, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountVirtualDNSRecord(ctx context.Context, obj *models.VirtualDNSRecord, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountVirtualDNS(ctx context.Context, obj *models.VirtualDNS, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountVirtualIP(ctx context.Context, obj *models.VirtualIP, prj *models.Project) (int64, error) {
	return 0, nil
}
func (svc *QuotaCheckerService) CountVirtualMachineInterface(ctx context.Context, obj *models.VirtualMachineInterface, prj *models.Project) (int64, error) {
	return 0, nil
}

/*
func (svc *QuotaCheckerService) CreateVirtualNetwork(ctx context.Context, request *models.CreateVirtualNetworkRequest) (*models.CreateVirtualNetworkResponse, error) {
	// TODO This need to be implemented with logic specific to this resource - this type has no parent
	return svc.BaseService.CreateVirtualNetwork(ctx, request)
}
*/
// CountVirtualNetwork counts all the virtual networks that are in a project
func (svc *QuotaCheckerService) CountVirtualNetwork(ctx context.Context, obj *models.VirtualNetwork, prj *models.Project) (int64, error) {
	log.Debugf("Counting VirtualNetworks (VNs) in Project %v - %+v", prj.GetUUID(), prj.VirtualNetworks)
	count := int64(len(prj.GetVirtualNetworks()))
	return count, nil
	//return len(prj.VirtualNetworks), nil
}

func (svc *QuotaCheckerService) CountVirtualRouter(ctx context.Context, obj *models.VirtualRouter, prj *models.Project) (int64, error) {
	return 0, nil
}
