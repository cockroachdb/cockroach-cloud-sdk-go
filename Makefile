# `brew install openapi-generator` on Mac
# Installation doc for openapi-generator: https://openapi-generator.tech/docs/installation/
.PHONY: generate-openapi-client
generate-openapi-client:
	rm -rf ./pkg/client/*.go docs
	openapi-generator generate -g go -i internal/spec/swagger.json -o internal/ccloud \
		-c internal/spec/config.yaml -t internal/spec
	mv internal/ccloud/docs ./
	mv internal/ccloud/*.md ./docs/
	mv ./internal/ccloud/api_cockroach_cloud.go ./internal/ccloud/service.go
	go fmt ./internal/ccloud/...
	mv ./internal/ccloud/*.go pkg/client/

# Add boilerplate header to all pkg golang files.
.PHONY: add-boilerplate
add-boilerplate:
	./internal/boilerplaterize.sh ./internal/boilerplate.txt \
		`find . -name '*.go' -path './pkg/*' -exec echo '{}' \;`
