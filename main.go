package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"io/ioutil"

	"github.com/anemiq/anemiq/config"
	"github.com/anemiq/anemiq/database"
	"github.com/anemiq/anemiq/schema"
	"github.com/graphql-go/graphql"
)

func main() {

	//Read configuration
	conf, err := config.Read()
	if err != nil {
		fatal(err)
	}

	//Open database
	db, err := database.Open(conf.Conn)
	if err != nil {
		fatal(err)
	}
	defer db.Close()

	//Collect tables
	tables := []*database.Table{}
	for _, tableName := range conf.Tables {
		table, _ := db.Table(tableName)
		tables = append(tables, table)
	}

	//Generate GraphQL schema
	sch := schema.ForTables(db, tables)

	//Run server
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		//Get query body
		//TODO improve error responses
		query, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(query) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Perform query
		params := graphql.Params{
			Schema:        sch,
			RequestString: string(query),
		}
		result := graphql.Do(params)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)

}

func fatal(err error) {
	log.Fatal(err.Error())
	os.Exit(1)
}
