package routers

import (
	"encoding/json"
	"github.com/brayanzv/FamiliarBudget2/bd"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/brayanzv/FamiliarBudget2/models"
)

//valida la informacion enviada en el body y sus credenciales
func DetailRegister(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	idUs:=vars["user_id"]
	idU,_:=strconv.Atoi(idUs)

	if uint(idU)!=IdUsers{
		http.Error(w, "id invalido",400)
		return
	}

	var t models.Details

	err := json.NewDecoder(r.Body).Decode(&t)
	t.Id_user=idU
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	status, err :=bd.InsertDetail(t)
	if status==false{
		return
		http.Error(w,"No se pudo registrar el Detalle ", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
