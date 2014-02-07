package main

import (
	"github.com/saggymac/rural/server"
	"github.com/saggymac/rural/apns"
	"github.com/saggymac/rural/storage"	
	"fmt"
	)

func main() {
	apns.Initialize( AppConfig.Apns)
	storage.Initialize( AppConfig.Storage)

	// Handle any panics by calling Shutdown
    defer func() {
        if err := recover(); err != nil {
    		Shutdown()
        }
    }()

	server.Start()
}

func Shutdown() {
	fmt.Println( "shutting down")
	apns.Shutdown()
	storage.Shutdown()
}