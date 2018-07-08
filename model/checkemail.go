package model

type ReferrerSignupEmailRequest struct {
	Email string `required:"true" json:"Email" example:"admin@jogchat.com"`
}

type ReferrerSignupResponseSuccess struct {
	Message string `json:"message" example:"verification email sent"`
}

type ReferrerSignupResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"email cannot be empty"`
}

type ReferrerSignupResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"email already registered"`
}

type ApplicantSignupEmailRequest struct {
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
}

type ApplicantSignupResponseSuccess struct {
	Message string `json:"message" example:"verification email sent"`
}

type ApplicantSignupResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"email cannot be empty"`
}

type ApplicantSignupResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"email already registered"`
}
