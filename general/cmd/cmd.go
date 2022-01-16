package cmd

import (
	"bytes"
	"github.com/abelwei/abel.special3/general/encode"
	"github.com/sirupsen/logrus"
	"os/exec"
)

type Cmd struct {
	Windows bool
	cmdObj  *exec.Cmd
}

func NewCmd(isWindows bool) *Cmd {
	return &Cmd{
		Windows: isWindows,
	}
}

//升级后Linux或windows用
func (self *Cmd) ExecOut(cmdName string, runCommand ...string) (error, string) {

	if self.Windows {
		runCommand = append([]string{"/C"}, runCommand...) //在头部添加元素/C命令，windows专有
		self.cmdObj = exec.Command("cmd", runCommand...)
	} else {
		self.cmdObj = exec.Command(cmdName, runCommand...)
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	self.cmdObj.Stdout = &out
	self.cmdObj.Stderr = &stderr
	err := self.cmdObj.Run()
	if err != nil {
		bStderr := stderr.Bytes()
		if self.Windows {
			bStderr = encode.GbkToUtf8(bStderr)
		}
		sMsg := string(bStderr)
		logrus.Error("RunCommand:", runCommand)
		logrus.Error("Cmd.Exec error:", sMsg)
		return err, ""
	} else {
		msg := out.String()
		return nil, msg
	}

}

//windows专用
func (self *Cmd) Exec(runCommand string) (error, string) {
	//var commResu Result
	cmd := exec.Command("cmd", "/C", runCommand)
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
	} else {
		msg := out.String()
		return nil, msg
	}
}
