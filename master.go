package main

import (
	"github.com/emicklei/go-restful"
	"github.com/HeTaixiang/master/factory"
	"net/http"
	"github.com/golang/glog"
)

func main() {
	glog.V(1).Infof("create restful container...")
	container := restful.NewContainer()
	for service := range factory.GetServices() {
		service(container)
	}
	server := &http.Server{":8080", container}
	glog.Fatal(server.ListenAndServe())

}
