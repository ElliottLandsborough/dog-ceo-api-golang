package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	breedUtil "../breedUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	result := breedUtil.ListSubBreeds(request)

	resultJSON, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		Body:       string(resultJSON),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
