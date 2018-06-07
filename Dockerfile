# build stage
FROM golang:alpine
RUN apk add --update git
ADD . /go/src/github/PGURL/pgurl-core
WORKDIR /go/src/github/PGURL/pgurl-core

# Get govendor & fresh
RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/pilu/fresh

# govendor

## Download external package to vendor
RUN govendor sync -v
## Add local package to vendor
RUN govendor add +e
RUN govendor list

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#  
# # final stage
# # FROM scratch
# # COPY --from=build-env /src/app /
ENTRYPOINT ["fresh"]
