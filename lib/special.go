/*
 * CliGo Project
 * Copyright (c) 2021 Alessio Saltarin
 * MIT License
 */

package lib

// SpecialMessage Implements cli.Command interfaces
type SpecialMessage struct {
	msg string
}

func (m *SpecialMessage) Help() string {
	return m.msg
}

func (m *SpecialMessage) Run(args []string) int {
	println("Bravo, this is your default message!")
	return 0
}

func (m *SpecialMessage) Synopsis() string {
	return "Get a default or a special message"
}

func NewSpecialMessage(message string) *DefaultMessage {
	t := &DefaultMessage{}
	t.msg = message
	return t
}
