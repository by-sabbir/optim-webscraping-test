.PHONY:

SERVICE_NAME := scraper
VERSION := v0.0.1

build:
	docker build \
		-t $(SERVICE_NAME):$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

binary:
	go build -ldflags "-X main.build=${BUILD_REF}"