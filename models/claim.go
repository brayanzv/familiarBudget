package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/*claim es la escructura usada para usar el JWT*/
type Claim struct {
	Id    uint   `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Id_role int `json:"id_role,omitempty"`
	Cod_familiar string `json:"cod_familiar,omitempty"`
	jwt.StandardClaims
}
