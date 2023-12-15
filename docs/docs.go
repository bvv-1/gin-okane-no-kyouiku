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
                "tags": [
                    "hello"
                ],
                "summary": "Hello Worldのエンドポイント",
                "operationId": "helloWorld",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/goals": {
            "get": {
                "description": "ユーザーの現在の目標を確認する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goals"
                ],
                "summary": "現在の目標を確認するエンドポイント",
                "operationId": "checkGoal",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.GoalResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/plans/accept": {
            "post": {
                "description": "ユーザーが提案されたデイリープランを受け入れる",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "提案されたデイリープランを受け入れるエンドポイント",
                "operationId": "acceptSuggestedPlans",
                "parameters": [
                    {
                        "description": "受け入れリクエストのボディ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.AcceptRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.OkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/plans/check": {
            "get": {
                "description": "ユーザーのデイリープランが順調かどうかを確認する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "デイリープランが順調かどうかを確認するエンドポイント",
                "operationId": "checkProgress",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.AdjustmentResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/plans/submit": {
            "post": {
                "description": "ユーザーがデイリータスクデータを提出する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "submit"
                ],
                "summary": "デイリータスクデータを提出するエンドポイント",
                "operationId": "submitDailyTasks",
                "parameters": [
                    {
                        "description": "提出リクエストのボディ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SubmitRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.OkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/plans/today": {
            "post": {
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
                "operationId": "getDailyPlansOld",
                "parameters": [
                    {
                        "description": "取得する日の番号",
                        "name": "day",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.GetDailyPlansRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.DailyPlansResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/points": {
            "get": {
                "description": "ユーザーの現在のポイントを取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "points"
                ],
                "summary": "ユーザーのポイントを取得するエンドポイント",
                "operationId": "getUserPoints",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.PointsResponse"
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
                            "$ref": "#/definitions/httputil.HTTPError"
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
                "operationId": "getDailyPlans",
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
                            "$ref": "#/definitions/main.DailyPlansResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.SuggestRequest": {
            "type": "object",
            "properties": {
                "goal": {
                    "type": "string"
                },
                "goal_points": {
                    "type": "integer"
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
        "httputil.HTTPError": {
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
        "main.AcceptRequest": {
            "type": "object",
            "properties": {
                "plans_ids_id": {
                    "type": "integer"
                },
                "tasks_ids_id": {
                    "type": "integer"
                }
            }
        },
        "main.AdjustmentResponse": {
            "type": "object",
            "properties": {
                "adjusted_plans": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SuggestedPlan"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "main.DailyPlansResponse": {
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
        "main.GetDailyPlansRequest": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                }
            }
        },
        "main.GoalResponse": {
            "type": "object",
            "properties": {
                "goal": {
                    "type": "string"
                },
                "goal_points": {
                    "type": "integer"
                }
            }
        },
        "main.OkResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "main.PointsResponse": {
            "type": "object",
            "properties": {
                "points": {
                    "type": "integer"
                }
            }
        },
        "main.SubmitRequest": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "total_points": {
                    "type": "integer"
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
                "point": {
                    "type": "integer"
                },
                "task": {
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
