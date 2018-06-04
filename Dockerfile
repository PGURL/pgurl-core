# build stage
FROM golang:alpine AS build-env
RUN apk add --update git
ADD . /src
WORKDIR /src
RUN go get -d .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
 
# final stage
FROM centurylink/ca-certs
COPY --from=build-env /src/app /
ENTRYPOINT ["/app"]
