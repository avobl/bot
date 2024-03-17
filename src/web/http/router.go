package http

import (
	"context"
	"net/http"

	"github.com/avobl/bot/src/config"
	"github.com/avobl/bot/src/db"
	"github.com/avobl/bot/src/external/todoist"
	"github.com/avobl/bot/src/service/token"
	"github.com/avobl/bot/src/service/token/repository"
	tknhandler "github.com/avobl/bot/src/web/http/handler/token"
)

func LoadRoutes(_ context.Context) *http.ServeMux {
	// http clients
	client := &http.Client{}
	todoistClient := todoist.NewClient(client, config.C.Todoist.DevBaseURL, config.C.Todoist.AuthzBaseURL)

	// repositories
	tokenRepo := repository.NewTokenRepository(db.Connection)

	// services
	tokenService := token.NewService(todoistClient, tokenRepo)

	// handlers
	tokenHandler := tknhandler.NewHandler(tokenService)

	// routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /bot/v1/token", tokenHandler.Authorize)

	return mux
}
