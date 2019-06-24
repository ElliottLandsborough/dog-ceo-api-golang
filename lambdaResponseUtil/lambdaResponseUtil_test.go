package lib

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestJsonResponse(t *testing.T) {
	statusCode := 418
	json := "{ \"name\":\"John\", \"age\":30, \"car\":null }"

	response := jsonResponse(statusCode, json)

	wanted := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       json,
		StatusCode: statusCode,
	}

	if reflect.DeepEqual(response, wanted) != true {
		t.Errorf("Incorrect, got: %T, want: %T.", response, wanted)
	}
}
