package routers

import (
	"encoding/json"
	"fmt"
	"github.com/brayanzv/FamiliarBudget2/bd"
	"github.com/brayanzv/FamiliarBudget2/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func UsersListCodFamily(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IDCodF := vars["codFamily"]

	if IDCodF == ""{
		http.Error(w, "Tiene que enviar un codigo de familia ", 400)
		return
	}
	if Id_roles!= 1 && Id_roles != 2{
		http.Error(w, "Usuario no autorizado", 401)
		return
	}

	if CodFamiliar != IDCodF && Id_roles !=1{
		http.Error(w, "Debe ingresar correctamente su codigo de familia", 401)
		return
	}
	fmt.Println(IDCodF+"handler")
	fmt.Println(CodFamiliar+"este")

	result,status,c:=bd.GetUsers()
	if status== true {
		var search []*models.UserAPI

		for i := 0; i < int(c); i++ {

			tempUserApi := models.UserAPI{}
			if result[i].Cod_familiar==IDCodF {
				tempUserApi.Id = int(result[i].Id)
				tempUserApi.Name = result[i].Name
				tempUserApi.Last_name = result[i].Last_name
				tempUserApi.Id_role = result[i].Id_role
				tempUserApi.Cod_familiar = result[i].Cod_familiar
				tempUserApi.Status = result[i].Status

				search = append(search, &tempUserApi)
			}
		}
		if search ==  nil{
			http.Error(w, "No hay familiares registrados con es codigo de familia",400)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(search)
	}
}

func UsersFamilyModify(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	codFamily := vars["codFamily"]
	idUauxi :=vars["user_id"]
	idU,_ := strconv.Atoi(idUauxi)
	codFamily=strings.TrimSpace(codFamily)

	if codFamily != CodFamiliar {
		http.Error(w, "El codigo de familia que envio no le corresponde", 400)
		return
	}

	t:= models.Users{}

	err:=json.NewDecoder(r.Body).Decode(&t)

	if err!=nil{
		http.Error(w,"no se pudo decodificar el json"+err.Error(),400)
		return
	}
	status, err :=bd.FamilyModifyUser(t,uint(idU),codFamily)
	if err != nil {
		http.Error(w, "Ocurrio un error al modificar el registro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro modificar el registro vuelva a intentarlo"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
