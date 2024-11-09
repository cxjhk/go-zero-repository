package rabitmq

import (
	"context"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"go-zero-repository/pkg/stores/redisc"
	"log"
	"testing"
)

func TestPublish(t *testing.T) {
	publisher := NewPublisher(
		NewConn("amqp://guest:guest@localhost"),
		PublisherConfig{
			PoolSize: 1,
		},
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)

	var l = 10000
	//var wg sync.WaitGroup
	//wg.Add(l)
	for i := 0; i < l; i++ {
		//go func(i int) {
		//	defer wg.Done()
		err := publisher.Publish(
			[]byte(uuid.New().String()),
			[]string{"my_routing_key"},
			rabbitmq.WithPublishOptionsExchange("events"),
			rabbitmq.WithPublishOptionsPersistentDelivery,
		)
		if err != nil {
			log.Println(err)
		}
		//}(i)
	}
	//wg.Wait()
}

func TestConsumer(t *testing.T) {
	consumer, err := rabbitmq.NewConsumer(
		NewConn("amqp://guest:guest@localhost"),
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsConcurrency(100),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	var redisClient = redisc.NewEngine(redisc.Config{
		Addrs: []string{"localhost:6379"},
	})
	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		redisClient.Incr(context.Background(), "test")
		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal(err)
	}
}
