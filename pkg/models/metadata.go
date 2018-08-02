package models

type FQNameUUIDPair struct {
	UUID   string
	FQName []string
}

// MetaData represents resource meta data.
type MetaData struct {
	UUID   string
	FQName []string
	Type   string
}
