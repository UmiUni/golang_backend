package main

import (
	"code.jogchat.internal/golang_backend/handler"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/middleware"
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
	r.Use(middleware.CORSMiddleware())

	r.POST("/signup", handler.Signup(env))
	r.POST("/activate", handler.ActivateEmail(env))
	r.POST("/login", handler.Signin(env))

	r.POST("/reset_request", handler.ResetRequest(env))
	r.POST("/reset_password", handler.ResetPassword(env))

	r.Run(env.Port)
}
