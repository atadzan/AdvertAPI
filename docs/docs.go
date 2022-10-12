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
                        "type": "string",
                        "description": "page info",
                        "name": "page",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add Advert to DB",
                "consumes": [
                    "multipart/form-data"
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
                        "name": "category",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "description",
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
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "user_id",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get User Favourite List",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Get User Favourite List",
                "operationId": "get_fav",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        "/api/advert/search": {
            "get": {
                "description": "Search Adverts by Title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "Search",
                "operationId": "search_adv",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status",
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
                            "$ref": "#/definitions/AdvertAPI.AdvertInfo"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                        "type": "integer",
                        "name": "comment_count",
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
                        "type": "integer",
                        "name": "user_id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        "/api/advert/{id}/comment/{comment_id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Advert Comment By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Update Comment",
                "operationId": "update comment",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete Advert Comment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Delete Comment",
                "operationId": "delete_comment",
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
        "/api/{id}/comment": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add Comment to Advert",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Add Comment",
                "operationId": "add_comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "advert id",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        "/auth/sign-in": {
            "post": {
                "description": "Sign in app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
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
                "description": "Create account in app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
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
        "AdvertAPI.AdvertImage": {
            "type": "object",
            "properties": {
                "advertId": {
                    "type": "integer"
                },
                "fname": {
                    "type": "string"
                },
                "fsize": {
                    "type": "integer"
                },
                "ftype": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "AdvertAPI.AdvertInfo": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "comment_count": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/AdvertAPI.Comment"
                    }
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
                "user_id": {
                    "type": "integer"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "AdvertAPI.Comment": {
            "type": "object",
            "properties": {
                "advert_id": {
                    "type": "integer"
                },
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
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
                 "username": {
                    "type": "string"
                },
				"password": {
                    "type": "string"
                }
            }
        },
        "AdvertAPI.SignUpInput": {
            "type": "object",
            "properties": {
               	 "username": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
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
