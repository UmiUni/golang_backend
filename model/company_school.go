package model

type AddCompanyRequest struct {
	Name string `required:"true" json:"name" example:"Google"`
	Domain string `required:"true" json:"domain" example:"google"`
}

type AddCompanyRequestError0 struct {
	ErrorMessage string `json:"error" example:"email not registered"`
}

type AddSchoolRequest struct {
	Email string `required:"true" json:"Email" example:"wang374@uiuc.edu"`
}

type AddSchoolRequestError0 struct {
	ErrorMessage string `json:"error" example:"schemaless add school fail"`
}
