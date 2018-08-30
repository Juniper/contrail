package basedb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

func TestBuildMetadataFilter(t *testing.T) {
	tests := []struct {
		name           string
		args           []*basemodels.Metadata
		want           string
		mysqlExpect    string
		postgresExpect string
		fails          bool
	}{
		{
			name: "Get multiple metadatas using UUID and FQName",
			args: []*basemodels.Metadata{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
			mysqlExpect:    " ( uuid = ? )  or  ( type = ? and fq_name = ? ) ",
			postgresExpect: " ( uuid = $1 )  or  ( type = $2 and fq_name = $3 ) ",
		},
		{
			name: "Get multiple metadatas using UUIDs",
			args: []*basemodels.Metadata{
				{
					UUID: "uuid-b",
				},
				{
					UUID: "uuid-c",
				},
			},
			mysqlExpect:    " ( uuid = ? )  or  ( uuid = ? ) ",
			postgresExpect: " ( uuid = $1 )  or  ( uuid = $2 ) ",
		},
		{
			name: "Get multiple metadatas using FQNames",
			args: []*basemodels.Metadata{
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
				{
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
			mysqlExpect:    " ( type = ? and fq_name = ? )  or  ( type = ? and fq_name = ? ) ",
			postgresExpect: " ( type = $1 and fq_name = $2 )  or  ( type = $3 and fq_name = $4 ) ",
		},
		{
			name: "Provide only FQNames - fail",
			args: []*basemodels.Metadata{
				{
					FQName: []string{"default", "uuid-b"},
				},
				{
					FQName: []string{"default", "uuid-c"},
				},
			},
		},
		{
			name: "Get metadata using FQName",
			args: []*basemodels.Metadata{
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
			mysqlExpect:    " ( type = ? and fq_name = ? ) ",
			postgresExpect: " ( type = $1 and fq_name = $2 ) ",
		},
		{
			name: "Get metadata using UUID",
			args: []*basemodels.Metadata{
				{
					UUID: "uuid-b",
				},
			},
			mysqlExpect:    " ( uuid = ? ) ",
			postgresExpect: " ( uuid = $1 ) ",
		},

		{
			name: "Get single metadata using UUID and FQName",
			args: []*basemodels.Metadata{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
			mysqlExpect:    " ( uuid = ? )  or  ( type = ? and fq_name = ? ) ",
			postgresExpect: " ( uuid = $1 )  or  ( type = $2 and fq_name = $3 ) ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mysqlDialect := NewDialect(MYSQL)
			postgresDialect := NewDialect(POSTGRES)

			mysql, _, err1 := buildMetadataFilter(mysqlDialect, tt.args)
			postgres, _, err2 := buildMetadataFilter(postgresDialect, tt.args)

			if tt.fails {
				assert.Error(t, err1)
				assert.Error(t, err2)
				return
			}

			assert.Equal(t, tt.mysqlExpect, mysql)
			assert.Equal(t, tt.postgresExpect, postgres)
		})
	}
}
