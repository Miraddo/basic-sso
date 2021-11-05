SHELL := /bin/bash

# ==============================================================================
# Running, Test and Tidy Go

run:
	go run app/services/sso-api/main.go

test:
	go test ./... -count=1
	staticcheck -checks=all ./...

tidy:
	go mod tidy
	go mod vendor

# ==============================================================================
# Building containers

VERSION := 1.0

sso-build:
	docker build \
		-f zarf/docker/dockerfile.sso-api \
		-t sso-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
sso-up:
	docker-compose \
	-f zarf/docker/docker-compose.yaml \
	up -d
	
sso-down:
	docker-compose \
	-f zarf/docker/docker-compose.yaml \
	down

sso-run:
	make sso-down
	make sso-build
	make sso-up

sso-logs:
	docker-compose \
	-f zarf/docker/docker-compose.yaml \
	logs --follow --tail 10
# ==============================================================================
# Test Service

service-test-up:
	make sso-down
	docker-compose \
	-f zarf/docker/docker-compose.test.yaml \
	up -d

service-test-down:
	docker-compose \
	-f zarf/docker/docker-compose.test.yaml \
	down

service-test-logs:
	docker-compose \
	-f zarf/docker/docker-compose.test.yaml \
	logs --follow --tail 10