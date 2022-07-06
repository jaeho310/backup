package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.KinesisFirehoseEvent) (events.KinesisFirehoseResponse, error) {
	var responseRecord []events.KinesisFirehoseResponseRecord
	for _, record := range event.Records {
		result := events.KinesisFirehoseTransformedStateDropped
		logMessages := "test"
		// logMessages, err := service.LogProcessor{}.ExecuteLogProcessor(record.Data)
		// if err != nil || len(logMessages) == 0 {
		// log.Println(err)
		// result = events.KinesisFirehoseTransformedStateProcessingFailed
		// }
		responseRecord = append(responseRecord, events.KinesisFirehoseResponseRecord{
			RecordID: record.RecordID,
			Result:   result,
			Data:     []byte(logMessages),
			Metadata: events.KinesisFirehoseResponseRecordMetadata{},
		})
	}
	return events.KinesisFirehoseResponse{Records: responseRecord}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
