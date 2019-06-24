package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsUtil "../../libraries/aws"
	breedUtil "../../libraries/breed"
	lambdaResponseUtil "../../libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// the breed from the {breed1} section of url
	masterBreed := request.PathParameters["breed1"]
	// the breed from the {breed2} section of url
	subBreed := request.PathParameters["breed2"]
	breed := masterBreed + "-" + subBreed

	key := breedUtil.GenerateBreedYamlKey(breed)
	object, err := awsUtil.GetObject(key)

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
