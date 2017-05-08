package factory

import (
	"sync"

	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

// Factory is type of the webservice register factory function
type Factory func(service *restful.WebService)

var (
	container    = restful.NewContainer()
	serviceMutex sync.Mutex
	services     = make(map[string]Factory)
)

// RegisterService is for register webservice
func RegisterService(path string, service Factory) {
	serviceMutex.Lock()
	defer serviceMutex.Unlock()
	_, found := services[path]
	if found {
		glog.Errorf("service factory %q has been registered", path)
	}
	glog.V(1).Infof("registered service %q", path)
	services[path] = service
}

// GetServices is get all register webservice
func GetServices() []Factory {
	serviceMutex.Lock()
	defer serviceMutex.Unlock()
	wsServices := []Factory{}
	for _, s := range services {
		wsServices = append(wsServices, s)
	}
	return wsServices
}

// NewContainer add WebService and return restful.Container
func NewContainer() *restful.Container {
	for _, service := range GetServices() {
		s := new(restful.WebService)
		service(s)
		container.Add(s)
	}
	return container
}
