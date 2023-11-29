package main

import (
	"goremote/internal/app"
	"goremote/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.New()
	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath)
	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("application stopping", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()
	log.Info("application stopped")

	// TODO: инициализровать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return log
}
