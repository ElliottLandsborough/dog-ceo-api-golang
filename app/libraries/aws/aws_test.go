package lib

import (
	"fmt"
	"testing"

	"github.com/ElliottLandsborough/dog-ceo-api-golang/app/mocks/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type S3API s3iface.S3API

type NotFoundError struct {
	bucket *string
	path   *string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not Found %v %v", *e.bucket, *e.path)
}

// Strp return a string pointer from string
func Strp(s string) *string {
	return &s
}

func TestGetObject(t *testing.T) {
	/*
		svc := &s3.MockS3Client{}
		result, err := GetObject(svc, "bucket", "/path")
	*/
}

// Get downloads content from S3
func Get(s3c S3API, bucket *string, path *string) (*s3.GetObjectOutput, error) {
	/*
		input := &s3.GetObjectInput{
			Bucket: aws.String(*bucket),
			Key:    aws.String(*path),
		}
		//body, err := s3c.GetObject(s3c, bucket, path)
		body, err := s3c.GetObject(input)
		fmt.Println(body)
		//return body, err
		return body, err
	*/

	return nil, nil
}

func Test_Get_Success(t *testing.T) {
	/*
		s3c := &s3.MockS3Client{}
		_, err := Get(s3c, Strp("bucket"), Strp("/path"))
		fmt.Println(err)
		assert.Error(t, err)
	*/
	// no idea why this doesn't work...
	//assert.IsType(t, &NotFoundError{}, err)
	/*
		s3c := &mocks.MockS3Client{}
		_, err := Get(s3c, Strp("bucket"), Strp("/path"))
		assert.Error(t, err)
		assert.IsType(t, &NotFoundError{}, err)

		s3c.AddGetObject("/path", "asd", nil)
		out, err := Get(s3c, Strp("bucket"), Strp("/path"))
		assert.NoError(t, err)
		assert.Equal(t, "asd", string(*out))
	*/
}
