package db

import (
	"testing"
)

func TestRedisPing(t *testing.T) {
	ping := Ping()
	if ping != "PONG" {
		t.Error("redis connect error")
	}
	t.Logf("%s", ping)
}

func TestRedisCR(t *testing.T) {
	// write some test
	hash := "i3s9f81"
	url := "https://www.google.com"
	createErr := Create(hash, url)
	if createErr != nil {
		t.Error("redis created is error")
	} else {
		t.Logf("Create is success")
	}
	rurl, readErr := Read(hash)

	if readErr == nil && rurl == url {
		t.Logf("Read is success")
	}
}
