package logic

import (
	"net"

	"github.com/pkg/errors"
)

// HostPrefixes is a combination of Next Hop IP Addresses and
// destination CIDRs that can be reached by those Next Hops.
type hostPrefixes map[string][]string

func makeHostPrefixes() hostPrefixes {
	return make(map[string][]string)
}

func getHostPrefixes(hostRoutes []*RouteTableType, subnetCIDR string) (hostPrefixes, error) {
	solved, unresolved, err := splitRoutesByReachAbility(hostRoutes, subnetCIDR)
	if err != nil {
		return nil, err
	}

	hostPrefs := makeHostPrefixes()

	for _, route := range solved {
		hostPrefs.addDestinationToIP(route.Destination, route.Nexthop)
		reached, notReached, err := extractDestinationsReachableFromCIDR(route.Destination, unresolved)
		if err != nil {
			return nil, err
		}
		unresolved = notReached
		hostPrefs.addDestinationsToIP(reached, route.Nexthop)
	}
	return hostPrefs, nil
}

func extractDestinationsReachableFromCIDR(
	cidr string, unresolvedRoutes []*RouteTableType,
) ([]string, []*RouteTableType, error) {
	reachedDestinations := []string{}

	reached, notReached, err := splitRoutesByReachAbility(unresolvedRoutes, cidr)
	if err != nil {
		return nil, nil, err
	}

	for _, route := range reached {
		reachedDestinations = append(reachedDestinations, route.Destination)
		destinations, unresolved, err := extractDestinationsReachableFromCIDR(route.Destination, notReached)
		if err != nil {
			return nil, nil, err
		}
		reachedDestinations = append(reachedDestinations, destinations...)
		notReached = unresolved
	}
	return reachedDestinations, notReached, nil
}

func splitRoutesByReachAbility(
	routes []*RouteTableType, cidr string,
) ([]*RouteTableType, []*RouteTableType, error) {
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, nil, err
	}

	possible := []*RouteTableType{}
	notPossible := []*RouteTableType{}

	for _, route := range routes {
		nextHopIP := net.ParseIP(route.Nexthop)
		if nextHopIP == nil {
			return nil, nil, errors.Errorf("Following NextHop route cannot be parsed: %v", route.Nexthop)
		}
		if subnet.Contains(nextHopIP) {
			possible = append(possible, route)
		} else {
			notPossible = append(notPossible, route)
		}
	}

	return possible, notPossible, nil
}

func (hp hostPrefixes) getDestinationsForIP(ip string) []string {
	return hp[ip]
}

func (hp hostPrefixes) addDestinationToIP(destination string, ip string) {
	hp[ip] = append(hp[ip], destination)
}

func (hp hostPrefixes) addDestinationsToIP(destinations []string, ip string) {
	hp[ip] = append(hp[ip], destinations...)
}

func (hp hostPrefixes) removeDestinationsForIP(ip string) {
	delete(hp, ip)
}

func (hp hostPrefixes) getIPAddresses() []string {
	ips := []string{}
	for ip := range hp {
		if len(hp[ip]) > 0 {
			ips = append(ips, ip)
		}
	}
	return ips
}

func (hp hostPrefixes) hasAnyDestinationsForIP(ip string) bool {
	return len(hp[ip]) > 0
}

func (hp hostPrefixes) isEmpty() bool {
	for _, addr := range hp.getIPAddresses() {
		if hp.hasAnyDestinationsForIP(addr) {
			return false
		}
	}
	return true
}
