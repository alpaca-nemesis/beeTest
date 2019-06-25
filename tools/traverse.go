package tools

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type readerIn interface {
	getFile(string) error
	readAll() error
}


var formats = []string{"docx", "pdf", "xlsx", "txt", "rtf"}
var uploadDir = beego.AppConfig.String("uploaddir")

func Traverse() error{
	fileSuffix := ""
	eR := excelReader{}
	dR := docReader{}
	eR.esC = &esClient{}
	dR.esC = &esClient{}
	files, err := GetAllFiles(uploadDir, formats)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files{
		fileSuffix = path.Ext(file)
		if fileSuffix == ".xlsx"{
			err = handler(&eR, file)
		}else{
			err = handler(&dR, file)
		}
		if err != nil{
			log.Println(err)
			return err
		}
	}
	return nil
}

func handler(r readerIn, file string)error{
	err := r.getFile(file)
	if err != nil{
		return err
	}else{
		err = r.readAll()
		if err != nil{
			return err
		}
	}
	return nil
}


func GetAllFiles(dirPth string, formats []string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			// 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
		} else {
			// 过滤指定格式
			ok := false
			for _, format := range formats{
				//忽略大小写
				ok = ok || strings.HasSuffix(strings.ToLower(fi.Name()), format)
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
