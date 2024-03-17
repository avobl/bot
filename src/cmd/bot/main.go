package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/avobl/bot/src/app"
	"github.com/avobl/bot/src/web/http/server"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-signals
		slog.WarnContext(ctx, "received signal", slog.Any("signal", sig))
		cancel()
	}()

	if err := app.Load(ctx); err != nil {
		slog.Error("loading app", slog.String("error", err.Error()))
		return
	}

	server.Start(ctx, cancel)
}
