package main

import (
	. "github.com/guildenstern70/cligo/lib"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

// Returns the command and the error in case the command failed to load
func messageFactory() (cli.Command, error) {
	var messageCommand = NewMessage("A simple message")
	return messageCommand, nil
}

func main() {
	c := cli.NewCLI("CliGo", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"message": messageFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
