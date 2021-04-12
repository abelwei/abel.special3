package requests

import (
	"github.com/abelwei/abel.special3/general/dirfile"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestRequests_Post(t *testing.T) {
	dirfile.SetLogrusCoutColor()

	rqst := Requests{}
	rqst.IntiRequests()

	postData := map[string]string{
		"mobile":"111",
	}
	strHtml, err := rqst.Post("http://localhost:8921/auth/batchRegister", postData)
	if err == nil {
		if strings.Contains(`"code":1`, strHtml){
			logrus.Info("注册成功:")
		}else{
			logrus.Error(strHtml)
		}
	}else{
		logrus.Error(err)
	}
}
