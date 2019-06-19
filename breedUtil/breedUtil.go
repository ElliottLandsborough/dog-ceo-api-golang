package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetBreedPrefixesFromS3 gets all the breed fixes from s3
func GetBreedPrefixesFromS3() *s3.ListObjectsV2Output {
	// @todo: deal with the errors in this function
	bucket := os.Getenv("IMAGE_BUCKET_NAME")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)

	fmt.Println(os.Getenv("IMAGE_BUCKET_NAME"))
	fmt.Println(os.Getenv("BUCKET_REGION"))

	svc := s3.New(sess)

	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Delimiter: aws.String("/"),
		Prefix:    aws.String(""),
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

// ListAllBreeds gets all breeds (master and sub)
func ListAllBreeds() map[string][]string {
	response := GetBreedPrefixesFromS3()

	// create map of string arrays
	twoDimensionalArray := map[string][]string{}

	// loop through aws result
	for _, c := range response.CommonPrefixes {
		// remove the trailing slash
		breed := strings.TrimRight(*c.Prefix, "/")

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

// Contains checks if a string exists in a slice of strings
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ListBreeds gets all master breeds
func ListBreeds() []string {
	response := GetBreedPrefixesFromS3()

	// create map of string arrays
	breedArray := []string{}

	// loop through aws result
	for _, c := range response.CommonPrefixes {
		// remove the trailing slash
		breed := strings.TrimRight(*c.Prefix, "/")

		// explode by -
		exploded := strings.Split(breed, "-")

		if !Contains(breedArray, exploded[0]) {
			// append to breeds array
			breedArray = append(breedArray, exploded[0])
		}
	}

	return breedArray
}

// ListSubBreeds gets all sub breeds by master breed name
func ListSubBreeds(request events.APIGatewayProxyRequest) []string {
	response := GetBreedPrefixesFromS3()

	// create map of string arrays
	breedArray := []string{}

	// the breed from the {breed} section of url
	breedRequested := request.PathParameters["breed"]

	// loop through aws result
	for _, c := range response.CommonPrefixes {
		// remove the trailing slash
		breed := strings.TrimRight(*c.Prefix, "/")

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
