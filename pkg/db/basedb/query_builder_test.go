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
		{
			name: "Listing virtual networks with no child fields",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"uuid",
					// TODO Test with non-mandatory fields here.
				},
				childFields: map[string][]string{
					"access_control_list": []string{
						"uuid",
					},
				},
				table: "virtual_network",
			},
			spec: baseservices.ListSpec{},
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select " +
						"`virtual_network_t`.`uuid` " +
						"from virtual_network as virtual_network_t " +
						"order by `virtual_network_t`.`uuid`",
					values: []interface{}{},
				},
				POSTGRES: {
					query: `select ` +
						`"virtual_network_t"."uuid" ` +
						`from virtual_network as virtual_network_t ` +
						`order by "virtual_network_t"."uuid"`,
					values: []interface{}{},
				},
			},
		},
		{
			name: "Listing virtual networks with only one child field",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"uuid",
				},
				childFields: map[string][]string{
					"access_control_list": []string{
						"uuid",
					},
					// TODO Test with multiple kinds of children here.
				},
				table: "virtual_network",
			},
			spec: baseservices.ListSpec{
				Fields: []string{
					"access_control_lists",
				},
			},
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select " +
						"`virtual_network_t`.`uuid`," +
						"(select group_concat(JSON_OBJECT('uuid',`access_control_list_t`.`uuid`)) as `access_control_list_ref` " +
						"from `access_control_list` as access_control_list_t " +
						"where `virtual_network_t`.`uuid` = `access_control_list_t`.`parent_uuid` " +
						"group by `access_control_list_t`.`parent_uuid` ) " +
						"from virtual_network as virtual_network_t " +
						"order by `virtual_network_t`.`uuid`",
					values: []interface{}{},
				},
				POSTGRES: {
					query: `select ` +
						`"virtual_network_t"."uuid",` +
						`(select json_agg(row_to_json("access_control_list_t")) as "access_control_list_ref" ` +
						`from "access_control_list" as access_control_list_t ` +
						`where "virtual_network_t"."uuid" = "access_control_list_t"."parent_uuid" ` +
						`group by "access_control_list_t"."parent_uuid" ) ` +
						`from virtual_network as virtual_network_t ` +
						`order by "virtual_network_t"."uuid"`,
					values: []interface{}{},
				},
			},
		},
		{
			name: "Listing virtual networks with Detail",
			queryBuilderParams: queryBuilderParams{
				fields: []string{
					"uuid",
				},
				childFields: map[string][]string{
					"access_control_list": []string{
						"uuid",
					},
				},
				table: "virtual_network",
			},
			spec: baseservices.ListSpec{
				Detail: true,
			},
			expected: map[string]expectedResult{
				MYSQL: {
					query: "select " +
						"`virtual_network_t`.`uuid`," +
						"(select group_concat(JSON_OBJECT('uuid',`access_control_list_t`.`uuid`)) as `access_control_list_ref` " +
						"from `access_control_list` as access_control_list_t " +
						"where `virtual_network_t`.`uuid` = `access_control_list_t`.`parent_uuid` " +
						"group by `access_control_list_t`.`parent_uuid` ) " +
						"from virtual_network as virtual_network_t " +
						"order by `virtual_network_t`.`uuid`",
					values: []interface{}{},
				},
				POSTGRES: {
					query: `select ` +
						`"virtual_network_t"."uuid",` +
						`(select json_agg(row_to_json("access_control_list_t")) as "access_control_list_ref" ` +
						`from "access_control_list" as access_control_list_t ` +
						`where "virtual_network_t"."uuid" = "access_control_list_t"."parent_uuid" ` +
						`group by "access_control_list_t"."parent_uuid" ) ` +
						`from virtual_network as virtual_network_t ` +
						`order by "virtual_network_t"."uuid"`,
					values: []interface{}{},
				},
			},
		},
		// TODO Test backrefs as well.
		// TODO Test refs as well.
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
