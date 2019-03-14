package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/keystone"
)

// Keystone is a keystone client.
type Keystone struct {
	URL        string
	HTTPClient *http.Client
}

func (k *Keystone) getURL(path string) string {
	return k.URL + path
}

type projectResponse struct {
	Project keystone.Project `json:"project"`
}

type projectListResponse struct {
	Projects []*keystone.Project `json:"projects"`
}

// GetProjects fetch all projects.
func (k *Keystone) GetProjects(ctx context.Context, token string) ([]*keystone.Project, error) {
	var response projectListResponse
	err := k.makeProjectRequest(ctx, token, "/projects", &response)
	if err != nil {
		return nil, err
	}
	return response.Projects, nil
}

// GetProject fetch single project by ID.
func (k *Keystone) GetProject(ctx context.Context, token string, id string) (*keystone.Project, error) {
	var response projectResponse
	err := k.makeProjectRequest(ctx, token, "/projects/"+id, &response)
	if err != nil {
		return nil, err
	}
	return &response.Project, nil
}

func (k *Keystone) makeProjectRequest(ctx context.Context, token, url string, out interface{}) error {
	request, err := http.NewRequest(echo.GET, k.getURL(url), nil)
	if err != nil {
		return errors.Wrap(err, "creating HTTP request failed")
	}
	request = auth.SetXClusterIDInHeader(ctx, request.WithContext(ctx))
	request.Header.Set("X-Auth-Token", token)

	resp, err := k.HTTPClient.Do(request)
	if err != nil {
		return errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{http.StatusOK}, resp.StatusCode); err != nil {
		return errorFromResponse(err, resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	return nil
}

// GetProjectIDByName finds project id using project name.
func (k *Keystone) GetProjectIDByName(ctx context.Context,
	id, password, projectName string, domain *keystone.Domain) (string, error) {
	// Fetch unscoped token
	resp, err := k.ObtainUnScopedToken(ctx, id, password, domain)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}
	token := resp.Header.Get("X-Subject-Token")
	// Get project list with unscoped token
	request, err := http.NewRequest(echo.GET, k.getURL("/auth/projects"), nil)
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

// ObtainToken gets authentication token.
func (k *Keystone) ObtainToken(
	ctx context.Context, id, password string, scope *keystone.Scope,
) (*http.Response, error) {
	if k.URL == "" {
		return nil, nil
	}

	var err error
	var dataJSON []byte
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
	if err != nil {
		return nil, err
	}
	return k.FetchToken(ctx, dataJSON)
}

// ObtainUnScopedToken gets unscoped authentication token.
func (k *Keystone) ObtainUnScopedToken(
	ctx context.Context, id, password string, domain *keystone.Domain,
) (*http.Response, error) {
	if k.URL == "" {
		return nil, nil
	}

	var err error
	var dataJSON []byte
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
	if err != nil {
		return nil, err
	}
	return k.FetchToken(ctx, dataJSON)
}

// FetchToken gets scoped/unscoped token
func (k *Keystone) FetchToken(ctx context.Context, dataJSON []byte,
) (*http.Response, error) {
	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	if err != nil {
		return nil, err
	}
	request = auth.SetXAuthTokenInHeader(ctx, request)
	request = auth.SetXClusterIDInHeader(ctx, request)
	request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")

	startedAt := time.Now()
	resp, err := k.HTTPClient.Do(request)
	durationInUsec := time.Since(startedAt) / time.Microsecond
	if err != nil {
		return nil, errorFromResponse(err, resp)
	}

	if c := collector.FromContext(ctx); c != nil {
		c.Send(analytics.VncAPILatencyStatsLog(ctx, "VALIDATE", "KEYSTONE", int64(durationInUsec)))
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
