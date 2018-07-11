package cql

import "fmt"

func Read(hash string) (string, error) {
	iter := cqlClient.Query(`select url from shorturl where hashurl=?`, hash).Iter()
	res, err := iter.SliceMap()
	if err == nil && len(res) > 0 {
		url := res[0]["url"]
		return url.(string), nil
	}

	return "", fmt.Errorf("read err")
}
