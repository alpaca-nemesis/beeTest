package tools

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var formats []string = []string{"docx", "pdf", "xlsx", "txt", "rtf"}

var uploadDir = beego.AppConfig.String("uploaddir")

func traverse() {
	files, err := GetAllFiles(uploadDir, formats)
	if err != nil {
		fmt.Println(err)
	}
	var fileSuffix string
	for _, file := range files{
		fileSuffix = path.Ext(file)
		fmt.Println(fileSuffix)
	}
}


func GetAllFiles(dirPth string, formats []string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() {
			// 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
		} else {
			// 过滤指定格式
			ok := false
			for _, format := range formats{
				ok = ok || strings.HasSuffix(fi.Name(), format)
				if ok {
					break
				}
			}
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}
	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table, formats)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}
	return files, nil
}
