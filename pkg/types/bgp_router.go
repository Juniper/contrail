package types

import (
	"context"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateBGPRouter creates BGP Router and checks if provided asn
// is equal to asn in linked SubCluster.
func (sv *ContrailTypeLogicService) CreateBGPRouter(
	ctx context.Context, request *services.CreateBGPRouterRequest,
) (*services.CreateBGPRouterResponse, error) {

	var response *services.CreateBGPRouterResponse
	bgpRouter := request.GetBGPRouter()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err := sv.checkSubClusterASN(ctx, bgpRouter)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.CreateBGPRouter(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getSubCluster(
	ctx context.Context, id string,
) (*models.SubCluster, error) {

	subClusterResponse, err := sv.ReadService.GetSubCluster(
		ctx,
		&services.GetSubClusterRequest{
			ID: id,
		},
	)

	return subClusterResponse.GetSubCluster(), err
}

func (sv *ContrailTypeLogicService) checkSubClusterASN(
	ctx context.Context, bgpRouter *models.BGPRouter,
) error {

	asn := bgpRouter.GetBGPRouterParameters().GetAutonomousSystem()
	subClusterRefs := bgpRouter.GetSubClusterRefs()

	if len(subClusterRefs) != 0 && asn != 0 {
		subCluster, err := sv.getSubCluster(ctx, subClusterRefs[0].GetUUID())
		if err != nil {
			return err
		}

		if asn != subCluster.GetSubClusterAsn() {
			return errutil.ErrorBadRequestf("subcluster asn and bgp asn should be same")
		}
	}

	return nil
}
