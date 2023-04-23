package routers

import (
	Bd "TRABAJO2/Twittor/bd"

	Models "TRABAJO2/Twittor/models"
	"encoding/json"
	"net/http"
)

func ConsultoRelacion (w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	var t Models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp Models.RespuestaConsultaRelacion

	status, err := Bd.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}