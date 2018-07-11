package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"golang.org/x/net/context"
)

//CreateInstanceIP ...
func (sv *ContrailTypeLogicService) CreateInstanceIP(
	ctx context.Context,
	request *services.CreateInstanceIPRequest) (*services.CreateInstanceIPResponse, error) {

	var response *services.CreateInstanceIPResponse
	instanceIP := request.GetInstanceIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			virtualNetworkRefs := instanceIP.GetVirtualNetworkRefs() //TODO sprawdzac czy zwraca nil?
			virtualRouterRefs := instanceIP.GetVirtualRouterRefs()
			networkIpamRefs := instanceIP.GetNetworkIpamRefs()
			var virtualNetwork *models.VirtualNetwork
			var ipamRefs []*models.VirtualRouterNetworkIpamRef

			if virtualRouterRefs != nil && networkIpamRefs != nil {
				return common.ErrorBadRequestf("virtual_router_refs and network_ipam_refs are not allowed")
			}

			if virtualRouterRefs != nil && virtualNetworkRefs != nil {
				return common.ErrorBadRequestf("virtual_router_refs and virtual_network_refs are not allowed")
			}

			if virtualNetworkRefs != nil {
				virtualNetwork, err := sv.getVirtualNetworkFromInstanceIP(ctx, instanceIP)

				//TODO ip_fabric and link_local fq_name check

				ipamRefs = nil //TODO co z tym?

			} else {
				virtualNetwork = nil //TODO chyba mozna usunac
				if virtualRouterRefs != nil {
					if len(virtualRouterRefs) > 1 {
						return common.ErrorBadRequestf("Instance-ip can not refer to multiple vrouters")
					}

					networkImapRefsTemp, err := sv.getIpamRefsFromInstanceIPVirtualRouterRef(ctx, virtualRouterRefs[0])

					ipamRefs = virtualRouterResponse.GetVirtualRouter().GetNetworkIpamRefs()
				} else {
					ipamRefs = instanceIP.GetNetworkIpamRefs()
				}

			}

			// allocateIPRequest := &ipam.AllocateIPRequest{
			// 	VirtualNetwork: virtualNetwork,
			// 	SubnetUUID:     floatingIPPoolSubnetUUID,
			// 	IPAddress:      floatingIP.GetFloatingIPAddress(),
			// }

			// loatingIPAddress, _, err = sv.AddressManager.AllocateIP(ctx, allocateIPParams)

			response, err = sv.BaseService.CreateInstanceIP(ctx, request)
			return err
		})

	return response, err
}

//DeleteInstanceIP ...
func (sv *ContrailTypeLogicService) DeleteInstanceIP(
	ctx context.Context,
	request *services.DeleteInstanceIPRequest) (*services.DeleteInstanceIPResponse, error) {

	var response *services.DeleteInstanceIPResponse
	//instanceIP := request.GetInstanceIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			response, err = sv.BaseService.DeleteInstanceIP(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromInstanceIP(
	ctx context.Context, instanceIP *models.InstanceIP) (*models.VirtualNetwork, error) {

	instanceIPVirtualNetworkRef := instanceIP.GetVirtualNetworkRefs()[0] //TODO dodac do glownej?
	virtualNetworkFQName := instanceIPVirtualNetworkRef.GetTo()

	metadata, err := sv.FQNameUUIDTranslator.TranslateBetweenFQNameUUID(ctx, "", virtualNetworkFQName)
	if err != nil {
		return nil, err
	}

	virtualNetworkResponse, err := sv.DataService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: metadata.UUID,
		})
	if err != nil {
		return nil, err
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) getIpamRefsFromInstanceIPVirtualRouterRef(
	ctx context.Context, virtualRouterRef *models.InstanceIPVirtualRouterRef) ([]*models.VirtualRouterNetworkIpamRef, error) {

	virtualRouterFQName := virtualRouterRef.GetTo()

	metadata, err := sv.FQNameUUIDTranslator.TranslateBetweenFQNameUUID(ctx, "", virtualRouterFQName)
	if err != nil {
		return nil, err
	}

	virtualRouterResponse, err := sv.DataService.GetVirtualRouter(ctx,
		&services.GetVirtualRouterRequest{
			ID: metadata.UUID,
		})
	if err != nil {
		return nil, err
	}

	return virtualRouterResponse.GetVirtualRouter().GetNetworkIpamRefs(), nil
}
