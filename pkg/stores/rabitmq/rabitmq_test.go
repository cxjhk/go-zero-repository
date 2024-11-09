package rabitmq

import (
	"context"
	"github.com/google/uuid"
	"github.com/panjf2000/ants"
	"github.com/wagslane/go-rabbitmq"
	"go-zero-repository/pkg/stores/redisc"
	"log"
	"sync"
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

	var l = 50000
	var wg sync.WaitGroup
	wg.Add(l)
	var pool, _ = ants.NewPool(1000)
	for i := 0; i < l; i++ {
		pool.Submit(func() {
			defer wg.Done()
			err := publisher.PublishWithContext(
				context.Background(),
				[]byte(uuid.New().String()),
				[]string{"my_routing_key"},
				rabbitmq.WithPublishOptionsContentType("application/json"),
				rabbitmq.WithPublishOptionsMandatory,
				rabbitmq.WithPublishOptionsPersistentDelivery,
				rabbitmq.WithPublishOptionsExchange("events"),
			)
			if err != nil {
				log.Println(err)
			}
		})
	}
	wg.Wait()
}

func TestConsumer(t *testing.T) {
	consumer, err := rabbitmq.NewConsumer(
		NewConn("amqp://guest:guest@localhost"),
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
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
