package neutron

import (
	"context"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

type keystoneClient interface {
	GetProject(ctx context.Context, token string, id string) (*keystone.Project, error)
}

// Service handles neutron specific logic
type Service struct {
	services.BaseService
	Keystone          keystoneClient
	ReadService       services.ReadService
	MetadataGetter    baseservices.MetadataGetter
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
				p, kerr := sv.getProjectFromKeystone(ctx, request.GetID())
				if kerr != nil {
					return errutil.MultiError{err, kerr}
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

func (sv *Service) getProjectFromKeystone(ctx context.Context, id string) (*keystone.Project, error) {
	token := auth.GetIdentity(ctx).AuthToken()
	if token == "" {
		return nil, errors.New("expected auth token in context")
	}
	p, err := sv.Keystone.GetProject(ctx, token, logic.VncUUIDToNeutronID(id))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get project from keystone")
	}

	return p, nil
}

// CreateProject creates the project and ensures its default security group exists.
func (sv *Service) CreateProject(
	ctx context.Context, request *services.CreateProjectRequest,
) (*services.CreateProjectResponse, error) {
	var response *services.CreateProjectResponse
	err := sv.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var err error

		response, err = sv.BaseService.CreateProject(ctx, request)
		if err != nil {
			return err
		}

		project := response.GetProject()

		_, err = sv.WriteService.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
			SecurityGroup: project.DefaultSecurityGroup(),
		})
		return err

		// TODO: Create default firewall group.
	})
	return response, err
}

// DeleteProject deletes the project with its default security group.
func (sv *Service) DeleteProject(
	ctx context.Context, request *services.DeleteProjectRequest,
) (*services.DeleteProjectResponse, error) {
	var response *services.DeleteProjectResponse

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			if err = sv.deleteDefaultSecurityGroup(ctx, request.GetID()); err != nil {
				return err
			}

			// TODO: Delete default firewall group.

			response, err = sv.BaseService.DeleteProject(ctx, request)
			return err
		})

	return response, err
}

func (sv *Service) deleteDefaultSecurityGroup(ctx context.Context, projectUUID string) error {
	projectResponse, err := sv.ReadService.GetProject(ctx, &services.GetProjectRequest{
		ID:     projectUUID,
		Fields: []string{models.ProjectFieldFQName},
	})
	if err != nil {
		return err
	}

	metadata, err := sv.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{
		FQName: projectResponse.GetProject().DefaultSecurityGroupFQName(),
		Type:   models.KindSecurityGroup,
	})
	if err != nil {
		return errors.Wrap(err, "default SecurityGroup not found")
	}

	_, err = sv.WriteService.DeleteSecurityGroup(ctx, &services.DeleteSecurityGroupRequest{
		ID: metadata.UUID,
	})
	return err
}
