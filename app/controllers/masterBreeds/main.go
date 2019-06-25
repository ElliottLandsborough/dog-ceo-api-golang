package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breed "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	region := os.Getenv("BUCKET_REGION")
	svc, _ := aws.S3svc(region)
	bucket := os.Getenv("IMAGE_BUCKET_NAME")

	slice := aws.GetRootPrefixes(svc, bucket)
	result := breed.ListMasterBreeds(slice)
	return response.BreedResponseOneDimensional(result), nil
}

func main() {
	lambda.Start(handler)
}
