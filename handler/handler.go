package handler

import (
	"net/http"
	"errors"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/utils"
	//model "code.jogchat.internal/golang_backend/model"
)


// Env holds application-wide configuration.
type Env struct {
	IP	string
	Port	string
	Secret	string
	Domain	string
	Email	string
	PrivateKey	string
	PublicKey	string
}

// @Title Signup
// @Summary Signup
// @Description Signup
// @Accept  json
// @Param body body model.ReferrerSignupEmailStruct true "Body JSON"
// @Example "{"email":"superchaoran@gmail.com"}"
// @Success 200 {object} model.ReferrerSignupSuccessStruct "Success: verification email sent"
// @Failure 400 {object} model.APIError0 "email cannot be empty"
// @Failure 400 {object} model.APIError1 "email already registered"
// @Router /referrer_signup [post]
func ReferrerSignup(env *Env) func(ctx *gin.Context) {
	return Signup(env, "referrer")
}

func ApplicantSignup(env *Env) func(ctx *gin.Context) {
	return Signup(env, "applicant")
}

func Signup(env *Env, category string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		token := utils.GetToken(env.Secret, email)
		if email == "" {
			handleFailure(errors.New("email cannot be empty"), ctx)
		} else {
			successful, err := schemaless.SignupDB(category, email, token)
			if !successful {
				handleFailure(err, ctx)
			} else {
				ctx.JSON(http.StatusOK, map[string]interface{} {
					"message": "verification email sent",
				})
				go sendVerificationEmail(env, email, token)
			}
		}
	}
}

func ActivateEmail(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		username := params["Username"]
		password := params["Password"]
		token := params["Token"]
		info, successful, err := schemaless.ActivateEmail(email, username, password, token)
		if !successful {
			handleFailure(err, ctx)
		} else {
			addCredentials(env, ctx, info["id"], info["username"], info["email"])
		}
	}
}

// Login captures the data posted to the /login route
func Signin(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		password := params["Password"]

		info, successful, err := schemaless.SigninDB(email, password)
		if !successful {
			handleFailure(err, ctx)
		} else {
			addCredentials(env, ctx, info["id"], info["username"], info["email"])
		}
	}
}

// Used by frontend to send a request to reset password
// send email to user with token
func ResetRequest(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		token := utils.GetToken(env.Secret, email)
		found, err := schemaless.ResetRequest(email, token)
		if !found {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, map[string]interface{} {
				"message": "reset email sent",
			})
			go sendResetPasswordEmail(env, email, token)
		}
	}
}

// Used by frontend after user click the password reset link in email
func ResetPassword(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		password := params["Password"]
		token := params["Token"]
		info, successful, err := schemaless.ResetPassword(email, password, token)
		if !successful {
			handleFailure(err, ctx)
		} else {
			addCredentials(env, ctx, info["id"], info["username"], info["email"])
		}
	}
}

func AddCompany(env *Env) func(ctx *gin.Context) {
	return AddCompanySchool(env, "companies")
}

func AddSchool(env *Env) func(ctx *gin.Context) {
	return AddCompanySchool(env, "schools")
}

func AddCompanySchool(env *Env, category string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		name := params["Name"]
		domain := params["Domain"]
		successful, err := schemaless.AddCompanySchool(category, name, domain)
		if !successful {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, map[string]interface{} {
				"message": category + " added",
			})
		}
	}
}
