package services

import (
	"bytes"
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// SanitizerService fills up missing properties based on resources logic and metadata
// TODO: Move logic from ContrailService when validation will be a separate service
type SanitizerService struct {
	BaseService
	MetadataGetter baseservices.MetadataGetter
}

func (sv *SanitizerService) sanitizeRefs(
	ctx context.Context,
	refs []basemodels.Reference,
) error {
	fqNameToRef := make(map[string]basemodels.Reference)
	var metadatas []*basemodels.Metadata
	for _, ref := range refs {
		if ref.GetUUID() != "" {
			continue
		}
		fqNameToRef[basemodels.FQNameToString(ref.GetTo())] = ref
		metadatas = append(metadatas, &basemodels.Metadata{FQName: ref.GetTo(), Type: ref.GetReferredKind()})
	}

	if len(metadatas) == 0 {
		return nil
	}

	ml, err := sv.MetadataGetter.ListMetadata(ctx, metadatas)
	if err != nil {
		return err
	}

	if len(ml) != len(metadatas) {
		for _, metadata := range ml {
			delete(fqNameToRef, basemodels.FQNameToString(metadata.FQName))
		}
		var missingRefs bytes.Buffer
		for _, ref := range fqNameToRef {
			missingRefs.WriteString(fmt.Sprintf(" {type: %v, to: %v}", ref.GetReferredKind(), ref.GetTo()))
		}
		return errors.Errorf("couldn't get metadatas for references:%v", missingRefs.String())
	}

	for _, metadata := range ml {
		fqNameToRef[basemodels.FQNameToString(metadata.FQName)].SetUUID(metadata.UUID)
	}
	return nil
}
