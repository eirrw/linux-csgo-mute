package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"virunus.com/linux-csgo-mute/app"
	"virunus.com/linux-csgo-mute/config"
)

func main() {
	testPtr := flag.Bool("t", false, "Print the current config and exit")
	writePtr := flag.Bool("C", false, "Write the current config to disk and exit, creating or overwriting the config file as needed")

	flag.Parse()

	if flag.NFlag() > 1 {
		log.Print("Too many flags")
		os.Exit(1)
	}

	// test config
	if *testPtr {
		c := config.New()

		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(c); err != nil {
			log.Fatal(err)
		}

		fmt.Println(buf.String())

		os.Exit(0)
	}

	// write config
	if *writePtr {
		c := config.New()
		if err := c.WriteFile(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	// run app
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
