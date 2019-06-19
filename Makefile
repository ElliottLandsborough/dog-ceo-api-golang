.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./bin/listAllBreeds
	rm -rf ./bin/listBreeds

build:
	GOOS=linux GOARCH=amd64 go build -o bin/listAllBreeds ./listAllBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listBreeds ./listBreeds

start:
	sam local start-api --env-vars environment_variables.json
