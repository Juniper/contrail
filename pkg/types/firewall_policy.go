package types

import (
	"context"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateFirewallPolicy performs pre create type specific validation.
func (sv *ContrailTypeLogicService) CreateFirewallPolicy(
	ctx context.Context,
	request *services.CreateFirewallPolicyRequest,
) (response *services.CreateFirewallPolicyResponse, err error) {
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			fp := request.GetFirewallPolicy()
			if err = checkDraftModeState(ctx, fp); err != nil {
				return err
			}

			if err = sv.complementRefs(ctx, fp); err != nil {
				return err
			}

			if err = fp.CheckAssociatedRefsInSameScope(fp.GetFQName()); err != nil {
				return errutil.ErrorBadRequest(err.Error())
			}

			response, err = sv.BaseService.CreateFirewallPolicy(ctx, request)
			return err
		})

	return response, err
}

// UpdateFirewallPolicy performs pre update checks for the firewall policy.
func (sv *ContrailTypeLogicService) UpdateFirewallPolicy(
	ctx context.Context,
	request *services.UpdateFirewallPolicyRequest,
) (response *services.UpdateFirewallPolicyResponse, err error) {
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			fp := request.GetFirewallPolicy()
			if err = checkDraftModeState(ctx, fp); err != nil {
				return err
			}

			var fqName []string
			fqName, err = sv.getFirewallPolicyFQName(ctx, fp)
			if err != nil {
				return err
			}

			if err = sv.complementRefs(ctx, fp); err != nil {
				return err
			}

			if err = fp.CheckAssociatedRefsInSameScope(fqName); err != nil {
				return errutil.ErrorBadRequest(err.Error())
			}

			response, err = sv.BaseService.UpdateFirewallPolicy(ctx, request)
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
