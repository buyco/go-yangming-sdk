DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v5.4.0 generate -i sdk/api_files/yangming_track_and_trace.json
GO_CLIENT := -g go -o /sdk/yangming \
			--git-repo-id=go-yangming-sdk --git-user-id=buyco \
			--additional-properties=packageName=yangming \
			--additional-properties=isGoSubmodule=true \
			--additional-properties=generateInterfaces=true

## generate: Clean and generate SDK from file.
generate:
	${MAKE} clean
	${MAKE} go-sdk

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./one

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo