package common

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	retryDB     = 10
	retryDBWait = 10
)

//ListSpec is configuraion option for select query.
type ListSpec struct {
	Table           string
	Filter          Filter
	Limit           int
	Offset          int
	Detail          bool
	Count           bool
	Shared          bool
	ExcludeHrefs    bool
	ParentFQName    []string
	ParentType      string
	ParentUUIDs     []string
	BackRefUUIDs    []string
	ObjectUUIDs     []string
	RequestedFields []string
	Fields          []string
	RefFields       map[string][]string
	BackRefFields   map[string][]string
	Auth            *AuthContext
	Values          []interface{}
	Columns         Columns
	columnParts     []string
	where           []string
	joins           []string
	groupBy         []string
	query           *bytes.Buffer
}

//Init initializes ListSpec.
func (spec *ListSpec) Init() {
	var query bytes.Buffer
	spec.query = &query
	spec.Columns = Columns{}
	spec.columnParts = []string{}
	spec.Values = []interface{}{}
	spec.where = []string{}
	spec.joins = []string{}
	spec.groupBy = []string{}
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

func (spec *ListSpec) buildFilterParts(column string, filterValues []string) string {
	var where string
	if len(filterValues) == 1 {
		where = column + " = ?"
		spec.Values = append(spec.Values, filterValues[0])
	} else {
		var filterQuery bytes.Buffer
		filterQuery.WriteString(column)
		filterQuery.WriteString(" in (")
		last := len(filterValues) - 1
		for _, value := range filterValues[:last] {
			filterQuery.WriteString("?,")
			spec.Values = append(spec.Values, value)
		}
		filterQuery.WriteString("?)")

		where = filterQuery.String()
		spec.Values = append(spec.Values, filterValues[last])
	}
	return where
}

func (spec *ListSpec) buildFilterQuery() {
	filter := spec.Filter
	filter.AppendValues("uuid", spec.ObjectUUIDs)
	filter.AppendValues("parent_uuid", spec.ParentUUIDs)
	if spec.ParentType != "" {
		filter.AppendValues("parent_type", []string{spec.ParentType})
	}
	for key, filterValues := range filter {
		if !spec.isValidField(key) {
			continue
		}
		column := fmt.Sprintf("`%s`.`%s`", spec.Table, key)
		where := spec.buildFilterParts(column, filterValues)
		spec.where = append(spec.where, where)
	}
	if len(spec.BackRefUUIDs) > 0 {
		where := []string{}
		for refTable := range spec.BackRefFields {
			column := fmt.Sprintf("`%s`.`uuid`", refTable)
			wherePart := spec.buildFilterParts(column, spec.BackRefUUIDs)
			where = append(where, wherePart)
		}
		spec.where = append(spec.where, "("+strings.Join(where, " or ")+")")
	}
}

func (spec *ListSpec) buildAuthQuery() {
	auth := spec.Auth
	where := []string{}

	if !auth.IsAdmin() {
		where = append(where, fmt.Sprintf("`%s`.`owner` = ?", spec.Table))
		spec.Values = append(spec.Values, auth.ProjectID())
	}
	if spec.Shared {
		shareTables := []string{"domain_share_" + spec.Table, "tenant_share_" + spec.Table}
		for _, shareTable := range shareTables {
			spec.joins = append(spec.joins,
				fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`uuid`",
					shareTable,
					spec.Table,
					shareTable,
				))
			spec.groupBy = append(spec.groupBy, spec.Table+".`uuid`")
			where = append(where, fmt.Sprintf("(`%s`.to = ? and `%s`.access >= 4)", shareTable, shareTable))
		}
		spec.Values = append(spec.Values, auth.DomainID(), auth.ProjectID())
	}
	if len(where) > 0 {
		spec.where = append(spec.where, fmt.Sprintf("(%s)", strings.Join(where, " or ")))
	}
}

