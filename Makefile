#==============================================================================
#
# Makefile for building Docker images and pushing them to AWS ECR registries.
#
#==============================================================================

# Set app identity.
APP=devbox
PKG=github.com/mojochao/$(APP)
VERSION := $(shell cat VERSION | tr -d '\n')

# Discover OS for default static build configuration.
OS = shell("uname")
ifeq ($(OS),Linux)
    OS=linux
else
    OS=darwin
endif

# Set GOOS for static builds using discovered OS if not provided.
GOOS ?= $(OS)

# Set Docker image build and run configuration.
DOCKER_FILE ?= Dockerfile
IMAGE = mojochao/$(APP)

#==============================================================================
#
# Define help targets with descriptions provided in trailing `##` comments.
#
# Note that the '## description' is used in generating documentation when 'make'
# is invoked with no arguments.
#
# See https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html for
# additional details on how this works.
#
#==============================================================================

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

#==============================================================================
#
# Define Golang targets.
#
#==============================================================================

.PHONY: prep
prep: ## Prepare Golang tools needed for builds
	@echo 'Installing govvv'
	go get github.com/ahmetb/govvv
	go get golang.org/x/lint/golint

.PHONY: build
build: ## Build the application
	@echo 'building $(APP)'
	CGO_ENABLED=0 GOOS=$(GOOS) govvv build -a -installsuffix cgo -ldflags '-extldflags "-static"' -pkg $(PKG)/internal/build -o $(APP) .

.PHONY: lint
lint: ## Lint the application
	@echo 'linting $(APP)'
	golint ./...

.PHONY: test
test: ## Run all tests
	@echo 'testing $(APP)'
	go test -v ./...

.PHONY: clean
clean: ## Clean build artifacts
	@echo 'cleaning $(APP)'
	rm -f $(APP)

#==============================================================================
#
# Define docker targets.
#
#==============================================================================

.PHONY: docker-build
docker-build: ## Build docker image
	@echo 'building docker image $(IMAGE):latest'
	DOCKER_BUILDKIT=1 docker build -f $(DOCKER_FILE) -t $(IMAGE):latest .
	docker tag $(IMAGE):latest $(IMAGE):$(VERSION)

.PHONY: docker-push
docker-push: ## Push docker images to Docker Hub
	@echo 'pushing docker image $(IMAGE):$(TAG) to Docker Hub'
	docker push $(IMAGE):latest
	docker push $(IMAGE):$(VERSION)
