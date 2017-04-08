package main

import (
	"log"
	"os"

	"github.com/anemiq/config"
)

func main() {

	_, err := config.Read()

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

}
