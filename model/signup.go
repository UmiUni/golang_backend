package model

// @Example "{"Email":"wang374@uiuc.edu", "Username":"wang374", "Password":"wang374password", "Token":""}"
type ActivateAndSignupRequest struct {
	Email string `required:"true" json:"Email" example:"admin@jogchat.com"`
	Username string `required:"true" json:"Username" example:"admin374"`
	Password string `required:"true" json:"Password" example:"admin374password"`
	Token string `required:"true" json:"Token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEzNTIyMDAsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoid2FuZzM3NEB1aXVjLmVkdSJ9.gC7dTl64XDe5BwlS8PuZxBxGes1ujcCWFbe23r0xOXM"`
}

// return would be a credential
type ActivateAndSignupResponseSuccess struct {
	Username string `required:"true" json:"Username" example:"admin374"`
	UserId string `required:"true" json:"UserId" example:"ce57e12a-fe27-43a2-9a1f-0792b3d36f2e"`
	Email string `required:"true" json:"Email" example:"admin@jogchat.com"`
	AuthToken string `required:"true" json:"AuthToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEzNTM1NjUsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoid2FuZzM3NEB1aXVjLmVkdSJ9.XwmDhW1b99E9jwGatN_6y1tYpLGBcAqywS9fI23Oxxo"`
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