package router

import (
	"crawler-backend/internal/crawler/infrastructure/transport"
	"github.com/bearname/http-server/pkg/server"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/pprof"
)

func Router(controller *transport.Controller) http.Handler {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	api := router.PathPrefix("/api/v1/urls").Subrouter()

	router.HandleFunc("/health", server.HealthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/ready", server.ReadyCheckHandler).Methods(http.MethodGet)
	api.HandleFunc("", controller.Create).Methods(http.MethodGet)

	pprofRouter := router.PathPrefix("/debug/pprof").Subrouter()
	pprofRouter.HandleFunc("/", pprof.Index)
	pprofRouter.HandleFunc("/cmdline", pprof.Cmdline)
	pprofRouter.HandleFunc("/symbol", pprof.Symbol)
	pprofRouter.HandleFunc("/trace", pprof.Trace)

	profile := pprofRouter.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("", pprof.Profile)
	profile.Handle("/goroutine", pprof.Handler("goroutine"))
	profile.Handle("/threadcreate", pprof.Handler("threadcreate"))
	profile.Handle("/heap", pprof.Handler("heap"))
	profile.Handle("/block", pprof.Handler("block"))
	profile.Handle("/mutex", pprof.Handler("mutex"))

	return server.LogMiddleware(router)
}
