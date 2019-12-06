package keystone

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestKeystone_CreateUser(t *testing.T) {
	tests := []struct {
		name     string
		HTTPDoer doer
		ctx      context.Context
		want     User
		wantErr  bool
	}{{
		name:     "keystone returns sample user",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte(`{"user":{"id":"ff4e51"}}`))},
		want:     User{ID: "ff4e51"},
	}, {
		name:     "keystone returns invalid fields",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte(`{"foo":{"bar":"foobar"}}`))},
	}, {
		name:     "keystone returns bad response",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte(`{"foo":"bar":"foobar"}}`))},
		wantErr:  true,
	}, {
		name:     "got empty response from keystone",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusCreated, []byte{})},
		wantErr:  true,
	}, {
		name:     "keystone returns StatusForbidden",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusForbidden, nil)},
		wantErr:  true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ctx == nil {
				tt.ctx = context.Background()
			}
			k := &Client{
				HTTPDoer: tt.HTTPDoer,
			}
			got, err := k.CreateUser(tt.ctx, User{})
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

func TestKeystone_AssignProjectRoleOnUser(t *testing.T) {
	tests := []struct {
		HTTPDoer doer
		ctx      context.Context
		name     string
		wantErr  bool
	}{{
		name:     "keystone returns StatusBadRequest",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusBadRequest, nil)},
		wantErr:  true,
	}, {
		name:     "keystone returns StatusForbidden",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusForbidden, nil)},
		wantErr:  true,
	}, {
		name:     "keystone returns StatusNoContent",
		HTTPDoer: &mockDoer{Response: newResponse(http.StatusNoContent, nil)},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ctx == nil {
				tt.ctx = context.Background()
			}
			k := &Client{
				HTTPDoer: tt.HTTPDoer,
			}
			if err := k.AssignProjectRoleOnUser(tt.ctx, User{ID: ""}, Role{ID: "", Project: &Project{ID: ""}}); (err != nil) != tt.wantErr {
				t.Errorf("Keystone.AssignProjectRoleOnUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
