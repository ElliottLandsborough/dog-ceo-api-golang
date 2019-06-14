.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./bin/getAllBreeds
	
build:
	GOOS=linux GOARCH=amd64 go build -o bin/getAllBreeds ./getAllBreeds