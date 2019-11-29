package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/pkg/errors"
)

// NeutronPost sends Neutron request
func (h *HTTP) NeutronPost(ctx context.Context, r *logic.Request, expected []int) (logic.Response, error) {
	response, err := logic.MakeResponse(r.GetType())
	if err != nil {
		return nil, errors.Errorf("failed to get response type for request %v", r)
	}
	_, err = h.Do(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/neutron/%s", r.Context.Type),
		nil,
		r,
		&response,
		expected,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
