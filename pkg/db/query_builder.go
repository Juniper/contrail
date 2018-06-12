package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

const (
	//MYSQL db type
	MYSQL = "mysql"
	//POSTGRES db type
	POSTGRES = "postgres"
)

//QueryBuilder builds list query.
type QueryBuilder struct {
	Dialect
	Fields        []string
	Table         string
	RefFields     map[string][]string
	BackRefFields map[string][]string
}

type queryContext struct {
	auth        *common.AuthContext
	values      []interface{}
	columns     Columns
	columnParts []string
	where       []string
	joins       []string
	query       *bytes.Buffer
	spec        *services.ListSpec
}

func newQueryContext() *queryContext {
	var query bytes.Buffer
	return &queryContext{
		query:       &query,
		columns:     Columns{},
		columnParts: []string{},
		values:      []interface{}{},
		where:       []string{},
		joins:       []string{},
	}
}

//NewQueryBuilder makes a query builder.
func NewQueryBuilder(dialect Dialect, table string,
	fields []string, refFields map[string][]string,
	backRefFields map[string][]string) *QueryBuilder {
	qb := &QueryBuilder{}
	qb.Dialect = dialect
	qb.Table = table
	qb.Fields = fields
	qb.RefFields = refFields
	qb.BackRefFields = backRefFields
	return qb
}

//NewDialect creates NewDialect objects.
func NewDialect(mode string) Dialect {
	switch mode {
	case MYSQL:
		return Dialect{
			Name:             MYSQL,
			QuoteRune:        "`",
			JSONAggFuncStart: "group_concat(distinct JSON_OBJECT(",
			JSONAggFuncEnd:   "))",
			AnyValueString:   "ANY_VALUE(",
			PlaceHolderIndex: false,
		}
	default:
		return Dialect{
			Name:             POSTGRES,
			QuoteRune:        `"`,
			JSONAggFuncStart: "json_agg(json_build_object(",
			JSONAggFuncEnd:   "))",
			AnyValueString:   "",
			PlaceHolderIndex: true,
		}
	}
}

//Dialect represents database dialect.
type Dialect struct {
	Name             string
	QuoteRune        string
	JSONAggFuncStart string
	JSONAggFuncEnd   string
	AnyValueString   string
	PlaceHolderIndex bool
}

func (d *Dialect) quote(params ...string) string {
	query := ""
	l := len(params)
	for i := 0; i < l-1; i++ {
		query += d.QuoteRune + strings.ToLower(params[i]) + d.QuoteRune + "."
	}
	query += d.QuoteRune + strings.ToLower(params[l-1]) + d.QuoteRune
	return query
}

func (d *Dialect) placeholder(i int) string {
	if d.PlaceHolderIndex {
		return "$" + strconv.Itoa(i)
	}
	return "?"
}

func (d *Dialect) values(params ...string) string {
	query := ""
	l := len(params)
	for i := 0; i < l-1; i++ {
		query += d.placeholder(i+1) + ","
	}
	query += d.placeholder(l)
	return query
}

func (d *Dialect) quoteSep(params ...string) string {
	query := ""
	l := len(params)
	for i := 0; i < l-1; i++ {
		query += d.QuoteRune + strings.ToLower(params[i]) + d.QuoteRune + ","
	}
	query += d.QuoteRune + strings.ToLower(params[l-1]) + d.QuoteRune
	return query
}

func (d *Dialect) jsonAgg(table string, params ...string) string {
	if d.Name == POSTGRES {
		return "json_agg(row_to_json(" + d.quote(table) + "))"
	}
	query := ""
	l := len(params)
	query += d.JSONAggFuncStart
	for i := 0; i < l-1; i++ {
		query += "'" + params[i] + "'" + "," + d.quote(table, params[i]) + ","
	}
	query += "'" + params[l-1] + "'" + "," + d.quote(table, params[l-1]) + d.JSONAggFuncEnd
	return query
}

