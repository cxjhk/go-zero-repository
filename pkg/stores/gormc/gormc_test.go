package gormc

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	Nickname string `json:"nickname"`
}

func TestClickhouse(t *testing.T) {
	//docker run -d -p 19000:9000 -p 18123:8123 -p 19440:9440 --ulimit nofile=262144:262144 clickhouse/clickhouse-server
	db, err := gorm.Open(clickhouse.Open(
		"clickhouse://localhost:19000/default?dial_timeout=10s&read_timeout=20s",
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Nickname: "test"})
}
