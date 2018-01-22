package main

import (
	"fmt"
	"os"

	"github.com/anemiq/config"
	"github.com/anemiq/database"
	"github.com/anemiq/schema"
	"github.com/anemiq/server"
)

func main() {

	fmt.Printf("reading config from anemiq.yml..." + "\n")
	conf, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("done" + "\n")

	fmt.Printf("connecting to database..." + "\n")
	db, err := database.Open(conf.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Printf("done" + "\n")

	tables, err := db.Tables(conf.Tables)
	if err != nil {
		panic(err)
	}

	sch := schema.ForTables(tables)

	sv := server.New(conf, sch)
	fmt.Printf("server started at http://localhost:" + conf.Server.Port + "/graphql" + "\n")
	sv.Start()
}
