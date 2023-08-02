package infrastructure

import (
	"net/http"
)

func (api *API) newServer(handler http.Handler) error {
	httpServer := &http.Server{
		Addr:                         ":" + api.config.ServerPort,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	return httpServer.ListenAndServe()
}
