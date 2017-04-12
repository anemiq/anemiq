package gql

import (
	"github.com/anemiq/database"
	"github.com/graphql-go/graphql"
)

func SchemaForTables(db *database.Database, tables []*database.Table) graphql.Schema {
	rootQueryFields := graphql.Fields{}
	for _, table := range tables {
		tableType := buildTableType(table)
		tableField := &graphql.Field{
			Type: graphql.NewList(tableType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return table.SelectAll(), nil
			},
		}
		rootQueryFields[table.Name] = tableField
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: rootQueryFields,
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
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
