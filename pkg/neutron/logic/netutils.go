package logic

import (
	"net"
)

const (
	ipV4 = 4
	ipV6 = 6
)

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
