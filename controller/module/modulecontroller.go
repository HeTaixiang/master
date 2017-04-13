package module

import (
	"github.com/HeTaixiang/master/factory"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

func init() {
	factory.RegisterService("/modules", func(container *restful.Container) {
		glog.V(1).Info("call back register server")
		service := new(restful.WebService)
		service.Path("/modules").Doc("manager modules").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
		defer container.Add(service)

		service.Route(service.GET("/").To(func(request *restful.Request, response *restful.Response) {
			response.WriteAsJson([]string{"key", "value"})
		}))

	})
}
