package rabitmq

import (
	"crypto/md5"
	"errors"
	"github.com/wagslane/go-rabbitmq"
	"github.com/zeromicro/go-zero/core/syncx"
)

/// rabbitmq publisher封装

type Publisher struct {
	c    PublisherConfig
	pool *syncx.Pool
}

func NewPublisher(conn *rabbitmq.Conn, c PublisherConfig, optionFuncs ...func(*rabbitmq.PublisherOptions)) *Publisher {
	return &Publisher{
		c: c,
		pool: syncx.NewPool(c.PoolSize, func() any {
			publisher, err2 := rabbitmq.NewPublisher(
				conn,
				optionFuncs...,
			)
			if err2 != nil {
				panic(err2)
			}
			return publisher
		}, func(pool any) {
			if item, ok := pool.(*rabbitmq.Publisher); ok {
				item.Close()
			}
		}),
	}
}

func (p *Publisher) Publish(
	data []byte,
	routingKeys []string,
	optionFuncs ...func(*rabbitmq.PublishOptions),
) error {
	pool := p.pool.Get()
	if publisher, ok := pool.(*rabbitmq.Publisher); ok {
		defer p.pool.Put(pool)
		return publisher.Publish(data, routingKeys, append(optionFuncs, func(options *rabbitmq.PublishOptions) {
			options.MessageID = p.msgId(data)
		})...)
	}
	return errors.New("publisher is nil")
}

func (p *Publisher) msgId(data []byte) string {
	md5hash := md5.New()
	md5hash.Write(data)
	return string(md5hash.Sum(nil))
}
