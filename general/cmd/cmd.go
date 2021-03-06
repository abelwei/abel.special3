package cmd

import (
	"bytes"
	"github.com/abelwei/abel.special3/general/encode"
	"github.com/sirupsen/logrus"
	"os/exec"
)

type Cmd struct {
	Windows bool
}

func (self *Cmd) Exec(runCommand string) (error, string) {
	//var commResu Result
	cmd := exec.Command("cmd","/C", runCommand)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		bStderr := stderr.Bytes()
		if self.Windows {
			bStderr = encode.GbkToUtf8(bStderr)
		}
		sMsg := string(bStderr)
		logrus.Error("RunCommand:", runCommand)
		logrus.Error("Cmd.Exec error:", sMsg)
		return err, ""
	}else{
		msg :=  out.String()
		return nil, msg
	}
}
