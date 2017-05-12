package gqltype

import (
	"testing"

	"github.com/anemiq/anemiq/test"
	"github.com/graphql-go/graphql"
)

func TestConversionToGraphQLTypes(t *testing.T) {
	test.AssertEqual(t, FromColType("varchar"), graphql.String)
	test.AssertEqual(t, FromColType("char"), graphql.String)

	test.AssertEqual(t, FromColType("integer"), graphql.Int)
	test.AssertEqual(t, FromColType("int"), graphql.Int)
	test.AssertEqual(t, FromColType("smallint"), graphql.Int)

	test.AssertEqual(t, FromColType("numeric"), graphql.Float)
	test.AssertEqual(t, FromColType("float"), graphql.Float)
	test.AssertEqual(t, FromColType("double"), graphql.Float)
	test.AssertEqual(t, FromColType("dec"), graphql.Float)
	test.AssertEqual(t, FromColType("fixed"), graphql.Float)
}
