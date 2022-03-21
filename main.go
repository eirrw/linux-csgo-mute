package main

import (
	"fmt"
	"github.com/dank/go-csgsi"
)

var flashed, health int
var alive bool
var bomb, phase string

func main() {
	game := csgsi.New(10)

	go func() {
		for state := range game.Channel {
			if state.Auth == nil || state.Auth.Token != "VALVEPLSSPAREMYEARS" {
				fmt.Println("bad")
				continue
			}

			handleState(state)
		}
	}()

	game.Listen(":3202")
}

func handleState(state csgsi.State) {
	if state.Player.State.Flashed != flashed {
		flashed = state.Player.State.Flashed
		if flashed > 0 {
			fmt.Println("flashed")
			// set vol to flash
		} else if flashed < 200 {
			fmt.Println("not flashed")
			// vol to max
		}
	}

	if state.Player.State.Health != health {
		health = state.Player.State.Health

		if health == 0 {
			fmt.Println("dead")
			alive = false
			// set vol to death
		} else if health >= 100 {
			fmt.Println("alive")
			alive = true
			// set vol to max
		}
	}

	if state.Round.Bomb != bomb {
		bomb = state.Round.Bomb
		if bomb == "exploded" {
			fmt.Println("exploded")
			// set vol to bomb
		}
	}

	if state.Round.Phase != phase {
		phase = state.Round.Phase
		if phase == "freezetime" {
			fmt.Println("freezetime")
			// set vol to max
		}
	}
}
