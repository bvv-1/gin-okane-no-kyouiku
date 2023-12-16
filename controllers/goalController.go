package controllers

import (
	"gin-okane-no-kyouiku/models"
	"net/http"

	httputil "gin-okane-no-kyouiku/utils"

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

// GetGoal godoc
// @Summary Get goals
// @Description Get a list of goals
// @ID GetGoal
// @Tags goals
// @Accept  json
// @Produce json
// @Success 200 {object} GoalResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/v2/goals [get]
func GetGoal(c *gin.Context) {
	// モックデータを使用してレスポンスを生成
	response := GoalResponse{
		Goal: models.Goal{
			Name:  "computer",
			Point: 100,
		},
	}

	c.JSON(200, response)
}

// SetGoal godoc
// @Summary Set a goal with tasks
// @Description Set a goal with associated tasks
// @ID SetGoal
// @Tags goals
// @Accept  json
// @Produce json
// @Param goal body GoalAndTasks true "Goal and Tasks object"
// @Success 200 {string} httputil.SuccessResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/v1/goals [post]
func SetGoal(c *gin.Context) {
	// Mock request data
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

	response := httputil.SuccessResponse{Message: "Goal and tasks set successfully"}

	// Mock success response
	c.JSON(http.StatusOK, response)
}
