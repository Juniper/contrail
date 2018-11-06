package services

type FQNameToIDRequest struct {
	FQName []string `json:"fq_name"`
	Type   string   `json:"type"`
}

// FQNameToIDResponse defines FqNameToID response format.
type FQNameToIDResponse struct {
	UUID string `json:"uuid"`
}
