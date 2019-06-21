package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ghodss/yaml"
)

// ListObjectsFromS3 gets all the breed fixes from s3
func ListObjectsFromS3(delimeter string, prefix string) *s3.ListObjectsV2Output {
	bucket := os.Getenv("IMAGE_BUCKET_NAME")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)

	svc := s3.New(sess)

	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Delimiter: aws.String(delimeter),
		Prefix:    aws.String(prefix),
		MaxKeys:   aws.Int64(1000000),
	}

	response, err := svc.ListObjectsV2(input)

	// handle the error...
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(aerr.Error())
				os.Exit(1)
			case s3.ErrCodeNoSuchKey:
				fmt.Println(aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	return response
}

func getObjectFromS3(key string) (*s3.GetObjectOutput, error) {
	bucket := os.Getenv("FILE_BUCKET_NAME")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)

	svc := s3.New(sess)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	response, err := svc.GetObject(input)

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

	return response, err
}

func getObjectContents(object *s3.GetObjectOutput) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	s := buf.String() // Does a complete copy of the bytes in the buffer.

	return s
}

// GetRootPrefixesFromS3 gets all breed prefixes from s3
func GetRootPrefixesFromS3() []string {
	return prefixesToSlice(ListObjectsFromS3("/", ""))
}

// converts listObjectsV2Output response with prefixes to string slice
func prefixesToSlice(listObjectsV2Output *s3.ListObjectsV2Output) []string {
	breeds := []string{}

	// loop through aws result
	for _, c := range listObjectsV2Output.CommonPrefixes {
		breed := strings.TrimRight(*c.Prefix, "/")
		breeds = append(breeds, breed)
	}

	return breeds
}

// get objects from s3 which start with string
func getObjectsByPrefix(prefix string) []string {
	// get all objects from prefix* on s3
	response := ListObjectsFromS3("", prefix)

	// slice of strings
	objects := []string{}

	// loop through results
	for _, c := range response.Contents {
		cdn := os.Getenv("CDN_DOMAIN_PREFIX")
		url := cdn + *c.Key
		// append result to slice
		objects = append(objects, url)
	}

	return objects
}

// sliceContainsString checks if a string exists in a slice of strings
func sliceContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func getRandomItemFromSliceString(slice []string) string {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// pick random string from slice
	return slice[rand.Intn(len(slice))]
}

// get all breeds from s3 and pick a random one
func getRandomPrefix() string {
	return getRandomItemFromSliceString(GetRootPrefixesFromS3())
}

// ListAllBreeds gets all breeds (master and sub)
func ListAllBreeds() map[string][]string {
	// get all breeds from s3
	breeds := GetRootPrefixesFromS3()

	// create map of string arrays
	twoDimensionalArray := map[string][]string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		// master breed will always be at 0
		master := exploded[0]

		_, ok := twoDimensionalArray[master]

		// master breed isn't in 2d array yet, add it
		if !ok {
			twoDimensionalArray[master] = []string{}
		}

		// sub breed exists?
		if len(exploded) > 1 {
			// sub will always be 1
			sub := exploded[1]

			// append item to slice
			twoDimensionalArray[master] = append(twoDimensionalArray[master], sub)
		}
	}

	return twoDimensionalArray
}

// ListBreeds gets all master breeds
func ListBreeds() []string {
	// get all breeds from s3
	breeds := GetRootPrefixesFromS3()

	// slice of strings
	s := []string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		if !sliceContainsString(s, exploded[0]) {
			// append to breeds
			s = append(s, exploded[0])
		}
	}

	return s
}

// ListSubBreeds gets all sub breeds by master breed name
func ListSubBreeds(request events.APIGatewayProxyRequest) []string {
	// the breed from the {breed} section of url
	breedRequested := request.PathParameters["breed1"]

	// get all breeds from s3
	breeds := GetRootPrefixesFromS3()

	// slice of strings
	s := []string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		// primary breed will always be there
		primary := exploded[0]

		// does the url segment match this item?
		if breedRequested == primary {
			// sub breed exists?
			if len(exploded) > 1 {
				// sub will always be 1
				sub := exploded[1]

				// append item to slice
				s = append(s, sub)
			}
		}
	}

	return s
}

