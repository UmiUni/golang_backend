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

	r.POST("/login", handler.Signin(env))
	r.POST("/signup", handler.Signup(env))
	r.GET("/activate", handler.VerifyEmail(env))

	r.POST("/insert_news", handler.InsertNews(env))
	r.GET("/get_news", handler.GetNews(env))

	r.POST("/comment_on", handler.CommentOn(env))
	r.GET("/get_comment", handler.GetComment(env))

	r.Run(env.Port)
}
