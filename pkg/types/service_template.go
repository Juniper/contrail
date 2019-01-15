package types

import (
	"context"

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
				domainUUID, err = sv.getParentUUIDFromFQName(ctx, serviceTemplate)
				if err != nil {
					return err
				}
			}

			serviceTemplate.Perms2, err = sv.enableDomainSharing(ctx, serviceTemplate.GetPerms2(), domainUUID, basemodels.PermsRX)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.CreateServiceTemplate(ctx, request)
			return err
		})

	return response, err
}
