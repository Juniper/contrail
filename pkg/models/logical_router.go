package models

import (
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
)

const (
	noneString       = "None"
	internalVNPrefix = "__contrail_lr_internal_vn_"
)

// GetVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() (string, error) {
	id := lr.GetVxlanNetworkIdentifier()
	if id == noneString || id == "" {
		return "", nil
	}

	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", common.ErrorBadRequestf("vxlan network id must be a number(%s)", id)
	}

	return id, nil
}

func (lr *LogicalRouter) GetInternalVNName() string {
	name := []string{internalVNPrefix, lr.GetUUID(), "__"}
	return strings.Join(name, "")
}
