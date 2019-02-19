package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/keystone"
)

// Keystone is a keystone client.
type Keystone struct {
	URL        string
	HTTPClient *http.Client
}

type projectResponse struct {
	Project keystone.Project `json:"project"`
}

// GetProject gets project.
func (k *Keystone) GetProject(ctx context.Context, token string, id string) (*keystone.Project, error) {
	request, err := http.NewRequest(echo.GET, getURL(k.URL, "/projects/"+id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating HTTP request failed")
	}
	request = auth.SetXClusterIDInHeader(ctx, request.WithContext(ctx))
	request.Header.Set("X-Auth-Token", token)
	var output projectResponse

	resp, err := k.HTTPClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "issuing HTTP request failed")
	}
	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{http.StatusOK}, resp.StatusCode); err != nil {
		return nil, errorFromResponse(err, resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, errors.Wrapf(errorFromResponse(err, resp), "decoding response body failed")
	}

	return &output.Project, nil
}

// ObtainToken gets authentication token.
func (k *Keystone) ObtainToken(
	ctx context.Context, id string, password string, scope *keystone.Scope,
) (*http.Response, error) {
	if k.URL == "" {
		return nil, nil
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
		return nil, err
	}

	request, err := http.NewRequest("POST", k.URL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	if err != nil {
		return nil, err
	}
	request = auth.SetXClusterIDInHeader(ctx, request.WithContext(ctx))
	request.Header.Set("Content-Type", "application/json")

	startedAt := time.Now()
	resp, err := k.HTTPClient.Do(request)
	durationInUsec := time.Since(startedAt) / time.Microsecond
	if err != nil {
		return nil, errorFromResponse(err, resp)
	}

	if c := collector.FromContext(ctx); c != nil {
		c.Send(analytics.VncAPILatencyStatsLog(ctx, "VALIDATE", "KEYSTONE", int64(durationInUsec)))
	}

	defer resp.Body.Close() // nolint: errcheck

	if err = checkStatusCode([]int{200, 201}, resp.StatusCode); err != nil {
		return resp, errorFromResponse(err, resp)
	}

	var authResponse keystone.AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return resp, errorFromResponse(err, resp)
	}

	return resp, nil
}
