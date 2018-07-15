package main

import (
	"code.jogchat.internal/golang_backend/handler"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/middleware"
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files

    _ "./docs" // docs is generated by Swag CLI, you have to import it.
)


// @title ReferHelper API
// @version 1.0
// @description This is a ReferHelper API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url admin@jogchat.com
// @contact.email admin@jogchat.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 178.128.0.108:3001
// @BasePath /

func main() {
	// Initialise our app-wide environment data we'll send to the handler
	env := &handler.Env{
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

	v1 := r.Group("/v1")
	v1.Use(middleware.VerifyToken(env))
	{
		v1.POST("/upload_resume", handler.UploadResume(env))
		v1.GET("/get_resume", handler.GetResume(env))
		v1.POST("/post_job", handler.PostPosition(env))
		v1.POST("/comment_on", handler.CommentOn(env))
	}

	r.POST("/referrer_check_signup_email", handler.ReferrerCheckSignupEmail(env))
	r.POST("/applicant_check_signup_email", handler.ApplicantCheckSignupEmail(env))
	r.POST("/resend_activation_email", handler.ResendActivationEmail(env))
	r.POST("/activate_and_signup", handler.ActivateAndSignup(env))
	r.POST("/signin", handler.Signin(env))
	r.POST("/reset_request", handler.SendResetPasswordEmail(env))
	r.POST("/reset_password", handler.ResetPasswordForm(env))

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/add_company", handler.AddCompany(env))
	r.POST("/add_school", handler.AddSchool(env))
	r.GET("/get_all_companies", handler.GetAllCompanies(env))
	r.GET("/get_all_schools", handler.GetAllSchools(env))
	r.GET("/get_company", nil)
	r.GET("/get_school", nil)

	r.Run(env.Port)
}
