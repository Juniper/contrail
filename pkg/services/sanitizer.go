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
	refsWithoutUUID := refs.Filter(func(r basemodels.Reference) bool{return  r.GetUUID() == ""})
	if len(refsWithoutUUID) == 0 {
		return nil
	}
	foundMetadata, err := sv.MetadataGetter.ListMetadata(ctx, refsToMetadatas(refsWithoutUUID))
	if err != nil {
		return err
	}
	if len(foundMetadata) != len(refsWithoutUUID) {
		if notFound := getRefsNotFound(refsWithoutUUID, foundMetadata); len(notFound) != 0 {
			return errors.Errorf("couldn't get metadatas for references:%v", listNotFoundEvents(notFound))
		}
	}
	fillUUIDs(refsWithoutUUID, foundMetadata)
	return nil
}

func listNotFoundEvents(notFound basemodels.References) string {
	var results []string
	for _, ref := range notFound {
		results = append(results, fmt.Sprintf("{type: %v, to: %v}", ref.GetReferredKind(), ref.GetTo()))
	}
	return strings.Join(results, " ")
}

func fillUUIDs(refs basemodels.References, foundMetadata []*basemodels.Metadata) {
	fqNameToUUID := make(map[string]string)
	for _, metadata := range foundMetadata {
		fqNameToUUID[basemodels.FQNameToString(metadata.FQName)] = metadata.UUID
	}
	for id := range refs {
		refs[id].SetUUID(fqNameToUUID[basemodels.FQNameToString(refs[id].GetTo())])
	}
}

func getRefsNotFound(refs basemodels.References, foundMetadata []*basemodels.Metadata) basemodels.References {
	var notFound basemodels.References
	found := metadatasToFQNames(foundMetadata)
	for id, ref := range refs {
		if !found[basemodels.FQNameToString(ref.GetTo())]{
			notFound = append(notFound, refs[id])
		}
	}
	return notFound
}

func metadatasToFQNames(metadatas []*basemodels.Metadata) map[string]bool {
	fqNames := make(map[string]bool, len(metadatas))
	for _, m := range metadatas {
		fqNames[basemodels.FQNameToString(m.FQName)] = true
	}
	return fqNames
}

func refsToMetadatas(refs basemodels.References) []*basemodels.Metadata {
	var metadatas []*basemodels.Metadata
	for _, ref := range refs {
		metadatas = append(metadatas, &basemodels.Metadata{FQName: ref.GetTo(), Type: ref.GetReferredKind()})
	}
	return metadatas
}
