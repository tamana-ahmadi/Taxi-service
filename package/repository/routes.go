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

func GetAllRoutes(isresp, isdeleted bool, price int) (route []models.Route, err error) {
	err = db.GetconnectDB().Raw("Select r.from, r.into, r.distance, r.pricekm,(r.distance * r.pricekm) as all_price, r.client_id, r.driver_id,r.is_response FROM routes r Where r.is_response=? AND r.is_deleted=? AND all_price<=?", isresp, isdeleted, price).Scan(&route).Error
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
