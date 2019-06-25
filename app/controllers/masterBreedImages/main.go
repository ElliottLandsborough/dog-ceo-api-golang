package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]
	result := aws.GetObjectsByPrefix(breed)
	return response.ImageResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
