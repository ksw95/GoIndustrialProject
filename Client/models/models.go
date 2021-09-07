package models

type Account struct {
	Username string
	Password string
}

type UserCond struct {
	Username    string
	LastLogin   string
	MaxCalories int
	Diabetic    string
	Halal       string
	Vegan       string
}

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

type History struct {
	ID             int
	Username       string
	FoodPurchase   string
	Price          float64
	DeliveryMode   string
	Distance       float64
	CaloriesBurned int
}

type Restaurant struct {
	ID          int
	Name        string
	Description string
	Address     string
	PostalCode  int
}

type TokenPayload struct {
	tokenString string
	payload     UserCond
}
