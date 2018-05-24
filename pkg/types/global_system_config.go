package types

import (
	"fmt"

	"regexp"

	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

func (service *ContrailTypeLogicService) CreateGlobalSystemConfig(
	ctx context.Context,
	request *models.CreateGlobalSystemConfigRequest) (resp *models.CreateGlobalSystemConfigResponse, err error) {
	return nil, fmt.Errorf("Trying to call create GlobalSystemConfig outside of DBInit")
}

func (service *ContrailTypeLogicService) UpdateGlobalSystemConfig(
	ctx context.Context,
	request *models.UpdateGlobalSystemConfigRequest) (resp *models.UpdateGlobalSystemConfigResponse, err error) {

	oldObj, err := service.GetGlobalSystemConfig(ctx, &models.GetGlobalSystemConfigRequest{})
	if err != nil || oldObj == nil {
		return nil, fmt.Errorf("No GlobalSystemConfig found to update")
	}

	err = service.checkUdc(request)
	if err != nil {
		return nil, err
	}

	err = service.checkAsn(ctx, request, oldObj.GlobalSystemConfig)
	if err != nil {
		return nil, err
	}

	err = service.checkBgpaasPorts(ctx, request, oldObj.GlobalSystemConfig)
	if err != nil {
		return nil, err
	}

	if service.Next() != nil {

		resp, err = service.Next().UpdateGlobalSystemConfig(ctx, request)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (service *ContrailTypeLogicService) checkUdc(request *models.UpdateGlobalSystemConfigRequest) (err error) {
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

// TODO DELETE AFTER ROUTE_TARGET.GO CREATED, ADD TESTS
func (service *ContrailTypeLogicService) isRouteTargetUserDefined(routeTarget string, asn int64) (bool, error) {
	return true, nil
}

func (service *ContrailTypeLogicService) checkAsn(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest, obj *models.GlobalSystemConfig) (err error) {

	globalAsn := request.GlobalSystemConfig.AutonomousSystem
	if globalAsn == 0 {
		return nil
	}

	vnList, err := service.DB.ListVirtualNetwork(ctx, &models.ListVirtualNetworkRequest{&models.ListSpec{Fields: []string{"route_target_list"}}})
	if err != nil {
		return err
	}

	for _, vn := range vnList.VirtualNetworks {
		rtList := vn.RouteTargetList
		for _, rt := range rtList.RouteTarget {
			userDefed, err := service.isRouteTargetUserDefined(rt, globalAsn) // Temporary implementation
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

func (service *ContrailTypeLogicService) checkBgpaasPorts(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest, obj *models.GlobalSystemConfig) (err error) {
	newObj := request.GlobalSystemConfig
	if newObj.GetBgpaasParameters() == nil {
		return nil
	}

	err = newObj.BgpaasParameters.ValidatePortRange()
	if err != nil {
		return err
	}

	newPortsRange := newObj.GetBgpaasParameters()
	oldPortsRange := models.GetDefaultBGPaaServiceParameters()

	if obj.GetBgpaasParameters() != nil {
		oldPortsRange = obj.GetBgpaasParameters()
	}

	numBgpaas, err := service.DB.ListBGPAsAService(ctx, &models.ListBGPAsAServiceRequest{&models.ListSpec{Count: true}})

	if err != nil {
		return err
	}

	if len(numBgpaas.BGPAsAServices) > 0 {

		if !newPortsRange.ContainsRange(oldPortsRange) {
			return fmt.Errorf("BGP Port range cannot be shrunk")
		}
	}
	return nil
}
