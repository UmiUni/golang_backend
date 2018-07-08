package model

// @Example "{"Email":"wang374@uiuc.edu", "Username":"wang374", "Password":"wang374password", "Token":""}"
type ActivateAndSignupRequest struct {
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
	Username string `required:"true" json:"Username" example:"wang374"`
	Password string `required:"true" json:"Password" example:"wang374password"`
	Token string `required:"true" json:"Token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEyOTE4MTgsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoic3VwZXJjaGFvcmFuQGdtYWlsLmNvbSJ9.yYQOAIyHQJJUntMGtRaAm2bXF-HvvsK6vHjhe0SDsHg"`
}

// return would be a credential
type ActivateAndSignupResponseSuccess struct {
	UserId string `required:"true" json:"UserId" example:""`
	Username string `required:"true" json:"Username" example:"wang374"`
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
	AuthToken string `required:"true" json:"AuthToken" example:""`
}

type ActivateAndSignupResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"username already in use"`
}

type ActivateAndSignupResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"invalid token"`
}

type ActivateAndSignupResponseAPIError2 struct {
	ErrorMessage string `json:"error" example:"email already activated"`
}