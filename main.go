package main

import (
	"code.jogchat.internal/golang_backend/handler"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
)


func main() {
	// Initialise our app-wide environment data we'll send to the handler
	env := &handler.Env{
		//IP: "localhost",
		IP: "178.128.0.108",
		Port: ":3001",
		Secret: "biscuits and gravy",
		Domain: "jogchat.com",
		Email: "admin@jogchat.com",
		PrivateKey: "key-f44fa8c4e93f293b34bffd8df6269870",
		PublicKey: "pubkey-1fd2593b1993e76a28fd0ba2420b9333",
	}

	schemaless.InitDB()
	defer schemaless.CloseDB()

	r := gin.Default()

	// Test this with
	//    curl -v -X POST -d "{\"username\":\"odewahn\", \"password\":\"password\"}" --header "X-Authentication: eddieTheYeti" localhost:3000/login
	r.POST("/login", handler.Signin(env))
	r.POST("/signup", handler.Signup(env))
	r.GET("/activate", handler.VerifyEmail(env))

	r.Run(env.Port)
}
