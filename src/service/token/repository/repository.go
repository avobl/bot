package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/avobl/bot/src/db"
)

var (
	// ErrNotFound is returned when token is not found in db
	ErrNotFound = errors.New("token not found")
)

// TokenRepository is a repository for user's oauth2 token for Todoist
type TokenRepository struct {
	db db.Provider
}

// Token represents user's oauth2 token for Todoist
type Token struct {
	UserID      int    `db:"user_id"`
	Code        string `db:"code"`
	AccessToken string `db:"access_token"`
}

// NewTokenRepository returns new TokenRepository
func NewTokenRepository(db db.Provider) *TokenRepository {
	return &TokenRepository{db: db}
}

// GetByUser retrieves user token from db
func (r *TokenRepository) GetByUser(ctx context.Context, userID int) (Token, error) {
	query := `SELECT user_id, code, access_token FROM user_tokens WHERE user_id = ?`

	var token Token

	row := r.db.QueryRowContext(ctx, query, userID)
	err := row.Scan(&token.UserID, &token.Code, &token.AccessToken)
	if err != nil {
		if errors.Is(err, db.ErrNoRows) {
			return Token{}, ErrNotFound
		}

		return Token{}, fmt.Errorf("scan row: %w", err)
	}

	return token, nil
}

// Save saves user token to db
func (r *TokenRepository) Save(ctx context.Context, userID int, code, token string) error {
	query := `INSERT INTO user_tokens (user_id, code, access_token) VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, userID, code, token)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
