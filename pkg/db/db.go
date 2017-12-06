package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

//ListSpec is configuraion option for select query.
type ListSpec struct {
	Table     string
	Filter    map[string]interface{}
	Limit     int
	Offset    int
	Detail    bool
	Fields    []string
	RefFields map[string][]string
}

//Columns represents column index
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
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = do(tx)
	return err
}

//BuildListQuery makes query using list spec.
func BuildListQuery(spec *ListSpec) (string, Columns, []interface{}) {
	var query bytes.Buffer
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
	pagenationQuery := fmt.Sprintf(" limit %d offset %d ", spec.Limit, spec.Offset)

	for key, value := range spec.Filter {
		where = append(where, fmt.Sprintf("`%s`.`%s` = ?", spec.Table, key))
		values = append(values, value)
	}
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
		query.WriteString(strings.Join(where, ","))
	}
	if len(groupBy) > 0 {
		query.WriteString(" group by ")
		query.WriteString(strings.Join(groupBy, ","))
	}
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	return query.String(), columns, values
}
