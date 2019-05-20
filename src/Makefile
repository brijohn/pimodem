PACKAGE     = github.com/brijohn/pimodem

DATE       ?= $(strip $(shell date -u +"%FT%TZ"))
GIT_COMMIT  = $(strip $(shell git rev-parse --short HEAD))
VERSION    ?= $(shell git describe --always --exact-match --match=v* 2> /dev/null || \
			head -n 1 $(CURDIR)/VERSION 2> /dev/null || echo devel)
ifeq ($(strip $(shell git status --porcelain 2>/dev/null)),)
	GIT_STATE=clean
else
	GIT_STATE=dirty
endif

GOPATH   = $(CURDIR)/.gopath
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))

GO       = go
GOFMT    = gofmt

GOARCH     ?= $(shell $(GO) env GOARCH)
GOOS       ?= $(shell $(GO) env GOOS)
CGOCC      ?= $(shell $(GO) env CC)
GOHOSTARCH  = $(shell $(GO) env GOHOSTARCH)
GOHOSTOS    = $(shell $(GO) env GOHOSTOS)

V ?= 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: fmt vendor main-packr.go | $(BASE) ; $(info $(M) building executable…) @ ## Build program binary
	$Q cd $(BASE) && CC=$(CGOCC) GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=1 $(GO) build \
		-gcflags '-e' \
		-ldflags '-X main.version=$(VERSION) -X main.buildDate=$(DATE) -X main.gitCommit=${GIT_COMMIT} -X main.gitState=${GIT_STATE}' \
		-o bin/$(PACKAGE) .

$(BASE): ; $(info $(M) setting GOPATH…)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@


GOLINT = $(BIN)/golint
$(BIN)/golint: | $(BASE) ; $(info $(M) building golint…)
	$Q GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GO) get github.com/golang/lint/golint

GODEP = $(BIN)/dep
$(BIN)/dep: | $(BASE) ; $(info $(M) building godep…)
	$Q GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GO) get github.com/golang/dep/cmd/dep

GOPACKR = $(BIN)/packr
$(BIN)/packr: | $(BASE) ; $(info $(M) building packr…)
	$Q GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GO) get github.com/gobuffalo/packr/packr

.PHONY: lint
lint: vendor | $(BASE) $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

Gopkg.toml: | $(BASE) $(GODEP) ; $(info $(M) init dependencies…)
	$Q cd $(BASE) && $(GODEP) init

Gopkg.lock: Gopkg.toml | $(BASE) $(GODEP) ; $(info $(M) updating dependencies…)
	@touch $@
	$Q cd $(BASE) && $(GODEP) ensure -update -no-vendor

main-packr.go: | $(BASE) $(GOPACKR) ; $(info $(M) packing assets…)
	$Q cd $(BASE) && $(GOPACKR) -i audio.go

vendor: Gopkg.lock | $(BASE) $(GODEP) ; $(info $(M) retrieving dependencies…) @ ## Update vendor directory
	$Q cd $(BASE) && $(GODEP) ensure
	@touch $@

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(GOPATH)
	@rm -rf vendor
	@rm -rf bin
	@rm -rf *-packr.go


.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
