package common

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	retryDB     = 10
	retryDBWait = 10
)

//ListQueryBuilder builds list query.
type ListQueryBuilder struct {
	Spec          *models.ListSpec
	Fields        []string
	Table         string
	RefFields     map[string][]string
	BackRefFields map[string][]string
	Auth          *AuthContext
	Values        []interface{}
	Columns       Columns
	columnParts   []string
	where         []string
	joins         []string
	groupBy       []string
	query         *bytes.Buffer
}

//Init initializes ListSpec.
func (qb *ListQueryBuilder) Init() {
	var query bytes.Buffer
	qb.query = &query
	qb.Columns = Columns{}
	qb.columnParts = []string{}
	qb.Values = []interface{}{}
	qb.where = []string{}
	qb.joins = []string{}
	qb.groupBy = []string{}
}

//Columns represents column index.
type Columns map[string]int

//DoInTransaction run a function inside of DB transaction
func DoInTransaction(db *sql.DB, do func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Error(err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = do(tx)
	return err
}

func (qb *ListQueryBuilder) buildFilterParts(column string, filterValues []string) string {
	var where string
	if len(filterValues) == 1 {
		where = column + " = ?"
		qb.Values = append(qb.Values, filterValues[0])
	} else {
		var filterQuery bytes.Buffer
		filterQuery.WriteString(column)
		filterQuery.WriteString(" in (")
		last := len(filterValues) - 1
		for _, value := range filterValues[:last] {
			filterQuery.WriteString("?,")
			qb.Values = append(qb.Values, value)
		}
		filterQuery.WriteString("?)")

		where = filterQuery.String()
		qb.Values = append(qb.Values, filterValues[last])
	}
	return where
}

func (qb *ListQueryBuilder) buildFilterQuery() {
	spec := qb.Spec
	filter := spec.Filter
	filter.AppendValues("uuid", spec.ObjectUUIDs)
	filter.AppendValues("parent_uuid", spec.ParentUUIDs)
	if spec.ParentType != "" {
		filter.AppendValues("parent_type", []string{spec.ParentType})
	}
	for key, filterValues := range filter {
		if !qb.isValidField(key) {
			continue
		}
		column := fmt.Sprintf("`%s`.`%s`", qb.Table, key)
		where := qb.buildFilterParts(column, filterValues)
		qb.where = append(qb.where, where)
	}
	if len(spec.BackRefUUIDs) > 0 {
		where := []string{}
		for refTable := range qb.BackRefFields {
			column := fmt.Sprintf("`%s`.`uuid`", refTable)
			wherePart := qb.buildFilterParts(column, spec.BackRefUUIDs)
			where = append(where, wherePart)
		}
		qb.where = append(qb.where, "("+strings.Join(where, " or ")+")")
	}
}

func (qb *ListQueryBuilder) buildAuthQuery() {
	auth := qb.Auth
	spec := qb.Spec
	where := []string{}

	if !auth.IsAdmin() {
		where = append(where, fmt.Sprintf("`%s`.`owner` = ?", qb.Table))
		qb.Values = append(qb.Values, auth.ProjectID())
	}
	if spec.Shared {
		shareTables := []string{"domain_share_" + qb.Table, "tenant_share_" + qb.Table}
		for _, shareTable := range shareTables {
			qb.joins = append(qb.joins,
				fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`uuid`",
					shareTable,
					qb.Table,
					shareTable,
				))
			qb.groupBy = append(qb.groupBy, qb.Table+".`uuid`")
			where = append(where, fmt.Sprintf("(`%s`.to = ? and `%s`.access >= 4)", shareTable, shareTable))
		}
		qb.Values = append(qb.Values, auth.DomainID(), auth.ProjectID())
	}
	if len(where) > 0 {
		qb.where = append(qb.where, fmt.Sprintf("(%s)", strings.Join(where, " or ")))
	}
}

func (qb *ListQueryBuilder) buildQuery() {
	spec := qb.Spec
	query := qb.query

	query.WriteString("select ")
	if len(qb.columnParts) != len(qb.Columns) {
		log.Fatal("unmatch")
	}
	query.WriteString(strings.Join(qb.columnParts, ","))
	query.WriteString(" from ")
	query.WriteString(qb.Table)
	query.WriteRune(' ')
	if len(qb.joins) > 0 {
		query.WriteString(strings.Join(qb.joins, " "))
	}
	if len(qb.where) > 0 {
		query.WriteString(" where ")
		query.WriteString(strings.Join(qb.where, " and "))
	}
	if len(qb.groupBy) > 0 {
		query.WriteString(" group by ")
		query.WriteString(strings.Join(qb.groupBy, ","))
	}
	query.WriteRune(' ')
	pagenationQuery := fmt.Sprintf(" limit %d offset %d ", spec.Limit, spec.Offset)
	query.WriteString(pagenationQuery)
}

