package controllers

import (
	"gin-okane-no-kyouiku/db"
	"gin-okane-no-kyouiku/models"
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

	if err := models.InsertUser(db.GetDB(), request.Email, request.Password); err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Failed to register: %w", err).Error())
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse{Message: "OK"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @ID Login
// @Tags auth
// @Accept  json
// @Produce json
// @Param user body LoginRequest true "LoginRequest object"
// @Success 200 {string} LoginResponse
// @Failure 400 {object} utils.HTTPError
// @Router /api/v2/login [post]
func Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Invalid data format: %w", err).Error())
		return
	}

	token, err := models.LoginCheck(db.GetDB(), request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, xerrors.Errorf("Failed to login: %w", err).Error())
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
