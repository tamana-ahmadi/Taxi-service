package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func OrdersReport(isresp, isdeletedr, isblocku, isdeletedu bool) (report []models.Route, err error) {
	err = db.GetconnectDB().Raw(`Select routes.from ,routes.into, routes.is_response,
	COUNT(DISTINCT CASE WHEN users.role='user' THEN users.id END) as client_id,
	COUNT(DISTINCT CASE WHEN users.role='driver' THEN users.id END) as driver_id,
	SUM(routes.distance*routes.price) as incomes FROM routes, users WHERE users.id=routes.driver_id AND routes.is_response=? AND routes.is_deleted=?
	AND users.is_blocked=? AND users.is_deleted=? GROUP BY routes.from,routes.into,routes.is_response ORDER BY incomes DESC`, isresp, isdeletedr, isblocku, isdeletedu).Scan(&report).Error
	if err != nil {
		logger.Error.Printf("[repository.OrdersReport]error in  orders report %s\n", err.Error())
		return report, err
	}
	return report, nil
}
