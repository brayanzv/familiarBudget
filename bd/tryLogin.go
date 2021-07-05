package bd

import (
	"github.com/brayanzv/FamiliarBudget2/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(login string, password string) (models.Users, bool) {
	Usu, found, _ := CheckUser(login)
	if !found {
		return Usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(Usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return Usu, false
	}
	return Usu, true
}
