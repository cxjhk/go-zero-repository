package redisc

type Config struct {
	Addrs      []string
	Debug      bool   `json:",default=false"`
	MasterName string `json:",optional"`
	Username   string `json:",optional"`
	Password   string `json:",optional"`
	DB         int    `json:",default=0"`
}
