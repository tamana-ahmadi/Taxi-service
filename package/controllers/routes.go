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

// CreateRoute
// @Summary Create Route
// @Security AKA
// @Tags routes
// @Description create new route
// @ID create-route
// @Accept json
// @Produce json
// @Param input body models.Route true "new route info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes [post]
func CreateRoute(c *gin.Context) {
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

// GetAllRoutes
// @Summary Get All Routes
// @Security AKA
// @Tags routes
// @Description get list of all routes
// @ID get-all-routes
// @Produce json
// @Param q query string false "fill if you need search"
// @Param is_response query bool false "fill if you need search"
// @Param price query int true "fill if you need search"
// @Success 200 {array} models.Route
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes [get]
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
	routes, err := service.PrintAllRoutes(isResp, false, price)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"routes": routes})
}

// GetRoutesByID
// @Summary Get Route By ID
// @Security AKA
// @Tags routes
// @Description get route by ID
// @ID get-route-by-id
// @Produce json
// @Param id path integer true "id of the route"
// @Success 200 {object} models.Route
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes/{id} [get]
func GetAllRoutesByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	rid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetAllRoutesByID] invalid route_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	route, err := service.PrintAllRouteByID(false, uint(rid))
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"route": route})
}

// UpdateRouteByID
// @Summary Update Route
// @Security AKA
// @Tags routes
// @Description update existed route
// @ID update-route
// @Accept json
// @Produce json
// @Param id path integer true "id of the route"
// @Param input body models.Route true "route update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes/{id} [put]
func UpdateRouteByID(c *gin.Context) {
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
	if urole != "driver" {
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

	err = service.UpdateRoute(route, int(userID), id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

// ChecksRouteasResponse
// @Summary Check route as response
// @Security AKA
// @Tags routes
// @Description  check as response existed route
// @ID check-route-as-response
// @Accept json
// @Produce json
// @Param id path integer true "id of the route"
// @Param input body models.Route true " check route as response info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes/{id} [patch]
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
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	err = service.CheckRouteasResponse(true, int(userID), id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check as response is succesfuly"})

}

// DeleteRouteByID
// @Summary Delete Route By ID
// @Security AKA
// @Tags routes
// @Description delete route by ID
// @ID delete-route-by-id
// @Param id path integer true "id of the route"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/routes/{id} [delete]
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
