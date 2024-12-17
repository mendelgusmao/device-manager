GO=$(shell which go)
DOCKER=$(shell which docker)

TAG=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
BRANCH=$(if $(TAG),$(TAG),$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))
HASH=$(shell git rev-parse --short=7 HEAD 2>/dev/null)
TIMESTAMP=$(shell git show -s --date=format:'%Y%m%d%H%M' --format=%cd $(HASH))
GIT_REV=$(shell printf "%s-%s-%s" "$(BRANCH)" "$(HASH)" "$(TIMESTAMP)")
REV=$(if $(filter --,$(GIT_REV)),latest,$(GIT_REV))

run:
	$(GO) run cmd/device-manager-api/*.go

setup:
	$(GO) mod tidy

build:
	$(GO) build -ldflags "-X internal.application.version.revision=$(REV) -s -w" -o device-manager-api cmd/device-manager-api/*.go

test:
	$(GO) test -race ./...

docker/build-image:
	$(DOCKER) build -t mendelgusmao/device-manager-api:development .

docker/run: docker/build-image
	$(DOCKER) compose up
