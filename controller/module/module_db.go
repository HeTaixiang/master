package module

import (
	"github.com/HeTaixiang/master/controller"
)

// Manager used to manage Module with database
type Manager interface {
	Get() ([]*controller.Module, error)

	GetByID(ID int64) (*controller.Module, error)

	GetSubModulesByModuleID(ID int64) ([]*controller.SubModule, error)

	GetSubModuleByID(ID int64) (*controller.SubModule, error)

	Add(module *controller.Module) error

	Update(module *controller.Module) error

	DeleteByID(ID int16) error
}

type moduleManager struct {
}
