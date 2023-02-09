// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/movie": {
            "post": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "Admin create programme",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin create programme",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Movie": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "watch_url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AdminAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "UserAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Movie API",
	Description:      "This is collection of movie endpoint.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}