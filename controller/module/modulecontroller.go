package module

import (
	"github.com/HeTaixiang/master/db"
	"github.com/HeTaixiang/master/factory"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

func init() {
	factory.RegisterService("/modules", func(service *restful.WebService) {
		glog.V(1).Info("call back register server")
		service.Path("/modules").Doc("manager modules").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

		service.Route(service.GET("/").To(func(request *restful.Request, response *restful.Response) {
			moduleManager := NewManager(db.GetSessionManager())
			modules, err := moduleManager.Get()
			if err != nil {
				//TODO:

			}
			response.WriteAsJson(modules)
		}))

	})
}
