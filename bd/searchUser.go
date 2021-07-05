package bd

import (
	"strconv"

	"github.com/brayanzv/FamiliarBudget2/models"
)

func SearchUser(Id string) (models.Users, error) {
	search, _ := strconv.Atoi(Id)

	users := models.Users{}
	var conn = ConectionBD()
	err:=conn.Where("id = ?", search).Find(&users).Error
	conn.Close()
	users.Password = ""
	if search != int(users.Id) || users.Id==0{
		return users, err
	}
	return users, nil
}
