package schema

import (
	"github.com/anemiq/anemiq/database"
	"github.com/graphql-go/graphql"
)

//ForTables returns GraphQL schema for given tables
func ForTables(db *database.Database, tables []*database.Table) graphql.Schema {
	rootQueryFields := graphql.Fields{}
	for _, table := range tables {
		rootQueryFields[table.Name] = buildTableField(table, buildTableType(table))
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: rootQueryFields,
		}),
	})

	if err != nil {
		panic(err)
	}

	return schema
}

func buildTableType(t *database.Table) graphql.Type {
	fields := graphql.Fields{}
	for _, col := range t.Cols {
		fields[col.Name] = &graphql.Field{
			Type: col.ColType,
		}
	}
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   t.Name,
		Fields: fields,
	})
}

func buildTableField(table *database.Table, tableType graphql.Type) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(tableType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return table.SelectAll(), nil
		},
	}
}
