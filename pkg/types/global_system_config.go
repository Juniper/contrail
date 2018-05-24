package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

//CreateGlobalSystemConfig by design should never be called.
// GlobalSystemConfig can only be created by DBInit, never by user request
func (sv *ContrailTypeLogicService) CreateGlobalSystemConfig(
	ctx context.Context,
	request *models.CreateGlobalSystemConfigRequest) (resp *models.CreateGlobalSystemConfigResponse, err error) {
	return nil, common.ErrorBadRequest("Trying to call create GlobalSystemConfig outside of DBInit")
}

//UpdateGlobalSystemConfig performs type specific checks for update.
func (sv *ContrailTypeLogicService) UpdateGlobalSystemConfig(
	ctx context.Context,
	request *models.UpdateGlobalSystemConfigRequest) (*models.UpdateGlobalSystemConfigResponse, error) {
	var resp *models.UpdateGlobalSystemConfigResponse

	err := db.DoInTransaction(
		ctx,
		sv.DB.DB(),
		func(ctx context.Context) (err error) {
			oldObjResp, err := sv.DB.GetGlobalSystemConfig(ctx, &models.GetGlobalSystemConfigRequest{ID: request.GetGlobalSystemConfig().GetUUID()})
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

			err = sv.checkAsn(ctx, updateObj, oldObj)
			if err != nil {
				return err
			}

			err = sv.checkBgpaasPorts(ctx, updateObj, oldObj)
			if err != nil {
				return err
			}

			resp, err = sv.Next().UpdateGlobalSystemConfig(ctx, request)
			if err != nil {
				return err
			}

			return nil
		})

	if err != nil {
		return nil, errors.Wrap(err, "GlobalSystemConfig update validation failed")
	}
	return resp, nil
}

func (sv *ContrailTypeLogicService) checkUdc(updateObj *models.GlobalSystemConfig) (err error) {
	if updateObj.UserDefinedLogStatistics == nil {
		return nil
	}
	return updateObj.UserDefinedLogStatistics.ValidateRegexps()
}

func (sv *ContrailTypeLogicService) checkAsn(ctx context.Context, updateObj *models.GlobalSystemConfig, obj *models.GlobalSystemConfig) (err error) {

	globalAsn := updateObj.AutonomousSystem
	if globalAsn == 0 {
		return nil
	}

	vnList, err := sv.DB.ListVirtualNetwork(ctx, &models.ListVirtualNetworkRequest{Spec: &models.ListSpec{Fields: []string{"route_target_list"}}})
	if err != nil {
		return err
	}

	for _, vn := range vnList.VirtualNetworks {
		rtList := vn.RouteTargetList
		for _, rt := range rtList.RouteTarget {
			userDefed, err := models.IsStringRouteTargetUserDefined(rt, globalAsn)
			if err != nil {
				return err
			}
			if userDefed {
				return errors.Errorf("Cannot update ASN as there are ASN based Virtual Networks")
			}
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkBgpaasPorts(ctx context.Context, updateObj *models.GlobalSystemConfig, obj *models.GlobalSystemConfig) (err error) {
	if updateObj.BgpaasParameters == nil {
		return nil
	}

	err = updateObj.BgpaasParameters.ValidatePortRange()
	if err != nil {
		return err
	}

	newPortsRange := updateObj.GetBgpaasParameters()
	oldPortsRange := models.GetDefaultBGPaaServiceParameters()

	if obj.BgpaasParameters != nil {
		oldPortsRange = obj.BgpaasParameters
	}

	bgpaasList, err := sv.DB.ListBGPAsAService(ctx, &models.ListBGPAsAServiceRequest{Spec: &models.ListSpec{Count: true}})

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
