package controller

// Item is the base struct
type Item struct {
	// name
	Name string `json:"name,omitempty"`
	// Icon
	Icon string `json:"icon,omitempty"`
	// Weight used for order
	Weight int `json:"weight,omitempty"`
	// Path used for route
	Path string `json:"route,omitempty"`
}

// Module is the main navigator
type Module struct {
	// Item
	Item `json:",inline"`
	// IsMain is used for the Main Web
	IsMain bool `json:"isMain,omitempty"`
	// SubModules
	SubModules []SubModule `json:"subModules"`
}

// SubModule used for sub navigation
type SubModule Item

// Menu for paged scope navigation
type Menu struct {
	// Item
	Item `json:",inline"`
	// MenuItems subMenu
	MenuItems []MenuItem `json:"menuItems"`
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
