package types

import (
	"fmt"
	"math"

	"regexp"

	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

// CreateGlobalSystemConfig
func (service *ContrailTypeLogicService) CreateGlobalSystemConfig(
	ctx context.Context,
	request *models.CreateGlobalSystemConfigRequest) (resp *models.CreateGlobalSystemConfigResponse, err error) {

	oldObj, err := service.GetGlobalSystemConfig(ctx, &models.GetGlobalSystemConfigRequest{})
	if err != nil || oldObj != nil {
		return nil, fmt.Errorf("GlobalSystemConfig exists")
	}

	if service.Next() == nil {
		return nil, nil
	}

	resp, err = service.Next().CreateGlobalSystemConfig(ctx, request)

	return
}

// UpdateGlobalSystemConfig
func (service *ContrailTypeLogicService) UpdateGlobalSystemConfig(
	ctx context.Context,
	request *models.UpdateGlobalSystemConfigRequest) (resp *models.UpdateGlobalSystemConfigResponse, err error) {

	oldObj, err := service.GetGlobalSystemConfig(ctx, &models.GetGlobalSystemConfigRequest{})
	if err != nil || oldObj == nil {
		return nil, fmt.Errorf("No GlobalSystemConfig found to update")
	}

	err = service.checkUdc(request)
	if err != nil {
		return
	}

	err = service.checkAsn(ctx, request, oldObj.GlobalSystemConfig)
	if err != nil {
		return
	}

	err = service.checkBgpaasPorts(ctx, request, oldObj.GlobalSystemConfig)
	if err != nil {
		return
	}

	if service.Next() != nil {

		resp, err = service.Next().UpdateGlobalSystemConfig(ctx, request)
		if err != nil {
			return nil, err
		}
	}

	return
}

func (service *ContrailTypeLogicService) checkUdc(request *models.UpdateGlobalSystemConfigRequest) (err error) {
	// How to work with this? We need testing by fire
	//request.FieldMask.GetPaths() should filter these

	if request.GlobalSystemConfig.UserDefinedLogStatistics == nil {
		return
	}
	for _, udc := range request.GlobalSystemConfig.UserDefinedLogStatistics.Statlist {
		_, err = regexp.Compile(udc.Pattern)
		if err != nil {
			return
		}
	}

	return
}

// TODO DELETE AFTER ROUTE_TARGET.GO CREATED, ADD TESTS
func (service *ContrailTypeLogicService) isRouteTargetUserDefined(routeTarget string, asn int64) (bool, error) {
	return true, nil
}

// global asn in python implementation = current globalConfig.asn
func (service *ContrailTypeLogicService) checkAsn(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest, obj *models.GlobalSystemConfig) (err error) {

	globalAsn := request.GlobalSystemConfig.AutonomousSystem
	if globalAsn == 0 {
		return nil
	}

	vnList, err := service.DB.ListVirtualNetwork(ctx, &models.ListVirtualNetworkRequest{&models.ListSpec{Fields: []string{"route_target_list"}}})
	if err != nil {
		return
	}

	foundedUsingAsn := false

	for _, vn := range vnList.VirtualNetworks {
		rtList := vn.RouteTargetList
		for _, rt := range rtList.RouteTarget {
			userDefed, err := service.isRouteTargetUserDefined(rt, globalAsn) // Temporary implementation
			if err != nil {
				return err
			}
			if userDefed {
				foundedUsingAsn = true
			}
		}
	}

	if !foundedUsingAsn {
		return
	}

	return fmt.Errorf("Cannot update ASN as there are ASN based VN's") // Can add more specific output using route target details
}

func (service *ContrailTypeLogicService) checkBgpaasPorts(ctx context.Context, request *models.UpdateGlobalSystemConfigRequest, obj *models.GlobalSystemConfig) (err error) {
	newObj := request.GlobalSystemConfig
	if newObj.GetBgpaasParameters() == nil {
		return nil
	}

	newPortStart := newObj.GetBgpaasParameters().PortStart
	newPortEnd := newObj.GetBgpaasParameters().PortEnd

	err = checkValidPortRange(newPortStart, newPortEnd)
	if err != nil {
		return err
	}

	var oldPortStart int64 = 50000
	var oldPortEnd int64 = 50512

	if obj.GetBgpaasParameters() != nil {
		oldPortStart = obj.GetBgpaasParameters().PortStart
		oldPortEnd = obj.GetBgpaasParameters().PortEnd
	}

	numBgpaas, err := service.DB.ListBGPAsAService(ctx, &models.ListBGPAsAServiceRequest{&models.ListSpec{Count: true}})

	if err != nil {
		return
	}

	if len(numBgpaas.BGPAsAServices) > 0 {

		if (newPortStart > oldPortStart) || (newPortEnd < oldPortEnd) {
			return fmt.Errorf("BGP Port range cannot be shrunk")
		}
	}
	return
}

func checkValidPortRange(portStart, portEnd int64) error {
	if (portStart > portEnd) || (portStart <= 0) || (portEnd > math.MaxUint16) {
		return fmt.Errorf("Invalid Port range specified")
	}
	return nil
}
