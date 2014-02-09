
package server

import (
	"net/http"
	"time"
	"log"
	"github.com/saggymac/rural/storage"		
	)

var (
	serverInstance *http.Server = nil
)


func Start() {
	serverInstance := &http.Server{
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc( "/device", deviceHandler)
	http.HandleFunc( "/push", pushHandler)
	http.HandleFunc( "/pushraw", pushRawHandler)	

	log.Fatal( serverInstance.ListenAndServe())
}



// This endpoint registers (or updates) a device in our device tables
// When a device registers, it will send its appId (e.g., iOS bundle ID) and it's device token.
// We then have to send that into the database, or update it's timestamp it if already exists.
// The appId,deviceId tuple is the uniqueness constraint on the databases table.
//
// For this handler, we are expecting a POST message with HTML FORM encoded values for:
//   appId
//   deviceId
//   appVersion
//
func deviceHandler ( writer http.ResponseWriter, request *http.Request ) {

	appId := request.FormValue( "appId")
	deviceId := request.FormValue( "deviceId")
	appVersion := request.FormValue( "appVersion")

	badRequest :=  len( appId) <= 0
	badRequest = badRequest || (len( deviceId) <= 0)
	badRequest = badRequest || (len( appVersion) <= 0)

	if badRequest {
		// todo: send an error response
	}

	storage.UpdateDevice( appId, deviceId, appVersion)

	// todo: send success response
}



// Normal, HTML FORM encoded endpoint for push messaging
//
func pushHandler ( writer http.ResponseWriter, request *http.Request ) {

	// body, err := ioutil.ReadAll( request.Body)
	// if err != nil {
	// 	// todo: handle error
	// }


}


// In this case, the user must have sent a valid JSON payload as
// defined the APNS specs.
//
func pushRawHandler ( writer http.ResponseWriter, request *http.Request ) {

	// body, err := ioutil.ReadAll( request.Body)
	// if err != nil {
	// 	// todo: handle error
	// }


}


