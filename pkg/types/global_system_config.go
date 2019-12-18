package types

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const defaultGSCName = "default-global-system-config"

// CreateGlobalSystemConfig by design should never be called.
// GlobalSystemConfig can only be created by DBInit, never by user request
func (sv *ContrailTypeLogicService) CreateGlobalSystemConfig(
	ctx context.Context, request *services.CreateGlobalSystemConfigRequest,
) (resp *services.CreateGlobalSystemConfigResponse, err error) {
	return nil, errutil.ErrorBadRequest("Trying to call create GlobalSystemConfig outside of DBInit")
}

// UpdateGlobalSystemConfig performs type specific checks for update.
func (sv *ContrailTypeLogicService) UpdateGlobalSystemConfig(
	ctx context.Context, request *services.UpdateGlobalSystemConfigRequest,
) (*services.UpdateGlobalSystemConfigResponse, error) {
	var resp *services.UpdateGlobalSystemConfigResponse

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) (err error) {
			// TODO(Jan Darowski) JBE-431 - Add proper uuid update when using draft object
			// (enable_security_policy_draft)
			oldObjResp, err := sv.ReadService.GetGlobalSystemConfig(ctx, &services.GetGlobalSystemConfigRequest{
				ID: request.GetGlobalSystemConfig().GetUUID(),
			})
			updateObj := request.GlobalSystemConfig

			if err != nil {
				return err
			}

			if oldObjResp == nil {
				return errors.Errorf("No GlobalSystemConfig found to update")
			}
			oldObj := oldObjResp.GlobalSystemConfig

			err = sv.checkUdc(updateObj)
			if err != nil {
				return err
			}

			err = sv.checkAsn(ctx, updateObj)
			if err != nil {
				return err
			}

			err = sv.checkBgpaasPorts(ctx, updateObj, oldObj)
			if err != nil {
				return err
			}

			resp, err = sv.BaseService.UpdateGlobalSystemConfig(ctx, request)
			return err

		})

	if err != nil {
		return nil, errors.Wrap(err, "GlobalSystemConfig update validation failed")
	}
	return resp, nil
}

func (sv *ContrailTypeLogicService) checkUdc(obj *models.GlobalSystemConfig) (err error) {
	if obj.UserDefinedLogStatistics == nil {
		return nil
	}
	return obj.UserDefinedLogStatistics.ValidateRegexps()
}

func (sv *ContrailTypeLogicService) checkAsn(ctx context.Context, updateObj *models.GlobalSystemConfig) (err error) {

	globalAsn := updateObj.AutonomousSystem
	if globalAsn == 0 {
		return nil
	}

	vnList, err := sv.ReadService.ListVirtualNetwork(ctx, &services.ListVirtualNetworkRequest{Spec: &baseservices.ListSpec{
		Fields: []string{models.VirtualNetworkFieldRouteTargetList}},
	})
	if err != nil {
		return err
	}

	var multiError errutil.MultiError
	for _, vn := range vnList.VirtualNetworks {
		rtList := vn.RouteTargetList
		for _, rt := range rtList.RouteTarget {
			userDefined, err := models.IsRouteTargetUserDefined(rt, globalAsn)
			if err != nil {
				return err
			}
			if !userDefined {
				multiError = append(multiError, errors.Errorf("\t- %s (%s) have route target %s configured\n",
					vn.FQName, vn.UUID, rt))
			}
		}
	}

	if multiError != nil {
		return errors.Wrapf(multiError, "Virtual networks are configured with a route target with this ASN %d "+
			"and route target value in the same range as used by automatically allocated route targets:\n", globalAsn)
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkBgpaasPorts(ctx context.Context, updateObj *models.GlobalSystemConfig,
	oldObj *models.GlobalSystemConfig) (err error) {
	if updateObj.BgpaasParameters == nil {
		return nil
	}

	err = updateObj.BgpaasParameters.ValidatePortRange()
	if err != nil {
		return err
	}

	newPortsRange := updateObj.GetBgpaasParameters()
	oldPortsRange := models.GetDefaultBGPaaServiceParameters()

	if oldObj.BgpaasParameters != nil {
		oldPortsRange = oldObj.BgpaasParameters
	}

	bgpaasList, err := sv.ReadService.ListBGPAsAService(ctx, &services.ListBGPAsAServiceRequest{
		Spec: &baseservices.ListSpec{Count: true},
	})

	if err != nil {
		return err
	}

	if len(bgpaasList.BGPAsAServices) > 0 {
		if !newPortsRange.EnclosesRange(oldPortsRange) {
			return errors.Errorf("BGP Port range cannot be shrunk. Old: (%d %d) New: (%d %d)",
				oldPortsRange.PortStart,
				oldPortsRange.PortEnd,
				newPortsRange.PortStart,
				newPortsRange.PortEnd)
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) getDefaultGlobalSystemConfigUUID(ctx context.Context) (string, error) {
	if sv.defaultGSCUUID != "" {
		return sv.defaultGSCUUID, nil
	}

	defaultGSCFqName := sv.getDefaultGlobalSystemConfigFqName()
	defaultGSCMeta, err := sv.MetadataGetter.GetMetadata(
		ctx,
		basemodels.Metadata{
			FQName: defaultGSCFqName,
			Type:   models.KindGlobalSystemConfig,
		},
	)
	if err != nil {
		return "", errutil.ErrorBadRequestf("Cannot resolve Global System Config with FQName %v: %v",
			defaultGSCFqName, err)
	}
	sv.defaultGSCUUID = defaultGSCMeta.UUID
	return sv.defaultGSCUUID, nil
}

func (sv *ContrailTypeLogicService) getDefaultGlobalSystemConfigFqName() []string {
	return []string{defaultGSCName}
}
