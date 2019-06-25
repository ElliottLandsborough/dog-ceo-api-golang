package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breedUtil "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get all breeds from s3
	breeds := aws.GetRootPrefixes()

	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	result := breedUtil.ListSubBreeds(breed, breeds)

	return response.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
