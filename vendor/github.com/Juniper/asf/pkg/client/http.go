package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Juniper/asf/pkg/httputil"
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	retryCount = 2
)

// HTTPConfig contains HTTP client configuration.
type HTTPConfig struct {
	ID       string          `yaml:"id"`
	Password string          `yaml:"password"`
	Endpoint string          `yaml:"endpoint"`
	AuthURL  string          `yaml:"authurl"`
	Scope    *keystone.Scope `yaml:"scope"`
	Insecure bool            `yaml:"insecure"`
}

// LoadGlobalHTTPConfig creates new config object based on global Viper configuration.
func LoadGlobalHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		ID:       viper.GetString("client.id"),
		Password: viper.GetString("client.password"),
		Endpoint: viper.GetString("client.endpoint"),
		AuthURL:  viper.GetString("keystone.authurl"),
		Scope: keystone.NewScope(
			viper.GetString("client.domain_id"),
			viper.GetString("client.domain_name"),
			viper.GetString("client.project_id"),
			viper.GetString("client.project_name"),
		),
		Insecure: viper.GetBool("insecure"),
	}
}

// SetCredentials sets custom credentials in HTTPConfig.
func (c *HTTPConfig) SetCredentials(username, password string) {
	c.ID = username
	c.Password = password
}

// HTTP represents API Server HTTP client.
type HTTP struct {
	httpClient *http.Client
	Keystone   *keystone.Client

	HTTPConfig `yaml:",inline"`

	AuthToken string
}

// NewHTTP makes API Server HTTP client.
func NewHTTP(c *HTTPConfig) *HTTP {
	hc := &http.Client{Transport: httputil.DefaultTransport(c.Insecure)}
	return &HTTP{
		httpClient: hc,
		Keystone: &keystone.Client{
			HTTPDoer: hc,
			URL:      c.AuthURL,
		},
		HTTPConfig: *c,
	}
}

// NewHTTPFromConfig makes API Server HTTP client with viper config
func NewHTTPFromConfig() *HTTP {
	return NewHTTP(LoadGlobalHTTPConfig())
}

// Login refreshes authentication token.
func (h *HTTP) Login(ctx context.Context) error {
	token, err := h.Keystone.ObtainToken(ctx, h.ID, h.Password, h.Scope)
	if err != nil {
		return err
	}

	h.AuthToken = token
	return nil
}

// Batch execution.
func (h *HTTP) Batch(ctx context.Context, requests []*Request) error {
	for i, request := range requests {
		_, err := h.DoRequest(ctx, request)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%dth request failed.", i))
		}
	}
	return nil
}

// Create send a create API request.
func (h *HTTP) Create(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	return h.Do(ctx, http.MethodPost, path, nil, data, output, []int{http.StatusOK})
}

// Read send a get API request.
func (h *HTTP) Read(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	return h.Do(ctx, http.MethodGet, path, nil, nil, output, []int{http.StatusOK})
}

// ReadWithQuery send a get API request with a query.
func (h *HTTP) ReadWithQuery(
	ctx context.Context, path string, query url.Values, output interface{},
) (*http.Response, error) {
	return h.Do(ctx, http.MethodGet, path, query, nil, output, []int{http.StatusOK})
}

// Update send an update API request.
func (h *HTTP) Update(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	return h.Do(ctx, http.MethodPut, path, nil, data, output, []int{http.StatusOK})
}

// Delete send a delete API request.
func (h *HTTP) Delete(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	return h.Do(ctx, http.MethodDelete, path, nil, nil, output, []int{http.StatusOK})
}

// EnsureDeleted send a delete API request.
func (h *HTTP) EnsureDeleted(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	return h.Do(ctx, http.MethodDelete, path, nil, nil, output, []int{http.StatusOK, http.StatusNotFound})
}

// RefUpdate sends a create/update API request
func (h *HTTP) RefUpdate(ctx context.Context, data interface{}, output interface{}) (*http.Response, error) {
	return h.Do(ctx, http.MethodPost, services.RefUpdatePath, nil, data, output, []int{http.StatusOK})
}

