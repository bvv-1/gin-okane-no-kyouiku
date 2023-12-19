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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "GETリクエストに対して{\"message\": \"Hello, World!\"}を返す",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Hello Worldのエンドポイント",
                "operationId": "helloWorld",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/goals": {
            "post": {
                "description": "Set a goal with associated tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goals"
                ],
                "summary": "Set a goal with tasks",
                "operationId": "SetGoal",
                "parameters": [
                    {
                        "description": "Goal and Tasks object",
                        "name": "goal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GoalAndTasks"
                        }
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
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/goals/progress": {
            "get": {
                "description": "Get the goal details, accumulated points, and whether it's on track",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goals"
                ],
                "summary": "Check the progress of a goal",
                "operationId": "CheckProgress",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ProgressResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/plans": {
            "get": {
                "description": "Get a list of plans",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Get plans",
                "operationId": "GetPlans",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Plan"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/plans/suggested": {
            "get": {
                "description": "Get a list of suggested plans",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Get suggested plans",
                "operationId": "GetSuggestedPlans",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Plan"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Accept the suggested plans and update the status to \"inprogress\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Accept suggested plans",
                "operationId": "AcceptSuggestedPlans",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/goals": {
            "get": {
                "description": "Get a list of goals",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goals"
                ],
                "summary": "Get goals",
                "operationId": "GetGoal",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GoalResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v2/plans/suggest": {
            "post": {
                "description": "ユーザーが設定した目標とタスクに基づいて日々のお手伝いプランを生成する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "日々のお手伝いプランを生成するエンドポイント",
                "operationId": "suggestDailyPlans",
                "parameters": [
                    {
                        "description": "提案リクエストのボディ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SuggestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuggestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v2/plans/today": {
            "get": {
                "description": "ユーザーが指定した日のデイリープランを取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "指定された日のデイリープランを取得するエンドポイント",
                "operationId": "GetTodayPlan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "取得する日の番号",
                        "name": "day",
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
                                "$ref": "#/definitions/models.Plan"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Submit the progress of tasks for today's plan and store in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Submit progress for today's plan",
                "operationId": "SubmitTodayProgress",
                "parameters": [
                    {
                        "description": "Progress request object",
                        "name": "progress",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ProgressRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.GoalAndTasks": {
            "type": "object",
            "properties": {
                "goal": {
                    "$ref": "#/definitions/models.Goal"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Task"
                    }
                }
            }
        },
        "controllers.GoalResponse": {
            "type": "object",
            "properties": {
                "goal": {
                    "$ref": "#/definitions/models.Goal"
                }
            }
        },
        "controllers.ProgressRequest": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "task_progress": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.TaskAndStatus"
                    }
                }
            }
        },
        "controllers.ProgressResponse": {
            "type": "object",
            "properties": {
                "goal": {
                    "$ref": "#/definitions/models.Goal"
                },
                "on_track": {
                    "type": "boolean"
                },
                "total_point": {
                    "type": "integer"
                }
            }
        },
        "controllers.SuggestRequest": {
            "type": "object",
            "properties": {
                "goal": {
                    "$ref": "#/definitions/models.Goal"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Task"
                    }
                }
            }
        },
        "controllers.SuggestResponse": {
            "type": "object",
            "properties": {
                "plans": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SuggestedPlan"
                    }
                }
            }
        },
        "controllers.TaskAndStatus": {
            "type": "object",
            "properties": {
                "is_done": {
                    "type": "boolean"
                },
                "task": {
                    "$ref": "#/definitions/models.Task"
                }
            }
        },
        "models.Goal": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                }
            }
        },
        "models.Plan": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "tasks_today": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Task"
                    }
                }
            }
        },
        "models.SuggestedPlan": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "plans_today": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Task"
                    }
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "utils.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "okane no kyouiku API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
