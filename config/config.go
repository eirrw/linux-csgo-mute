package config

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

const ConfigFile = "csgo-mute/config.toml"
const DefaultToken = "VALVEPLSSPAREMYEARS"

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
	c := Config{
		Gsi: gsi{
			Port:  3202,
			Token: DefaultToken,
		},
		Volume: volume{
			Flash:   0.2,
			Death:   0.2,
			Bomb:    0.2,
			Default: 1.0,
		},
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
