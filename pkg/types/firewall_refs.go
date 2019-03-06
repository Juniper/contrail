package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// complementRefs checks if to fields in resource refs are filled
func (sv *ContrailTypeLogicService) complementRefs(ctx context.Context, obj basemodels.Object) error {
	for _, ref := range obj.GetReferences() {
		if err := sv.fillToFieldInRef(ctx, ref); err != nil {
			return err
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) fillToFieldInRef(ctx context.Context, ref basemodels.Reference) error {
	if len(ref.GetTo()) != 0 {
		return nil
	}

	m, err := sv.MetadataGetter.GetMetadata(
		ctx,
		basemodels.Metadata{
			UUID: ref.GetUUID(),
			Type: ref.GetToKind(),
		},
	)
	if err != nil {
		return errutil.ErrorNotFoundf("cannot find %s ref with uuid %s: %v", ref.GetToKind(), ref.GetUUID(), err)
	}

	ref.SetTo(m.FQName)
	return nil
}
