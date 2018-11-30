package models

func (n *NetworkRequest) Create(ctx CreateContext) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
