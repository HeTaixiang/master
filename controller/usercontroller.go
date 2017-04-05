package controller

import (
	"github.com/emicklei/go-restful"
	"github.com/HeTaixiang/master/factory"
	"github.com/golang/glog"
)
func init() {
	factory.RegisterService("/users", func(container *restful.Container) {
		glog.V(1).Info("call back register server")
		service := new(restful.WebService)
		service.Path("/users").Doc("manager user").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
		service.Route(service.GET("/{userid}").To(func(request *restful.Request, response *restful.Response) {
			response.WriteAsJson(map[string]string{"key":"value"})
		}))
		container.Add(service)
	})
}
