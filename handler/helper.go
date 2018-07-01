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

func readParams(ctx *gin.Context) map[string]string {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var params map[string]string
	json.Unmarshal(data, &params)
	return params
}

func handleFailure(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
