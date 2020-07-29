package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

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

// LoadHTTPConfig creates new config object based on viper configuration.
func LoadHTTPConfig() *HTTPConfig {
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
	return NewHTTP(LoadHTTPConfig())
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

// DoRequest requests based on request object.
func (h *HTTP) DoRequest(ctx context.Context, request *Request) (*http.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return h.Do(ctx, request.Method, request.Path, nil, request.Data, &request.Output, request.Expected)
}

// Request represents API request to the server.
type Request struct {
	Method   string      `yaml:"method"`
	Path     string      `yaml:"path,omitempty"`
	Expected []int       `yaml:"expected,omitempty"`
	Data     interface{} `yaml:"data,omitempty"`
	Output   interface{} `yaml:"output,omitempty"`
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
	return h.Do(ctx, http.MethodPost, "/"+services.RefUpdatePath, nil, data, output, []int{http.StatusOK})
}

// CreateIntPool sends a create int pool request to remote int-pools.
func (h *HTTP) CreateIntPool(ctx context.Context, pool string, start int64, end int64) error {
	_, err := h.Do(
		ctx,
		http.MethodPost,
		"/"+services.IntPoolsPath,
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

	_, err := h.Do(ctx, http.MethodGet, "/"+services.IntPoolPath, q, nil, &output, []int{http.StatusOK})
	return output.Owner, errors.Wrap(err, "error getting int pool owner via HTTP")
}

// DeleteIntPool sends a delete int pool request to remote int-pools.
func (h *HTTP) DeleteIntPool(ctx context.Context, pool string) error {
	_, err := h.Do(
		ctx,
		http.MethodDelete,
		"/"+services.IntPoolsPath,
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
		"/"+services.IntPoolPath,
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
		"/"+services.IntPoolPath,
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
		"/"+services.IntPoolPath,
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
func (h *HTTP) Do(
	ctx context.Context,
	method, path string,
	query url.Values,
	data, output interface{},
	expected []int,
) (*http.Response, error) {
	request, err := h.prepareHTTPRequest(ctx, method, path, data, query)
	if err != nil {
		return nil, err
	}

	resp, err := h.doHTTPRequestRetryingOn401(ctx, request, data)
	if err != nil {
		return resp, httputil.ErrorFromResponse(err, resp)
	}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			logrus.WithError(err).Debug("client.HTTP: Error closing response body")
		}
	}()

	err = httputil.CheckStatusCode(expected, resp.StatusCode)
	if err != nil {
		return resp, httputil.ErrorFromResponse(err, resp)
	}

	d := json.NewDecoder(resp.Body)
	d.UseNumber()
	err = d.Decode(&output)
	switch err {
	default:
		return resp, errors.Wrapf(httputil.ErrorFromResponse(err, resp), "decoding response body failed")
	case io.EOF:
	case nil:
	}
	return resp, nil
}

func (h *HTTP) prepareHTTPRequest(
	ctx context.Context, method, path string, data interface{}, query url.Values,
) (*http.Request, error) {
	var request *http.Request
	if data != nil {
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "encoding request data failed")
		}
		request, err = http.NewRequest(method, h.getURL(path), bytes.NewBuffer(dataJSON))
		if err != nil {
			return nil, errors.Wrap(err, "creating HTTP request failed")
		}
		request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	} else {
		var err error
		request, err = http.NewRequest(method, h.getURL(path), nil)
		if err != nil {
			return nil, errors.Wrap(err, "creating HTTP request failed")
		}
		request = request.WithContext(ctx) // TODO(mblotniak): use http.NewRequestWithContext after go 1.13 upgrade
	}
	if len(query) > 0 {
		request.URL.RawQuery = query.Encode()
	}
	request.Header.Set("Content-Type", "application/json")
	if h.AuthToken != "" {
		request.Header.Set("X-Auth-Token", h.AuthToken)
	}
	return request, nil
}

func (h *HTTP) getURL(path string) string {
	return h.Endpoint + path
}

// nolint: gocyclo
func (h *HTTP) doHTTPRequestRetryingOn401(
	ctx context.Context, request *http.Request, data interface{},
) (*http.Response, error) {
	logrus.WithFields(logrus.Fields{
		"method": request.Method,
		"url":    request.URL,
		"header": request.Header,
		"data":   data,
	}).Debug("Executing API Server request")

	httputil.SetContextHeaders(request)
	var resp *http.Response
	for i := 0; i < retryCount; i++ {
		var err error
		if resp, err = h.httpClient.Do(request); err != nil {
			return resp, errors.Wrap(err, "issuing HTTP request failed")
		}
		if resp.StatusCode != http.StatusUnauthorized || h.ID == "" {
			break
		}
		// If there is no keystone scope setup then the retry won't help.
		if resp.StatusCode == http.StatusUnauthorized && h.Scope == nil {
			return resp, errors.Wrap(err, "no keystone present cannot reauth")
		}
		// token might be expired, refresh token and retry
		// skip refresh token after last retry
		if i < retryCount-1 {
			if err = resp.Body.Close(); err != nil {
				return resp, errors.Wrap(err, "closing response body failed")
			}

			// refresh token and use the new token in request header
			if err := h.Login(ctx); err != nil {
				return resp, err
			}
			if h.AuthToken != "" {
				request.Header.Set("X-Auth-Token", h.AuthToken)
			}
		}
	}
	return resp, nil
}
