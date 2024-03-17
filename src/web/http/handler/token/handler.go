package token

import (
	"context"
	"net/http"
)

type Service interface {
	Authorize(ctx context.Context, userID int, clientID, clientSecret, code, redirectURL string) (string, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Authorize handles user authorization
func (h *Handler) Authorize(w http.ResponseWriter, r *http.Request) {

	userID := 1
	clientID := "client_id"
	clientSecret := "client_secret"
	code := r.URL.Query().Get("code")
	redirectURL := r.URL.Query().Get("redirect_uri")

	token, err := h.service.Authorize(r.Context(), userID, clientID, clientSecret, code, redirectURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte(token))
}
