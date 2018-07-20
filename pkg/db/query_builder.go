package db

// TODO: replace this file with some ORM framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/schema"
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
	TableAlias    string
	RefFields     map[string][]string
	ChildFields   map[string][]string
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
func NewQueryBuilder(
	dialect Dialect, table string, fields []string, refFields map[string][]string,
	childFields map[string][]string, backRefFields map[string][]string,
) *QueryBuilder {
	return &QueryBuilder{
		Dialect:       dialect,
		Table:         table,
		TableAlias:    table + "_t",
		Fields:        fields,
		RefFields:     refFields,
		ChildFields:   childFields,
		BackRefFields: backRefFields,
	}
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
			IPLiteralPrefix:  "INET6_ATON('",
			PpLiteralSuffix:  "')",
			SelectIPPrefix:   "INET6_NTOA(`",
			SelectIPSuffix:   "`)",
		}
	default:
		return Dialect{
			Name:             POSTGRES,
			QuoteRune:        `"`,
			JSONAggFuncStart: "json_agg(json_build_object(",
			JSONAggFuncEnd:   "))",
			AnyValueString:   "",
			PlaceHolderIndex: true,
			IPLiteralPrefix:  "inet '",
			PpLiteralSuffix:  "'",
			SelectIPPrefix:   `"`,
			SelectIPSuffix:   `"`,
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
	IPLiteralPrefix  string
	PpLiteralSuffix  string
	SelectIPPrefix   string
	SelectIPSuffix   string
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

func (d *Dialect) literalIP(ip net.IP) string {
	return d.IPLiteralPrefix + StringIPv6(ip) + d.PpLiteralSuffix
}

func (d *Dialect) selectIP(columnName string) string {
	return d.SelectIPPrefix + columnName + d.SelectIPSuffix
}

//Columns represents column index.
type Columns map[string]int

func (qb *QueryBuilder) buildFilterParts(ctx *queryContext, column string, filterValues []string) string {
	var where string
	if len(filterValues) == 1 {
		ctx.values = append(ctx.values, filterValues[0])
		where = column + " = " + qb.placeholder(len(ctx.values))
	} else {
		var filterQuery bytes.Buffer
		writeStrings(&filterQuery, column, " in (")

		last := len(filterValues) - 1
		for _, value := range filterValues[:last] {
			ctx.values = append(ctx.values, value)
			writeStrings(&filterQuery, qb.placeholder(len(ctx.values)), ",")
		}
		ctx.values = append(ctx.values, filterValues[last])
		writeStrings(&filterQuery, qb.placeholder(len(ctx.values)), ")")

		where = filterQuery.String()
	}
	return where
}

func (qb *QueryBuilder) join(fromTable, fromProperty, toTable string) string {
	return "left join " + qb.quote(fromTable) + " on " +
		qb.quote(toTable, "uuid") + " = " + qb.quote(fromTable, fromProperty)
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
		column := qb.quote(qb.TableAlias, filter.Key)
		where := qb.buildFilterParts(ctx, column, filter.Values)
		ctx.where = append(ctx.where, where)
	}
	// use join if backrefuuids
	if len(spec.BackRefUUIDs) > 0 {
		where := []string{}
		for backrefTable := range qb.BackRefFields {
			refTable := schema.ReferenceTableName(schema.RefPrefix, backrefTable, qb.Table)
			ctx.joins = append(ctx.joins, qb.join(refTable, "to", qb.TableAlias))
			wherePart := qb.buildFilterParts(ctx, qb.quote(refTable, "from"), spec.BackRefUUIDs)
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
		where = append(where, qb.quote(qb.TableAlias, "owner")+" = "+qb.placeholder(len(ctx.values)))
	}
	if spec.Shared {
		shareTables := []string{"domain_share_" + qb.Table, "tenant_share_" + qb.Table}
		for i, shareTable := range shareTables {
			ctx.joins = append(ctx.joins,
				qb.join(shareTable, "uuid", qb.TableAlias))
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
	writeString(query, "select ")

	if len(ctx.columnParts) != len(ctx.columns) {
		log.Fatal("unmatch")
	}
	writeStrings(query, strings.Join(ctx.columnParts, ","), " from ", qb.as(qb.Table, qb.TableAlias), " ")

	if len(ctx.joins) > 0 {
		writeString(query, strings.Join(ctx.joins, " "))
	}
	if len(ctx.where) > 0 {
		writeStrings(query, " where ", strings.Join(ctx.where, " and "))
	}
	if spec.Shared || len(spec.BackRefUUIDs) > 0 {
		writeStrings(query, " group by ", qb.quote(qb.TableAlias, "uuid"))
	}
	writeString(query, " ")
	if spec.Limit > 0 {
		writeStrings(
			query,
			" limit ",
			strconv.FormatInt(spec.Limit, 10),
			" offset ",
			strconv.FormatInt(spec.Offset, 10),
			" ",
		)
	}
}

func (qb *QueryBuilder) islinkToInField(ctx *queryContext, linkTo string) bool {
	spec := ctx.spec
	if len(spec.Fields) == 0 {
		return true
	}
	for _, field := range spec.Fields {
		if field == linkTo {
			return true
		}
	}
	return false
}

func (qb *QueryBuilder) buildRefQuery(ctx *queryContext) {
	spec := ctx.spec
	if !spec.Detail {
		return
	}
	for linkTo, refFields := range qb.RefFields {
		if !qb.islinkToInField(ctx, linkTo+"_refs") {
			continue
		}
		refTable := schema.ReferenceTableName(schema.RefPrefix, qb.Table, linkTo)
		refFields = append(refFields, "from")
		refFields = append(refFields, "to")
		subQuery := "(select " +
			qb.as(qb.jsonAgg(refTable+"_t", refFields...), qb.quote(refTable+"_ref")) +
			" from " + qb.as(qb.quote(refTable), refTable+"_t") +
			" where " + qb.quote(qb.TableAlias, "uuid") + " = " + qb.quote(refTable+"_t", "from") +
			" group by " + qb.quote(refTable+"_t", "from") + " )"
		ctx.columnParts = append(
			ctx.columnParts,
			subQuery)
		ctx.columns["ref_"+linkTo] = len(ctx.columns)
	}
}

func (qb *QueryBuilder) buildChildQuery(ctx *queryContext) {
	spec := ctx.spec
	if !spec.Detail {
		return
	}
	for child, childFields := range qb.ChildFields {
		if !qb.islinkToInField(ctx, child+"s") {
			continue
		}
		child = strings.ToLower(child)
		subQuery := "(select " +
			qb.as(qb.jsonAgg(child+"_t", childFields...), qb.quote(child+"_ref")) +
			" from " + qb.as(qb.quote(child), child+"_t") +
			" where " + qb.quote(qb.TableAlias, "uuid") + " = " + qb.quote(child+"_t", "parent_uuid") +
			" group by " + qb.quote(child+"_t", "parent_uuid") + " )"
		ctx.columnParts = append(
			ctx.columnParts,
			subQuery)
		ctx.columns[schema.ChildColumnName(child, qb.Table)] = len(ctx.columns)
	}
}

func (qb *QueryBuilder) buildBackRefQuery(ctx *queryContext) {
	spec := ctx.spec
	if !spec.Detail {
		return
	}
	for backrefTable, backrefFields := range qb.BackRefFields {
		if !qb.islinkToInField(ctx, backrefTable+"backrefs") {
			continue
		}
		refTable := schema.ReferenceTableName(schema.RefPrefix, backrefTable, qb.Table)
		backrefTable = strings.ToLower(backrefTable)
		subQuery := "(select " +
			qb.as(qb.jsonAgg(backrefTable+"_t", backrefFields...), qb.quote(refTable+"_backref")) +
			" from " + qb.as(qb.quote(backrefTable), backrefTable+"_t") +
			" inner join " + qb.as(refTable, refTable+"_t") +
			" on " + qb.quote(refTable+"_t", "from") + " = " + qb.quote(backrefTable+"_t", "uuid") +
			" where " + qb.quote(refTable+"_t", "to") + " = " + qb.quote(qb.TableAlias, "uuid") + " )"
		ctx.columnParts = append(
			ctx.columnParts,
			subQuery)
		ctx.columns[schema.BackRefColumnName(backrefTable, qb.Table)] = len(ctx.columns)
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
			ctx.columnParts = append(ctx.columnParts, qb.anyValue(qb.TableAlias, column))
		}
	} else {
		for _, column := range fields {
			ctx.columns[column] = len(ctx.columns)
			ctx.columnParts = append(ctx.columnParts, qb.quote(qb.TableAlias, column))
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
	qb.buildChildQuery(ctx)
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

//CreateRefQuery makes a reference.
func (qb *QueryBuilder) CreateRefQuery(linkTo string) string {
	fields := append([]string{"from", "to"}, qb.RefFields[linkTo]...)
	table := schema.ReferenceTableName(schema.RefPrefix, qb.Table, linkTo)
	return ("insert into " + table +
		" (" + qb.quoteSep(fields...) + ") values (" + qb.values(fields...) + ")")
}

//CreateParentRefQuery makes a reference to parent object.
func (qb *QueryBuilder) CreateParentRefQuery(linkTo string) string {
	fields := []string{"from", "to"}
	table := schema.ReferenceTableName(schema.ParentPrefix, qb.Table, linkTo)
	return ("insert into " + table +
		" (" + qb.quoteSep(fields...) + ") values (" + qb.values(fields...) + ")")
}

//DeleteQuery makes sql query.
func (qb *QueryBuilder) DeleteQuery() string {
	return "delete from " + qb.quote(qb.Table) + " where uuid = " + qb.placeholder(1)
}

//DeleteRefQuery makes sql query.
func (qb *QueryBuilder) DeleteRefQuery(linkTo string) string {
	table := schema.ReferenceTableName(schema.RefPrefix, qb.Table, linkTo)
	return "delete from " + table + " where " + qb.quote("from") + " = " + qb.placeholder(1)
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
func (qb *QueryBuilder) UpdateQuery(columns []string) string {
	var query bytes.Buffer
	writeStrings(&query, "update ", qb.quote(qb.Table), "set ")
	for i, column := range columns {
		writeStrings(&query, qb.quote(column), " = ", qb.placeholder(i+1))
		if i < len(columns)-1 {
			writeString(&query, ", ")
		}
	}
	writeStrings(&query, " where uuid = ", qb.placeholder(len(columns)+1))
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

// StringIPv6 serializes ip address, forces ipv6 format.
func StringIPv6(ip net.IP) string {
	if ip == nil || len(ip) == 0 {
		return ""
	}
	if ip.To4() == nil {
		return ip.String()
	}

	res := make(net.IP, len(ip))
	copy(res, ip)
	res = res.To16()
	res[1] = 1
	return res.String()[1:]
}

// writeStrings writes multiple strings to given buffer.
func writeStrings(b *bytes.Buffer, strings ...string) {
	for _, s := range strings {
		writeString(b, s)
	}
}

// writeString calls bytes.Buffer.WriteString() and strips its signature from redundant error,
// which  in current implementation is always nil.
// See: https://golang.org/pkg/bytes/#Buffer.WriteString
func writeString(b *bytes.Buffer, s string) {
	_, err := b.WriteString(s)
	if err != nil {
		log.Fatalf("unexpected bytes.Buffer.WriteString() error: %v", err)
	}
}
