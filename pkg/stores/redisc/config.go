package redisc

type Config struct {
	Addr     string
	Username string `json:",optional"`
	Password string `json:",optional"`
	DB       int    `json:",default=0"`
}
