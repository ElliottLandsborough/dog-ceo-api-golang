package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breedUtil "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rootPrefixes := aws.GetRootPrefixes()
	breed := breedUtil.GetRandomItemFromSliceString(rootPrefixes)
	images := aws.GetObjectsByPrefix(breed)
	result := breedUtil.ListBreedImageRandom(images)
	return response.ImageResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
