// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Pivotal Tracker to JIRA synchronization service.",
    "title": "Tracker 2 JIRA",
    "termsOfService": "https://github.com/king-jam/tracker2jira/blob/master/docs/TOS.md",
    "contact": {
      "email": "james.r.king4@gmail.com"
    },
    "license": {
      "name": "Mozilla Public License 2.0",
      "url": "https://www.mozilla.org/en-US/MPL/2.0/"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "general"
        ],
        "summary": "Returns the default home page",
        "operationId": "root",
        "responses": {
          "200": {
            "description": "The default home page for the application"
          }
        }
      }
    },
    "/projects": {
      "get": {
        "tags": [
          "projects"
        ],
        "summary": "Returns all the projects.",
        "operationId": "getProjects",
        "responses": {
          "200": {
            "description": "The list of current projects",
            "schema": {
              "$ref": "#/definitions/Projects"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      },
      "post": {
        "description": "Post a new project config",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "projects"
        ],
        "summary": "Adds a project configuration",
        "operationId": "postProject",
        "parameters": [
          {
            "description": "Project definition",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Project"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    },
    "/projects/{projectID}": {
      "get": {
        "description": "getting project object",
        "produces": [
          "application/json"
        ],
        "tags": [
          "projects"
        ],
        "summary": "gets the project from ID",
        "operationId": "getProjectByID",
        "parameters": [
          {
            "type": "string",
            "description": "ID of project to return",
            "name": "projectID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Project"
            }
          },
          "404": {
            "description": "Project not found"
          }
        }
      }
    },
    "/tasks": {
      "get": {
        "tags": [
          "projects"
        ],
        "summary": "Returns all the tasks.",
        "operationId": "getTasks",
        "responses": {
          "200": {
            "description": "The list of current tasks",
            "schema": {
              "$ref": "#/definitions/Tasks"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      },
      "post": {
        "description": "Post a new task config",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "projects"
        ],
        "summary": "Adds a task configuration",
        "operationId": "postTask",
        "parameters": [
          {
            "description": "Task definition",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    },
    "/tasks/{taskID}": {
      "get": {
        "description": "getting task object",
        "produces": [
          "application/json"
        ],
        "tags": [
          "projects"
        ],
        "summary": "gets the task from ID",
        "operationId": "getTaskByID",
        "parameters": [
          {
            "type": "string",
            "description": "ID of task to return",
            "name": "taskID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "404": {
            "description": "Task not found"
          }
        }
      }
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Returns all the users.",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "The list of current users",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      },
      "post": {
        "description": "Post a new user config",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Adds a new user configuration",
        "operationId": "postUser",
        "parameters": [
          {
            "description": "User definition",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    },
    "/users/{userID}": {
      "get": {
        "description": "getting user object",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "gets the user from ID",
        "operationId": "getUserByID",
        "parameters": [
          {
            "type": "string",
            "description": "ID of user to return",
            "name": "userID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "general"
        ],
        "summary": "Returns the current version running.",
        "operationId": "version",
        "responses": {
          "200": {
            "description": "The current version of the service",
            "schema": {
              "$ref": "#/definitions/Version"
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
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "LoginRequest": {
      "description": "Login Request Body",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "Project": {
      "type": "object",
      "properties": {
        "adminUserID": {
          "type": "string"
        },
        "externalID": {
          "type": "string"
        },
        "projectID": {
          "type": "string"
        },
        "projectOverrides": {
          "type": "object"
        },
        "projectType": {
          "type": "string"
        },
        "projectURL": {
          "type": "string"
        }
      }
    },
    "Projects": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Project"
      }
    },
    "Task": {
      "type": "object",
      "properties": {
        "currentStateMap": {
          "type": "object"
        },
        "master": {
          "type": "string"
        },
        "slave": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "storyFieldMap": {
          "type": "object"
        },
        "storyTypeMap": {
          "type": "object"
        },
        "taskID": {
          "type": "string"
        }
      }
    },
    "Tasks": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Task"
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "externalCredentials": {
          "type": "object"
        },
        "userID": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "Users": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/User"
      }
    },
    "Version": {
      "type": "object",
      "properties": {
        "buildDate": {
          "type": "string"
        },
        "commitHash": {
          "type": "string"
        },
        "releaseVersion": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "tags": [
    {
      "description": "General Application Settings \u0026 Information",
      "name": "general",
      "externalDocs": {
        "description": "Documentation",
        "url": "https://github.com/king-jam/tracker2jira/blob/master/docs/README.md"
      }
    },
    {
      "description": "Project configurations that are under management",
      "name": "projects"
    },
    {
      "description": "Operations about project members",
      "name": "users"
    }
  ]
}`))
}
