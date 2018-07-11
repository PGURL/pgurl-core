package short

import (
	"fmt"

	"github.com/PGURL/pgurl-core/pgurl/db"
	"github.com/PGURL/pgurl-core/pgurl/encrypt"
	"github.com/PGURL/pgurl-core/pgurl/gen"
	urltool "github.com/PGURL/pgurl-core/pgurl/url"
)

func ShortURL(url string) (string, error) {
	if urltool.IsValidUrl(url) {
		hashURL, _ := encrypt.Encrypt([]byte(aesKey), url)
		// key, return to user.
		urlKey := gen.GenRandomStr(6)
		createErr := db.Create(urlKey, hashURL)
		if createErr == nil {
			return urlKey, nil
		}
	}
	return "", fmt.Errorf("url error")
}
