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
		metadata models.MetaData
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
				UUID:       "4789f49b-a6df-4744-1ecf-60b0958e45e6",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
			},
			metadata: models.MetaData{FQName: []string{"default-domain", "default-project", "default-virtual-network"}},
			want: models.AccessControlList{
				UUID:       "4789f49b-a6df-4744-1ecf-60b0958e45e6",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
				// Default filled fields below
				Name: "default-access-control-list",
				FQName: []string{
					"default-domain", "default-project", "default-virtual-network", "default-access-control-list"},
				Perms2: &models.PermType2{Owner: "default-project"},
				IDPerms: &models.IdPermsType{
					Enable: true,
					UUID: &models.UuidType{
						UUIDMslong: 5154920197859002180,
						UUIDLslong: 2220099452856583654,
					},
				},
			},
		},
		{
			name: "fill default display name",
			model: models.AccessControlList{
				UUID:       "4789f49b-a6df-4744-1ecf-60b0958e45e6",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
				Name:       "some-name",
				FQName: []string{
					"default-domain", "default-project", "default-virtual-network", "default-access-control-list"},
				Perms2: &models.PermType2{Owner: "default-project"},
			},
			want: models.AccessControlList{
				UUID:       "4789f49b-a6df-4744-1ecf-60b0958e45e6",
				ParentUUID: "parent-uuid",
				ParentType: "virtual-network",
				Name:       "some-name",
				FQName: []string{
					"default-domain", "default-project", "default-virtual-network", "default-access-control-list"},
				Perms2: &models.PermType2{Owner: "default-project"},
				// Default filled fields below
				IDPerms: &models.IdPermsType{
					Enable: true,
					UUID: &models.UuidType{
						UUIDMslong: 5154920197859002180,
						UUIDLslong: 2220099452856583654,
					},
				},
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

type mockMetadataGetter models.MetaData

func (m *mockMetadataGetter) GetMetaData(_ context.Context, _ string, _ []string) (*models.MetaData, error) {
	return (*models.MetaData)(m), nil
}

func (m *mockMetadataGetter) ListMetadata(
	ctx context.Context, fqNameUUIDPairs []*models.FQNameUUIDPair,
) ([]*models.MetaData, error) {
	return []*models.MetaData{(*models.MetaData)(m)}, nil
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
