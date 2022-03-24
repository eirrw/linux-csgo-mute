package config

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

const (
	ConfigFile      = "csgo-mute/config.toml"
	DefaultToken    = "VALVEPLSSPAREMYEARS"
	DefaultCsgoNode = "csgo_linux64"

	AppKey    = "app"
	GsiKey    = "gsi"
	VolumeKey = "volume"

	CsgoNodeKey = "csgoNodeName"
	PortKey     = "port"
	TokenKey    = "token"
	FlashEndKey = "flashEnd"
	FlashKey    = "flash"
	DeathKey    = "death"
	BombKey     = "bomb"
	DefaultKey  = "default"
)

type Config struct {
	App    *app               `toml:"app"`
	Gsi    *gsi               `toml:"gsi"`
	Volume map[string]float32 `toml:"volume"`
}

type app struct {
	CsgoNodeName string `toml:"csgoNodeName"`
}

type gsi struct {
	Port     int    `toml:"port"`
	Token    string `toml:"token"`
	FlashEnd int    `toml:"flashEnd"`
}

func New() *Config {
	// default Config
	c := &Config{
		App: &app{
			CsgoNodeName: DefaultCsgoNode,
		},
		Gsi: &gsi{
			Port:     3202,
			Token:    DefaultToken,
			FlashEnd: 200,
		},
		Volume: map[string]float32{
			FlashKey:   0.2,
			DeathKey:   0.2,
			BombKey:    0.2,
			DefaultKey: 1.0,
		},
	}

	if err := loadConfigFile(c); err != nil {
		log.Fatal(err)
	}

	if err := validateConfig(c); err != nil {
		log.Fatal(err)
	}

	return c
}

func (c Config) GetConfig() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c Config) WriteFile() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, ConfigFile)

	if err = os.MkdirAll(filepath.Dir(configPath), 0744); err != nil {
		return err
	}

	var str []byte
	if str, err = c.GetConfig(); err != nil {
		return err
	}

	if err = os.WriteFile(configPath, str, 0644); err != nil {
		return err
	}

	return nil
}

func loadConfigFile(defaultConfig *Config) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, ConfigFile)

	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}

	newConfig := Config{}
	var md toml.MetaData
	if md, err = toml.DecodeFile(configPath, &newConfig); err != nil {
		return err
	}

	// overwrite defaults with configured values
	// app
	if md.IsDefined(AppKey, CsgoNodeKey) {
		defaultConfig.App.CsgoNodeName = newConfig.App.CsgoNodeName
	}

	// gsi
	if md.IsDefined(GsiKey, PortKey) {
		defaultConfig.Gsi.Port = newConfig.Gsi.Port
	}
	if md.IsDefined(GsiKey, TokenKey) {
		defaultConfig.Gsi.Token = newConfig.Gsi.Token
	}
	if md.IsDefined(GsiKey, FlashEndKey) {
		defaultConfig.Gsi.FlashEnd = newConfig.Gsi.FlashEnd
	}

	// volume
	for k, v := range newConfig.Volume {
		defaultConfig.Volume[k] = v
	}

	return nil
}

func validateConfig(config *Config) error {
	if config.Gsi.FlashEnd < 0 || config.Gsi.FlashEnd > 255 {
		return errors.New(fmt.Sprintf("invalid value for 'flashEnd': '%d'. must be in range 0 - 255", config.Gsi.FlashEnd))
	}

	for k, v := range config.Volume {
		if v < 0 || v > 1 {
			return errors.New(fmt.Sprintf("invalid value for '%s': '%.1f'. must be in range 0.0 - 1.0", k, v))
		}
	}

	return nil
}
