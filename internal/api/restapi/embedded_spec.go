// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Backend service",
    "title": "Mproducto application",
    "version": "0.0.1"
  },
  "paths": {
    "/user": {
      "get": {
        "summary": "Show user profile",
        "responses": {
          "200": {
            "description": "User profile",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "$ref": "#/responses/Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "User": {
      "description": "The profile of the user",
      "type": "object",
      "required": [
        "name",
        "email"
      ],
      "properties": {
        "email": {
          "description": "user email",
          "type": "string",
          "format": "string"
        },
        "name": {
          "description": "user name",
          "type": "string",
          "format": "string"
        }
      }
    }
  },
  "responses": {
    "Error": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "api_key": null
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Backend service",
    "title": "Mproducto application",
    "version": "0.0.1"
  },
  "paths": {
    "/user": {
      "get": {
        "summary": "Show user profile",
        "responses": {
          "200": {
            "description": "User profile",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "User": {
      "description": "The profile of the user",
      "type": "object",
      "required": [
        "name",
        "email"
      ],
      "properties": {
        "email": {
          "description": "user email",
          "type": "string",
          "format": "string"
        },
        "name": {
          "description": "user name",
          "type": "string",
          "format": "string"
        }
      }
    }
  },
  "responses": {
    "Error": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "api_key": []
    }
  ]
}`))
}
