package config

import (
	"github.com/BurntSushi/toml"
	"github.com/iain17/ipLookup/src/utils/logger"
)

func Load(path string) *Config {
	cfg := new(Config)

	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		logger.Fatalf("Cannot read config file; %s", err)
	}

	return cfg
}
