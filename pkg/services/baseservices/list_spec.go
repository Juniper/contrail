package baseservices

import (
	"context"
)

// SimpleListSpec is a list spec for a common case where List is filtered only by uuids.
// Whole resources are fetched if no fields are specified
func SimpleListSpec(ctx context.Context, uuids []string, fields ...string) *ListSpec {
	return &ListSpec{
		Filters: []*Filter{{
			Key:    "uuid",
			Values: uuids,
		}},
		Detail: true,
		Fields: fields,
	}
}
