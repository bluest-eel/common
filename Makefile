VERSION = $(shell cat VERSION)
GO ?= go
GOFMT ?= $(GO)fmt
GOMOD=on
DOCKER_ORG = bluesteelabm

FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
DEFAULT_GOPATH = $(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN = $(DEFAULT_GOPATH)/bin
export PATH := $(PATH):$(DEFAULT_GOBIN)

GOLANGCI_LINT = $(DEFAULT_GOBIN)/golangci-lint
RICH_GO = $(DEFAULT_GOBIN)/richgo

default: all

all-common: lint
all: all-common test
all-cicd: all-common test-nocolor

#############################################################################
###   Source Code   #########################################################
#############################################################################
###
### Linting, building, testing, etc.
###

show-linter:
	@echo $(GOLANGCI_LINT)

show-version:
	@echo $(VERSION)

deps:
	@echo '>> Downloading deps ...'
	@GO111MODULE=$(GOMOD) $(GO) get -v -d ./...

$(GOLANGCI_LINT):
	@echo ">> Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.15.0

lint: $(GOLANGCI_LINT)
	@echo '>> Linting source code'
	@GL_DEBUG=linters_output GOPACKAGESPRINTGOLISTERRORS=1 $(GOLANGCI_LINT) \
	--enable=golint \
	--enable=gocritic \
	--enable=misspell \
	--enable=nakedret \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	run ./...

$(RICH_GO):
	@echo ">> Couldn't find $(RICH_GO); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u github.com/kyoh86/richgo

test: $(RICH_GO)
	@echo '>> Running all tests'
	@$(RICH_GO) test ./... -v

test-nocolor:
	@echo '>> Running all tests'
	@$(GO) test ./... -v

#############################################################################
###   Release Process   #####################################################
#############################################################################

tag:
	@echo "Tags:"
	@git tag
	@git tag "v$(VERSION)"
	@echo "New tag list:"
	@git tag

tag-and-push: tag
	@git push --tags

#############################################################################
###   Misc   ################################################################
#############################################################################

clean-cache:
	@echo '>> Purging Go mod cahce ...'
	@$(GO) clean -cache
	@$(GO) clean -modcache

show-targets:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | \
	awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
	sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'