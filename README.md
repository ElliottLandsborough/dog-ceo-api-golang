# dog.ceo golang lambda functions

[![CircleCI](https://dl.circleci.com/status-badge/img/circleci/AE794oGvVf6X3TA8Q3K9s/bYUpPd4Cc9ATboVrnJhRf/tree/main.svg?style=svg&circle-token=f466b05680041172041bf3f65dd74246cdbb49d1)](https://dl.circleci.com/status-badge/redirect/circleci/AE794oGvVf6X3TA8Q3K9s/bYUpPd4Cc9ATboVrnJhRf/tree/main)
[![codecov](https://codecov.io/gh/ElliottLandsborough/dog-ceo-api-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/ElliottLandsborough/dog-ceo-api-golang)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.x-success.svg)](https://golang.org/)
[![AWS Lambda](https://img.shields.io/badge/AWS-Lambda-orange.svg)](https://aws.amazon.com/lambda/)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/40c5e0b1db42449b91e0a4a0f5a0dcdf)](https://www.codacy.com/app/ElliottLandsborough/dog-ceo-api-golang?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ElliottLandsborough/dog-ceo-api-golang&amp;utm_campaign=Badge_Grade)

The AWS Lambda functions used for the [dog.ceo api](https://dog.ceo/api).

[Old version in node](https://github.com/ElliottLandsborough/dog-ceo-api-node).

## Quick start

```shell
go get -u github.com/ElliottLandsborough/dog-ceo-api-golang
cd $GOPATH/src/github.com/ElliottLandsborough/dog-ceo-api-golang
make deps
make clean
make build
make start
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Dependencies

```shell
go get ./...
```

### Compiling

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

### Testing

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
