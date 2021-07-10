package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Details struct {
	gorm.Model
	Id_user int `json:"id_user"`
	Type   string    `json:"type"`
	Amount float32   `json:"amount"`
	Detail string    `json:"detail"`
}
type GetDetailsID struct {
	Name string `json:"name"`
	Type string
	Amount float32
	Detail string
	CreatedAt time.Time
}
