package routers

import (
	Bd "TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
	"net/http"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t Models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := Bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}