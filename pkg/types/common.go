package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// EnableDomainSharing enables domain sharing for resource.
func enableDomainSharing(
	perms2 *models.PermType2,
	domainUUID string,
	accessLevel int64,
) (*models.PermType2, error) {

	if perms2 == nil {
		perms2 = &models.PermType2{}
	}

	perms2.Share = append(perms2.Share, &models.ShareType{
		Tenant:       "domain:" + domainUUID,
		TenantAccess: accessLevel,
	})

	return perms2, nil
}

func (sv *ContrailTypeLogicService) getUUIDFromFQName(
	ctx context.Context,
	fqName []string,
	resourceType string,
) (string, error) {

	metadata, err := sv.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: resourceType, FQName: fqName})
	if err != nil {
		return "", err
	}

	return metadata.UUID, nil
}
