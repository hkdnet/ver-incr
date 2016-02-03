package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var major, minor, patch bool

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "hkdnet"
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Action = func(c *cli.Context) {
		if len(c.Args()) > 0 {
			// Errorのはず
			fmt.Printf("Please use with file path.\n")
			return
		}

		if (major && minor) || (minor && patch) || (patch && major) {
			// Errorのはず
			fmt.Printf("Do not set more than one flag true\n")
			return
		}
		switch {
		case major:
			fmt.Printf("major\n")
		case minor:
			fmt.Printf("minor\n")
		default:
			fmt.Printf("patch\n")
		}
	}
	app.Run(os.Args)
}