// CreateIntPool sends a create int pool request to remote int-pools.
func (h *HTTP) CreateIntPool(ctx context.Context, pool string, start int64, end int64) error {
	_, err := h.Do(
		ctx,
		http.MethodPost,
		services.IntPoolsPath,
		nil,
		&services.CreateIntPoolRequest{
			Pool:  pool,
			Start: start,
			End:   end,
		},
		&struct{}{},
		[]int{http.StatusOK},
	)
	return errors.Wrap(err, "error creating int pool in int-pools via HTTP")
}

// GetIntOwner sends a get int pool owner request to remote int-owner.
func (h *HTTP) GetIntOwner(ctx context.Context, pool string, value int64) (string, error) {
	q := make(url.Values)
	q.Set("pool", pool)
	q.Set("value", strconv.FormatInt(value, 10))
	var output struct {
		Owner string `json:"owner"`
	}

	_, err := h.Do(ctx, http.MethodGet, services.IntPoolPath, q, nil, &output, []int{http.StatusOK})
	return output.Owner, errors.Wrap(err, "error getting int pool owner via HTTP")
}

// DeleteIntPool sends a delete int pool request to remote int-pools.
func (h *HTTP) DeleteIntPool(ctx context.Context, pool string) error {
	_, err := h.Do(
		ctx,
		http.MethodDelete,
		services.IntPoolsPath,
		nil,
		&services.DeleteIntPoolRequest{
			Pool: pool,
		},
		&struct{}{},
		[]int{http.StatusOK},
	)
	return errors.Wrap(err, "error deleting int pool in int-pools via HTTP")
}

// AllocateInt sends an allocate int request to remote int-pool.
func (h *HTTP) AllocateInt(ctx context.Context, pool, owner string) (int64, error) {
	var output struct {
		Value int64 `json:"value"`
	}
	_, err := h.Do(
		ctx,
		http.MethodPost,
		services.IntPoolPath,
		nil,
		&services.IntPoolAllocationBody{
			Pool:  pool,
			Owner: owner,
		},
		&output,
		[]int{http.StatusOK},
	)
	return output.Value, errors.Wrap(err, "error allocating int in int-pool via HTTP")
}

// SetInt sends a set int request to remote int-pool.
func (h *HTTP) SetInt(ctx context.Context, pool string, value int64, owner string) error {
	_, err := h.Do(
		ctx,
		http.MethodPost,
		services.IntPoolPath,
		nil,
		&services.IntPoolAllocationBody{
			Pool:  pool,
			Value: &value,
			Owner: owner,
		},
		&struct{}{},
		[]int{http.StatusOK},
	)
	return errors.Wrap(err, "error setting int in int-pool via HTTP")
}

// DeallocateInt sends a deallocate int request to remote int-pool.
func (h *HTTP) DeallocateInt(ctx context.Context, pool string, value int64) error {
	_, err := h.Do(
		ctx,
		http.MethodDelete,
		services.IntPoolPath,
		nil,
		&services.IntPoolAllocationBody{
			Pool:  pool,
			Value: &value,
		},
		&struct{}{},
		[]int{http.StatusOK},
	)
	return errors.Wrap(err, "error deallocating int in int-pool via HTTP")
}

// Do issues an API request.
// Deprecated: use DoRequest() instead.
func (h *HTTP) Do(
	ctx context.Context,
	method, path string,
	query url.Values,
	data, output interface{},
	expected []int,
) (*http.Response, error) {
	return h.DoRequest(ctx, &Request{
		Method:           method,
		Path:             path,
		Query:            query,
		RequestBody:      data,
		ResponseBody:     output,
		ExpectedStatuses: expected,
	})
}

