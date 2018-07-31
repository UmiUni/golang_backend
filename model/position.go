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

type CommentOnRequest struct {
	Username string `required:"true" json:"Username" example:"wang374"`
	PositionId string `required:"true" json:"PositionId" example:"67bebc0c-f0bd-4352-b588-08a056085e0a"]`
	ParentId string `required:"true" json:"ParentId" example:"ee3bc72c-9804-498b-be9d-c77471374e0a"`
	ParentType string `required:"true" json:"ParentType" example:""`
    Content string `required:"true" json:"Content" example:""`
}

type CommentOnResponseSuccess struct {
	Message string `required:"true" json:"message" example:"Success on commenting: status 200"`
}

type CommentOnResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"invalid parent type"`
}

type CommentOnResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"username does not exist"`
}

type CommentOnResponseAPIError2 struct {
	ErrorMessage string `json:"error" example:"invalid position id"`
}

type CommentOnResponseAPIError3 struct {
	ErrorMessage string `json:"error" example:"invalid parent id"`
}