package tools

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"github.com/tealeg/xlsx"
)

var ctx = context.Background()
var esHost = beego.AppConfig.String("eshost")

type excelReader struct{
	xlFile *xlsx.File
	filename string
	columnNum int
	headers map[string][]string
	esC *elastic.Client
}

func (eR *excelReader)getFile(fileName string) error{
	var err error
	eR.headers = make(map[string][]string)
	eR.filename = fileName
	eR.xlFile, err = xlsx.OpenFile(fileName)
	if err != nil{
		return err
	}
	err = eR.setHeaders()
	if err != nil{
		return err
	}
	return nil
}

func (eR *excelReader)setHeaders() error{
	for _, sheet := range eR.xlFile.Sheets {
		row0 := sheet.Rows[0]
		eR.columnNum = len(row0.Cells)

		var temp []string
		for _, cell := range row0.Cells{
			temp = append(temp, cell.String())
		}
		eR.headers[sheet.Name] = temp
	}
	return nil
}

func (eR *excelReader)readAll() error{
	esC := esClient{}
	err := esC.clientInit(esHost)
	ins := map[string]interface{}{}
	ins["agg"] = "fuck"
	if err != nil{
		return err
	}

	for _, sheet := range eR.xlFile.Sheets{
		sheetName := sheet.Name
		flag := true
		for _, row := range sheet.Rows{
			if flag == true{
				flag = false
				continue
			}else{
				for i, cell := range row.Cells{
					ins[eR.headers[sheetName][i]] = cell.String()
					fmt.Println(eR.headers[sheetName][i],cell.String())
				}
				err = esC.create(ins, sheetName, "")
				if err != nil{
					return err
				}
			}
		}
	}
	return nil
}



