package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

func (s *SVCInstance) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncSI, err := s.toVncForCreate(ctx, rp)
	if err != nil {
		return nil, err
	}

	if vncSI, err = rp.WriteService.CreateServiceInstance(ctx, &services.CreateServiceInstanceRequest{
		ServiceInstance: vncSI,
	}); err != nil {
		return err
	}

}

func (s *SVCInstance) toVncForCreate(ctx context.Context, rp RequestParameters) (*models.ServiceInstance, error) {
	vncSI := models.MakeServiceInstance()

	vncSI.ParentType = models.KindProject

	var err error
	vncSI.ParentUUID, err = neutronIDToVncUUID(s.TenantID)
	if err != nil {
		return nil, err
	}

	vnResp, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: s.NetID, Fields: []string{models.VirtualNetworkFieldFQName},
	})
	if err != nil {
		return nil, err
	}

	vncSI.ServiceInstanceProperties = &models.ServiceInstanceType{
		AutoPolicy:          true,
		RightVirtualNetwork: basemodels.FQNameToString(vnResp.GetVirtualNetwork().GetFQName()),
		ScaleOut: &models.ServiceScaleOutType{
			MaxInstances: 1,
			AutoScale:    false,
		},
	}

	return vncSI, nil
}

func makeServiceInstanceResponse(ctx context.Context, vncSI *models.ServiceInstance, rs services.ReadService) (*SVCInstanceResponse, error) {
	// TODO(Michal): find a way to implement `si_q_dict = self._obj_to_dict(si_obj)`

	s := &SVCInstanceResponse{
		ID:       vncSI.GetUUID(),
		TenantID: VncUUIDToNeutronID(vncSI.GetParentUUID()),
		Name:     vncSI.GetName(),
	}
	if vncSI.GetServiceInstanceProperties() != nil {
		vnFQName := basemodels.ParseFQName(vncSI.GetRightVirtualNetwork())
		vnResp, err := rs.ListVirtualNetwork(ctx, &services.ListVirtualNetworkRequest{})
	}

	return s
}

func (s *SVCInstance) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	panic("not implemented")
}

func (s *SVCInstance) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	_, err := rp.WriteService.DeleteServiceInstance(ctx, &services.DeleteServiceInstanceRequest{
		ID: id,
	})

	return &SVCInstanceResponse{}, err
}

func (s *SVCInstance) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	panic("not implemented")
}

func (s *SVCInstance) ReadAll(ctx context.Context, rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	panic("not implemented")
}
