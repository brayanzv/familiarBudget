package jwt

import (
	"time"

	"github.com/brayanzv/FamiliarBudget2/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera el encriptado con JWT */
func GenerateJWT(t models.Users) (string, error) {

	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"id":t.Id,
		"login":     t.Login,
		"name":      t.Name,
		"last_name": t.Last_name,
		"id_role":   t.Id_role,
		"cod_familiar": t.Cod_familiar,
		"status":    t.Status,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
