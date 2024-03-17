package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const (
	envPrefix = "BOT_"
)

var C *Configuration

// Load initializes the configuration from the file and environment variables.
func Load(_ context.Context) error {
	k := koanf.New(".")

	// load yml config
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		return fmt.Errorf("loading yaml config: %w", err)
	}

	// load env variables and merge with yml config
	err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, envPrefix)), "_", ".", -1)
	}), nil)
	if err != nil {
		return fmt.Errorf("loading env config: %w", err)
	}

	var conf Configuration
	if err = k.Unmarshal("", &conf); err != nil {
		return fmt.Errorf("unmarshaling config: %w", err)
	}

	C = &conf

	return nil
}
