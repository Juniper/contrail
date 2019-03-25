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

		expected map[string]expectedResult // map[dialect]expectedResult
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
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select `alert_t`.`global_access`,`alert_t`.`group_access` from alert " +
						"as alert_t order by `alert_t`.`uuid`",
					values: []interface{}{},
				},
				POSTGRES: {
					query: "select \"alert_t\".\"global_access\",\"alert_t\".\"group_access\" from alert " +
						"as alert_t order by \"alert_t\".\"uuid\"",
					values: []interface{}{},
				},
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
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select `virtual_network_t`.`uuid`,`virtual_network_t`.`name` from virtual_network " +
						"as virtual_network_t order by `virtual_network_t`.`uuid` limit 10",
					values: []interface{}{},
				},
				POSTGRES: {
					query: "select \"virtual_network_t\".\"uuid\",\"virtual_network_t\".\"name\" from virtual_network " +
						"as virtual_network_t order by \"virtual_network_t\".\"uuid\" limit 10",
					values: []interface{}{},
				},
			},
		},
		{
			name: "Test one ref filter from three ref fields",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"uuid",
					"name",
				},
				refFields: map[string][]string{
					"project": nil,
					"vmi":     nil,
					"tag":     nil,
				},
				table: "floating_ip",
			},
			spec: baseservices.ListSpec{
				RefUUIDs: map[string]*baseservices.UUIDs{
					"project_refs": {[]string{"proj_ref_uuid", "proj_ref_uuid_2"}},
				},
			},
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select `floating_ip_t`.`uuid`,`floating_ip_t`.`name` from floating_ip as floating_ip_t " +
						"left join `ref_floating_ip_project` on `floating_ip_t`.`uuid` = `ref_floating_ip_project`.`from` " +
						"where (CONVERT(`ref_floating_ip_project`.`to` USING utf8) in (?,?)) order by `floating_ip_t`.`uuid`",
					values: []interface{}{
						"proj_ref_uuid",
						"proj_ref_uuid_2",
					},
				},
				POSTGRES: {
					query: "select \"floating_ip_t\".\"uuid\",\"floating_ip_t\".\"name\" from floating_ip as floating_ip_t " +
						"left join \"ref_floating_ip_project\" on \"floating_ip_t\".\"uuid\" = \"ref_floating_ip_project\".\"from\" " +
						"where (\"ref_floating_ip_project\".\"to\" :: TEXT in ($1,$2)) order by \"floating_ip_t\".\"uuid\"",
					values: []interface{}{
						"proj_ref_uuid",
						"proj_ref_uuid_2",
					},
				},
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
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select `virtual_network_t`.`parent_type`,`virtual_network_t`.`parent_uuid` from virtual_network " +
						"as virtual_network_t where `virtual_network_t`.`uuid` > ? order by `virtual_network_t`.`uuid` limit 5",
					values: []interface{}{
						"marker_uuid",
					},
				},
				POSTGRES: {
					query: "select \"virtual_network_t\".\"parent_type\",\"virtual_network_t\".\"parent_uuid\" from virtual_network " +
						"as virtual_network_t where \"virtual_network_t\".\"uuid\" > $1 order by \"virtual_network_t\".\"uuid\" limit 5",
					values: []interface{}{
						"marker_uuid",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for dialect, expectedResult := range tt.expected {
				query, _, values := newQueryBuilder(dialect, tt.queryBuilderParams).ListQuery(nil, &tt.spec)
				assert.Equal(t, expectedResult.query, query)
				assert.Equal(t, expectedResult.values, values)
			}
		})
	}
}

func TestRelaxRefQuery(t *testing.T) {
	tests := []struct {
		name               string
		queryBuilderParams queryBuilderParams
		linkTo             string

		expected map[string]string // map[dialect]expectedQuery
	}{
		{
			name: "Reference from VirtualNetwork to NetworkPolicy",
			queryBuilderParams: queryBuilderParams{
				table: "virtual_network",
			},
			linkTo: "network_policy",
			expected: map[string]string{
				MYSQL:    "update ref_virtual_network_network_policy set `relaxed` = true where `from` = ? and `to` = ?",
				POSTGRES: `update ref_virtual_network_network_policy set "relaxed" = true where "from" = $1 and "to" = $2`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for dialect, expectedQuery := range tt.expected {
				query := newQueryBuilder(dialect, tt.queryBuilderParams).RelaxRefQuery(tt.linkTo)
				assert.Equal(t, expectedQuery, query)
			}
		})
	}
}

func TestDeleteRelaxedBackrefsQuery(t *testing.T) {
	tests := []struct {
		name               string
		queryBuilderParams queryBuilderParams
		linkFrom           string

		expected map[string]string // map[dialect]expectedQuery
	}{
		{
			name:     "References from VirtualNetwork to NetworkPolicy",
			linkFrom: "virtual_network",
			queryBuilderParams: queryBuilderParams{
				table: "network_policy",
			},
			expected: map[string]string{
				MYSQL:    "delete from ref_virtual_network_network_policy where `to` = ? and `relaxed` = true",
				POSTGRES: `delete from ref_virtual_network_network_policy where "to" = $1 and "relaxed" = true`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for dialect, expectedQuery := range tt.expected {
				query := newQueryBuilder(dialect, tt.queryBuilderParams).DeleteRelaxedBackrefsQuery(tt.linkFrom)
				assert.Equal(t, expectedQuery, query)
			}
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
