package main

import (
	"net/http"

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
// @BasePath /api/v1
func main() {
	r := gin.Default()

	r.GET("/", helloWorld)
	r.POST("/api/v1/plans/suggest", suggestDailyPlans)

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
// @Success 200 {object} map[string]string
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
// @Failure 400 {object} map[string]string
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
