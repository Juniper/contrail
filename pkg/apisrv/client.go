package apisrv

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Client represents a client.
type Client struct {
	Endpoint   string
	httpClient *http.Client
}

// Request represents API request to the server
type Request struct {
	Method   string      `yaml:"method"`
	Path     string      `yaml:"path"`
	Expected []int       `yaml:"expected"`
	Data     interface{} `yaml:"data"`
	Output   interface{}
}

// NewClient makes api srv client
func NewClient(endpoint string) *Client {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // nolint: gas
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	return &Client{
		Endpoint:   endpoint,
		httpClient: client,
	}
}

// Login refreshes authentication token
func (c *Client) Login() error {
	return nil
}

func checkStatusCode(expected []int, actual int) error {
	for _, expected := range expected {
		if expected == actual {
			return nil
		}
	}
	return fmt.Errorf("Unexpeced return code expected %v, actual %d", expected, actual)
}

// Create send a create API request.
func (c *Client) Create(path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusCreated}
	return c.Do(echo.POST, path, data, output, expected)
}

// Read send a create API request.
func (c *Client) Read(path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return c.Do(echo.GET, path, nil, output, expected)
}

// Update send an update API request.
func (c *Client) Update(path string, data interface{}, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusOK}
	return c.Do(echo.PUT, path, data, output, expected)
}

// Delete send a delete API request.
func (c *Client) Delete(path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusNoContent}
	return c.Do(echo.DELETE, path, nil, output, expected)
}

// EnsureDeleted send a delete API request.
func (c *Client) EnsureDeleted(path string, output interface{}) (*http.Response, error) {
	expected := []int{http.StatusNoContent, http.StatusNotFound}
	return c.Do(echo.DELETE, path, nil, output, expected)
}

// Do issue a API request.
func (c *Client) Do(method, path string, data interface{}, output interface{}, expected []int) (*http.Response, error) {
	var request *http.Request
	var err error
	endpoint := c.Endpoint + path
	if data == nil {
		request, err = http.NewRequest(method, endpoint, nil)
	} else {
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest(method, endpoint, bytes.NewBuffer(dataJSON))
	}
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = checkStatusCode(expected, resp.StatusCode)
	if err != nil {
		output, _ := httputil.DumpResponse(resp, true)
		log.Println(string(output))
		return resp, err
	}
	if method == echo.DELETE {
		return resp, err
	}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// DoRequest requests based on reqest object.
func (c *Client) DoRequest(request *Request) (*http.Response, error) {
	return c.Do(request.Method, request.Path, request.Data, &request.Output, request.Expected)
}

// Batch execution
func (c *Client) Batch(requests []*Request) error {
	for i, request := range requests {
		_, err := c.DoRequest(request)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%dth request failed.", i))
		}
	}
	return nil
}
