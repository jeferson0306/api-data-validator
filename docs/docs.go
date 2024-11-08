// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/validate": {
            "get": {
                "description": "Validates different types of data, including email, CPF, name, phone, RG, CEP, and credit card number",
                "tags": [
                    "Validation"
                ],
                "summary": "Validates provided data format",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email to be validated",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "CPF to be validated",
                        "name": "cpf",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name to be validated",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Phone number to be validated",
                        "name": "telephone",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Credit card number to be validated",
                        "name": "plastic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "RG to be validated",
                        "name": "rg",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "CEP (postal code) to be validated",
                        "name": "cep",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ValidationResponse": {
            "type": "object",
            "properties": {
                "execution_time_ms": {
                    "type": "integer"
                },
                "from_cache": {
                    "type": "boolean"
                },
                "is_valid": {
                    "type": "boolean"
                },
                "location_data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
                    "type": "string"
                },
                "parameter_key": {
                    "type": "string"
                },
                "parameter_value": {
                    "type": "string"
                },
                "raw_parameter_value": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
