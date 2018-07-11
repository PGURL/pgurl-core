package short

import (
	"fmt"

	"github.com/PGURL/pgurl-core/pgurl/db/cql"
	"github.com/PGURL/pgurl-core/pgurl/encrypt"
)

func ReShort(urlKey string) (string, error) {
	hashURL, readErr := cql.Read(urlKey)
	if readErr != nil {
		return "", fmt.Errorf("read url error")
	}
	url, _ := encrypt.Decrypt([]byte(aesKey), hashURL)
	return url, nil
}
