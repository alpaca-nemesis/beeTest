package tools

import (
	"code.sajari.com/docconv"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/tealeg/xlsx"
)

type docReader struct{
	xlFile *xlsx.File
	filename string
	columnNum int
	headers map[string][]string
	esC *elastic.Client
}


func (dR *docReader)toTxt(file string) (content string, res map[string]string, err error){
	res = make(map[string]string)
	fileRes, err := docconv.ConvertPath(file)
	if err != nil {
		fmt.Println(err)
	}
	return fileRes.Body, fileRes.Meta, err
}
