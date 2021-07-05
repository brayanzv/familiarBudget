package bd

import (
	"fmt"
	"github.com/brayanzv/FamiliarBudget2/models"
)

/*ChequeoYaExisteUsuario recibe un login de parámetro y chequea si ya está en la BD */
func CheckUser(login string) (models.Users, bool, uint) {

	var user models.Users
	search := login
	var conn = ConectionBD()
	err := conn.Where("login = ?", search).Find(&user).Error
	conn.Close()
	fmt.Errorf("error al buscar",err)
	id := user.Id
	if login == user.Login && login!= "" {

		return user, true, id
	}
	return user, false, id

}
