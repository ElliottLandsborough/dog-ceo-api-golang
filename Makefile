ENVSWITCH=dev
ifeq ("$(ENVIRONMENT)","production")
	ENVSWITCH=prod
endif

build: deps compile

compile:
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
	go get -u github.com/aws/aws-lambda-go/events
	go get -u github.com/aws/aws-sdk-go/aws
	go get -u github.com/aws/aws-sdk-go/aws/awserr
	go get -u github.com/aws/aws-sdk-go/aws/session
	go get -u github.com/aws/aws-sdk-go/service/s3
	go get -u github.com/ghodss/yaml
	go get -u github.com/aws/aws-lambda-go/events
	go get -u github.com/aws/aws-lambda-go/lambda

test:
	go test -v ./breedUtil
	# go test -v ./...

clean:
	rm -rf ./bin/listAllBreeds
	rm -rf ./bin/listBreeds

start:
	sam local start-api
	#sam local start-api --env-vars environment_variables.json

sendtoaws:
	sam package --output-template-file packaged.yaml --s3-bucket dog-ceo-api-golang-$(ENVSWITCH)-sam
	sam deploy --template-file packaged.yaml --stack-name dog-ceo-api-golang-$(ENVSWITCH)-sam --capabilities CAPABILITY_IAM

deploy: test compile sendtoaws