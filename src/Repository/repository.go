package database

import (
	model "github.com/Noexperience-Team/carsrent/src/Models"
	"github.com/jinzhu/gorm"
)

type Db struct {
	DB *gorm.DB
}

func (config *Db) AddCar(req *model.Car) *gorm.DB {
	return config.DB.Model(&model.Car{}).Create(&req)
}
func (config *Db) GetCars(cars *[]model.Car) *gorm.DB {
	return config.DB.Find(&cars)
}
func (config *Db) UpdateCar(car *model.Car) *gorm.DB {
	if car.Rented == false {
		return config.DB.Model(&model.Car{}).Where("registration=?", car.Registration).Updates(map[string]interface{}{"mileage": car.Mileage, "rented": false})
	} else {
		return config.DB.Model(&model.Car{}).Where("registration=?", car.Registration).Updates(map[string]interface{}{"rented": true})
	}

}
func (config *Db) FindCar(car *model.Car) *gorm.DB {
	return config.DB.Model(&model.Car{}).Where("registration=?", car.Registration).First(&car)
}
