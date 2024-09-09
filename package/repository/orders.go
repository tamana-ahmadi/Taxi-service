package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func InsertOrders(order models.Ord) error {
	err := db.GetconnectDB().Create(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.InsertOrders]error in added order %s\n", err.Error())
	}
	return nil
}
func EditOrders(routeid, clientid, driverid, taxicompid, id int) error {
	err := db.GetconnectDB().Save(&models.Ord{ID: id, ClientID: clientid, DriverID: driverid, RouteID: routeid, TaxicompID: taxicompid}).Error
	if err != nil {
		logger.Error.Printf("[repository.EditOrders]error in update order %s\n", err.Error())
	}
	return nil
}

func SoftDeleteOrders(isdeleted bool, id int) error {
	err := db.GetconnectDB().Model(&models.Ord{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[repository.SoftDeleteOrders]error in deleted order %s\n", err.Error())
	}
	return nil
}

func GetAllOrdersByID(isdeleted, isdone, isblocked bool, uid, oid uint) (order []models.Ord, err error) {
	err = db.GetconnectDB().Preload("User").Joins("Join users ON users.id=orders.user_id").Joins("Join taxicompanies ON taxicompanies.id=orders.taxicomp_id").Joins("Join routes ON routes.id=orders.route_id").Where("routes.is_deleted=?", isdeleted).Where("orders.is_done=? AND orders.is_deleted=?", isdone, isdeleted).Where("orders.id=?  AND orders.user_id=?", oid, uid).Order("orders.id").Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeleted).Find(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrdersByID]error in getting all order by id %s\n", err.Error())
		return order, err
	}
	return order, nil
}

func CheckOrdersasDone(isdone bool, id int) error {
	var order models.Ord
	err := db.GetconnectDB().Model(&order).Select("is_done").Where("id=?", id).Updates(models.Ord{IsDone: isdone}).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckOrdersasDone]error in checked order %s\n", err.Error())

	}
	return nil
}
func OrdersReport(isdone, isdeletedo, isdeletedt, isdeletedr, isresp, isblocked, isdeletedu bool) ([]models.Ord, error) {
	var order []models.Ord
	err := db.GetconnectDB().Select("taxicompanies.comp_title,orders.is_done, COUNT(DISTINCT CASE WHEN users.role='user' THEN users.id END) AS clientid, COUNT(DISTINCT CASE WHEN users.role='driver' THEN users.id END) AS driverid, COUNT(DISTINCT orders.route_id) AS routeid, SUM(routes.distance) as distance, SUM(routes.price) as price").Preload("User").Preload("Taxicomp").Preload("Route").Joins("Join taxicompanies  ON taxicompanies.id = orders.taxicomp_id").Joins("Join routes  ON routes.id=orders.route_id").Joins("Join users  ON users.id = orders.client_id OR users.id = orders.driver_id").Where("orders.is_done = ? AND orders.is_deleted = ? AND t.is_deleted = ?", isdone, isdeletedo, isdeletedt).Where("routes.is_deleted=? AND routes.is_response=?", isdeletedr, isresp).Where("users.is_blocked = ? AND users.is_deleted = ?", isblocked, isdeletedu).Group("taxicompanies.comp_title, orders.is_done").Order("clientid, driverid DESC").Find(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.OrdersReport]error in  order report %s\n", err.Error())
		return order, err
	}

	return order, nil
}
