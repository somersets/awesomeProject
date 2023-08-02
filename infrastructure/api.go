package infrastructure

import (
	"awesomeProject/infrastructure/router"
	"awesomeProject/registry/rootRegistry"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
}

func NewApi(config *Config) *API {
	return &API{config: config, logger: logrus.New()}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	postgre, err := api.newPostgre()
	if err != nil {
		return err
	}

	api.logger.Info("starting api server at port:", api.config.ServerPort)

	r := rootRegistry.NewRegistry(postgre)
	appController := r.NewAppController()

	err = api.newServer(router.InitRoutes(&appController))

	if err != nil {
		return err
	}
	return nil
}
