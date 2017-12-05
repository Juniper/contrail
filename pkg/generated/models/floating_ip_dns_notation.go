package models

// FloatingIpDnsNotation

type FloatingIpDnsNotation string

func MakeFloatingIpDnsNotation() FloatingIpDnsNotation {
	var data FloatingIpDnsNotation
	return data
}

func InterfaceToFloatingIpDnsNotation(data interface{}) FloatingIpDnsNotation {
	return data.(FloatingIpDnsNotation)
}

func InterfaceToFloatingIpDnsNotationSlice(data interface{}) []FloatingIpDnsNotation {
	list := data.([]interface{})
	result := MakeFloatingIpDnsNotationSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIpDnsNotation(item))
	}
	return result
}

func MakeFloatingIpDnsNotationSlice() []FloatingIpDnsNotation {
	return []FloatingIpDnsNotation{}
}
