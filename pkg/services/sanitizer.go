package services

import (
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
	refsWithoutUUID := getRefsWithoutUUID(refs)
	if len(refsWithoutUUID) == 0 {
		return nil
	}
	foundMetadata, err := sv.MetadataGetter.ListMetadata(ctx, refsIntoMetadatas(refsWithoutUUID))
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

func listNotFoundEvents(notFound []basemodels.Reference) string {
	var result string
	for _, ref := range notFound {
		result += fmt.Sprintf(" {type: %v, to: %v}", ref.GetReferredKind(), ref.GetTo())
	}
	return result
}

func fillUUIDs(refs []basemodels.Reference, foundMetadata []*basemodels.Metadata) {
	fqNameToUUID := make(map[string]string)
	for _, metadata := range foundMetadata {
		fqNameToUUID[basemodels.FQNameToString(metadata.FQName)] = metadata.UUID
	}
	for id := range refs {
		refs[id].SetUUID(fqNameToUUID[basemodels.FQNameToString(refs[id].GetTo())])
	}
}

func getRefsNotFound(refs []basemodels.Reference, foundMetadata []*basemodels.Metadata) []basemodels.Reference {
	var notFound []basemodels.Reference
	for id, ref := range refs {
		found := false
		for _, metadata := range foundMetadata {
			if basemodels.FQNameEquals(metadata.FQName, ref.GetTo()) {
				found = true
				break
			}
		}
		if !found {
			notFound = append(notFound, refs[id])
		}
	}
	return notFound
}

func refsIntoMetadatas(refs []basemodels.Reference) []*basemodels.Metadata {
	var metadatas []*basemodels.Metadata
	for _, ref := range refs {
		metadatas = append(metadatas, &basemodels.Metadata{FQName: ref.GetTo(), Type: ref.GetReferredKind()})
	}
	return metadatas
}


func getRefsWithoutUUID(refs []basemodels.Reference) []basemodels.Reference {
	var withoutUUID []basemodels.Reference
	for id, ref := range refs {
		if ref.GetUUID() == "" {
			withoutUUID = append(withoutUUID, refs[id])
		}
	}
	return withoutUUID
}