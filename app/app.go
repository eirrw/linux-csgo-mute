package app

import (
	_ "embed"
	"fmt"
	"github.com/dank/go-csgsi"
	"github.com/getlantern/systray"
	"log"
	"os"
	"time"
	"virunus.com/linux-csgo-mute/config"
	"virunus.com/linux-csgo-mute/pipewire"
)

//go:embed icon_blue.png
var iconBlue []byte

//go:embed icon_red.png
var iconRed []byte

// context manages context and state for the application
type context struct {
	game     *csgsi.Game        // Holds the underlying GSI server implementation
	pipewire *pipewire.Pipewire // Holds an instance of the local pipewire server
	config   *config.Config     // Current configuration of the app
	state    *state             // Manages state of the game/player as reported by game
	debug    bool               // Set if debug messages should be logged
	nodeId   int                // nodeId to use for CSGO, so we don't have to look it up every time
	enabled  bool
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
	ctx := context{
		game:    csgsi.New(0),
		config:  config.New(),
		state:   &state{},
		debug:   debug,
		enabled: true,
	}

	// system tray icon
	go func() {
		systray.Run(func() {
			systray.SetIcon(iconBlue)
			menuEnabled := systray.AddMenuItem("Disable", "Enable or disable volume adjustment")
			systray.AddSeparator()
			menuQuit := systray.AddMenuItem("Quit", "Exit application")

			for {
				select {
				case <-menuQuit.ClickedCh:
					systray.Quit()
					os.Exit(0)
					return
				case <-menuEnabled.ClickedCh:
					if !ctx.enabled {
						ctx.enabled = true
						systray.SetIcon(iconBlue)
						menuEnabled.SetTitle("Disable")
					} else {
						ctx.enabled = false
						systray.SetIcon(iconRed)
						menuEnabled.SetTitle("Enable")
					}
				}
			}
		}, nil)
	}()

	go processGsiRequest(&ctx)

	if err := ctx.game.Listen(fmt.Sprintf(":%d", ctx.config.Gsi.Port)); err != nil {
		return err
	}

	return nil
}

// processGsiRequest will asynchronously handle requests to the GSI server and perform the required actions
func processGsiRequest(ctx *context) {
	for {
		select {
		case state := <-ctx.game.Channel:
			if state.Auth == nil || state.Auth.Token != ctx.config.Gsi.Token {
				fmt.Println("bad")
				continue
			}

			if !ctx.state.connected {
				ctx.pipewire = pipewire.New()
				ctx.pipewire.GetNodeIdByName(ctx.config.App.CsgoNodeName)
				if ctx.nodeId == -1 {
					if ctx.debug {
						log.Println("Could not find audio node, reloading")
					}
					continue
				} else {
					ctx.state.connected = true
				}
			}

			if ctx.enabled {
				err := handleState(state, ctx)
				if err != nil {
					log.Println(err)
				}
			}
		case <-time.After(time.Minute):
			if ctx.debug {
				log.Println("timeout")
			}
			ctx.state.connected = false
		}
	}
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
