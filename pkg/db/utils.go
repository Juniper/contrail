package db

import "github.com/Juniper/contrail/pkg/services/baseservices"

func listSpecForGet(uuid string, fields []string) *baseservices.ListSpec {
	return &baseservices.ListSpec{
		Limit:  1,
		Detail: true,
		Fields: fields,
		Filters: []*baseservices.Filter{
			&baseservices.Filter{
				Key:    "uuid",
				Values: []string{uuid},
			},
		},
	}
}
