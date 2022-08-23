build:
	go build -v ./cmd/apiserver/main.go

run:
	go run ./cmd/apiserver/main.go

.DEFAULT_GOAL := run
