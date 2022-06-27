package Models

type Cars struct {
	Id           uint   `json:"id" gorm:"primarykey" db:"id"`
	Model        string `json:"model"  db:"model"`
	Registration string `json:"registration" db:"registration"`
	Mileage      int    `json:"mileage" db:"mileage"`
	Rented       bool   `json:"rented"  db:"rented"`
}
