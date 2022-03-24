package app

import (
	"fmt"
	"github.com/dank/go-csgsi"
	"log"
	"time"
	"virunus.com/linux-csgo-mute/config"
	"virunus.com/linux-csgo-mute/pipewire"
)

// context manages context and state for the application
type context struct {
	game     *csgsi.Game        // Holds the underlying GSI server implementation
	pipewire *pipewire.Pipewire // Holds an instance of the local pipewire server
	config   *config.Config     // Current configuration of the app
	state    *state             // Manages state of the game/player as reported by game
	debug    bool               // Set if debug messages should be logged
	nodeId   int                // nodeId to use for CSGO, so we don't have to look it up every time
}

// state manages the current states of the player/game as reported by the GSI
type state struct {
	alive     bool   // If the player is alive
	bomb      string // Current state of the bomb
	connected bool   // If the game if connected
	flashed   int    // Amount that the player is flashed
	health    int    // Amount of health that the player has
	phase     string // Current phase of the round
}

// Start loads the config and pipewire components then starts the GSI server
func Start(debug bool) error {
	context := context{
		game:   csgsi.New(0),
		config: config.New(),
		state:  &state{},
		debug:  debug,
	}

	go func() {
		for {
			select {
			case state := <-context.game.Channel:
				if state.Auth == nil || state.Auth.Token != context.config.Gsi.Token {
					fmt.Println("bad")
					continue
				}

				if !context.state.connected {
					context.pipewire = pipewire.New()
					context.nodeId = context.pipewire.GetNodeIdByName(context.config.App.CsgoNodeName)
					if context.nodeId == -1 {
						if context.debug {
							log.Println("Could not find audio node, reloading")
						}
						continue
					} else {
						context.state.connected = true
					}
				}

				err := handleState(state, &context)
				if err != nil {
					log.Println(err)
				}
			case <-time.After(time.Minute):
				if context.debug {
					log.Println("timeout")
				}
				context.state.connected = false
			}
		}
	}()

	if err := context.game.Listen(fmt.Sprintf(":%d", context.config.Gsi.Port)); err != nil {
		return err
	}

	return nil
}

// handleState reads the incoming state from the GSI and performs the necessary actions based on the input
func handleState(state csgsi.State, ctx *context) error {
	var err error
	if state.Player != nil && state.Round != nil {
		if state.Player.State.Flashed != ctx.state.flashed {
			ctx.state.flashed = state.Player.State.Flashed
			if ctx.state.flashed > 200 {
				if ctx.debug {
					log.Printf("flashed: %d\n", ctx.state.flashed)
				}
				err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.FlashKey])
			} else {
				if ctx.debug {
					log.Println("not flashed")
				}
				if ctx.state.alive {
					err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.DefaultKey])
				} else {
					err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.DeathKey])
				}
			}
		}

		if state.Player.State.Health != ctx.state.health {
			ctx.state.health = state.Player.State.Health

			if ctx.state.health == 0 {
				if ctx.debug {
					log.Println("dead")
				}
				ctx.state.alive = false
				err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.DeathKey])
			} else if ctx.state.health >= 100 {
				if ctx.debug {
					log.Println("alive")
				}
				ctx.state.alive = true
				err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.DefaultKey])
			}
		}

		if state.Round.Bomb != ctx.state.bomb {
			ctx.state.bomb = state.Round.Bomb
			if ctx.state.bomb == "exploded" {
				if ctx.debug {
					log.Println("exploded")
				}
				err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.BombKey])
			}
		}

		if state.Round.Phase != ctx.state.phase {
			ctx.state.phase = state.Round.Phase
			if ctx.state.phase == "freezetime" {
				if ctx.debug {
					log.Println("freezetime")
				}
				err = ctx.pipewire.SetVolume(ctx.nodeId, ctx.config.Volume[config.DefaultKey])
			}
		}
	}

	return err
}
