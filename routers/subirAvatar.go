package routers

import (
	"TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
	"io"
	"net/http"
	"os"
	"strings"
)


func SubirAvatar (w http.ResponseWriter, r *http.Request){

	file, handler, _:= r.FormFile("avatar")

	var extension = strings.Split(handler.Filename,".")[1]
	var archivo string = "upload/avatars/"+IDUsuario+"."+extension

	f, err := os.OpenFile(archivo, os.O_WRONLY| os.O_CREATE,0666)

	if err!=nil {
		http.Error(w,"Error al subir imagen ! "+ err.Error (),http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f,file)

	if err != nil {
		http.Error (w, "Error al copiar la imagen ! "+err.Error(),http.StatusBadRequest)
		return
	}
	var usuario Models.Usuario
	var status bool

	usuario.Avatar = IDUsuario+ "." +extension

	status, err = bd.ModificoRegistro(usuario,IDUsuario)

	if err != nil || status == false{
		http.Error(w, "Error al grabar el avatar en BD ! "+err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	

}