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

func sendVerificationEmail(env *Env, email string, token string) {
	subject := "[Jogchat] Activate your account"
	text := subject
	body := fmt.Sprintf(
		"<html>" +
		"<body>" +
		"<h2>Welcome to Jogchat.com.</h2>" +
		"<h2>This is your verification token: %s</h2>" +
		"</body> " +
		"</html>", token)
	sendEmail(env, email, subject, text, body)
}

func sendResetPasswordEmail(env *Env, email string, token string)  {
	subject := "[Jogchat] Reset password"
	text := subject
	body := fmt.Sprintf(
		"<html>" +
			"<body>" +
			"<h2>Welcome to Jogchat.com.</h2>" +
			"<h2>This is your verification token: %s</h2>" +
			"</body> " +
			"</html>", token)
	sendEmail(env, email, subject, text, body)
}

func sendEmail(env *Env, email string, subject string, text string, body string)  {
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
