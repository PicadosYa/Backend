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
        "/reservations": {
            "get": {
                "description": "Devuelve una lista paginada de reservas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Obtiene una lista de reservas",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Número de reservas a obtener",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Desplazamiento para paginación",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Reservation"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualiza una reserva existente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Actualiza una reserva",
                "parameters": [
                    {
                        "description": "Reserva a actualizar",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Crea una nueva reserva en el sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Crea una reserva",
                "parameters": [
                    {
                        "description": "Reserva a crear",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            }
        },
        "/reservations/{id}": {
            "get": {
                "description": "Devuelve una reserva específica por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Obtiene una reserva",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la reserva",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Elimina una reserva específica por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Elimina una reserva",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la reserva",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.responseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "api.responseMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Reservation": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "end_time": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "field_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "reservation_date": {
                    "type": "string"
                },
                "start_time": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "PicadosYa API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
