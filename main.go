package main

import (
	"net/http"

	"strconv"

	"golang.org/x/xerrors"

	_ "gin-okane-no-kyouiku/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title okane no kyouiku API
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/", helloWorld)
	r.POST("/api/v1/plans/suggest", suggestDailyPlans)
	r.POST("/api/v1/plans/accept", acceptSuggestedPlans)
	r.GET("/api/v1/goals", checkGoal)
	r.GET("/api/v1/plans/check", checkProgress)
	r.GET("/api/v1/plans/today", getDailyPlans)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// helloWorld godoc
// @Summary Hello Worldのエンドポイント
// @Description GETリクエストに対して{"message": "Hello, World!"}を返す
// @ID helloWorld
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {object} httputil.HTTPError
// @Router / [get]
func helloWorld(c *gin.Context) {
	data := map[string]string{"message": "Hello, World!"}
	c.JSON(http.StatusOK, data)
}

type Task struct {
	Task  string `json:"task"`
	Point int    `json:"point"`
}

type SuggestedPlan struct {
	Day        int    `json:"day"`
	PlansToday []Task `json:"plans_today"`
}

type SuggestRequest struct {
	Goal       string `json:"goal"`
	GoalPoints int    `json:"goal_points"`
	Tasks      []Task `json:"tasks"`
}

type SuggestResponse struct {
	Plans []SuggestedPlan `json:"plans"`
}

// @Summary 日々のお手伝いプランを生成するエンドポイント
// @Description ユーザーが設定した目標とタスクに基づいて日々のお手伝いプランを生成する
// @ID suggestDailyPlans
// @Tags plans
// @Accept json
// @Produce json
// @Param request body SuggestRequest true "提案リクエストのボディ"
// @Success 200 {object} SuggestResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/v1/plans/suggest [post]
func suggestDailyPlans(c *gin.Context) {
	var request SuggestRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": xerrors.Errorf("Invalid data format: %w", err).Error()})
		return
	}

	// モックデータを使用してレスポンスを生成
	response := SuggestResponse{
		Plans: []SuggestedPlan{
			{Day: 1, PlansToday: []Task{{Task: "cleaning", Point: 5}}},
			{Day: 2, PlansToday: []Task{}},
		},
	}

	c.JSON(http.StatusOK, response)
}

type AcceptRequest struct {
	PlanID int `json:"plans_ids_id"`
	TaskID int `json:"tasks_ids_id"`
}

type OkResponse struct {
	Message string `json:"message"`
}

type HTTPError struct {
	Code         int
	WrappedError error
}

// AcceptSuggestedPlans godoc
// @Summary 提案されたデイリープランを受け入れるエンドポイント
// @Description ユーザーが提案されたデイリープランを受け入れる
// @ID acceptSuggestedPlans
// @Tags plans
// @Accept json
// @Produce json
// @Param request body AcceptRequest true "受け入れリクエストのボディ"
// @Success 200 {object} OkResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/v1/plans/accept [post]
func acceptSuggestedPlans(c *gin.Context) {
	var request AcceptRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": xerrors.Errorf("Invalid data format: %w", err).Error()})
		return
	}

	// モックデータを使用してレスポンスを生成
	response := OkResponse{
		Message: "Plan accepted",
	}

	c.JSON(http.StatusOK, response)
}

type GoalResponse struct {
	Goal       string `json:"goal"`
	GoalPoints int    `json:"goal_points"`
}

// @Summary 現在の目標を確認するエンドポイント
// @Description ユーザーの現在の目標を確認する
// @ID checkGoal
// @Tags goals
// @Accept json
// @Produce json
// @Success 200 {object} GoalResponse
// @Router /api/v1/goals [get]
func checkGoal(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	response := GoalResponse{
		Goal:       "computer",
		GoalPoints: 100,
	}

	c.JSON(http.StatusOK, response)
}
func getAdjustedPlans() []SuggestedPlan {
	return []SuggestedPlan{
		{Day: 1, PlansToday: []Task{{Task: "cleaning", Point: 5}}},
		// 他の日のプランも同様に追加
	}
}

type AdjustmentResponse struct {
	Message       string          `json:"message"`
	AdjustedPlans []SuggestedPlan `json:"adjusted_plans"`
}

// @Summary デイリープランが順調かどうかを確認するエンドポイント
// @Description ユーザーのデイリープランが順調かどうかを確認する
// @ID checkProgress
// @Tags plans
// @Accept json
// @Produce json
// @Success 200 {object} OkResponse
// @Success 200 {object} AdjustmentResponse
// @Router /api/v1/plans/check [get]
func checkProgress(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	isOnTrack := true

	if isOnTrack {
		// デイリープランが順調な場合のレスポンス
		okResponse := OkResponse{
			Message: "Plans are on track",
		}
		c.JSON(http.StatusOK, okResponse)
	} else {
		// デイリープランが調整が必要な場合のレスポンス
		adjustmentResponse := AdjustmentResponse{
			Message:       "Plans need adjustment",
			AdjustedPlans: getAdjustedPlans(),
		}
		c.JSON(http.StatusOK, adjustmentResponse)
	}
}

type DailyPlansResponse struct {
	Day        int    `json:"day"`
	PlansToday []Task `json:"plans_today"`
}

// @Summary 指定された日のデイリープランを取得するエンドポイント
// @Description ユーザーが指定した日のデイリープランを取得する
// @ID getDailyPlans
// @Tags plans
// @Accept json
// @Produce json
// @Param day query int true "取得する日の番号"
// @Success 200 {object} DailyPlansResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/v1/plans/today [get]
func getDailyPlans(c *gin.Context) {
	dayStr, ok := c.GetQuery("day")
	if !ok {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Query parameter 'day' is required").Error())
		return
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Invalid data format: %w", err).Error())
		return
	}

	// モックデータを使用してレスポンスを生成
	response := DailyPlansResponse{
		Day:        day,
		PlansToday: []Task{{Task: "cleaning", Point: 5}},
	}

	c.JSON(http.StatusOK, response)
}
