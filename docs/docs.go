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
        "/installment": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Membuat transaksi baru",
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
                            "$ref": "#/definitions/models.InstallmentRequest"
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
        "/login": {
            "post": {
                "description": "Autentikasi user dengan email: \"iqmalr@gmail.com\" dan password: \"password\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Email dan Password",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token JWT",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Input tidak valid",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Login gagal",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Menghapus token JWT",
                "tags": [
                    "auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "Logout berhasil",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/merchants": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Mengambil daftar semua pembayaran",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Get all payment",
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
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/models.PaymentRequest"
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
        "/recap": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Mengambil rekap cicilan dan status pembayaran",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "installment"
                ],
                "summary": "Get installment recap",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InstallmentRecap"
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
        "/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
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
        "models.InstallmentRecap": {
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
                "payments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Payment"
                    }
                },
                "remaining": {
                    "type": "number"
                },
                "status": {
                    "type": "string"
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
        "models.InstallmentRequest": {
            "type": "object",
            "properties": {
                "merchant_id": {
                    "type": "integer"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "iqmalr@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "password"
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
                "payment_method": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.PaymentRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 200000
                },
                "payment_method": {
                    "type": "string",
                    "example": "bank_transfer"
                },
                "transaction_id": {
                    "type": "integer",
                    "example": 1
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
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Masukkan token dengan format **\"{token}\"**",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
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
