package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breed "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rootPrefixes := aws.GetRootPrefixes()
	randomPrefix := breed.GetRandomItemFromSliceString(rootPrefixes)
	slice := aws.GetObjectsByPrefix(randomPrefix)
	count := request.PathParameters["count"]
	result := breed.ListAnyBreedMultiImageRandom(slice, count)
	return response.ImageResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
