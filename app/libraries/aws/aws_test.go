package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

type ObjectStore struct {
	Client s3iface.S3API
	URL    string
}

func (svc ObjectStore) GetObject(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {

	return svc.Client.GetObject(in)
}

type mockGetObject struct {
	s3iface.S3API
	Resp s3.GetObjectOutput
}

func (m mockGetObject) GetObject(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return &m.Resp, nil
}

func TestS3svc(t *testing.T) {
	svc, _ := S3svc("fsdfsadfas")

	expected := "*s3.S3"
	got := reflect.TypeOf(svc).String()

	assert.Equal(t, expected, got)
}

func TestObjectInputGen(t *testing.T) {
	result := ObjectInputGen("bucket", "key")

	assert.Equal(t, reflect.TypeOf(result).String(), "*s3.GetObjectInput")
}

func TestObjectsV2InputGen(t *testing.T) {
	result := ObjectsV2InputGen("bucket", "delimeter", "prefix")

	assert.Equal(t, reflect.TypeOf(result).String(), "*s3.ListObjectsV2Input")
}

func TestGetObjectContents(t *testing.T) {
	// Create mock receiver
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	expected := "hello world"

	// r type is io.ReadCloser
	r := ioutil.NopCloser(bytes.NewReader([]byte(expected)))

	// mock svc
	svc := ObjectStore{
		Client: mockGetObject{Resp: s3.GetObjectOutput{Body: r}},
		URL:    ts.URL,
	}

	input := ObjectInputGen("foo", "bar")

	resp, err := svc.GetObject(input)

	if err != nil {
		t.Fatal(err)
	}

	got := GetObjectContents(resp)

	assert.Equal(t, expected, got)
}

func TestPrefixesToSlice(t *testing.T) {
	bucketName := "testBucket"
	Prefix1 := "Prefix1"
	Prefix2 := "Prefix2"

	output := &s3.ListObjectsV2Output{
		Name: &bucketName,
		CommonPrefixes: []*s3.CommonPrefix{
			{
				Prefix: &Prefix1,
			},
			{
				Prefix: &Prefix2,
			},
		},
	}

	got := PrefixesToSlice(output)

	expected := []string{Prefix1, Prefix2}

	assert.Equal(t, got, expected)
}

func TestObjectsToSlice(t *testing.T) {
	bucketName := "testBucket"
	expectedKey1 := "Object1"
	expectedKey2 := "Object2"

	output := &s3.ListObjectsV2Output{
		Name: &bucketName,
		Contents: []*s3.Object{
			{
				Key: &expectedKey1,
			},
			{
				Key: &expectedKey2,
			},
		},
	}

	expected := []string{expectedKey1, expectedKey2}

	assert.Equal(t, expected, ObjectsToSlice(output))
}

// Define a mock struct to be used in your unit tests of myFunc.
type mockS3Client struct {
	s3iface.S3API
}

// Example from https://godoc.org/github.com/aws/aws-sdk-go/service/s3/s3iface
func (m *mockS3Client) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	bucketName := "testBucket"
	expectedKey1 := "breed-name/image1.jpg"
	expectedKey2 := "breed-name/image2.jpg"

	// mock response/functionality
	output := &s3.ListObjectsV2Output{
		Name: &bucketName,
		Contents: []*s3.Object{
			{
				Key: &expectedKey1,
			},
			{
				Key: &expectedKey2,
			},
		},
	}

	return output, nil
}

func TestGetObjectsByDelimeterAndPrefix(t *testing.T) {
	// test bad bucket response
	svc1, _ := S3svc("eu-west-1")
	result1 := GetObjectsByDelimeterAndPrefix(svc1, "testBucket", "", "breed-name")
	output1 := &s3.ListObjectsV2Output{}
	assert.Equal(t, output1, result1)

	// test good bucket response
	svc2 := &mockS3Client{}
	result2 := GetObjectsByDelimeterAndPrefix(svc2, "testBucket", "", "breed-name")
	slice := ObjectsToSlice(result2)
	assert.Equal(t, []string{"breed-name/image1.jpg", "breed-name/image2.jpg"}, slice)
}
