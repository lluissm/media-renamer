.PHONY: all clean build test test-e2e license

all: build test

BIN_PATH = bin
VERSION = $(shell git describe --tags)
CMD = media-renamer
BUILD_PATH = ./cmd/$(CMD)
LD_FLAGS = -ldflags="-X 'main.version=${VERSION}'"
TOOLS := $(CURDIR)/.tools

# remove bin folder
clean:
	@echo ">> Cleaning project"
	rm -rf ${BIN_PATH}

# Build locally in bin folder
build:
	@echo ">> Building project"
	go mod tidy
	go build -v ${LD_FLAGS} -o ${BIN_PATH}/${CMD} ${BUILD_PATH}

# Run unit tests
test:
	@echo ">> Running unit tests"
	go test -cover ./...

# Run end to end tests
test-e2e: build
	@echo ">> Running end to end tests"
	cd test && bash test.sh ../${BIN_PATH}/${CMD}

# Add missing licenses
license:
	@echo ">> Adding missing licenses"
	if [[ ! -f "./bin/license-header-checker" ]]; then curl -s https://raw.githubusercontent.com/lluissm/license-header-checker/master/install.sh | bash; fi
	./bin/license-header-checker -a -r -i testdata ./license_header.txt . go && [[ -z `git status -s` ]]

# Execute golangci-lint
lint:
	@echo ">> Executing golangci-lint"
	if [[ ! -f "./bin/golangci-lint" ]]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.49.0; fi
	./bin/golangci-lint run