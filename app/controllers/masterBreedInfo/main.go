package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breedUtil "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	key := breedUtil.GenerateBreedYamlKey(breed)
	object, err := aws.GetObject(key)

	if err != nil {
		return response.KeyNotFoundErrorResponse(), nil
	}

	yaml := aws.GetObjectContents(object)
	json := breedUtil.ParseYamlToJSON(yaml)

	return response.InfoResponseFromString(json), nil
}

func main() {
	lambda.Start(handler)
}
