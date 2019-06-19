.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./bin/listAllBreeds
	rm -rf ./bin/listBreeds

build:
	GOOS=linux GOARCH=amd64 go build -o bin/listAllBreeds ./listAllBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listBreeds ./listBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreeds ./listSubBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listMasterBreedImages ./listMasterBreedImages
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreedImages ./listSubBreedImages
	GOOS=linux GOARCH=amd64 go build -o bin/listMasterBreedImageRandom ./listMasterBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreedImageRandom ./listSubBreedImageRandom

start:
	sam local start-api --env-vars environment_variables.json
