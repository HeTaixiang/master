package module

import (
	"github.com/HeTaixiang/master/controller"
)

// Manager used to manage Module with database
type Manager interface {
	Get() []*controller.Module

	GetSubModulesByID(ID int64) []*controller.SubModule
}

type moduleManager struct {
}
