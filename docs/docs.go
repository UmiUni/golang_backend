// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-07-07 23:50:00.9419465 -0700 PDT m=+0.023726531

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/referrer_signup": {
            "post": {
                "description": "Signup",
                "consumes": [
                    "application/json"
                ],
                "summary": "Signup",
                "parameters": [
                    {
                        "description": "Body JSON",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ReferrerSignupEmailStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: verification email sent",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ReferrerSignupSuccessStruct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ReferrerSignupEmailStruct": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "superchaoran@gmail.com"
                }
            }
        },
        "model.ReferrerSignupSuccessStruct": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "verification email sent"
                }
            }
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
