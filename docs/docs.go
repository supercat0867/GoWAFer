// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/waf/api/v1/auth/dologin": {
            "post": {
                "description": "dologin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "处理登录",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "api_handler.LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_handler.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            }
        },
        "/waf/api/v1/auth/getCaptcha": {
            "get": {
                "description": "获取图片验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "获取图片验证码",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            }
        },
        "/waf/api/v1/ip": {
            "get": {
                "description": "分页查询IP",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IP"
                ],
                "summary": "分页查询IP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询IP",
                        "name": "keywords",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IP类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页面大小",
                        "name": "perPage",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "新增IP记录",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "IP"
                ],
                "summary": "新增IP记录",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "api_handler.CreateIPRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_handler.CreateIPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            }
        },
        "/waf/api/v1/ip/{id}": {
            "delete": {
                "description": "删除IP",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IP"
                ],
                "summary": "删除IP",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IP主键ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "编辑IP",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IP"
                ],
                "summary": "编辑IP",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IP主键ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求体",
                        "name": "api_handler.UpdateIPRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_handler.UpdateIPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            }
        },
        "/waf/api/v1/log": {
            "get": {
                "description": "查询指定天数和小时数的日志记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "查询指定天数和小时数的日志记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询范围天数",
                        "name": "days",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_handler.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_handler.CreateIPRequest": {
            "type": "object",
            "required": [
                "ip-address"
            ],
            "properties": {
                "expirationAt": {
                    "type": "string"
                },
                "ip-address": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "api_handler.LoginRequest": {
            "type": "object",
            "required": [
                "captcha",
                "captchaId",
                "password",
                "username"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captchaId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api_handler.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "api_handler.UpdateIPRequest": {
            "type": "object",
            "required": [
                "ip_address"
            ],
            "properties": {
                "expiration_at": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
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
	Version:     "v0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "GoWAFer",
	Description: "Golang编写的一款基于反向代理模式的web防火墙应用 By supercat0867",
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
	swag.Register(swag.Name, &s{})
}
