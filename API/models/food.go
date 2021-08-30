package models

type Food struct {
	ID          int
	Name        string
	ShopID      int
	Calories    int
	Description string
	Sugary      string
	Halal       string
	Vegan       string
}
