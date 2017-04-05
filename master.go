package main

import (
	"github.com/emicklei/go-restful"
	factory "github.com/HeTaixiang/master/factory"
	_ "github.com/HeTaixiang/master/controller"
	"net/http"
	"github.com/golang/glog"
)

func main() {
	glog.V(1).Infof("create restful container...")
	container := restful.NewContainer()
	for _, service := range factory.GetServices() {
		service(container)
	}
	server := &http.Server{Addr:":8080", Handler:container}
	glog.Fatal(server.ListenAndServe())

}
