// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Duy Nguyen",
            "email": "duynguyenngoc@hotmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/celery/status/{task_id}": {
            "get": {
                "description": "show result celery status base on task id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Celery"
                ],
                "summary": "Celery Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apientities.CeleryStatusResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apientities.CeleryStatus": {
            "type": "object",
            "properties": {
                "general": {
                    "type": "string",
                    "format": "string",
                    "example": "SUCCESS"
                },
                "ml": {
                    "type": "string",
                    "format": "string",
                    "example": "SUCCESS"
                },
                "upload": {
                    "type": "string",
                    "format": "string",
                    "example": "SUCCESS"
                }
            }
        },
        "apientities.CeleryStatusResult": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/apientities.CeleryStatus"
                },
                "task_id": {
                    "type": "string",
                    "format": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "times": {
                    "$ref": "#/definitions/apientities.CeleryTimes"
                }
            }
        },
        "apientities.CeleryTimes": {
            "type": "object",
            "properties": {
                "end_ml": {
                    "type": "string",
                    "format": "string",
                    "example": "2021-12-29 10:11:22"
                },
                "end_upload": {
                    "type": "string",
                    "format": "string",
                    "example": "2021-12-29 10:11:22"
                },
                "start_ml": {
                    "type": "string",
                    "format": "string",
                    "example": "2021-12-29 10:11:22"
                },
                "start_upload": {
                    "type": "string",
                    "format": "string",
                    "example": "2021-12-29 10:11:22"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger x-api",
	Description: "This is a x-ml api server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
