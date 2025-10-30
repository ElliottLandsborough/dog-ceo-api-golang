ENVSWITCH=dev
ifeq ("$(ENVIRONMENT)","production")
	ENVSWITCH=prod
endif

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/allBreeds/bootstrap ./app/controllers/allBreeds
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/anyBreedMultiImageRandom/bootstrap ./app/controllers/anyBreedMultiImageRandom
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/breedImageRandom/bootstrap ./app/controllers/breedImageRandom
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/breedImages/bootstrap ./app/controllers/breedImages
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/breedInfo/bootstrap ./app/controllers/breedInfo
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/masterBreeds/bootstrap ./app/controllers/masterBreeds
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bin/subBreeds/bootstrap ./app/controllers/subBreeds

deps:
	go mod download

test:
	GOTOOLCHAIN=go1.25.3+auto go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

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
	# sam local start-api --env-vars environment_variables.json

sendtoaws:
	sam package --output-template-file packaged.yaml --s3-bucket dog-ceo-api-golang-$(ENVSWITCH)-sam
	sam deploy --template-file packaged.yaml --stack-name dog-ceo-api-golang-$(ENVSWITCH)-sam --capabilities CAPABILITY_IAM

# make deploy
# -- OR --
# make ENVIRONMENT=production deploy
deploy: clean test clean compile sendtoaws

build: test compile

fullstart: clean compile test start
