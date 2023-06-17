// Code generated by swaggo/swag. DO NOT EDIT.

package poolservice

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Khoi Nguyen",
            "email": "khoindq@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/pool/insertorappend": {
            "post": {
                "description": "Inserts or appends a pool to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pool"
                ],
                "summary": "Insert or append a pool",
                "parameters": [
                    {
                        "description": "Pool object",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/poolmodel.Pool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.successRes"
                        }
                    }
                }
            }
        },
        "/pool/quantile/get": {
            "post": {
                "description": "GetQuantile of a pool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pool"
                ],
                "summary": "GetQuantile of a pool",
                "parameters": [
                    {
                        "description": "PoolQuantileGet object",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/poolmodel.PoolQuantileGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.successRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.successRes": {
            "type": "object",
            "properties": {
                "data": {},
                "filter": {
                    "description": "for future  using"
                },
                "paging": {
                    "description": "for future  using"
                }
            }
        },
        "poolmodel.Pool": {
            "type": "object",
            "required": [
                "poolId",
                "poolValues"
            ],
            "properties": {
                "poolId": {
                    "description": "ID of the pool",
                    "type": "integer"
                },
                "poolValues": {
                    "description": "ID of the pool",
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "poolmodel.PoolQuantileGet": {
            "type": "object",
            "required": [
                "percentile",
                "poolId"
            ],
            "properties": {
                "percentile": {
                    "description": "Desired percentile",
                    "type": "number"
                },
                "poolId": {
                    "description": "ID of the pool",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0 d",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Pool service",
	Description:      "A tcp homework backend server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
