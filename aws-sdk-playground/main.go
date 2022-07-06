package main

import (
	"aws-sdk-playground/firehose"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
)

var (
	client *s3.Client
)

func init() {
	log.Println("load config")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("zigzag-alpha"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create s3 client")
	client = s3.NewFromConfig(cfg)
}
func main() {
	firehose.PutRecord()
}

func getBucketLifeCycleRule() {
	bucketName := "ad-display-alpha-serverlessdeploymentbucket-1t49h5mfvblq0"
	input := &s3.GetBucketLifecycleConfigurationInput{
		Bucket: &bucketName,
	}
	configuration, err := client.GetBucketLifecycleConfiguration(context.TODO(), input)
	if err != nil {
		panic(err)
	}
	fmt.Println(configuration)
}

func executeChangeBucketLifeCycle(listBucketOutput *s3.ListBucketsOutput) {
	var failureList []string
	for _, bucket := range listBucketOutput.Buckets {
		err := putBucketLifecycleRule(bucket.Name)
		if err != nil {
			log.Println(err)
			failureList = append(failureList, *bucket.Name)
		}
	}
	for _, item := range failureList {
		log.Println("버킷의 라이프사이클 적용에 실패했습니다. bucket: ", item)
	}
}

func putBucketLifecycleRule(bucketName *string) error {
	itRuleId := "intelligent-tiering-lifecycle-rule"
	dpRuleId := "delete-mpu-lifecycle-rule"
	var nilValue string
	memberPrefix := types.LifecycleRuleFilterMemberPrefix{Value: nilValue}

	configurationInput := s3.PutBucketLifecycleConfigurationInput{
		Bucket: bucketName,
		LifecycleConfiguration: &types.BucketLifecycleConfiguration{
			Rules: []types.LifecycleRule{
				{
					Status: types.ExpirationStatusEnabled,
					Transitions: []types.Transition{
						{
							Days:         1,
							StorageClass: types.TransitionStorageClassIntelligentTiering,
						},
					},
					ID: &itRuleId,
					//Prefix: &nilValue,
					Filter: &memberPrefix,
				},
				{
					Status: types.ExpirationStatusEnabled,
					AbortIncompleteMultipartUpload: &types.AbortIncompleteMultipartUpload{
						DaysAfterInitiation: int32(1),
					},
					ID: &dpRuleId,
					//Prefix: &nilValue,
					Filter: &memberPrefix,
				},
			},
		},
	}
	_, err := client.PutBucketLifecycleConfiguration(context.TODO(), &configurationInput)
	if err != nil {
		return err
	}
	log.Println("버킷의 라이프사이클 적용에 성공했습니다. bucket: ", *bucketName)
	return nil
}

func createPresignedUrl() {
	fmt.Println("Create Presign client")
	presignClient := s3.NewPresignClient(client)
	key := "test135"

	req, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("rookie-test-bucket"),
		Key:    &key,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(req.URL, req.Method, req.SignedHeader)
}

func getBucketList() (*s3.ListBucketsOutput, error) {
	buckets, err := client.ListBuckets(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return buckets, nil
}
