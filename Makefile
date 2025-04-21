.PHONY: all docs build

docs:
	swag init -g cmd/oxign/oxign.go --output api/oxign

run_api:
	go run cmd/oxign/oxign.go

run_worker:
	go run cmd/oxign/oxignd.go

build:
	go build -o out cmd/oxign/oxign.go
	go build -o out cmd/oxignd/oxignd.go
