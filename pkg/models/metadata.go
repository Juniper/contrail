package models

type FQNameUUIDPair struct {
	UUID   string
	FQName []string
}

// MetaData represents resource meta data.
type MetaData struct {
	FQNameUUIDPair
	Type string
}
