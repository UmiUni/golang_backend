package model

import "time"

type PostPositionRequest struct {
	Username string `required:"true" json:"Username" example:"admin374"`
	Company string `required:"true" json:"Company" example:"Jogchat"`
	Position string `required:"true" json:"Position" example:"Software Engineer"`
	Description string `required:"true" json:"Description" example:"Build a microservice platform for Jogchat. A position requires microservice knowledge and past experience in Golang."`
}

type PostPositionResponseSuccess struct {
	Message string `required:"true" json:"message" example:"Success: status 200 with position id {"id":"1528edfd-2cbd-451f-9053-a89e2e806cbe"} "`
}

type PostPositionResponseAPIError0 struct {
	ErrorMessage string `json:"error" example:"username does not exist"`
}

type PostPositionResponseAPIError1 struct {
	ErrorMessage string `json:"error" example:"construct cell failure"`
}

type CommentOnRequest struct {
	Username string `required:"true" json:"Username" example:"admin374"`
	PositionId string `required:"true" json:"PositionId" example:"67bebc0c-f0bd-4352-b588-08a056085e0a"]`
	ParentId string `required:"true" json:"ParentId" example:"67bebc0c-f0bd-4352-b588-08a056085e0a"`
	ParentType string `required:"true" json:"ParentType" example:"position"`
    Content string `required:"true" json:"Content" example:"这个Position很适合我背景，请联系superchaoran@gmail.com"`
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

type GetPositionsRequest struct {
	Companies string `required:"true" json:"Companies" example:"Jogchat,Google"`
	Duration string `required:"true" json:"Duration" example:"day"`
	Limit string `required:"true" json:"Limit" example:"2"`
}

type Position struct {
	PostedBy string `required:"true" json:"Companies" example:"superchaoran"`
	Company string `required:"true" json:"Duration" example:"Jogchat"`
	Position string `required:"true" json:"Limit" example:"Software Engineer"`
	PostedAt time.Time `required:"true" json:"Limit" example:"1537765825"`
	Description string `required:"true" json:"Limit" example:"Software Engineer position at Jogchat.com"`
}

type GetPositionsResponse struct {
   Postions []Position `required:"true" json:"Positions" example:""`
}