package charlib

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"math/rand"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const (
	RandLetterAll = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	RandLetterLower = "abcdefghijklmnopqrstuvwxyz"
	RandLetterUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandLetterNumber = "0123456789"
)

//英文首字母大写
//或者原生 caseChar := "abc_def"
//strings.Title(caseChar)
func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122  {
		strArry[0] -=  32
	}
	return string(strArry)
}


//将首字母变成小写
func Lowercase(str string) (rst string) {
	if len(str) < 1 {
		return
	}
	strArry := []rune(str)
	if strArry[0] >= 65 && strArry[0] <= 90  {
		strArry[0] +=  32
	}
	rst = string(strArry)
	return
}

//将首字母变成大写
func Capitalize(str string) (rst string) {
	if len(str) < 1 {
		return
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122  {
		strArry[0] -=  32
	}
	rst = string(strArry)
	return
}

//下划线转驼峰
func Case2Camel(caseChar string) string {
	camelChar := strings.Replace(caseChar, "_", " ", -1)
	camelChar = strings.Title(camelChar)
	return strings.Replace(camelChar, " ", "", -1)
}

//驼峰转下划线
func Camel2Case(CamelChar string) string {
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchAllCap.ReplaceAllString(CamelChar, "${1}_${2}")
	return strings.ToLower(snake)
}


//返回随机字符串
func RandStringBytes(n int, letterBytes string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//是否是JSON格式
func IsJson(str string) bool {
	var jsn json.RawMessage
	return json.Unmarshal([]byte(str), &jsn) == nil
}

//将切片数组转为分隔符的字符串
func Slice2Separator(sli []string) (str string) {
	for _, val := range sli {
		str = str + val + "."
	}
	if str != "" {
		str = str[:len(str)-1]
	}
	return
}



func Template2Text(textTmpl string, srtDefault interface{}) (err error, rst string) {
	if textTmpl == "" {
		return errors.New("textTmpl is empty"), ""
	}
	tempCode, err := template.New("create").Parse(textTmpl)
	if err != nil {
		logrus.Error("template.create error:", err)
		return
	}
	tempBuf := new(bytes.Buffer)
	err = tempCode.Execute(tempBuf, srtDefault)
	if err != nil {
		logrus.Error("template.Execute error:",err)
		return
	}
	rst = tempBuf.String()
	return
}

func Template2TextWithFunc(tempCode string, srtDefault interface{}, funcs map[string]interface{}) (err error, rst string) {
	if tempCode == "" {
		return errors.New("tempCode is empty"), ""
	}
	paser, err := template.New("create").Funcs(funcs).Parse(tempCode)
	if err != nil {
		logrus.Error("template.create error:", err)
		return
	}
	tempBuf := new(bytes.Buffer)
	err = paser.Execute(tempBuf, srtDefault)
	if err != nil {
		logrus.Error("template.Execute error:",err)
		return
	}
	rst = tempBuf.String()
	return
}