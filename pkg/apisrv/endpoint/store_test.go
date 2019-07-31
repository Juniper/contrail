package endpoint_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/endpoint"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
)

// TODO(Daniel): improve test coverage

func TestTargetStoreRead(t *testing.T) {
	for _, tt := range []struct {
		name     string
		expected *models.Endpoint
	}{
		{
			name: "empty store",
		},
	} {
		t.Run(tt.name, func(t1 *testing.T) {
			ts := endpoint.NewTargetStore()

			e := ts.Read("id")

			assert.Equal(t, tt.expected, e)
		})
	}
}

func TestTargetStoreReadAll(t *testing.T) {
	fooE, barE, spockE := newEndpoint("foo"), newEndpoint("bar"), newEndpoint("spock")

	for _, tt := range []struct {
		name      string
		scope     string
		expected1 []*endpoint.Endpoint
		expected2 []*endpoint.Endpoint
	}{
		{
			name:  "invalid scope",
			scope: "invalid",
		},
		{
			name:  "private scope",
			scope: endpoint.Private,
			expected1: []*endpoint.Endpoint{
				//{
				//	URL:      fooE.PrivateURL,
				//	Username: fooE.Username,
				//	Password: fooE.Password,
				//},
				//{
				//	URL:      barE.PrivateURL,
				//	Username: barE.Username,
				//	Password: barE.Password,
				//},
				{
					URL:      spockE.PrivateURL,
					Username: spockE.Username,
					Password: spockE.Password,
				},
			},
			expected2: []*endpoint.Endpoint{
				//{
				//	URL:      fooE.PrivateURL,
				//	Username: fooE.Username,
				//	Password: fooE.Password,
				//},
				//{
				//	URL:      barE.PrivateURL,
				//	Username: barE.Username,
				//	Password: barE.Password,
				//},
				{
					URL:      spockE.PrivateURL,
					Username: spockE.Username,
					Password: spockE.Password,
				},
			},
		},
		//{
		//	name:  "public scope",
		//	scope: endpoint.Public,
		//},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ts := endpoint.NewTargetStore()
			ts.Write("foo", fooE)
			ts.Write("bar", barE)
			ts.Write("spock", spockE)

			targets1 := ts.ReadAll(tt.scope)
			targets2 := ts.ReadAll(tt.scope)

			assert.Equal(t, tt.expected1, targets1, "invalid targets1")
			assert.Equal(t, tt.expected2, targets2, "invalid targets2")
		})
	}
}

func newEndpoint(name string) *models.Endpoint {
	return &models.Endpoint{
		PrivateURL: name + "-private",
		PublicURL:  name + "-public",
		Username:   name + "-username",
		Password:   name + "-password",
	}
}
