package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/port-scanner/cmd/server/cfg"
	"github.com/port-scanner/cmd/server/handlers"
)

func Get(c *cfg.Cfg) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/open-ports", handlers.ScanOpenPorts(c))
	return mux
}
