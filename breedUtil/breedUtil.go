package lib

import (
	"bytes"
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
	// @todo: deal with the errors in this function
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

	//var responseCode int
	//var responseBody string

	// handle the error...
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
				//responseCode = 404
				//responseBody = aerr.Error()
			default:
				fmt.Println(aerr.Error())
				//responseCode = 500
				//responseBody = aerr.Error()
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			//responseCode = 500
			//responseBody = err.Error()
		}
		/*return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", responseBody),
			StatusCode: responseCode,
		}, nil*/
	}

	return response
}

func getObjectFromS3(key string) *s3.GetObjectOutput {
	// @todo: deal with the errors in this function
	bucket := os.Getenv("FILE_BUCKET_NAME")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)

	svc := s3.New(sess)

	response, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	//var responseCode int
	//var responseBody string

	// handle the error...
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
				//responseCode = 404
				//responseBody = aerr.Error()
			default:
				fmt.Println(aerr.Error())
				//responseCode = 500
				//responseBody = aerr.Error()
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			//responseCode = 500
			//responseBody = err.Error()
		}
		/*return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", responseBody),
			StatusCode: responseCode,
		}, nil*/
	}

	return response
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

	// create map of string arrays
	objects := []string{}

	// loop through results
	for _, c := range response.Contents {
		// append result to slice
		objects = append(objects, *c.Key)
	}

	return objects
}

// contains checks if a string exists in a slice of strings
func contains(a []string, x string) bool {
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

	// create map of string arrays
	breedArray := []string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		if !contains(breedArray, exploded[0]) {
			// append to breeds array
			breedArray = append(breedArray, exploded[0])
		}
	}

	return breedArray
}

// ListSubBreeds gets all sub breeds by master breed name
func ListSubBreeds(request events.APIGatewayProxyRequest) []string {
	// the breed from the {breed} section of url
	breedRequested := request.PathParameters["breed"]

	// get all breeds from s3
	breeds := GetRootPrefixesFromS3()

	// create map of string arrays
	breedArray := []string{}

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
				breedArray = append(breedArray, sub)
			}
		}
	}

	return breedArray
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
	breed := request.PathParameters["breed"]

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
	breed := request.PathParameters["breed"]

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

func getBreedInfo(breed string) string {
	key := generateBreedYamlKey(breed)
	object := getObjectFromS3(key)
	yaml := getObjectContents(object)
	json := parseYamlToJSON(yaml)

	return json
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
func ListMasterBreedInfo(request events.APIGatewayProxyRequest) string {
	// the breed from the {breed} section of url
	breed := request.PathParameters["breed"]

	return getBreedInfo(breed)
}

// ListSubBreedInfo gets the yaml file from s3 and converts it to json
func ListSubBreedInfo(request events.APIGatewayProxyRequest) string {
	// the breed from the {breed1} section of url
	masterBreed := request.PathParameters["breed1"]
	// the breed from the {breed2} section of url
	subBreed := request.PathParameters["breed2"]
	breed := masterBreed + "-" + subBreed

	return getBreedInfo(breed)
}
