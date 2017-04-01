package controller

import (
	"github.com/emicklei/go-restful"
	"github.com/HeTaixiang/master/factory"
)
func init() {
	factory.RegisterService("user", func(container restful.Container) {

	})
}
