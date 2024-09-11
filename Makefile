.PHONY: *

all: speakeasy

speakeasy: check-speakeasy
	speakeasy run -o console --skip-versioning
	@git clean -fd examples docs > /dev/null
	@git checkout -- README.md examples/README.md
	@rm USAGE.md

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }