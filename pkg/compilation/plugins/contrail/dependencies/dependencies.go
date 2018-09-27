/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Contrail Plugin Implementation
 *  The Dependencies map
 *  - object dependencies which needs to be evaluated when objects are
 *    created/modified or deleted
 *
 */

package dependencies

// ReactionMap - map of object dependencies
var ReactionMap = map[string]map[string]map[string]struct{}{
	"RoutingInstance": {
		"Self": {"VirtualNetwork": {}},
	},
	"VirtualMachineInterface": {
		"Self":           {"VirtualMachine": {}, "PortTuple": {}, "VirtualNetwork": {}, "BGPAsAService": {}},
		"VirtualNetwork": {"VirtualMachine": {}, "PortTuple": {}, "BGPAsAService": {}},
		"LogicalRouter":  {"VirtualNetwork": {}},
		"InstanceIP":     {"VirtualMachine": {}, "PortTuple": {}, "BGPAsAService": {}, "VirtualNetwork": {}},
		"FloatingIP":     {"VirtualMachine": {}, "PortTuple": {}},
		"AliasIP":        {"VirtualMachine": {}, "PortTuple": {}},
		"VirtualMachine": {"VirtualNetwork": {}},
		"PortTuple":      {"VirtualNetwork": {}},
		"BGPAsAService":  {},
	},
	"VirtualNetwork": {
		"Self":                    {"NetworkPolicy": {}, "RouteTable": {}, "VirtualNetwork": {}},
		"VirtualNetwork":          {},
		"RoutingInstance":         {"NetworkPolicy": {}},
		"NetworkPolicy":           {},
		"VirtualMachineInterface": {},
		"RouteTable":              {},
		"bgpvpn":                  {},
		"RoutingPolicy":           {},
	},
	"VirtualMachine": {
		"Self": {"ServiceInstance": {}},
		"VirtualMachineInterface": {"ServiceInstance": {}},
		"ServiceInstance":         {"VirtualMachineInterface": {}},
	},
	"PortTuple": {
		"Self": {"ServiceInstance": {}},
		"VirtualMachineInterface": {"ServiceInstance": {}},
		"ServiceInstance":         {"VirtualMachineInterface": {}},
	},
	"ServiceInstance": {
		"Self":           {"NetworkPolicy": {}, "VirtualMachine": {}, "PortTuple": {}},
		"RouteTable":     {"NetworkPolicy": {}},
		"RoutingPolicy":  {"NetworkPolicy": {}},
		"RouteAggregate": {"NetworkPolicy": {}},
		"VirtualMachine": {"NetworkPolicy": {}},
		"PortTuple":      {"NetworkPolicy": {}},
		"NetworkPolicy":  {"VirtualMachine": {}, "PortTuple": {}},
	},
	"NetworkPolicy": {
		"Self":            {"SecurityLoggingObject": {}, "VirtualNetwork": {}, "NetworkPolicy": {}, "ServiceInstance": {}},
		"ServiceInstance": {"VirtualNetwork": {}},
		"NetworkPolicy":   {"VirtualNetwork": {}},
		"VirtualNetwork":  {"VirtualNetwork": {}, "NetworkPolicy": {}, "ServiceInstance": {}},
	},
	"SecurityGroup": {
		"Self":          {"SecurityGroup": {}, "SecurityLoggingObject": {}},
		"SecurityGroup": {},
	},
	"SecurityLoggingObject": {
		"Self":          {},
		"NetworkPolicy": {},
		"SecurityGroup": {},
	},
	"RouteTable": {
		"Self":           {"VirtualNetwork": {}, "ServiceInstance": {}, "LogicalRouter": {}},
		"VirtualNetwork": {"ServiceInstance": {}},
		"LogicalRouter":  {"ServiceInstance": {}},
	},
	"LogicalRouter": {
		"Self": {"RouteTable": {}},
		"VirtualMachineInterface": {},
		"RouteTable":              {},
		"bgpvpn":                  {},
	},
	"FloatingIP": {
		"Self": {"VirtualMachineInterface": {}},
	},
	"AliasIP": {
		"Self": {"VirtualMachineInterface": {}},
	},
	"InstanceIP": {
		"Self": {"VirtualMachineInterface": {}},
	},
	"BGPAsAService": {
		"Self": {"BGPRouter": {}},
		"VirtualMachineInterface": {"BGPRouter": {}},
	},
	"BGPRouter": {
		"Self":          {},
		"BGPAsAService": {},
	},
	"GlobalSystemConfig": {
		"Self": {},
	},
	"RoutingPolicy": {
		"Self": {"ServiceInstance": {}, "VirtualNetwork": {}},
	},
	"RouteAggregate": {
		"Self": {"ServiceInstance": {}},
	},
	"bgpvpn": {
		"Self":           {"VirtualNetwork": {}, "LogicalRouter": {}},
		"VirtualNetwork": {},
		"LogicalRouter":  {},
	},
}
