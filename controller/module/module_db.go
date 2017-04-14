package module

import (
	"github.com/HeTaixiang/master/controller"
	"github.com/HeTaixiang/master/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Manager used to manage Module with database
type Manager interface {
	Get() ([]*controller.Module, error)

	GetByID(ID string) (*controller.Module, error)

	GetSubModulesByModuleID(ID string) ([]*controller.SubModule, error)

	GetSubModuleByID(ID string) (*controller.SubModule, error)

	Add(module *controller.Module) (string, error)

	Update(module *controller.Module) error

	DeleteByID(ID string) error
}

const moduleName = "module"
const subModuleName = "submodule"

// NewManager create module manager
func NewManager(session db.SessionInterface) Manager {
	return &moduleManager{session}
}

type moduleManager struct {
	session db.SessionInterface
}

func (manager *moduleManager) Get() ([]*controller.Module, error) {
	var modules []*controller.Module
	err := manager.session.ExecuteFor(moduleName, func(c *mgo.Collection) error {
		return c.Find(nil).All(modules)
	})
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (manager *moduleManager) GetByID(ID string) (*controller.Module, error) {
	var module *controller.Module
	err := manager.session.ExecuteFor(moduleName, func(c *mgo.Collection) error {
		return c.FindId(bson.ObjectIdHex(ID)).One(module)
	})
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (manager *moduleManager) GetSubModulesByModuleID(ID string) ([]*controller.SubModule, error) {
	var module *controller.Module
	err := manager.session.ExecuteFor(moduleName, func(c *mgo.Collection) error {
		return c.FindId(bson.ObjectIdHex(ID)).One(module)
	})
	if err != nil {
		return nil, err
	}
	return module.SubModules, nil
}

func (manager *moduleManager) GetSubModuleByID(ID string) (*controller.SubModule, error) {
	return new(controller.SubModule), nil
}

func (manager *moduleManager) Add(module *controller.Module) (string, error) {
	module.ID = bson.NewObjectId()
	err := manager.session.ExecuteFor(moduleName, func(c *mgo.Collection) error {
		return c.Insert(module)
	})
	if err != nil {
		return "", err
	}
	return module.ID.Hex(), nil
}

func (manager *moduleManager) Update(module *controller.Module) error {
	return manager.session.ExecuteFor(moduleName, func(c *mgo.Collection) error {
		return c.UpdateId(module.ID, module)
	})
}

func (manager *moduleManager) DeleteByID(ID string) error {
	return manager.session.ExecuteFor(moduleName, func(c *mgo.Collection) error {
		return c.RemoveId(bson.IsObjectIdHex(ID))
	})
}
