package server

import (
	"log"
	"net/http"

	"github.com/XeniaBgd/CleanArch/internal/transport/http/server/params"
)

func StartServer(handler http.Handler, conf params.Conf) {
	server := http.Server{
		Addr:         conf.Addr,
		Handler:      handler,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		IdleTimeout:  conf.IdleTimeout,
	}

	log.Println("Server is started")
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
