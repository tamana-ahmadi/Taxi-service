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
	err := db.GetconnectDB().Where("id=?", id).Save(&models.Route{ID: id, From: from, Into: into, Distance: distance, Pricekm: price, IsResponse: isresp, DriverID: userid}).Error
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

func GetAllRoutes(isresp, isdeletedr, isblocked, isdeletedu bool, price int) (route []models.GetRoutes, err error) {
	err = db.GetconnectDB().Raw("Select r.from, r.into, r.is_response, COUNT(DISTINCT r.distance) as distance ,COUNT(DISTINCT r.pricekm) as pricekm, MAX(r.all_price) as all_price, COUNT(DISTINCT CASE WHEN u.role='user' THEN u.id END) as client_id, COUNT(DISTINCT CASE WHEN u.role='driver' THEN u.id END) as driver_id FROM routes r, users u Where r.client_id=u.id OR r.driver_id=u.id AND r.is_response=? AND r.is_deleted=? AND u.is_blocked=? AND u.is_deleted=? AND all_price<=? GROUP BY r.from,r.into,r.is_response ORDER BY distance,all_price DESC", isresp, isdeletedr, isblocked, isdeletedu, price).Scan(&route).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllRoutes]error in getting all routes %s\n", err.Error())
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
