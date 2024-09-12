package controllers

import (
	"Taxi_service/models"
	"Taxi_service/package/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp
// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.SwagSignUp true "account info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		HandleError(c, err)
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

// SignIn
// @Summary SignIn
// @Tags auth
// @Description sign in to account
// @ID sign-in-to-account
// @Accept json
// @Produce json
// @Param input body models.SwagSignIn true "sign-in info"
// @Success 200 {object} accessTokenResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		HandleError(c, err)
		return
	}

	accessToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, accessTokenResponse{accessToken})
}
