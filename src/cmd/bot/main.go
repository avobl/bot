package main

import (
	"github.com/avobl/bot/src/config"
	"log"

	"github.com/avobl/bot/src/db/sqlite"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("config: %v\n", err)
	}

	db := sqlite.GetProvider(&sqlite.Config{
		DBName:       conf.SQLite.Dbname,
		MaxIdleConns: conf.SQLite.MaxIdleConns,
		MaxOpenConns: conf.SQLite.MaxOpenConns,
	})
	if err = db.Init(); err != nil {
		log.Fatalf("sqlite init: %v\n", err)
	}

	defer func() { _ = db.Close() }()

}
