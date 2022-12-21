# `brew install openapi-generator` on Mac
# Installation doc for openapi-generator: https://openapi-generator.tech/docs/installation/
.PHONY: generate-openapi-client
generate-openapi-client:
	rm -rf ./pkg/client/*.go docs
	docker-compose -f ./docker-compose.yml run --rm \
		openapi-generator generate \
			-g go \
			-i internal/spec/openapi.json \
			-o internal/openapi-generator \
			-c internal/spec/config.yaml
	mv internal/openapi-generator/docs ./
	mv internal/openapi-generator/*.md ./docs/
	go fmt ./internal/openapi-generator/...
	mv ./internal/openapi-generator/*.go pkg/client/
	@$(MAKE) add-boilerplate

.PHONY: fetch-latest-spec
fetch-latest-spec:
	curl https://cockroachlabs.cloud/assets/docs/api/latest/openapi.json > internal/spec/openapi.json

# Add boilerplate header to all pkg golang files.
.PHONY: add-boilerplate
add-boilerplate:
	./internal/boilerplaterize.sh ./internal/boilerplate.txt \
		`find . -name '*.go' -path './pkg/*' -exec echo '{}' \;`
