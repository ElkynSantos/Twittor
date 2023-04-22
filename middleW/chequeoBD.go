package middlew

import (
	Bd "TRABAJO2/Twittor/bd"
	"net/http"
)

func ChequeoBD (next http.HandlerFunc) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){

		if Bd.Checkconnection() == 0{
			http.Error(w, "Conexion Perdida con la base de datos", 500)
		} 

		next.ServeHTTP(w,r)
	}
}