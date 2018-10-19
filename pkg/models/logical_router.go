package models

import (
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
)

const (
	noneString       = "None"
	internalVNPrefix = "__contrail_lr_internal_vn_"
)

// GetVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() string {
	id := lr.GetVxlanNetworkIdentifier()
	if id == noneString || id == "" {
		return ""
	}

	return id
}

// GetInternalVNName returns proper internal virtual network name
func (lr *LogicalRouter) GetInternalVNName() string {
	return internalVNPrefix + lr.GetUUID() + "__"
}

// ConvertVXLanIDToInt converts vxlan network id form string to int
func (lr *LogicalRouter) ConvertVXLanIDToInt() (int64, error) {
	vxlanNetworkID := lr.GetVxlanNetworkIdentifier()
	id, err := strconv.ParseInt(vxlanNetworkID, 10, 64)
	if err != nil {
		return 0, errutil.ErrorBadRequestf("vxlan network id must be a number(%s): %v", vxlanNetworkID, err)
	}

	return id, nil
}
