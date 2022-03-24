package config

import (
	"bytes"
	_ "embed"
	"errors"
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
	App    *app    `toml:"app"`
	Gsi    *gsi    `toml:"gsi"`
	Volume *volume `toml:"volume"`
}

type app struct {
	CsgoNodeName string `toml:"csgoNodeName"`
}

type gsi struct {
	Port     int    `toml:"port"`
	Token    string `toml:"token"`
	FlashEnd int    `toml:"flashEnd"`
}

type volume struct {
	Flash   float32 `toml:"flash"`
	Death   float32 `toml:"death"`
	Bomb    float32 `toml:"bomb"`
	Default float32 `toml:"default"`
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
		Volume: &volume{
			Flash:   0.2,
			Death:   0.2,
			Bomb:    0.2,
			Default: 1.0,
		},
	}

	if err := loadConfigFile(c); err != nil {
		log.Fatal(err)
	}

	return c
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

	buf := new(bytes.Buffer)
	if err = toml.NewEncoder(buf).Encode(c); err != nil {
		return err
	}

	if err = os.WriteFile(configPath, buf.Bytes(), 0644); err != nil {
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
	if md.IsDefined(AppKey, CsgoNodeKey) {
		defaultConfig.App.CsgoNodeName = newConfig.App.CsgoNodeName
	}

	if md.IsDefined(GsiKey, PortKey) {
		defaultConfig.Gsi.Port = newConfig.Gsi.Port
	}
	if md.IsDefined(GsiKey, TokenKey) {
		defaultConfig.Gsi.Token = newConfig.Gsi.Token
	}
	if md.IsDefined(GsiKey, FlashEndKey) {
		defaultConfig.Gsi.FlashEnd = newConfig.Gsi.FlashEnd
	}

	if md.IsDefined(VolumeKey, FlashKey) {
		defaultConfig.Volume.Flash = newConfig.Volume.Flash
	}
	if md.IsDefined(VolumeKey, DeathKey) {
		defaultConfig.Volume.Death = newConfig.Volume.Death
	}
	if md.IsDefined(VolumeKey, BombKey) {
		defaultConfig.Volume.Bomb = newConfig.Volume.Bomb
	}
	if md.IsDefined(VolumeKey, DefaultKey) {
		defaultConfig.Volume.Default = newConfig.Volume.Default
	}

	return nil
}
