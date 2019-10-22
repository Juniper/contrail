package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/asf/pkg/models"
	"github.com/pkg/errors"
)

// Sanitizer is an object that allows setting missing values on objects based on theirs metadata.
type Sanitizer struct {
	MetadataGetter MetadataGetter
}

// NamedObject is object that has methods for getting and setting name, display_name and fq_name.
type NamedObject interface {
	GetName() string
	GetDisplayName() string
	GetFQName() []string
}

// SanitizeDisplayNameAndName sets name and display_name based on fq_name.
func (sv *Sanitizer) SanitizeDisplayNameAndName(ctx context.Context, obj NamedObject) (name, displayName string) {
	name, displayName = obj.GetName(), obj.GetDisplayName()
	if name == "" {
		fqName := obj.GetFQName()
		if len(fqName) > 0 {
			name = fqName[len(fqName)-1]
		}
	}

	if displayName == "" {
		displayName = name
	}

	return name, displayName
}

// SanitizeRefs fills all object's refs with uuids if empty.
func (sv *Sanitizer) SanitizeRefs(ctx context.Context, refs models.References) error {
	refsWithoutUUID := refs.Filter(func(r models.Reference) bool { return r.GetUUID() == "" })
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

func refsToMetadatas(refs models.References) []*models.Metadata {
	var metadatas []*models.Metadata
	for _, ref := range refs {
		metadatas = append(metadatas, &models.Metadata{
			FQName: ref.GetTo(),
			Type:   ref.GetToKind(),
		})
	}
	return metadatas
}

func getRefsNotFound(
	refs models.References, mds []*models.Metadata,
) models.References {
	found := metadatasToFQNames(mds)
	return refs.Filter(func(r models.Reference) bool {
		return !found[models.FQNameToString(r.GetTo())]
	})
}

func metadatasToFQNames(mds []*models.Metadata) map[string]bool {
	fqNames := make(map[string]bool, len(mds))
	for _, m := range mds {
		fqNames[models.FQNameToString(m.FQName)] = true
	}
	return fqNames
}

func listNotFoundEvents(notFound models.References) string {
	var results []string
	for _, ref := range notFound {
		results = append(results, fmt.Sprintf("{type: %v, to: %v}", ref.GetToKind(), ref.GetTo()))
	}
	return strings.Join(results, " ")
}

func fillUUIDs(refs models.References, foundMetadata []*models.Metadata) {
	fqNameToUUID := make(map[string]string)
	for _, metadata := range foundMetadata {
		fqNameToUUID[models.FQNameToString(metadata.FQName)] = metadata.UUID
	}
	for _, ref := range refs {
		ref.SetUUID(fqNameToUUID[models.FQNameToString(ref.GetTo())])
	}
}
