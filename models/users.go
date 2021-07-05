package models

import "github.com/jinzhu/gorm"



type Users struct {
	gorm.Model
	Name         string `json:"name,omitempty"`
	Last_name    string `json:"last_name,omitempty"`
	Login        string `json:"login,omitempty"`
	Password     string `json:"password,omitempty"`
	Id_role      int    `json:"id_role,omitempty"`
	Cod_familiar string `json:"cod_familiar,omitempty"`
	Status       bool   `json:"status,omitempty"`
}
type UserAPI struct {
	Id int
	Name string
	Last_name string
	Id_role int
	Cod_familiar string
	Status bool
}