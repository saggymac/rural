package main

import (
	"github.com/saggymac/rural/server"
	"github.com/saggymac/rural/apns"
	"github.com/saggymac/rural/storage"	
	)

func main() {
	apns.Initialize( AppConfig.Apns)
	storage.Initialize( AppConfig.Storage)

	server.Start()
}
