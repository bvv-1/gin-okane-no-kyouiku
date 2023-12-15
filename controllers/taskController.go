package controllers

import (
	"gin-okane-no-kyouiku/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type SuggestRequest struct {
	Goal       string        `json:"goal"`
	GoalPoints int           `json:"goal_points"`
	Tasks      []models.Task `json:"tasks"`
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
			{Day: 1, PlansToday: []models.Task{{Task: "cleaning", Point: 5}}},
			{Day: 2, PlansToday: []models.Task{}},
		},
	}

	c.JSON(http.StatusOK, response)
}
