package main

import (
	"fmt"

	"github.com/HeTaixiang/master/controller"
	"github.com/HeTaixiang/master/db"
	"github.com/HeTaixiang/master/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	config, err := utils.ReadToConfig("../../db.json")
	if err != nil {
		fmt.Printf("parse db.json failure %v", err)

	}
	sessionManager, err := db.NewManager(config)
	if err != nil {
		fmt.Printf("create mongodb session failure %v", err)

	}
	// var module = new(controller.Module)
	// module.ID = bson.NewObjectId()
	// module.Name = "test"
	// module.IsMain = true
	// module.Icon = "/img.png"
	// module.Path = "/test"
	// module.Weight = 500
	// submodule := &controller.SubModule{
	// 	Weight: 1000,
	// 	Name:   "test",
	// 	Icon:   "/img.png",
	// 	Path:   "/test"}
	// module.SubModules = append(module.SubModules, submodule)

	// err = sessionManager.ExecuteFor("module", func(c *mgo.Collection) error {
	// 	return c.Insert(module)
	// })

	modules := make([]*controller.Module, 0)
	err = sessionManager.ExecuteFor("module", func(c *mgo.Collection) error {
		return c.Find(nil).All(&modules)
	})
	for _, m := range modules {
		fmt.Printf("the result value %v\n", m)
	}

	var m = new(controller.Module)
	err = sessionManager.ExecuteFor("module", func(c *mgo.Collection) error {
		return c.FindId(bson.ObjectIdHex("590aedb8c2485ba3dd23c709")).One(m)
	})
	for _, sb := range m.SubModules {
		fmt.Printf("------ %v -----\n", sb.Icon)
	}

	// fmt.Printf("select data %v \n", module)
}
