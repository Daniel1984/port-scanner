package main

import (
	"github.com/port-scanner/cmd/server/cfg"
	"github.com/port-scanner/cmd/server/router"
	"github.com/port-scanner/pkg/gracefulshutdown"
	"github.com/port-scanner/pkg/server"
)

func main() {
	cfg := cfg.New()

	// configure server instance
	srv := server.
		New().
		WithAddr(cfg.GetAPIPort()).
		WithRouter(router.Get()).
		WithErrLogger(cfg.Errlog)

	// start server in separate goroutine so that it doesn't block graceful
	// shutdown handler
	go func() {
		cfg.Infolog.Printf("starting server at %s", cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			cfg.Errlog.Printf("starting server: %s", err)
		}
	}()

	// initiate graceful shutdown handler that will listen to api crash signals
	// and perform cleanup
	gracefulshutdown.Init(cfg.Errlog, func() {
		if err := srv.Close(); err != nil {
			cfg.Errlog.Printf("closing server: %s", err)
		}
	})
}
