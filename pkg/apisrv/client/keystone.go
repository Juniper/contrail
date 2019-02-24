package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/keystone"
)

// Keystone is a keystone client.
type Keystone struct {
	URL        string
	HTTPClient *http.Client
}

type projectResponse struct {
	Project keystone.Project `json:"project"`
}

type projectListResponse struct {
	Projects []*keystone.Project `json:"projects"`
}

// GetProject gets project.
func (k *Keystone) GetProject(ctx context.Context, token string, id string) (*keystone.Project, error) {
	request, err := http.NewRequest(echo.GET, getURL(k.URL, "/projects/"+id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}
	request = auth.SetXClusterIDInHeader(ctx, request.WithContext(ctx))
	request.Header.Set("X-Auth-Token", token)
	var output projectResponse

	resp, err := k.HTTPClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{http.StatusOK}, resp.StatusCode); err != nil {
		return nil, errorFromResponse(err, resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	return &output.Project, nil
}

// ObtainToken gets authentication token.
func (k *Keystone) ObtainToken(
	ctx context.Context, id, password string, scope *keystone.Scope,
) (*http.Response, error) {
	return k.FetchToken(ctx, id, password, nil, scope)
}

func (k *Keystone) FetchToken(
	ctx context.Context, id, password string, domain *keystone.Domain,
	scope *keystone.Scope) (*http.Response, error) {
	if k.URL == "" {
		return nil, nil
	}

	var err error
	var dataJSON []byte
	if scope != nil {
		dataJSON, err = json.Marshal(&keystone.ScopedAuthRequest{
			Auth: &keystone.ScopedAuth{
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
	} else if domain != nil {
		dataJSON, err = json.Marshal(&keystone.UnScopedAuthRequest{
			Auth: &keystone.UnScopedAuth{
				Identity: &keystone.Identity{
					Methods: []string{"password"},
					Password: &keystone.Password{
						User: &keystone.User{
							Name:     id,
							Password: password,
							Domain:   domain,
						},
					},
				},
			},
		})
	} else {
		return nil, errors.Wrap(err, "unable to determine token type")
	}
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	if err != nil {
		return nil, err
	}
	request = auth.SetXClusterIDInHeader(ctx, request.WithContext(ctx))
	request.Header.Set("Content-Type", "application/json")

	resp, err := k.HTTPClient.Do(request)
	if err != nil {
		return nil, errorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{200, 201}, resp.StatusCode); err != nil {
		return resp, errorFromResponse(err, resp)
	}

	var authResponse keystone.AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return resp, errorFromResponse(err, resp)
	}

	return resp, nil
}

// GetProjectIDByName finds project id using project name.
func (k *Keystone) GetProjectIDByName(ctx context.Context,
	id, password, projectName string, domain *keystone.Domain) (string, error) {
	// Fetch unscoped token
	resp, err := k.FetchToken(ctx, id, password, domain, nil)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}
	token := resp.Header.Get("X-Subject-Token")
	// Get project list with unscoped token
	request, err := http.NewRequest(echo.GET, getURL(k.URL, "/auth/projects"), nil)
	if err != nil {
		return "", errors.Wrap(err, "creating HTTP request failed")
	}
	request = auth.SetXClusterIDInHeader(ctx, request.WithContext(ctx))
	request.Header.Set("X-Auth-Token", token)
	var output *projectListResponse
	resp, err = k.HTTPClient.Do(request)
	if err != nil {
		return "", errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{http.StatusOK}, resp.StatusCode); err != nil {
		return "", errorFromResponse(err, resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return "", errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	for _, project := range output.Projects {
		if project.Name == projectName {
			return project.ID, nil
		}
	}
	return "", fmt.Errorf("'%s' not a valid project name", projectName)
}
