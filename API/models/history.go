package models

type History struct {
	ID             int
	Username       string
	FoodPurchase   string
	Price          float64
	DeliveryMode   string
	Distance       float64
	CaloriesBurned int
}
