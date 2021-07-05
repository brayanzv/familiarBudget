package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brayanzv/FamiliarBudget2/bd"
	"github.com/brayanzv/FamiliarBudget2/jwt"
	"github.com/brayanzv/FamiliarBudget2/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "aplication/json")

	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Login) == 0 {
		http.Error(w, "Debe ingresar su login", 400)
		return
	}

	doc, exist := bd.TryLogin(t.Login, t.Password)

	if doc.Status == false && t.Cod_familiar != doc.Cod_familiar {
		http.Error(w, "Espere a que habiliten su cuenta esto tomara unos minutos", 400)
		return
	}
	if !exist {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correxpondiente"+err.Error(), 400)
		return
	}
	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
