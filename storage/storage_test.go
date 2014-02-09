package storage

import (
	. "launchpad.net/gocheck"
	"testing"
	)


type StorageSuite struct { 
	deviceId string
	appId string
	appVersion string
}
var _ = Suite( &StorageSuite{})


// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }


func (s *StorageSuite) SetUpSuite(c *C) {
	s.deviceId = "testDeviceId"
	s.appId = "testApp"
	s.appVersion = "0.9"

	dbPath := c.MkDir() + "/test.db"

	Initialize( Config{ dbPath})
}



func (s *StorageSuite) TestDeviceInsert( c *C ) {
	err := UpdateDevice( s.appId, s.deviceId, s.appVersion)
	c.Check( err, IsNil)
}

func (s *StorageSuite) TestDeviceUpdate( c *C ) {
	err := UpdateDevice( s.appId, s.deviceId, s.appVersion)
	c.Check( err, IsNil)
}
