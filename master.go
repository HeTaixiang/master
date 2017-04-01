package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/log"
	"github.com/HeTaixiang/master/factory"
)

func main() {
	log.Printf("create restful container")
	container := restful.NewContainer()
	for service := range factory.GetServices() {
		service(container)
	}



}
