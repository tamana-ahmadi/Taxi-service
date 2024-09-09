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

func AddRoute(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "driver" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	var newroute models.Route
	err := c.BindJSON(&newroute)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	newroute.DriverID = int(userID)
	logger.Info.Printf("[controllers.AddRoute] add route is succesful")

	err = service.AddRoute(newroute)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}

func GetAllRoutes(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}

	isRespStr := c.Query("is_response")
	isResp, err := strconv.ParseBool(isRespStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	priceStr := c.Query("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	routes, err := service.PrintAllRoutes(false, isResp, false, price, uint(userID))
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"routes": routes})
}

func GetAllRoutesByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	isRespStr := c.Query("is_response")
	isResp, err := strconv.ParseBool(isRespStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	rid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetAllRoutesByID] invalid route_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	route, err := service.PrintAllRouteByID(false, isResp, false, uint(userID), uint(rid))
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": route})
}

func UpdateRouteByID(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "driver" && urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateRouteByID] invalid route_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	var route models.Route
	err = c.BindJSON(&route)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.UpdateRoute(route, id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

func ChecksRouteasResponse(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "user" && urole != "driver" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.ChecksRouteasResponse] invalid route_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	var checkroute models.Route
	err = c.BindJSON(&checkroute)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	err = service.CheckRouteasResponse(true, checkroute, id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check as response is succesfuly"})

}
func DeleteRouteByID(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "driver" && urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteRouteByID] invalid route_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	err = service.DeleteRoute(true, id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})
}
