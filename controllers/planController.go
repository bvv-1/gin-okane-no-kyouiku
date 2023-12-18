package controllers

import (
	"fmt"
	"gin-okane-no-kyouiku/models"
	"gin-okane-no-kyouiku/utils"
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

type ProgressRequest struct {
	Day          int             `json:"day"`
	TaskProgress []TaskAndStatus `json:"task_progress"`
}

type TaskAndStatus struct {
	Task   models.Task `json:"task"`
	IsDone bool        `json:"is_done"`
}

// @Summary 日々のお手伝いプランを生成するエンドポイント
// @Description ユーザーが設定した目標とタスクに基づいて日々のお手伝いプランを生成する
// @ID suggestDailyPlans
// @Tags plans
// @Accept json
// @Produce json
// @Param request body SuggestRequest true "提案リクエストのボディ"
// @Success 200 {object} SuggestResponse
// @Failure 400 {object} utils.HTTPError
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

// GetPlans godoc
// @Summary Get plans
// @Description Get a list of plans
// @ID GetPlans
// @Tags plans
// @Accept  json
// @Produce json
// @Success 200 {array} models.Plan
// @Failure 400 {object} utils.HTTPError
// @Router /api/v1/plans [get]
func GetPlans(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	response := []models.Plan{
		{Day: 1, TasksToday: []models.Task{{Name: "Task 1", Point: 5}, {Name: "Task 2", Point: 10}}},
		{Day: 2, TasksToday: []models.Task{{Name: "Task 3", Point: 15}, {Name: "Task 2", Point: 10}}},
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

// AcceptSuggestedPlans godoc
// @Summary Accept suggested plans
// @Description Accept the suggested plans and update the status to "inprogress"
// @ID AcceptSuggestedPlans
// @Tags plans
// @Accept json
// @Produce json
// @Success 200 {string} utils.SuccessResponse
// @Router /api/v1/plans/suggested [put]
func AcceptSuggestedPlans(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	response := utils.SuccessResponse{Message: "Suggested plans accepted"}

	c.JSON(http.StatusOK, response)
}

// @Summary 指定された日のデイリープランを取得するエンドポイント
// @Description ユーザーが指定した日のデイリープランを取得する
// @ID GetTodayPlan
// @Tags plans
// @Accept json
// @Produce json
// @Param day query int true "取得する日の番号"
// @Success 200 {array} models.Plan
// @Failure 400 {object} utils.HTTPError
// @Router /api/v2/plans/today [get]
func GetTodayPlan(c *gin.Context) {
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
	response := models.Plan{
		Day:        day,
		TasksToday: []models.Task{{Name: "Task 1", Point: 5}, {Name: "Task 2", Point: 10}},
	}

	c.JSON(http.StatusOK, response)
}

// SubmitTodayProgress godoc
// @Summary Submit progress for today's plan
// @Description Submit the progress of tasks for today's plan and store in the database
// @ID SubmitTodayProgress
// @Tags plans
// @Accept json
// @Produce json
// @Param progress body ProgressRequest true "Progress request object"
// @Success 200 {string} utils.SuccessResponse
// @Router /api/v2/plans/today [post]
func SubmitTodayProgress(c *gin.Context) {
	var progressRequest ProgressRequest

	if err := c.ShouldBindJSON(&progressRequest); err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Invalid data format: %w", err).Error())
		return
	}

	// progressRequest.TaskProgressをfor
	for _, taskAndStatus := range progressRequest.TaskProgress {
		if taskAndStatus.IsDone {
			fmt.Printf("Task: %s, IsDone: %t\n", taskAndStatus.Task.Name, taskAndStatus.IsDone)
		}
	}

	// モックデータを使用してレスポンスを生成
	response := utils.SuccessResponse{Message: "Progress submitted"}

	c.JSON(http.StatusOK, response)
}
