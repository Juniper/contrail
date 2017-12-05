package models

// PeeringServiceType

type PeeringServiceType string

func MakePeeringServiceType() PeeringServiceType {
	var data PeeringServiceType
	return data
}

func InterfaceToPeeringServiceType(data interface{}) PeeringServiceType {
	return data.(PeeringServiceType)
}

func InterfaceToPeeringServiceTypeSlice(data interface{}) []PeeringServiceType {
	list := data.([]interface{})
	result := MakePeeringServiceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPeeringServiceType(item))
	}
	return result
}

func MakePeeringServiceTypeSlice() []PeeringServiceType {
	return []PeeringServiceType{}
}
