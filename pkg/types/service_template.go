package types

import (
	"context"

	"github.com/Juniper/asf/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateServiceTemplate validates domain UUID and enables domain sharing for service template.
func (sv *ContrailTypeLogicService) CreateServiceTemplate(
	ctx context.Context,
	request *services.CreateServiceTemplateRequest,
) (response *services.CreateServiceTemplateResponse, err error) {
	svcTemplate := request.GetServiceTemplate()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			domainUUID := svcTemplate.GetParentUUID()
			if domainUUID == "" {
				domainUUID, err = sv.FQNameToUUID(ctx, svcTemplate.GetFQName(), svcTemplate.GetParentType())
				if err != nil {
					return err
				}
			}
			err = svcTemplate.GetPerms2().EnableDomainSharing(domainUUID, models.PermsRX)
			if err != nil {
				return err
			}
			response, err = sv.BaseService.CreateServiceTemplate(ctx, request)
			return err
		})

	return response, err
}
