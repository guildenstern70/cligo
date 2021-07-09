/*
 * CliGo Project
 * Copyright (c) 2021 Alessio Saltarin
 * MIT License
 */

package main

import (
	. "github.com/guildenstern70/cligo/lib"
	"github.com/guildenstern70/cligo/lib/termcolor"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

// Returns the command and the error in case the command failed to load
func messageFactory() (cli.Command, error) {
	var messageCommand = NewMessage("Get a message")
	return messageCommand, nil
}

func main() {
	c := cli.NewCLI("cligo", "0.0.1")
	c.Args = os.Args[1:]
	println(termcolor.Purple + "cligo" + termcolor.Reset)
	println("")

	c.Commands = map[string]cli.CommandFactory{
		"message": messageFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
