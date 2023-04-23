package routers

import (
	Bd "TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
	"net/http"
)

func AltaRelacion (w http.ResponseWriter, r* http.Request){
	ID := r.URL.Query().Get("id")

	if len(ID)<1{
		http.Error(w,"El parametro ID es obligatorio",http.StatusBadRequest)
		return
	}

	var t Models.Relacion

	t.UsuarioID=IDUsuario
	t.UsuarioRelacionID=ID

	status,err := Bd.InsertoRelacion(t)

	if err!= nil{
		http.Error(w, "Ocurrio un error al insertar la relacion"+err.Error(),http.StatusBadRequest)
		return
	}

	if status ==false {
		http.Error(w,"No se ha logrado insertar la relacion",http.StatusBadRequest)

	}

	w.WriteHeader(http.StatusCreated)

}