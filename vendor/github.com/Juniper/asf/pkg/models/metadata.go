package models

// Metadata represents resource meta data.
type Metadata struct {
	UUID   string
	FQName []string
	Type   string
}

// FQNameMetadata creates minimal valid Metadata with FQName and type.
func FQNameMetadata(fqName []string, typeName string) Metadata {
	return Metadata{FQName: fqName, Type: typeName}
}

// UUIDMetadata creates minimal valid Metadata with UUID.
func UUIDMetadata(uuid string) Metadata {
	return Metadata{UUID: uuid}
}
