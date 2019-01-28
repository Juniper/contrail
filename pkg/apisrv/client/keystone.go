package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"

	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/siddontang/go/log"
)

// Keystone is a keystone client
type Keystone struct {
	URL        string `yaml:"authurl"`
	httpClient *http.Client
}

type projectResponse struct {
	Project keystone.Project `json:"project"`
}

// NewKeystone creates a new keystone client
func NewKeystone(url string, inSecure bool, tr *http.Transport) *Keystone {
	k := &Keystone{
		httpClient: &http.Client{
			Transport: &http.Transport{
				Dial:            (&net.Dialer{}).Dial,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: inSecure},
			},
		},
		URL: url,
	}

	return k
}

// GetProject gets projects
func (k *Keystone) GetProject(ctx context.Context, token string, id string) (*keystone.Project, error) {
	log.Error("GetProject")
	request, err := http.NewRequest(echo.GET, getURL(k.URL, "/project/"+id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}

	request.Header.Set("X-Auth-Token", token)
	var output projectResponse

	resp, err := k.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}

	err = checkStatusCode([]int{http.StatusOK}, resp.StatusCode)
	if err != nil {
		return nil, errorFromResponse(err, resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	defer resp.Body.Close()
	log.Errorf("GetProject %+v", output.Project)
	return &output.Project, nil
}

// ObtainToken gets authentication token.
func (k *Keystone) ObtainToken(
	ctx context.Context, id string, password string, scope *keystone.Scope,
) (string, error) {
	if k.URL == "" {
		return "", nil
	}

	dataJSON, err := json.Marshal(&keystone.AuthRequest{
		Auth: &keystone.Auth{
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
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	request = request.WithContext(ctx)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := k.httpClient.Do(request)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	err = checkStatusCode([]int{200, 201}, resp.StatusCode)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}

	var authResponse keystone.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return "", errorFromResponse(err, resp)
	}

	return resp.Header.Get("X-Subject-Token"), nil
}
