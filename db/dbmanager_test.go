package db

import (
	"testing"

	"github.com/HeTaixiang/master/controller"
	"github.com/HeTaixiang/master/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestNewManager(t *testing.T) {
	config, _ := utils.ReadToConfig("../db.json")
	_, err := NewManager(config)
	if err != nil {
		t.Logf("create mongodb connect failure %v", err)
	} else {
		t.Log("create mongodb connect success !")
	}
}

func TestExecuteFor(t *testing.T) {
	config, err := utils.ReadToConfig("../db.json")
	if err != nil {
		t.Logf("parse db.json failure v%", err)
		t.Fail()
	}
	sessionManager, err := NewManager(config)
	if err != nil {
		t.Logf("create mongodb session failure v%", err)
		t.Fail()
	}
	var module *controller.Module
	err = sessionManager.ExecuteFor("module", func(c *mgo.Collection) error {

		return c.Find(bson.M{}).One(module)
	})
	t.Logf("select data v%", module)

}
