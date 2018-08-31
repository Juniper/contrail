package baseservices

import (
	"context"
)

func FieldsListSpec(ctx context.Context, uuids []string, fields ...string) *ListSpec {
	return &ListSpec{
		Filters: []*Filter{{
			Key:    "uuid",
			Values: uuids,
		}},
		Detail: true,
		Fields: fields,
	}
}
