package factory

import (
	"github.com/emicklei/go-restful"
	"sync"
	"github.com/golang/glog"
)

type Factory func(container *restful.Container)

var (
	serviceMutex sync.Mutex
	services = make(map[string]Factory)
)

func RegisterService(path string, service Factory)  {
	serviceMutex.Lock()
	defer serviceMutex.Unlock()
	_, found := services[path]
	if found {
		glog.Fatal("service factory %q has been registered", path)
	}
	glog.V(1).Infof("registered service %q", path)
	services[path] = service
}

func GetServices()[]Factory  {
	serviceMutex.Lock()
	defer serviceMutex.Unlock()
	wsServices := []Factory{}
	for _, s := range services{
		wsServices = append(wsServices, s)
	}
	return wsServices
}
