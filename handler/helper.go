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
			"<a href=\"https://referhelper.com/signup?email=%s&token=%s\">activate account</a>"+
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
			"<a href=\"https://referhelper.com/reset?email=%s&token=%s\">reset password</a>"+
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
	return "tmp/resume" + username + "_" + filename
}

//func getIcons(domain string) (err error) {
//	// The session the S3 Downloader will use
//	sess := session.Must(session.NewSession())
//
//	// Create a downloader with the session and default options
//	downloader := s3manager.NewDownloader(sess)
//	getS3(downloader, domain + "/1.0.png")
//	return nil
//}
//
//func getS3(downloader *s3manager.Downloader, filename string) (content []byte, err error) {
//	// Create a file to write the S3 Object contents to.
//	f, err := os.Create(filename)
//	if err != nil {
//		return fmt.Errorf("failed to create file %q, %v", filename, err)
//	}
//
//	// Write the contents of S3 Object to the file
//	n, err := downloader.Download(f, &s3.GetObjectInput{
//		Bucket: aws.String("<YOUR_BUCKET_NAME>"),
//		Key:    aws.String("<YOUR_ITEM_NAME>"),
//	})
//	_, err = f.Read(content)
//	if err != nil {
//		return fmt.Errorf("failed to download file, %v", err)
//	}
//	fmt.Printf("file downloaded, %d bytes\n", n)
//}
