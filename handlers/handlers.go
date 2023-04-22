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
    PORT := os.Getenv("PORT")

    if PORT == "" {
        PORT = "8080"
    }

    handler := cors.Default().Handler(router)

    log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
