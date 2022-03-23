package config

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

const ConfigFile = "csgo-mute/Config.toml"

//go:embed config.toml
var defaultConfig []byte

type Config struct {
	Gsi    gsi
	Volume volume
}

type gsi struct {
	Port  int
	Token string
}

type volume struct {
	Flash   float32
	Death   float32
	Bomb    float32
	Default float32
}

func New() *Config {
	// default Config
	c := Config{}
	if _, err := toml.Decode(string(defaultConfig), &c); err != nil {
		log.Fatal(err)
	}

	return &c
}

func loadConfigFile() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(configDir, ConfigFile)

	c := Config{}
	if _, err := toml.DecodeFile(configPath, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
