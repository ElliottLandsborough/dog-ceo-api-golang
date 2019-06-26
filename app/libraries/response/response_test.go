package lib

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, wanted, response)
}

func TestBreedResponseOneDimensional(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := BreedResponseOneDimensional(s)

	assert.Equal(t, 200, result.StatusCode)

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	assert.Equal(t, expectedHeaders, result.Headers)

	expectedJSON := `{"message":["string1","string2","string3","string4","string5"],"status":"success"}`

	assert.Equal(t, expectedJSON, result.Body)
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

	assert.Equal(t, 200, result.StatusCode)

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	assert.Equal(t, expectedHeaders, result.Headers)

	expectedJSON := `{"message":{"breed1":["subbreed1","subbreed2"],"breed2":[]},"status":"success"}`

	assert.Equal(t, expectedJSON, result.Body)
}

func TestImageResponseOneDimensional(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := ImageResponseOneDimensional(s)

	assert.Equal(t, 200, result.StatusCode)

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	assert.Equal(t, expectedHeaders, result.Headers)

	expectedJSON := `{"message":["string1","string2","string3","string4","string5"],"status":"success"}`

	assert.Equal(t, expectedJSON, result.Body)
}

func TestInfoResponseFromString(t *testing.T) {
	someJSON := `{"message":["string1","string2","string3","string4","string5"],"status":"success"}`
	result := InfoResponseFromString(someJSON)

	assert.Equal(t, 200, result.StatusCode)

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	assert.Equal(t, expectedHeaders, result.Headers)

	expectedJSON := `{"message":{"message":["string1","string2","string3","string4","string5"],"status":"success"},"status":"success"}`

	assert.Equal(t, expectedJSON, result.Body)

	someBadJSON := `%$£`
	result2 := InfoResponseFromString(someBadJSON)

	assert.Equal(t, 500, result2.StatusCode)

	assert.Equal(t, expectedHeaders, result2.Headers)

	expectedJSON2 := `{"message":"data is badly formatted","status":"error"}`

	assert.Equal(t, expectedJSON2, result2.Body)
}

func TestKeyNotFoundErrorResponse(t *testing.T) {
	result := KeyNotFoundErrorResponse()

	assert.Equal(t, 404, result.StatusCode)

	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	assert.Equal(t, expectedHeaders, result.Headers)

	expectedJSON := `{"message":"Breed not found.","status":"error"}`

	assert.Equal(t, expectedJSON, result.Body)
}
