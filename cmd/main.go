package main

import (
	"chatbotbasic/cmd/config"
	"chatbotbasic/internal/platform/webapi/router"
	"chatbotbasic/pkg/customlog"
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	if err := Run(); err != nil {
		panic(err.Error())
	}
}

func Run() error {
	conf, err := config.GetConfig()
	if err != nil {
		log.Panic().Msg(err.Error())
		return err
	}

	// setup logger
	log.Logger = customlog.Init(conf.LogInfo.Path, conf.LogInfo.Size, conf.LogInfo.Backup, conf.LogError.Path, conf.LogError.Size, conf.LogError.Backup)

	// MongoDB
	db := config.NewMongoDatabase(conf.DB)

	// setup server
	s := &http.Server{
		Addr:         conf.AppConfig.Address,
		ReadTimeout:  conf.AppConfig.ReadTimeout,
		WriteTimeout: conf.AppConfig.WriteTimeout,
	}

	// setup modules
	modules := Bootstrap(db)

	// router & init run
	e, err := router.RunRouter(*modules)
	if err != nil {
		log.Error().Str("main", "running router").Msg(err.Error())
		return err
	}

	// run server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("shutting down server...")
		}
		//log.Info().Str("server", "server running").Msg(fmt.Sprintf("running in %s", s.Addr))
	}()

	// handler error run server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), conf.AppConfig.ShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Error().Str("main", "timeout runner").Msg(err.Error())
		return err
	}

	return nil
}
