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

func AddOrder(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	var neworder models.Order
	err := c.BindJSON(&neworder)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	logger.Info.Printf("[controllers.AddOrder] add order is succesful")

	err = service.AddOrder(neworder)
	if err != nil {
		HandleError(c, errs.ErrOrdersNotFound)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}

func GetAllOrdersByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	isDoneStr := c.Query("is_done")
	isDone, err := strconv.ParseBool(isDoneStr)
	if err != nil {
		HandleError(c, errs.ErrOrdersNotFound)
		return
	}
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetTasksByID] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, err)
		return
	}
	ord, err := service.PrintAllOrderByID(false, isDone, false, uint(userID), uint(oid))
	if err != nil {
		HandleError(c, errs.ErrOrdersNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": ord})
}

func UpdateOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateOrderByID] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	var ord models.Order
	err = c.BindJSON(&ord)
	if err != nil {
		HandleError(c, errs.ErrOrdersNotFound)
		return
	}

	err = service.UpdateOrder(ord, id)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

func ChecksOrderasDone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.CheckOrderasDone] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	err = service.CheckOrderasDone(true, id)
	if err != nil {
		HandleError(c, errs.ErrOrdersNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check as done is succesfuly"})

}
func DeleteOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteOrderByID] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, err)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	err = service.DeleteOrder(true, id)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})

}
func OrdersReport(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	isDoneStr := c.Query("is_done")
	isDone, err := strconv.ParseBool(isDoneStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	isRespStr := c.Query("is_response")
	isResp, err := strconv.ParseBool(isRespStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	ords, err := service.ReportofOrder(isDone, false, isResp, true)
	if err != nil {
		HandleError(c, errs.ErrOrdersNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders report": ords})

}
