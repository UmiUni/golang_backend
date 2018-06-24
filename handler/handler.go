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

		if params["Username"] == "" || params["Email"] == "" || params["Password"] == "" {
			handleFailure(errors.New("username, email and password cannot be empty"), ctx)
		} else {
			successful, err := schemaless.SignupDB(params["Username"], params["Email"], params["Password"])
			if !successful {
				handleFailure(err, ctx)
			} else {
				addCredentials(env, ctx, params["Username"], params["Email"])
			}
		}
	}
}

func SendEmail(env *Env, email string, token string) {
	mg := mailgun.NewMailgun(env.Domain, env.PrivateKey, env.PublicKey)
	subject := "[Jogchat] Activate your account"
	body := "<html>" +
		"<body>" +
		"<h1>'.$randomAdd.' Welcome to Jogchat.com '.$randomAdd.'</h1>" +
		"<h1>Please click on the following link to activate your account: </h1>" +
		"<h1>'.$randomAdd.'" +
		"<a href ='.$tokenLink.'>link</a>'.$randomAdd.'" +
		"</h1> </body> </html> "
	message := mg.NewMessage(env.Email, subject, body, email)
	_, _, err := mg.Send(message)
	utils.CheckErr(err)
}


func handleFailure(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
