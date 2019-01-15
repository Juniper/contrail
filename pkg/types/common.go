package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// EnableDomainSharing enables domain sharing for resource.
func (sv *ContrailTypeLogicService) enableDomainSharing(
	ctx context.Context,
	perms2 *models.PermType2,
	domainUUID string,
	accessLevel int64,
) (*models.PermType2, error) {

	if perms2 == nil {
		perms2 = &models.PermType2{}
	}

	perms2.Share = append(perms2.GetShare(), &models.ShareType{
		Tenant:       "domain:" + domainUUID,
		TenantAccess: accessLevel,
	})

	return perms2, nil
}

func (sv *ContrailTypeLogicService) getParentUUIDFromFQName(
	ctx context.Context,
	obj basemodels.Object,
) (string, error) {
	fqRequest := services.FQNameToIDRequest{
		FQName: obj.GetFQName(),
		Type:   obj.GetParentType(),
	}

	metadata, err := sv.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: fqRequest.Type, FQName: fqRequest.FQName})
	if err != nil {
		return "", err
	}

	return metadata.UUID, nil
}
