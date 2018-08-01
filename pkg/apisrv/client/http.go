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
	Domain    string          `yaml:"domain"`
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

// NewHTTP makes API Server HTTP client.
func NewHTTP(endpoint, authURL, id, password, domain string, insecure bool, scope *keystone.Scope) *HTTP {
	c := &HTTP{
		ID:       id,
		Password: password,
		AuthURL:  authURL,
		Endpoint: endpoint,
		Scope:    scope,
		Domain:   domain,
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

	dataJSON, err := json.Marshal(&keystone.AuthRequest{
		Auth: &keystone.Auth{
			Identity: &keystone.Identity{
				Methods: []string{"password"},
				Password: &keystone.Password{
					User: &keystone.User{
						Name:     h.ID,
						Password: h.Password,
						Domain: &keystone.Domain{
							ID: h.Domain,
						},
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
		logErrorAndResponse(err, resp)
		return err
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode([]int{200}, resp.StatusCode)
	if err != nil {
		logErrorAndResponse(err, resp)
		return err
	}

	var authResponse keystone.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		logErrorAndResponse(err, resp)
		return err
	}

	h.AuthToken = resp.Header.Get("X-Subject-Token")
	return nil
}

// Create send a create API request.
func (h *HTTP) Create(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.POST, path, data, output, expected)
}

// Read send a get API request.
func (h *HTTP) Read(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.GET, path, nil, output, expected)
}

// Update send an update API request.
func (h *HTTP) Update(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.PUT, path, data, output, expected)
}

// Delete send a delete API request.
func (h *HTTP) Delete(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.DELETE, path, nil, output, expected)
}

// RefUpdate sends a create/update API request/
func (h *HTTP) RefUpdate(ctx context.Context, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.POST, "/ref-update", data, output, expected)
}

// EnsureDeleted send a delete API request.
func (h *HTTP) EnsureDeleted(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK, http.StatusNotFound}
	return h.Do(ctx, echo.DELETE, path, nil, output, expected)
}

// Do issues an API request.
func (h *HTTP) Do(ctx context.Context,
	method, path string, data interface{}, output interface{}, expected []int) (*http.Response, error) {
	request, err := h.prepareHTTPRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	resp, err := h.doHTTPRequestRetryingOn401(ctx, request, data)
	if err != nil {
		logErrorAndResponse(err, resp)
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode(expected, resp.StatusCode)
	if err != nil {
		logErrorAndResponse(err, resp)
		return resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err == io.EOF {
		return resp, nil
	} else if err != nil {
		logErrorAndResponse(err, resp)
		return resp, errors.Wrap(err, "decoding response body failed")
	}

	if h.Debug {
		log.WithFields(log.Fields{
			"response": resp,
			"output":   output,
		}).Debug("API Server response")
	}
	return resp, nil
}

func (h *HTTP) prepareHTTPRequest(method, path string, data interface{}) (*http.Request, error) {
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

func logErrorAndResponse(err error, response *http.Response) {
	if err == nil {
		return
	}

	if response == nil {
		log.WithError(err).WithField("response", "nil response").Info("Request failed")
		return
	}

	r, dErr := httputil.DumpResponse(response, true)
	if dErr != nil {
		log.WithError(err).WithField("response", "error dumping response").Info("Request failed")
	} else {
		log.WithError(err).WithField("response", string(r)).Info("Request failed")
	}
}

// DoRequest requests based on request object.
func (h *HTTP) DoRequest(ctx context.Context, request *Request) (*http.Response, error) {
	return h.Do(ctx, request.Method, request.Path, request.Data, &request.Output, request.Expected)
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
