package apisrv

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	retryCount = 2
)

// Client represents a client.
// nolint
type Client struct {
	ID         string `yaml:"id"`
	Password   string `yaml:"password"`
	AuthURL    string `yaml:"authurl"`
	Endpoint   string `yaml:"endpoint"`
	Debug      bool   `yaml:"debug"`
	httpClient *http.Client
	AuthToken  string          `yaml:"-"`
	InSecure   bool            `yaml:"insecure"`
	Domain     string          `yaml:"domain"`
	Scope      *keystone.Scope `yaml:"scope"`
}

// Request represents API request to the server
type Request struct {
	Method   string      `yaml:"method"`
	Path     string      `yaml:"path,omitempty"`
	Expected []int       `yaml:"expected,omitempty"`
	Data     interface{} `yaml:"data,omitempty"`
	Output   interface{} `yaml:"output,omitempty"`
}

// NewClient makes api srv client.
func NewClient(endpoint, authURL, id, password, domain string, insecure bool, scope *keystone.Scope) *Client {
	c := &Client{
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
func (c *Client) Init() {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			//Timeout: 5 * time.Second,
		}).Dial,
		//TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.InSecure},
	}
	client := &http.Client{
		Transport: tr,
		//Timeout:   time.Second * 10,
	}
	c.httpClient = client
}

// Login refreshes authentication token
func (c *Client) Login() error {
	if c.AuthURL == "" {
		return nil
	}
	authURL := c.AuthURL + "/auth/tokens"
	authRequest := &keystone.AuthRequest{
		Auth: &keystone.Auth{
			Identity: &keystone.Identity{
				Methods: []string{"password"},
				Password: &keystone.Password{
					User: &keystone.User{
						Name:     c.ID,
						Password: c.Password,
						Domain: &keystone.Domain{
							ID: c.Domain,
						},
					},
				},
			},
			Scope: c.Scope,
		},
	}
	authResponse := &keystone.AuthResponse{}
	dataJSON, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", authURL, bytes.NewBuffer(dataJSON))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint: errcheck
	err = checkStatusCode([]int{201}, resp.StatusCode)
	if err != nil {
		output, _ := httputil.DumpResponse(resp, true) // nolint: gas
		log.Println(string(output))
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return err
	}
	c.AuthToken = resp.Header.Get("X-Subject-Token")
	return nil
}

func checkStatusCode(expected []int, actual int) error {
	for _, expected := range expected {
		if expected == actual {
			return nil
		}
	}
	return errors.Errorf("unexpected return code: expected %v, actual %v", expected, actual)
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

// Do issues an API request.
func (c *Client) Do(method, path string, data interface{}, output interface{}, expected []int) (*http.Response, error) {
	request, err := c.prepareHTTPRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	resp, err := c.doHTTPRequestRetryingOn401(request, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode(expected, resp.StatusCode)
	if err != nil {
		output, _ := httputil.DumpResponse(resp, true) // nolint:  gas
		log.Println(string(output))
		return resp, err
	}
	if method == echo.DELETE {
		return resp, nil
	}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return resp, errors.Wrap(err, "decoding response body failed")
	}
	if c.Debug {
		log.WithFields(log.Fields{
			"response": resp,
			"output":   output,
		}).Debug("API Server response")
	}
	return resp, err
}

func (c *Client) prepareHTTPRequest(method, path string, data interface{}) (*http.Request, error) {
	var request *http.Request
	if data == nil {
		var err error
		request, err = http.NewRequest(method, getURL(c.Endpoint, path), nil)
		if err != nil {
			return nil, errors.Wrap(err, "creating HTTP request failed")
		}
	} else {
		var dataJSON []byte
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "encoding request data failed")
		}

		request, err = http.NewRequest(method, getURL(c.Endpoint, path), bytes.NewBuffer(dataJSON))
		if err != nil {
			return nil, errors.Wrap(err, "creating HTTP request failed")
		}
	}

	request.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		request.Header.Set("X-Auth-Token", c.AuthToken)
	}
	return request, nil
}

func getURL(endpoint, path string) string {
	return endpoint + path
}

func (c *Client) doHTTPRequestRetryingOn401(request *http.Request, data interface{}) (*http.Response, error) {
	if c.Debug {
		log.WithFields(log.Fields{
			"method": request.Method,
			"url":    request.URL,
			"header": request.Header,
			"data":   data,
		}).Debug("Executing API Server request")
	}
	var resp *http.Response
	for i := 0; i < retryCount; i++ {
		var err error
		resp, err = c.httpClient.Do(request)
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

			err = c.Login() // refresh token
			if err != nil {
				return nil, err
			}
		}
	}
	return resp, nil
}

// DoRequest requests based on reqest object.
func (c *Client) DoRequest(request *Request) (*http.Response, error) {
	return c.Do(request.Method, request.Path, request.Data, &request.Output, request.Expected)
}

// Batch execution.
func (c *Client) Batch(requests []*Request) error {
	for i, request := range requests {
		_, err := c.DoRequest(request)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%dth request failed.", i))
		}
	}
	return nil
}
