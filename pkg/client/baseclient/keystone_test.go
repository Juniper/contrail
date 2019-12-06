package baseclient_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/Juniper/asf/pkg/keystone"
	. "github.com/Juniper/contrail/pkg/client/baseclient"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestKeystoneClient(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		RepoRootPath: "../../..",
	})
	defer s.CloseT(t)

	k := &Keystone{
		URL: viper.GetString("keystone.authurl"),
		HTTPDoer: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: viper.GetBool("keystone.insecure")},
			},
		},
	}

	token, err := k.ObtainToken(context.Background(), integration.AdminUserID, integration.AdminUserPassword, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	p, err := k.GetProject(context.Background(), token, integration.AdminProjectID)
	assert.NoError(t, err)
	assert.Equal(t, integration.AdminProjectID, p.ID)
	assert.Equal(t, integration.AdminProjectName, p.Name)
}

func TestKeystone_CreateUser(t *testing.T) {
	tests := []struct {
		name string

		HTTPDoer doer
		ctx      context.Context

		want    *keystone.User
		wantErr bool
	}{{
		name:     "got empty response from keystone",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte{})},
		wantErr:  true,
	}, {
		name:     "keystone returns sample user",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte(`{"user":{"id":"ff4e51"}}`))},
		want:     &keystone.User{ID: "ff4e51"},
	}, {
		name:     "keystone returns StatusForbidden",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusForbidden, nil)},
		wantErr:  true,
	}, {
		name:     "keystone returns invalid fields",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte(`{"foo":{"bar":"foobar"}}`))},
		want:     &keystone.User{},
	}, {
		name:     "keystone returns bad response",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte(`{"foo":"bar":"foobar"}}`))},
		wantErr:  true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ctx == nil {
				tt.ctx = context.Background()
			}
			k := &Keystone{
				HTTPDoer: tt.HTTPDoer,
			}
			got, err := k.CreateUser(tt.ctx, keystone.User{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Keystone.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keystone.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type mockDoer struct {
	*http.Response
}

func (m *mockDoer) Do(req *http.Request) (*http.Response, error) {
	return m.Response, nil
}

func newResponse(statusCode int, body []byte) *http.Response {
	return &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewReader(body))}
}

func TestKeystone_AddRole(t *testing.T) {
	tests := []struct {
		URL      string
		HTTPDoer doer
		ctx      context.Context
		user     keystone.User
		role     keystone.Role
		name     string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keystone{
				URL:      tt.fields.URL,
				HTTPDoer: tt.fields.HTTPDoer,
			}
			if err := k.AddRole(tt.args.ctx, tt.args.user, tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("Keystone.AddRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeystone_CreateServiceUser(t *testing.T) {
	tests := []struct {
		URL      string
		HTTPDoer doer
		ctx      context.Context
		user     keystone.User
		name     string
		want     *keystone.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keystone{
				URL:      tt.fields.URL,
				HTTPDoer: tt.fields.HTTPDoer,
			}
			got, err := k.CreateServiceUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keystone.CreateServiceUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keystone.CreateServiceUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
