package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// CreateNetworkPolicy does pre-check for network policy create
func (sv *ContrailTypeLogicService) CreateNetworkPolicy(
	ctx context.Context,
	request *services.CreateNetworkPolicyRequest,
) (response *services.CreateNetworkPolicyResponse, err error) {
	requestNetworkPolicy := request.GetNetworkPolicy()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err = CheckPolicyRules(requestNetworkPolicy.GetNetworkPolicyEntries(), true)
			if err != nil {
				return errors.Wrapf(err, "failed to check Policy Rules")
			}
			response, err = sv.BaseService.CreateNetworkPolicy(ctx, request)
			return err
		})
	return response, err
}

// UpdateNetworkPolicy does pre-check for network policy update
func (sv *ContrailTypeLogicService) UpdateNetworkPolicy(
	ctx context.Context,
	request *services.UpdateNetworkPolicyRequest,
) (response *services.UpdateNetworkPolicyResponse, err error) {
	requestNetworkPolicy := request.GetNetworkPolicy()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err = CheckPolicyRules(requestNetworkPolicy.GetNetworkPolicyEntries(), true)
			if err != nil {
				return errors.Wrapf(err, "failed to check Policy Rules")
			}
			response, err = sv.BaseService.UpdateNetworkPolicy(ctx, request)
			return err
		})
	return response, err
}
