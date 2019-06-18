package controllers

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"log"
	"reflect"
	"strings"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
var ctx = context.Background()
var esHost = beego.AppConfig.String("eshost")

func (c *SearchController) clientInit() error{

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL(esHost))
	if err != nil {
		c.client = nil
	}else{
		c.client = client
	}
	return err
}

func (c *SearchController) searchContent(str string) string{

	termQuery := elastic.NewTermQuery("user", str)
	//searchResult, err := c.client.Search().

	searchResult, err := c.client.Search().
		Index("twitter").            // search in index "tweets"
		Query(termQuery).           // specify the query
		//Sort("user.keyword", true). // sort by "user" field, ascending
		//From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(ctx)    // execute

	if err != nil {
		fmt.Println("error:", err)
	}
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		//fmt.Println(reflect.TypeOf(item))
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
			return t.Message
		}
	}
	hehe := "caonima"
	return hehe
}



//------------------------AddController

func (c *AddContentController) clientInit() error{

	// Create a client
	var client, err = elastic.NewClient(elastic.SetURL(esHost))
	if err != nil {
		c.client = nil
		return err
	}else{
		c.client = client
		return nil
	}
}


func (c *AddContentController) addIndex(body interface{}, index string, id string) (string, error){
	temp := c.client.Index().
		Index(index)
	if id != ""{
		temp = temp.Id(id)
	}
	res, err := temp.
		BodyJson(body).
		Do(ctx)
	if err != nil {
		return "", err
	}
	fmt.Printf("Indexed document %s to index %s, type %s\n", res.Id, res.Index, res.Type)

	//flush the index
	_, err = c.client.Flush().Index(index).Do(ctx)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

func (c *AddContentController) deleteIndex(index string, id string) error{

	res, err := c.client.Delete().Index(index).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("delete result %s\n", res.Result)
	return nil
}

func (c *AddContentController) catIndices() elastic.CatIndicesResponse{
	indices, err := c.client.CatIndices().Do(ctx)
	if err != nil{
		log.Println("client err: ", err)
	}
	return indices
}


func getIndexName(res elastic.CatIndicesResponse) []string{
	var indexNames []string
	for _, val := range(res){
		if strings.HasPrefix(val.Index, "."){
			continue
		}else{
			indexNames = append(indexNames, val.Index)
		}
	}
	return indexNames
}

