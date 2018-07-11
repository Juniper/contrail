package types

import (
	"github.com/Juniper/contrail/pkg/services"
	"golang.org/x/net/context"
)

//CreateInstanceIP ...
func (sv *ContrailTypeLogicService) CreateInstanceIP(
	ctx context.Context,
	request *services.CreateInstanceIPRequest) (*services.CreateInstanceIPResponse, error) {

	var response *services.CreateInstanceIPResponse
	//instanceIP := request.GetInstanceIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

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
