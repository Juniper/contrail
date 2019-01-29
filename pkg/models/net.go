package models

import (
	"net"

	"github.com/pkg/errors"
)

func parseIP(ipString string) (net.IP, error) {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return nil, errors.Errorf("couldn't parse ip address: " + ipString)
	}
	return ip, nil
}

func isIPInSubnet(subnet *net.IPNet, ipString string) error {
	ip, err := parseIP(ipString)
	if err != nil {
		return err
	}
	if !subnet.Contains(ip) {
		return errors.Errorf("address is out of cidr: " + subnet.String())
	}
	return nil
}
