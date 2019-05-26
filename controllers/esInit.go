package controllers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)


func (c *SearchController) clientInit() error{

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL("http://192.168.8.4:9200"))
	if err != nil {
		fmt.Println("error:", err)
	}
	c.client = client
	return err
}

func createIndex(this *SearchController){

	// Create an index
	var _, err = this.client.CreateIndex("tweets").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}

