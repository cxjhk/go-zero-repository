package rabitmq

import "github.com/wagslane/go-rabbitmq"

func NewConn(url string, opts ...func(*rabbitmq.ConnectionOptions)) *rabbitmq.Conn {
	conn, err := rabbitmq.NewConn(
		url,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	return conn
}
