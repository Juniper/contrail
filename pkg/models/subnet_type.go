package models

import (
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
)

// ValidateWithEthertype checks if IP version from CIDR matches ethertype
// and throws an error if it doesn't.
func (s *SubnetType) ValidateWithEthertype(ethertype string) error {
	if s != nil {
		IPPrefix := s.GetIPPrefix()
		IPPrefixLen := strconv.Itoa(int(s.GetIPPrefixLen()))
		version, err := resolveIPVersionFromCIDR(IPPrefix, IPPrefixLen)
		if err != nil {
			return err
		}
		if ethertype != version {
			return errutil.ErrorBadRequestf("Rule subnet %v doesn't match ethertype %v",
				IPPrefix+"/"+IPPrefixLen, ethertype)
		}
	}
	return nil
}
