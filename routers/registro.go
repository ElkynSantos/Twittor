package routers

import (
	Bd "TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
	"encoding/json"
	"net/http"
)

func Registro (w http.ResponseWriter, r *http.Request){
	var t Models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err !=nil{
		http.Error(w,"Error en los datos recibidos +"+err.Error(),400)
		return
	}

	if len(t.Email)==0{
		http.Error(w,"El email del usuario es requerido",400)
		return

	}

	if len (t.Password)<6{
		http.Error(w,"debe de ser de 6 caracteres",400)
	}


	_,encontrado,_:=Bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado==true{
		http.Error(w,"Ya existe un usuario con ese email",400)
		return
	}

	_, status, err := Bd.InsertoRegistro(t)

	if err != nil {
		
		http.Error(w,"Ocurrio un error al intentar el registro "+ err.Error(),400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}