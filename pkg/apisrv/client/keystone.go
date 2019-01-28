package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// Keystone is a keystone client
type Keystone struct {
	URL        string
	HTTPClient *http.Client
}

type projectResponse struct {
	Project keystone.Project `json:"project"`
}

// GetProject gets projects
func (k *Keystone) GetProject(ctx context.Context, token string, id string) (*keystone.Project, error) {
	request, err := http.NewRequest(echo.GET, getURL(k.URL, "/projects/"+id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}

	request.Header.Set("X-Auth-Token", token)
	var output projectResponse

	resp, err := k.HTTPClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}

	err = checkStatusCode([]int{http.StatusOK}, resp.StatusCode)
	if err != nil {
		return nil, errorFromResponse(err, resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	defer resp.Body.Close()
	return &output.Project, nil
}

// ObtainToken gets authentication token.
func (k *Keystone) ObtainToken(
	ctx context.Context, id string, password string, scope *keystone.Scope,
) (string, error) {
	if k.URL == "" {
		return "", nil
	}

	dataJSON, err := json.Marshal(&keystone.AuthRequest{
		Auth: &keystone.Auth{
			Identity: &keystone.Identity{
				Methods: []string{"password"},
				Password: &keystone.Password{
					User: &keystone.User{
						Name:     id,
						Password: password,
						Domain:   scope.GetDomain(),
					},
				},
			},
			Scope: scope,
		},
	})
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	request = request.WithContext(ctx)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := k.HTTPClient.Do(request)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode([]int{200, 201}, resp.StatusCode)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}

	var authResponse keystone.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}

	return resp.Header.Get("X-Subject-Token"), nil
}
