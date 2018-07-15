package model

type AddCompanyRequest struct {
	Name string `required:"true" json:"name" example:"Google"`
	Domain string `required:"true" json:"domain" example:"google.com"`
}

type AddCompanyRequestError0 struct {
	ErrorMessage string `json:"error" example:"schemaless add company fail"`
}

type AddSchoolRequest struct {
	Name string `required:"true" json:"name" example:"University of Illinois at Urbana-Champaign"`
	Domain string `required:"true" json:"domain" example:"illinois.edu"`
}

type AddSchoolRequestError0 struct {
	ErrorMessage string `json:"error" example:"schemaless add school fail"`
}
