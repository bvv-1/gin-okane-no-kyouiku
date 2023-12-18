package controllers

import (
	"gin-okane-no-kyouiku/models"
	"gin-okane-no-kyouiku/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type GoalResponse struct {
	Goal models.Goal `json:"goal"`
}

type GoalAndTasks struct {
	Goal  models.Goal   `json:"goal"`
	Tasks []models.Task `json:"tasks"`
}

type ProgressResponse struct {
	Goal       models.Goal `json:"goal"`
	TotalPoint int         `json:"total_point"`
	OnTrack    bool        `json:"on_track"`
}

// GetGoal godoc
// @Summary Get goals
// @Description Get a list of goals
// @ID GetGoal
// @Tags goals
// @Accept  json
// @Produce json
// @Success 200 {object} GoalResponse
// @Failure 400 {object} utils.HTTPError
// @Router /api/v2/goals [get]
func GetGoal(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	response := GoalResponse{
		Goal: models.Goal{
			Name:  "computer",
			Point: 100,
		},
	}

	c.JSON(http.StatusOK, response)
}

// SetGoal godoc
// @Summary Set a goal with tasks
// @Description Set a goal with associated tasks
// @ID SetGoal
// @Tags goals
// @Accept  json
// @Produce json
// @Param goal body GoalAndTasks true "Goal and Tasks object"
// @Success 200 {string} utils.SuccessResponse
// @Failure 400 {object} utils.HTTPError
// @Router /api/v1/goals [post]
func SetGoalAndTasks(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	var goalAndTasks GoalAndTasks
	if err := c.ShouldBindJSON(&goalAndTasks); err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Invalid data format: %w", err).Error())
		return
	}

	// Mock DB insertion (assuming models.SetGoalWithTasks is a function in your models package)
	// err := models.SetGoalWithTasks(goalAndTasks.Goal, goalAndTasks.Tasks)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to set goal and tasks in the database"})
	// 	return
	// }

	response := utils.SuccessResponse{Message: "Goal and tasks set successfully"}

	// Mock success response
	c.JSON(http.StatusOK, response)
}

// CheckProgress godoc
// @Summary Check the progress of a goal
// @Description Get the goal details, accumulated points, and whether it's on track
// @ID CheckProgress
// @Tags goals
// @Produce json
// @Success 200 {object} ProgressResponse
// @Router /api/v1/goals/progress [get]
func CheckProgress(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	goal := models.Goal{Name: "My Goal", Point: 100}
	totalPoints := 75
	onTrack := true

	response := ProgressResponse{
		Goal:       goal,
		TotalPoint: totalPoints,
		OnTrack:    onTrack,
	}

	c.JSON(http.StatusOK, response)
}
