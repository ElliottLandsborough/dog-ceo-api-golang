# dog.ceo golang lambda functions

The AWS Lambda functions used for the https://dog.ceo api.

Old version in node: https://github.com/ElliottLandsborough/dog-ceo-api-node

## Quick start

```
$ make deps
$ make clean
$ make build
$ make start
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Dependencies

```shell
go get -u github.com/aws/aws-lambda-go/...
go get -u github.com/aws/aws-sdk-go
go get github.com/ghodss/yaml
```

### Building

```shell
make
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatibility.

### Local development

**Invoking function locally through local API Gateway**

```shell
make build && make start
```

## Packaging and deployment

```shell
make deploy
-- OR --
make ENVIRONMENT=production deploy
```

### Testing (not yet implemented)

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
make test
```

### Example environment_variables.json
```json
{
  "listAllBreeds": {
  	"IMAGE_BUCKET_NAME": "dog-ceo-stanford-files",
    "FILE_BUCKET_NAME": "dog-ceo-api-static-content-dev",
    "BUCKET_REGION": "eu-west-1",
    "CDN_DOMAIN_PREFIX": "https://images.dog.ceo/breeds/"
  },
  "listBreeds": {
  	"IMAGE_BUCKET_NAME": "dog-ceo-stanford-files",
    "FILE_BUCKET_NAME": "dog-ceo-api-static-content-dev",
    "BUCKET_REGION": "eu-west-1",
    "CDN_DOMAIN_PREFIX": "https://images.dog.ceo/breeds/"
  }
}
```
