package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../awsUtil"
	breedUtil "../breedUtil"
	lambdaResponseUtil "../lambdaResponseUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get all breeds from s3
	breeds := awsUtil.GetRootPrefixesFromS3()

	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	result := breedUtil.ListSubBreeds(breed, breeds)

	return lambdaResponseUtil.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
