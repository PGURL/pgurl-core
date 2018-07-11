package short

import "github.com/PGURL/pgurl-core/pgurl/config"

var aesKey = getAesKey()

func getAesKey() string {
	config := config.GetConfig()
	key := config.GetString("key")
	return key
}
