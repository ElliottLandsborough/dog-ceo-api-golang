ENVSWITCH=dev
ifeq ("$(ENVIRONMENT)","production")
	ENVSWITCH=prod
endif

build:
	GOOS=linux GOARCH=amd64 go build -o bin/listAllBreeds ./listAllBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listBreeds ./listBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreeds ./listSubBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/listMasterBreedImages ./listMasterBreedImages
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreedImages ./listSubBreedImages
	GOOS=linux GOARCH=amd64 go build -o bin/listMasterBreedImageRandom ./listMasterBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreedImageRandom ./listSubBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/listAnyBreedImageRandom ./listAnyBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/listAnyBreedMultiImageRandom ./listAnyBreedMultiImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/listMasterBreedInfo ./listMasterBreedInfo
	GOOS=linux GOARCH=amd64 go build -o bin/listSubBreedInfo ./listSubBreedInfo

deps:
	go get -u ./...

clean:
	rm -rf ./bin/listAllBreeds
	rm -rf ./bin/listBreeds

start:
	sam local start-api --env-vars environment_variables.json

deploy:
	sam package --output-template-file packaged.yaml --s3-bucket dog-ceo-api-golang-$(ENVSWITCH)-sam
	sam deploy --template-file packaged.yaml --stack-name dog-ceo-api-golang-$(ENVSWITCH)-sam --capabilities CAPABILITY_IAM