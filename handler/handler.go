package handler

import (
	"net/http"
	"errors"
	"code.jogchat.internal/golang_backend/schemaless"
	"github.com/gin-gonic/gin"
	"code.jogchat.internal/golang_backend/utils"
	"os"
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
			return
		}
		successful, err := schemaless.SignupDB(category, email, token)
		if !successful {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "verification email sent",
			})
			go sendVerificationEmail(env, email, token)
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
			ctx.JSON(http.StatusOK, gin.H{
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

func UploadResume(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.PostForm("Username")
		if username == "" {
			handleFailure(errors.New("invalid username"), ctx)
			return
		}
		file, err := ctx.FormFile("Resume")
		if err != nil {
			handleFailure(err, ctx)
			return
		}
		filename := resumePath(username, file.Filename)
		err = ctx.SaveUploadedFile(file, filename)
		if err != nil {
			handleFailure(err, ctx)
			return
		}
		sucessful, err := schemaless.UploadResume(username, filename)
		if !sucessful {
			os.Remove(filename)
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "resume uploaded",
			})
		}
	}
}

func GetResume(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.Query("Username")
		filename, found, err := schemaless.GetResume(username)
		if !found {
			handleFailure(err, ctx)
		} else {
			ctx.Header("Content-Description", "File Transfer")
			ctx.Header("Content-Transfer-Encoding", "binary")
			ctx.Header("Content-Disposition", "attachment; filename=" + filename)
			ctx.File(filename)
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
			ctx.JSON(http.StatusOK, gin.H{
				"message": category + " added",
			})
		}
	}
}
