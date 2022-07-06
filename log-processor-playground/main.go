package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func HandleRequest(ctx context.Context, event events.KinesisFirehoseEvent) (events.KinesisFirehoseResponse, error) {
	var responseRecord []events.KinesisFirehoseResponseRecord
	for _, record := range event.Records {
		result := events.KinesisFirehoseTransformedStateOk
		test := "{\"key\": \"val\"}"
		log.Println(string(record.Data))
		responseRecord = append(responseRecord, events.KinesisFirehoseResponseRecord{
			RecordID: record.RecordID,
			Result:   result,
			Data:     []byte(test),
			Metadata: events.KinesisFirehoseResponseRecordMetadata{},
		})
	}
	return events.KinesisFirehoseResponse{Records: responseRecord}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
