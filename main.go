package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"virunus.com/linux-csgo-mute/app"
	"virunus.com/linux-csgo-mute/config"
)

func main() {
	testPtr := flag.Bool("t", false, "Print the current config and exit")
	writePtr := flag.Bool("C", false, "Write the current config to disk and exit, creating or overwriting the config file as needed")
	debugFlag := flag.Bool("d", false, "Print debug info")

	flag.Parse()

	if flag.NFlag() > 1 {
		log.Print("Too many flags")
		os.Exit(1)
	}

	// test config
	if *testPtr {
		c := config.New()

		str, err := c.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(str))

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
	err := app.Start(*debugFlag)
	if err != nil {
		log.Fatal(err)
	}
}
