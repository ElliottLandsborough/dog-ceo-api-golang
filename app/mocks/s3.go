package mocks

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// heavily inspired by https://github.com/coinbase/step/blob/d787f293f5d480bb28f96327e470ef8a705f4cd1/aws/s3/s3_test.go

// GetObjectResponse
type GetObjectResponse struct {
	Resp  *s3.GetObjectOutput
	Body  string
	Error error
}

// MockS3Client
type MockS3Client struct {
	s3iface.S3API

	GetObjectResp map[string]*GetObjectResponse
}

func (m *MockS3Client) init() {
	if m.GetObjectResp == nil {
		m.GetObjectResp = map[string]*GetObjectResponse{}
	}
}

// MakeS3Body
func MakeS3Body(ret string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(ret))
}

// AWSS3NotFoundError
func AWSS3NotFoundError() error {
	return awserr.New(s3.ErrCodeNoSuchKey, "not found", nil)
}

// GetObject
func (m *MockS3Client) GetObject(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	m.init()
	resp := m.GetObjectResp[*in.Key]

	if resp == nil {
		return nil, AWSS3NotFoundError()
	}

	resp.Resp.Body = MakeS3Body(resp.Body)
	return resp.Resp, resp.Error
}

// ListObjects
func (m *MockS3Client) ListObjects(in *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return nil, nil
}
