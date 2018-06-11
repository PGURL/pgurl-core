package url

import (
	"net/url"
	"regexp"
)

func IsValidUrl(toTestUrl string) bool {
        res := true
        //  Use parse requests uri to check toTest string.
	_, err := url.ParseRequestURI(toTestUrl)
        if (err != nil){
	    res = false
        }

        //  Check http or https include in toTest string.
	match, _ := regexp.MatchString("(http|https)://", toTestUrl)
	if !match {
            res = false
	}

        return res
}
