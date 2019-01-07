package logic

import (
	"net"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	ipV4 = 4
	ipV6 = 6
)

func getIPVersionFromCIDR(cidr string) (int8, error) {
	prefix := strings.Split(cidr, "/")
	return getIPVersion(prefix[0])
}

func getIPVersion(ipAddress string) (int8, error) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return -1, errors.Errorf("Invalid ip address: %s", ipAddress)
	}

	if v := ip.To4(); v != nil {
		return ipV4, nil
	}

	return ipV6, nil
}

func getIPPrefixAndPrefixLen(cidr string) (string, int64, error) {
	splittedCIDR := strings.Split(cidr, "/")
	if len(splittedCIDR) < 2 {
		return "", 0, errors.Errorf("Invalid cidr: %s", cidr)
	}

	ipPrefix := splittedCIDR[0]
	ipPrefixLen, err := strconv.ParseInt(splittedCIDR[1], 10, 64)
	if err != nil {
		return "", 0, err
	}

	return ipPrefix, ipPrefixLen, nil
}
