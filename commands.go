package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar:      "ENV_MAJOR",
		Name:        "major, M",
		Usage:       "",
		Destination: &major,
	},
	cli.BoolFlag{
		EnvVar:      "ENV_MINOR",
		Name:        "minor, m",
		Usage:       "",
		Destination: &minor,
	},
	cli.BoolFlag{
		EnvVar:      "ENV_PATCH",
		Name:        "patch, p",
		Usage:       "default true",
		Destination: &patch,
	},
}

var Commands = []cli.Command{}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
