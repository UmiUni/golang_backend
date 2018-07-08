package model

type ReferrerSignupEmailStruct struct {
	Email string `required:"true" json:"Email" example:"admin@jogchat.com"`
}

type ReferrerSignupSuccessStruct struct {
	Message string `json:"message" example:"verification email sent"`
}

type ReferrerSignupAPIError0 struct {
	ErrorMessage string `json:"error" example:"email cannot be empty"`
}

type ReferrerSignupAPIError1 struct {
	ErrorMessage string `json:"error" example:"email already registered"`
}

type ApplicantSignupEmailStruct struct {
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
}

type ApplicantSignupSuccessStruct struct {
	Message string `json:"message" example:"verification email sent"`
}

type ApplicantSignupAPIError0 struct {
	ErrorMessage string `json:"error" example:"email cannot be empty"`
}

type ApplicantSignupAPIError1 struct {
	ErrorMessage string `json:"error" example:"email already registered"`
}
