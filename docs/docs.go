// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-08-13 09:18:41.764409685 -0700 PDT m=+0.063184178

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a ReferHelper API server.",
        "title": "ReferHelper API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "admin@jogchat.com",
            "email": "admin@jogchat.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "178.128.0.108:3001",
    "basePath": "/",
    "paths": {
        "/activate_and_signup": {
            "post": {
                "description": "When user click on the GET link in user email, it will hit a frontend page as a GET request with {Email, Token} as parameters. The frontend page should then provide user with a form that ask for (Email(prefilled), Username, password, token(prefilled and hidden)). Once frontend gather all infos from the user, frontend should POST call this [ActivateAndSignup endpoint] with a post request that has {email, username, password, token} as JSON to sign the user up. This endpoint will both signup the user and activate their account.",
                "consumes": [
                    "application/json"
                ],
                "summary": "ActivateAndSignup",
                "parameters": [
                    {
                        "description": "ActivateAndSignupRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ActivateAndSignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: verification email sent",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ActivateAndSignupResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "email already activated",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ActivateAndSignupResponseAPIError2"
                        }
                    }
                }
            }
        },
        "/add_company": {
            "post": {
                "description": "AddCompany is an endpoint that adds company json(id(generated), name, domain) to schemaless database",
                "consumes": [
                    "application/json"
                ],
                "summary": "AddCompany",
                "parameters": [
                    {
                        "description": "AddCompanyRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddCompanyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: schemaless add company success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddCompanyResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Failure: schemaless add company fail",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddCompanyResponseError0"
                        }
                    }
                }
            }
        },
        "/add_school": {
            "post": {
                "description": "AddSchool is an endpoint that adds school json(id(generated), name, domain) to schemaless database",
                "consumes": [
                    "application/json"
                ],
                "summary": "AddSchool",
                "parameters": [
                    {
                        "description": "AddSchoolRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddSchoolRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: schemaless add school success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddSchoolResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Failure: schemaless add school fail",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddSchoolResponseError0"
                        }
                    }
                }
            }
        },
        "/applicant_check_signup_email": {
            "post": {
                "description": "Provide a school/university edu email to sign up for the applicant portal, if the email does not exists in schemaless database, we will send the email an activation link",
                "consumes": [
                    "application/json"
                ],
                "summary": "ApplicantCheckSignupEmail",
                "parameters": [
                    {
                        "description": "ApplicantSignupEmailRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApplicantSignupEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: verification email sent",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApplicantSignupResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "email already registered",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApplicantSignupResponseAPIError1"
                        }
                    }
                }
            }
        },
        "/get_all_companies": {
            "get": {
                "description": "GetAllCompanies is an endpoint that returns companies list from schemaless database",
                "summary": "Get All Companies"
            }
        },
        "/get_all_schools": {
            "get": {
                "description": "GetAllSchools is an endpoint that returns schools list from schemaless database",
                "summary": "Get All Schools"
            }
        },
        "/referrer_check_signup_email": {
            "post": {
                "description": "Onboarding user will provide a company email to sign up for the referral portal, if the email does not exists in schemaless database, we will send the email an activation link",
                "consumes": [
                    "application/json"
                ],
                "summary": "ReferrerCheckSignupEmail",
                "parameters": [
                    {
                        "description": "ReferrerSignupEmailRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ReferrerSignupEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: verification email sent",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ReferrerSignupResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "email already registered",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ReferrerSignupResponseAPIError1"
                        }
                    }
                }
            }
        },
        "/reset_password_form": {
            "post": {
                "description": "After user clicks on reset password link(GET with email and token) in email, front-end/mobile will provide user with a form, {Email(prefilled), Password, Token(prefilled)}. After user filled the form, front-end/mobile will call this endpoint with a JSON wrapped {Email(prefilled), Password, Token(prefilled)} POST to reset password. If the user is not activated at the point of click on reset_password, an email titled reset_password with activation instruction will be sent.",
                "consumes": [
                    "application/json"
                ],
                "summary": "ResetPasswordForm",
                "parameters": [
                    {
                        "description": "ResetPasswordFormRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ResetPasswordFormRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: message: reset email sent",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ResetPasswordFormResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Failure: email not registered",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ResetPasswordFormResponseAPIError0"
                        }
                    }
                }
            }
        },
        "/send_reset_password_email": {
            "post": {
                "description": "When user click on reset password button with an email filled in a form above, front-end will call this endpoint with a JSON wrapped {Email, Token} to sent reset password email, a hacker cannot hack this end point by repeatedly calling and our system and spam send email. Requiring a session {Email, AuthToken} combination and this endpoint will only be able to sent email to this session's Email.",
                "consumes": [
                    "application/json"
                ],
                "summary": "SendResetPasswordEmail",
                "parameters": [
                    {
                        "description": "ResetPasswordButtonRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SendResetPasswordEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: message: reset email sent",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SendResetPasswordEmailResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Failure: email not registered",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SendResetPasswordEmailResponseAPIError0"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "After user click on sign-in button, front-end will call this endpoint with a JSON wrapped {Email and Password}, the end point will then return an AuthToken on success. Front-end should store the authtoken for user either in session or cookie for user. To access password protect resource later, front-end needs to pass (username+AuthToken) to backend to verify user identity. This is called JWT Auth flow.",
                "consumes": [
                    "application/json"
                ],
                "summary": "Signin",
                "parameters": [
                    {
                        "description": "SigninRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SigninRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: sign in request succeed",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SigninResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "invalid password",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SigninResponseAPIError2"
                        }
                    }
                }
            }
        },
        "/v1/comment_on": {
            "post": {
                "description": "GetPositions is an endpoint called to get all the positions",
                "consumes": [
                    "application/json"
                ],
                "summary": "GetPositions",
                "parameters": [
                    {
                        "description": "CommentOnRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CommentOnRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success on commenting",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CommentOnResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "invalid parent type",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CommentOnResponseAPIError0"
                        }
                    }
                }
            }
        },
        "/v1/post_position": {
            "post": {
                "description": "PostPosition is an endpoint called when an referral create a job position to refer with dedicated JSON.",
                "consumes": [
                    "application/json"
                ],
                "summary": "PostPosition",
                "parameters": [
                    {
                        "description": "PostPositionRequest is a POST JSON type",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.PostPositionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success with position id: {\"id\":\"1528edfd-2cbd-451f-9053-a89e2e806cbe\"}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.PostPositionResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "construct cell failure",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.PostPositionResponseAPIError1"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ActivateAndSignupRequest": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                },
                "Password": {
                    "type": "string",
                    "example": "admin374password"
                },
                "Token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEzNTIyMDAsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoid2FuZzM3NEB1aXVjLmVkdSJ9.gC7dTl64XDe5BwlS8PuZxBxGes1ujcCWFbe23r0xOXM"
                },
                "Username": {
                    "type": "string",
                    "example": "admin374"
                }
            }
        },
        "model.ActivateAndSignupResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "username already in use"
                }
            }
        },
        "model.ActivateAndSignupResponseAPIError1": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "invalid token"
                }
            }
        },
        "model.ActivateAndSignupResponseAPIError2": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email already activated"
                }
            }
        },
        "model.ActivateAndSignupResponseSuccess": {
            "type": "object",
            "properties": {
                "AuthToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEzNTM1NjUsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoid2FuZzM3NEB1aXVjLmVkdSJ9.XwmDhW1b99E9jwGatN_6y1tYpLGBcAqywS9fI23Oxxo"
                },
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                },
                "UserId": {
                    "type": "string",
                    "example": "ce57e12a-fe27-43a2-9a1f-0792b3d36f2e"
                },
                "Username": {
                    "type": "string",
                    "example": "admin374"
                }
            }
        },
        "model.AddCompanyRequest": {
            "type": "object",
            "properties": {
                "Domain": {
                    "type": "string",
                    "example": "jogchat.com"
                },
                "Name": {
                    "type": "string",
                    "example": "Jogchat"
                }
            }
        },
        "model.AddCompanyResponseError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "schemaless add company fail"
                }
            }
        },
        "model.AddCompanyResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "status 200"
                }
            }
        },
        "model.AddSchoolRequest": {
            "type": "object",
            "properties": {
                "Domain": {
                    "type": "string",
                    "example": "illinois.edu"
                },
                "Name": {
                    "type": "string",
                    "example": "University of Illinois at Urbana-Champaign"
                }
            }
        },
        "model.AddSchoolResponseError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "schemaless add school fail"
                }
            }
        },
        "model.AddSchoolResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "status 200"
                }
            }
        },
        "model.ApplicantSignupEmailRequest": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "wang374@uiuc.edu"
                }
            }
        },
        "model.ApplicantSignupResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email cannot be empty"
                }
            }
        },
        "model.ApplicantSignupResponseAPIError1": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email already registered"
                }
            }
        },
        "model.ApplicantSignupResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "verification email sent"
                }
            }
        },
        "model.CommentOnRequest": {
            "type": "object",
            "properties": {
                "Content": {
                    "type": "string",
                    "example": "这个Position很适合我背景，请联系superchaoran@gmail.com"
                },
                "ParentId": {
                    "type": "string",
                    "example": "67bebc0c-f0bd-4352-b588-08a056085e0a"
                },
                "ParentType": {
                    "type": "string",
                    "example": "position"
                },
                "PositionId": {
                    "type": "string",
                    "example": "67bebc0c-f0bd-4352-b588-08a056085e0a"
                },
                "Username": {
                    "type": "string",
                    "example": "admin374"
                }
            }
        },
        "model.CommentOnResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "invalid parent type"
                }
            }
        },
        "model.CommentOnResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Success on commenting: status 200"
                }
            }
        },
        "model.PostPositionRequest": {
            "type": "object",
            "properties": {
                "Company": {
                    "type": "string",
                    "example": "Jogchat"
                },
                "Description": {
                    "type": "string",
                    "example": "Build a microservice platform for Jogchat. A position requires microservice knowledge and past experience in Golang."
                },
                "Position": {
                    "type": "string",
                    "example": "Software Engineer"
                },
                "Username": {
                    "type": "string",
                    "example": "admin374"
                }
            }
        },
        "model.PostPositionResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "username does not exist"
                }
            }
        },
        "model.PostPositionResponseAPIError1": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "construct cell failure"
                }
            }
        },
        "model.PostPositionResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Success: status 200 with position id {"
                }
            }
        },
        "model.ReferrerSignupEmailRequest": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                }
            }
        },
        "model.ReferrerSignupResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email cannot be empty"
                }
            }
        },
        "model.ReferrerSignupResponseAPIError1": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email already registered"
                }
            }
        },
        "model.ReferrerSignupResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "verification email sent"
                }
            }
        },
        "model.ResetPasswordFormRequest": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                },
                "Password": {
                    "type": "string",
                    "example": "admin374newpassword"
                },
                "Token": {
                    "type": "string"
                }
            }
        },
        "model.ResetPasswordFormResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email not registered"
                }
            }
        },
        "model.ResetPasswordFormResponseSuccess": {
            "type": "object",
            "properties": {
                "AuthToken": {
                    "type": "string"
                },
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                },
                "UserId": {
                    "type": "string",
                    "example": "ce57e12a-fe27-43a2-9a1f-0792b3d36f2e"
                },
                "Username": {
                    "type": "string",
                    "example": "admin374"
                }
            }
        },
        "model.SendResetPasswordEmailRequest": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                }
            }
        },
        "model.SendResetPasswordEmailResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email not registered"
                }
            }
        },
        "model.SendResetPasswordEmailResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "reset email sent"
                }
            }
        },
        "model.SigninRequest": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                },
                "Password": {
                    "type": "string",
                    "example": "admin374password"
                }
            }
        },
        "model.SigninResponseAPIError0": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "email not registered"
                }
            }
        },
        "model.SigninResponseAPIError1": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "please verify your email"
                }
            }
        },
        "model.SigninResponseAPIError2": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "invalid password"
                }
            }
        },
        "model.SigninResponseSuccess": {
            "type": "object",
            "properties": {
                "AuthToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEzNTM2MzgsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoid2FuZzM3NEB1aXVjLmVkdSJ9.RhRUpHJbIfid1hiJOTtStuxc86v0isnWny85COG9Mek"
                },
                "Email": {
                    "type": "string",
                    "example": "admin@umiuni.com"
                },
                "UserId": {
                    "type": "string",
                    "example": "ce57e12a-fe27-43a2-9a1f-0792b3d36f2e"
                },
                "Username": {
                    "type": "string",
                    "example": "admin374"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
