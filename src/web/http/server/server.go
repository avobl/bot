package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/avobl/bot/src/config"
	router "github.com/avobl/bot/src/web/http"
)

func Start(ctx context.Context, cancel context.CancelFunc) {
	mux := router.LoadRoutes(ctx)

	server := http.Server{
		Addr:              fmt.Sprintf(":%s", config.C.Server.Port),
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.ErrorContext(ctx, "listen and serve", slog.Any("error", err))
			cancel()
		}
	}()

	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "shutdown: %v", slog.Any("error", err))
	}
}
