package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	driverName     = flag.String("driver", "mysql", "database driver name")
	dataSourceName = flag.String("source", "root@(localhost:3306)/", "data source name")
)

func main() {
	flag.Parse()
	db, err := sql.Open(*driverName, *dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("stats:%+v\n", db.Stats())
}
