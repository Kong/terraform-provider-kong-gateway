.PHONY: *

all: speakeasy

format:
	@npx oas-toolkit filter-annotation --keep x-speakeasy-entity-operation,x-keep-sdk openapi.yaml > foo
	@mv foo openapi.yaml
	@npx oas-toolkit remove-unused-components openapi.yaml > foo
	@mv foo openapi.yaml
	@npx openapi-format --sortFile .openapi-format-sort.json openapi.yaml -o openapi.yaml

speakeasy: check-speakeasy
	speakeasy run --skip-versioning --output console --minimal
	@go generate .
	@git clean -fd docs/data-sources examples > /dev/null
	@git checkout -- README.md examples/README.md
	@rm USAGE.md

acceptance:
	@TF_ACC=1 go test -v ./tests -count 1

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }
