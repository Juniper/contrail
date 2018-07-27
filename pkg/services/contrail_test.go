package services

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestBaseObjectDefaultValuesOnCreateAccessControlList(t *testing.T) {
	tests := []struct {
		name     string
		model    models.AccessControlList
		metadata MetaData
		want     models.AccessControlList
		fails    bool
	}{
		{name: "empty", fails: true},
		{
			name:  "missing parent type - ambigious",
			model: models.AccessControlList{ParentUUID: "parent-uuid"},
			fails: true,
		},
		{
			name: "parent uuid and type provided",
			model: models.AccessControlList{
				UUID:       "foo-uuid",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
			},
			metadata: MetaData{FQName: []string{"default-domain", "default-project", "default-virtual-network"}},
			want: models.AccessControlList{
				UUID:       "foo-uuid",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
				// Default filled fields below
				Name:        "default-access-control-list",
				DisplayName: "default-access-control-list",
				FQName: []string{
					"default-domain", "default-project", "default-virtual-network", "default-access-control-list"},
				Perms2: &models.PermType2{Owner: "default-project"},
			},
		},
		{
			name: "fill default display name",
			model: models.AccessControlList{
				UUID:       "foo-uuid",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
				Name:       "some-name",
				FQName: []string{
					"default-domain", "default-project", "default-virtual-network", "default-access-control-list"},
				Perms2: &models.PermType2{Owner: "default-project"},
			},
			want: models.AccessControlList{
				UUID:       "foo-uuid",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
				Name:       "some-name",
				FQName: []string{
					"default-domain", "default-project", "default-virtual-network", "default-access-control-list"},
				Perms2: &models.PermType2{Owner: "default-project"},
				// Default filled fields below
				DisplayName: "some-name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spy := &serviceSpy{}
			tv, err := models.NewTypeValidatorWithFormat()
			assert.NoError(t, err)

			service := &ContrailService{
				BaseService:    BaseService{next: spy},
				MetadataGetter: (*mockMetadataGetter)(&tt.metadata),
				TypeValidator:  tv,
			}
			_, err = service.CreateAccessControlList(
				common.NoAuth(context.Background()),
				&CreateAccessControlListRequest{&tt.model},
			)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, *spy.acl)
			}
		})
	}
}

type mockMetadataGetter MetaData

func (m *mockMetadataGetter) GetMetaData(_ context.Context, _ string, _ []string) (*MetaData, error) {
	return (*MetaData)(m), nil
}

type serviceSpy struct {
	BaseService
	acl *models.AccessControlList
}

func (s *serviceSpy) CreateAccessControlList(
	ctx context.Context,
	request *CreateAccessControlListRequest,
) (*CreateAccessControlListResponse, error) {
	s.acl = request.AccessControlList
	return nil, nil
}
