package neutron

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/services"
)

type keystoneClient interface {
	GetProject(ctx context.Context, token string, id string) (*keystone.Project, error)
}

// Service handles neutron specific logic
type Service struct {
	services.BaseService
	Keystone          keystoneClient
	WriteService      services.WriteService
	InTransactionDoer services.InTransactionDoer
}

// GetProject ensures that projects in keystone exists in contrail and returns it
func (sv *Service) GetProject(
	ctx context.Context, request *services.GetProjectRequest,
) (*services.GetProjectResponse, error) {

	var response *services.GetProjectResponse
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			response, err = sv.BaseService.GetProject(ctx, request)
			if errutil.IsNotFound(err) {
				token := auth.GetAuthCTX(ctx).AuthToken()
				if token == "" {
					return errors.New("expected auth token in context")
				}
				p, kerr := sv.Keystone.GetProject(ctx, token, logic.VncUUIDToNeutronID(request.GetID()))
				if kerr != nil {
					var multiError errutil.MultiError
					multiError = append(multiError, err, errors.Wrap(kerr, "couldn't get project from keystone:"))
					return multiError
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

			return err
		})

	return response, err
}
