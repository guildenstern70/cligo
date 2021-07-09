/*
 * CliGo Project
 * Copyright (c) 2021 Alessio Saltarin
 * MIT License
 */

package lib

import (
	"github.com/guildenstern70/cligo/lib/termcolor"
	"testing"
)

func TestSpecialMessage_Run(t *testing.T) {

	var cmd = NewSpecialMessage("")
	var errorCode = cmd.Run(nil)
	if errorCode == 0 {
		t.Logf("Default Message Cmd " + termcolor.Cyan + "OK" + termcolor.Reset)
	} else {
		t.Errorf("Default Message Cmd  " + termcolor.Red + "not working" + termcolor.Reset)
	}
}
