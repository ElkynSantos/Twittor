package routers

import (
	Bd "TRABAJO2/Twittor/bd"
	Models "TRABAJO2/Twittor/models"
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ProcesoToken (tk string) (*Models.Claim, bool, string,error){

	miClave := []byte("HALAMADRID")
	claims := &Models.Claim{}
	
	splitToken :=strings.Split(tk, "Bearer")

	if len(splitToken) !=2 {
		return claims,false,string(""),errors.New("Formato de Token invalido")

	}

	tk = strings.TrimSpace(splitToken[1])
	
	tkn, err := jwt.ParseWithClaims(tk,claims,func(token *jwt.Token)(interface{},error){

	return miClave, nil
	})

	if err == nil {
		_, encontrado, _:= Bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario= claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil

	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims,false,string(""),err

}
