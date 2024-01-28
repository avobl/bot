package token

import (
	"context"
	"errors"
	"fmt"

	"github.com/avobl/bot/src/service/token/repository"
)

var (
	ErrNotFound = errors.New("token not found")
)

type TodoistClient interface {
	RetrieveAccessToken(ctx context.Context, clientID, clientSecret, code, redirectURL string) (string, error)
}

type Repository interface {
	GetByUser(ctx context.Context, userID int) (repository.Token, error)
	Save(ctx context.Context, userID int, code, token string) error
}

type Service struct {
	todoistClient TodoistClient
	tokenRepo     Repository
}

func NewService(todoistClient TodoistClient, tokenRepo Repository) *Service {
	return &Service{
		todoistClient: todoistClient,
		tokenRepo:     tokenRepo,
	}
}

// GetToken returns user's access token stored in db
func (s *Service) GetToken(ctx context.Context, userID int) (string, error) {
	storedToken, err := s.tokenRepo.GetByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", fmt.Errorf("get token: %v", err)
	}

	return storedToken.AccessToken, nil
}

// Authorize retrieves and saves user token to db
func (s *Service) Authorize(ctx context.Context, userID int, clientID, clientSecret, code, redirectURL string) (string, error) {
	token, err := s.todoistClient.RetrieveAccessToken(ctx, clientID, clientSecret, code, redirectURL)
	if err != nil {
		return "", fmt.Errorf("retrieve access token: %v", err)
	}

	err = s.tokenRepo.Save(ctx, userID, code, token)
	if err != nil {
		return "", fmt.Errorf("save token: %v", err)
	}

	return token, nil
}
