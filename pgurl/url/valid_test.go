package url

import (
	"testing"
)

func TestValidUrl(t *testing.T) {

	testurl := "http://www.google.com"
	if IsValidUrl(testurl) {
		t.Logf("%s is valid url", testurl)
	} else {
		t.Error("$s -> isValidUrl error", testurl)
	}

	testurl = "https://www.google.com"
	if IsValidUrl(testurl) {
		t.Logf("%s is valid url", testurl)
	} else {
		t.Error("$s -> isValidUrl error", testurl)
	}

	testurl = "http:/www.google.com"
	if !IsValidUrl(testurl) {
		t.Logf("%s is not valid url", testurl)
	} else {
		t.Error("$s -> isValidUrl error", testurl)
	}

	testurl = "htt//www.google.com"
	if !IsValidUrl(testurl) {
		t.Logf("%s is not valid url", testurl)
	} else {
		t.Error("$s -> isValidUrl error", testurl)
	}

	testurl = ""
	if !IsValidUrl(testurl) {
		t.Logf("%s is not valid url", testurl)
	} else {
		t.Error("$s -> isValidUrl error", testurl)
	}

}
