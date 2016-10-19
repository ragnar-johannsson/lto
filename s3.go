package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func generateSignedURL(key string) string {
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(Config.S3Region),
		Credentials: credentials.NewStaticCredentials(
			Config.S3AccessKey,
			Config.S3SecretKey,
			"",
		),
	}))

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(Config.S3Bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(Config.UrlTTL)
	if err != nil {
		log.Fatal("Error generating a signed URL: ", err)
	}

	return url
}
