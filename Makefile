ENVSWITCH=dev
ifeq ("$(ENVIRONMENT)","production")
	ENVSWITCH=prod
endif

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/allBreeds ./app/controllers/allBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/anyBreedMultiImageRandom ./app/controllers/anyBreedMultiImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/breedImageRandom ./app/controllers/breedImageRandom
	GOOS=linux GOARCH=amd64 go build -o bin/breedImages ./app/controllers/breedImages
	GOOS=linux GOARCH=amd64 go build -o bin/breedInfo ./app/controllers/breedInfo
	GOOS=linux GOARCH=amd64 go build -o bin/masterBreeds ./app/controllers/masterBreeds
	GOOS=linux GOARCH=amd64 go build -o bin/subBreeds ./app/controllers/subBreeds

deps:
	go get ./...

test:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

clean:
	rm -rf ./bin/allBreeds
	rm -rf ./bin/anyBreedMultiImageRandom
	rm -rf ./bin/breedImageRandom
	rm -rf ./bin/breedImages
	rm -rf ./bin/breedInfo
	rm -rf ./bin/masterBreeds
	rm -rf ./bin/subBreeds

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
