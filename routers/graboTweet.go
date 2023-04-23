package routers

import (
	"encoding/json"
	"net/http"
	"time"

	Bd "TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
)

func GraboTweet (w http.ResponseWriter, r *http.Request){
	var mensaje Models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := Models.GraboTweet{
		UserID: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := Bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro"+err.Error(),400)
		return
	}

	if status ==false{
		http.Error(w, "No se ha logrado a insertar",400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}