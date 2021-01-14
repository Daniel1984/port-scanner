package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/port-scanner/cmd/server/handlers"
)

func Get() *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/open-ports", handlers.ScanOpenPorts)
	return mux
}
