package middleware

import (
	"fmt"
	"net/http"
)

type ErrMiddleware struct {
}

func NewErrMiddleware() *ErrMiddleware {
	return &ErrMiddleware{}
}

func (m *ErrMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		next(w, r)
	}
}
