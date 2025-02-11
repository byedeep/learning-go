package services

import (
	"api/api/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	var cars config.Cars
	dsn := "host=localhost user=postgres password=postgres dbname=cars port=3000 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connet to database!")
	}
	DB.AutoMigrate(&cars)
}

func CreateCar(c *fiber.Ctx) error {
	car := new(config.Cars)
	if err := c.BodyParser(car); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}
	DB.Create(&car)
	return c.Status(201).JSON(car)
}

func GetCars(c *fiber.Ctx) error {
	var cars []config.Cars
	DB.Find(&cars)
	return c.JSON(cars)
}

func GetCar(c *fiber.Ctx) error {
	var cars config.Cars
	id := c.Params("id")

	result := DB.First(&cars, id)

	if result.Error != nil {
		c.Status(404).JSON(fiber.Map{
			"error": "id not found",
		})
	}
	return c.JSON(cars)
}

func UpdateCar(c *fiber.Ctx) error {
	var cars config.Cars
	id := c.Params("id")

	result := DB.First(&cars, id)

	if result.Error != nil {
		c.Status(404).JSON(fiber.Map{
			"error": "id not found",
		})
	}

	var UpdateData config.Cars

	if err := c.BodyParser(&UpdateData); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": "Failed to parser body",
		})
	}

	cars.Brand = UpdateData.Brand
	cars.CarType = UpdateData.CarType
	cars.FuelType = UpdateData.FuelType
	cars.Transmission = UpdateData.Transmission
	DB.Save(&cars)
	return c.JSON(cars)
}

func DeleteCar(c *fiber.Ctx) error {
	var cars config.Cars
	id := c.Params("id")

	result := DB.First(&cars, id)
	if result.Error != nil {
		c.Status(400).JSON(fiber.Map{
			"error": "Failed to find id",
		})
	}
	DB.Delete(&cars)
	return c.SendStatus(201)
}
