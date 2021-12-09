package main

import (
	"Short"
	"Short/internal/config"
	"Short/internal/server"
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func initEnvs() {
	// Set root env
	if err := Short.SetRootEnvs("ROOT_PATH"); err != nil {
		logrus.Info(err.Error())
		if err = Short.SetDefaultEnvs(); err != nil {
			logrus.Fatalln(err.Error())
		}
	}
}

func main() {
	var configPath string

	initEnvs()
	flag.StringVar(&configPath, "config-path", os.Getenv("ROOT_PATH")+"/"+"configs/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Info(err.Error())

		cfg = config.Default()
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	s := server.New(ctx, cfg)
	s.Run(ctx, stop)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				logrus.Fatal("graceful shutdown timed out.. forcing exit")
			}
		}()

		err := s.Shutdown(shutdownCtx)
		if err != nil {
			logrus.Fatal(err)
		}

		stop()
		cancel()
	}()

	<-ctx.Done()
}
