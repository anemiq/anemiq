package main

import (
	"fmt"
	"os"

	"github.com/anemiq/config"
	"github.com/anemiq/database"
	"github.com/anemiq/schema"
	"github.com/anemiq/server"
)

func log(msg string) {
	fmt.Printf(msg + "\n")
}

func main() {

	log("reading config from anemiq.yml...")
	conf, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	log("done!")

	log("connecting to database...")
	db, err := database.Open(conf.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log("done!")

	//Collect tables
	tables, err := db.Tables(conf.Tables)
	if err != nil {
		panic(err)
	}

	//Generate GraphQL schema
	sch := schema.ForTables(tables)

	sv := server.New(conf, sch)
	log("server started at http://localhost:" + conf.Server.Port + "/graphql")
	sv.Start()
}