// DoRequest preforms the request based on given request object.
// request.RequestBodyJSON or request.RequestBody is used as a request body - RequestBodyJSON variant
// has higher priority.
// TODO(dfurman): refactor the function to implement Doer interface: func(http.Request) (http.Response, error)
// TODO(dfurman): use the builder pattern: h.ExpectStatuses([]int{200}).DecodeResponseTo(&response).Do(http.Request)
// TODO(dfurman): improve the test coverage
func (h *HTTP) DoRequest(ctx context.Context, r *Request) (*http.Response, error) {
	if r == nil {
		return nil, errors.Errorf("request cannot be nil")
	}

	resp, err := h.doHTTPRequestRetryingOn401(ctx, r)
	if err != nil {
		return resp, httputil.ErrorFromResponse(err, resp)
	}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			// TODO(dfurman): append cErr to err
			logrus.WithError(err).Debug("client.HTTP: Error closing response body")
		}
	}()

	err = httputil.CheckStatusCode(r.ExpectedStatuses, resp.StatusCode)
	if err != nil {
		return resp, httputil.ErrorFromResponse(err, resp)
	}

	d := json.NewDecoder(resp.Body)
	d.UseNumber()
	err = d.Decode(&r.ResponseBody)
	switch err {
	default:
		return resp, errors.Wrapf(httputil.ErrorFromResponse(err, resp), "decoding response body failed")
	case io.EOF:
	case nil:
	}
	return resp, nil
}

func (h *HTTP) doHTTPRequestRetryingOn401(ctx context.Context, r *Request) (*http.Response, error) {
	var resp *http.Response
	for i := 0; i < retryCount; i++ {
		httpR, err := r.newHTTPRequest(ctx, h.Endpoint, h.AuthToken)
		if err != nil {
			return nil, err
		}
		logrus.WithFields(logrus.Fields{
			"attempt": i + 1,
			"request": httpR,
		}).Debug("Doing HTTP request")

		if resp, err = h.httpClient.Do(httpR); err != nil {
			return resp, errors.Wrap(err, "issuing HTTP request failed")
		}
		if resp.StatusCode != http.StatusUnauthorized || h.ID == "" {
			break
		}
		// If there is no keystone scope setup then the retry won't help.
		if resp.StatusCode == http.StatusUnauthorized && h.Scope == nil {
			return resp, errors.Wrap(err, "no keystone present - cannot re-authenticate")
		}
		// token might be expired, refresh token and retry
		// skip refresh token after last retry
		if i < retryCount-1 {
			if err = resp.Body.Close(); err != nil {
				return resp, errors.Wrap(err, "closing response body failed")
			}

			// refresh token and use the new token in request newHeader
			if err := h.Login(ctx); err != nil {
				return resp, err
			}
			if h.AuthToken != "" {
				httpR.Header.Set("X-Auth-Token", h.AuthToken)
			}
		}
	}
	return resp, nil
}

// Request holds HTTP request data.
type Request struct {
	Method           string      `yaml:"method,omitempty"`
	Path             string      `yaml:"path,omitempty"`
	Query            url.Values  `yaml:"query,omitempty"`
	RequestBody      interface{} `yaml:"data,omitempty"`
	RequestBodyJSON  string      `yaml:"request_body_json,omitempty"`
	ResponseBody     interface{} `yaml:"output,omitempty"`
	ExpectedStatuses []int       `yaml:"expected,omitempty"`
}

func (r *Request) newHTTPRequest(ctx context.Context, endpoint, authToken string) (request *http.Request, err error) {
	body, err := r.resolveBody()
	if err != nil {
		return nil, err
	}
	request, err = http.NewRequestWithContext(ctx, r.Method, endpoint+r.Path, body)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}

	if len(r.Query) > 0 {
		request.URL.RawQuery = r.Query.Encode()
	}

	request.Header = newHeader(authToken)
	httputil.SetContextHeaders(request)
	return request, nil
}

func (r *Request) resolveBody() (body io.Reader, err error) {
	if r.RequestBodyJSON != "" {
		body = strings.NewReader(r.RequestBodyJSON)
	} else if r.RequestBody != nil {
		j, err := json.Marshal(r.RequestBody)
		if err != nil {
			return nil, errors.Wrap(err, "encode request body")
		}
		body = strings.NewReader(string(j))
	}
	return body, err
}

func newHeader(authToken string) http.Header {
	header := http.Header{
		"Content-Type": []string{"application/json"},
	}
	if authToken != "" {
		header.Set("X-Auth-Token", authToken)
	}
	return header
}
