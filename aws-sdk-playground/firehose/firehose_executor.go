package firehose

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"log"
)

var client *kinesis.Client

func init() {
	log.Println("load config")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("zigzag-alpha"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create s3 client")
	client = kinesis.NewFromConfig(cfg)
}

func PutRecord() {
	data := "test"
	input := &kinesis.PutRecordInput{
		Data:         []byte(data),
		PartitionKey: aws.String("rookie-partitionKey"),
		StreamName:   aws.String("rookie-os-stream"),
	}
	record, err := client.PutRecord(context.TODO(), input)
	if err != nil {
		panic(err)
	}
	fmt.Println(record)
}

func GetRecord() {
	records, err := client.GetRecords(context.TODO(), &kinesis.GetRecordsInput{
		ShardIterator: aws.String("test"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(records)
}
