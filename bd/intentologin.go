package bd

import (
	Models "TRABAJO2/Twittor/models"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin (email string, password string) (Models.Usuario, bool){

	usu, encontrado, _:= ChequeoYaExisteUsuario(email)

	if encontrado == false {
		return usu,false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {

		return usu,false
	}

	return usu,true
}