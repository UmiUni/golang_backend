package model

type AddCompanyRequest struct {
	Name string `required:"true" json:"Name" example:"Google"`
	Domain string `required:"true" json:"Domain" example:"google.com"`
}

type AddCompanyResponseSuccess struct {
	Message string `json:"message" example:"status 200"`
}

type AddCompanyResponseError0 struct {
	ErrorMessage string `json:"error" example:"schemaless add company fail"`
}

type AddSchoolRequest struct {
	Name string `required:"true" json:"Name" example:"University of Illinois at Urbana-Champaign"`
	Domain string `required:"true" json:"Domain" example:"illinois.edu"`
}

type AddSchoolResponseSuccess struct {
	Message string `json:"message" example:"status 200"`
}

type AddSchoolResponseError0 struct {
	ErrorMessage string `json:"error" example:"schemaless add school fail"`
}
