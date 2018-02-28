package db

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestQueryBuffer(t *testing.T) {
	qb := &QueryBuilder{
		Table:  "virtual_networks",
		Fields: []string{"uuid", "fq_name"},
		RefFields: map[string][]string{
			"subnet": {"uuid", "fq_name"},
		},
		BackRefFields: map[string][]string{
			"project": {"uuid", "fq_name"},
		},
	}
	spec := &models.ListSpec{
		Limit: 1,
	}
	qb.Dialect = NewDialect("mysql")
	assert.Equal(t, "`a`.`b`.`c`", qb.quote("a", "b", "c"))
	query, _, _ := qb.ListQuery(nil, spec)
	assert.Equal(t, "select `virtual_networks`.`uuid`,`virtual_networks`.`fq_name` from virtual_networks   limit 1 offset 0 ", query) // nolint
	spec.Detail = true
	query, _, _ = qb.ListQuery(nil, spec)
	assert.Equal(t, "select ANY_VALUE(`virtual_networks`.`uuid`),ANY_VALUE(`virtual_networks`.`fq_name`),group_concat(distinct JSON_OBJECT('uuid',`ref_virtual_networks_subnet`.`uuid`,'fq_name',`ref_virtual_networks_subnet`.`fq_name`,'from',`ref_virtual_networks_subnet`.`from`,'to',`ref_virtual_networks_subnet`.`to`)) as `ref_virtual_networks_subnet_ref`,group_concat(distinct JSON_OBJECT('uuid',`project`.`uuid`,'fq_name',`project`.`fq_name`)) as `project_ref` from virtual_networks left join `ref_virtual_networks_subnet` on `virtual_networks`.`uuid` = `ref_virtual_networks_subnet`.`from` left join `project` on `virtual_networks`.`uuid` = `project`.`parent_uuid` group by `ref_virtual_networks_subnet`.`from`,`project`.`uuid`  limit 1 offset 0 ", query) // nolint
	// try postgres query
	qb.Dialect = NewDialect("postgres")
	spec.Detail = false
	assert.Equal(t, "\"a\".\"b\".\"c\"", qb.quote("a", "b", "c"))
	query, _, _ = qb.ListQuery(nil, spec)
	assert.Equal(t, "select \"virtual_networks\".\"uuid\",\"virtual_networks\".\"fq_name\" from virtual_networks   limit 1 offset 0 ", query) // nolint
	spec.Detail = true
	query, _, _ = qb.ListQuery(nil, spec)
	assert.Equal(t, "select \"virtual_networks\".\"uuid\",\"virtual_networks\".\"fq_name\",json_agg(json_build_object('uuid',\"ref_virtual_networks_subnet\".\"uuid\",'fq_name',\"ref_virtual_networks_subnet\".\"fq_name\",'from',\"ref_virtual_networks_subnet\".\"from\",'to',\"ref_virtual_networks_subnet\".\"to\")) as \"ref_virtual_networks_subnet_ref\",json_agg(json_build_object('uuid',\"project\".\"uuid\",'fq_name',\"project\".\"fq_name\")) as \"project_ref\" from virtual_networks left join \"ref_virtual_networks_subnet\" on \"virtual_networks\".\"uuid\" = \"ref_virtual_networks_subnet\".\"from\" left join \"project\" on \"virtual_networks\".\"uuid\" = \"project\".\"parent_uuid\" group by \"ref_virtual_networks_subnet\".\"from\",\"project\".\"uuid\"  limit 1 offset 0 ", query) // nolint
}
