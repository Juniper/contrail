package types

import (
	"fmt"

	"regexp"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

//CreateGlobalSystemConfig by design should never be called.
// GlobalSystemConfig can only be created by DBInit, never by user request
func (sv *ContrailTypeLogicService) CreateGlobalSystemConfig(
	ctx context.Context,
	request *models.CreateGlobalSystemConfigRequest) (resp *models.CreateGlobalSystemConfigResponse, err error) {
	return nil, fmt.Errorf("Trying to call create GlobalSystemConfig outside of DBInit")
}

//UpdateGlobalSystemConfig performs type specific checks for update
func (sv *ContrailTypeLogicService) UpdateGlobalSystemConfig(
	ctx context.Context,
	request *models.UpdateGlobalSystemConfigRequest) (resp *models.UpdateGlobalSystemConfigResponse, err error) {

	err = db.DoInTransaction(
		ctx,
		sv.DB.DB(),
		func(ctx context.Context) (inTxErr error) {
			// Probably uuid required
			oldObj, inTxErr := sv.DB.GetGlobalSystemConfig(ctx, &models.GetGlobalSystemConfigRequest{})

			if inTxErr != nil {
				return inTxErr
			}
			if oldObj == nil {
				return fmt.Errorf("No GlobalSystemConfig found to update")
			}

			inTxErr = sv.checkUdc(request)
			if inTxErr != nil {
				return inTxErr
			}

			inTxErr = sv.checkAsn(ctx, request, oldObj.GlobalSystemConfig)
			if inTxErr != nil {
				return inTxErr
			}

			inTxErr = sv.checkBgpaasPorts(ctx, request, oldObj.GlobalSystemConfig)
			if inTxErr != nil {
				return inTxErr
			}

			if sv.Next() != nil {

				resp, inTxErr = sv.Next().UpdateGlobalSystemConfig(ctx, request)
				if inTxErr != nil {
					return inTxErr
				}
			}

			return nil
		})

	return resp, err
}

func (sv *ContrailTypeLogicService) checkUdc(request *models.UpdateGlobalSystemConfigRequest) (err error) {
	if request.GlobalSystemConfig.UserDefinedLogStatistics == nil {
		return nil
	}
	for _, udc := range request.GlobalSystemConfig.UserDefinedLogStatistics.Statlist {
		_, err = regexp.Compile(udc.Pattern)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO (Jan Darowski) DELETE AFTER ROUTE_TARGET.GO CREATED, ADD TESTS
// THIS IS MOCK VERSION
func (sv *ContrailTypeLogicService) isRouteTargetUserDefined(routeTarget string, asn int64) (bool, error) {
	return true, nil
}

func (sv *ContrailTypeLogicService) checkAsn(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest, obj *models.GlobalSystemConfig) (err error) {

	globalAsn := request.GlobalSystemConfig.AutonomousSystem
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
			userDefed, err := sv.isRouteTargetUserDefined(rt, globalAsn) // Temporary implementation
			if err != nil {
				return err
			}
			if userDefed {
				return fmt.Errorf("Cannot update ASN as there are ASN based Virtual Networks")
			}
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkBgpaasPorts(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest, obj *models.GlobalSystemConfig) (err error) {
	newObj := request.GlobalSystemConfig
	if newObj.BgpaasParameters == nil {
		return nil
	}

	err = newObj.BgpaasParameters.ValidatePortRange()
	if err != nil {
		return err
	}

	newPortsRange := newObj.GetBgpaasParameters()
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
			return fmt.Errorf("BGP Port range cannot be shrunk. Old: (%d %d) New: (%d %d)",
				oldPortsRange.PortStart,
				oldPortsRange.PortEnd,
				newPortsRange.PortStart,
				newPortsRange.PortEnd)
		}
	}
	return nil
}
