package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../awsUtil"
	breedUtil "../breedUtil"
	lambdaResponseUtil "../lambdaResponseUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slice := awsUtil.GetRootPrefixesFromS3()
	result := breedUtil.ListAllBreeds(slice)
	return lambdaResponseUtil.BreedResponseTwoDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
