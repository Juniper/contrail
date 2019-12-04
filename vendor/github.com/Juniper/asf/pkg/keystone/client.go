package keystone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/httputil"
)

const (
	// xAuthTokenHeader is a header used by keystone to store user auth tokens.
	xAuthTokenHeader = "X-Auth-Token"
	// xSubjectTokenHeader is a header used by keystone to return new tokens.
	xSubjectTokenHeader = "X-Subject-Token"

	contentTypeHeader    = "Content-Type"
	applicationJSONValue = "application/json"
)

// WithXAuthToken creates child context with Auth Token
func WithXAuthToken(ctx context.Context, token string) context.Context {
	return httputil.WithHTTPHeader(ctx, xAuthTokenHeader, token)
}

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a keystone client.
type Client struct {
	URL      string
	HTTPDoer doer
}

type projectResponse struct {
	Project Project `json:"project"`
}

type projectListResponse struct {
	Projects []*Project `json:"projects"`
}

// GetProject gets project.
func (k *Client) GetProject(ctx context.Context, token string, id string) (*Project, error) {
	request, err := http.NewRequest(http.MethodGet, k.getURL("/projects/"+id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	httputil.SetContextHeaders(request)
	request.Header.Set(xAuthTokenHeader, token)
	var output projectResponse

	resp, err := k.HTTPDoer.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = httputil.CheckStatusCode([]int{http.StatusOK}, resp.StatusCode); err != nil {
		return nil, httputil.ErrorFromResponse(err, resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, errors.Wrapf(httputil.ErrorFromResponse(err, resp), "decoding response body failed")
	}

	return &output.Project, nil
}

// GetProjectIDByName finds project id using project name.
func (k *Client) GetProjectIDByName(
	ctx context.Context, id, password, projectName string, domain *Domain) (string, error) {
	// Fetch unscoped token
	token, err := k.obtainUnscopedToken(ctx, id, password, domain)
	if err != nil {
		return "", err
	}
	// Get project list with unscoped token
	request, err := http.NewRequest(http.MethodGet, k.getURL("/auth/projects"), nil)
	if err != nil {
		return "", errors.Wrap(err, "creating HTTP request failed")
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	httputil.SetContextHeaders(request)
	request.Header.Set(xAuthTokenHeader, token)

	var output *projectListResponse
	resp, err := k.HTTPDoer.Do(request)
	if err != nil {
		return "", errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = httputil.CheckStatusCode([]int{http.StatusOK}, resp.StatusCode); err != nil {
		return "", httputil.ErrorFromResponse(err, resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return "", errors.Wrapf(httputil.ErrorFromResponse(err, resp), "decoding response body failed")
	}

	for _, project := range output.Projects {
		if project.Name == projectName {
			return project.ID, nil
		}
	}
	return "", fmt.Errorf("'%s' not a valid project name", projectName)
}

func (k *Client) getURL(path string) string {
	return k.URL + path
}

// obtainUnscopedToken gets unscoped authentication token.
func (k *Client) obtainUnscopedToken(
	ctx context.Context, id, password string, domain *Domain,
) (string, error) {
	if k.URL == "" {
		return "", nil
	}
	return k.fetchToken(ctx, &UnScopedAuthRequest{
		Auth: &UnScopedAuth{
			Identity: &Identity{
				Methods: []string{"password"},
				Password: &Password{
					User: &User{
						Name:     id,
						Password: password,
						Domain:   domain,
					},
				},
			},
		},
	})
}

// ObtainToken gets authentication token.
func (k *Client) ObtainToken(ctx context.Context, id, password string, scope *Scope) (string, error) {
	if k.URL == "" {
		return "", nil
	}
	return k.fetchToken(ctx, &ScopedAuthRequest{
		Auth: &ScopedAuth{
			Identity: &Identity{
				Methods: []string{"password"},
				Password: &Password{
					User: &User{
						Name:     id,
						Password: password,
						Domain:   scope.GetDomain(),
					},
				},
			},
			Scope: scope,
		},
	})
}

// fetchToken gets scoped/unscoped token.
func (k *Client) fetchToken(ctx context.Context, authRequest interface{}) (string, error) {
	d, err := json.Marshal(authRequest)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewReader(d))
	if err != nil {
		return "", err
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	httputil.SetContextHeaders(request)
	request.Header.Set(contentTypeHeader, applicationJSONValue)

	resp, err := k.HTTPDoer.Do(request)
	if err != nil {
		return "", httputil.ErrorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = httputil.CheckStatusCode([]int{200, 201}, resp.StatusCode); err != nil {
		return "", httputil.ErrorFromResponse(err, resp)
	}

	return resp.Header.Get(xSubjectTokenHeader), nil
}
