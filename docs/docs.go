// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "m.iqmal.riffai@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/installment": {
            "get": {
                "description": "Menampilkan daftar semua cicilan",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "installment"
                ],
                "summary": "Get all installments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Installment"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Membuat cicilan baru",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "installment"
                ],
                "summary": "Create new installment",
                "parameters": [
                    {
                        "description": "Data Cicilan",
                        "name": "installment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Installment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Installment"
                        }
                    },
                    "400": {
                        "description": "Data tidak valid",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Kesalahan server",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchants": {
            "get": {
                "description": "Mengambil daftar semua merchant",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "merchant"
                ],
                "summary": "Get all merchants",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Merchant"
                            }
                        }
                    },
                    "500": {
                        "description": "Kesalahan server",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/payments": {
            "get": {
                "description": "Mengambil daftar semua pembayaran",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Get all payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Payment"
                            }
                        }
                    },
                    "500": {
                        "description": "Kesalahan server",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Membuat pembayaran baru",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Create new payment",
                "parameters": [
                    {
                        "description": "Data Pembayaran",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Payment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Payment"
                        }
                    },
                    "400": {
                        "description": "Data tidak valid",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Kesalahan server",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Menampilkan daftar semua pengguna",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Installment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "total_amount": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Merchant": {
            "type": "object",
            "properties": {
                "bank_account": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merchant_name": {
                    "type": "string"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "payment_method": {
                    "type": "string"
                },
                "transaction_id": {
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
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "utils.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MNC Test API",
	Description:      "API untuk test MNC",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
