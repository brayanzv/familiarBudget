package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Detail struct {
	gorm.Model
	Type   string    `json: "type"`
	Amount float32   `json: "amount"`
	Date   time.Time `json: "date"`
	Detail string    `json: "detail"`
}
