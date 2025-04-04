.PHONY: *

all: speakeasy

format:
	@npx oas-toolkit filter-annotation --keep x-speakeasy-entity-operation openapi.yaml > foo
	@mv foo openapi.yaml
	@npx oas-toolkit remove-unused-components openapi.yaml > foo
	@mv foo openapi.yaml
	@npx openapi-format --sortFile .openapi-format-sort.json openapi.yaml -o openapi.yaml

speakeasy: check-speakeasy
	speakeasy run -o console --skip-versioning
	@go generate .
	@git checkout -- README.md examples/README.md
	@rm USAGE.md

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }
