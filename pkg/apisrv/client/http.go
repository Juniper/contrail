package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
)

const (
	retryCount = 2
)

// HTTP represents API Server HTTP client.
type HTTP struct {
	*common.BaseHTTP

	ID        string          `yaml:"id"`
	Password  string          `yaml:"password"`
	AuthURL   string          `yaml:"authurl"`
	AuthToken string          `yaml:"-"`
	InSecure  bool            `yaml:"insecure"`
	Debug     bool            `yaml:"debug"`
	Scope     *keystone.Scope `yaml:"scope"`
}

// Request represents API request to the server.
type Request struct {
	Method   string      `yaml:"method"`
	Path     string      `yaml:"path,omitempty"`
	Expected []int       `yaml:"expected,omitempty"`
	Data     interface{} `yaml:"data,omitempty"`
	Output   interface{} `yaml:"output,omitempty"`
}

// GetKeystoneScope returns the project/domain scope
func GetKeystoneScope(domainID, domainName, projectID, projectName string) *keystone.Scope {
	scope := &keystone.Scope{
		Project: &keystone.Project{
			Domain: &keystone.Domain{},
		},
	}
	if domainID != "" {
		scope.Project.Domain.ID = domainID
	} else if domainName != "" {
		scope.Project.Domain.Name = domainName
	}
	if projectID != "" {
		scope.Project.ID = projectID
	} else if projectName != "" {
		scope.Project.Name = projectName
	}
	return scope
}

// NewHTTP makes API Server HTTP client.
func NewHTTP(endpoint, authURL, id, password string, insecure bool, scope *keystone.Scope) *HTTP {
	c := &HTTP{
		BaseHTTP: &common.BaseHTTP{
			Endpoint: endpoint,
			InSecure: insecure,
		},
		ID:       id,
		Password: password,
		AuthURL:  authURL,
		Scope:    scope,
	}
	c.Preparer = c.prepareHTTPRequest
	c.Requester = c.doHTTPRequestRetryingOn401
	c.Init()
	return c
}

// Login refreshes authentication token.
func (h *HTTP) Login(ctx context.Context) error {
	if h.AuthURL == "" {
		return nil
	}

	var domain *keystone.Domain
	if h.Scope.Domain != nil {
		domain = h.Scope.Domain
	} else if h.Scope.Project != nil {
		domain = h.Scope.Project.Domain
	}
	dataJSON, err := json.Marshal(&keystone.AuthRequest{
		Auth: &keystone.Auth{
			Identity: &keystone.Identity{
				Methods: []string{"password"},
				Password: &keystone.Password{
					User: &keystone.User{
						Name:     h.ID,
						Password: h.Password,
						Domain:   domain,
					},
				},
			},
			Scope: h.Scope,
		},
	})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", h.AuthURL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	request = request.WithContext(ctx)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := h.HttpClient.Do(request)
	if err != nil {
		return common.ErrorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	err = common.CheckStatusCode([]int{200}, resp.StatusCode)
	if err != nil {
		return common.ErrorFromResponse(err, resp)
	}

	var authResponse keystone.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return common.ErrorFromResponse(err, resp)
	}

	h.AuthToken = resp.Header.Get("X-Subject-Token")
	return nil
}
func (h *HTTP) prepareHTTPRequest(method, path string, data interface{}, query url.Values) (*http.Request, error) {
	request, err := h.PrepareHTTPRequest(method, path, data, query)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	if h.AuthToken != "" {
		request.Header.Set("X-Auth-Token", h.AuthToken)
	}
	return request, nil
}

func (h *HTTP) doHTTPRequestRetryingOn401(
	ctx context.Context,
	request *http.Request, data interface{}) (*http.Response, error) {
	if h.Debug {
		log.WithFields(log.Fields{
			"method": request.Method,
			"url":    request.URL,
			"header": request.Header,
			"data":   data,
		}).Debug("Executing API Server request")
	}
	request = request.WithContext(ctx)
	var resp *http.Response
	for i := 0; i < retryCount; i++ {
		var err error
		resp, err = h.HttpClient.Do(request)
		if err != nil {
			return nil, errors.Wrap(err, "issuing HTTP request failed")
		}
		if resp.StatusCode != 401 {
			break
		}
		// token might be expired, refresh token and retry
		// skip refresh token after last retry
		if i < retryCount-1 {
			err = resp.Body.Close()
			if err != nil {
				return nil, errors.Wrap(err, "closing response body failed")
			}

			// refresh token and use the new token in request header
			err = h.Login(ctx)
			if err != nil {
				return nil, err
			}
			if h.AuthToken != "" {
				request.Header.Set("X-Auth-Token", h.AuthToken)
			}
		}
	}
	return resp, nil
}
