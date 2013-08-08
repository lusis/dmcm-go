all: dmcm cli

clean-all: clean all

dmcm:
	@go install dmcm/geography
	@go install dmcm/infrastructure
	@go install dmcm/utils

cli:
	@mkdir -p bin/
	@go get ./...
	@go install dmcm-cli

clean:
	@rm -rf bin/ pkg/ src/github.com/

.PHONY: all clean clean-all dmcm cli
