package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/youkyll/tw/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "login",
		Usage:  "",
		Action: command.CmdLogin,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "logout",
		Usage:  "",
		Action: command.CmdLogout,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "post",
		Usage:  "",
		Action: command.CmdPost,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "whoami",
		Usage:  "",
		Action: command.CmdWhoami,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
