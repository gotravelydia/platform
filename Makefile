GIT_ROOT = $(shell git rev-parse --show-toplevel)
GOPATH := $(realpath $(GIT_ROOT)/../../../..)

run-api:
	go install ./...
	BIT_ROLE=api $(GOPATH)/bin/runner

deploy-platform-prod:
	eb use platform-prod
	eb setenv BIT_ROLE=api BIT_ENV=prod && eb deploy --verbose --timeout 20 platform-prod

deploy-platform-staging:
	eb use platform-staging
	eb setenv BIT_ROLE=api BIT_ENV=staging && eb deploy -v --timeout 20 platform-staging
