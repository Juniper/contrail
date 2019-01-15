package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

//CreateServiceTemplate does pre check for service template.
func (sv *ContrailTypeLogicService) CreateServiceTemplate(
	ctx context.Context,
	request *services.CreateServiceTemplateRequest,
) (response *services.CreateServiceTemplateResponse, err error) {

	serviceTemplate := request.GetServiceTemplate()

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			domainUUID := serviceTemplate.GetParentUUID()
			if domainUUID == "" {
				domainUUID, err = sv.getParentDomainUUIDFromFQName(ctx, serviceTemplate)
				if err != nil {
					return err
				}
			}

			perms2 := serviceTemplate.GetPerms2()
			if perms2 == nil {
				serviceTemplate.Perms2 = &models.PermType2{}
				perms2 = serviceTemplate.GetPerms2()
			}

			perms2.Share = append(perms2.GetShare(), &models.ShareType{
				Tenant:       "domain:" + domainUUID,
				TenantAccess: 5, // RX_PERMS
			})

			response, err = sv.BaseService.CreateServiceTemplate(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getParentDomainUUIDFromFQName(
	ctx context.Context,
	st *models.ServiceTemplate,
) (string, error) {
	fqRequest := services.FQNameToIDRequest{
		FQName: st.GetFQName(),
		Type:   st.GetParentType(),
	}

	metadata, err := sv.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: fqRequest.Type, FQName: fqRequest.FQName})
	if err != nil {
		return "", err
	}

	return metadata.UUID, nil
}
