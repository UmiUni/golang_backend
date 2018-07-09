package handler

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"code.jogchat.internal/golang_backend/utils"
	"io/ioutil"
	"encoding/json"
)


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

func sendVerificationEmail(env *Env, email string, token string) {
	subject := "[Jogchat] Activate your account"
	text := subject
	body := fmt.Sprintf(
		"<html>" +
		"<body>" +
		"<h2>Welcome to ReferHelper.com.</h2>" +
		"<h2>This is your activation link:</h2>" +
			"<a href=\"https://referhelper.com/signup/email?=%stoken?=%s\">activate account</a>"+
		"</body> " +
		"</html>", email, token)
	sendEmail(env, email, subject, text, body)
}

func sendResetPasswordEmail(env *Env, email string, token string)  {
	subject := "[Jogchat] Reset password"
	text := subject
	body := fmt.Sprintf(
		"<html>" +
			"<body>" +
			"<h2>Welcome to ReferHelper.com.</h2>" +
			"<h2>This is your reset password link:</h2>" +
			"<a href=\"https://referhelper.com/reset/email?=%stoken?=%s\">reset password</a>"+
			"</body> " +
			"</html>", email, token)
	sendEmail(env, email, subject, text, body)
}

func sendEmail(env *Env, email string, subject string, text string, body string) {
	mg := mailgun.NewMailgun(env.Domain, env.PrivateKey, env.PublicKey)
	message := mg.NewMessage(env.Email, subject, text, email)
	message.SetHtml(body)
	_, _, err := mg.Send(message)
	utils.CheckErr(err)
}

func readParams(ctx *gin.Context) map[string]string {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var params map[string]string
	json.Unmarshal(data, &params)
	return params
}

func handleFailure(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

// TODO: change path to S3 bucket
func resumePath(username string, filename string) string {
	return "tmp/" + username + "_" + filename
}
