package main

import (
	"fmt"
	"github.com/dank/go-csgsi"
	"log"
	"os/exec"
	"sync"
	"time"
	"virunus.com/linux-csgo-mute/app"
	"virunus.com/linux-csgo-mute/config"
	"virunus.com/linux-csgo-mute/pipewire"
)

const CsgoNodeName = "csgo_linux64"

func main() {
	context := app.Context{
		Game:   csgsi.New(0),
		Mu:     &sync.Mutex{},
		Config: config.New(),
		State:  &app.State{},
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
					context.State.NodeId = getNodeIdByName(CsgoNodeName, context)
					if context.State.NodeId == -1 {
						log.Print("Could not find audio node, reloading")
						continue
					} else {
						context.State.Connected = true
					}
				}

				handleState(state, &context)
			case <-time.After(time.Minute):
				log.Print("timeout...")
				context.State.Connected = false
			}
		}
	}()

	context.Game.Listen(fmt.Sprintf(":%d", context.Config.Gsi.Port))
}

func handleState(state csgsi.State, ctx *app.Context) {
	if state.Player != nil && state.Round != nil {
		if state.Player.State.Flashed != ctx.State.Flashed {
			ctx.State.Flashed = state.Player.State.Flashed
			if ctx.State.Flashed > 0 {
				fmt.Println("flashed")
				setVolume(ctx.State.NodeId, ctx.Config.Volume.Flash, ctx)
			} else if ctx.State.Flashed < 200 {
				fmt.Println("not flashed")
				if ctx.State.Alive {
					setVolume(ctx.State.NodeId, ctx.Config.Volume.Default, ctx)
				} else {
					setVolume(ctx.State.NodeId, ctx.Config.Volume.Death, ctx)
				}
			}
		}

		if state.Player.State.Health != ctx.State.Health {
			ctx.State.Health = state.Player.State.Health

			if ctx.State.Health == 0 {
				fmt.Println("dead")
				ctx.State.Alive = false
				setVolume(ctx.State.NodeId, ctx.Config.Volume.Death, ctx)
			} else if ctx.State.Health >= 100 {
				fmt.Println("alive")
				ctx.State.Alive = true
				setVolume(ctx.State.NodeId, ctx.Config.Volume.Default, ctx)
			}
		}

		if state.Round.Bomb != ctx.State.Bomb {
			ctx.State.Bomb = state.Round.Bomb
			if ctx.State.Bomb == "exploded" {
				fmt.Println("exploded")
				setVolume(ctx.State.NodeId, ctx.Config.Volume.Bomb, ctx)
			}
		}

		if state.Round.Phase != ctx.State.Phase {
			ctx.State.Phase = state.Round.Phase
			if ctx.State.Phase == "freezetime" {
				fmt.Println("freezetime")
				setVolume(ctx.State.NodeId, ctx.Config.Volume.Default, ctx)
			}
		}
	}
}

func getNodeIdByName(nodeName string, ctx app.Context) int {
	for _, node := range ctx.Pw.Nodes {
		if nodeName == node.Info.Props.NodeName {
			return node.Id
		}
	}

	return -1
}

func setVolume(nodeId int, vol float32, ctx *app.Context) {
	ctx.Mu.Lock()
	defer ctx.Mu.Unlock()
	cmd := exec.Command(
		pipewire.PwCliCmd,
		pipewire.PwCliSetOpt,
		fmt.Sprintf("%d", nodeId),
		pipewire.PwCliPropsOpt,
		fmt.Sprintf(pipewire.PwCliVolProp, vol),
	)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
