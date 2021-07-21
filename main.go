package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yh-server/internal/app/injector"
	"yh-server/internal/config"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := config.Load("config.toml"); err != nil {
		log.Error(errors.Wrap(err, "fail to load config file"))
		os.Exit(1)
	}
}

// @title c2server
// @version 1.0
// @description c2server
// @host localhost:8091
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Cookie
// @BasePath /
func main() {

	Injector, cleanFunc, err := injector.BuildInjector()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer cleanFunc()

	server := http.Server{
		Addr:    config.GetHttpAddress(),
		Handler: Injector.Engine,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("listening [%s] ...", config.GetHttpAddress())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Info("shutdown web server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error(errors.Wrap(err, "fail to shutdown web server"))
		os.Exit(1)
	}

	log.Info("web server exit!!!")

}