func (d *Dialect) anyValue(params ...string) string {
	if d.AnyValueString != "" {
		return d.AnyValueString + d.quote(params...) + ")"
	}
	return d.quote(params...)
}

//Columns represents column index.
type Columns map[string]int

func (qb *QueryBuilder) buildFilterParts(ctx *queryContext, column string, filterValues []string) string {
	var where string
	var err error
	if len(filterValues) == 1 {
		ctx.values = append(ctx.values, filterValues[0])
		where = column + " = " + qb.placeholder(len(ctx.values))
	} else {
		var filterQuery bytes.Buffer
		_, err = filterQuery.WriteString(column)
		_, err = filterQuery.WriteString(" in (")
		last := len(filterValues) - 1
		for _, value := range filterValues[:last] {
			ctx.values = append(ctx.values, value)
			_, err = filterQuery.WriteString(qb.placeholder(len(ctx.values)) + ",")
		}
		ctx.values = append(ctx.values, filterValues[last])
		_, err = filterQuery.WriteString(qb.placeholder(len(ctx.values)) + ")")

		where = filterQuery.String()
	}
	if err != nil {
		log.Fatal(err)
	}
	return where
}

func (qb *QueryBuilder) join(fromTable, fromProperty, toTable string) string {
	return "left join " + qb.quote(
		fromTable) + " on " + qb.quote(
		toTable, "uuid") + " = " + qb.quote(fromTable, fromProperty)
}

func (qb *QueryBuilder) as(a, b string) string {
	return a + " as " + b
}

func (qb *QueryBuilder) buildFilterQuery(ctx *queryContext) {
	spec := ctx.spec
	filters := spec.Filters
	filters = services.AppendFilter(filters, "uuid", spec.ObjectUUIDs...)
	filters = services.AppendFilter(filters, "parent_uuid", spec.ParentUUIDs...)
	if spec.ParentType != "" {
		filters = services.AppendFilter(filters, "parent_type", spec.ParentType)
	}
	for _, filter := range filters {
		if !qb.isValidField(filter.Key) {
			continue
		}
		column := qb.quote(qb.Table, filter.Key)
		where := qb.buildFilterParts(ctx, column, filter.Values)
		ctx.where = append(ctx.where, where)
	}
	if len(spec.BackRefUUIDs) > 0 {
		where := []string{}
		for refTable := range qb.BackRefFields {
			column := qb.quote(refTable, "uuid")
			wherePart := qb.buildFilterParts(ctx, column, spec.BackRefUUIDs)
			where = append(where, wherePart)
		}
		ctx.where = append(ctx.where, "("+strings.Join(where, " or ")+")")
	}
}

func (qb *QueryBuilder) buildAuthQuery(ctx *queryContext) {
	auth := ctx.auth
	spec := ctx.spec
	where := []string{}

	if !auth.IsAdmin() {
		ctx.values = append(ctx.values, auth.ProjectID())
		where = append(where, qb.quote(qb.Table, "owner")+" = "+qb.placeholder(len(ctx.values)))
	}
	if spec.Shared {
		shareTables := []string{"domain_share_" + qb.Table, "tenant_share_" + qb.Table}
		for i, shareTable := range shareTables {
			ctx.joins = append(ctx.joins,
				qb.join(shareTable, "uuid", qb.Table))
			where = append(where, fmt.Sprintf("(%s.to = %s and %s.access >= 4)",
				qb.quote(shareTable), qb.placeholder(len(ctx.values)+i+1), qb.quote(shareTable)))
		}
		ctx.values = append(ctx.values, auth.DomainID(), auth.ProjectID())
	}
	if len(where) > 0 {
		ctx.where = append(ctx.where, fmt.Sprintf("(%s)", strings.Join(where, " or ")))
	}
}

