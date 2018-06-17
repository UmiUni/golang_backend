package main

import (
	"net/http"
	"code.jogchat.internal/golang_backend/handler"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gorilla/mux"
	"log"
)


func main() {
	// Initialise our app-wide environment data we'll send to the handler
	env := &handler.Env{
		Secret: "biscuits and gravy",
	}

	schemaless.InitDB()
	defer schemaless.CloseDB()

	r := mux.NewRouter()

	// Test this with
	//    curl -v -X POST -d "{\"username\":\"odewahn\", \"password\":\"password\"}" --header "X-Authentication: eddieTheYeti" localhost:3000/login
	r.Handle("/login", handler.Handler{env, handler.Login}).Methods("POST", "OPTIONS")
	r.Handle("/signup", handler.Handler{env, handler.Signup}).Methods("POST", "OPTIONS")

	port := "3001" // this is the gin port, but the app port is exposed at 3000
	log.Fatal(http.ListenAndServe(":"+port, r))
}
