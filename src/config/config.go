package config

import (
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

// New returns new Config
func New() (*Config, error) {
	k := koanf.New(".")

	// load yml config
	if err := k.Load(file.Provider("./config.yml"), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("loading yaml config: %w", err)
	}

	// load env variables and merge with yml config
	err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, envPrefix)), "_", ".", -1)
	}), nil)
	if err != nil {
		return nil, fmt.Errorf("loading env config: %w", err)
	}

	var conf Config
	if err := k.Unmarshal("", &conf); err != nil {
		return nil, fmt.Errorf("unmarshaling config: %w", err)
	}

	return &conf, nil
}
