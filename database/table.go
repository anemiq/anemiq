package database

import _ "github.com/go-sql-driver/mysql"
import "github.com/graphql-go/graphql"
import "fmt"
import "strings"

type Table struct {
	Name string
	Cols []Column
}

type Column struct {
	Name string
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

	var so *graphql.Scalar
	so = graphql.Int
	fmt.Print(so)

	return &Table{name, cols}, nil
}

func newColumn(name, colTypeStr string) Column {
	return Column{name, buildGraphQlType(colTypeStr)}
}

func buildGraphQlType(colType string) *graphql.Scalar {
	if isChar(colType) {
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
	colTypeUpper := strings.ToUpper(colType)
	return strings.Contains(colTypeUpper, "INT")
}

func isFloat(colType string) bool {
	colTypeUpper := strings.ToUpper(colType)
	return strings.Contains(colTypeUpper, "DEC") ||  strings.Contains(colTypeUpper, "FIXED") || 
		strings.Contains(colTypeUpper, "NUMERIC") || strings.Contains(colTypeUpper, "FLOAT") || 
		strings.Contains(colTypeUpper, "DOUBLE")
}