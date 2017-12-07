package models

// FloatingIpDnsNotation

type FloatingIpDnsNotation string

// MakeFloatingIpDnsNotation makes FloatingIpDnsNotation
func MakeFloatingIpDnsNotation() FloatingIpDnsNotation {
	var data FloatingIpDnsNotation
	return data
}

// InterfaceToFloatingIpDnsNotation makes FloatingIpDnsNotation from interface
func InterfaceToFloatingIpDnsNotation(data interface{}) FloatingIpDnsNotation {
	return data.(FloatingIpDnsNotation)
}

// InterfaceToFloatingIpDnsNotationSlice makes a slice of FloatingIpDnsNotation from interface
func InterfaceToFloatingIpDnsNotationSlice(data interface{}) []FloatingIpDnsNotation {
	list := data.([]interface{})
	result := MakeFloatingIpDnsNotationSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIpDnsNotation(item))
	}
	return result
}

// MakeFloatingIpDnsNotationSlice() makes a slice of FloatingIpDnsNotation
func MakeFloatingIpDnsNotationSlice() []FloatingIpDnsNotation {
	return []FloatingIpDnsNotation{}
}
