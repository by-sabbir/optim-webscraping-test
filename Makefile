.PHONY:

SERVICE_NAME := scraper
VERSION := $(shell git describe --tags --abbrev=0)

build:
	docker build \
		-t $(SERVICE_NAME):$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

binary:
	go build -ldflags "-X main.build=${BUILD_REF}"

test:
	go test -cover ./...