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

// CreateTaxicomp
// @Summary Create Taxicomp
// @Security AKA
// @Tags taxi companies
// @Description create new taxi company
// @ID create-taxi-company
// @Accept json
// @Produce json
// @Param input body models.TaxiComp true "new taxicomp info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/taxicomps [post]
func CreateTaxicomp(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	var newtxcm models.TaxiComp
	err := c.BindJSON(&newtxcm)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	logger.Info.Printf("[controllers.AddTaxiComp] add taxi company is succesful")

	err = service.AddTaxicomp(newtxcm)
	if err != nil {
		HandleError(c, errs.ErrTaxicompsNotFound)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}

// GetAllTaxiComp
// @Summary Get All Taxi Companies
// @Security AKA
// @Tags taxi companies
// @Description get list of all taxi companies
// @ID get-all-taxi-companies
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.TaxiComp
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/taxicomps [get]
func GetAllTaxiComp(c *gin.Context) {
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
	if urole != "admin" && urole != "driver" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	txcms, err := service.PrintAllTaxiComps(false)
	if err != nil {
		HandleError(c, errs.ErrTaxicompsNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"taxicompanies": txcms})
}

// GetAllTaxiCompByID
// @Summary Get Taxi Company By ID
// @Security AKA
// @Tags taxi companies
// @Description get taxi company by ID
// @ID get-taxi-company-by-id
// @Produce json
// @Param id path integer true "id of the taxicomp"
// @Success 200 {object} models.TaxiComp
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/taxicomps/{id} [get]
func GetAllTaxiCompByID(c *gin.Context) {
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
	if urole != "admin" && urole != "driver" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetAllTaxiCompByID] invalid taxi company_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	txcm, err := service.PrintAllTaxiCompByID(false, id)
	if err != nil {
		HandleError(c, errs.ErrTaxicompsNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"taxicompany": txcm})
}

// UpdateTaxiCompByID
// @Summary Update taxi Company
// @Security AKA
// @Tags taxi companies
// @Description update existed taxi company
// @ID update-taxi-company
// @Accept json
// @Produce json
// @Param id path integer true "id of the taxi company"
// @Param input body models.TaxiComp true "taxi company update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/taxicomps/{id} [put]
func UpdateTaxiCompByID(c *gin.Context) {
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
		logger.Error.Printf("[controllers.UpdateTaxiCompByID] invalid taxi company_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	var txcm models.TaxiComp
	err = c.BindJSON(&txcm)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.UpdateTaxiComp(txcm, id)
	if err != nil {
		HandleError(c, errs.ErrTaxicompsNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

// DeleteTaxiCompByID
// @Summary Delete Taxi Company By ID
// @Security AKA
// @Tags taxi companies
// @Description delete taxi company by ID
// @ID delete-taxi-company-by-id
// @Param id path integer true "id of the taxi company"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/taxicomps/{id} [delete]
func DeleteTaxiCompByID(c *gin.Context) {
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
		logger.Error.Printf("[controllers.DeleteTaxiCompByID] invalid taxicompanies_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	err = service.DeleteTaxiComp(true, id)
	if err != nil {
		HandleError(c, errs.ErrTaxicompsNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})

}
