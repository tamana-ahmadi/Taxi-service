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
func EditRoutes(from, into string, userid, distance, price, id int) error {
	err := db.GetconnectDB().Save(&models.Route{ID: id, From: from, Into: into, Distance: distance, Price: price, DriverID: userid}).Error
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

func GetAllRoutes(isdeleted, isresp, isblocked bool, price int, uid uint) (route []models.Route, err error) {
	err = db.GetconnectDB().Preload("User").Joins("Join users ON users.id=routes.user_id").Where("routes.is_response=? AND routes.is_deleted=? AND routes.price=?", isresp, isdeleted, price).Where("routes.user_id=?", uid).Order("routes.id").Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeleted).Find(&route).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllRoutes]error in getting all orders %s\n", err.Error())
		return route, err
	}
	return route, nil
}
func GetAllRoutesByID(isdeleted, isresp, isblocked bool, uid, rid uint) (route []models.Route, err error) {
	err = db.GetconnectDB().Preload("User").Joins("Join users ON users.id=routes.user_id").Where("routes.is_response=? AND routes.is_deleted=?", isresp, isdeleted).Where("routes.id=?  AND orders.user_id=?", rid, uid).Order("routes.id").Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeleted).Find(&route).Error
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
