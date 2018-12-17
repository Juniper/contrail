package logic

// Create logic
func (n *Network) Create(rp RequestParameters) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
