package apns


//
// what I want to achieve here is a dynamic connection pool.
// Meaning, we would maintain up to a max number of connections
// to APNS (defined by config) that will decay with reduced load; 
// grow up to the max as load demands.
//


const (
	DELIVERY_IMMEDIATE = iota
	DELIVERY_DEFERRED
)


type Config struct {
	ProdServer string "prodServer"
	SandboxServer string "sandboxServer"
	ProdFeedbackServer string "prodFeedbackServer"
	SandboxFeedbackServer string "sandboxFeedbackServer" 	
}

var config = Config{}

type PushItem struct {

}

type PushMessage struct {
	// 32 bytes
	DeviceToken string

	// up to 256 bytes of json
	Payload string

	// Ignoring expiration, 4 bytes

	// Use the DELIVERY_* package constants, 1 byte
	DeliveryMethod int
} 


func Shutdown() {

}

func Initialize( cnf Config ) {
	config = cnf
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