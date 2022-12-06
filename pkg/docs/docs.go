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
        "contact": {
            "name": "envd maintainers",
            "url": "https://github.com/tensorchord/envd",
            "email": "envd-maintainers@tensorchord.ai"
        },
        "license": {
            "name": "MPL 2.0",
            "url": "https://mozilla.org/MPL/2.0/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/config": {
            "post": {
                "description": "It is called by the containerssh webhook. and is not expected to be used externally.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ssh-internal"
                ],
                "summary": "Update the config of containerssh.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login to the server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "login the user.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AuthNRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.AuthNResponse"
                        }
                    }
                }
            }
        },
        "/pubkey": {
            "post": {
                "description": "It is called by the containerssh webhook. and is not expected to be used externally.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ssh-internal"
                ],
                "summary": "authenticate the public key.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.PublicKeyAuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.ResponseBody"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "register the user for the given public key.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "register the user.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AuthNRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.AuthNResponse"
                        }
                    }
                }
            }
        },
        "/users/{login_name}/environments": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "description": "List the environment.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "environment"
                ],
                "summary": "List the environment.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"alice\"",
                        "description": "login name",
                        "name": "login_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.EnvironmentListResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "description": "Create the environment.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "environment"
                ],
                "summary": "Create the environment.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"alice\"",
                        "description": "login name",
                        "name": "login_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.EnvironmentCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.EnvironmentCreateResponse"
                        }
                    }
                }
            }
        },
        "/users/{login_name}/environments/{name}": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "description": "Get the environment with the given environment name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "environment"
                ],
                "summary": "Get the environment.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"alice\"",
                        "description": "login name",
                        "name": "login_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"pytorch-example\"",
                        "description": "environment name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.EnvironmentGetResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "description": "Remove the environment.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "environment"
                ],
                "summary": "Remove the environment.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"alice\"",
                        "description": "login name",
                        "name": "login_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"pytorch-example\"",
                        "description": "environment name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.EnvironmentRemoveResponse"
                        }
                    }
                }
            }
        },
        "/users/{login_name}/images": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "description": "List the images.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "List the images.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"alice\"",
                        "description": "login name",
                        "name": "login_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"pytorch-example\"",
                        "description": "image name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ImageListResponse"
                        }
                    }
                }
            }
        },
        "/users/{login_name}/images/{name}": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "description": "Get the image with the given image name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Get the image.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"alice\"",
                        "description": "login name",
                        "name": "login_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"pytorch-example\"",
                        "description": "image name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ImageGetResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.PublicKeyAuthRequest": {
            "type": "object",
            "properties": {
                "clientVersion": {
                    "description": "ClientVersion contains the version string the connecting client sent if any. May be empty if the client did not\nprovide a client version.\n\nrequired: false\nin: body",
                    "type": "string"
                },
                "connectionId": {
                    "description": "ConnectionID is an opaque ID to identify the SSH connection in question.\n\nrequired: true\nin: body",
                    "type": "string"
                },
                "environment": {
                    "description": "Environment is a set of key-value pairs provided by the authentication or configuration system and may be\nexposed by the backend.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.Value"
                    }
                },
                "files": {
                    "description": "Files is a key-value pair of file names and their content set by the authentication or configuration system\nand consumed by the backend.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.BinaryValue"
                    }
                },
                "metadata": {
                    "description": "Metadata is a set of key-value pairs that carry additional information from the authentication and configuration\nsystem to the backends. Backends can expose this information as container labels, environment variables, or\nother places.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.Value"
                    }
                },
                "publicKey": {
                    "description": "PublicKey is the key in the authorized key format.\n\nrequired: true",
                    "type": "string"
                },
                "remoteAddress": {
                    "description": "RemoteAddress is the IP address and port of the user trying to authenticate.\n\nrequired: true\nin: body",
                    "type": "string"
                },
                "username": {
                    "description": "Username is the username provided on login by the client. This may, but must not necessarily match the\nauthenticated username.\n\nrequired: true\nin: body",
                    "type": "string"
                }
            }
        },
        "auth.ResponseBody": {
            "type": "object",
            "properties": {
                "authenticatedUsername": {
                    "description": "AuthenticatedUsername contains the username that was actually verified. This may differ from LoginUsername when,\nfor example OAuth2 or Kerberos authentication is used. This field is empty until the authentication phase is\ncompleted.\n\nrequired: false\nin: body",
                    "type": "string"
                },
                "clientVersion": {
                    "description": "ClientVersion contains the version string the connecting client sent if any. May be empty if the client did not\nprovide a client version.\n\nrequired: false\nin: body",
                    "type": "string"
                },
                "connectionId": {
                    "description": "ConnectionID is an opaque ID to identify the SSH connection in question.\n\nrequired: true\nin: body",
                    "type": "string"
                },
                "environment": {
                    "description": "Environment is a set of key-value pairs provided by the authentication or configuration system and may be\nexposed by the backend.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.Value"
                    }
                },
                "files": {
                    "description": "Files is a key-value pair of file names and their content set by the authentication or configuration system\nand consumed by the backend.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.BinaryValue"
                    }
                },
                "metadata": {
                    "description": "Metadata is a set of key-value pairs that carry additional information from the authentication and configuration\nsystem to the backends. Backends can expose this information as container labels, environment variables, or\nother places.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.Value"
                    }
                },
                "remoteAddress": {
                    "description": "RemoteAddress is the IP address and port of the user trying to authenticate.\n\nrequired: true\nin: body",
                    "type": "string"
                },
                "success": {
                    "description": "Success indicates if the authentication was successful.\n\nrequired: true\nin: body",
                    "type": "boolean"
                },
                "username": {
                    "description": "Username is the username provided on login by the client. This may, but must not necessarily match the\nauthenticated username.\n\nrequired: true\nin: body",
                    "type": "string"
                }
            }
        },
        "config.Request": {
            "type": "object",
            "properties": {
                "authenticatedUsername": {
                    "description": "AuthenticatedUsername contains the username that was actually verified. This may differ from LoginUsername when,\nfor example OAuth2 or Kerberos authentication is used. This field is empty until the authentication phase is\ncompleted.\n\nrequired: false\nin: body",
                    "type": "string"
                },
                "clientVersion": {
                    "description": "ClientVersion contains the version string the connecting client sent if any. May be empty if the client did not\nprovide a client version.\n\nrequired: false\nin: body",
                    "type": "string"
                },
                "connectionId": {
                    "description": "ConnectionID is an opaque ID to identify the SSH connection in question.\n\nrequired: true\nin: body",
                    "type": "string"
                },
                "environment": {
                    "description": "Environment is a set of key-value pairs provided by the authentication or configuration system and may be\nexposed by the backend.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.Value"
                    }
                },
                "files": {
                    "description": "Files is a key-value pair of file names and their content set by the authentication or configuration system\nand consumed by the backend.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.BinaryValue"
                    }
                },
                "metadata": {
                    "description": "Metadata is a set of key-value pairs that carry additional information from the authentication and configuration\nsystem to the backends. Backends can expose this information as container labels, environment variables, or\nother places.\n\nrequired: false\nin: body",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/metadata.Value"
                    }
                },
                "remoteAddress": {
                    "description": "RemoteAddress is the IP address and port of the user trying to authenticate.\n\nrequired: true\nin: body",
                    "type": "string"
                },
                "username": {
                    "description": "Username is the username provided on login by the client. This may, but must not necessarily match the\nauthenticated username.\n\nrequired: true\nin: body",
                    "type": "string"
                }
            }
        },
        "metadata.BinaryValue": {
            "type": "object",
            "properties": {
                "sensitive": {
                    "description": "Sensitive indicates that the metadata value contains sensitive data and should not be transmitted to\nservers unnecessarily.\n\nrequired: false\nin: body",
                    "type": "boolean"
                },
                "value": {
                    "description": "Value contains the binary data for the current value.\n\nrequired: true\nin: body\nswagger:strfmt: byte",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "metadata.Value": {
            "type": "object",
            "properties": {
                "sensitive": {
                    "description": "Sensitive indicates that the metadata value contains sensitive data and should not be transmitted to\nservers unnecessarily.",
                    "type": "boolean"
                },
                "value": {
                    "description": "Value contains the string for the current value.",
                    "type": "string"
                }
            }
        },
        "types.AuthNRequest": {
            "type": "object",
            "properties": {
                "login_name": {
                    "description": "LoginName is used to authenticate the user and get\nan access token for the registry.",
                    "type": "string",
                    "example": "alice"
                },
                "password": {
                    "description": "Password stores the hashed password.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "public_key": {
                    "type": "string"
                }
            }
        },
        "types.AuthNResponse": {
            "type": "object",
            "properties": {
                "identity_token": {
                    "description": "An opaque token used to authenticate a user after a successful login\nRequired: true",
                    "type": "string",
                    "example": "a332139d39b89a241400013700e665a3"
                },
                "login_name": {
                    "description": "LoginName is used to authenticate the user and get\nan access token for the registry.",
                    "type": "string",
                    "example": "alice"
                },
                "status": {
                    "description": "The status of the authentication\nRequired: true",
                    "type": "string",
                    "example": "Login successfully"
                }
            }
        },
        "types.Environment": {
            "type": "object",
            "properties": {
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "resource": {
                    "$ref": "#/definitions/types.ResourceSpec"
                },
                "spec": {
                    "$ref": "#/definitions/types.EnvironmentSpec"
                },
                "status": {
                    "$ref": "#/definitions/types.EnvironmentStatus"
                }
            }
        },
        "types.EnvironmentCreateRequest": {
            "type": "object",
            "properties": {
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "resource": {
                    "$ref": "#/definitions/types.ResourceSpec"
                },
                "spec": {
                    "$ref": "#/definitions/types.EnvironmentSpec"
                },
                "status": {
                    "$ref": "#/definitions/types.EnvironmentStatus"
                }
            }
        },
        "types.EnvironmentCreateResponse": {
            "type": "object",
            "properties": {
                "environment": {
                    "$ref": "#/definitions/types.Environment"
                },
                "warnings": {
                    "description": "Warnings encountered when creating the pod\nRequired: true",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "types.EnvironmentGetResponse": {
            "type": "object",
            "properties": {
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "resource": {
                    "$ref": "#/definitions/types.ResourceSpec"
                },
                "spec": {
                    "$ref": "#/definitions/types.EnvironmentSpec"
                },
                "status": {
                    "$ref": "#/definitions/types.EnvironmentStatus"
                }
            }
        },
        "types.EnvironmentListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Environment"
                    }
                }
            }
        },
        "types.EnvironmentPort": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "types.EnvironmentRemoveResponse": {
            "type": "object"
        },
        "types.EnvironmentSpec": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "env": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "image": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "ports": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.EnvironmentPort"
                    }
                }
            }
        },
        "types.EnvironmentStatus": {
            "type": "object",
            "properties": {
                "jupyter_addr": {
                    "type": "string"
                },
                "phase": {
                    "type": "string"
                },
                "rstudio_server_addr": {
                    "type": "string"
                }
            }
        },
        "types.ImageGetResponse": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "integer"
                },
                "digest": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string",
                    "example": "pytorch-cuda:dev"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "types.ImageListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.ImageMeta"
                    }
                }
            }
        },
        "types.ImageMeta": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "integer"
                },
                "digest": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string",
                    "example": "pytorch-cuda:dev"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "types.ResourceSpec": {
            "type": "object",
            "properties": {
                "cpu": {
                    "type": "string"
                },
                "gpu": {
                    "type": "string"
                },
                "memory": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authentication": {
            "type": "apiKey",
            "name": "JWT",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.12",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "envd server API",
	Description:      "envd backend server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
