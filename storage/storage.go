package storage 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"errors"
	)

type Config struct {
	Datasource string "datasource"
}

var config = Config{}
var db *sql.DB = nil

func Initialize( cnf Config ) {
	config = cnf

	var err error
	db, err = sql.Open( "sqlite3", config.Datasource)
	if err != nil {
		panic( err)
	}

	// TODO: create schema if necessary

	defer db.Close()
}


// Create a new device record, or update its timestamp
func UpdateDevice( deviceId string ) error {
	if db == nil {
		return errors.New( "no database")
	}

	if len( deviceId) <= 0 {
		return errors.New( "invalid deviceId")
	}

	_, err := db.Exec( "INSERT OR REPLACE INTO devices (deviceId) VALUES (?)", deviceId)
	if err != nil {
		return errors.New( "unable to insert device into db")
	}

	return nil
}