func (spec *ListSpec) buildQuery() {
	query := spec.query

	query.WriteString("select ")
	if len(spec.columnParts) != len(spec.Columns) {
		log.Fatal("unmatch")
	}
	query.WriteString(strings.Join(spec.columnParts, ","))
	query.WriteString(" from ")
	query.WriteString(spec.Table)
	query.WriteRune(' ')
	if len(spec.joins) > 0 {
		query.WriteString(strings.Join(spec.joins, " "))
	}
	if len(spec.where) > 0 {
		query.WriteString(" where ")
		query.WriteString(strings.Join(spec.where, " and "))
	}
	if len(spec.groupBy) > 0 {
		query.WriteString(" group by ")
		query.WriteString(strings.Join(spec.groupBy, ","))
	}
	query.WriteRune(' ')
	pagenationQuery := fmt.Sprintf(" limit %d offset %d ", spec.Limit, spec.Offset)
	query.WriteString(pagenationQuery)
}

func (spec *ListSpec) buildRefQuery() {
	if !spec.Detail {
		return
	}
	for linkTo, refFields := range spec.RefFields {
		refColumns := []string{}
		refTable := "ref_" + spec.Table + "_" + linkTo
		refFields = append(refFields, "from")
		refFields = append(refFields, "to")
		for _, field := range refFields {
			refColumns = append(refColumns, fmt.Sprintf("'%s', `%s`.`%s`", field, refTable, field))
		}
		spec.columnParts = append(spec.columnParts, fmt.Sprintf("group_concat(distinct JSON_OBJECT(%s)) as `%s_ref`", strings.Join(refColumns, ","), refTable))
		spec.Columns["ref_"+linkTo] = len(spec.Columns)
		spec.joins = append(spec.joins,
			fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`from`",
				refTable,
				spec.Table,
				refTable,
			))
		spec.groupBy = append(spec.groupBy, refTable+".`from`")
	}
}

func (spec *ListSpec) buildBackRefQuery() {
	for refTable, refFields := range spec.BackRefFields {
		refColumns := []string{}
		for _, field := range refFields {
			refColumns = append(refColumns, fmt.Sprintf("'%s', `%s`.`%s`", field, refTable, field))
		}
		if spec.Detail {
			spec.columnParts = append(spec.columnParts, fmt.Sprintf("group_concat(distinct JSON_OBJECT(%s)) as `%s_ref`", strings.Join(refColumns, ","), refTable))
			spec.Columns["backref_"+refTable] = len(spec.Columns)
		}
		if spec.Detail || len(spec.BackRefUUIDs) > 0 {
			spec.joins = append(spec.joins,
				fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`parent_uuid`",
					refTable,
					spec.Table,
					refTable,
				))
			spec.groupBy = append(spec.groupBy, refTable+".`uuid`")
		}
	}

}

func (spec *ListSpec) isValidField(requestedField string) bool {
	for _, field := range spec.Fields {
		if field == requestedField {
			return true
		}
	}
	return false
}

func (spec *ListSpec) checkRequestedFields() bool {
	if len(spec.RequestedFields) == 0 {
		return false
	}
	for _, requestedField := range spec.RequestedFields {
		if !spec.isValidField(requestedField) {
			return false
		}
	}
	return true
}

func (spec *ListSpec) buildColumns() {
	columnTemplate := "`%s`.`%s`"
	if spec.Detail || len(spec.BackRefUUIDs) > 0 || spec.Shared {
		columnTemplate = "ANY_VALUE(`%s`.`%s`)"
	}

	fields := spec.Fields

	if spec.checkRequestedFields() {
		fields = spec.RequestedFields
	}

	for _, column := range fields {
		spec.Columns[column] = len(spec.Columns)
		spec.columnParts = append(spec.columnParts, fmt.Sprintf(columnTemplate, spec.Table, column))
	}
}

//BuildQuery makes sql query.
func (spec *ListSpec) BuildQuery() string {
	spec.Init()
	spec.buildColumns()
	spec.buildFilterQuery()
	spec.buildAuthQuery()
	spec.buildRefQuery()
	spec.buildBackRefQuery()
	spec.buildQuery()
	return spec.query.String()
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
