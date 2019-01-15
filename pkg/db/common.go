package db

import (
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/format"
)

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
