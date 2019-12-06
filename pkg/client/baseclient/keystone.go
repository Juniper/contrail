package baseclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

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

	serviceProjectName = "service"
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
	if _, err := k.do(
		ctx, http.MethodGet, "/projects/"+id, []int{http.StatusOK}, nil, &response,
	); err != nil {
		return nil, err
	}
	return &response.Project, nil
}

func (k *Keystone) do(
	ctx context.Context, method, requestPath string, expectedCodes []int, input, output interface{},
) (*http.Response, error) {
	var payload io.Reader
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling keystone request")
		}
		payload = bytes.NewReader(b)
	}
	request, err := http.NewRequest(method, k.URL+requestPath, payload)
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

	if output != nil {
		if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
			return nil, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
		}
	}

	return resp, nil
}

// GetProjectIDByName finds project id using project name.
func (k *Keystone) GetProjectIDByName(
	ctx context.Context, projectName string, domain *keystone.Domain,
) (string, error) {
	var response projectListResponse
	if _, err := k.do(
		ctx, http.MethodGet, fmt.Sprintf("/projects?name=%s", projectName), []int{http.StatusOK}, nil, &response,
	); err != nil {
		return "", err
	}

	for _, project := range response.Projects {
		if project.Name == projectName {
			return project.ID, nil
		}
	}
	return "", fmt.Errorf("project '%s' does not exist", projectName)
}

// ObtainUnscopedToken gets unscoped authentication token.
func (k *Keystone) ObtainUnscopedToken(
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

// fetchToken gets scoped/unscoped token.
func (k *Keystone) fetchToken(ctx context.Context, authRequest interface{}) (string, error) {
	ctx = WithHTTPHeader(ctx, contentTypeHeader, applicationJSONValue)
	var out map[string]interface{}
	resp, err := k.do(
		ctx, http.MethodPost, "/auth/tokens", []int{http.StatusOK, http.StatusCreated}, authRequest, &out,
	)
	if err != nil {
		return "", err
	}
	return resp.Header.Get(xSubjectTokenHeader), nil
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

// CreateUserRequest represents a keystone user creation request.
type CreateUserRequest struct {
	keystone.User `json:"user"`
}

// CreateUserResponse represents a keystone user creation response.
type CreateUserResponse CreateUserRequest

// CreateUser creates user in keystone.
func (k *Keystone) CreateUser(ctx context.Context, user keystone.User) (keystone.User, error) {
	var response CreateUserResponse
	httpResp, err := k.do(
		ctx,
		http.MethodPost,
		"/users/",
		[]int{http.StatusCreated, http.StatusConflict},
		CreateUserRequest{User: user},
		&response,
	)
	if err != nil {
		return keystone.User{}, err
	}
	if httpResp.StatusCode == http.StatusConflict {
		return k.GetUserByName(ctx, user.Name)
	}
	return response.User, nil
}

type userListResponse struct {
	Users []keystone.User `json:"users" yaml:"users"`
}

// GetRoleByName looks for role with provided name in keystone.
func (k *Keystone) GetUserByName(ctx context.Context, userName string) (keystone.User, error) {
	var users userListResponse
	if _, err := k.do(
		ctx, http.MethodGet, fmt.Sprintf("/users?name=%v", userName), []int{http.StatusOK}, nil, &users,
	); err != nil {
		return keystone.User{}, err
	}
	for _, user := range users.Users {
		if user.Name == userName {
			return user, nil
		}
	}
	return keystone.User{}, fmt.Errorf("user '%s' does not exist", userName)
}

// CreateServiceUser creates service user in keystone.
func (k *Keystone) CreateServiceUser(ctx context.Context, user keystone.User) (keystone.User, error) {
	projectID, err := k.GetProjectIDByName(ctx, serviceProjectName, keystone.DefaultDomain())
	if err != nil {
		return keystone.User{}, err
	}
	user, err = k.CreateUser(ctx, user)
	if err != nil {
		return keystone.User{}, err
	}
	role, err := k.GetRoleByName(ctx, "admin")
	if err != nil {
		return keystone.User{}, err
	}
	role.Project = &keystone.Project{ID: projectID}
	if err := k.AssignProjectRoleOnUser(ctx, user, role); err != nil {
		return keystone.User{}, err
	}

	return user, nil
}

// AssignProjectRoleOnUser adds role to user in keystone.
func (k *Keystone) AssignProjectRoleOnUser(ctx context.Context, user keystone.User, role keystone.Role) error {
	path := path.Join("/projects", role.Project.ID, "users", user.ID, "roles", role.ID)
	_, err := k.do(ctx, http.MethodPut, path, []int{http.StatusNoContent}, nil, nil)
	return err
}

type rolesListResponse struct {
	Roles []keystone.Role `json:"roles" yaml:"roles"`
}

// GetRoleByName looks for role with provided name in keystone.
func (k *Keystone) GetRoleByName(ctx context.Context, roleName string) (keystone.Role, error) {
	var roles rolesListResponse
	if _, err := k.do(
		ctx, http.MethodGet, fmt.Sprintf("/roles?name=%v", roleName), []int{http.StatusOK}, nil, &roles,
	); err != nil {
		return keystone.Role{}, err
	}

	for _, role := range roles.Roles {
		if role.Name == roleName {
			return role, nil
		}
	}
	return keystone.Role{}, fmt.Errorf("role '%s' does not exist", roleName)
}
