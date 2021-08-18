package models

type UserCond struct {
	Username    string
	MaxCalories int
	Diabetic    bool
	Halal       bool
	Vegan       bool
	Address     string
	PostalCode  int
}
