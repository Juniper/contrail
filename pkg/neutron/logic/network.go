package logic

import "context"

// Create logic
func (n *Network) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
