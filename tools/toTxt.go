package tools

import (
	"code.sajari.com/docconv"
	"fmt"
)

func toTxt(file string) (content string, res map[string]string, err error){
	res = make(map[string]string)
	fileRes, err := docconv.ConvertPath(file)
	if err != nil {
		fmt.Println(err)
	}
	return fileRes.Body, fileRes.Meta, err
}
