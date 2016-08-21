GIT_ROOT = $(shell git rev-parse --show-toplevel)
GOPATH := $(realpath $(GIT_ROOT)/../../../..)

run-api:
	go install ./...
	$(GOPATH)/bin/runner

deploy-prod:
	eb use platform-prod
	eb deploy --verbose --timeout 20 platform-prod

deploy-staging:
	eb use platform-staging
	eb deploy -v --timeout 20 platform-staging