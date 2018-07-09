package model

type SigninRequest struct {
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
	Password string `required:"true" json:"Password" example:"wang374password"`
}

type SigninResponseSuccess struct {
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
	Password string `required:"true" json:"Password" example:"wang374password"`
}

type SigninResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"email not registered"`
}

type SigninResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"please verify your email"`
}

type SigninResponseAPIError2 struct {
	ErrorMessage string `json:"error" example:"invalid password"`
}