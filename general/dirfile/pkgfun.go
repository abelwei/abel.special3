package dirfile

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetFiles(dir string) (rtsFiles []string) {
	files, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, file := range files {
			rtsFiles = append(rtsFiles, file.Name())
		}
		return rtsFiles
	}
	return nil
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		logrus.Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func GetParentDirectory(currentDir string, lay int) string {
	sLay := "/"
	for i := 0; i < lay; i++ {
		sLay = sLay + "../"
	}
	return path.Dir(currentDir + sLay)
}

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		return string(data)
	} else {
		logrus.Error(err)
		return ""
	}

}

func WriteFile(filename, text string) bool {
	bt := []byte(text)
	err := ioutil.WriteFile(filename, bt, 0644)
	if err == nil {
		return true
	} else {
		logrus.Error(err)
		return false
	}
}

func WriterFile4Path(pathfile, text string) bool {
	dir := GetDir4Path(pathfile)
	if MkdirAll(dir) {
		return WriteFile(pathfile, text)
	} else {
		return false
	}
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func DirExist(dir string) bool {
	if info, err := os.Stat(dir); err == nil {
		if info.IsDir() {
			return true
		}
	} else {
		if os.IsNotExist(err) {
			return false
		} else {
			logrus.Error(err)
		}
	}
	return false
}

func FileExist(path string) bool {
	if info, err := os.Stat(path); err == nil {
		if info.IsDir() == false {
			return true
		}
	} else {
		if os.IsNotExist(err) {
			return false
		} else {
			logrus.Error(err)
		}
	}
	return false
}

func GetDir4Path(pathFile string) string {
	if dir, err := filepath.Abs(filepath.Dir(pathFile)); err == nil {
		return dir
	} else {
		logrus.Error(err)
		return ""
	}
}

func MkdirAll(dir string) bool {
	if PathExist(dir) == false {
		var appFs = afero.NewOsFs()
		if err := appFs.MkdirAll(dir, 0777); err == nil {
			return true
		} else {
			logrus.Error(err)
			return false
		}
	} else {
		return true
	}
}

func GetFileNameExt(filePath string) (filenameFull, filenameOnly, fileExt string) {
	filenameFull = path.Base(filePath)
	fileExt = path.Ext(filenameFull)
	filenameOnly = strings.TrimSuffix(filenameFull, fileExt)
	return
}

func GetFileNameFull(filePath string) string {
	str, _, _ := GetFileNameExt(filePath)
	return str
}

func GetFileExt(filePath string) string {
	_, _, str := GetFileNameExt(filePath)
	if str != "" {
		str = str[1:]
	}
	return str
}

func GetFileNameOnly(filePath string) string {
	_, str, _ := GetFileNameExt(filePath)
	return str
}

func ReadfileLine(filePath string) (rstss []string) {
	fi, err := os.Open(filePath)
	if err != nil {
		logrus.Errorf("dirfile.ReadfileLine Error: %s", err.Error())
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, _, err2 := br.ReadLine()
		if err2 != nil {
			if err2 != io.EOF { //如果不是读取到文本尽头的错误是非常规错误
				logrus.Errorf("dirfile.ReadfileLine Error2: %s", err2.Error())
			}
			break
		}
		rstss = append(rstss, string(line))
	}
	return
}
