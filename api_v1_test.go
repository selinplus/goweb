package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"testing"
)

func TestApiV1GetTag(t *testing.T){
	cli := gorequest.New()

	_,b,_ := cli.Get("http://localhost:8000/api/v1/tags").End()

	log.Println(b)
}
func TestApiV1AddTag(t *testing.T){
	cli := gorequest.New()

	_,b,_ := cli.Post("http://localhost:8000/api/v1/tags?name=1&state=1&created_by=test").End()

	log.Println(b)
}
