package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func InsertRoutes(route models.Route) error {
	err := db.GetconnectDB().Create(&route).Error
	if err != nil {
		logger.Error.Printf("[repository.InsertRoutes]error in added route %s\n", err.Error())
	}
	return nil
}
func EditRoutes(from, into string, distance, price int, isresp bool, userid int, id int) error {
	err := db.GetconnectDB().Where("id=?", id).Save(&models.Route{ID: id, From: from, Into: into, Distance: distance, Price: price, IsResponse: isresp, DriverID: userid}).Error
	if err != nil {
		logger.Error.Printf("[repository.EditRoutes]error in update route %s\n", err.Error())
	}
	return nil
}

func SoftDeleteRoutes(isdeleted bool, id int) error {
	err := db.GetconnectDB().Model(&models.Route{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[repository.SofDeleteRoutes]error in deleted route %s\n", err.Error())
	}
	return nil
}

func GetAllRoutes(isresp, isdeleted bool, price int) (route []models.Route, err error) {
	err = db.GetconnectDB().Where("routes.is_response=? AND routes.is_deleted=? AND price<=?", isresp, isdeleted, price).Find(&route).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllRoutes]error in getting all orders %s\n", err.Error())
		return route, err
	}
	return route, nil
}
func GetAllRoutesByID(isdeleted bool, id uint) (route []models.Route, err error) {
	err = db.GetconnectDB().Where("routes.is_deleted=?", isdeleted).Where("routes.id=?", id).Find(&route).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllRoutesByID]error in getting all order by id %s\n", err.Error())
		return route, err
	}
	return route, nil
}

func CheckRoutesAsResponse(isresp bool, cid, id int) error {
	err := db.GetconnectDB().Model(&models.Route{}).Select("is_response", "client_id").Where("id=?", id).Updates(models.Route{IsResponse: isresp, ClientID: cid}).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckRoutesAsResponse]error in checked route %s\n", err.Error())

	}
	return nil
}

func OrdersReport(isresp, isdeletedr, isblocku, isdeletedu, isdeletedt bool) (report []models.OrdersReport, err error) {
	err = db.GetconnectDB().Raw(`SELECT t.comp_title as comp_title, r.is_response as is_response,
	COUNT(DISTINCT CASE WHEN u.role='user' THEN u.id END) as count_clients,
	COUNT(DISTINCT CASE WHEN u.role='driver' THEN u.id END) as count_drivers
	SUM(r.distance*r.price) as incomes
	FROM routes r, taxicompanies t,users u
	WHERE r.client_id=u.id AND t.user_id=u.id AND 
	r.is_response=? AND r.is_deleted=?
	u.is_blocked=? AND u.is_deleted=? 
	AND t.is_deleted=?
	GROUP BY comp_title,is_response
	 ORDER BY incomes DESC`, isresp, isdeletedr, isblocku, isdeletedu, isdeletedt).Scan(&report).Error
	if err != nil {
		logger.Error.Printf("[repository.OrdersReport]error in checked route %s\n", err.Error())
		return report, err
	}
	return report, nil
}