func getRandomBreedImageByBreedString(breed string) string {
	// get all images of a breed
	images := getObjectsByPrefix(breed)

	// pick random image from slice
	image := getRandomItemFromSliceString(images)

	return image
}

func shuffleSlice(slice []string) []string {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// shuffle the items
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}

func getMultipleRandomItemsFromSliceString(slice []string, amount int) []string {
	// shuffle the items
	slice = shuffleSlice(slice)

	// dont bother if we want all the items
	if amount > len(slice) {
		return slice
	}

	// return {amount} items from slice
	return slice[0:amount]
}

// ListMasterBreedImages gets all images from a master breed
func ListMasterBreedImages(request events.APIGatewayProxyRequest) []string {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	return getObjectsByPrefix(breed)
}

// ListSubBreedImages gets all images from a sub breed
func ListSubBreedImages(request events.APIGatewayProxyRequest) []string {
	// the breed from the {breed1} section of url
	masterBreed := request.PathParameters["breed1"]
	// the breed from the {breed2} section of url
	subBreed := request.PathParameters["breed2"]
	breed := masterBreed + "-" + subBreed

	return getObjectsByPrefix(breed)
}

// ListMasterBreedImageRandom gets a random image from all the master breed images
func ListMasterBreedImageRandom(request events.APIGatewayProxyRequest) []string {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	return []string{getRandomBreedImageByBreedString(breed)}
}

// ListSubBreedImageRandom gets a random image from all the sub breed images
func ListSubBreedImageRandom(request events.APIGatewayProxyRequest) []string {
	// the breed from the {breed1} section of url
	masterBreed := request.PathParameters["breed1"]
	// the breed from the {breed2} section of url
	subBreed := request.PathParameters["breed2"]
	breed := masterBreed + "-" + subBreed

	return []string{getRandomBreedImageByBreedString(breed)}
}

// ListAnyBreedImageRandom gets random breed, gets all images, returns random image
func ListAnyBreedImageRandom() []string {
	return []string{getRandomBreedImageByBreedString(getRandomPrefix())}
}

// ListAnyBreedMultiImageRandom gets all images from a random breed, returns {count} images
func ListAnyBreedMultiImageRandom(request events.APIGatewayProxyRequest) []string {
	s := request.PathParameters["count"]

	// string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		i = 1
	}

	return getMultipleRandomItemsFromSliceString(getObjectsByPrefix(getRandomPrefix()), i)
}

func generateBreedYamlKey(breed string) string {
	return "breed-info/" + breed + ".yaml"
}

func getBreedInfo(breed string) (string, error) {
	key := generateBreedYamlKey(breed)
	object, err := getObjectFromS3(key)

	if err != nil {
		return "{}", err
	}

	yaml := getObjectContents(object)
	json := parseYamlToJSON(yaml)

	return json, nil
}

func parseYamlToJSON(yamlString string) string {

	data, err := yaml.YAMLToJSON([]byte(yamlString))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "error"
	}

	return string(data)
}

// ListMasterBreedInfo gets the yaml file from s3 and converts it to json
func ListMasterBreedInfo(request events.APIGatewayProxyRequest) (string, error) {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed1"]

	info, err := getBreedInfo(breed)

	return info, err
}

// ListSubBreedInfo gets the yaml file from s3 and converts it to json
func ListSubBreedInfo(request events.APIGatewayProxyRequest) (string, error) {
	// the breed from the {breed1} section of url
	masterBreed := request.PathParameters["breed1"]
	// the breed from the {breed2} section of url
	subBreed := request.PathParameters["breed2"]
	breed := masterBreed + "-" + subBreed

	info, err := getBreedInfo(breed)

	return info, err
}

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
