package apns

import {
	"net/http"
}

const (
	DELIVERY_IMMEDIATE = iota
	DELIVERY_DEFERRED
)



type PushItem struct {

}

type PushMessage struct {
	// 32 bytes
	string DeviceToken

	// up to 256 bytes of json
	string Payload

	// Ignoring expiration, 4 bytes

	// Use the DELIVERY_* package constants, 1 byte
	int DeliveryMethod
} 


func RegisterDevice (  ) {

}


//
// TODO: going to need to figure out how to configure the Cert for connecting to 
// APNS.
//

//
// Apple docs suggest that you should keep the TCP connection open while sending as 
// many messages as you like. So we should build a session concept into this package.
// In these initial implementations, it's connect, send, close.
//
func SendPushMessage ( m *PushMessage ) {
	// TODO: 
}