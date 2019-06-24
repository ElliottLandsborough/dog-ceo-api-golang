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

	key := breedUtil.GenerateBreedYamlKey(breed)
	object, err := awsUtil.GetObjectFromS3(key)

	if err != nil {
		return lambdaResponseUtil.KeyNotFoundErrorResponse(), nil
	}

	yaml := awsUtil.GetObjectContents(object)
	json := breedUtil.ParseYamlToJSON(yaml)

	return lambdaResponseUtil.InfoResponseFromString(json), nil
}

func main() {
	lambda.Start(handler)
}
