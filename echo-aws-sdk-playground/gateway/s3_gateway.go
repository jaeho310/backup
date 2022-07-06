package gateway

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

var (
	client          *s3.Client
	presignedClient *s3.PresignClient
)

func init() {
	log.Println("load config")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create s3 client")
	client = s3.NewFromConfig(cfg)
	presignedClient = s3.NewPresignClient(client)
}

func GetPreSignedUrlForUpload(key string) (*v4.PresignedHTTPRequest, error) {
	req, err := presignedClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("rookie-test-bucket"),
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	log.Println(req.URL, req.Method, req.SignedHeader)
	return req, nil
}
