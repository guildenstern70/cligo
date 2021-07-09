/*
 * CliGo Project
 * Copyright (c) 2021 Alessio Saltarin
 * MIT License
 */

package lib

import (
	"github.com/posener/complete"
	"strings"
)

// Message Implements cli.Command and cli.Error interfaces
type Message struct {
	msg string
}

func (m *Message) Help() string {
	return m.helpMessage()
}

func (m *Message) Run(args []string) int {
	println(m.helpMessage())
	return 0
}

func (m *Message) Synopsis() string {
	return "Get a default or a special message"
}

func (m *Message) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (*Message) AutocompleteFlags() complete.Flags {
	return complete.Flags{
		"-color": complete.PredictNothing,
		"-debug": complete.PredictNothing,
	}
}

func (m *Message) helpMessage() string {

	helpText := `
Usage: cligo message [default|special]

   Will print your customized message. The message can be the default one, or a special one.

Options:

  -color=false                  Disable color output. (Default: color)
  -debug                        Debug mode enabled for builds.

`

	return strings.TrimSpace(helpText)
}

func NewMessage(message string) *Message {
	t := &Message{}
	t.msg = message
	return t
}
