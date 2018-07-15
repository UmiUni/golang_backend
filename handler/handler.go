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

// @Title ReferrerCheckSignupEmail
// @Summary ReferrerCheckSignupEmail
// @Description Onboarding user will provide a company email to sign up for the referral portal, if the email does not exists in schemaless database, we will send the email an activation link
// @Accept  json
// @Param body body model.ReferrerSignupEmailRequest true "ReferrerSignupEmailRequest is a POST JSON type"
// @Example "{"Email":"chaoran@uber.com"}"
// @Success 200 {object} model.ReferrerSignupResponseSuccess "Success: verification email sent"
// @Failure 400 {object} model.ReferrerSignupResponseAPIError0 "email cannot be empty"
// @Failure 400 {object} model.ReferrerSignupResponseAPIError1 "email already registered"
// @Router /referrer_check_signup_email [post]
func ReferrerCheckSignupEmail(env *Env) func(ctx *gin.Context) {
	return CheckSignupEmail(env, "referrer")
}

// @Title ApplicantCheckSignupEmail
// @Summary ApplicantCheckSignupEmail
// @Description Provide a school/university edu email to sign up for the applicant portal, if the email does not exists in schemaless database, we will send the email an activation link
// @Accept  json
// @Param body body model.ApplicantSignupEmailRequest true "ApplicantSignupEmailRequest is a POST JSON type"
// @Example "{"Email":"wang374@uiuc.edu"}"
// @Success 200 {object} model.ApplicantSignupResponseSuccess "Success: verification email sent"
// @Failure 400 {object} model.ApplicantSignupResponseAPIError0 "email cannot be empty"
// @Failure 400 {object} model.ApplicantSignupResponseAPIError1 "email already registered"
// @Router /applicant_check_signup_email [post]
func ApplicantCheckSignupEmail(env *Env) func(ctx *gin.Context) {
	return CheckSignupEmail(env, "applicant")
}

func CheckSignupEmail(env *Env, category string) func(ctx *gin.Context) {
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

func ResendActivationEmail(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		email := params["Email"]
		token := utils.GetToken(env.Secret, email)
		if email == "" {
			handleFailure(errors.New("email cannot be empty"), ctx)
			return
		}
		successful, err := schemaless.ReverifyEmail(email, token)
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

// @Title ActivateAndSignup
// @Summary ActivateAndSignup
// @Description When user click on the GET link in user email, it will hit a frontend page as a GET request with {Email, Token} as parameters. The frontend page should then provide user with a form that ask for (Email(prefilled), Username, password, token(prefilled and hidden)). Once frontend gather all infos from the user, frontend should POST call this [ActivateAndSignup endpoint] with a post request that has {email, username, password, token} as JSON to sign the user up. This endpoint will both signup the user and activate their account.
// @Accept  json
// @Param body body model.ActivateAndSignupRequest true "ActivateAndSignupRequest is a POST JSON type"
// @Success 200 {object} model.ActivateAndSignupResponseSuccess "Success: verification email sent"
// @Failure 400 {object} model.ActivateAndSignupResponseAPIError0 "username already in use"
// @Failure 400 {object} model.ActivateAndSignupResponseAPIError1 "invalid token"
// @Failure 400 {object} model.ActivateAndSignupResponseAPIError2 "email already activated"
// @Router /activate_and_signup [post]
func ActivateAndSignup(env *Env) func(ctx *gin.Context) {
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

// @Title Signin
// @Summary Signin
// @Description After user click on sign-in button, front-end will call this endpoint with a JSON wrapped {Email and Password}, the end point will then return an AuthToken on success. Front-end should store the authtoken for user either in session or cookie for user. To access password protect resource later, front-end needs to pass (username+AuthToken) to backend to verify user identity. This is called JWT Auth flow.
// @Accept json
// @Param body body model.SigninRequest true "SigninRequest is a POST JSON type"
// @Success 200 {object} model.SigninResponseSuccess "Success: sign in request succeed"
// @Failure 400 {object} model.SigninResponseAPIError0 "email not registered"
// @Failure 400 {object} model.SigninResponseAPIError1 "please verify your email"
// @Failure 400 {object} model.SigninResponseAPIError2 "invalid password"
// @Router /signin [post]
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

// @Title SendResetPasswordEmail
// @Summary SendResetPasswordEmail
// @Description When user click on reset password button with an email filled in a form above, front-end will call this endpoint with a JSON wrapped {Email} to sent reset password email, a hacker can hack this end point currently by repeatedly calling and our system will spam send email. Next step would be requiring a session {Email, AuthToken} combination and this endpoint will only be able to sent email to this session's Email.
// @Accept json
// @Param body body model.SendResetPasswordEmailRequest true "ResetPasswordButtonRequest is a POST JSON type"
// @Success 200 {object} model.SendResetPasswordEmailResponseSuccess "Success: message: reset email sent"
// @Failure 400 {object} model.SendResetPasswordEmailResponseAPIError0 "Failure: email not registered"
// @Router /send_reset_password_email [post]
func SendResetPasswordEmail(env *Env) func(ctx *gin.Context) {
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

// @Title ResetPasswordForm
// @Summary ResetPasswordForm
// @Description After user clicks on reset password link(GET with email and token) in email, front-end/mobile will provide user with a form, {Email(prefilled), Password, Token(prefilled)}. After user filled the form, front-end/mobile will call this endpoint with a JSON wrapped {Email(prefilled), Password, Token(prefilled)} POST to reset password. If the user is not activated at the point of click on reset_password, an email titled reset_password with activation instruction will be sent.
// @Accept json
// @Param body body model.ResetPasswordFormRequest true "ResetPasswordButtonRequest is a POST JSON type"
// @Success 200 {object} model.ResetPasswordFormResponseSuccess "Success: message: reset email sent"
// @Failure 400 {object} model.ResetPasswordFormResponseAPIError0 "Failure: email not registered"
// @Router /reset_password_form [post]
func ResetPasswordForm(env *Env) func(ctx *gin.Context) {
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
		name := ctx.PostForm("Name")
		domain := ctx.PostForm("Domain")
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

func PostPosition(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		username := params["Username"]
		company := params["Company"]
		position := params["Position"]
		description := params["Description"]
		info, successful, err := schemaless.PostPosition(username, company, position, description)
		if !successful {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, info)
		}
	}
}

func CommentOn(env *Env) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := readParams(ctx)
		username := params["Username"]
		positionId := params["PositionId"]
		parentId := params["ParentId"]
		parentType := params["ParentType"]
		content := params["Content"]
		if parentType != "position" && parentType != "comment" {
			handleFailure(errors.New("invalid parent type"), ctx)
			return
		}
		info, successful, err := schemaless.CommentOn(username, positionId, parentId, parentType, content)
		if !successful {
			handleFailure(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, info)
		}
	}
}
