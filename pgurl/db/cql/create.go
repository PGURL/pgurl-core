package cql

func Create(hash string, url string) error {
	err := cqlClient.Query(`INSERT INTO shorturl (hashurl, url, createdtime) VALUES (?, ?, dateof(now()))`,
		hash, url).Exec()
	if err != nil {
		return err
	}
	return nil
}
