package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/utils"
	"strconv"
	"github.com/satori/go.uuid"
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
				go sendVerificationEmail(env, email, token)
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
			err := schemaless.ActivateEmail(rowKey)
			if err != nil {
				handleFailure(err, ctx)
			} else {
				ctx.JSON(http.StatusOK, gin.H{"Congratulations": "You've activated your email."})
			}
		}
	}
}

func InsertNews(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, _ := ioutil.ReadAll(ctx.Request.Body)

		var params map[string]string
		json.Unmarshal(data, &params)

		domain := params["Domain"]
		timestamp, _ := strconv.ParseInt(params["Timestamp"], 10, 64)
		author := params["Author"]
		summary := params["Summary"]
		title := params["Title"]
		text := params["Text"]
		url := params["URL"]

		successful, err := schemaless.InsertNews(domain, timestamp, author, summary, title, text, url)
		if !successful {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"Congratulations": "News successfully added."})
		}
	}
}

func GetNews(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var news[]map[string]interface{}
		var err error
		var id uuid.UUID

		id_str, exist := ctx.GetQuery("id")
		if exist {
			id, err = uuid.FromString(id_str)
			if err == nil {
				news, _, err = schemaless.GetNewsByField("id", id.Bytes())
			}
		} else {
			domain := ctx.Query("domain")
			news, _, err = schemaless.GetNewsByField("domain", domain)
		}

		if err != nil {
			handleFailure(err, ctx)
		} else {
			results := make(map[string]interface{})
			results["news"] = news
			ctx.JSON(http.StatusOK, results)
		}
	}
}

func CommentOn(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func GetComment(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
