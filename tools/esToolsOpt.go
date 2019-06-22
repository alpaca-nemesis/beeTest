package tools

import (
	"fmt"
	"github.com/olivere/elastic"
)

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