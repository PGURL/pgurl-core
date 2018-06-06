# build stage
FROM golang:alpine
RUN apk add --update git
ADD . /go/src/pgurl-core
WORKDIR /go/src/pgurl-core

# Get govendor & fresh
RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/pilu/fresh

# govendor

## Add local package to vendor
RUN govendor add +local
RUN go get -d .

## Add external package to vendor
RUN govendor add +external
RUN govendor list

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#  
# # final stage
# # FROM scratch
# # COPY --from=build-env /src/app /
ENTRYPOINT ["fresh"]
