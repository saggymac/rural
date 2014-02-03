package storage 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	)

type Config struct {
	Datasource string "datasource"
}

var config = Config{}
var db *sql.DB

func Initialize( cnf Config ) {
	config = cnf

	var err error
	db, err = sql.Open( "sqlite3", config.Datasource)
	if err != nil {
		panic( err)
	}

	defer db.Close()
}
