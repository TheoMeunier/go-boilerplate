package storage

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3StorageAdapter struct {
	Client     *s3.Client
	BucketName string
}

func (s *S3StorageAdapter) Upload(path string, data []byte) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(path),
		Body:   bytes.NewReader(data),
	})
	return err
}

func (s *S3StorageAdapter) Download(path string) ([]byte, error) {
	obj, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()
	return io.ReadAll(obj.Body)
}

func (s *S3StorageAdapter) Delete(path string) error {
	_, err := s.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(path),
	})
	return err
}
