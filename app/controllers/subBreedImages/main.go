package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../../libraries/aws"
	lambdaResponseUtil "../../libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// the breed from the {breed1} section of url
	masterBreed := request.PathParameters["breed1"]
	// the breed from the {breed2} section of url
	subBreed := request.PathParameters["breed2"]
	breed := masterBreed + "-" + subBreed

	result := awsUtil.GetObjectsByPrefix(breed)

	return lambdaResponseUtil.ImageResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
