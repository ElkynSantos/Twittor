package handlers

import (
	Middlew "TRABAJO2/Twittor/middleW"
	"TRABAJO2/Twittor/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {

    router := mux.NewRouter()

    
    router.HandleFunc("/registro", Middlew.ChequeoBD(routers.Registro)).Methods("POST")
    
    router.HandleFunc("/login", Middlew.ChequeoBD(routers.Login)).Methods("POST")
    router.HandleFunc("/verperfil", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
    router.HandleFunc("/modificarPerfil", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
    router.HandleFunc("/graboTweet", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
    PORT := os.Getenv("PORT")

    if PORT == "" {
        PORT = "8080"
    }

    handler := cors.Default().Handler(router)

    log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
