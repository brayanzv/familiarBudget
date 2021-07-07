package routers

import (
	"errors"
	"strings"

	"github.com/brayanzv/FamiliarBudget2/bd"
	"github.com/brayanzv/FamiliarBudget2/models"
	jwt "github.com/dgrijalva/jwt-go"
)
var Id_roles int
var IdUsers uint
var CodFamiliar string

func TokenProcess(tk string) (*models.Claim, bool, uint, error) {
	keyCode := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, 0, errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return keyCode, nil
	})

	//bla bla bkabkabashf sfh
	if err == nil {

		_, found, _ := bd.CheckUser(claims.Login)
		if found == true {
			CodFamiliar = claims.Cod_familiar

			Id_roles= claims.Id_role
			IdUsers = claims.Id
		}

		return claims, found,0 , nil
	}
	if !tkn.Valid {
		return claims, false, 0, errors.New("token invalido")
	}
	return claims, false, 0, err
}
