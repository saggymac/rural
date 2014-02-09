package storage 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"errors"
	"io/ioutil"
	)

type Config struct {
	Datasource string "datasource"
}

var config = Config{}
var db *sql.DB = nil

func readAndExecSqlFromFile( sqlFilePath string ) error {
	dat, err := ioutil.ReadFile( sqlFilePath)
	if err != nil {
		return err
	}

	_, err = db.Exec( string( dat))
	if err != nil {
		return err
	}

	return nil	
}


func setupTablesIfNecessary() error {

	err := readAndExecSqlFromFile( "devices.sql")
	if err != nil {
		return err
	}

	err = readAndExecSqlFromFile( "messages.sql")
	if err != nil {
		return err
	}

	return nil
}


func Shutdown() {
	if db != nil {
		db.Close()
	}
}


func Initialize( cnf Config ) {
	config = cnf

	var err error
	db, err = sql.Open( "sqlite3", config.Datasource)
	if err != nil {
		panic( err)
	}

	err = setupTablesIfNecessary()
	if err != nil {
		panic( err)
	}

}


// We either create a new record with these values, or we
// update an existing one.
// 
func UpdateDevice( appId string, deviceId string, appVersion string ) error {
	if db == nil {
		return errors.New( "no database")
	}

	// todo: more validation
	if len( deviceId) <= 0 {
		return errors.New( "invalid deviceId")
	}

	_, err := db.Exec( "INSERT OR REPLACE INTO devices (appId,deviceId,appVersion) VALUES (?,?,?)", appId, deviceId, appVersion)
	if err != nil {
		return err
	}

	return nil
}