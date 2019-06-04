package controllers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
	"strconv"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
var ctx = context.Background()

func (c *SearchController) clientInit() error{

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL("http://192.168.38.132:9200"))
	if err != nil {
		fmt.Println("error:", err)
	}
	c.client = client
	return err
}

func (c *SearchController) searchContent(str string) string{

	termQuery := elastic.NewTermQuery("user", str)
	//searchResult, err := c.client.Search().

	searchResult, err := c.client.Search().
		Index("twitter").            // search in index "tweets"
		Query(termQuery).           // specify the query
		//Sort("user.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(ctx)    // execute

	if err != nil {
		fmt.Println("error:", err)
	}
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
			return t.Message
		}
	}
	hehe := "caonima"
	return hehe
}

func (c *AddContentController) addIndex(body interface{}, index string, id int) error{
	//str := `{"user" : "olive777re", "message" : "It777's a Raggy Waltz","sex":2,"hobby":"swimming, dota"}`
	temp := c.client.Index().
		Index(index)
	if id != 0{
		temp = temp.Id(strconv.Itoa(id))
	}
	put2, err := temp.
		BodyJson(body).
		Do(ctx)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Printf("Indexed document %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	_, err = c.client.Flush().Index(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

