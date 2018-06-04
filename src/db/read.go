package db

func Read(hash string) (string, string) {
    url, err := redis_client.Get(hash).Result()
    if err != nil {
        return "err", "err"
    }
    return url, "success"
}
