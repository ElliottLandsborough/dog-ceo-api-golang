package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../../libraries/aws"
	breedUtil "../../libraries/breed"
	lambdaResponseUtil "../../libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get all breeds from s3
	breeds := awsUtil.GetRootPrefixes()

	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	result := breedUtil.ListSubBreeds(breed, breeds)

	return lambdaResponseUtil.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
