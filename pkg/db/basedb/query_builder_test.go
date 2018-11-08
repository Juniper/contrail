package basedb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/services/baseservices"
)

type queryBuilderParams struct {
	table         string
	fields        []string
	refFields     map[string][]string
	childFields   map[string][]string
	backRefFields map[string][]string
}

func TestQueryBuilder(t *testing.T) {
	type expectedResult struct {
		query  string
		values []interface{}
	}

	tests := []struct {
		name               string
		queryBuilderParams queryBuilderParams
		spec               baseservices.ListSpec

		mysqlExpect    expectedResult
		postgresExpect expectedResult
	}{
		{
			name: "Collecion of alarms",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"global_access",
					"group_access",
				},
				table: "alert",
			},
			spec: baseservices.ListSpec{},
			mysqlExpect: expectedResult{
				query: "select `alert_t`.`global_access`,`alert_t`.`group_access` from alert " +
					"as alert_t order by `alert_t`.`uuid`",
				values: []interface{}{},
			},
			postgresExpect: expectedResult{
				query: "select \"alert_t\".\"global_access\",\"alert_t\".\"group_access\" from alert " +
					"as alert_t order by \"alert_t\".\"uuid\"",
				values: []interface{}{},
			},
		},
		{
			name: "Collecion of virtual networks are limited",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"uuid",
					"name",
				},
				table: "virtual_network",
			},
			spec: baseservices.ListSpec{
				Limit: 10,
			},
			mysqlExpect: expectedResult{
				query: "select `virtual_network_t`.`uuid`,`virtual_network_t`.`name` from virtual_network " +
					"as virtual_network_t order by `virtual_network_t`.`uuid` limit 10",
				values: []interface{}{},
			},
			postgresExpect: expectedResult{
				query: "select \"virtual_network_t\".\"uuid\",\"virtual_network_t\".\"name\" from virtual_network " +
					"as virtual_network_t order by \"virtual_network_t\".\"uuid\" limit 10",
				values: []interface{}{},
			},
		},
		{
			name: "Collecion of virtual networks are limited and started with marker",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"parent_type",
					"parent_uuid",
				},
				table: "virtual_network",
			},
			spec: baseservices.ListSpec{
				Limit:  5,
				Marker: "marker_uuid",
			},
			mysqlExpect: expectedResult{
				query: "select `virtual_network_t`.`parent_type`,`virtual_network_t`.`parent_uuid` from virtual_network " +
					"as virtual_network_t where `virtual_network_t`.`uuid` > ? order by `virtual_network_t`.`uuid` limit 5",
				values: []interface{}{
					"marker_uuid",
				},
			},
			postgresExpect: expectedResult{
				query: "select \"virtual_network_t\".\"parent_type\",\"virtual_network_t\".\"parent_uuid\" from virtual_network " +
					"as virtual_network_t where \"virtual_network_t\".\"uuid\" > $1 order by \"virtual_network_t\".\"uuid\" limit 5",
				values: []interface{}{
					"marker_uuid",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mysql, _, mysqlValues := newQueryBuilder(MYSQL, tt.queryBuilderParams).ListQuery(nil, &tt.spec)
			postgres, _, postgresValues := newQueryBuilder(POSTGRES, tt.queryBuilderParams).ListQuery(nil, &tt.spec)

			assert.Equal(t, tt.mysqlExpect.query, mysql)
			assert.Equal(t, tt.mysqlExpect.values, mysqlValues)
			assert.Equal(t, tt.postgresExpect.query, postgres)
			assert.Equal(t, tt.postgresExpect.values, postgresValues)
		})
	}
}

func TestRelaxRefQuery(t *testing.T) {
	tests := []struct {
		name               string
		queryBuilderParams queryBuilderParams
		linkTo             string

		mysqlExpect    string
		postgresExpect string
	}{
		{
			name: "Reference from VirtualNetwork to NetworkPolicy",
			queryBuilderParams: queryBuilderParams{
				table: "virtual_network",
			},
			linkTo:         "network_policy",
			mysqlExpect:    "update ref_virtual_network_network_policy set `relaxed` = true where `from` = ? and `to` = ?",
			postgresExpect: `update ref_virtual_network_network_policy set "relaxed" = true where "from" = $1 and "to" = $2`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mysql := newQueryBuilder(MYSQL, tt.queryBuilderParams).RelaxRefQuery(tt.linkTo)
			postgres := newQueryBuilder(POSTGRES, tt.queryBuilderParams).RelaxRefQuery(tt.linkTo)

			assert.Equal(t, tt.mysqlExpect, mysql)
			assert.Equal(t, tt.postgresExpect, postgres)
		})
	}
}

func TestDeleteRelaxedBackrefsQuery(t *testing.T) {
	tests := []struct {
		name               string
		queryBuilderParams queryBuilderParams
		linkFrom           string

		mysqlExpect    string
		postgresExpect string
	}{
		{
			name:     "References from VirtualNetwork to NetworkPolicy",
			linkFrom: "virtual_network",
			queryBuilderParams: queryBuilderParams{
				table: "network_policy",
			},
			mysqlExpect:    "delete from ref_virtual_network_network_policy where `to` = ? and `relaxed` = true",
			postgresExpect: `delete from ref_virtual_network_network_policy where "to" = $1 and "relaxed" = true`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mysql := newQueryBuilder(MYSQL, tt.queryBuilderParams).DeleteRelaxedBackrefsQuery(tt.linkFrom)
			postgres := newQueryBuilder(POSTGRES, tt.queryBuilderParams).DeleteRelaxedBackrefsQuery(tt.linkFrom)

			assert.Equal(t, tt.mysqlExpect, mysql)
			assert.Equal(t, tt.postgresExpect, postgres)
		})
	}
}

func newQueryBuilder(dialect string, p queryBuilderParams) *QueryBuilder {
	return NewQueryBuilder(
		NewDialect(dialect),
		p.table,
		p.fields,
		p.refFields,
		p.childFields,
		p.backRefFields)
}
