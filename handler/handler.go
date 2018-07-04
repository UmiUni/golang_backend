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
		"AuthToken": utils.GetToken(env.Secret, username),
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
		username := params["Username"]
		email := params["Email"]
		password := params["Password"]
		category := params["Category"]
		if username == "" || email == "" || password == "" {
			handleFailure(errors.New("username, email and password cannot be empty"), ctx)
		} else if category != "referrer" && category != "applicant" {
			handleFailure(errors.New("invalid category"), ctx)
		} else {
			token := utils.GetToken(env.Secret, username)
			info, successful, err := schemaless.SignupDB(username, email, password, category, token)
			if !successful {
				handleFailure(err, ctx)
			} else {
				addCredentials(env, ctx, info["id"], info["username"], info["email"])
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

func UploadResume(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func InsertNews(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
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
		var (
			news[]map[string]interface{}
			err error
		)

		id, exist := ctx.GetQuery("id")
		if exist {
			news, _, err = schemaless.GetNewsByField("id", id)
		} else {
			domain := ctx.Query("domain")
			news, _, err = schemaless.GetNewsByField("domain", domain)
		}

		if err != nil {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, map[string]interface{} {
				"news": news,
			})
		}
	}
}

// If replying to other comment, parent_id should be the comment it's replying to
// If commenting on news directly, parent_id should be the news_id
// CommentOn can be either comment or news
func CommentOn(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		user_id := params["UserId"]
		news_id := params["NewsId"]
		parent_id := params["ParentId"]
		content := params["Content"]
		comment_on := params["CommentOn"]
		timestamp, _ := strconv.ParseInt(params["Timestamp"], 10, 64)

		comment_id, successful, err := schemaless.CommentOn(user_id, news_id, parent_id, comment_on, content, timestamp)
		if !successful {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, map[string]string {
				"comment_id": comment_id,
			})
		}
	}
}

func GetComment(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		parent_id := ctx.Query("parent_id")

		comments, found, err := schemaless.GetComment(parent_id)
		if !found {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, map[string]interface{} {
				"comments": comments,
			})
		}
	}
}
