package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/utils"
	"github.com/mailgun/mailgun-go"
	"fmt"
)

// Creds holds the credentials we send back
type Creds struct {
	Status      string
	AccountType string
	Email       string
	AuthToken   string
	IsLoggedIn  bool
}


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
func GetCredentials(env *Env, username string, email string) Creds {
	credentials := Creds{
		Status:      "OK",
		AccountType: "user",
		Email:       email,
		IsLoggedIn:  true,
	}
	// Now create a JWT for user
	// Create the token
	credentials.AuthToken = utils.GetToken(env.Secret, username)
	return credentials
}

func addCredentials(env *Env, ctx *gin.Context, username string, email string) {
	credentials := GetCredentials(env, username, email)
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
			addCredentials(env, ctx, info["username"], info["email"])
		}
	}
}

func Signup(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, _ := ioutil.ReadAll(ctx.Request.Body)

		var params map[string]string
		json.Unmarshal(data, &params)

		username := params["Username"]
		email := params["Email"]
		password := params["Password"]
		if username == "" || email == "" || password == "" {
			handleFailure(errors.New("username, email and password cannot be empty"), ctx)
		} else {
			token := utils.GetToken(env.Secret, username)
			successful, err := schemaless.SignupDB(username, email, password, token)
			if !successful {
				handleFailure(err, ctx)
			} else {
				addCredentials(env, ctx, username, email)
				go SendVerificationEmail(env, email, token)
			}
		}
	}
}

func VerifyEmail(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		email := ctx.Query("email")
		token := ctx.Query("token")
		rowKey, successful, err := schemaless.VerifyEmail(email, token)
		if !successful {
			handleFailure(err, ctx)
		} else {
			schemaless.ActivateEmail(rowKey)
			ctx.JSON(http.StatusOK, gin.H{"Congratulations": "You've activated your email."})
		}
	}
}

func SendVerificationEmail(env *Env, email string, token string) {
	link := fmt.Sprintf("http://%s%s/activate?email=%s&token=%s", env.IP, env.Port, email, token)
	mg := mailgun.NewMailgun(env.Domain, env.PrivateKey, env.PublicKey)
	subject := "[Jogchat] Activate your account"
	message := mg.NewMessage(env.Email, subject, "[Jogchat] Activate your account", email)
	message.SetHtml(fmt.Sprintf(
		"<html>" +
			"<body>" +
			"<h2>Welcome to Jogchat.com.</h2>" +
			"<h2>Please click on the following link to activate your account: </h2>" +
			"<h2><a href =\"%s\">link</a></h2>" +
			"</body> " +
			"</html>", link))
	_, _, err := mg.Send(message)
	utils.CheckErr(err)
}


func handleFailure(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
