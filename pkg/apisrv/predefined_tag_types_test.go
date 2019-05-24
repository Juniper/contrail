package apisrv_test

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
)

const (
	tagTypeTypeName = "tag-type"
)

func TestPredefinedTagTypes(t *testing.T) {
	c := integration.NewTestingHTTPClient(t, server.URL(), integration.AdminUserID)

	predefinedTags := []struct {
		fqName    []string
		tagTypeID string
	}{
		{
			fqName:    []string{"label"},
			tagTypeID: "0x0000",
		},
		{
			fqName:    []string{"application"},
			tagTypeID: "0x0001",
		},
		{
			fqName:    []string{"tier"},
			tagTypeID: "0x0002",
		},
		{
			fqName:    []string{"deployment"},
			tagTypeID: "0x0003",
		},
		{
			fqName:    []string{"site"},
			tagTypeID: "0x0004",
		},
	}
	for _, tag := range predefinedTags {
		uuid := c.FQNameToID(t, tag.fqName, tagTypeTypeName)
		assert.NotEmpty(t, uuid)

		resp, err := c.GetTagType(context.Background(), &services.GetTagTypeRequest{ID: uuid})
		assert.NoError(t, err)
		assert.Equal(t, tag.tagTypeID, resp.GetTagType().GetTagTypeID())
	}
}
