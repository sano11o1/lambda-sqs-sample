package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.SQSEvent) (string, error) {
	log.Print("๐๐๐๐ messageใฎๆฐ", len(event.Records))
	for _, message := range event.Records {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		number := *message.MessageAttributes["number"].StringValue
		if number == "A-4" {
			time.Sleep(time.Second * 10)
		}
		log.Print("๐๐๐๐", number, "โฑ๏ธ", time.Now().In(loc))
	}
	return "success", nil
}

func main() {
	lambda.Start(HandleRequest)
}
