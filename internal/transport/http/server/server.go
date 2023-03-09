package server

import (
	"context"
	"log"
	"net/http"
)

type HTTPServer struct {
	Server http.Server
}

func NewServer(handler http.Handler, conf Conf) HTTPServer {
	return HTTPServer{
		Server: http.Server{
			Addr:         conf.Addr,
			Handler:      handler,
			ReadTimeout:  conf.ReadTimeout,
			WriteTimeout: conf.WriteTimeout,
			IdleTimeout:  conf.IdleTimeout,
		},
	}
}

func (s *HTTPServer) Start(ctx context.Context, halt <-chan struct{}) error {
	errShutdown := make(chan error, 1)
	go func() {
		defer close(errShutdown)
		select {
		case <-halt:
		case <-ctx.Done():
		}

		if err := s.Server.Shutdown(ctx); err != nil {
			errShutdown <- err
		}
	}()

	log.Println("Server is started")
	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		log.Println(err)
		return err
	}
	log.Println("Server is stopped")
	if err, ok := <-errShutdown; ok {
		return err
	}

	return nil
}
