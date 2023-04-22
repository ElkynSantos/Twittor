package routers

import (
	Bd "TRABAJO2/Twittor/bd"
	"TRABAJO2/Twittor/jwt"

	Models "TRABAJO2/Twittor/models"
	"encoding/json"
	"net/http"
	"time"
)

func Login (w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type"," application/json")

	var t Models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil{
		http.Error(w,"Usuario y/o Contrasena invalido"+ err.Error(),400)
		return
	}



	if len(t.Email) == 0 {
		http.Error(w,"El email del usuario requerido"+ err.Error(),400)
		return
	}

	documento, existe := Bd.IntentoLogin(t.Email, t.Password)

	if existe == false {

		http.Error(w,"Usuario y/o Contrasena invalido",400)
	
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w,"Ocurrio un error en el token"+ err.Error(),400)
	
		return
	}
	
	resp := Models.RespuestaLogin{
		Token:jwtKey,

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime:= time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}