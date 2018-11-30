package logic

func (n *Network) Create(ctx Context) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
