package app

import (
	"fmt"
	"github.com/dank/go-csgsi"
	"log"
	"sync"
	"time"
	"virunus.com/linux-csgo-mute/config"
	"virunus.com/linux-csgo-mute/pipewire"
)

type Context struct {
	Game   *csgsi.Game
	Mu     *sync.Mutex
	Pw     *pipewire.Pipewire
	Config *config.Config
	State  *State
	Debug  bool
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

func Start(debug bool) error {
	context := Context{
		Game:   csgsi.New(0),
		Config: config.New(),
		State:  &State{},
		Debug:  debug,
	}

	go func() {
		for {
			select {
			case state := <-context.Game.Channel:
				if state.Auth == nil || state.Auth.Token != context.Config.Gsi.Token {
					fmt.Println("bad")
					continue
				}

				if !context.State.Connected {
					context.Pw = pipewire.New()
					context.State.NodeId = context.Pw.GetNodeIdByName(context.Config.App.CsgoNodeName)
					if context.State.NodeId == -1 {
						if context.Debug {
							log.Println("Could not find audio node, reloading")
						}
						continue
					} else {
						context.State.Connected = true
					}
				}

				err := handleState(state, &context)
				if err != nil {
					log.Println(err)
				}
			case <-time.After(time.Minute):
				if context.Debug {
					log.Println("timeout")
				}
				context.State.Connected = false
			}
		}
	}()

	if err := context.Game.Listen(fmt.Sprintf(":%d", context.Config.Gsi.Port)); err != nil {
		return err
	}

	return nil
}

func handleState(state csgsi.State, ctx *Context) error {
	var err error
	if state.Player != nil && state.Round != nil {
		if state.Player.State.Flashed != ctx.State.Flashed {
			ctx.State.Flashed = state.Player.State.Flashed
			if ctx.State.Flashed > 200 {
				if ctx.Debug {
					log.Printf("flashed: %d\n", ctx.State.Flashed)
				}
				err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.FlashKey])
			} else {
				if ctx.Debug {
					log.Println("not flashed")
				}
				if ctx.State.Alive {
					err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.DefaultKey])
				} else {
					err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.DeathKey])
				}
			}
		}

		if state.Player.State.Health != ctx.State.Health {
			ctx.State.Health = state.Player.State.Health

			if ctx.State.Health == 0 {
				if ctx.Debug {
					log.Println("dead")
				}
				ctx.State.Alive = false
				err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.DeathKey])
			} else if ctx.State.Health >= 100 {
				if ctx.Debug {
					log.Println("alive")
				}
				ctx.State.Alive = true
				err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.DefaultKey])
			}
		}

		if state.Round.Bomb != ctx.State.Bomb {
			ctx.State.Bomb = state.Round.Bomb
			if ctx.State.Bomb == "exploded" {
				if ctx.Debug {
					log.Println("exploded")
				}
				err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.BombKey])
			}
		}

		if state.Round.Phase != ctx.State.Phase {
			ctx.State.Phase = state.Round.Phase
			if ctx.State.Phase == "freezetime" {
				if ctx.Debug {
					log.Println("freezetime")
				}
				err = ctx.Pw.SetVolume(ctx.State.NodeId, ctx.Config.Volume[config.DefaultKey])
			}
		}
	}

	return err
}
