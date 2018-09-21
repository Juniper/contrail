package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
)

const (
	retryCount = 2
)

// HTTP represents API Server HTTP client.
type HTTP struct {
	httpClient *http.Client

	ID        string          `yaml:"id"`
	Password  string          `yaml:"password"`
	AuthURL   string          `yaml:"authurl"`
	Endpoint  string          `yaml:"endpoint"`
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
		ID:       id,
		Password: password,
		AuthURL:  authURL,
		Endpoint: endpoint,
		Scope:    scope,
		InSecure: insecure,
	}
	c.Init()
	return c
}

//Init is used to initialize a client.
func (h *HTTP) Init() {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			//Timeout: 5 * time.Second,
		}).Dial,
		//TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: h.InSecure},
	}
	client := &http.Client{
		Transport: tr,
		//Timeout:   time.Second * 10,
	}
	h.httpClient = client
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

	resp, err := h.httpClient.Do(request)
	if err != nil {
		return errorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode([]int{200}, resp.StatusCode)
	if err != nil {
		return errorFromResponse(err, resp)
	}

	var authResponse keystone.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return errorFromResponse(err, resp)
	}

	h.AuthToken = resp.Header.Get("X-Subject-Token")
	return nil
}

// Create send a create API request.
func (h *HTTP) Create(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.POST, path, nil, data, output, expected)
}

// Read send a get API request.
func (h *HTTP) Read(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.GET, path, nil, nil, output, expected)
}

// ReadWithQuery send a get API request with a query.
func (h *HTTP) ReadWithQuery(
	ctx context.Context, path string, query url.Values, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.GET, path, query, nil, output, expected)
}

// Update send an update API request.
func (h *HTTP) Update(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.PUT, path, nil, data, output, expected)
}

// Delete send a delete API request.
func (h *HTTP) Delete(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.DELETE, path, nil, nil, output, expected)
}

// RefUpdate sends a create/update API request/
func (h *HTTP) RefUpdate(ctx context.Context, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.POST, "/ref-update", nil, data, output, expected)
}

// EnsureDeleted send a delete API request.
func (h *HTTP) EnsureDeleted(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK, http.StatusNotFound}
	return h.Do(ctx, echo.DELETE, path, nil, nil, output, expected)
}

// AllocateInt sends an allocate int request to remote int-pool.
func (h *HTTP) AllocateInt(ctx context.Context, pool string) (int64, error) {
	var output struct {
		Value int64 `json:"value"`
	}
	expected := []int{http.StatusOK}
	_, err := h.Do(ctx, echo.POST, path.Join("/int-pool", pool), nil, nil, &output, expected)
	return output.Value, errors.Wrap(err, "error allocating int in int-pool via HTTP")
}

// SetInt sends a set int request to remote int-pool.
func (h *HTTP) SetInt(ctx context.Context, pool string, value int64) error {
	var output struct{}
	expected := []int{http.StatusOK}
	_, err := h.Do(ctx, echo.POST, path.Join("/int-pool", pool, fmt.Sprint(value)), nil, nil, &output, expected)
	return errors.Wrap(err, "error setting int in int-pool via HTTP")
}

// DeallocateInt sends a deallocate int request to remote int pool.
func (h *HTTP) DeallocateInt(ctx context.Context, pool string, value int64) error {
	var output struct{}
	expected := []int{http.StatusOK}
	_, err := h.Do(ctx, echo.DELETE, path.Join("/int-pool", pool, fmt.Sprint(value)), nil, nil, &output, expected)
	return errors.Wrap(err, "error deallocating int in int-pool via HTTP")
}

// Do issues an API request.
func (h *HTTP) Do(ctx context.Context,
	method, path string, query url.Values, data interface{}, output interface{}, expected []int) (*http.Response, error) {
	request, err := h.prepareHTTPRequest(method, path, data, query)
	if err != nil {
		return nil, err
	}

	resp, err := h.doHTTPRequestRetryingOn401(ctx, request, data)
	if err != nil {
		return nil, errorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode(expected, resp.StatusCode)
	if err != nil {
		return resp, errorFromResponse(err, resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err == io.EOF {
		return resp, nil
	} else if err != nil {
		return resp, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	return resp, nil
}

func (h *HTTP) prepareHTTPRequest(method, path string, data interface{}, query url.Values) (*http.Request, error) {
	var request *http.Request
	if data == nil {
		var err error
		request, err = http.NewRequest(method, getURL(h.Endpoint, path), nil)
		if err != nil {
			return nil, errors.Wrap(err, "creating HTTP request failed")
		}
	} else {
		var dataJSON []byte
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "encoding request data failed")
		}

		request, err = http.NewRequest(method, getURL(h.Endpoint, path), bytes.NewBuffer(dataJSON))
		if err != nil {
			return nil, errors.Wrap(err, "creating HTTP request failed")
		}
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

func getURL(endpoint, path string) string {
	return endpoint + path
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
		resp, err = h.httpClient.Do(request)
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

func checkStatusCode(expected []int, actual int) error {
	for _, e := range expected {
		if e == actual {
			return nil
		}
	}
	return errors.Errorf("unexpected return code: expected %v, actual %v", expected, actual)
}

// DoRequest requests based on request object.
func (h *HTTP) DoRequest(ctx context.Context, request *Request) (*http.Response, error) {
	return h.Do(ctx, request.Method, request.Path, nil, request.Data, &request.Output, request.Expected)
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

func errorFromResponse(e error, r *http.Response) error {
	if r == nil {
		return errors.Wrap(e, "response is nil, error")
	}
	b, err := httputil.DumpResponse(r, true)
	if err != nil {
		return errors.Wrapf(e, "response: failed to dump (%s)", err)
	}
	return errors.Wrapf(e, "response:\n%v", string(b))
}
