package lib

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// returns a json response with status code
func jsonResponse(statusCode int, json string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       json,
		StatusCode: statusCode,
	}
}

// BreedResponseOneDimensional returns a json response
func BreedResponseOneDimensional(data []string) events.APIGatewayProxyResponse {
	successData := map[string]interface{}{
		"status":  "success",
		"message": data,
	}

	resultJSON, _ := json.Marshal(successData)

	return jsonResponse(200, string(resultJSON))
}

// BreedResponseTwoDimensional returns a json response
func BreedResponseTwoDimensional(data map[string][]string) events.APIGatewayProxyResponse {
	successData := map[string]interface{}{
		"status":  "success",
		"message": data,
	}

	resultJSON, _ := json.Marshal(successData)

	return jsonResponse(200, string(resultJSON))
}

// ImageResponseOneDimensional returns a json response
func ImageResponseOneDimensional(data []string) events.APIGatewayProxyResponse {
	successData := map[string]interface{}{
		"status":  "success",
		"message": data,
	}

	resultJSON, _ := json.Marshal(successData)

	return jsonResponse(200, string(resultJSON))
}

// InfoResponseFromString returns a json response
func InfoResponseFromString(data string) events.APIGatewayProxyResponse {
	byt := []byte(data)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		fail := map[string]interface{}{
			"status":  "error",
			"message": "data is badly formatted",
		}
		failJSON, _ := json.Marshal(fail)
		return jsonResponse(500, string(failJSON))
	}

	successData := map[string]interface{}{
		"status":  "success",
		"message": dat,
	}

	resultJSON, _ := json.Marshal(successData)

	return jsonResponse(200, string(resultJSON))
}

// KeyNotFoundErrorResponse is what happens when a breed doesnt exist
func KeyNotFoundErrorResponse() events.APIGatewayProxyResponse {
	fail := map[string]interface{}{
		"status":  "error",
		"message": "Breed not found.",
	}
	failJSON, _ := json.Marshal(fail)
	return jsonResponse(404, string(failJSON))
}
