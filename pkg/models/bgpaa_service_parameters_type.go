package models

import (
	"fmt"
	"math"
)

const (
	DefaultPortRangeStart = 50000
	DefaultPortRangeEnd   = 50512
)

func (ports *BGPaaServiceParametersType) ValidatePortRange() error {

	if (ports.PortStart > ports.PortEnd) || (ports.PortStart <= 0) || (ports.PortEnd > math.MaxUint16) {
		return fmt.Errorf("Invalid Port range specified (%d : %d)", ports.PortStart, ports.PortEnd)
	}
	return nil
}

func GetDefaultBGPaaServiceParameters() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{PortStart: DefaultPortRangeStart, PortEnd: DefaultPortRangeEnd}
}

func (ports *BGPaaServiceParametersType) EnclosesRange(other *BGPaaServiceParametersType) bool {
	if ports != nil && other != nil {
		return ports.PortStart <= other.PortStart && ports.PortEnd >= other.PortEnd
	}
	return false
}
