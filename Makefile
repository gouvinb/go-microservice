## Copyright 2016 gouvinb. All rights reserved.
## Use of this source code is governed by a BSD-style
## license that can be found in the LICENSE.md file.

# Makefile for Go
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_CLEAN=$(GO_CMD) clean
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_CMD) get -d -v -u
GO_FMT=$(GO_CMD) fmt
GO_GENERATE=$(GO_CMD) generate
GO_INSTALL=$(GO_CMD) install -v
GO_LINT=golint -min_confidence=0
GO_RUN=$(GO_CMD) run
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_VET=$(GO_CMD) vet -v

# Packages
TOP_PACKAGE_DIR := github.com/gouvinb
PACKAGE := go-microservice/

# Publish
ARGS=main.go
FILE=.

# Command
all: build

build: lint vet generate
	@echo "==> Build $(PACKAGE) ..."; \
	$(GO_BUILD) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

build-race: lint vet generate
	@echo "==> Build race $(PACKAGE) ..."; \
	$(GO_BUILD_RACE) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

clean:
	@echo "==> Clean $(PACKAGE) ..."; \
	$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$(PACKAGE); \
	rm -fv ./vendor/config/bindata.go #./go-microservice.db

deps:
	@echo "==> Install dependencies for $(PACKAGE) ..."; \
	$(GO_DEPS) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

# TODO: add all file for remove and run in a new base
fclean:
	@echo "==> Clean $(PACKAGE) ..."; \
	$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$(PACKAGE); \
	rm -fv ./vendor/config/bindata.go ./go-microservice.db

fmt:
	@echo "==> Formatting $(PACKAGE) ..."; \
	$(GO_FMT) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

generate:
	@echo "==> Generate $(PACKAGE) ..."; \
	$(GO_GENERATE) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

install: fmt lint vet generate
	@echo "==> Install $(PACKAGE) ..."; \
	$(GO_INSTALL) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

lint: clean
	@echo "==> Lint $(PACKAGE) ..."; \
	$(GO_LINT) ./...;

publish: fmt lint vet generate build fclean
	@echo "==> Publish $(PACKAGE) ..."; \
	git add $(FILE);
	@read -ep "Commit message: " MESSAGE; \
	git commit -am "$$MESSAGE";
	@read -ep "Branch name: " NAME; \
	git push origin "$$NAME";

run: lint vet generate
	@echo "==> Run $(PACKAGE) ..."; \
	$(GO_RUN) $(ARGS);

test: deps lint vet generate
	@echo "==> Unit Testing $(PACKAGE) ..."; \
	$(GO_TEST) $(TOP_PACKAGE_DIR)/$(PACKAGE);

test-verbose: deps lint vet generate
	@echo "==> Unit Testing $(PACKAGE) ..."; \
	$(GO_TEST_VERBOSE) $(TOP_PACKAGE_DIR)/$(PACKAGE);

update-deps:
	@echo "==> Update dependencies for $(PACKAGE) ..."; \
	$(GO_DEPS_UPDATE) $(TOP_PACKAGE_DIR)/$(PACKAGE);

vet: clean
	@echo "==> Vet $(PACKAGE) ..."; \
	$(GO_VET) ./...;

# Secure command
.PHONY: all build build-race clean deps fclean fmt generate install lint \
	 publish run test test-verbose update-deps vet
