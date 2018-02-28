package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallSequence makes FirewallSequence
// nolint
func MakeFirewallSequence() *FirewallSequence {
	return &FirewallSequence{
		//TODO(nati): Apply default
		Sequence: "",
	}
}

// MakeFirewallSequence makes FirewallSequence
// nolint
func InterfaceToFirewallSequence(i interface{}) *FirewallSequence {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallSequence{
		//TODO(nati): Apply default
		Sequence: common.InterfaceToString(m["sequence"]),
	}
}

// MakeFirewallSequenceSlice() makes a slice of FirewallSequence
// nolint
func MakeFirewallSequenceSlice() []*FirewallSequence {
	return []*FirewallSequence{}
}

// InterfaceToFirewallSequenceSlice() makes a slice of FirewallSequence
// nolint
func InterfaceToFirewallSequenceSlice(i interface{}) []*FirewallSequence {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallSequence{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallSequence(item))
	}
	return result
}
