.PHONY: all clean build test license

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

# Add missing licenses
license:
	@echo ">> Adding missing licenses"
	curl -s https://raw.githubusercontent.com/lluissm/license-header-checker/master/install.sh | bash && ./bin/license-header-checker -a -r -i testdata ./license_header.txt . go && [[ -z `git status -s` ]]
