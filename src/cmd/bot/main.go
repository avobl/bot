package main

import (
	"log"

	"github.com/avobl/bot/src/db/sqlite"
)

func main() {
	db := sqlite.GetProvider(&sqlite.Config{})
	err := db.Init()
	if err != nil {
		log.Fatalf("sqlite init: %v\n", err)
	}

	defer func() { _ = db.Close() }()

}
