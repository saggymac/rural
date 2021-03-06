package storage 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"errors"
	"io/ioutil"
	"os"
	)

type Config struct {
	Datasource string "datasource"
}

var config = Config{}
var db *sql.DB = nil

func readAndExecSqlFromFile( sqlFilePath string ) error {

	var fullPath = ""

	envHome := os.Getenv( "RURAL_HOME")
	if len( envHome) <= 0 {
		fullPath = sqlFilePath		
	} else {
		fullPath = envHome + "/conf/" + sqlFilePath
	}
		

	dat, err := ioutil.ReadFile( fullPath)
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

	// Validate 
	if len( deviceId) <= 0 {
		return errors.New( "invalid deviceId")
	}

	if len( appId) <= 0 {
		return errors.New( "invalid appId")
	}

	if len( appVersion) <= 0 {
		return errors.New( "invalid appVersion")
	}


	// Update the DB
	_, err := db.Exec( "INSERT OR REPLACE INTO devices (appId,deviceId,appVersion) VALUES (?,?,?)", appId, deviceId, appVersion)
	if err != nil {
		return err
	}


	return nil
}