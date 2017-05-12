package schema

import (
	"github.com/anemiq/anemiq/database"
	"github.com/graphql-go/graphql"
)

//ForTables returns GraphQL schema for given tables
func ForTables(tables []*database.Table) graphql.Schema {
	rootQueryFields := graphql.Fields{}
	for _, table := range tables {
		rootQueryFields[table.Name] = buildTableField(table, buildTableType(table))
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "rootQuery",
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

func buildTableField(t *database.Table, tableType graphql.Type) *graphql.Field {
	args := graphql.FieldConfigArgument{}
	for _, col := range t.Cols {
		args[col.Name] = &graphql.ArgumentConfig{
			Type:        col.ColType,
			Description: col.Name,
		}
	}
	return &graphql.Field{
		Type: graphql.NewList(tableType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if shouldFilter(p) {
				return t.SelectWhere(p.Args)
			}
			return t.SelectAll()
		},
		Args: args,
	}
}

func shouldFilter(p graphql.ResolveParams) bool {
	for _ = range p.Args {
		return true
	}
	return false
}
