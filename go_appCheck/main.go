package main

import (
	"appcheck/internal/app/injector"
	"appcheck/internal/config"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := config.Load("config.toml"); err != nil {
		log.Error(errors.Wrap(err, "fail to load config file"))
		os.Exit(1)
	}
}

// @title appcheck
// @version 1.0
// @description appcheck
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
