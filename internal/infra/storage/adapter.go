package storage

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/net/context"
)

type FileStorage interface {
	Upload(path string, data []byte) error
	Download(path string) ([]byte, error)
	Delete(path string) error
}

func InitStorage() (FileStorage, error) {
	env := os.Getenv("FILE_STORAGE_ENV")

	switch env {
	case "local":
		return &LocalStorageAdapter{
			BasePath: "./local_storage",
		}, nil

	case "s3":
		region := os.Getenv("AWS_REGION")
		bucketName := os.Getenv("AWS_BUCKET_NAME")
		endpoint := os.Getenv("S3_ENDPOINT")
		accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				accessKey,
				secretKey,
				"",
			)),
		)

		if err != nil {
			return nil, fmt.Errorf("erreur de configuration AWS: %w", err)
		}

		client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			if endpoint != "" {
				o.BaseEndpoint = aws.String(endpoint)
				o.UsePathStyle = true
			}
		})

		return &S3StorageAdapter{
			Client:     client,
			BucketName: bucketName,
		}, nil

	default:
		return nil, fmt.Errorf("FILE_STORAGE_ENV inconnu: %s (valeurs possibles: local, s3)", env)
	}
}
