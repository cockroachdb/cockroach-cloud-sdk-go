# TOOLPATH is the location of tools used during builds.
TOOLPATH := $(abspath bin)

# `brew install openapi-generator` on Mac
# Installation doc for openapi-generator: https://openapi-generator.tech/docs/installation/
.PHONY: generate-openapi-client
generate-openapi-client: bin/goimports
	rm -rf ./pkg/client/*.go docs
	docker-compose -f ./docker-compose.yml run --rm \
		jq -M -f filter.jq internal/spec/openapi.json > internal/openapi-generator/api/openapi-modified.json
	docker-compose -f ./docker-compose.yml run --rm \
		openapi-generator generate \
			-g go \
			-i internal/openapi-generator/api/openapi-modified.json \
			-o internal/openapi-generator \
			-c internal/spec/config.yaml
	mv internal/openapi-generator/docs ./
	mv internal/openapi-generator/*.md ./docs/
	cp ./docs/README.md ./
	bin/goimports -w ./internal/openapi-generator/
	go fmt ./internal/openapi-generator/...
	mv ./internal/openapi-generator/*.go pkg/client/
	@$(MAKE) add-boilerplate

bin/goimports:
	@TOOLPKG=github.com/cockroachdb/gostdlib/x/tools/cmd/goimports@v1.19.0 $(MAKE) build-tool

.PHONY: fetch-latest-spec
fetch-latest-spec:
	curl -L https://cockroachlabs.cloud/assets/docs/api/latest/openapi.json > internal/spec/openapi.json

# Add boilerplate header to all pkg golang files.
.PHONY: add-boilerplate
add-boilerplate:
	./internal/boilerplaterize.sh ./internal/boilerplate.txt \
		`find . -name '*.go' -path './pkg/*' -exec echo '{}' \;`

# Validate that the generated package is valid and free of syntax errors.
.PHONY: validate
validate:
	go run main.go

default: generate-openapi-client validate

# build-tool is a helper that builds $TOOL_PKG in a temp directory so that it
# can use its own isolated go.mod file.
.PHONY: build-tool
build-tool:
	@{ \
	TMP_DIR=$$(mktemp -d); \
	cd $$TMP_DIR; \
	go mod init tmp; \
	go get $(TOOLPKG); \
	GOBIN=$(TOOLPATH) go install $(TOOLPKG); \
	rm -rf $$TMP_DIR; \
	}
