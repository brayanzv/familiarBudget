package bd

import (
	"fmt"
	"github.com/brayanzv/FamiliarBudget2/models"
)

func CodFamiliar(codf string) bool {

	var user models.Users

	search := codf
	var conn = ConectionBD()
	if codf != "" {
		err := conn.Where("cod_familiar = ?", search).Where("id_role = ?", 2).Find(&user).Error
		conn.Close()
		fmt.Println(err)
		Names := user.Cod_familiar
		if codf == Names {
			return true
		}
	}
	conn.Close()
	return false
}
