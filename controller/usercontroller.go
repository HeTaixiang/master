package controller

import (
	"github.com/emicklei/go-restful"
	factory "github.com/HeTaixiang/master/factory"
	"github.com/emicklei/go-restful/log"
)
func init() {
	factory.RegisterServiceFactory(func(container *restful.Container) {
		log.Printf("callback register services")
		service := new(restful.WebService)
		service.Path("/users").Doc("manager user").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
		service.Route(service.GET("/{userid}").To(func(request *restful.Request, response *restful.Response) {
			response.WriteAsJson(map[string]string{"key":"value"})
		}))
		container.Add(service)
	})
}
