package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../../libraries/aws"
	breedUtil "../../libraries/breed"
	lambdaResponseUtil "../../libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slice := awsUtil.GetRootPrefixesFromS3()
	result := breedUtil.ListMasterBreeds(slice)
	return lambdaResponseUtil.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
