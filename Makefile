.PHONY: clean build

clean: 
	rm -rf ./bin/getAllBreeds
	
build:
	GOOS=linux GOARCH=amd64 go build -o bin/getAllBreeds ./getAllBreeds