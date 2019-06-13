package tools

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/tealeg/xlsx"
)

var ctx = context.Background()
var esHost = "http://192.168.8.5:9200"

type excelReader struct{
	xlFile *xlsx.File
	filename string
	columnNum int
	headers map[string][]string
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


//************************esClient********************
type esClient struct {
	client *elastic.Client
	flushCount int
}


func (c *esClient) clientInit(host string) error {

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		fmt.Println("error:", err)
	}
	c.client = client
	return err
}

func (c *esClient) create(body interface{}, index string, id string) error {
	//str := `{"user" : "olive777re", "message" : "It777's a Raggy Waltz","sex":2,"hobby":"swimming, dota"}`
	temp := c.client.Index().
		Index(index)
	if id != "" {
		temp = temp.Id(id)
	}
	put2, err := temp.
		BodyJson(body).
		Do(ctx)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return err
	}
	c.flushCount += 1
	if c.flushCount >= 20{
		err = c.flush(index)
		if err != nil{
			return err
		}
	}
	fmt.Printf("Indexed document %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	return nil
}

func (c *esClient) flush(index string) error {
	_, err := c.client.Flush().Index(index).Do(ctx)
	return err
}
