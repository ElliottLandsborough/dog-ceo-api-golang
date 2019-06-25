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
