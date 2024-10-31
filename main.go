package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/zayver/cybercomplaint-server/config"
	"github.com/zayver/cybercomplaint-server/controller"
	"github.com/zayver/cybercomplaint-server/middleware"
	"github.com/zayver/cybercomplaint-server/repository"
	"github.com/zayver/cybercomplaint-server/router"
	"github.com/zayver/cybercomplaint-server/service"
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