package models

import (
	"net"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
)

// ValidateWithEthertype checks if IP version from CIDR matches ethertype
// and throws an error if it doesn't.
func (s *SubnetType) ValidateWithEthertype(ethertype string) error {
	if s == nil {
		return nil
	}
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
	return nil
}

func resolveIPVersionFromCIDR(IPPrefix, IPPrefixLen string) (string, error) {
	network, _, err := net.ParseCIDR(IPPrefix + "/" + IPPrefixLen)
	if err != nil {
		return "", errutil.ErrorBadRequestf("Cannot parse address %v/%v. %v.",
			IPPrefix, IPPrefixLen, err)
	}
	switch {
	case network.To4() != nil:
		return "IPv4", nil
	case network.To16() != nil:
		return "IPv6", nil
	default:
		return "", errutil.ErrorBadRequestf("Cannot resolve ip version %v/%v.",
			IPPrefix, IPPrefixLen)
	}
}