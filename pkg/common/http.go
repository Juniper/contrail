package common

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

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	retryCount = 2
)

// BaseHTTP represents API Server HTTP client.
type BaseHTTP struct {
	HttpClient *http.Client

	Preparer func(
		method, path string,
		data interface{},
		query url.Values) (*http.Request, error)
	Requester func(
		ctx context.Context,
		request *http.Request,
		data interface{}) (*http.Response, error)

	Endpoint string `yaml:"endpoint"`
	InSecure bool   `yaml:"insecure"`
	Debug    bool   `yaml:"debug"`
}

// Request represents API request to the server.
type Request struct {
	Method   string      `yaml:"method"`
	Path     string      `yaml:"path,omitempty"`
	Expected []int       `yaml:"expected,omitempty"`
	Data     interface{} `yaml:"data,omitempty"`
	Output   interface{} `yaml:"output,omitempty"`
}

// NewBaseHTTP makes HTTP client.
func NewBaseHTTP(endpoint string, insecure bool) *BaseHTTP {
	c := &BaseHTTP{
		Endpoint: endpoint,
		InSecure: insecure,
		Debug:    viper.GetBool("server.log_api"),
	}
	c.Preparer = c.PrepareHTTPRequest
	c.Requester = c.DoHTTPRequest
	c.Init()
	return c
}

//Init is used to initialize a client.
func (h *BaseHTTP) Init() {
	if h.getProtocol() == "https" {
		tr := &http.Transport{
			Dial: (&net.Dialer{
				//Timeout: 5 * time.Second,
			}).Dial,
			//TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: h.InSecure},
		}
		h.HttpClient = &http.Client{
			Transport: tr,
			//Timeout:   time.Second * 10,
		}
	} else {
		h.HttpClient = &http.Client{}
	}
}

// Create send a create API request.
func (h *BaseHTTP) Create(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.POST, path, nil, data, output, expected)
}

// Read send a get API request.
func (h *BaseHTTP) Read(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.GET, path, nil, nil, output, expected)
}

// ReadWithQuery send a get API request with a query.
func (h *BaseHTTP) ReadWithQuery(
	ctx context.Context, path string, query url.Values, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.GET, path, query, nil, output, expected)
}

// Update send an update API request.
func (h *BaseHTTP) Update(ctx context.Context, path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.PUT, path, nil, data, output, expected)
}

// Delete send a delete API request.
func (h *BaseHTTP) Delete(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.DELETE, path, nil, nil, output, expected)
}

// RefUpdate sends a create/update API request/
func (h *BaseHTTP) RefUpdate(ctx context.Context, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return h.Do(ctx, echo.POST, "/ref-update", nil, data, output, expected)
}

// EnsureDeleted send a delete API request.
func (h *BaseHTTP) EnsureDeleted(ctx context.Context, path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK, http.StatusNotFound}
	return h.Do(ctx, echo.DELETE, path, nil, nil, output, expected)
}

// Do issues an API request.
func (h *BaseHTTP) Do(ctx context.Context,
	method, path string, query url.Values, data interface{}, output interface{}, expected []int) (*http.Response, error) {
	request, err := h.Preparer(method, path, data, query)
	if err != nil {
		return nil, err
	}

	resp, err := h.Requester(ctx, request, data)
	if err != nil {
		return nil, ErrorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	err = CheckStatusCode(expected, resp.StatusCode)
	if err != nil {
		return resp, ErrorFromResponse(err, resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err == io.EOF {
		return resp, nil
	} else if err != nil {
		return resp, errors.Wrapf(ErrorFromResponse(err, resp), "decoding response body failed")
	}

	return resp, nil
}

// PrepareHTTPRequest frames the http request
func (h *BaseHTTP) PrepareHTTPRequest(method, path string, data interface{}, query url.Values) (*http.Request, error) {
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

	return request, nil
}

func getURL(endpoint, path string) string {
	return endpoint + path
}

func (h *BaseHTTP) getProtocol() string {
	u, _ := url.Parse(h.Endpoint)
	return u.Scheme
}

// DoHTTPRequest sends http requst to the endpoint
func (h *BaseHTTP) DoHTTPRequest(
	ctx context.Context,
	request *http.Request, data interface{}) (*http.Response, error) {
	if h.Debug {
		log.WithFields(log.Fields{
			"method": request.Method,
			"url":    request.URL,
			"header": request.Header,
			"data":   data,
		}).Debug("Executing request")
	}
	request = request.WithContext(ctx)
	resp, err := h.HttpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}
	return resp, nil
}

// CheckStatusCode compares http status code with expected code.
func CheckStatusCode(expected []int, actual int) error {
	for _, e := range expected {
		if e == actual {
			return nil
		}
	}
	return errors.Errorf("unexpected return code: expected %v, actual %v", expected, actual)
}

// DoRequest requests based on request object.
func (h *BaseHTTP) DoRequest(ctx context.Context, request *Request) (*http.Response, error) {
	return h.Do(ctx, request.Method, request.Path, nil, request.Data, &request.Output, request.Expected)
}

// Batch execution.
func (h *BaseHTTP) Batch(ctx context.Context, requests []*Request) error {
	for i, request := range requests {
		_, err := h.DoRequest(ctx, request)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%dth request failed.", i))
		}
	}
	return nil
}

// ErrorFromResponse fromats error from the http response
func ErrorFromResponse(e error, r *http.Response) error {
	if r == nil {
		return errors.Wrapf(e, "response: nil")
	}
	b, err := httputil.DumpResponse(r, true)
	if err != nil {
		errors.Wrap(e, "response: failed to dump")
	}
	return errors.Wrapf(e, "response:\n%v", string(b))
}
