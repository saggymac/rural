package main

import (
	"github.com/saggymac/rural/server"
	"github.com/saggymac/rural/apns"
	)

func main() {
	apns.Initialize( AppConfig.apns)

	server.Start()
}
