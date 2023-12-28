package controllers

import (
	"gin-okane-no-kyouiku/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register a user
// @Description Register a user
// @ID Register
// @Tags auth
// @Accept  json
// @Produce json
// @Param user body RegisterRequest true "RegisterRequest object"
// @Success 200 {string} utils.SuccessResponse
// @Failure 400 {object} utils.HTTPError
// @Router /api/v2/register [post]
func Register(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Invalid data format: %w", err).Error())
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse{Message: "OK"})
}
