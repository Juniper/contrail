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
		return fmt.Errorf("Invalid Port range specified")
	}
	return nil
}

func GetDefaultBGPaaServiceParameters() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{DefaultPortRangeStart, DefaultPortRangeEnd}
}

func (ports *BGPaaServiceParametersType) ContainsRange(other *BGPaaServiceParametersType) bool {
	return ports.PortStart <= other.PortStart && ports.PortEnd >= other.PortEnd
}
