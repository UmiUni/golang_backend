package main

import (
	"code.jogchat.internal/golang_backend/handler"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
)


func main() {
	// Initialise our app-wide environment data we'll send to the handler
	env := &handler.Env{
		Secret: "biscuits and gravy",
	}

	schemaless.InitDB()
	defer schemaless.CloseDB()

	r := gin.Default()

	// Test this with
	//    curl -v -X POST -d "{\"username\":\"odewahn\", \"password\":\"password\"}" --header "X-Authentication: eddieTheYeti" localhost:3000/login
	r.POST("/login", handler.Signin(env))
	r.POST("/signup", handler.Signup(env))

	r.Run(":3001")
}
