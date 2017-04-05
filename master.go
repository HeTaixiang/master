package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/log"
	factory "github.com/HeTaixiang/master/factory"
	_ "github.com/HeTaixiang/master/controller"
	"net/http"
)

func main() {
	log.Printf("create restful container")
	container := restful.NewContainer()
	for _, service := range factory.S {
		service(container)
	}

	server := &http.Server{Addr:":8080", Handler:container}
	log.Print(server.ListenAndServe())

}
