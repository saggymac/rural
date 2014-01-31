
package server

import (
	"net/http"
	"fmt"
	"time"
	"log"
	)

var (
	serverInstance *http.Server = nil
)


func Setup() {
	serverInstance := &http.Server{
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc( "/device", deviceHandler)

	log.Fatal( serverInstance.ListenAndServe())
}


func deviceHandler ( writer http.ResponseWriter, request *http.Request ) {
	// todo
	fmt.Fprintf( writer, "hello")
}