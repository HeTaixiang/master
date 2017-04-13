package main

import (
	"flag"
	"net/http"

	_ "github.com/HeTaixiang/master/controller"
	factory "github.com/HeTaixiang/master/factory"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.V(1).Infof("create restful container...")
	container := restful.NewContainer()
	for _, service := range factory.GetServices() {
		service(container)
	}
	server := &http.Server{Addr: ":8080", Handler: container}
	glog.Fatal(server.ListenAndServe())

}
