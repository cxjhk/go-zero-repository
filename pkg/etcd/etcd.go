package etcd

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/discov"
)

type Etcd[T any] struct {
	c            Config
	configurator configurator.Configurator[T]
}

func NewEtcd[T any](c discov.EtcdConf) *Etcd[T] {
	return &Etcd[T]{
		configurator: configurator.MustNewConfigCenter[T](configurator.Config{
			Type: "json",
		}, subscriber.MustNewEtcdSubscriber(c)),
	}
}

func (ctr *Etcd[T]) GetConfig() (T, error) {
	return ctr.configurator.GetConfig()
}

func (ctr *Etcd[T]) Listener(listener func()) {
	ctr.configurator.AddListener(listener)
}
