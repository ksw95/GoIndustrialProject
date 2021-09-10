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

type Restaurant struct {
	ID          int
	Name        string
	Description string
	Address     string
	PostalCode  int
}

type Account struct {
	Username string
	Password string
}

type History struct {
	ID             int
	Username       string
	FoodPurchase   string
	Price          float64
	DeliveryMode   string
	Distance       float64
	CaloriesBurned int
}

type UserCond struct {
	Username    string
	MaxCalories int
	Diabetic    bool
	Halal       bool
	Vegan       bool
}
