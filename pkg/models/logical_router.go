package models

import (
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
)

const (
<<<<<<< f18964bdfa7131de7c589afcbbd0b15d03e208ac
<<<<<<< 8e8d69379981e06d9ed655e8e3edb3397cb7f007
	noneString = "None"
=======
	NoneString = "None"
>>>>>>> Implementing pre create/update logical-router type validation
=======
	noneString       = "None"
	internalVNPrefix = "__contrail_lr_internal_vn_"
>>>>>>> Adding post create/update logical router validation
)

// GetVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() (string, error) {
	id := lr.GetVxlanNetworkIdentifier()
<<<<<<< f18964bdfa7131de7c589afcbbd0b15d03e208ac
<<<<<<< 8e8d69379981e06d9ed655e8e3edb3397cb7f007
	if id == noneString {
=======
	if id == NoneString {
>>>>>>> Implementing pre create/update logical-router type validation
		return ""
=======
	if id == noneString || id == "" {
		return "", nil
>>>>>>> Adding post create/update logical router validation
	}

	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return id, common.ErrorBadRequestf("vxlan network id must be a number(%s)", id)
	}

	return id, nil
}

//GetInternalVNName returns proper internal virtual network name
func (lr *LogicalRouter) GetInternalVNName() string {
	name := []string{internalVNPrefix, lr.GetUUID(), "__"}
	return strings.Join(name, "")
}
