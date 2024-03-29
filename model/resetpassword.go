package model

type SendResetPasswordEmailRequest struct {
	Email string `required:"true" json:"Email" example:"admin@umiuni.com"`
}

type SendResetPasswordEmailResponseSuccess struct {
	Message string `required:"true" json:"message" example:"reset email sent"`
}

type SendResetPasswordEmailResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"email not registered"`
}

type ResetPasswordFormRequest struct {
	Email string `required:"true" json:"Email" example:"admin@umiuni.com"`
	Password string `required:"true" json:"Password" example:"admin374newpassword"`
	token string `required:"true" json:"Token" example:""`
}

type ResetPasswordFormResponseSuccess struct {
	UserId string `required:"true" json:"UserId" example:"ce57e12a-fe27-43a2-9a1f-0792b3d36f2e"`
	Username string `required:"true" json:"Username" example:"admin374"`
	Email string `required:"true" json:"Email" example:"admin@umiuni.com"`
	AuthToken string `required:"true" json:"AuthToken" example:""`
}

type ResetPasswordFormResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"email not registered"`
}

