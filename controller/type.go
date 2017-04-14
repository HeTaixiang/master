package controller

import (
	"gopkg.in/mgo.v2/bson"
)

// Item is the base struct
type Item struct {
	// name
	Name string `json:"name,omitempty" bson:"name"`
	// Icon
	Icon string `json:"icon,omitempty" bson:"icon"`
	// Weight used for order
	Weight int `json:"weight,omitempty" bson:"weight"`
	// Path used for route
	Path string `json:"route,omitempty" bson:"path"`
}

// Module is the main navigator
type Module struct {
	//ID
	ID bson.ObjectId `json:"id" bson:"_id"`
	// Item
	Item `json:",inline" bson:",inline"`
	// IsMain is used for the Main Web
	IsMain bool `json:"isMain,omitempty" bson:"isMain"`
	// SubModules
	SubModules []*SubModule `json:"subModules" bson:"subModules"`
}

// SubModule used for sub navigation
type SubModule Item

// Menu for paged scope navigation
type Menu struct {
	// Item
	Item `json:",inline"`
	// MenuItems subMenu
	MenuItems []*MenuItem `json:"menuItems"`
}

// MenuItem used for navigation
type MenuItem Item

// Product used for show product information
type Product struct {
	// ID
	ID int64
	// Name
	Name string
	// Description
	Description string
	// GroupID
	GroupID int64
	//Properties
	Properties map[string]string
}

// Group used for classify product
type Group struct {
	//ID
	ID int64
	//Name
	Name string
}
