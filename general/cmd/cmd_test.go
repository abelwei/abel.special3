package cmd

import "testing"

func TestCmd_ExecOut(t *testing.T) {
	err, sOut := NewCmd(false).ExecOut("ls", "-a")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(sOut)
	}
}
