package db

func Create(hash string, url string) error {
	err := redis_client.Set(hash, url, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
