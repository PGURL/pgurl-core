# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=pgurl-core
BINARY_UNIX=$(BINARY_NAME)_unix
DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build
DOCKERRUN=$(DOCKERCMD) run
DOCKERPUSH=$(DOCKERCMD) push
DOCKER_IMAGE=pgurl/pgurl-core
DOCKER_TAG=k8s_beta
PRODUCTION_ENV_MODE=PRODUCTION
DEV_ENV_MODE=DEV

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v -cover ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync -v
	govendor add +e
	govendor list
	export ENV_MODE=$(DEV_ENV_MODE)
docker_build:
	$(DOCKERBUILD) -t $(DOCKER_IMAGE):$(DOCKER_TAG) --build-arg ENV_MODE=$(PRODUCTION_ENV_MODE) .
docker_run:
	$(DOCKERRUN) -d -e ENV_MODE=$(DEV_ENV_MODE) -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)
docker_push:
	$(DOCKERPUSH) $(DOCKER_IMAGE):$(DOCKER_TAG)