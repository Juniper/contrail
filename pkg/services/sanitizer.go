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

func (sv *SanitizerService) getIncompleteRefs(
	refs []basemodels.Reference,
) map[string]basemodels.Reference {
	incompleteRefsMap := make(map[string]basemodels.Reference)

	for _, ref := range refs {
		switch {
		case ref.GetUUID() == "":
			incompleteRefsMap[basemodels.FQNameToString(ref.GetTo())] = ref
		case len(ref.GetTo()) == 0:
			incompleteRefsMap[ref.GetUUID()] = ref
		default:
			continue
		}
	}

	return incompleteRefsMap
}

func (sv *SanitizerService) completeRefs(
	incompleteRefsMap map[string]basemodels.Reference, metadatas []*basemodels.Metadata,
) {
	for _, metadata := range metadatas {
		if ref, ok := incompleteRefsMap[basemodels.FQNameToString(metadata.FQName)]; ok {
			ref.SetUUID(metadata.UUID)
		}

		if ref, ok := incompleteRefsMap[metadata.UUID]; ok {
			ref.SetTo(metadata.FQName)
		}
	}
}

func (sv *SanitizerService) sanitizeRefs(
	ctx context.Context,
	refs []basemodels.Reference,
) error {
	incompleteRefsMap := sv.getIncompleteRefs(refs)
	var metadatas []*basemodels.Metadata

	if len(incompleteRefsMap) == 0 {
		return nil
	}

	for _, ref := range incompleteRefsMap {
		metadatas = append(metadatas,
			&basemodels.Metadata{FQName: ref.GetTo(), UUID: ref.GetUUID(), Type: ref.GetReferredKind()})
	}

	ml, err := sv.MetadataGetter.ListMetadata(ctx, metadatas)
	if err != nil {
		return err
	}

	if len(ml) != len(metadatas) {
		for _, metadata := range ml {
			delete(incompleteRefsMap, basemodels.FQNameToString(metadata.FQName))
			delete(incompleteRefsMap, metadata.UUID)
		}

		var missingRefs bytes.Buffer
		for _, ref := range incompleteRefsMap {
			missingRefs.WriteString(fmt.Sprintf(" {type: %v, to: %v}", ref.GetReferredKind(), ref.GetTo()))
		}

		return errors.Errorf("couldn't get metadatas for references:%v", missingRefs.String())
	}

	sv.completeRefs(incompleteRefsMap, ml)

	return nil
}
