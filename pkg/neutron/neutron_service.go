package neutron

import (
	"context"

	"github.com/Juniper/contrail/pkg/neutron/logic"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// NeutronService handles neutron specific logic
type NeutronService struct {
	services.BaseService
	keystone          *client.Keystone
	WriteService      services.WriteService
	InTransactionDoer services.InTransactionDoer
}

func NewNeutronService(
	keystoneUrl string, ws services.WriteService, itd services.InTransactionDoer,
) *NeutronService {
	sv := &NeutronService{
		// TODO use keystone client instead
		keystone:          client.NewKeystone(keystoneUrl, true, nil),
		WriteService:      ws,
		InTransactionDoer: itd,
	}
	return sv
}

// GetProject gets
func (sv *NeutronService) GetProject(
	ctx context.Context, request *services.GetProjectRequest,
) (*services.GetProjectResponse, error) {

	var response *services.GetProjectResponse
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			response, err = sv.BaseService.GetProject(ctx, request)
			if errutil.IsNotFound(err) {
				var tokenKey interface{} = "token"
				token := ctx.Value(tokenKey).(string)
				p, kerr := sv.keystone.GetProject(ctx, token, logic.ContrailUUIDToNeutronID(request.GetID()))
				if kerr != nil {
					return errors.Wrapf(err, "kerr %v", kerr)
				}

				_, err = sv.WriteService.CreateProject(
					ctx,
					&services.CreateProjectRequest{
						Project: &models.Project{
							UUID:        request.ID,
							DisplayName: p.Name,
							Name:        p.Name,
							ParentType:  models.KindDomain,
							FQName:      []string{"default-domain", p.Name},
						},
					},
				)
				if err != nil {
					return err
				}

				response, err = sv.BaseService.GetProject(ctx, request)
			}

			if err != nil {
				return err
			}

			return err
		})

	return response, err
}
