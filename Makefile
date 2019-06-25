ENVSWITCH=dev
ifeq ("$(ENVIRONMENT)","production")
	ENVSWITCH=prod
endif

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/allBreeds ./app/controllers/allBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/masterBreeds ./app/controllers/masterBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/subBreeds ./app/controllers/subBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/masterBreedImages ./app/controllers/masterBreedImages
	GOOS=linux GOARCH=amd64 go build -o bin/subBreedImages ./app/controllers/subBreedImages
	GOOS=linux GOARCH=amd64 go build -o bin/masterBreedImageRandom ./app/controllers/masterBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/subBreedImageRandom ./app/controllers/subBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/anyBreedImageRandom ./app/controllers/anyBreedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/anyBreedMultiImageRandom ./app/controllers/anyBreedMultiImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/masterBreedInfo ./app/controllers/masterBreedInfo
	GOOS=linux GOARCH=amd64 go build -o bin/subBreedInfo ./app/controllers/subBreedInfo

deps:
	go get ./...

test:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

clean:
	rm -rf ./bin/allBreeds
	rm -rf ./bin/masterBreeds
	rm -rf ./bin/subBreeds
	rm -rf ./bin/masterBreedImages
	rm -rf ./bin/subBreedImages
	rm -rf ./bin/masterBreedImageRandom
	rm -rf ./bin/subBreedImageRandom
	rm -rf ./bin/anyBreedImageRandom
	rm -rf ./bin/anyBreedMultiImageRandom
	rm -rf ./bin/masterBreedInfo
	rm -rf ./bin/subBreedInfo

start:
	sam local start-api
	#sam local start-api --env-vars environment_variables.json

sendtoaws:
	sam package --output-template-file packaged.yaml --s3-bucket dog-ceo-api-golang-$(ENVSWITCH)-sam
	sam deploy --template-file packaged.yaml --stack-name dog-ceo-api-golang-$(ENVSWITCH)-sam --capabilities CAPABILITY_IAM

# make deploy
# -- OR --
# make ENVIRONMENT=production deploy
deploy: test compile sendtoaws

build: test compile

fullstart: clean compile test start
