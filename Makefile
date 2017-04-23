all:
	@go build -v
test:
	@go test -test.v ./...
clean:
	@go clean -x

.PHONY: \
	all \
	test \
	clean
