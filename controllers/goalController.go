package controllers

import (
	"gin-okane-no-kyouiku/models"

	"github.com/gin-gonic/gin"
)

type GoalResponse struct {
	Goal models.Goal `json:"goal"`
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
