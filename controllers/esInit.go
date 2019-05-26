package controllers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func (c *SearchController) clientInit() error{

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL("http://192.168.38.129:9200"))
	if err != nil {
		fmt.Println("error:", err)
	}
	c.client = client
	return err
}

func (c *SearchController) searchContent(str string) string{

	termQuery := elastic.NewTermQuery("user", str)
	//searchResult, err := c.client.Search().

	_, err := c.client.Search().
		Index("tweets").            // search in index "tweets"
		Query(termQuery).           // specify the query
		Sort("user.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(context.Background())    // execute
	if err != nil {
		fmt.Println("error:", err)
	}
	//var ttyp Tweet
	//for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	//	if t, ok := item.(Tweet); ok {
	//		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//		return t.Message
	//	}
	//}
	hehe := "caonima"
	return hehe
}

func (c *AddContentController) addIndex(str string) error{
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err := c.client.Index().
		Index("tweets").
		Type("doc").
		Id("1").
		BodyJson(tweet).
		Refresh("wait_for").
		Do(context.Background())
	return err
}

