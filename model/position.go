package model

type PostPositionRequest struct {
	Username string `required:"true" json:"Username" example:"chaoran@Uber.com"`
	Company string `required:"true" json:"Company" example:"Uber"`
	Position string `required:"true" json:"Position" example:"Software Engineer"`
	Description string `required:"true" json:"Description" example:"Build a microservice platform for Uber. A position require microservice knowledge and past experience in Golang."`
}

type PostPositionResponseSuccess struct {
	Message string `required:"true" json:"message" example:"Success: status 200"`
}

type PostPositionResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"username does not exist"`
}

type PostPositionResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"construct cell failure"`
}

