package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateFirewallPolicy performs pre create type specific validation
func (sv *ContrailTypeLogicService) CreateFirewallPolicy(
	ctx context.Context,
	request *services.CreateFirewallPolicyRequest,
) (*services.CreateFirewallPolicyResponse, error) {

	var response *services.CreateFirewallPolicyResponse
	fp := request.GetFirewallPolicy()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			if err := checkDraftModeState(ctx, fp); err != nil {
				return err
			}

			if err := fp.CheckAssociatedRefsInSameScope(fp.GetFQName()); err != nil {
				return err
			}

			response, err = sv.Next().CreateFirewallPolicy(ctx, request)
			return err
		})

	return response, err
}

// UpdateFirewallPolicy does pre update checks for the firewall policy
func (sv *ContrailTypeLogicService) UpdateFirewallPolicy(
	ctx context.Context,
	request *services.UpdateFirewallPolicyRequest,
) (*services.UpdateFirewallPolicyResponse, error) {

	var response *services.UpdateFirewallPolicyResponse
	fp := request.GetFirewallPolicy()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			if err := checkDraftModeState(ctx, fp); err != nil {
				return err
			}

			fqName, err := sv.getFirewallPolicyFQName(ctx, fp)
			if err != nil {
				return err
			}

			if err := fp.CheckAssociatedRefsInSameScope(fqName); err != nil {
				return err
			}

			response, err = sv.Next().UpdateFirewallPolicy(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getFirewallPolicyFQName(
	ctx context.Context, fp *models.FirewallPolicy,
) ([]string, error) {

	if len(fp.GetFQName()) > 0 {
		return fp.GetFQName(), nil
	}

	firewallPolicyResponse, err := sv.ReadService.GetFirewallPolicy(
		ctx,
		&services.GetFirewallPolicyRequest{
			ID:     fp.GetUUID(),
			Fields: []string{models.FirewallPolicyFieldFQName},
		},
	)

	return firewallPolicyResponse.GetFirewallPolicy().GetFQName(), err
}
