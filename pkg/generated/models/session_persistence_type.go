package models

// SessionPersistenceType

type SessionPersistenceType string

// MakeSessionPersistenceType makes SessionPersistenceType
func MakeSessionPersistenceType() SessionPersistenceType {
	var data SessionPersistenceType
	return data
}

// InterfaceToSessionPersistenceType makes SessionPersistenceType from interface
func InterfaceToSessionPersistenceType(data interface{}) SessionPersistenceType {
	return data.(SessionPersistenceType)
}

// InterfaceToSessionPersistenceTypeSlice makes a slice of SessionPersistenceType from interface
func InterfaceToSessionPersistenceTypeSlice(data interface{}) []SessionPersistenceType {
	list := data.([]interface{})
	result := MakeSessionPersistenceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSessionPersistenceType(item))
	}
	return result
}

// MakeSessionPersistenceTypeSlice() makes a slice of SessionPersistenceType
func MakeSessionPersistenceTypeSlice() []SessionPersistenceType {
	return []SessionPersistenceType{}
}
