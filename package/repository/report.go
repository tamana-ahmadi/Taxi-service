package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func OrdersReport(isresp, isdeletedr, isblocku, isdeletedu, isdeletedt bool) (report []models.Route, err error) {
	err = db.GetconnectDB().Select(`taxicompanies.comp_title as comp_title, routes.is_response as is_response,
	COUNT(DISTINCT CASE WHEN users.role='user' THEN users.id END) as count_clients,
	COUNT(DISTINCT CASE WHEN users.role='driver' THEN users.id END) as count_drivers,
	SUM(routes.distance*routes.price) as incomes`).Joins("Join taxicompanies.user_id=users.id").Joins("Join routes.driver_id=users.id").Where(`routes.is_response=? AND routes.is_deleted=?
	AND users.is_blocked=? AND users.is_deleted=? 
	AND taxicompanies.is_deleted=?`, isresp, isdeletedr, isblocku, isdeletedu, isdeletedt).Group("comp_title,is_response").Order("incomes DESC").Find(&report).Error
	if err != nil {
		logger.Error.Printf("[repository.OrdersReport]error in  orders report %s\n", err.Error())
		return report, err
	}
	return report, nil
}
