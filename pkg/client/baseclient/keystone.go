package baseclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	ctx = WithXAuthToken(ctx, token)

	var response projectResponse
	if err := k.doREST(
		ctx, http.MethodGet, k.getURL("/projects/"+id), []int{http.StatusOK}, nil, &response,
	); err != nil {
		return nil, err
	}
	return &response.Project, nil
}

// GetProjectIDByName finds project id using project name.
func (k *Keystone) GetProjectIDByName(
	ctx context.Context, id, password, projectName string, domain *keystone.Domain,
) (string, error) {
	// Fetch unscoped token
	token, err := k.obtainUnscopedToken(ctx, id, password, domain)
	if err != nil {
		return "", err
	}
	ctx = WithXAuthToken(ctx, token)

	var response projectListResponse
	if err := k.doREST(
		ctx, http.MethodGet, k.getURL("/auth/projects"), []int{http.StatusOK}, nil, &response,
	); err != nil {
		return "", err
	}

	for _, project := range response.Projects {
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
	ctx = WithHTTPHeader(ctx, contentTypeHeader, applicationJSONValue)
	resp, err := k.do(
		ctx, http.MethodPost, k.getURL("/auth/tokens"), []int{http.StatusOK, http.StatusCreated}, authRequest,
	)
	if err != nil {
		return "", err
	}
	return resp.Header.Get(xSubjectTokenHeader), nil
}

// CreateUser creates user in keystone.
func (k *Keystone) CreateUser(ctx context.Context, user keystone.User) (*keystone.User, error) {
	var response keystone.CreateUserResponse
	if err := k.doREST(
		ctx,
		http.MethodPost,
		k.getURL("/users/"),
		[]int{http.StatusCreated},
		keystone.CreateUserRequest{User: user},
		&response,
	); err != nil {
		return nil, err
	}
	return &response.User, nil
}

// AddRole adds role to user in keystone.
func (k *Keystone) AddRole(ctx context.Context, user keystone.User, role keystone.Role) error {
	url := k.getURL("/roles/" + role.Project.ID + "/" + user.ID + "/" + role.ID)
	return k.doREST(ctx, http.MethodPut, url, []int{http.StatusNoContent}, nil, nil)
}

func (k *Keystone) do(
	ctx context.Context, method, url string, expectedCodes []int, input interface{},
) (*http.Response, error) {
	var payload io.Reader
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling CreateUserRequest")
		}
		payload = bytes.NewReader(b)
	}
	request, err := http.NewRequest(method, url, payload)
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

	if err := checkStatusCode(expectedCodes, resp.StatusCode); err != nil {
		return nil, errorFromResponse(err, resp)
	}

	return resp, nil
}

func (k *Keystone) doREST(
	ctx context.Context, method, url string, expectedCodes []int, input, output interface{},
) error {
	resp, err := k.do(ctx, method, url, expectedCodes, input)
	if err != nil {
		return err
	}
	if output != nil {
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
		}
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
