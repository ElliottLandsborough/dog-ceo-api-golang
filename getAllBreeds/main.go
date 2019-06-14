package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"

    breedUtil "../breedUtil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var responseCode int
	var responseBody string
	var bucket = os.Getenv("IMAGE_BUCKET_NAME");

	sess, err := session.NewSession(&aws.Config{
    	Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)

	svc := s3.New(sess)

	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		Delimiter:  aws.String("/"),
		Prefix:  aws.String(""),
		MaxKeys: aws.Int64(1000000),
	}

	response, err := svc.ListObjectsV2(input)

	// handle the error...
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
				responseCode = 404
				responseBody = aerr.Error()
			default:
				fmt.Println(aerr.Error())
				responseCode = 500
				responseBody = aerr.Error()
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			responseCode = 500
			responseBody = err.Error()
		}
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", responseBody),
			StatusCode: responseCode,
		}, nil
	}

	var result []string
	for _, c := range response.CommonPrefixes {
		item := strings.TrimRight(*c.Prefix, "/")
		//item = strings.TrimLeft(item, "page/")
		result = append(result, item)
	}

	resultJson, _ := json.Marshal(result)

	breedUtil.Demo()

	return events.APIGatewayProxyResponse{
		Body:       string(resultJson),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
