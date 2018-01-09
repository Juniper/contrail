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
	Table         string
	Filter        Filter
	Limit         int
	Offset        int
	Detail        bool
	Count         bool
	Shared        bool
	ExcludeHrefs  bool
	ParentFQName  []string
	ParentType    string
	ParentUUIDs   []string
	BackrefUUIDs  []string
	ObjectUUIDs   []string
	Fields        []string
	RefFields     map[string][]string
	BackRefFields map[string][]string
	Auth          *AuthContext
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
			err = tx.Commit()
		}
	}()
	err = do(tx)
	return err
}

func buildFilterQuery(spec *ListSpec, where []string, values []interface{}) ([]string, []interface{}) {
	for key, filterValues := range spec.Filter {
		if len(filterValues) == 1 {
			where = append(where, fmt.Sprintf("`%s`.`%s` = ?", spec.Table, key))
			values = append(values, filterValues[0])
		} else {
			var filterQuery bytes.Buffer
			filterQuery.WriteString("`%s`.`%s` in (")
			last := len(filterValues) - 1
			for _, value := range filterValues[:last] {
				filterQuery.WriteString("?,")
				values = append(values, value)
			}
			filterQuery.WriteString("?)")

			where = append(where, filterQuery.String())
			values = append(values, filterValues[last])
		}
	}
	return where, values
}

func buildAuthQuery(spec *ListSpec, where []string, values []interface{}) ([]string, []interface{}) {
	auth := spec.Auth
	if !auth.IsAdmin() {
		where = append(where, fmt.Sprintf("`%s`.`%s` = ?", spec.Table, "owner"))
		values = append(values, auth.ProjectID())
	}
	return where, values
}

func buildQuery(spec *ListSpec, columns Columns, columnParts []string, where []string, values []interface{}, joins []string, groupBy []string) (string, Columns, []interface{}) {
	var query bytes.Buffer

	where, values = buildFilterQuery(spec, where, values)
	where, values = buildAuthQuery(spec, where, values)

	query.WriteString("select ")
	if len(columnParts) != len(columns) {
		log.Fatal("unmatch")
	}
	query.WriteString(strings.Join(columnParts, ","))
	query.WriteString(" from ")
	query.WriteString(spec.Table)
	query.WriteRune(' ')
	if len(joins) > 0 {
		query.WriteString(strings.Join(joins, " "))
	}
	if len(where) > 0 {
		query.WriteString(" where ")
		query.WriteString(strings.Join(where, " and "))
	}
	if len(groupBy) > 0 {
		query.WriteString(" group by ")
		query.WriteString(strings.Join(groupBy, ","))
	}
	query.WriteRune(' ')
	pagenationQuery := fmt.Sprintf(" limit %d offset %d ", spec.Limit, spec.Offset)
	query.WriteString(pagenationQuery)
	return query.String(), columns, values
}

func buildSimpleQuery(spec *ListSpec) (string, Columns, []interface{}) {
	values := []interface{}{}
	columns := Columns{}
	columnParts := []string{}
	where := []string{}

	for _, column := range spec.Fields {
		columns[column] = len(columns)
		columnParts = append(columnParts, fmt.Sprintf("`%s`.`%s`", spec.Table, column))
	}

	return buildQuery(spec, columns, columnParts, where, values, nil, nil)
}

func buildDetailQuery(spec *ListSpec) (string, Columns, []interface{}) {
	values := []interface{}{}
	columns := Columns{}
	columnParts := []string{}
	where := []string{}
	joins := []string{}
	groupBy := []string{}
	for _, column := range spec.Fields {
		columns[column] = len(columns)
		columnParts = append(columnParts, fmt.Sprintf("ANY_VALUE(`%s`.`%s`)", spec.Table, column))
	}

	for linkTo, refFields := range spec.RefFields {
		refColumns := []string{}
		refTable := "ref_" + spec.Table + "_" + linkTo
		refFields = append(refFields, "from")
		refFields = append(refFields, "to")
		for _, field := range refFields {
			refColumns = append(refColumns, fmt.Sprintf("'%s', `%s`.`%s`", field, refTable, field))
		}
		columnParts = append(columnParts, fmt.Sprintf("group_concat(distinct JSON_OBJECT(%s)) as `%s_ref`", strings.Join(refColumns, ","), refTable))
		columns["ref_"+linkTo] = len(columns)
		joins = append(joins,
			fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`from`",
				refTable,
				spec.Table,
				refTable,
			))
		groupBy = append(groupBy, refTable+".`from`")
	}
	for refTable, refFields := range spec.BackRefFields {
		refColumns := []string{}
		for _, field := range refFields {
			refColumns = append(refColumns, fmt.Sprintf("'%s', `%s`.`%s`", field, refTable, field))
		}
		columnParts = append(columnParts, fmt.Sprintf("group_concat(distinct JSON_OBJECT(%s)) as `%s_ref`", strings.Join(refColumns, ","), refTable))
		columns["backref_"+refTable] = len(columns)
		joins = append(joins,
			fmt.Sprintf("left join `%s` on `%s`.`uuid` = `%s`.`parent_uuid`",
				refTable,
				spec.Table,
				refTable,
			))
		groupBy = append(groupBy, refTable+".`uuid`")
	}

	return buildQuery(spec, columns, columnParts, where, values, joins, groupBy)
}

//BuildListQuery makes query using list spec.
func BuildListQuery(spec *ListSpec) (string, Columns, []interface{}) {
	if spec.Detail {
		return buildDetailQuery(spec)
	}
	return buildSimpleQuery(spec)
}

//SetLogLevel set global log level using viper configuraion.
func SetLogLevel() {
	logLevel := viper.GetString("log_level")
	switch logLevel {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
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