func (qb *ListQueryBuilder) buildRefQuery() {
	spec := qb.Spec
	if !spec.Detail {
		return
	}
	for linkTo, refFields := range qb.RefFields {
		refColumns := []string{}
		refTable := "ref_" + qb.Table + "_" + linkTo
		refFields = append(refFields, "from")
		refFields = append(refFields, "to")
		for _, field := range refFields {
			refColumns = append(refColumns, fmt.Sprintf("'%s', `%s`.`%s`", field, refTable, field))
		}
		qb.columnParts = append(
			qb.columnParts,
			fmt.Sprintf(
				"group_concat(distinct JSON_OBJECT(%s)) as `%s_ref`",
				strings.Join(refColumns, ","),
				refTable,
			),
		)
		qb.Columns["ref_"+linkTo] = len(qb.Columns)
		qb.joins = append(qb.joins,
			fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`from`",
				refTable,
				qb.Table,
				refTable,
			))
		qb.groupBy = append(qb.groupBy, refTable+".`from`")
	}
}

func (qb *ListQueryBuilder) buildBackRefQuery() {
	spec := qb.Spec
	for refTable, refFields := range qb.BackRefFields {
		refColumns := []string{}
		for _, field := range refFields {
			refColumns = append(refColumns, fmt.Sprintf("'%s', `%s`.`%s`", field, refTable, field))
		}
		if spec.Detail {
			qb.columnParts = append(
				qb.columnParts,
				fmt.Sprintf(
					"group_concat(distinct JSON_OBJECT(%s)) as `%s_ref`",
					strings.Join(refColumns, ","),
					refTable,
				),
			)
			qb.Columns["backref_"+refTable] = len(qb.Columns)
		}
		if spec.Detail || len(spec.BackRefUUIDs) > 0 {
			qb.joins = append(qb.joins,
				fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`parent_uuid`",
					refTable,
					qb.Table,
					refTable,
				))
			qb.groupBy = append(qb.groupBy, refTable+".`uuid`")
		}
	}

}

func (qb *ListQueryBuilder) isValidField(requestedField string) bool {
	for _, field := range qb.Fields {
		if field == requestedField {
			return true
		}
	}
	return false
}

func (qb *ListQueryBuilder) checkRequestedFields() bool {
	spec := qb.Spec
	for _, requestedField := range spec.Fields {
		if !qb.isValidField(requestedField) {
			return false
		}
	}
	return true
}

func (qb *ListQueryBuilder) buildColumns() {
	spec := qb.Spec
	columnTemplate := "`%s`.`%s`"
	if spec.Detail || len(spec.BackRefUUIDs) > 0 || spec.Shared {
		columnTemplate = "ANY_VALUE(`%s`.`%s`)"
	}

	fields := qb.Fields

	if len(spec.Fields) > 0 && qb.checkRequestedFields() {
		fields = spec.Fields
	}

	for _, column := range fields {
		qb.Columns[column] = len(qb.Columns)
		qb.columnParts = append(qb.columnParts, fmt.Sprintf(columnTemplate, qb.Table, column))
	}
}

//BuildQuery makes sql query.
func (qb *ListQueryBuilder) BuildQuery() string {
	qb.Init()
	qb.buildColumns()
	qb.buildFilterQuery()
	qb.buildAuthQuery()
	qb.buildRefQuery()
	qb.buildBackRefQuery()
	qb.buildQuery()
	return qb.query.String()
}

//ConnectDB connect to the db based on viper configuration.
func ConnectDB() (*sql.DB, error) {
	databaseConnection := viper.GetString("database.connection")
	maxConn := viper.GetInt("database.max_open_conn")
	db, err := sql.Open("mysql", databaseConnection)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxConn)
	for i := 0; i < retryDB; i++ {
		err = db.Ping()
		if err == nil {
			log.Info("connected to the database")
			return db, nil
		}
		time.Sleep(retryDBWait * time.Second)
		log.Printf("Retrying db connection... (%s)", err)
	}
	return nil, fmt.Errorf("failed to open db connection")
}
