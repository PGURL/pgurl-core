## RUN test

```
go test -v -cover ./...
```

## For developer

- run redis local

```
docker run --name some-redis -d redis
```

## RUN k8s

```
kubectl create namespace pgurl
kubectl -n pgurl apply -f k8s
kubectl -n pgurl get service
```
