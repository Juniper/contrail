package services

//FQNameToIDRequest defines fq_name to id request format
type FQNameToIDRequest struct {
	FQName []string `json:"fq_name"`
	Type   string   `json:"type"`
}

// FQNameToIDResponse defines FqNameToID response format.
type FQNameToIDResponse struct {
	UUID string `json:"uuid"`
}
