package gqltype

import (
	"strings"

	"github.com/graphql-go/graphql"
)

func FromColType(colType string) *graphql.Scalar {
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
