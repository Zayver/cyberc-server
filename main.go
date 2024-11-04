package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/zayver/cyberc-server/config"
	"github.com/zayver/cyberc-server/controller"
	"github.com/zayver/cyberc-server/middleware"
	"github.com/zayver/cyberc-server/repository"
	"github.com/zayver/cyberc-server/router"
	"github.com/zayver/cyberc-server/service"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controller.Module,
	router.Module,
	config.Module,
	service.Module,
	middleware.Module,
	repository.Module,
	fx.Invoke(Run),
)

func Run(router router.Routes){
	router.Init().Run()
}

func main() {
	app := fx.New(CommonModules)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal("Error bootstraping app: ", err)
	}
}