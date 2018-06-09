# build stage
FROM golang:alpine AS build-env
RUN apk add --update git
ADD . /go/src/github.com/PGURL/pgurl-core
WORKDIR /go/src/github.com/PGURL/pgurl-core

# Get govendor
RUN go get -u github.com/kardianos/govendor

# govendor

## Download external package to vendor
RUN govendor sync -v
## Add local package to vendor
RUN govendor add +e
RUN govendor list

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final stage
FROM scratch
COPY --from=build-env /go/src/github.com/PGURL/pgurl-core/app /
COPY --from=build-env /go/src/github.com/PGURL/pgurl-core/confs/ /confs

# Start Set ENV
ARG ENV_MODE=development
ENV ENV_MODE=$ENV_MODE
# END Set ENV
ENTRYPOINT ["/app"]
