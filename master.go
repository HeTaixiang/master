package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/log"
)

func main() {
	log.Printf("create restful container")
	container := restful.NewContainer()
	for _, service := range S {
		service(container)
	}



}
