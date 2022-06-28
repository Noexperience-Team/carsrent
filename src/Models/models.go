package Models

type Car struct {
	Id           uint   `json:"id" gorm:"primarykey" db:"id"`
	Model        string `json:"model"  db:"model"`
	Registration string `json:"registration" db:"registration"`
	Mileage      int    `json:"mileage" db:"mileage"`
	Rented       bool   `json:"rented"  db:"rented"`
}

// type Rent struct {
// 	Id     uint `json:"id" gorm:"primarykey" db:"id"`
// 	Rented bool `json:"rented"  db:"rented"`
// 	CarId  uint `json:"car_id"  gorm:"car_id" db:"car_id" `
// 	Car    Car  `gorm:"foreignKey:car_id;"`
// }
type Response struct {
	Error string
	Data  interface{}
}
