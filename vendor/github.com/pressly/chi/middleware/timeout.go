package middleware

import (
	"net/http"
	"time"

	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

// Timeout is a middleware that cancels ctx after a given timeout and return
// a 504 Gateway Timeout error to the client.
//
// It's required that you select the ctx.Done() channel to check for the signal
// if the context has reached its deadline and return, otherwise the timeout
// signal will be just ignored.
//
// ie. a route/handler may look like:
//
// r.Get("/long", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
// 	processTime := time.Duration(rand.Intn(4)+1) * time.Second
//
// 	select {
// 	case <-ctx.Done():
// 		return
//
// 	case <-time.After(processTime):
// 		// The above channel simulates some hard work.
// 	}
//
// 	w.Write([]byte("done"))
// })
//
func Timeout(timeout time.Duration) func(next chi.Handler) chi.Handler {
	return func(next chi.Handler) chi.Handler {
		fn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer func() {
				cancel()
				if ctx.Err() == context.DeadlineExceeded {
					w.WriteHeader(http.StatusGatewayTimeout)
				}
			}()

			next.ServeHTTPC(ctx, w, r)
		}
		return chi.HandlerFunc(fn)
	}
}
