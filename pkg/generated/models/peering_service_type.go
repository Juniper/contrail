package models

// PeeringServiceType

type PeeringServiceType string

// MakePeeringServiceType makes PeeringServiceType
func MakePeeringServiceType() PeeringServiceType {
	var data PeeringServiceType
	return data
}

// InterfaceToPeeringServiceType makes PeeringServiceType from interface
func InterfaceToPeeringServiceType(data interface{}) PeeringServiceType {
	return data.(PeeringServiceType)
}

// InterfaceToPeeringServiceTypeSlice makes a slice of PeeringServiceType from interface
func InterfaceToPeeringServiceTypeSlice(data interface{}) []PeeringServiceType {
	list := data.([]interface{})
	result := MakePeeringServiceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPeeringServiceType(item))
	}
	return result
}

// MakePeeringServiceTypeSlice() makes a slice of PeeringServiceType
func MakePeeringServiceTypeSlice() []PeeringServiceType {
	return []PeeringServiceType{}
}
