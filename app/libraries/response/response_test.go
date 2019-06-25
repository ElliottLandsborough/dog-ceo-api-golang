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

func TestBreedResponseOneDimensional(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := BreedResponseOneDimensional(s)

	if result.StatusCode != 200 {
		t.Errorf("Incorrect, got: %d, want: %d.", result.StatusCode, 200)
	}

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if reflect.DeepEqual(expectedHeaders, result.Headers) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Headers, expectedHeaders)
	}

	expectedJSON := `{"message":["string1","string2","string3","string4","string5"],"status":"success"}`

	if expectedJSON != result.Body {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Body, expectedJSON)
	}
}

func TestBreedResponseTwoDimensional(t *testing.T) {
	s := map[string][]string{
		"breed1": []string{
			"subbreed1",
			"subbreed2",
		},
		"breed2": []string{},
	}

	result := BreedResponseTwoDimensional(s)

	if result.StatusCode != 200 {
		t.Errorf("Incorrect, got: %d, want: %d.", result.StatusCode, 200)
	}

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if reflect.DeepEqual(expectedHeaders, result.Headers) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Headers, expectedHeaders)
	}

	expectedJSON := `{"message":{"breed1":["subbreed1","subbreed2"],"breed2":[]},"status":"success"}`

	if expectedJSON != result.Body {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Body, expectedJSON)
	}
}

func TestImageResponseOneDimensional(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := ImageResponseOneDimensional(s)

	if result.StatusCode != 200 {
		t.Errorf("Incorrect, got: %d, want: %d.", result.StatusCode, 200)
	}

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if reflect.DeepEqual(expectedHeaders, result.Headers) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Headers, expectedHeaders)
	}

	expectedJSON := `{"message":["string1","string2","string3","string4","string5"],"status":"success"}`

	if expectedJSON != result.Body {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Body, expectedJSON)
	}
}

func TestInfoResponseFromString(t *testing.T) {
	someJSON := `{"message":["string1","string2","string3","string4","string5"],"status":"success"}`
	result := InfoResponseFromString(someJSON)

	if result.StatusCode != 200 {
		t.Errorf("Incorrect, got: %d, want: %d.", result.StatusCode, 200)
	}

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if reflect.DeepEqual(expectedHeaders, result.Headers) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Headers, expectedHeaders)
	}

	expectedJSON := `{"message":{"message":["string1","string2","string3","string4","string5"],"status":"success"},"status":"success"}`

	if expectedJSON != result.Body {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Body, expectedJSON)
	}

	someBadJSON := `%$£`
	result2 := InfoResponseFromString(someBadJSON)

	if result2.StatusCode != 500 {
		t.Errorf("Incorrect, got: %d, want: %d.", result2.StatusCode, 500)
	}

	if reflect.DeepEqual(expectedHeaders, result2.Headers) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", result2.Headers, expectedHeaders)
	}

	expectedJSON2 := `{"message":"data is badly formatted","status":"error"}`

	if expectedJSON2 != result2.Body {
		t.Errorf("Incorrect, got: %s, want: %s.", result2.Body, expectedJSON2)
	}
}

func TestKeyNotFoundErrorResponse(t *testing.T) {
	result := KeyNotFoundErrorResponse()

	if result.StatusCode != 404 {
		t.Errorf("Incorrect, got: %d, want: %d.", result.StatusCode, 404)
	}

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if reflect.DeepEqual(expectedHeaders, result.Headers) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Headers, expectedHeaders)
	}

	expectedJSON := `{"message":"Breed not found.","status":"error"}`

	if expectedJSON != result.Body {
		t.Errorf("Incorrect, got: %s, want: %s.", result.Body, expectedJSON)
	}
}
