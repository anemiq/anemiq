package database

import _ "github.com/go-sql-driver/mysql"
import "github.com/graphql-go/graphql"

import "strings"

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
	return Column{name, buildGraphQlType(colTypeStr)}
}

func buildGraphQlType(colType string) *graphql.Scalar {
	if isChar(colType) || isDate(colType) {
		return graphql.String
	} else if isInt(colType) {
		return graphql.Int
	} else if isFloat(colType) {
		return graphql.Float
	}
	return nil
}

func isChar(colType string) bool {
	return strings.Contains(strings.ToUpper(colType), "CHAR")
}

func isInt(colType string) bool {
	return strings.Contains(strings.ToUpper(colType), "INT")
}

func isFloat(colType string) bool {
	colTypeUpper := strings.ToUpper(colType)
	return strings.Contains(colTypeUpper, "DEC") || strings.Contains(colTypeUpper, "FIXED") ||
		strings.Contains(colTypeUpper, "NUMERIC") || strings.Contains(colTypeUpper, "FLOAT") ||
		strings.Contains(colTypeUpper, "DOUBLE")
}

func isDate(colType string) bool {
	return strings.Contains(strings.ToUpper(colType), "TIME")
}
