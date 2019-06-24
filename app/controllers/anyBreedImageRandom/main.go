package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../../libraries/aws"
	breedUtil "../../libraries/breed"
	lambdaResponseUtil "../../libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rootPrefixes := awsUtil.GetRootPrefixesFromS3()
	breed := breedUtil.GetRandomItemFromSliceString(rootPrefixes)
	images := awsUtil.GetObjectsByPrefix(breed)
	result := breedUtil.ListBreedImageRandom(images)
	return lambdaResponseUtil.ImageResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
