package database

import (
	"testing"

	"github.com/anemiq/config"
	"github.com/anemiq/test"
	"github.com/graphql-go/graphql"
)

func TestBuildDataSourceName(t *testing.T) {
	conn := config.Conn{
		Host:     "localhost",
		Port:     "3306",
		Database: "mydb",
		User:     "user",
		Pass:     "pass"}
	dataSource := buildDataSourceName(conn)
	test.AssertEqual(t, dataSource, "user:pass@tcp(localhost:3306)/mydb")
}

func TestConversionToGraphQLTypes(t *testing.T) {
	test.AssertEqual(t, buildGraphQlType("varchar(50)"), graphql.String)
	test.AssertEqual(t, buildGraphQlType("char(4)"), graphql.String)

	test.AssertEqual(t, buildGraphQlType("integer"), graphql.Int)
	test.AssertEqual(t, buildGraphQlType("smallint"), graphql.Int)
	
	test.AssertEqual(t, buildGraphQlType("numeric"), graphql.Float)
	test.AssertEqual(t,buildGraphQlType("float"), graphql.Float)
	test.AssertEqual(t,buildGraphQlType("double"), graphql.Float)
	test.AssertEqual(t,buildGraphQlType("dec"), graphql.Float)
	test.AssertEqual(t,buildGraphQlType("fixed"), graphql.Float)
}