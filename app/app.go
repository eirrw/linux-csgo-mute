package app

import (
	"github.com/dank/go-csgsi"
	"sync"
	"virunus.com/linux-csgo-mute/config"
	"virunus.com/linux-csgo-mute/pipewire"
)

type Context struct {
	Game   *csgsi.Game
	Mu     *sync.Mutex
	Pw     *pipewire.Pipewire
	Config *config.Config
	State  *State
}

type State struct {
	Alive     bool
	Bomb      string
	Connected bool
	Flashed   int
	Health    int
	NodeId    int
	Phase     string
}
