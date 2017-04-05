package factory

import "github.com/emicklei/go-restful"

type RegisterServiceFn func(container *restful.Container)
type Services []RegisterServiceFn

var S = make(Services, 0)

func RegisterServiceFactory(service RegisterServiceFn)  {
	S = append(S, service)
}
