package models

import (
	"strings"
)

var reservedQFXL2VlanTags = []int64{1, 2, 4094}

func (m *LogicalInterface) IsReservedQFXVlanTag() bool {
	if strings.ToLower(m.LogicalInterfaceType) != "l2" {
		return false
	}
	for _, reservedTag := range reservedQFXL2VlanTags {
		if reservedTag == m.LogicalInterfaceVlanTag {
			return true
		}
	}
	return false
}
