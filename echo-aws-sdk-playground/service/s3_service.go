package service

import (
	"echo-aws-sdk-playground/gateway"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

func GetPreSignedUrlForUpload(key string) (*v4.PresignedHTTPRequest, error) {
	return gateway.GetPreSignedUrlForUpload(key)
}
