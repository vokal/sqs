package sqs

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"strconv"
	"testing"
)

func TestReadFromSQS(t *testing.T) {
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	sqs := New(auth, aws.USWest2)
	queue, err := sqs.GetQueue("tracker-alpha-matt")
	if err != nil {
		panic(err)
	}

	results, err := queue.ReceiveMessage([]string{"All"}, 10, 600)
	count := len(results.Messages)
	deleteMessageBatch := make([]DeleteMessageBatch, count)
	for index, message := range results.Messages {
		fmt.Println(message.Body)
		deleteMessageBatch[index] = DeleteMessageBatch{
			Id:            strconv.Itoa(index),
			ReceiptHandle: message.ReceiptHandle,
		}
	}

	fmt.Println(fmt.Sprintf("%#v", deleteMessageBatch))
	response, err := queue.DeleteMessageBatch(deleteMessageBatch)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", response))
}
