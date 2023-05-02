package tools

import (
	"os"
	"runtime"
	"strings"
)

const (
	AppNamePlaceholder = "{{.pkgname}}"
)

// IsExists 判断文件是否存在
func IsExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return os.IsNotExist(err)
}

// IsExistsDirectoryAndCreate 判断文件夹是否存在,不存在则创建
func IsExistsDirectoryAndCreate(dir string) bool {
	_, err := os.ReadDir(dir)
	if err != nil {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

// GetGOPath 获取gopath路径
func GetGOPath() string {
	return os.Getenv("GOPATH")
}

// GetGoVersion 获取golang版本
func GetGoVersion() string {
	strArray := strings.Split(runtime.Version()[2:], `.`)
	return strArray[0] + `.` + strArray[1]
}

// ReplaceAppNameTemplate 替换模版中的appname
func ReplaceAppNameTemplate(template string, appName string) string {
	tempStr := strings.Replace(template, "#", "`", -1)
	return strings.Replace(tempStr, AppNamePlaceholder, appName, -1)
}

// WriteToFile 写入文件
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}
