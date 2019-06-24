package lib

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListObjects search s3 for objects with delimeter, string or both
func ListObjects(delimeter string, prefix string) *s3.ListObjectsV2Output {
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

// GetObject gets obect from s3 which matches 'key'
func GetObject(key string) (*s3.GetObjectOutput, error) {
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

// GetObjectContents gets the contents of an object from s3
func GetObjectContents(object *s3.GetObjectOutput) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	s := buf.String() // Does a complete copy of the bytes in the buffer.

	return s
}

// PrefixesToSlice converts listObjectsV2Output response with prefixes to string slice
func PrefixesToSlice(listObjectsV2Output *s3.ListObjectsV2Output) []string {
	prefixes := []string{}

	// loop through aws result
	for _, c := range listObjectsV2Output.CommonPrefixes {
		prefix := strings.TrimRight(*c.Prefix, "/")
		prefixes = append(prefixes, prefix)
	}

	return prefixes
}

// GetRootPrefixes gets all root prefixes from s3
func GetRootPrefixes() []string {
	return PrefixesToSlice(ListObjects("/", ""))
}

// GetObjectsByPrefix gets objects from s3 which start with string
func GetObjectsByPrefix(prefix string) []string {
	// get all objects from prefix* on s3
	response := ListObjects("", prefix)

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
