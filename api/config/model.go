package config

type Cars struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Brand        string `json:"brand"`
	CarType      string `json:"car_type"`
	FuelType     string `json:"fuel_type"`
	Transmission string `json:"transmission"`
}
