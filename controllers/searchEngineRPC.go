package controllers

import "github.com/go-ego/riot/types"


type SearchRequest struct {
	Content string
}

type SearchResponse struct {
	Content types.SearchResp
}