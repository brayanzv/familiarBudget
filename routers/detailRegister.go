package routers

import (
	"encoding/json"
	"net/http"

	"github.com/brayanzv/FamiliarBudget2/models"
)

func DetailRegister(w http.ResponseWriter, r *http.Request) {
	var t models.Detail

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

}
