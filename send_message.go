package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	queueUrl := "https://sqs.ap-northeast-1.amazonaws.com/087391343737/MyFifo.fifo"
	for i := 1; i <= 5; i++ {
		groupID := "A"
		messageDeduplicationId := groupID + "-" + strconv.Itoa(i)
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			MessageGroupId:         &groupID,
			MessageDeduplicationId: &messageDeduplicationId,
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"number": {
					DataType:    aws.String("String"),
					StringValue: aws.String(groupID + "-" + strconv.Itoa(i)),
				},
			},
			MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
			QueueUrl:    &queueUrl,
		})
		fmt.Println(err)
	}
	for i := 1; i <= 5; i++ {
		groupID := "B"
		messageDeduplicationId := groupID + "-" + strconv.Itoa(i)
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			MessageGroupId:         &groupID,
			MessageDeduplicationId: &messageDeduplicationId,
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"number": {
					DataType:    aws.String("String"),
					StringValue: aws.String(groupID + "-" + strconv.Itoa(i)),
				},
			},
			MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
			QueueUrl:    &queueUrl,
		})
		fmt.Println(err)
	}
	for i := 6; i <= 10; i++ {
		groupID := "A"
		messageDeduplicationId := groupID + "-" + strconv.Itoa(i)
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			MessageGroupId:         &groupID,
			MessageDeduplicationId: &messageDeduplicationId,
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"number": {
					DataType:    aws.String("String"),
					StringValue: aws.String(groupID + "-" + strconv.Itoa(i)),
				},
			},
			MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
			QueueUrl:    &queueUrl,
		})
		fmt.Println(err)
	}
}

// func buildBatchMessages(groupID string, length int) []*sqs.SendMessageBatchRequestEntry {
// 	var entries []*sqs.SendMessageBatchRequestEntry
// 	for i := 0; i < length; i++ {
// 		messageDeduplicationId := groupID + "-" + strconv.Itoa(i)
// 		entry := &sqs.SendMessageBatchRequestEntry{
// 			Id:                     &messageDeduplicationId,
// 			MessageGroupId:         &groupID,
// 			MessageDeduplicationId: &messageDeduplicationId,
// 			MessageAttributes: map[string]*sqs.MessageAttributeValue{
// 				"number": {
// 					DataType:    aws.String("String"),
// 					StringValue: aws.String(groupID + "-" + strconv.Itoa(i)),
// 				},
// 			},
// 			MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
// 		}
// 		entries = append(entries, entry)
// 	}
// 	return entries
// }
