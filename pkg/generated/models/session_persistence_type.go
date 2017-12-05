package models

// SessionPersistenceType

type SessionPersistenceType string

func MakeSessionPersistenceType() SessionPersistenceType {
	var data SessionPersistenceType
	return data
}

func InterfaceToSessionPersistenceType(data interface{}) SessionPersistenceType {
	return data.(SessionPersistenceType)
}

func InterfaceToSessionPersistenceTypeSlice(data interface{}) []SessionPersistenceType {
	list := data.([]interface{})
	result := MakeSessionPersistenceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSessionPersistenceType(item))
	}
	return result
}

func MakeSessionPersistenceTypeSlice() []SessionPersistenceType {
	return []SessionPersistenceType{}
}
