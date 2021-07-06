package models

import (
	"github.com/jinzhu/gorm"
)

type Details struct {
	gorm.Model
	Id_user int `json:"id_user"`
	Type   string    `json:"type"`
	Amount float32   `json:"amount"`
	Detail string    `json:"detail"`







	
}
