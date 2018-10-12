package basedb

import (
	"testing"

	"github.com/Juniper/contrail/pkg/services/baseservices"

	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder(t *testing.T) {
	type expectedResult struct {
		query  string
		values []interface{}
	}

	tests := []struct {
		name           string
		table          string
		fields         []string
		refFields      map[string][]string
		childFields    map[string][]string
		backRefFields  map[string][]string
		spec           baseservices.ListSpec
		mysqlExpect    expectedResult
		postgresExpect expectedResult
	}{
		{
			name: "Collecion of alarms",
			fields: []string{
				"global_access",
				"group_access",
			},
			table: "alert",
			spec:  baseservices.ListSpec{},
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
			fields: []string{
				"uuid",
				"name",
			},
			table: "virtual_network",
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
			fields: []string{
				"parent_type",
				"parent_uuid",
			},
			table: "virtual_network",
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
			mysqlDialect := NewDialect(MYSQL)
			postgresDialect := NewDialect(POSTGRES)

			mysql, _, mysqlValues :=
				NewQueryBuilder(mysqlDialect, tt.table, tt.fields, tt.refFields, tt.childFields, tt.refFields).
					ListQuery(nil, &tt.spec)
			postgres, _, postgresValues :=
				NewQueryBuilder(postgresDialect, tt.table, tt.fields, tt.refFields, tt.childFields, tt.refFields).
					ListQuery(nil, &tt.spec)

			assert.Equal(t, tt.mysqlExpect.query, mysql)
			assert.Equal(t, tt.mysqlExpect.values, mysqlValues)
			assert.Equal(t, tt.postgresExpect.query, postgres)
			assert.Equal(t, tt.postgresExpect.values, postgresValues)
		})
	}
}
