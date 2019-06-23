package tools

import (
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"log"
)

//************************esClient********************
type esClient struct {
	client *elastic.Client
	flushCount int
}

var esHost = beego.AppConfig.String("eshost")

func (c *esClient) clientInit() error {

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL(esHost))
	if err != nil {
		//log.Println("error:", err)
		return err
	}
	c.client = client
	return err
}

func (c *esClient) create(body interface{}, index string, id string) error {
	if c.client == nil{
		err := c.clientInit()
		if err != nil{
			return err
		}
	}
	temp := c.client.Index().
		Index(index)
	if id != "" {
		temp = temp.Id(id)
	}
	put2, err := temp.
		BodyJson(body).
		Do(ctx)
	if err != nil {
		return err
	}
	c.flushCount += 1
	if c.flushCount >= 20{
		err = c.flush(index)
		if err != nil{
			return err
		}
	}
	log.Printf("Indexed document %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	return nil
}

func (c *esClient) flush(index string) error {
	_, err := c.client.Flush().Index(index).Do(ctx)
	return err
}