package database

import (
	"github.com/anemiq/anemiq/gqltype"
	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
)

type Table struct {
	Db   *Database
	Name string
	Cols []Column
}

func (t *Table) SelectAll() interface{} {

	rows, err := t.Db.db.Query("select * from " + t.Name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	registers := []map[string]interface{}{}

	for rows.Next() {
		columns := make([]string, len(cols))
		columnsPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnsPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnsPointers...); err != nil {
			panic(err)
		}

		m := make(map[string]interface{})

		for i, colName := range cols {
			val := columnsPointers[i].(*string)
			m[colName] = *val
		}

		registers = append(registers, m)

	}
	return registers
}

type Column struct {
	Name    string
	ColType *graphql.Scalar
}

func newTable(db *Database, name string) (*Table, error) {
	rows, err := db.db.Query("DESCRIBE " + name)
	if err != nil {
		return nil, err
	}

	var field, colType, allowNull, key, isDefault, extra string
	var cols []Column

	for rows.Next() {
		rows.Scan(&field, &colType, &allowNull, &key, &isDefault, &extra)
		col := newColumn(field, colType)
		cols = append(cols, col)
	}
	return &Table{db, name, cols}, nil
}

func newColumn(name, colTypeStr string) Column {
	return Column{name, gqltype.FromColType(colTypeStr)}
}
