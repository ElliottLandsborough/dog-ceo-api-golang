package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	breedUtil "../breedUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := breedUtil.ListSubBreedInfo(request)

	if err != nil {
		return breedUtil.KeyNotFoundErrorResponse(), nil
	}

	return breedUtil.InfoResponseFromString(result), nil
}

func main() {
	lambda.Start(handler)
}
