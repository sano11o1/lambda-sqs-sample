package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.SQSEvent) error {
	log.Print("😈😈😈😈 messageの数", len(event.Records))
	for _, e := range event.Records {
		fmt.Println("イベント😇", e.MessageId)
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
