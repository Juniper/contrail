package basedb

import "fmt"

// DatabaseData is a map where keys are table names and values are their contents.
type DatabaseData map[string]TableData

// ForEachRow calls f for each row of a database.
func (d DatabaseData) ForEachRow(f func(schemaID string, row RowData) error) error {
	for schemaID, table := range d {
		for _, row := range table {
			if err := f(schemaID, row); err != nil {
				return err
			}
		}
	}
	return nil
}

// RowsCount counts rows in a DatabaseData object.
func (d DatabaseData) RowsCount() (count int) {
	for _, table := range d {
		count += len(table)
	}
	return count
}

// TableData is a table representation in which rows are stored as maps.
type TableData []RowData

// RowData is a map representation of a DB row.
type RowData map[string]interface{}

// PK returns value of primary key for a given row data.
func (d RowData) PK() []string {
	return []string{fmt.Sprint(d[UUIDColumn])}
}
