package controllers

import (
	"Taxi_service/errs"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameorPassword) ||
		errors.Is(err, errs.ErrOrdersNotFound) ||
		errors.Is(err, errs.ErrRoutesNotFound) ||
		errors.Is(err, errs.ErrTaxicompsNotFound) ||
		errors.Is(err, errs.ErrValidationFailed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong})
	}
}
