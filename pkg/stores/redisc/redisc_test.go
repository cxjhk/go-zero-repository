package redisc

import (
	"context"
	"sync"
	"testing"
)

func TestNewEngine(t *testing.T) {
	engine := NewEngine(Config{
		Addrs: []string{"127.0.0.1:6379"},
	})

	var (
		l  = 500000
		wg sync.WaitGroup
	)
	wg.Add(l)
	for i := 0; i < l; i++ {
		go func() {
			defer wg.Done()
			engine.Incr(context.TODO(), "test")
		}()
	}
	wg.Wait()
}
