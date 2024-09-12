package controllers

import (
	"Taxi_service/errs"
	"Taxi_service/logger"
	"Taxi_service/models"
	"Taxi_service/package/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser
// @Summary Create User
// @Security ApiKeyAuth
// @Tags users
// @Description create new user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.User true "new route info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users [post]
func CreateUsers(c *gin.Context) {
	var newuser models.User
	err := c.BindJSON(&newuser)
	if err != nil {
		HandleError(c, err)
		return
	}

	err = service.CreateUser(newuser)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}

// UpdateUserByID
// @Summary Update User
// @Security ApiKeyAuth
// @Tags users
// @Description update existed user
// @ID update-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.User true "user update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes/{id} [put]
func EditUsers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, err)
		return
	}

	var user models.User
	err = c.BindJSON(&user)
	if err != nil {
		HandleError(c, err)
		return
	}

	err = service.UpdateUser(user, id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

// EditUsersRating
// @Summary edit user rating
// @Security ApiKeyAuth
// @Tags users
// @Description  edit  rating existed user
// @ID edit-users-rating
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.User true "edit users rating"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users/{id} [patch]
func EditUsersRating(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, err)
		return
	}
	var user models.User
	err = c.BindJSON(&user)
	if err != nil {
		HandleError(c, err)
		return
	}
	err = service.UpdateUserRating(user, id)
	if err != nil {
		logger.Error.Printf("[controllers.EditUsersRating] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Edit user`s password  is succesfuly"})
}

// BlockUsers
// @Summary Block User By ID
// @Security ApiKeyAuth
// @Tags users
// @Description block user by ID
// @ID block-user-by-id
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users/{id} [delete]
func BlockUsers(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUsers] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, err)
		return
	}
	err = service.IsBlockedUser(true, id)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blocked is succesfuly"})

}

// DeleteUsers
// @Summary Delete User By ID
// @Security ApiKeyAuth
// @Tags users
// @Description delete user by ID
// @ID delete-user-by-id
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users/{id} [delete]
func DeleteUsers(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUsers] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, err)
		return
	}
	err = service.IsDeletedUser(true, id)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})

}

// PrintUsers
// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags users
// @Description get list of all users
// @ID get-all-users
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users [get]
func PrintUsers(c *gin.Context) {
	role := c.Query("role")
	logger.Info.Printf("Client with ip: [%s] requested list of users\n", c.ClientIP())
	users, err := service.PrintAllUsers(false, false, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
	logger.Info.Printf("Client with ip: [%s] got list of users\n", c.ClientIP())
}

// PrintUsersByID
// @Summary Get user By ID
// @Security ApiKeyAuth
// @Tags users
// @Description get user by ID
// @ID get-user-by-id
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users/{id} [get]
func PrintUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.PrintUsersByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, err)
		return
	}
	user, err := service.PrintAllUsersByID(false, false, id)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
