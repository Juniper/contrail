package baseclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/pkg/errors"
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
	return WithHTTPHeader(ctx, xAuthTokenHeader, token)
}

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Keystone is a keystone client.
type Keystone struct {
	URL      string
	HTTPDoer doer
}

type projectResponse struct {
	Project keystone.Project `json:"project"`
}

type projectListResponse struct {
	Projects []*keystone.Project `json:"projects"`
}

// GetProject gets project.
func (k *Keystone) GetProject(ctx context.Context, token string, id string) (*keystone.Project, error) {
	request, err := http.NewRequest(http.MethodGet, k.getURL("/projects/"+id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	SetContextHeaders(request)
	request.Header.Set(xAuthTokenHeader, token)
	var output projectResponse

	resp, err := k.HTTPDoer.Do(request)
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

// GetProjectIDByName finds project id using project name.
func (k *Keystone) GetProjectIDByName(
	ctx context.Context, id, password, projectName string, domain *keystone.Domain) (string, error) {
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
	SetContextHeaders(request)
	request.Header.Set(xAuthTokenHeader, token)

	var output *projectListResponse
	resp, err := k.HTTPDoer.Do(request)
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

func (k *Keystone) getURL(path string) string {
	return k.URL + path
}

// obtainUnscopedToken gets unscoped authentication token.
func (k *Keystone) obtainUnscopedToken(
	ctx context.Context, id, password string, domain *keystone.Domain,
) (string, error) {
	if k.URL == "" {
		return "", nil
	}
	return k.fetchToken(ctx, &keystone.UnScopedAuthRequest{
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
}

// ObtainToken gets authentication token.
func (k *Keystone) ObtainToken(ctx context.Context, id, password string, scope *keystone.Scope) (string, error) {
	if k.URL == "" {
		return "", nil
	}
	return k.fetchToken(ctx, &keystone.ScopedAuthRequest{
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
}

// fetchToken gets scoped/unscoped token.
func (k *Keystone) fetchToken(ctx context.Context, authRequest interface{}) (string, error) {
	d, err := json.Marshal(authRequest)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewReader(d))
	if err != nil {
		return "", err
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	SetContextHeaders(request)
	request.Header.Set(contentTypeHeader, applicationJSONValue)

	resp, err := k.HTTPDoer.Do(request)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{200, 201}, resp.StatusCode); err != nil {
		return "", errorFromResponse(err, resp)
	}

	return resp.Header.Get(xSubjectTokenHeader), nil
}

// CreateUser creates user in keystone.
func (k *Keystone) CreateUser(ctx context.Context, user keystone.User) (*keystone.User, error) {
	b, err := json.Marshal(keystone.CreateUserRequest{User: user})
	if err != nil {
		return nil, errors.Wrap(err, "marshalling CreateUserRequest")
	}
	request, err := http.NewRequest(http.MethodPost, k.getURL("/users/"), bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	SetContextHeaders(request)

	resp, err := k.HTTPDoer.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err := checkStatusCode([]int{http.StatusCreated}, resp.StatusCode); err != nil {
		return nil, errorFromResponse(err, resp)
	}

	var userResp keystone.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	return &userResp.User, nil
}

// AddRole adds role to user in keystone.
func (k *Keystone) AddRole(ctx context.Context, user keystone.User, role keystone.Role) error {
	request, err := http.NewRequest(http.MethodPut, k.getURL("/roles/"+role.Project.ID+"/"+user.ID+"/"+role.ID), nil)

	if err != nil {
		return errors.Wrap(err, "creating HTTP request failed")
	}
	request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	SetContextHeaders(request)

	resp, err := k.HTTPDoer.Do(request)
	if err != nil {
		return errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err := checkStatusCode([]int{http.StatusNoContent}, resp.StatusCode); err != nil {
		return errorFromResponse(err, resp)
	}

	return nil
}

// CreateServiceUser creates service user in keystone.
func (k *Keystone) CreateServiceUser(ctx context.Context, user keystone.User) (*keystone.User, error) {
	u, err := k.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if err := k.AddRole(ctx, user, keystone.Role{Name: "admin", Project: &keystone.Project{}}); err != nil {
		return nil, err
	}

	return u, nil
}
