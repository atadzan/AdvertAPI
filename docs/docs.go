// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/advert": {
            "get": {
                "description": "Get adverts by page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Get Adverts",
                "operationId": "get_adverts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Get adverts",
                        "name": "pageStr",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/AdvertAPI.AdvertInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "post": {
                "description": "Add Advert to DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Add Advert",
                "operationId": "add_advert",
                "parameters": [
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/api/advert/fav": {
            "get": {
                "description": "Get User Favourite List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Get User Favourite List",
                "operationId": "get_fav",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user",
                        "name": "default",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/api/advert/fav/{id}": {
            "post": {
                "description": "Add Advert to Favourite List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Add Advert to Favourite List",
                "operationId": "add_fav",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "credentials",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Advert from Favourite List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Delete Advert from Favourite List",
                "operationId": "del_fav",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "advert",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/api/advert/{id}": {
            "get": {
                "description": "Get Advert by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Get Advert by ID",
                "operationId": "get_advert",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "advert ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/AdvertAPI.AdvertInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Advert",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Update Advert",
                "operationId": "update_advert",
                "parameters": [
                    {
                        "type": "string",
                        "name": "category",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "images_count",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "location",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "phone_number",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "price",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "publish_date",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "views",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Advert",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Delete Advert",
                "operationId": "del_advert",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "advert ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "operationId": "login",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/AdvertAPI.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/AdvertAPI.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AdvertAPI.AdvertInfo": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/AdvertAPI.ImageUrl"
                    }
                },
                "images_count": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "publish_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "views": {
                    "type": "string"
                }
            }
        },
        "AdvertAPI.ImageUrl": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "AdvertAPI.SignInInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "AdvertAPI.SignUpInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Advert App API",
	Description:      "API Server for Advert Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}