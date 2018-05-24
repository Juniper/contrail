package models

import (
	"strings"

	"net"

	"strconv"

	"github.com/graph-gophers/graphql-go/errors"
)

const (
	BgpRtgtMinId = 8000000
)

const (
	RouteTargetPrefix = "target"
)

// Check if route target was defined by user
func (rt *RouteTarget) IsUserDefined(globalAsn int64) (bool, error) {
	return IsRouteTargetUserDefined(rt.FQName, globalAsn)
}

// Check if route target represented as a string was user defined
func IsStringRouteTargetUserDefined(routeTarget string, globalAsn int64) (bool, error) {
	return IsRouteTargetUserDefined(strings.Split(routeTarget, ":"), globalAsn)
}

// Check if route target represented as a slice was user defined
func IsRouteTargetUserDefined(routeTarget []string, globalAsn int64) (bool, error) {
	ip, asn, target, err := parseRouteTarget(routeTarget)
	if err != nil {
		return false, err
	}

	// If ip is specified, rt is user defined for sure
	if ip != nil {
		return true, nil
	}
	if int64(asn) == globalAsn && target >= BgpRtgtMinId {
		return false, nil
	}

	return true, nil
}

func parseRouteTarget(routeTarget []string) (ip net.IP, asn int, target int, err error) {

	if len(routeTarget) != 3 || routeTarget[0] != RouteTargetPrefix {
		return nil, 0, 0, errors.Errorf("Invalid RouteTarget specified: %v \nRoute target must be of the format 'target:<asn>:<number>' or 'target:<ip>:<number>'", routeTarget)
	}

	ip = net.ParseIP(routeTarget[1])
	if ip == nil {
		asn, err = strconv.Atoi(routeTarget[1])
		if err != nil {
			return nil, 0, 0, errors.Errorf("Invalid RouteTarget specified: %v \nInvalid asn (should be ip or int) %v", routeTarget, err)
		}
	}
	target, err = strconv.Atoi(routeTarget[2])
	if err != nil {
		return nil, 0, 0, errors.Errorf("Invalid RouteTarget specified: %v \nInvalid target id (should be int) %v", routeTarget, err)
	}

	return ip, asn, target, nil
}
