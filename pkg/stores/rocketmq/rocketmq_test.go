package rocketmq

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"strconv"
	"testing"
	"time"
)

const (
	Topic     = "testTopic"
	Endpoint  = "127.0.0.1:9876"
	AccessKey = "xxxxxx"
	SecretKey = "xxxxxx"
)

func TestPublisher(t *testing.T) {
	// In most case, you don't need to create many producers, singleton pattern is more recommended.
	producer, err := rmq_client.NewProducer(
		&rmq_client.Config{
			Endpoint: Endpoint,
		},
		rmq_client.WithTopics(Topic),
	)
	if err != nil {
		panic(err)
	}
	// start producer
	err = producer.Start()
	if err != nil {
		panic(err)
	}
	// graceful stop producer
	defer producer.GracefulStop()

	for i := 0; i < 10; i++ {
		// new a message
		msg := &rmq_client.Message{
			Topic: Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")
		// send message in sync
		resp, err := producer.Send(context.TODO(), msg)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
