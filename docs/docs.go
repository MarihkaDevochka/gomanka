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
        "/filter": {
            "get": {
                "description": "Find Manga Chapter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manga"
                ],
                "summary": "Get a chapter",
                "operationId": "Filter-anime",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Manga",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Chapter of the Manga",
                        "name": "genres",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name of the Manga",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Chapter of the Manga",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "field of the Manga",
                        "name": "orderField",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort of the Manga",
                        "name": "orderSort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page not 0",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "perPage",
                        "name": "perPage",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.MangaSwag"
                            }
                        }
                    }
                }
            }
        },
        "/manga": {
            "get": {
                "description": "Retrieve a manga by its name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manga"
                ],
                "summary": "Get a manga by name",
                "operationId": "get-manga-by-name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Manga",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.MangaSwag"
                        }
                    }
                }
            }
        },
        "/manga/{name}/{chapter}": {
            "get": {
                "description": "Find Manga Chapter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manga"
                ],
                "summary": "Get a chapter",
                "operationId": "get-chapter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Manga",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Chapter of the Manga",
                        "name": "chapter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ChapterSwag"
                        }
                    }
                }
            }
        },
        "/mangas": {
            "get": {
                "description": "Retrieve a list of all mangas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manga"
                ],
                "summary": "Get all mangas",
                "operationId": "get-all-mangas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.MangaSwag"
                            }
                        }
                    }
                }
            }
        },
        "/popular": {
            "get": {
                "description": "Retrieve a list of popular mangas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manga"
                ],
                "summary": "Get popular mangas",
                "operationId": "get-popular-manga",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.MangaSwag"
                            }
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "Create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create or cheack user",
                "operationId": "create-or-cheack-user",
                "parameters": [
                    {
                        "description": "Auth Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UserSwag"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "description": "Delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete user by email",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/user/favorite/list": {
            "get": {
                "description": "User Favorites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User favorite Mangas",
                "operationId": "get-user-list-manga",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.MangaSwag"
                            }
                        }
                    }
                }
            }
        },
        "/user/favorite/one": {
            "get": {
                "description": "User Favorite",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User favorite Manga",
                "operationId": "get-user-favorite-manga",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.FavoriteResponse"
                        }
                    }
                }
            }
        },
        "/user/favorite/{name}/{email}": {
            "post": {
                "description": "Toggle manga",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Toggle Favorite manga",
                "operationId": "toggle-favorite-manga",
                "parameters": [
                    {
                        "type": "string",
                        "description": "manga name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/user/{email}": {
            "get": {
                "description": "Retrieve a user its email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get a user by email",
                "operationId": "get-user-by-email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UserSwag"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ChapterSwag": {
            "type": "object",
            "properties": {
                "animeName": {
                    "type": "string"
                },
                "chapter": {
                    "type": "integer"
                },
                "created": {
                    "type": "string"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.FavoriteResponse": {
            "type": "object",
            "properties": {
                "isFavorite": {
                    "type": "boolean"
                }
            }
        },
        "handler.MangaSwag": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "averageRating": {
                    "type": "number"
                },
                "chapters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.ChapterSwag"
                    }
                },
                "country": {
                    "type": "string"
                },
                "describe": {
                    "type": "string"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "img": {
                    "type": "string"
                },
                "imgHeader": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "popularity": {
                    "type": "integer"
                },
                "published": {
                    "type": "integer"
                },
                "ratingCount": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "handler.SuccessResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "string"
                }
            }
        },
        "handler.UserSwag": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "favorite": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Manka Api",
	Description:      "Manga search",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
