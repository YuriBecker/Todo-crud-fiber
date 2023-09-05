.PHONY: run build clean

default: run

build:
	@go build cmd/api/main.go

run:
	@go run cmd/api/main.go

clean:
	@rm ./main