func (qb *QueryBuilder) buildQuery(ctx *queryContext) {
	spec := ctx.spec
	query := ctx.query
	var err error
	_, err = query.WriteString("select ")
	if len(ctx.columnParts) != len(ctx.columns) {
		log.Fatal("unmatch")
	}
	_, err = query.WriteString(strings.Join(ctx.columnParts, ","))
	_, err = query.WriteString(" from ")
	_, err = query.WriteString(qb.Table)
	_, err = query.WriteRune(' ')
	if len(ctx.joins) > 0 {
		_, err = query.WriteString(strings.Join(ctx.joins, " "))
	}
	if len(ctx.where) > 0 {
		_, err = query.WriteString(" where ")
		_, err = query.WriteString(strings.Join(ctx.where, " and "))
	}
	if spec.Shared || len(spec.BackRefUUIDs) > 0 {
		_, err = query.WriteString(" group by ")
		_, err = query.WriteString(qb.quote(qb.Table, "uuid"))
	}
	_, err = query.WriteRune(' ')
	if spec.Limit > 0 {
		pagenationQuery := fmt.Sprintf(" limit %d offset %d ", spec.Limit, spec.Offset)
		_, err = query.WriteString(pagenationQuery)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (qb *QueryBuilder) buildRefQuery(ctx *queryContext) {
	spec := ctx.spec
	if !spec.Detail {
		return
	}
	for linkTo, refFields := range qb.RefFields {
		refTable := strings.ToLower("ref_" + qb.Table + "_" + linkTo)
		refFields = append(refFields, "from")
		refFields = append(refFields, "to")
		subQuery := "(select " +
			qb.as(qb.jsonAgg(refTable+"_t", refFields...), qb.quote(refTable+"_ref")) +
			" from " + qb.as(qb.quote(refTable), refTable+"_t") +
			" where " + qb.quote(qb.Table, "uuid") + " = " + qb.quote(refTable+"_t", "from") +
			" group by " + qb.quote(refTable+"_t", "from") + " )"
		ctx.columnParts = append(
			ctx.columnParts,
			subQuery)
		ctx.columns["ref_"+linkTo] = len(ctx.columns)
	}
}

func (qb *QueryBuilder) buildBackRefQuery(ctx *queryContext) {
	spec := ctx.spec
	// use join if backrefuuids
	if len(spec.BackRefUUIDs) > 0 {
		for refTable, refFields := range qb.BackRefFields {
			refTable = strings.ToLower(refTable)
			if spec.Detail {
				ctx.columnParts = append(
					ctx.columnParts,
					qb.as(qb.jsonAgg(refTable, refFields...), qb.quote(refTable+"_ref")),
				)
				ctx.columns["backref_"+refTable] = len(ctx.columns)
			}
			ctx.joins = append(ctx.joins,
				qb.join(refTable, "parent_uuid", qb.Table))
		}
		return
	}
	if !spec.Detail {
		return
	}
	// use sub query if no backrefuuids
	for refTable, refFields := range qb.BackRefFields {
		refTable = strings.ToLower(refTable)
		subQuery := "(select " +
			qb.as(qb.jsonAgg(refTable+"_t", refFields...), qb.quote(refTable+"_ref")) +
			" from " + qb.as(qb.quote(refTable), refTable+"_t") +
			" where " + qb.quote(qb.Table, "uuid") + " = " + qb.quote(refTable+"_t", "parent_uuid") +
			" group by " + qb.quote(refTable+"_t", "parent_uuid") + " )"
		ctx.columnParts = append(
			ctx.columnParts,
			subQuery)
		ctx.columns["backref_"+refTable] = len(ctx.columns)
	}
}

func (qb *QueryBuilder) isValidField(requestedField string) bool {
	for _, field := range qb.Fields {
		if field == requestedField {
			return true
		}
	}
	return false
}

func (qb *QueryBuilder) checkRequestedFields(ctx *queryContext) bool {
	spec := ctx.spec
	for _, requestedField := range spec.Fields {
		if !qb.isValidField(requestedField) {
			return false
		}
	}
	return true
}

func (qb *QueryBuilder) buildColumns(ctx *queryContext) {
	spec := ctx.spec
	fields := qb.Fields

	if len(spec.Fields) > 0 && qb.checkRequestedFields(ctx) {
		fields = spec.Fields
	}

	if spec.Shared || len(spec.BackRefUUIDs) > 0 {
		for _, column := range fields {
			ctx.columns[column] = len(ctx.columns)
			ctx.columnParts = append(ctx.columnParts, qb.anyValue(qb.Table, column))
		}
	} else {
		for _, column := range fields {
			ctx.columns[column] = len(ctx.columns)
			ctx.columnParts = append(ctx.columnParts, qb.quote(qb.Table, column))
		}
	}
}

//ListQuery makes sql query.
func (qb *QueryBuilder) ListQuery(auth *common.AuthContext, spec *services.ListSpec) (string, Columns, []interface{}) {
	ctx := newQueryContext()
	ctx.auth = auth
	ctx.spec = spec
	qb.buildColumns(ctx)
	qb.buildFilterQuery(ctx)
	qb.buildAuthQuery(ctx)
	qb.buildRefQuery(ctx)
	qb.buildBackRefQuery(ctx)
	qb.buildQuery(ctx)
	return ctx.query.String(), ctx.columns, ctx.values
}

//CreateQuery makes sql query.
func (qb *QueryBuilder) CreateQuery() string {
	query := ("insert into " + qb.quote(qb.Table) + "(" +
		qb.quoteSep(qb.Fields...) + ") values (" + qb.values(qb.Fields...) + ")")
	return query
}

//CreateRefQuery makes references.
func (qb *QueryBuilder) CreateRefQuery(linkto string) string {
	fields := append([]string{"from", "to"}, qb.RefFields[linkto]...)
	return ("insert into ref_" + qb.Table + "_" + linkto +
		" (" + qb.quoteSep(fields...) + ") values (" + qb.values(fields...) + ")")
}

//DeleteQuery makes sql query.
func (qb *QueryBuilder) DeleteQuery() string {
	return "delete from " + qb.quote(qb.Table) + " where uuid = " + qb.placeholder(1)
}

//DeleteRefQuery makes sql query.
func (qb *QueryBuilder) DeleteRefQuery(linkto string) string {
	return "delete from ref_" + qb.Table + "_" + linkto + " where " + qb.quote("from") + " = " + qb.placeholder(1)
}

//SelectAuthQuery makes sql query.
func (qb *QueryBuilder) SelectAuthQuery(admin bool) string {
	query := "select count(uuid) from " + qb.quote(qb.Table) + " where uuid = " + qb.placeholder(1)
	if !admin {
		query += " and owner = " + qb.placeholder(2)
	}
	return query
}

//UpdateQuery makes sql query for update.
// nolint
func (qb *QueryBuilder) UpdateQuery(columns []string) string {
	var query bytes.Buffer
	query.WriteString("update ")
	query.WriteString(qb.quote(qb.Table))
	query.WriteString("set ")
	for i, column := range columns {
		query.WriteString(qb.quote(column))
		query.WriteString(" = ")
		query.WriteString(qb.placeholder(i + 1))
		if i < len(columns)-1 {
			query.WriteString(", ")
		}
	}
	query.WriteString(" where uuid = ")
	query.WriteString(qb.placeholder(len(columns) + 1))
	return query.String()
}

func (qb *QueryBuilder) scanResourceList(value interface{}) []interface{} {
	var resources []interface{}
	stringValue := common.InterfaceToString(value)
	if stringValue == "" {
		return nil
	}
	switch qb.Dialect.Name {
	case MYSQL:
		err := json.Unmarshal([]byte("["+stringValue+"]"), &resources)
		if err != nil {
			log.Debug(err)
			return nil
		}
	case POSTGRES:
		err := json.Unmarshal([]byte(stringValue), &resources)
		if err != nil {
			log.Debug(err)
			return nil
		}
	default:
		log.Fatal("unsupported db dialect")
	}
	return resources
}
