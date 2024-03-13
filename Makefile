.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: build
build:
	go build ./cmd/monitoringlinter

.PHONY: test
test: build
	./tests/e2e.sh