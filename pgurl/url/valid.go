package url

import (
	"net/url"
	"regexp"
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		match, _ := regexp.MatchString("(http|https)://", toTest)
		if match {
			return true
		} else {
			return false
		}
	}
}
