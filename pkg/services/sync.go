package services

//RESTResource represents REST resource request.
type RESTResource struct {
	Kind      string      `json:"kind"`
	Data      interface{} `json:"data"`
	Operation string      `json:"operation"`
}

//RESTSyncRequest has multiple rest requests.
type RESTSyncRequest struct {
	Resources []*RESTResource `json:"resources"`
}
