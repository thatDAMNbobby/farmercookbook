package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func NewTCPServerWithRouter(ctx context.Context, port int, router http.Handler) *http.Server {
	server := NewServerWithRouter(ctx, router)
	server.Addr = fmt.Sprintf(":%d", port)
	return server
}

func NewServerWithRouter(ctx context.Context, router http.Handler) *http.Server {
	return &http.Server{
		// Good practice to set timeouts to avoid Slowloris attacks.
		// https://en.wikipedia.org/wiki/Slowloris_(computer_security)
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
	}
}
