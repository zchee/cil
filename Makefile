SHELL = /usr/bin/env bash

APP = cil
REPOSITORY = github.com/zchee/${APP}

GCFLAGS ?=
LDFLAGS ?=

PACKAGES = $(shell go list ./... | grep -v -e 'swagger')
FMT_PACKAGES = $(shell go list -f '{{.Dir}}' ./...)
GO_TEST_FLAGS ?=
GO_BENCH_FUNCS ?= .
GO_BENCH_FLAGS ?= $(GO_TEST_FLAGS) -run=^$$ -bench=${GO_BENCH_FUNCS}

define target
  @printf "+ \\033[32m$@\\033[0m\\n"
endef

all: $(APP)

.PHONY: $(APP)
$(APP):
	go build -v -o ./$(APP) $(GCFLAGS) $(LDFLAGS) $(REPOSITORY)/cmd/$(APP)

.PHONY: swagger
swagger:
	swagger-codegen generate -i provider/circleci/schema/swagger.yaml -l go -D 'packageName=circleci' -o provider/circleci/swagger
	command cp -f ./provider/circleci/types/* ./provider/circleci/swagger

.PHONY: clean
clean:
	${RM} -r $(APP) ./tools *.test *.out

.PHONY: deps
deps:
	@go install -v -x $(shell go list -deps ./...)

.PHONY: dep
dep: $(shell command -v dep)
	dep ensure -v

.PHONY: dep.update
dep.update: $(shell command -v dep)
	dep ensure -v -update

.PHONY: test
test:
	go test -v $(GO_TEST_FLAGS) $(TEST_SRC_PACKAGES)

.PHONY: benchmark
benchmark: benchmark
	go test -v $(GO_BENCH_FLAGS) $(TEST_SRC_PACKAGES)

.PHONY: coverage
coverage:
	go test -v -race -covermode=atomic -coverpkg=$(REPOSITORY)/... -coverprofile=coverage.out ./...

.PHONY: lint
lint: fmt vet goimports golint megacheck errcheck

.PHONY: lint.install
lint.install: tools/goimports tools/golint tools/megacheck tools/errcheck

.PHONY: fmt
fmt:
	$(call target)
	@gofmt -s -l $(FMT_PACKAGES) | grep -v -e 'swagger' | grep -E '.'; test $$? -eq 1

.PHONY: vet
vet:
	$(call target)
	@go vet -all $(PACKAGES)

tools:
	mkdir ./tools

tools/goimports: tools
	go build -o $@ ./vendor/golang.org/x/tools/cmd/goimports

.PHONY: goimports
goimports: tools/goimports
	$(call target)
	@./tools/goimports -l $(FMT_PACKAGES) | grep -v -e 'swagger' | grep -E '.'; test $$? -eq 1

tools/golint: tools
	go build -o $@ ./vendor/golang.org/x/lint/golint

.PHONY: golint
golint: tools/golint
	$(call target)
	@./tools/golint -min_confidence=0.3 -set_exit_status $(PACKAGES)

tools/megacheck: tools
	go build -o $@ ./vendor/honnef.co/go/tools/cmd/megacheck

.PHONY: megacheck
megacheck: tools/megacheck
	$(call target)
	@./tools/megacheck $(PACKAGES)

tools/errcheck: tools
	go build -o $@ ./vendor/github.com/kisielk/errcheck

.PHONY: errcheck
errcheck: tools/errcheck
	$(call target)
	@./tools/errcheck -exclude .errcheckignore $(PACKAGES)
