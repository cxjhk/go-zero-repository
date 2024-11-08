package etcd

import (
	"github.com/jinzhu/copier"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
)

type Etcd[T any] struct {
	c            Config
	configurator configurator.Configurator[T]
}

func NewEtcd[T any](c Config) *Etcd[T] {
	var cc subscriber.EtcdConf
	_ = copier.Copy(&cc, &c)
	return &Etcd[T]{
		c: c,
		configurator: configurator.MustNewConfigCenter[T](configurator.Config{
			Type: "json",
		}, subscriber.MustNewEtcdSubscriber(cc)),
	}
}

func (ctr *Etcd[T]) GetConfig() (T, error) {
	return ctr.configurator.GetConfig()
}

func (ctr *Etcd[T]) Listener(listener func()) {
	ctr.configurator.AddListener(listener)
}
