package models

<<<<<<< 8e8d69379981e06d9ed655e8e3edb3397cb7f007
const (
	noneString = "None"
)

// GetVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() string {
	id := lr.GetVxlanNetworkIdentifier()
	if id == noneString {
		return ""
	}

	return id
=======
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
		return id, common.ErrorBadRequestf("vxlan network id must be a number(%s)", id)
	}

	return id, nil
}

// GetInternalVNName returns proper internal virtual network name
func (lr *LogicalRouter) GetInternalVNName() string {
	name := []string{internalVNPrefix, lr.GetUUID(), "__"}
	return strings.Join(name, "")
>>>>>>> [WIP]Adding post create/update logical router validation
}
