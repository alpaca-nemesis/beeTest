package tools

import (
	"code.sajari.com/docconv"
	"path"
)

type docReader struct{
	docFile *docconv.Response
	filename string
	docType string
	esC *esClient
}


func (dR *docReader) getFile(filename string) error{
	var err error
	dR.filename = filename
	dR.docType = path.Ext(filename)[1:]
	dR.docFile, err = docconv.ConvertPath(filename)
	return err
	//return fileRes.Body, fileRes.Meta, err
}

func (dR *docReader) readAll() error{
	ins := map[string]interface{}{}
	ins["content"] = dR.docFile.Body
	ins["filePath"] = dR.filename
	for key, value := range dR.docFile.Meta{
		if key == ""{
			continue
		}
		ins[key] = value
	}
	err := dR.esC.create(ins, dR.docType, "")
	return err
}
