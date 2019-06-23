package tools

import (
	"context"
	"github.com/tealeg/xlsx"
)

var ctx = context.Background()

type excelReader struct{
	xlFile *xlsx.File
	filename string
	columnNum int
	headers map[string][]string
	esC *esClient
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
	var err error
	ins := map[string]interface{}{}

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
					//log.Println(eR.headers[sheetName][i],cell.String())
				}
				err = eR.esC.create(ins, sheetName, "")
				if err != nil{
					return err
				}
			}
		}
	}
	return nil
}



