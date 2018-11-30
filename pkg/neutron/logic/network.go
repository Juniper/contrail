package logic

// Create logic
func (n *Network) Create(ctx RequestParameters) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
