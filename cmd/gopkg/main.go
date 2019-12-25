package main

import (
	"github.com/emicklei/go-restful"

	"pgxs.io/chassis"
	"pgxs.io/chassis/config"
	"pgxs.io/chassis/log"

	"pgxs.io/gopkg/pkg/apiserver"
	"pgxs.io/gopkg/pkg/migrations"
)

func main() {
	var log = log.New()
	lApiserver := log.WithField("component", "gopkg")
	lApiserver.Infof("API server starting on http://127.0.0.1:%d", config.Server().Port)
	lApiserver.Infof("API docs on http://127.0.0.1:%d%s", config.Server().Port, config.Openapi().UI.Entrypoint)

	if err := migrations.Migrate(); err != nil {
		lApiserver.Fatalln(err)
	}

	apiserver.RegistPackageWebservice()

	chassis.Start(restful.RegisteredWebServices())

}
