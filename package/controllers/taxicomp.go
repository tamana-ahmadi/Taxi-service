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

func AddTaxicomp(c *gin.Context) {
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
