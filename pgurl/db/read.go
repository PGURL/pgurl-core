package db

func Read(hash string) (string, error) {
	url, err := redis_client.Get(hash).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}
