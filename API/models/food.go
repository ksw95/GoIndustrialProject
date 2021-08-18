package models

type Food struct {
	Name        string
	ShopID      int
	Calories    int
	Sugary      bool
	Description string
	Halal       bool
	Vegan       bool
}
