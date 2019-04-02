package hlsp

func getSchema() []byte {
	return []byte(`
{
    "title": "AsyncAPI 2.0.0 schema.",
    "$schema": "http://json-schema.org/draft-07/schema#",
    "type": "object",
    "required": [
        "asyncapi",
        "id",
        "info"
    ],
    "oneOf": [
        {
            "required": [
                "channels"
            ]
        },
        {
            "required": [
                "stream"
            ]
        }
    ],
    "additionalProperties": false,
    "patternProperties": {
        "^x-": {
            "$ref": "#/definitions/vendorExtension"
        }
    },
    "properties": {
        "asyncapi": {
            "type": "string",
            "enum": [
                "2.0.0"
            ],
            "description": "The AsyncAPI specification version of this document."
        },
        "id": {
            "type": "string",
            "description": "A unique id representing the application.",
            "format": "uri-reference"
        },
        "info": {
            "$ref": "#/definitions/info"
        },
        "servers": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/server"
            },
            "uniqueItems": true
        },
        "defaultContentType": {
            "type": "string"
        },
        "channels": {
            "$ref": "#/definitions/channels"
        },
        "stream": {
            "$ref": "#/definitions/stream",
            "description": "The list of messages a consumer can read or write from/to a streaming API."
        },
        "components": {
            "$ref": "#/definitions/components"
        },
        "tags": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/tag"
            },
            "uniqueItems": true
        },
        "externalDocs": {
            "$ref": "#/definitions/externalDocs"
        }
    },
    "definitions": {
        "Reference": {
            "type": "object",
            "required": [
                "$ref"
            ],
            "properties": {
                "$ref": {
                    "type": "string",
                    "format": "uri"
                }
            }
        },
        "info": {
            "type": "object",
            "description": "General information about the API.",
            "required": [
                "version",
                "title"
            ],
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "properties": {
                "title": {
                    "type": "string",
                    "description": "A unique and precise title of the API."
                },
                "version": {
                    "type": "string",
                    "description": "A semantic version number of the API."
                },
                "description": {
                    "type": "string",
                    "description": "A longer description of the API. Should be different from the title. CommonMark is allowed."
                },
                "termsOfService": {
                    "type": "string",
                    "description": "A URL to the Terms of Service for the API. MUST be in the format of a URL.",
                    "format": "uri"
                },
                "contact": {
                    "$ref": "#/definitions/contact"
                },
                "license": {
                    "$ref": "#/definitions/license"
                }
            }
        },
        "contact": {
            "type": "object",
            "description": "Contact information for the owners of the API.",
            "additionalProperties": false,
            "properties": {
                "name": {
                    "type": "string",
                    "description": "The identifying name of the contact person/organization."
                },
                "url": {
                    "type": "string",
                    "description": "The URL pointing to the contact information.",
                    "format": "uri"
                },
                "email": {
                    "type": "string",
                    "description": "The email address of the contact person/organization.",
                    "format": "email"
                }
            },
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            }
        },
        "license": {
            "type": "object",
            "required": [
                "name"
            ],
            "additionalProperties": false,
            "properties": {
                "name": {
                    "type": "string",
                    "description": "The name of the license type. It's encouraged to use an OSI compatible license."
                },
                "url": {
                    "type": "string",
                    "description": "The URL pointing to the license.",
                    "format": "uri"
                }
            },
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            }
        },
        "server": {
            "type": "object",
            "description": "An object representing a Server.",
            "required": [
                "url",
                "scheme"
            ],
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "properties": {
                "url": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "scheme": {
                    "type": "string",
                    "description": "The transfer protocol."
                },
                "schemeVersion": {
                    "type": "string"
                },
                "variables": {
                    "$ref": "#/definitions/serverVariables"
                },
                "baseChannel": {
                    "type": "string",
                    "x-format": "uri-path"
                },
                "security": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/SecurityRequirement"
                    }
                }
            }
        },
        "serverVariables": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/serverVariable"
            }
        },
        "serverVariable": {
            "type": "object",
            "description": "An object representing a Server Variable for server URL template substitution.",
            "minProperties": 1,
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "properties": {
                "enum": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "uniqueItems": true
                },
                "default": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "examples": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "channels": {
            "type": "object",
            "propertyNames": {
                "type": "string",
                "format": "uri-template",
                "minLength": 1
            },
            "patternProperties": {
                "": {
                    "$ref": "#/definitions/channelItem"
                }
            }
        },
        "components": {
            "type": "object",
            "description": "An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.",
            "additionalProperties": false,
            "properties": {
                "schemas": {
                    "$ref": "#/definitions/schemas"
                },
                "messages": {
                    "$ref": "#/definitions/messages"
                },
                "securitySchemes": {
                    "type": "object",
                    "patternProperties": {
                        "^[a-zA-Z0-9\\.\\-_]+$": {
                            "oneOf": [
                                {
                                    "$ref": "#/definitions/Reference"
                                },
                                {
                                    "$ref": "#/definitions/SecurityScheme"
                                }
                            ]
                        }
                    }
                },
                "parameters": {
                    "$ref": "#/definitions/parameters"
                }
            }
        },
        "schemas": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/schema"
            },
            "description": "JSON objects describing schemas the API uses."
        },
        "messages": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/message"
            },
            "description": "JSON objects describing the messages being consumed and produced by the API."
        },
        "parameters": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/parameter"
            },
            "description": "JSON objects describing re-usable channel parameters."
        },
        "schema": {
            "type": "object",
            "description": "A deterministic version of a JSON Schema object.",
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "properties": {
                "$ref": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "title": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/title"
                },
                "description": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/description"
                },
                "default": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/default"
                },
                "multipleOf": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/multipleOf"
                },
                "maximum": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/maximum"
                },
                "exclusiveMaximum": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/exclusiveMaximum"
                },
                "minimum": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/minimum"
                },
                "exclusiveMinimum": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/exclusiveMinimum"
                },
                "maxLength": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
                },
                "minLength": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
                },
                "pattern": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/pattern"
                },
                "maxItems": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
                },
                "minItems": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
                },
                "uniqueItems": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/uniqueItems"
                },
                "maxProperties": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
                },
                "minProperties": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
                },
                "required": {
                    "$ref": "http://json-schema.org/draft-04/schema#/definitions/stringArray"
                },
                "enum": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/enum"
                },
                "deprecated": {
                    "type": "boolean",
                    "default": false
                },
                "additionalProperties": {
                    "anyOf": [
                        {
                            "$ref": "#/definitions/schema"
                        },
                        {
                            "type": "boolean"
                        }
                    ],
                    "default": {}
                },
                "type": {
                    "$ref": "http://json-schema.org/draft-04/schema#/properties/type"
                },
                "items": {
                    "anyOf": [
                        {
                            "$ref": "#/definitions/schema"
                        },
                        {
                            "type": "array",
                            "minItems": 1,
                            "items": {
                                "$ref": "#/definitions/schema"
                            }
                        }
                    ],
                    "default": {}
                },
                "allOf": {
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/schema"
                    }
                },
                "oneOf": {
                    "type": "array",
                    "minItems": 2,
                    "items": {
                        "$ref": "#/definitions/schema"
                    }
                },
                "anyOf": {
                    "type": "array",
                    "minItems": 2,
                    "items": {
                        "$ref": "#/definitions/schema"
                    }
                },
                "not": {
                    "$ref": "#/definitions/schema"
                },
                "properties": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/schema"
                    },
                    "default": {}
                },
                "discriminator": {
                    "type": "string"
                },
                "readOnly": {
                    "type": "boolean",
                    "default": false
                },
                "xml": {
                    "$ref": "#/definitions/xml"
                },
                "externalDocs": {
                    "$ref": "#/definitions/externalDocs"
                },
                "example": {}
            },
            "additionalProperties": false
        },
        "xml": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "prefix": {
                    "type": "string"
                },
                "attribute": {
                    "type": "boolean",
                    "default": false
                },
                "wrapped": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "externalDocs": {
            "type": "object",
            "additionalProperties": false,
            "description": "information about external documentation",
            "required": [
                "url"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "url": {
                    "type": "string",
                    "format": "uri"
                }
            },
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            }
        },
        "channelItem": {
            "type": "object",
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "minProperties": 1,
            "properties": {
                "$ref": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "uniqueItems": true,
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/parameter"
                    }
                },
                "publish": {
                    "$ref": "#/definitions/operation"
                },
                "subscribe": {
                    "$ref": "#/definitions/operation"
                },
                "deprecated": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "parameter": {
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "properties": {
                "description": {
                    "type": "string",
                    "description": "A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed."
                },
                "name": {
                    "type": "string",
                    "description": "The name of the parameter."
                },
                "schema": {
                    "$ref": "#/definitions/schema"
                },
                "$ref": {
                    "type": "string"
                }
            }
        },
        "operation": {
            "type": "object",
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "properties": {
                "summary": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tag"
                    },
                    "uniqueItems": true
                },
                "externalDocs": {
                    "$ref": "#/definitions/externalDocs"
                },
                "operationId": {
                    "type": "string"
                },
                "message": {
                    "oneOf": [
                        {
                            "$ref": "#/definitions/message"
                        },
                        {
                            "type": "object",
                            "required": [
                                "oneOf"
                            ],
                            "additionalProperties": false,
                            "patternProperties": {
                                "^x-": {
                                    "$ref": "#/definitions/vendorExtension"
                                }
                            },
                            "properties": {
                                "oneOf": {
                                    "type": "array",
                                    "minItems": 2,
                                    "items": {
                                        "$ref": "#/definitions/message"
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        },
        "stream": {
            "title": "Stream Object",
            "type": "object",
            "additionalProperties": false,
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "minProperties": 1,
            "properties": {
                "framing": {
                    "title": "Stream Framing Object",
                    "type": "object",
                    "patternProperties": {
                        "^x-": {
                            "$ref": "#/definitions/vendorExtension"
                        }
                    },
                    "minProperties": 1,
                    "oneOf": [
                        {
                            "additionalProperties": false,
                            "properties": {
                                "type": {
                                    "type": "string",
                                    "enum": [
                                        "chunked"
                                    ]
                                },
                                "delimiter": {
                                    "type": "string",
                                    "enum": [
                                        "\\r\\n",
                                        "\\n"
                                    ],
                                    "default": "\\r\\n"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "type": {
                                    "type": "string",
                                    "enum": [
                                        "sse"
                                    ]
                                },
                                "delimiter": {
                                    "type": "string",
                                    "enum": [
                                        "\\n\\n"
                                    ],
                                    "default": "\\n\\n"
                                }
                            }
                        }
                    ]
                },
                "read": {
                    "title": "Stream Read Object",
                    "type": "array",
                    "uniqueItems": true,
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/message"
                    }
                },
                "write": {
                    "title": "Stream Write Object",
                    "type": "array",
                    "uniqueItems": true,
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/message"
                    }
                }
            }
        },
        "message": {
            "anyOf": [
                {
                    "type": "object",
                    "additionalProperties": false,
                    "patternProperties": {
                        "^x-": {
                            "$ref": "#/definitions/vendorExtension"
                        }
                    }
                },
                {
                    "type": "object",
                    "additionalProperties": false,
                    "patternProperties": {
                        "^x-": {
                            "$ref": "#/definitions/vendorExtension"
                        }
                    },
                    "properties": {
                        "schemaFormat": {
                            "type": "string"
                        },
                        "contentType": {
                            "type": "string"
                        },
                        "headers": {
                            "$ref": "#/definitions/schema"
                        },
                        "payload": {},
                        "tags": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/tag"
                            },
                            "uniqueItems": true
                        },
                        "summary": {
                            "type": "string",
                            "description": "A brief summary of the message."
                        },
                        "name": {
                            "type": "string",
                            "description": "Name of the message."
                        },
                        "title": {
                            "type": "string",
                            "description": "A human-friendly title for the message."
                        },
                        "description": {
                            "type": "string",
                            "description": "A longer description of the message. CommonMark is allowed."
                        },
                        "externalDocs": {
                            "$ref": "#/definitions/externalDocs"
                        },
                        "deprecated": {
                            "type": "boolean",
                            "default": false
                        },
                        "example": {}
                    }
                }
            ]
        },
        "vendorExtension": {
            "description": "Any property starting with x- is valid.",
            "additionalProperties": true,
            "additionalItems": true
        },
        "tag": {
            "type": "object",
            "additionalProperties": false,
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "externalDocs": {
                    "$ref": "#/definitions/externalDocs"
                }
            },
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            }
        },
        "SecurityScheme": {
            "oneOf": [
                {
                    "$ref": "#/definitions/userPassword"
                },
                {
                    "$ref": "#/definitions/apiKey"
                },
                {
                    "$ref": "#/definitions/X509"
                },
                {
                    "$ref": "#/definitions/symmetricEncryption"
                },
                {
                    "$ref": "#/definitions/asymmetricEncryption"
                },
                {
                    "$ref": "#/definitions/HTTPSecurityScheme"
                },
                {
                    "$ref": "#/definitions/oauth2Flows"
                },
                {
                    "$ref": "#/definitions/openIdConnect"
                }
            ]
        },
        "userPassword": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "userPassword"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "apiKey": {
            "type": "object",
            "required": [
                "type",
                "in"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "apiKey"
                    ]
                },
                "in": {
                    "type": "string",
                    "enum": [
                        "user",
                        "password"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "X509": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "X509"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "symmetricEncryption": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "symmetricEncryption"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "asymmetricEncryption": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "asymmetricEncryption"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "HTTPSecurityScheme": {
            "oneOf": [
                {
                    "$ref": "#/definitions/NonBearerHTTPSecurityScheme"
                },
                {
                    "$ref": "#/definitions/BearerHTTPSecurityScheme"
                },
                {
                    "$ref": "#/definitions/APIKeyHTTPSecurityScheme"
                }
            ]
        },
        "NonBearerHTTPSecurityScheme": {
            "not": {
                "type": "object",
                "properties": {
                    "scheme": {
                        "type": "string",
                        "enum": [
                            "bearer"
                        ]
                    }
                }
            },
            "type": "object",
            "required": [
                "scheme",
                "type"
            ],
            "properties": {
                "scheme": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "http"
                    ]
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "BearerHTTPSecurityScheme": {
            "type": "object",
            "required": [
                "type",
                "scheme"
            ],
            "properties": {
                "scheme": {
                    "type": "string",
                    "enum": [
                        "bearer"
                    ]
                },
                "bearerFormat": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "http"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "APIKeyHTTPSecurityScheme": {
            "type": "object",
            "required": [
                "type",
                "name",
                "in"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "httpApiKey"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "in": {
                    "type": "string",
                    "enum": [
                        "header",
                        "query",
                        "cookie"
                    ]
                },
                "description": {
                    "type": "string"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "oauth2Flows": {
            "type": "object",
            "required": [
                "type",
                "flows"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "oauth2"
                    ]
                },
                "description": {
                    "type": "string"
                },
                "flows": {
                    "type": "object",
                    "properties": {
                        "implicit": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/oauth2Flow"
                                },
                                {
                                    "required": [
                                        "authorizationUrl",
                                        "scopes"
                                    ]
                                },
                                {
                                    "not": {
                                        "required": [
                                            "tokenUrl"
                                        ]
                                    }
                                }
                            ]
                        },
                        "password": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/oauth2Flow"
                                },
                                {
                                    "required": [
                                        "tokenUrl",
                                        "scopes"
                                    ]
                                },
                                {
                                    "not": {
                                        "required": [
                                            "authorizationUrl"
                                        ]
                                    }
                                }
                            ]
                        },
                        "clientCredentials": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/oauth2Flow"
                                },
                                {
                                    "required": [
                                        "tokenUrl",
                                        "scopes"
                                    ]
                                },
                                {
                                    "not": {
                                        "required": [
                                            "authorizationUrl"
                                        ]
                                    }
                                }
                            ]
                        },
                        "authorizationCode": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/oauth2Flow"
                                },
                                {
                                    "required": [
                                        "authorizationUrl",
                                        "tokenUrl",
                                        "scopes"
                                    ]
                                }
                            ]
                        }
                    },
                    "additionalProperties": false,
                    "minProperties": 1
                }
            },
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            }
        },
        "oauth2Flow": {
            "type": "object",
            "properties": {
                "authorizationUrl": {
                    "type": "string",
                    "format": "uri"
                },
                "tokenUrl": {
                    "type": "string",
                    "format": "uri"
                },
                "refreshUrl": {
                    "type": "string",
                    "format": "uri"
                },
                "scopes": {
                    "$ref": "#/definitions/oauth2Scopes"
                }
            },
            "patternProperties": {
                "^x-": {
                    "$ref": "#/definitions/vendorExtension"
                }
            },
            "additionalProperties": false
        },
        "oauth2Scopes": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "openIdConnect": {
            "type": "object",
            "required": [
                "type",
                "openIdConnectUrl"
            ],
            "properties": {
                "type": {
                    "type": "string",
                    "enum": [
                        "openIdConnect"
                    ]
                },
                "description": {
                    "type": "string"
                },
                "openIdConnectUrl": {
                    "type": "string",
                    "format": "uri"
                }
            },
            "patternProperties": {
                "^x-": {}
            },
            "additionalProperties": false
        },
        "SecurityRequirement": {
            "type": "object",
            "additionalProperties": {
                "type": "array",
                "items": {
                    "type": "string"
                },
                "uniqueItems": true
            }
        },
        "title": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/title"
        },
        "description": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/description"
        },
        "default": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/default"
        },
        "multipleOf": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/multipleOf"
        },
        "maximum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/maximum"
        },
        "exclusiveMaximum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/exclusiveMaximum"
        },
        "minimum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/minimum"
        },
        "exclusiveMinimum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/exclusiveMinimum"
        },
        "maxLength": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
        },
        "minLength": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
        },
        "pattern": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/pattern"
        },
        "maxItems": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
        },
        "minItems": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
        },
        "uniqueItems": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/uniqueItems"
        },
        "enum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/enum"
        }
    }
}
	`)
}