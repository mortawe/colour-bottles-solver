MODULE   = $(shell $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat .version 2> /dev/null || echo v0)
PKGS     = $(or $(PKG),$(shell $(GO) list ./...))
BIN      = bin

GO      = go
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell if [ "$$(tput colors 2> /dev/null || echo 0)" -ge 8 ]; then printf "\033[34;1m▶\033[0m"; else printf "▶"; fi)

GENERATED = # List of generated files

SOURCE = $(shell $(GO) list -f '{{$$d := .Dir}}{{range $$f := .GoFiles}}{{printf "%s/%s\n" $$d $$f}}{{end}}{{range $$f := .CgoFiles}}{{printf "%s/%s\n" $$d $$f}}{{end}}{{range $$f := .TestGoFiles}}{{printf "%s/%s\n" $$d $$f}}{{end}}' $(PKGS))


.SUFFIXES:
.PHONY: all
all: $(BIN) ; $(info $(M) building executable…) @ ## Build program binary
	$Q $(GO) build \
		-tags release \
		-ldflags '-X main.Version=$(VERSION) -X main.BuildDate=$(DATE)' \
		-o $(BIN)/solver .

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)…)
	$Q env GOBIN=$(abspath $(BIN)) $(GO) install $(PACKAGE)

GOIMPORTS = $(BIN)/goimports
$(BIN)/goimports: PACKAGE=golang.org/x/tools/cmd/goimports@latest

REVIVE = $(BIN)/revive
$(BIN)/revive: PACKAGE=github.com/mgechev/revive@v1.2.4

GOLINES = $(BIN)/golines
$(BIN)/golines: PACKAGE=github.com/segmentio/golines@v0.11.0

GOCOV = $(BIN)/gocov
$(BIN)/gocov: PACKAGE=github.com/axw/gocov/gocov@latest

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: PACKAGE=github.com/AlekSi/gocov-xml@latest

GOTESTSUM = $(BIN)/gotestsum
$(BIN)/gotestsum: PACKAGE=gotest.tools/gotestsum@latest

COVERAGE_MODE = atomic
.PHONY: test-coverage
test-coverage: | $(GOCOV) $(GOCOVXML) $(GOTESTSUM) ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p test
	$Q $(GOTESTSUM) -- \
		-coverpkg=$(shell echo $(PKGS) | tr ' ' ',') \
		-covermode=$(COVERAGE_MODE) \
		-coverprofile=test/profile.out $(PKGS)
	$Q $(GO) tool cover -html=test/profile.out -o test/coverage.html
	$Q $(GOCOV) convert test/profile.out | $(GOCOVXML) > test/coverage.xml
	@echo -n "Code coverage: "; \
		echo "scale=1;$$(sed -En 's/^<coverage line-rate="([0-9.]+)".*/\1/p' test/coverage.xml) * 100 / 1" | bc -q

.PHONY: lint
lint: | $(REVIVE) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(REVIVE) -formatter friendly -set_exit_status ./...

.PHONY: fmt
fmt: | $(GOIMPORTS) $(GOLINES)  ## Run formatters
	$(info $(M) running gofmt…) 
	$Q $(GOIMPORTS) -local $(MODULE) -w $(SOURCE) 
	$(info $(M) running golines…) 
	$Q $(GOLINES) -w -m 110 --shorten-comments $(SOURCE)
	
.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
