package app

import (
	"context"
	"fmt"

	"github.com/avobl/bot/src/config"
	"github.com/avobl/bot/src/db"
	"github.com/avobl/bot/src/logger"
)

func Load(ctx context.Context) error {
	logger.Load(ctx)

	if err := config.Load(ctx); err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	if err := db.Load(ctx); err != nil {
		return fmt.Errorf("loading db: %w", err)
	}

	return nil
}
