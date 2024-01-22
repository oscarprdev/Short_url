.PHONY: Short_url api_spec

serve:
	@echo Serving Api spec
	npx @redocly/cli build-docs index.yaml && open redoc-static.html

compile:
	multi-file-swagger -o yaml index.yaml > compiled_tmp1.yaml && \
	multi-file-swagger -o yaml compiled_tmp1.yaml > compiled_tmp2.yaml && \
	multi-file-swagger -o yaml compiled_tmp2.yaml > compiled_tmp3.yaml && \
	multi-file-swagger -o yaml compiled_tmp3.yaml > compiled.yaml && \
	rm compiled_tmp1.yaml compiled_tmp2.yaml compiled_tmp3.yaml && \
	make format

format:
	sed -i -E 's/\$ref: components\/\([^ ]*\)\.yaml/\$ref: \#\/components\/\1/; s/_index#//; s#\/\/#\/#; s/\$ref: \(.*\)/$ref: "\1"/' compiled.yaml

generate:
	make compile && \
	oapi-codegen \
		-generate 'types, server, strict-server' \
		-package short_url compiled.yaml > ../backend/pkg/infra/generated/api.go