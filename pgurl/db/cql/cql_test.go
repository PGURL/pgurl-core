package cql

import "testing"

func TestCqlCR(t *testing.T) {
	// write some test
	hash := "i3s9f81"
	url := "https://www.google.com"
	createErr := Create(hash, url)
	if createErr != nil {
		t.Error("cql created is error")
	} else {
		t.Logf("cql Create is success")
	}
	rurl, readErr := Read(hash)

	if readErr == nil && rurl == url {
		t.Logf("cql Read is success")
	} else {
		t.Errorf(readErr.Error())
	}
}
