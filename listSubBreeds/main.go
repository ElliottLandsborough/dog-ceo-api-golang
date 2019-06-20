package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	breedUtil "../breedUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result := breedUtil.ListSubBreeds(request)
	return breedUtil.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
