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
