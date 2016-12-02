## Copyright 2016 gouvinb. All rights reserved.
## Use of this source code is governed by a BSD-style
## license that can be found in the LICENSE.md file.

# Env
UNAME_S	:=	$(shell sh -c 'uname -s 2>/dev/null || echo not')

# Makefile for Go
GO_CMD					=	go
GO_BUILD				=	$(GO_CMD) build
GO_BUILD_RACE		=	$(GO_CMD) build -race
GO_CLEAN				=	$(GO_CMD) clean
GO_DEPS					=	$(GO_CMD) get -d -v
GO_DEPS_UPDATE	=	$(GO_CMD) get -d -v -u
GO_FMT					=	$(GO_CMD) fmt
GO_GENERATE			=	$(GO_CMD) generate
GO_IMPORTS			=	goimports
GO_INSTALL			=	$(GO_CMD) install -v
GO_LINT					=	golint --min_confidence=0.3
GO_RUN					=	$(GO_CMD) run
GO_TEST					=	$(GO_CMD) test
GO_TEST_VERBOSE	=	$(GO_CMD) test -v
GO_VET					=	$(GO_CMD) vet -v

# Packages
TOP_PACKAGE_DIR	:=	github.com/gouvinb
PACKAGE					:=	go-microservice/

# Publish
ARGS	=	main.go
FILE	=	.

# notifier
ifeq ($(UNAME_S),Darwin)
	NOTIFY	=	@terminal-notifier -title Makefile \
						-subtitle "Job Finished" \
						-message "Check output" \
						-sound default \
						-appIcon "https://code.visualstudio.com/images/favicon.ico"
endif

# Primary commands
all: build

build: clean imports fmt lint vet generate
	@echo "==> Build $(PACKAGE) ..."; \
	$(GO_BUILD) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

build-race: clean imports fmt lint vet generate
	@echo "==> Build race $(PACKAGE) ..."; \
	$(GO_BUILD_RACE) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

install: clean imports fmt lint vet generate
	@echo "==> Install $(PACKAGE) ..."; \
	$(GO_INSTALL) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

publish: build fclean
	@echo "==> Publish $(PACKAGE) ..."; \
	git add $(FILE);
	@read -ep "Commit message: " MESSAGE; \
	git commit -am "$$MESSAGE";
	@read -ep "Branch name: " NAME; \
	git push origin "$$NAME";

run: clean imports fmt lint vet generate
	@echo "==> Run $(PACKAGE) ..."; \
	$(GO_RUN) $(ARGS);

test: clean imports fmt lint vet generate
	@echo "==> Unit Testing $(PACKAGE) ..."; \
	$(GO_TEST) $(TOP_PACKAGE_DIR)/$(PACKAGE) ./...;

test-verbose: clean imports fmt lint vet generate
	@echo "==> Unit Testing $(PACKAGE) ..."; \
	$(GO_TEST_VERBOSE) $(TOP_PACKAGE_DIR)/$(PACKAGE) ./...;

# Secondary commands
clean:
	@echo "==> Clean $(PACKAGE) ..."; \
	$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$(PACKAGE); \
	rm -fv ./vendor/config/bindata.go #./go-microservice.db

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

imports:
	@echo "==> Imports added for $(PACKAGE) ..."; \
	$(GO_IMPORTS) -w $(FILE)

lint:
	@echo "==> Lint $(PACKAGE) ..."; \
	$(GO_LINT) ./...;

vet:
	@echo "==> Vet $(PACKAGE) ..."; \
	$(GO_VET) ./...;

# Other commands
deps:
	@echo "==> Install dependencies for $(PACKAGE) ..."; \
	$(GO_DEPS) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

update-deps:
	@echo "==> Update dependencies for $(PACKAGE) ..."; \
	$(GO_DEPS_UPDATE) $(TOP_PACKAGE_DIR)/$(PACKAGE);

notify:
	$(NOTIFY)

# Secure command
.PHONY: all build build-race install publish run test test-verbose clean \
	 fclean fmt generate imports lint vet deps update-deps