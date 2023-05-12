.PHONY: gen

gen:
	cd internal/transport/rest && mkdir -p gen; \
	oapi-codegen --config ./server.config.yaml ./openapi.yaml;