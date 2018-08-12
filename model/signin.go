package model

type SigninRequest struct {
	Email string `required:"true" json:"Email" example:"admin@umiuni.com"`
	Password string `required:"true" json:"Password" example:"admin374password"`
}

type SigninResponseSuccess struct {
	UserId string `required:"true" json:"UserId" example:"ce57e12a-fe27-43a2-9a1f-0792b3d36f2e"`
	Username string `required:"true" json:"Username" example:"admin374"`
	Email string `required:"true" json:"Email" example:"admin@umiuni.com"`
	AuthToken string `required:"true" json:"AuthToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEzNTM2MzgsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoid2FuZzM3NEB1aXVjLmVkdSJ9.RhRUpHJbIfid1hiJOTtStuxc86v0isnWny85COG9Mek"`
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