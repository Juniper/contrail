package db

import (
	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/services/baseservices"
)

func listSpecForGet(uuid string, fields []string) *baseservices.ListSpec {
	return &baseservices.ListSpec{
		Limit:  1,
		Detail: true,
		Shared: true,
		Fields: fields,
		Filters: []*baseservices.Filter{
			{
				Key:    "uuid",
				Values: []string{uuid},
			},
		},
	}
}

func resolveUUIDAndFQNameFromMap(m map[string]interface{}) (uuid string, fqName []string, err error) {
	uuid = format.InterfaceToString(m["to"])
	if uuid == "" {
		return "", nil, nil
	}
	fqNameStr := format.InterfaceToString(m["fq_name"])
	fqName, err = basedb.ParseFQName(fqNameStr)
	if err != nil {
		return "", nil, err
	}
	return uuid, fqName, nil
}
