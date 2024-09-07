package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func InsertOrders(order models.Order) error {
	err := db.GetconnectDB().Create(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.InsertOrders]error in added order %s\n", err.Error())
	}
	return nil
}
func EditOrders(routeid, clientid, driverid, taxicompid, id int) error {
	err := db.GetconnectDB().Save(&models.Order{ID: id, ClientID: clientid, DriverID: driverid, RouteID: routeid, TaxicompID: taxicompid}).Error
	if err != nil {
		logger.Error.Printf("[repository.EditOrders]error in update order %s\n", err.Error())
	}
	return nil
}

func SoftDeleteOrders(isdeleted bool, id int) error {
	err := db.GetconnectDB().Model(&models.Order{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[repository.SoftDeleteOrders]error in deleted order %s\n", err.Error())
	}
	return nil
}

func GetAllOrdersByID(isdeleted, isdone, isblocked bool, uid, oid uint) (order []models.Order, err error) {
	err = db.GetconnectDB().Preload("User").Joins("Join users ON users.id=orders.user_id").Joins("Join taxicompanies ON taxicompanies.id=orders.taxicomp_id").Joins("Join routes ON routes.id=orders.route_id").Where("routes.is_deleted=?", isdeleted).Where("orders.is_done=? AND orders.is_deleted=?", isdone, isdeleted).Where("orders.id=?  AND orders.user_id=?", oid, uid).Order("orders.id").Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeleted).Find(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrdersByID]error in getting all order by id %s\n", err.Error())
		return order, err
	}
	return order, nil
}

func CheckOrdersasDone(isdone bool, id int) error {
	var order models.Order
	err := db.GetconnectDB().Model(&order).Select("is_done").Where("id=?", id).Updates(models.Order{IsDone: isdone}).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckOrdersasDone]error in checked order %s\n", err.Error())

	}
	return nil
}
func OrdersReport(isdone, isdeleted, isresp, isblocked bool) ([]models.Order, error) {
	var order []models.Order
	ordersreport := `SELECT t.comptitle, o.is_done, COUNT(DISTINCT СASE WHEN u.role='user'THEN u.id END) AS clientid, COUNT(DISTINCT СASE WHEN u.role='driver'THEN u.id END) AS driverid COUNT(DISTINCT o.route_id) AS routeid, SUM(r.distance) as distance, SUM(r.price) as price FROM orders o LEFT JOIN taxicompanies t ON t.id = o.taxicomp_id LEFT JOIN routes ON routes.id=o.route_id  LEFT JOIN users u ON u.ID = o.user_id WHERE o.is_done = ? AND o.is_deleted = ? AND t.is_deleted = ? AND r.is_deleted=? AND r.is_response=?  AND u.is_blocked = ? AND u.is_deleted = ? GROUP BY t.comptitle, o.is_done ORDER BY count_of_clients, count_of_drivers DESC`
	err := db.GetconnectDB().Raw(ordersreport, isdone, isdeleted, isresp, isblocked).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.OrdersReport]error in  order report %s\n", err.Error())
		return order, err
	}
	return order, nil
}
