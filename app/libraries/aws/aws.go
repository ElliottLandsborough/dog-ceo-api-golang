package lib

import (
	"bytes"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3svc
func S3svc(region string) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := s3.New(sess)

	return svc, err
}

// ObjectInputGen
func ObjectInputGen(bucket string, key string) *s3.GetObjectInput {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	return input
}

// ObjectsV2InputGen
func ObjectsV2InputGen(bucket string, delimeter string, prefix string) *s3.ListObjectsV2Input {
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Delimiter: aws.String(delimeter),
		Prefix:    aws.String(prefix),
		MaxKeys:   aws.Int64(1000000),
	}

	return input
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

// GetObjectsByDelimeterAndPrefix gets objects from s3 which start with string
func GetObjectsByDelimeterAndPrefix(svc *s3.S3, bucket string, delimeter string, prefix string) []string {
	input := ObjectsV2InputGen(bucket, "", prefix)
	objects, _ := svc.ListObjectsV2(input)
	return PrefixesToSlice(objects)
}
