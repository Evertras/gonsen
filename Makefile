.PHONY: default
default:
	go run ./examples/basic/*.go

.PHONY: fmt
fmt:
	go fmt ./...

