package timec

import "time"

func SetTimeout(callback func(), delay time.Duration) *time.Timer {
	return time.AfterFunc(delay, callback)
}

type IntervalHandler struct {
	ticker *time.Ticker
	stop   chan bool
}

func SetInterval(callback func(), delay time.Duration) *IntervalHandler {
	ticker := time.NewTicker(delay)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				callback()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return &IntervalHandler{
		ticker,
		stop,
	}
}

func (h *IntervalHandler) Stop() {
	h.stop <- true
}
