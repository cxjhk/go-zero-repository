package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Etcd       discov.EtcdConf
	TestConfig TestConfig `json:",optional"`
}

type TestConfig struct {
	Name string `json:"name"`
}
