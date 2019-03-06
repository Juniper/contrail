package services

import (
	"context"
	"fmt"
	"strings"

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
	refs basemodels.References,
) error {
	refsWithoutUUID := refs.Filter(func(r basemodels.Reference) bool { return r.GetUUID() == "" })
	if len(refsWithoutUUID) == 0 {
		return nil
	}
	foundMetadata, err := sv.MetadataGetter.ListMetadata(ctx, refsToMetadatas(refsWithoutUUID))
	if err != nil {
		return err
	}
	if len(foundMetadata) != len(refsWithoutUUID) {
		if notFound := getRefsNotFound(refsWithoutUUID, foundMetadata); len(notFound) != 0 {
			return errors.Errorf("couldn't get metadata for references:%v", listNotFoundEvents(notFound))
		}
	}
	fillUUIDs(refsWithoutUUID, foundMetadata)
	return nil
}

func listNotFoundEvents(notFound basemodels.References) string {
	var results []string
	for _, ref := range notFound {
		results = append(results, fmt.Sprintf("{type: %v, to: %v}", ref.GetToKind(), ref.GetTo()))
	}
	return strings.Join(results, " ")
}

func fillUUIDs(refs basemodels.References, foundMetadata []*basemodels.Metadata) {
	fqNameToUUID := make(map[string]string)
	for _, metadata := range foundMetadata {
		fqNameToUUID[basemodels.FQNameToString(metadata.FQName)] = metadata.UUID
	}
	for _, ref := range refs {
		ref.SetUUID(fqNameToUUID[basemodels.FQNameToString(ref.GetTo())])
	}
}

func getRefsNotFound(
	refs basemodels.References, mds []*basemodels.Metadata,
) basemodels.References {
	found := metadatasToFQNames(mds)
	return refs.Filter(func(r basemodels.Reference) bool {
		return !found[basemodels.FQNameToString(r.GetTo())]
	})
}

func metadatasToFQNames(mds []*basemodels.Metadata) map[string]bool {
	fqNames := make(map[string]bool, len(mds))
	for _, m := range mds {
		fqNames[basemodels.FQNameToString(m.FQName)] = true
	}
	return fqNames
}

func refsToMetadatas(refs basemodels.References) []*basemodels.Metadata {
	var metadatas []*basemodels.Metadata
	for _, ref := range refs {
		metadatas = append(metadatas, &basemodels.Metadata{
			FQName: ref.GetTo(),
			Type:   ref.GetToKind(),
		})
	}
	return metadatas
}
