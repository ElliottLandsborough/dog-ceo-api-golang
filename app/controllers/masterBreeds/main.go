package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breed "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slice := aws.GetRootPrefixes()
	result := breed.ListMasterBreeds(slice)
	return response.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
