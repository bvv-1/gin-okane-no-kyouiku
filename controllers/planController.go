package controllers

import (
	"gin-okane-no-kyouiku/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type SuggestRequest struct {
	Goal  models.Goal   `json:"goal"`
	Tasks []models.Task `json:"tasks"`
}

type SuggestResponse struct {
	Plans []models.SuggestedPlan `json:"plans"`
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
// @Router /api/v2/plans/suggest [post]
func SuggestDailyPlans(c *gin.Context) {
	var request SuggestRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": xerrors.Errorf("Invalid data format: %w", err).Error()})
		return
	}

	// モックデータを使用してレスポンスを生成
	response := SuggestResponse{
		Plans: []models.SuggestedPlan{
			{Day: 1, PlansToday: []models.Task{{Name: "cleaning", Point: 5}}},
			{Day: 2, PlansToday: []models.Task{}},
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetSuggestedPlans godoc
// @Summary Get suggested plans
// @Description Get a list of suggested plans
// @ID GetSuggestedPlans
// @Tags plans
// @Produce json
// @Success 200 {array} models.Plan
// @Router /api/v1/plans/suggested [get]
func GetSuggestedPlans(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	response := []models.Plan{
		{Day: 1, TasksToday: []models.Task{{Name: "Task 1", Point: 5}, {Name: "Task 2", Point: 10}}},
		{Day: 2, TasksToday: []models.Task{{Name: "Task 3", Point: 15}, {Name: "Task 2", Point: 10}}},
	}

	c.JSON(200, response)
}

// @Summary 指定された日のデイリープランを取得するエンドポイント
// @Description ユーザーが指定した日のデイリープランを取得する
// @ID GetTodayPlans
// @Tags plans
// @Accept json
// @Produce json
// @Param day query int true "取得する日の番号"
// @Success 200 {object} models.DailyPlansResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/v2/plans/today [get]
func GetTodayPlans(c *gin.Context) {
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
	response := models.DailyPlansResponse{
		Day:        day,
		PlansToday: []models.Task{{Name: "cleaning", Point: 5}},
	}

	c.JSON(http.StatusOK, response)
}
