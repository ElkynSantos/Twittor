package routers

import (
	Bd "TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
	"encoding/json"
	"net/http"
)


func ModificarPerfil (w http.ResponseWriter, r *http.Request){

	var t Models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {

		http.Error (w, "Datos Incorrectos"+err.Error(),400)
		return
	}
	var status bool
	status,err = Bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro",400)

	}
	
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

	

}