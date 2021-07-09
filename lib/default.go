/*
 * CliGo Project
 * Copyright (c) 2021 Alessio Saltarin
 * MIT License
 */

package lib

// DefaultMessage Implements cli.Command interfaces
type DefaultMessage struct {
	msg string
}

func (m *DefaultMessage) Help() string {
	return m.msg
}

func (m *DefaultMessage) Run(args []string) int {
	println("Bravo, this is your default message!")
	return 0
}

func (m *DefaultMessage) Synopsis() string {
	return "Get a default or a special message"
}

func NewDefaultMessage(message string) *DefaultMessage {
	t := &DefaultMessage{}
	t.msg = message
	return t
}
