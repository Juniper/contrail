package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

/*
@classmethod
def pre_dbe_create(cls, tenant_name, obj_dict, db_conn):
   if (obj_dict.get('bgpaas_shared') == False) or (obj_dict.get('bgpaas_ip_address') != None):
   if (not obj_dict.get('bgpaas_shared') == True) or (obj_dict.get('bgpaas_ip_address') != None):
	  return True, ''
   return (False, (400, 'BGPaaS IP Address needs to be configured if BGPaaS is shared'))
# end pre_dbe_create
*/

// CreateBGPAsAService creates BGP as a service instance and ensures that shared one have IP address
func (sv *ContrailTypeLogicService) CreateBGPAsAService(
	ctx context.Context, request *services.CreateBGPAsAServiceRequest,
) (response *services.CreateBGPAsAServiceResponse, err error) {
	if err = sv.checkSharedAddress(request.GetBGPAsAService()); err != nil {
		return nil, errutil.ErrorBadRequest("BGPaaS IP Address needs to be configured if BGPaaS is shared")
	}
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			response, err = sv.BaseService.CreateBGPAsAService(ctx, request)
			return err
		})
	return response, err
}

/*
@classmethod
def pre_dbe_update(cls, id, fq_name, obj_dict, db_conn, **kwargs):
   if 'bgpaas_shared' in obj_dict:
	  ok, result = cls.dbe_read(db_conn, 'bgp_as_a_service', id)

	  if not ok:
		 return ok, result
	  if result.get('bgpaas_shared', False) != obj_dict['bgpaas_shared']:
		 return (False, (400, 'BGPaaS sharing cannot be modified'))
   return True, ""
# end pre_dbe_update
*/
func (sv *ContrailTypeLogicService) UpdateBGPAsAService(
	ctx context.Context, request *services.UpdateBGPAsAServiceRequest,
) (response *services.UpdateBGPAsAServiceResponse, err error) {
	id := request.GetBGPAsAService().GetUUID()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			fm := request.GetFieldMask()
			if basemodels.FieldMaskContains(&fm, models.BGPAsAServiceFieldBgpaasShared) {
				bgpaas, err := sv.GetBGPAsAService(ctx, &services.GetBGPAsAServiceRequest{ID: id})
				if err != nil {
					return err
				}
				if bgpaas.GetBGPAsAService().GetBgpaasShared() != request.GetBGPAsAService().GetBgpaasShared() {
					return errutil.ErrorBadRequest("BGPaaS sharing cannot be modified")
				}
			}
			response, err = sv.BaseService.UpdateBGPAsAService(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) checkSharedAddress(bgpaas *models.BGPAsAService) error {
	if bgpaas.BgpaasShared && len(bgpaas.GetBgpaasIPAddress()) == 0 {
		return errutil.ErrorBadRequest("BGPaaS IP Address needs to be configured if BGPaaS is shared")
	}
	return nil
}
