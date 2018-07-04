package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/utils"
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

// GetCredentials determines if the username and password is valid
// This is where logic would go to validate and return account info
func addCredentials(env *Env, ctx *gin.Context, id string, username string, email string) {
	credentials := map[string]string {
		"UserId": id,
		"Username": username,
		"Email": email,
		"AuthToken": utils.GetToken(env.Secret, email),
	}
	ctx.JSON(http.StatusOK, credentials)
}

// Login captures the data posted to the /login route
func Signin(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, _ := ioutil.ReadAll(ctx.Request.Body) // Read the body of the POST request
		// Unmarshall this into a map
		var params map[string]string
		json.Unmarshal(data, &params)

		info, successful, err := schemaless.SigninDB(params["Email"], params["Password"])
		if !successful {
			handleFailure(err, ctx)
		} else {
			addCredentials(env, ctx, info["id"], info["username"], info["email"])
		}
	}
}

func Signup(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		if email == "" {
			handleFailure(errors.New("email cannot be empty"), ctx)
		} else {
			token := utils.GetToken(env.Secret, email)
			successful, err := schemaless.SignupDB(email, token)
			if !successful {
				handleFailure(err, ctx)
			} else {
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
		category := params["Category"]
		token := params["Token"]
		info, successful, err := schemaless.ActivateEmail(email, username, password, category, token)
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
		token, found, err := schemaless.ResetRequest(email)
		if !found {
			handleFailure(err, ctx)
		} else {
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
