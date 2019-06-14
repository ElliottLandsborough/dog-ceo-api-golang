package main

import (
	//"errors"
	"fmt"
	//"io/ioutil"
	//"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	//"github.com/aws/aws-sdk-go/aws"
    //"github.com/aws/aws-sdk-go/aws/session"
)

var (
	/*
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
	*/
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
	sess, err := session.NewSession(&aws.Config{
    	Region: aws.String("eu-west-1")},
	)
	*/

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", os.Getenv("BUCKET_NAME")),
		StatusCode: 200,
	}, nil
	/*
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(ip)),
		StatusCode: 200,
	}, nil
	*/
}

func main() {
	lambda.Start(handler)
}
