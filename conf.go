package main

import (
	"encoding/json" 
	"fmt"
	"io/ioutil"
	"github.com/saggymac/rural/apns"
	"github.com/saggymac/rural/storage"	
	)


type Config struct {
	Apns apns.Config "apns"
	Storage storage.Config "storage"
}

var AppConfig = &Config{}

var parsedConfig map[string]interface{}

func init() {

	dat, err := ioutil.ReadFile( "conf.json")
	if err != nil {
		fmt.Printf( "unable to read conf.json; %v\n", err)
		return
	}

	err = json.Unmarshal( dat, &AppConfig)
	if err != nil {
		fmt.Printf( "unable to parse config file; %v\n", err)
		return
	}

	fmt.Println( "config loaded")
	return
}