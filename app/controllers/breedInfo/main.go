package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	aws "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/aws"
	breedUtil "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/breed"
	response "github.com/ElliottLandsborough/dog-ceo-api-golang/app/libraries/response"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	region := os.Getenv("BUCKET_REGION")
	svc, _ := aws.S3svc(region)
	bucket := os.Getenv("FILE_BUCKET_NAME")

	breed := breedUtil.GetBreedFromPathParams(request.PathParameters)

	key := breedUtil.GenerateBreedYamlKey(breed)

	input := aws.ObjectInputGen(bucket, key)

	object, err := svc.GetObject(input)

	/*
		// handle the error...
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeNoSuchBucket:
					fmt.Println(aerr.Error())
					os.Exit(1)
				case s3.ErrCodeNoSuchKey:
					fmt.Println(aerr.Error())
					return nil, aerr
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}
	*/

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
