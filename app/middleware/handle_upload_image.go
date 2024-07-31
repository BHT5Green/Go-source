package middleware

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func HandleUploadImage(file *multipart.FileHeader) (string, error) {
	f, openErr := file.Open()
	if openErr != nil {
		return "", errors.New("failed to open file")
	}
	defer f.Close()

	client := s3Init()
	if client == nil {
		return "", errors.New("failed to initialize S3 client")
	}

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func s3Init() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return nil
	}

	return s3.NewFromConfig(cfg)
}
