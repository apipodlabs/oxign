.PHONY: all docs build api worker

api:
	go run cmd/oxign/oxign.go

worker:
	go run cmd/oxign/oxignd.go

build:
	go build -o out cmd/oxign/oxign.go
	go build -o out cmd/oxignd/oxignd.go
