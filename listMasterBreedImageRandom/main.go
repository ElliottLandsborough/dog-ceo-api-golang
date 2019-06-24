package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../awsUtil"
	breedUtil "../breedUtil"
	lambdaResponseUtil "../lambdaResponseUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]
	images := awsUtil.GetObjectsByPrefix(breed)
	result := breedUtil.ListBreedImageRandom(images)
	return lambdaResponseUtil.ImageResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
