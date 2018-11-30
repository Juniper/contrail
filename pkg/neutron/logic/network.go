package logic

func (n *NetworkRequest) Create(ctx Context) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
