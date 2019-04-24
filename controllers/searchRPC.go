package controllers

import (
	"github.com/go-ego/riot/types"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)


type SearchRequest struct {
	Content string
}

type SearchResponse struct {
	Content types.SearchResp
}

type AddRequest struct {
	Compulsive bool
	Content string
}

type AddResponse struct {
	Content string
}

func rpcInit() *rpc.Client{
	rpcClient, err := jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("dailing error: ", err)
		return nil
	}else{
		return rpcClient
	}
}


func search(content string, rpcClient *rpc.Client) types.SearchResp {
	req := SearchRequest{content}
	var res SearchResponse
	err := rpcClient.Call("RPCEngine.Search", req, &res)
	if err != nil {
		log.Fatalln("search error: ", err)
	}
	return res.Content
}

func addContent(content string, rpcClient *rpc.Client) string {
	req := AddRequest{false, content}
	var res AddResponse
	err := rpcClient.Call("RPCEngine.AddContent", req, &res)
	if err != nil {
		log.Fatalln("search error: ", err)
	}
	return res.Content
}