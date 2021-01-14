.PHONY: run
run:
	go run cmd/goshogi/main.go

.PHONY: test
test:
	go test ./...

.PHONY: testv
testv:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: mod
mod:
	go mod download